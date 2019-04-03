package egressmapper

import (
	"context"
	"fmt"

	egressv1alpha1 "github.com/mkimuram/egress-mapper/pkg/apis/egress/v1alpha1"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_egressmapper")

// Add creates a new EgressMapper Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileEgressMapper{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("egressmapper-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource EgressMapper
	err = c.Watch(&source.Kind{Type: &egressv1alpha1.EgressMapper{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource DaemonSet and requeue the owner EgressMapper
	err = c.Watch(&source.Kind{Type: &appsv1.DaemonSet{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &egressv1alpha1.EgressMapper{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileEgressMapper{}

// ReconcileEgressMapper reconciles a EgressMapper object
type ReconcileEgressMapper struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a EgressMapper object and makes changes based on the state read
// and what is in the EgressMapper.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileEgressMapper) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling EgressMapper")

	// Fetch the EgressMapper instance
	instance := &egressv1alpha1.EgressMapper{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Sync keepalived-vip daemonset
	if err := syncKeepAlivedVip(r, instance, reqLogger); err != nil {
		return reconcile.Result{}, err
	}

	// Sync kube-egress daemonset
	if err := syncKubeEgress(r, instance, reqLogger); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func syncKeepAlivedVip(r *ReconcileEgressMapper, cr *egressv1alpha1.EgressMapper, reqLogger logr.Logger) error {
	// Define new keepalived-vip daemonset
	keepAlivedVipDS := newKeepAlivedVipDSForCR(cr)

	// Set EgressMapper instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, keepAlivedVipDS, r.scheme); err != nil {
		return err
	}

	// Check if keepalived-vip DaemonSet already exists
	found := &appsv1.DaemonSet{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: keepAlivedVipDS.Name, Namespace: keepAlivedVipDS.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new keepalived-vip daemonset")
		err = r.client.Create(context.TODO(), keepAlivedVipDS)
		if err != nil {
			reqLogger.Info("Creating a new keepalived-vip daemonset fails")
			return err
		}

		// keepalived-vip daemonSet created successfully - don't requeue
		reqLogger.Info("Creating a new keepalived-vip daemonset succeeded")
		return nil
	} else if err != nil {
		reqLogger.Info("Getting keepalived-vip daemonset fails")
		return err
	}

	// keepalived-vip daemonSet already exists - don't requeue
	reqLogger.Info("Skip reconcile: keepalived-vip daemonSet. Daemonset already exists.")
	return nil
}

func syncKubeEgress(r *ReconcileEgressMapper, cr *egressv1alpha1.EgressMapper, reqLogger logr.Logger) error {
	// Define new kube-egress daemonset
	kubeEgressDS := newKubeEgressDSForCR(cr)

	// Set EgressMapper instance as the owner and controller
	if err := controllerutil.SetControllerReference(cr, kubeEgressDS, r.scheme); err != nil {
		return err
	}

	// Check if kube-egress DaemonSet already exists
	found := &appsv1.DaemonSet{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: kubeEgressDS.Name, Namespace: kubeEgressDS.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new kube-egress daemonset")
		err = r.client.Create(context.TODO(), kubeEgressDS)
		if err != nil {
			reqLogger.Info("Creating a new kube-egress daemonset fails")
			return err
		}

		// kube-egress daemonSet created successfully - don't requeue
		reqLogger.Info("Creating a new kube-egress daemonset succeeded")
		return nil
	} else if err != nil {
		reqLogger.Info("Getting kube-egress daemonset fails")
		return err
	}

	// keepalived-vip daemonSet already exists - don't requeue
	reqLogger.Info("Skip reconcile: kube-egress daemonSet. Daemonset already exists.")
	return nil
}

// newKeepAlivedVipDSForCR returns a keepalived daemonset with the same name/namespace as the cr
func newKeepAlivedVipDSForCR(cr *egressv1alpha1.EgressMapper) *appsv1.DaemonSet {
	namespace := "default"
	dsName := "kube-keepalived-vip"
	saName := "kube-keepalived-vip"
	containerName := "kube-keepalived-vip"
	image := cr.Spec.KeepalivedVIPImage

	imagePullPolicy := corev1.PullIfNotPresent
	configmapName := "vip-configmap"

	configmapArg := fmt.Sprintf("--services-configmap=%s/%s", namespace, configmapName)
	isPrivileged := true
	labels := map[string]string{"name": "kube-keepalived-vip"}

	volumes := []corev1.Volume{
		{
			Name: "modules",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/lib/modules"},
			},
		},
		{
			Name: "dev",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/dev"},
			},
		},
	}

	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "modules",
			MountPath: "/lib/modules",
			ReadOnly:  true,
		},
		{
			Name:      "dev",
			MountPath: "/dev",
		},
	}

	env := []corev1.EnvVar{
		{
			Name: "POD_NAME",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			},
		},
		{
			Name: "POD_NAMESPACE",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.namespace",
				},
			},
		},
	}

	container := corev1.Container{
		Name:                     containerName,
		Image:                    image,
		ImagePullPolicy:          imagePullPolicy,
		TerminationMessagePolicy: corev1.TerminationMessageReadFile,
		SecurityContext:          &corev1.SecurityContext{Privileged: &isPrivileged},
		VolumeMounts:             volumeMounts,
		Env:                      env,
		Args:                     []string{configmapArg},
	}

	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dsName,
			Namespace: namespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels:      labels,
				MatchExpressions: []metav1.LabelSelectorRequirement{},
			},
			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
				Type: appsv1.OnDeleteDaemonSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					HostNetwork:        true,
					ServiceAccountName: saName,
					RestartPolicy:      corev1.RestartPolicyAlways,
					Containers:         []corev1.Container{container},
					Volumes:            volumes,
				},
			},
		},
	}
}

