package utils

import (
	"context"
	"fmt"
	"strconv"

	egressv1alpha1 "github.com/mkimuram/egress-mapper/pkg/apis/egress/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SyncConfigMapForVip syncs configMap for vip
func SyncConfigMapForVip(cl client.Client) error {
	// Get vip-configmap
	vipConfigMap, err := GetOrCreateConfigMap(cl, "vip-configmap", "default")
	if err != nil {
		return err
	}

	// Get vip-routeid-mappings
	vipRouteIDMap, err := GetOrCreateConfigMap(cl, "vip-routeid-mappings", "default")
	if err != nil {
		return err
	}

	// Get list of vip
	vipList, err := getVipList(cl)
	if err != nil {
		return err
	}

	// Update vip-configmap
	if err := updateVipConfigMap(cl, vipList, vipConfigMap); err != nil {
		return err
	}

	// Update vip-routeid-mappings
	if err := updateVipRouteIDMap(cl, vipList, vipRouteIDMap); err != nil {
		return err
	}

	return nil
}

func getVipList(cl client.Client) (*egressv1alpha1.VipList, error) {
	// Get the list of vips
	vipList := &egressv1alpha1.VipList{}
	opts := &client.ListOptions{}
	err := cl.List(context.TODO(), opts, vipList)
	if err != nil {
		return vipList, fmt.Errorf("getVipLIst fails in getting list of vips: %v", err)
	}
	return vipList, nil
}

func updateVipConfigMap(cl client.Client, vipList *egressv1alpha1.VipList, vipConfigMap *corev1.ConfigMap) error {
	// Create the latest data from list of vips and update configmap
	vipData, err := createVipData(vipList)
	if err != nil {
		return err
	}

	if err := UpdateConfigmapData(cl, vipConfigMap, vipData); err != nil {
		return err
	}

	return nil
}

func updateVipRouteIDMap(cl client.Client, vipList *egressv1alpha1.VipList, vipRouteIDMap *corev1.ConfigMap) error {
	// Create the latest data from list of vips and update configmap
	vipRouteIDData, err := createVipRoutIDData(vipList)
	if err != nil {
		return err
	}

	if err := UpdateConfigmapData(cl, vipRouteIDMap, vipRouteIDData); err != nil {
		return err
	}

	return nil
}

func createVipData(vipList *egressv1alpha1.VipList) (map[string]string, error) {
	// Create the latest data from list of vips
	vipData := map[string]string{}
	for _, v := range vipList.Items {
		vipData[v.Spec.IP] = ""
	}

	return vipData, nil
}

func createVipRoutIDData(vipList *egressv1alpha1.VipList) (map[string]string, error) {
	// Create the latest data from list of vips
	vipRouteIDData := map[string]string{}
	for _, v := range vipList.Items {
		vipRouteIDData[v.Spec.IP] = strconv.Itoa(v.Spec.RouteID)
	}

	return vipRouteIDData, nil
}
