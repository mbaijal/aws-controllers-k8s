/*
Copyright 2020 The Crossplane Authors.

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
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	Containers             []*ContainerDefinition `json:"containers,omitempty"`
	EnableNetworkIsolation *bool                  `json:"enableNetworkIsolation,omitempty"`
	// +kubebuilder:validation:Required
	ExecutionRoleARN *string `json:"executionRoleARN"`
	// +kubebuilder:validation:Required
	ModelName        *string              `json:"modelName"`
	PrimaryContainer *ContainerDefinition `json:"primaryContainer,omitempty"`
	Tags             []*Tag               `json:"tags,omitempty"`
	VPCConfig        *VPCConfig           `json:"vpcConfig,omitempty"`
}

// ModelStatus defines the observed state of Model
type ModelStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// Model is the Schema for the Models API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelSpec   `json:"spec,omitempty"`
	Status            ModelStatus `json:"status,omitempty"`
}

// ModelList contains a list of Model
// +kubebuilder:object:root=true
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}