// newKubeEgressDSForCR returns a kube-egress daemonset with the same name/namespace as the cr
func newKubeEgressDSForCR(cr *egressv1alpha1.EgressMapper) *appsv1.DaemonSet {
	namespace := "default"
	dsName := "kube-egress"
	containerName := "kube-egress"
	image := cr.Spec.KubeEgressImage
	imagePullPolicy := corev1.PullIfNotPresent
	isPrivileged := true
	labels := map[string]string{"app": "kube-egress"}
	directoryOrCreate := corev1.HostPathDirectoryOrCreate
	fileOrCreate := corev1.HostPathFileOrCreate
	vipRouteidMappings := "/etc/vip-routeid-mappings"
	podipVipMappings := "/etc/podip-vip-mappings"
	var terminationGracePeriodSeconds int64 = 10

	// Update args with user specified values
	args := getArgsForKubeEgressDS(cr)

	volumes := []corev1.Volume{
		{
			Name: "routing-tables",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/iproute2/rt_tables.d/",
					Type: &directoryOrCreate},
			},
		},
		{
			Name: "xtables-lock",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/run/xtables.lock",
					Type: &fileOrCreate},
			},
		},
		{
			Name: "vip-routeid-mappings",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "vip-routeid-mappings",
					},
				},
			},
		},
		{
			Name: "podip-vip-mappings",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "podip-vip-mappings",
					},
				},
			},
		},
	}

	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "routing-tables",
			MountPath: "/etc/iproute2/rt_tables.d/",
			ReadOnly:  false,
		},
		{
			Name:      "xtables-lock",
			MountPath: "/run/xtables.lock",
			ReadOnly:  false,
		},
		{
			Name:      "vip-routeid-mappings",
			MountPath: vipRouteidMappings,
		},
		{
			Name:      "podip-vip-mappings",
			MountPath: podipVipMappings,
		},
	}

	container := corev1.Container{
		Name:            containerName,
		Image:           image,
		ImagePullPolicy: imagePullPolicy,
		SecurityContext: &corev1.SecurityContext{Privileged: &isPrivileged},
		VolumeMounts:    volumeMounts,
		Args:            args,
	}

	// TODO: add some more values, like rollingUpdate policy and resource limit, from original spec.
	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dsName,
			Namespace: namespace,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels:      labels,
				MatchExpressions: []metav1.LabelSelectorRequirement{},
			},
			UpdateStrategy: appsv1.DaemonSetUpdateStrategy{
				Type: appsv1.RollingUpdateDaemonSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					HostNetwork:                   true,
					TerminationGracePeriodSeconds: &terminationGracePeriodSeconds,
					Containers:                    []corev1.Container{container},
					Volumes:                       volumes,
				},
			},
		},
	}
}

func getArgsForKubeEgressDS(cr *egressv1alpha1.EgressMapper) []string {
	args := []string{}

	if cr.Spec.PodSubnet != "" {
		args = append(args, fmt.Sprintf("--pod-subnet=%s", cr.Spec.PodSubnet))
	}
	if cr.Spec.ServiceSubnet != "" {
		args = append(args, fmt.Sprintf("--service-subnet=%s", cr.Spec.ServiceSubnet))
	}
	if cr.Spec.InterfaceName != "" {
		args = append(args, fmt.Sprintf("--interface=%s", cr.Spec.InterfaceName))
	}
	if cr.Spec.UpdateInterval != "" {
		args = append(args, fmt.Sprintf("--update-interval=%s", cr.Spec.UpdateInterval))
	}

	return args
}
