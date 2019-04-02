package utils

import (
	"context"
	"fmt"

	egressv1alpha1 "github.com/mkimuram/egress-mapper/pkg/apis/egress/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("utils")

// SyncConfigMapForEgress syncs configMap for egress
func SyncConfigMapForEgress(cl client.Client) error {
	// Get podip-vip-mappings
	egressConfigMap, err := GetOrCreateConfigMap(cl, "podip-vip-mappings", "default")
	if err != nil {
		return err
	}

	// Get list of egress
	egressList, err := getEgressList(cl)
	if err != nil {
		return err
	}

	// Update podip-vip-mappings
	if err := updateEgressConfigMap(cl, egressList, egressConfigMap); err != nil {
		return err
	}

	return nil
}

func getEgressList(cl client.Client) (*egressv1alpha1.EgressList, error) {
	// Get the list of egresses
	egressList := &egressv1alpha1.EgressList{}
	opts := &client.ListOptions{}
	err := cl.List(context.TODO(), opts, egressList)
	if err != nil {
		return egressList, fmt.Errorf("getEgressList fails in getting list of egresses: %v", err)
	}
	return egressList, nil
}

func updateEgressConfigMap(cl client.Client, egressList *egressv1alpha1.EgressList, egressConfigMap *corev1.ConfigMap) error {
	// Create the latest data from list of egresses and update configmap
	egressData, err := createEgressData(cl, egressList)
	if err != nil {
		return err
	}

	if err := UpdateConfigmapData(cl, egressConfigMap, egressData); err != nil {
		return err
	}

	return nil
}

func createEgressData(cl client.Client, egressList *egressv1alpha1.EgressList) (map[string]string, error) {
	reqLogger := log.WithValues("egressList", egressList)
	reqLogger.Info("createEgressData")

	// Create the latest data from list of egresses
	egressData := map[string]string{}
	for _, e := range egressList.Items {
		podIPs, err := getPodIPs(cl, e.Spec.Kind, e.Spec.Namespace, e.Spec.Name)
		if err != nil {
			reqLogger.Info(fmt.Sprintf("Skip adding mapping for resource(kind: %s, namespae: %s, name: %s), error: %v",
				e.Spec.Kind, e.Spec.Namespace, e.Spec.Name, err))
			continue
		}
		for _, podIP := range podIPs {
			if podIP == "" {
				reqLogger.Info(fmt.Sprintf("Skip adding mapping, because PodIP is empty for pod (kind: %s, namespace: %s, name: %s", e.Spec.Kind, e.Spec.Namespace, e.Spec.Name))
			} else {
				egressData[podIP] = e.Spec.IP
			}
		}
	}

	return egressData, nil
}
