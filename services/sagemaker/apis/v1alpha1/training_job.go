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

// TrainingJobSpec defines the desired state of TrainingJob
type TrainingJobSpec struct {
	// +kubebuilder:validation:Required
	AlgorithmSpecification                *AlgorithmSpecification   `json:"algorithmSpecification"`
	CheckpointConfig                      *CheckpointConfig         `json:"checkpointConfig,omitempty"`
	DebugHookConfig                       *DebugHookConfig          `json:"debugHookConfig,omitempty"`
	DebugRuleConfigurations               []*DebugRuleConfiguration `json:"debugRuleConfigurations,omitempty"`
	EnableInterContainerTrafficEncryption *bool                     `json:"enableInterContainerTrafficEncryption,omitempty"`
	EnableManagedSpotTraining             *bool                     `json:"enableManagedSpotTraining,omitempty"`
	EnableNetworkIsolation                *bool                     `json:"enableNetworkIsolation,omitempty"`
	ExperimentConfig                      *ExperimentConfig         `json:"experimentConfig,omitempty"`
	HyperParameters                       map[string]*string        `json:"hyperParameters,omitempty"`
	InputDataConfig                       []*Channel                `json:"inputDataConfig,omitempty"`
	// +kubebuilder:validation:Required
	OutputDataConfig *OutputDataConfig `json:"outputDataConfig"`
	// +kubebuilder:validation:Required
	ResourceConfig *ResourceConfig `json:"resourceConfig"`
	// +kubebuilder:validation:Required
	RoleARN *string `json:"roleARN"`
	// +kubebuilder:validation:Required
	StoppingCondition       *StoppingCondition       `json:"stoppingCondition"`
	Tags                    []*Tag                   `json:"tags,omitempty"`
	TensorBoardOutputConfig *TensorBoardOutputConfig `json:"tensorBoardOutputConfig,omitempty"`
	// +kubebuilder:validation:Required
	TrainingJobName *string    `json:"trainingJobName"`
	VPCConfig       *VPCConfig `json:"vpcConfig,omitempty"`
}

// TrainingJobStatus defines the observed state of TrainingJob
type TrainingJobStatus struct {
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

// TrainingJob is the Schema for the TrainingJobs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type TrainingJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrainingJobSpec   `json:"spec,omitempty"`
	Status            TrainingJobStatus `json:"status,omitempty"`
}

// TrainingJobList contains a list of TrainingJob
// +kubebuilder:object:root=true
type TrainingJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrainingJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrainingJob{}, &TrainingJobList{})
}
