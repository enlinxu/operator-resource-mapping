/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SourceResourceSpec defines the kind and components name controlled by the ORM instance
type SourceResourceSpec struct {
	// A string represent the type of the controller
	Kind string `json:"kind,omitempty"`
	// Array of componentNames controlled by the ORM instance
	ComponentNames []string `json:"componentNames,omitempty"`
}

// ResourceMappingTemplate defines the mapping between source path and destination path
type ResourceMappingTemplate struct {
	// A string represents the source path in the deployment
	SrcPath string `json:"srcPath,omitempty"`
	// A string represents the destination path in the CR (CustomResource)
	DestPath string `json:"destPath,omitempty"`
}

// ResourceMapping combines the SourceResourceSpec and the ResourceMappingTemplate
type ResourceMapping struct {
	// srcResourceSpec defines the kind and components name controlled by the ORM instance
	SrcResourceSpec SourceResourceSpec `json:"srcResourceSpec,omitempty"`
	// resourceMappingTemplates defines an array of the mapping between source path and destination path
	ResourceMappingTemplates []ResourceMappingTemplate `json:"resourceMappingTemplates,omitempty"`
}

// OperatorResourceMappingSpec defines the desired state of OperatorResourceMapping
type OperatorResourceMappingSpec struct {
	// An array of ResourceMapping definition
	ResourceMappings []ResourceMapping `json:"resourceMappings,omitempty"`
}

// OperatorResourceMappingStatus defines the observed state of OperatorResourceMapping
type OperatorResourceMappingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OperatorResourceMapping is the Schema for the operatorresourcemappings API
type OperatorResourceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OperatorResourceMappingSpec   `json:"spec,omitempty"`
	Status OperatorResourceMappingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OperatorResourceMappingList contains a list of OperatorResourceMapping
type OperatorResourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OperatorResourceMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OperatorResourceMapping{}, &OperatorResourceMappingList{})
}
