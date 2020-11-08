package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network describes a Network resource
type Network struct {
	// TypeMeta is the metadata for the resource, like kind and
	apiversion metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object,
	// including
	// things like...
	// - name
	// - namespace
	// - self link
	// - labels
	// - ... etc ...
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              networkspec `json:"spec"`
}

type networkspec struct {
	Cidr    string `json:"cidr"`
	Gateway string `json:"gateway"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList is a list of Network resources
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Network `json:"items"`
}
