package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EgressMapperSpec defines the desired state of EgressMapper
// +k8s:openapi-gen=true
type EgressMapperSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	KeepalivedVIPImage string `json:"KeepalivedVIPImage,omitempty"`
	NeedMountDev       bool   `json:"NeedMountDev,omitempty"`
	KubeEgressImage    string `json:"KubeEgressImage,omitempty"`
	PodSubnet          string `json:"PodSubnet,omitempty"`
	ServiceSubnet      string `json:"ServiceSubnet,omitempty"`
	InterfaceName      string `json:"InterfaceName,omitempty"`
	UpdateInterval     string `json:"UpdateInterval,omitempty"`
}

// EgressMapperStatus defines the observed state of EgressMapper
// +k8s:openapi-gen=true
type EgressMapperStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EgressMapper is the Schema for the egressmappers API
// +k8s:openapi-gen=true
type EgressMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EgressMapperSpec   `json:"spec,omitempty"`
	Status EgressMapperStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EgressMapperList contains a list of EgressMapper
type EgressMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EgressMapper `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EgressMapper{}, &EgressMapperList{})
}
