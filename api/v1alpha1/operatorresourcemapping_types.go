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

type SourceResourceSpec struct {
	// A string represent the type of the controller
	kind string `json:"kind,omitempty"`
	// Array of componentNames controlled by the ORM instance
	componentNames []string `json:"componentNames,omitempty"`
}

type ResourceMappingTemplate struct {
	srcPath string `json:"srcPath,omitempty"`
	destPath string `json:"destPath,omitempty"`
}

type ResourceMapping struct {
	// A string represent the type of the controller
	srcResourceSpec SourceResourceSpec `json:"srcResourceSpec,omitempty"`
	// Array of componentNames controlled by the ORM instance
	resourceMappingTemplates []ResourceMappingTemplate `json:"resourceMappingTemplates,omitempty"`
}

// OperatorResourceMappingSpec defines the desired state of OperatorResourceMapping
type OperatorResourceMappingSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// A string represent the type of the controller
	resourceMappings []ResourceMapping `json:"resourceMappings,omitempty"`
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
