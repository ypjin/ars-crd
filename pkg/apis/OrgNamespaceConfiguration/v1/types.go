package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrgNamespaceConfiguration describes a OrgNamespaceConfiguration resource
type OrgNamespaceConfiguration struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec OrgNamespaceConfigurationSpec `json:"spec"`
}

// OrgNamespaceConfigurationSpec is the spec for a OrgNamespaceConfiguration resource
type OrgNamespaceConfigurationSpec struct {
	// this is where you would put your custom resource data
	Config string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrgNamespaceConfigurationList is a list of OrgNamespaceConfiguration resources
type OrgNamespaceConfigurationList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []OrgNamespaceConfiguration `json:"items"`
}
