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

package training_job

import (
	"context"
	corev1 "k8s.io/api/core/v1"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/sagemaker/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.TrainingJob{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeTrainingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeTrainingJob", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "ValidationException" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AlgorithmSpecification != nil {
		f0 := &svcapitypes.AlgorithmSpecification{}
		if resp.AlgorithmSpecification.AlgorithmName != nil {
			f0.AlgorithmName = resp.AlgorithmSpecification.AlgorithmName
		}
		if resp.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != nil {
			f0.EnableSageMakerMetricsTimeSeries = resp.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries
		}
		if resp.AlgorithmSpecification.MetricDefinitions != nil {
			f0f2 := []*svcapitypes.MetricDefinition{}
			for _, f0f2iter := range resp.AlgorithmSpecification.MetricDefinitions {
				f0f2elem := &svcapitypes.MetricDefinition{}
				if f0f2iter.Name != nil {
					f0f2elem.Name = f0f2iter.Name
				}
				if f0f2iter.Regex != nil {
					f0f2elem.Regex = f0f2iter.Regex
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.MetricDefinitions = f0f2
		}
		if resp.AlgorithmSpecification.TrainingImage != nil {
			f0.TrainingImage = resp.AlgorithmSpecification.TrainingImage
		}
		if resp.AlgorithmSpecification.TrainingInputMode != nil {
			f0.TrainingInputMode = resp.AlgorithmSpecification.TrainingInputMode
		}
		ko.Spec.AlgorithmSpecification = f0
	}
	if resp.CheckpointConfig != nil {
		f3 := &svcapitypes.CheckpointConfig{}
		if resp.CheckpointConfig.LocalPath != nil {
			f3.LocalPath = resp.CheckpointConfig.LocalPath
		}
		if resp.CheckpointConfig.S3Uri != nil {
			f3.S3URI = resp.CheckpointConfig.S3Uri
		}
		ko.Spec.CheckpointConfig = f3
	}
	if resp.DebugHookConfig != nil {
		f5 := &svcapitypes.DebugHookConfig{}
		if resp.DebugHookConfig.CollectionConfigurations != nil {
			f5f0 := []*svcapitypes.CollectionConfiguration{}
			for _, f5f0iter := range resp.DebugHookConfig.CollectionConfigurations {
				f5f0elem := &svcapitypes.CollectionConfiguration{}
				if f5f0iter.CollectionName != nil {
					f5f0elem.CollectionName = f5f0iter.CollectionName
				}
				if f5f0iter.CollectionParameters != nil {
					f5f0elemf1 := map[string]*string{}
					for f5f0elemf1key, f5f0elemf1valiter := range f5f0iter.CollectionParameters {
						var f5f0elemf1val string
						f5f0elemf1val = *f5f0elemf1valiter
						f5f0elemf1[f5f0elemf1key] = &f5f0elemf1val
					}
					f5f0elem.CollectionParameters = f5f0elemf1
				}
				f5f0 = append(f5f0, f5f0elem)
			}
			f5.CollectionConfigurations = f5f0
		}
		if resp.DebugHookConfig.HookParameters != nil {
			f5f1 := map[string]*string{}
			for f5f1key, f5f1valiter := range resp.DebugHookConfig.HookParameters {
				var f5f1val string
				f5f1val = *f5f1valiter
				f5f1[f5f1key] = &f5f1val
			}
			f5.HookParameters = f5f1
		}
		if resp.DebugHookConfig.LocalPath != nil {
			f5.LocalPath = resp.DebugHookConfig.LocalPath
		}
		if resp.DebugHookConfig.S3OutputPath != nil {
			f5.S3OutputPath = resp.DebugHookConfig.S3OutputPath
		}
		ko.Spec.DebugHookConfig = f5
	}
	if resp.DebugRuleConfigurations != nil {
		f6 := []*svcapitypes.DebugRuleConfiguration{}
		for _, f6iter := range resp.DebugRuleConfigurations {
			f6elem := &svcapitypes.DebugRuleConfiguration{}
			if f6iter.InstanceType != nil {
				f6elem.InstanceType = f6iter.InstanceType
			}
			if f6iter.LocalPath != nil {
				f6elem.LocalPath = f6iter.LocalPath
			}
			if f6iter.RuleConfigurationName != nil {
				f6elem.RuleConfigurationName = f6iter.RuleConfigurationName
			}
			if f6iter.RuleEvaluatorImage != nil {
				f6elem.RuleEvaluatorImage = f6iter.RuleEvaluatorImage
			}
			if f6iter.RuleParameters != nil {
				f6elemf4 := map[string]*string{}
				for f6elemf4key, f6elemf4valiter := range f6iter.RuleParameters {
					var f6elemf4val string
					f6elemf4val = *f6elemf4valiter
					f6elemf4[f6elemf4key] = &f6elemf4val
				}
				f6elem.RuleParameters = f6elemf4
			}
			if f6iter.S3OutputPath != nil {
				f6elem.S3OutputPath = f6iter.S3OutputPath
			}
			if f6iter.VolumeSizeInGB != nil {
				f6elem.VolumeSizeInGB = f6iter.VolumeSizeInGB
			}
			f6 = append(f6, f6elem)
		}
		ko.Spec.DebugRuleConfigurations = f6
	}
	if resp.EnableInterContainerTrafficEncryption != nil {
		ko.Spec.EnableInterContainerTrafficEncryption = resp.EnableInterContainerTrafficEncryption
	}
	if resp.EnableManagedSpotTraining != nil {
		ko.Spec.EnableManagedSpotTraining = resp.EnableManagedSpotTraining
	}
	if resp.EnableNetworkIsolation != nil {
		ko.Spec.EnableNetworkIsolation = resp.EnableNetworkIsolation
	}
	if resp.ExperimentConfig != nil {
		f11 := &svcapitypes.ExperimentConfig{}
		if resp.ExperimentConfig.ExperimentName != nil {
			f11.ExperimentName = resp.ExperimentConfig.ExperimentName
		}
		if resp.ExperimentConfig.TrialComponentDisplayName != nil {
			f11.TrialComponentDisplayName = resp.ExperimentConfig.TrialComponentDisplayName
		}
		if resp.ExperimentConfig.TrialName != nil {
			f11.TrialName = resp.ExperimentConfig.TrialName
		}
		ko.Spec.ExperimentConfig = f11
	}
	if resp.HyperParameters != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range resp.HyperParameters {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		ko.Spec.HyperParameters = f14
	}
	if resp.InputDataConfig != nil {
		f15 := []*svcapitypes.Channel{}
		for _, f15iter := range resp.InputDataConfig {
			f15elem := &svcapitypes.Channel{}
			if f15iter.ChannelName != nil {
				f15elem.ChannelName = f15iter.ChannelName
			}
			if f15iter.CompressionType != nil {
				f15elem.CompressionType = f15iter.CompressionType
			}
			if f15iter.ContentType != nil {
				f15elem.ContentType = f15iter.ContentType
			}
			if f15iter.DataSource != nil {
				f15elemf3 := &svcapitypes.DataSource{}
				if f15iter.DataSource.FileSystemDataSource != nil {
					f15elemf3f0 := &svcapitypes.FileSystemDataSource{}
					if f15iter.DataSource.FileSystemDataSource.DirectoryPath != nil {
						f15elemf3f0.DirectoryPath = f15iter.DataSource.FileSystemDataSource.DirectoryPath
					}
					if f15iter.DataSource.FileSystemDataSource.FileSystemAccessMode != nil {
						f15elemf3f0.FileSystemAccessMode = f15iter.DataSource.FileSystemDataSource.FileSystemAccessMode
					}
					if f15iter.DataSource.FileSystemDataSource.FileSystemId != nil {
						f15elemf3f0.FileSystemID = f15iter.DataSource.FileSystemDataSource.FileSystemId
					}
					if f15iter.DataSource.FileSystemDataSource.FileSystemType != nil {
						f15elemf3f0.FileSystemType = f15iter.DataSource.FileSystemDataSource.FileSystemType
					}
					f15elemf3.FileSystemDataSource = f15elemf3f0
				}
				if f15iter.DataSource.S3DataSource != nil {
					f15elemf3f1 := &svcapitypes.S3DataSource{}
					if f15iter.DataSource.S3DataSource.AttributeNames != nil {
						f15elemf3f1f0 := []*string{}
						for _, f15elemf3f1f0iter := range f15iter.DataSource.S3DataSource.AttributeNames {
							var f15elemf3f1f0elem string
							f15elemf3f1f0elem = *f15elemf3f1f0iter
							f15elemf3f1f0 = append(f15elemf3f1f0, &f15elemf3f1f0elem)
						}
						f15elemf3f1.AttributeNames = f15elemf3f1f0
					}
					if f15iter.DataSource.S3DataSource.S3DataDistributionType != nil {
						f15elemf3f1.S3DataDistributionType = f15iter.DataSource.S3DataSource.S3DataDistributionType
					}
					if f15iter.DataSource.S3DataSource.S3DataType != nil {
						f15elemf3f1.S3DataType = f15iter.DataSource.S3DataSource.S3DataType
					}
					if f15iter.DataSource.S3DataSource.S3Uri != nil {
						f15elemf3f1.S3URI = f15iter.DataSource.S3DataSource.S3Uri
					}
					f15elemf3.S3DataSource = f15elemf3f1
				}
				f15elem.DataSource = f15elemf3
			}
			if f15iter.InputMode != nil {
				f15elem.InputMode = f15iter.InputMode
			}
			if f15iter.RecordWrapperType != nil {
				f15elem.RecordWrapperType = f15iter.RecordWrapperType
			}
			if f15iter.ShuffleConfig != nil {
				f15elemf6 := &svcapitypes.ShuffleConfig{}
				if f15iter.ShuffleConfig.Seed != nil {
					f15elemf6.Seed = f15iter.ShuffleConfig.Seed
				}
				f15elem.ShuffleConfig = f15elemf6
			}
			f15 = append(f15, f15elem)
		}
		ko.Spec.InputDataConfig = f15
	}
	if resp.OutputDataConfig != nil {
		f19 := &svcapitypes.OutputDataConfig{}
		if resp.OutputDataConfig.KmsKeyId != nil {
			f19.KMSKeyID = resp.OutputDataConfig.KmsKeyId
		}
		if resp.OutputDataConfig.S3OutputPath != nil {
			f19.S3OutputPath = resp.OutputDataConfig.S3OutputPath
		}
		ko.Spec.OutputDataConfig = f19
	}
	if resp.ResourceConfig != nil {
		f20 := &svcapitypes.ResourceConfig{}
		if resp.ResourceConfig.InstanceCount != nil {
			f20.InstanceCount = resp.ResourceConfig.InstanceCount
		}
		if resp.ResourceConfig.InstanceType != nil {
			f20.InstanceType = resp.ResourceConfig.InstanceType
		}
		if resp.ResourceConfig.VolumeKmsKeyId != nil {
			f20.VolumeKMSKeyID = resp.ResourceConfig.VolumeKmsKeyId
		}
		if resp.ResourceConfig.VolumeSizeInGB != nil {
			f20.VolumeSizeInGB = resp.ResourceConfig.VolumeSizeInGB
		}
		ko.Spec.ResourceConfig = f20
	}
	if resp.RoleArn != nil {
		ko.Spec.RoleARN = resp.RoleArn
	}
	if resp.StoppingCondition != nil {
		f24 := &svcapitypes.StoppingCondition{}
		if resp.StoppingCondition.MaxRuntimeInSeconds != nil {
			f24.MaxRuntimeInSeconds = resp.StoppingCondition.MaxRuntimeInSeconds
		}
		if resp.StoppingCondition.MaxWaitTimeInSeconds != nil {
			f24.MaxWaitTimeInSeconds = resp.StoppingCondition.MaxWaitTimeInSeconds
		}
		ko.Spec.StoppingCondition = f24
	}
	if resp.TensorBoardOutputConfig != nil {
		f25 := &svcapitypes.TensorBoardOutputConfig{}
		if resp.TensorBoardOutputConfig.LocalPath != nil {
			f25.LocalPath = resp.TensorBoardOutputConfig.LocalPath
		}
		if resp.TensorBoardOutputConfig.S3OutputPath != nil {
			f25.S3OutputPath = resp.TensorBoardOutputConfig.S3OutputPath
		}
		ko.Spec.TensorBoardOutputConfig = f25
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TrainingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TrainingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.TrainingJobName != nil {
		ko.Spec.TrainingJobName = resp.TrainingJobName
	}
	if resp.VpcConfig != nil {
		f33 := &svcapitypes.VPCConfig{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f33f0 := []*string{}
			for _, f33f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f33f0elem string
				f33f0elem = *f33f0iter
				f33f0 = append(f33f0, &f33f0elem)
			}
			f33.SecurityGroupIDs = f33f0
		}
		if resp.VpcConfig.Subnets != nil {
			f33f1 := []*string{}
			for _, f33f1iter := range resp.VpcConfig.Subnets {
				var f33f1elem string
				f33f1elem = *f33f1iter
				f33f1 = append(f33f1, &f33f1elem)
			}
			f33.Subnets = f33f1
		}
		ko.Spec.VPCConfig = f33
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.TrainingJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeTrainingJobInput, error) {
	res := &svcsdk.DescribeTrainingJobInput{}

	if r.ko.Spec.TrainingJobName != nil {
		res.SetTrainingJobName(*r.ko.Spec.TrainingJobName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateTrainingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTrainingJob", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TrainingJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TrainingJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateTrainingJobInput, error) {
	res := &svcsdk.CreateTrainingJobInput{}

	if r.ko.Spec.AlgorithmSpecification != nil {
		f0 := &svcsdk.AlgorithmSpecification{}
		if r.ko.Spec.AlgorithmSpecification.AlgorithmName != nil {
			f0.SetAlgorithmName(*r.ko.Spec.AlgorithmSpecification.AlgorithmName)
		}
		if r.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries != nil {
			f0.SetEnableSageMakerMetricsTimeSeries(*r.ko.Spec.AlgorithmSpecification.EnableSageMakerMetricsTimeSeries)
		}
		if r.ko.Spec.AlgorithmSpecification.MetricDefinitions != nil {
			f0f2 := []*svcsdk.MetricDefinition{}
			for _, f0f2iter := range r.ko.Spec.AlgorithmSpecification.MetricDefinitions {
				f0f2elem := &svcsdk.MetricDefinition{}
				if f0f2iter.Name != nil {
					f0f2elem.SetName(*f0f2iter.Name)
				}
				if f0f2iter.Regex != nil {
					f0f2elem.SetRegex(*f0f2iter.Regex)
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.SetMetricDefinitions(f0f2)
		}
		if r.ko.Spec.AlgorithmSpecification.TrainingImage != nil {
			f0.SetTrainingImage(*r.ko.Spec.AlgorithmSpecification.TrainingImage)
		}
		if r.ko.Spec.AlgorithmSpecification.TrainingInputMode != nil {
			f0.SetTrainingInputMode(*r.ko.Spec.AlgorithmSpecification.TrainingInputMode)
		}
		res.SetAlgorithmSpecification(f0)
	}
	if r.ko.Spec.CheckpointConfig != nil {
		f1 := &svcsdk.CheckpointConfig{}
		if r.ko.Spec.CheckpointConfig.LocalPath != nil {
			f1.SetLocalPath(*r.ko.Spec.CheckpointConfig.LocalPath)
		}
		if r.ko.Spec.CheckpointConfig.S3URI != nil {
			f1.SetS3Uri(*r.ko.Spec.CheckpointConfig.S3URI)
		}
		res.SetCheckpointConfig(f1)
	}
	if r.ko.Spec.DebugHookConfig != nil {
		f2 := &svcsdk.DebugHookConfig{}
		if r.ko.Spec.DebugHookConfig.CollectionConfigurations != nil {
			f2f0 := []*svcsdk.CollectionConfiguration{}
			for _, f2f0iter := range r.ko.Spec.DebugHookConfig.CollectionConfigurations {
				f2f0elem := &svcsdk.CollectionConfiguration{}
				if f2f0iter.CollectionName != nil {
					f2f0elem.SetCollectionName(*f2f0iter.CollectionName)
				}
				if f2f0iter.CollectionParameters != nil {
					f2f0elemf1 := map[string]*string{}
					for f2f0elemf1key, f2f0elemf1valiter := range f2f0iter.CollectionParameters {
						var f2f0elemf1val string
						f2f0elemf1val = *f2f0elemf1valiter
						f2f0elemf1[f2f0elemf1key] = &f2f0elemf1val
					}
					f2f0elem.SetCollectionParameters(f2f0elemf1)
				}
				f2f0 = append(f2f0, f2f0elem)
			}
			f2.SetCollectionConfigurations(f2f0)
		}
		if r.ko.Spec.DebugHookConfig.HookParameters != nil {
			f2f1 := map[string]*string{}
			for f2f1key, f2f1valiter := range r.ko.Spec.DebugHookConfig.HookParameters {
				var f2f1val string
				f2f1val = *f2f1valiter
				f2f1[f2f1key] = &f2f1val
			}
			f2.SetHookParameters(f2f1)
		}
		if r.ko.Spec.DebugHookConfig.LocalPath != nil {
			f2.SetLocalPath(*r.ko.Spec.DebugHookConfig.LocalPath)
		}
		if r.ko.Spec.DebugHookConfig.S3OutputPath != nil {
			f2.SetS3OutputPath(*r.ko.Spec.DebugHookConfig.S3OutputPath)
		}
		res.SetDebugHookConfig(f2)
	}
	if r.ko.Spec.DebugRuleConfigurations != nil {
		f3 := []*svcsdk.DebugRuleConfiguration{}
		for _, f3iter := range r.ko.Spec.DebugRuleConfigurations {
			f3elem := &svcsdk.DebugRuleConfiguration{}
			if f3iter.InstanceType != nil {
				f3elem.SetInstanceType(*f3iter.InstanceType)
			}
			if f3iter.LocalPath != nil {
				f3elem.SetLocalPath(*f3iter.LocalPath)
			}
			if f3iter.RuleConfigurationName != nil {
				f3elem.SetRuleConfigurationName(*f3iter.RuleConfigurationName)
			}
			if f3iter.RuleEvaluatorImage != nil {
				f3elem.SetRuleEvaluatorImage(*f3iter.RuleEvaluatorImage)
			}
			if f3iter.RuleParameters != nil {
				f3elemf4 := map[string]*string{}
				for f3elemf4key, f3elemf4valiter := range f3iter.RuleParameters {
					var f3elemf4val string
					f3elemf4val = *f3elemf4valiter
					f3elemf4[f3elemf4key] = &f3elemf4val
				}
				f3elem.SetRuleParameters(f3elemf4)
			}
			if f3iter.S3OutputPath != nil {
				f3elem.SetS3OutputPath(*f3iter.S3OutputPath)
			}
			if f3iter.VolumeSizeInGB != nil {
				f3elem.SetVolumeSizeInGB(*f3iter.VolumeSizeInGB)
			}
			f3 = append(f3, f3elem)
		}
		res.SetDebugRuleConfigurations(f3)
	}
	if r.ko.Spec.EnableInterContainerTrafficEncryption != nil {
		res.SetEnableInterContainerTrafficEncryption(*r.ko.Spec.EnableInterContainerTrafficEncryption)
	}
	if r.ko.Spec.EnableManagedSpotTraining != nil {
		res.SetEnableManagedSpotTraining(*r.ko.Spec.EnableManagedSpotTraining)
	}
	if r.ko.Spec.EnableNetworkIsolation != nil {
		res.SetEnableNetworkIsolation(*r.ko.Spec.EnableNetworkIsolation)
	}
	if r.ko.Spec.ExperimentConfig != nil {
		f7 := &svcsdk.ExperimentConfig{}
		if r.ko.Spec.ExperimentConfig.ExperimentName != nil {
			f7.SetExperimentName(*r.ko.Spec.ExperimentConfig.ExperimentName)
		}
		if r.ko.Spec.ExperimentConfig.TrialComponentDisplayName != nil {
			f7.SetTrialComponentDisplayName(*r.ko.Spec.ExperimentConfig.TrialComponentDisplayName)
		}
		if r.ko.Spec.ExperimentConfig.TrialName != nil {
			f7.SetTrialName(*r.ko.Spec.ExperimentConfig.TrialName)
		}
		res.SetExperimentConfig(f7)
	}
	if r.ko.Spec.HyperParameters != nil {
		f8 := map[string]*string{}
		for f8key, f8valiter := range r.ko.Spec.HyperParameters {
			var f8val string
			f8val = *f8valiter
			f8[f8key] = &f8val
		}
		res.SetHyperParameters(f8)
	}
	if r.ko.Spec.InputDataConfig != nil {
		f9 := []*svcsdk.Channel{}
		for _, f9iter := range r.ko.Spec.InputDataConfig {
			f9elem := &svcsdk.Channel{}
			if f9iter.ChannelName != nil {
				f9elem.SetChannelName(*f9iter.ChannelName)
			}
			if f9iter.CompressionType != nil {
				f9elem.SetCompressionType(*f9iter.CompressionType)
			}
			if f9iter.ContentType != nil {
				f9elem.SetContentType(*f9iter.ContentType)
			}
			if f9iter.DataSource != nil {
				f9elemf3 := &svcsdk.DataSource{}
				if f9iter.DataSource.FileSystemDataSource != nil {
					f9elemf3f0 := &svcsdk.FileSystemDataSource{}
					if f9iter.DataSource.FileSystemDataSource.DirectoryPath != nil {
						f9elemf3f0.SetDirectoryPath(*f9iter.DataSource.FileSystemDataSource.DirectoryPath)
					}
					if f9iter.DataSource.FileSystemDataSource.FileSystemAccessMode != nil {
						f9elemf3f0.SetFileSystemAccessMode(*f9iter.DataSource.FileSystemDataSource.FileSystemAccessMode)
					}
					if f9iter.DataSource.FileSystemDataSource.FileSystemID != nil {
						f9elemf3f0.SetFileSystemId(*f9iter.DataSource.FileSystemDataSource.FileSystemID)
					}
					if f9iter.DataSource.FileSystemDataSource.FileSystemType != nil {
						f9elemf3f0.SetFileSystemType(*f9iter.DataSource.FileSystemDataSource.FileSystemType)
					}
					f9elemf3.SetFileSystemDataSource(f9elemf3f0)
				}
				if f9iter.DataSource.S3DataSource != nil {
					f9elemf3f1 := &svcsdk.S3DataSource{}
					if f9iter.DataSource.S3DataSource.AttributeNames != nil {
						f9elemf3f1f0 := []*string{}
						for _, f9elemf3f1f0iter := range f9iter.DataSource.S3DataSource.AttributeNames {
							var f9elemf3f1f0elem string
							f9elemf3f1f0elem = *f9elemf3f1f0iter
							f9elemf3f1f0 = append(f9elemf3f1f0, &f9elemf3f1f0elem)
						}
						f9elemf3f1.SetAttributeNames(f9elemf3f1f0)
					}
					if f9iter.DataSource.S3DataSource.S3DataDistributionType != nil {
						f9elemf3f1.SetS3DataDistributionType(*f9iter.DataSource.S3DataSource.S3DataDistributionType)
					}
					if f9iter.DataSource.S3DataSource.S3DataType != nil {
						f9elemf3f1.SetS3DataType(*f9iter.DataSource.S3DataSource.S3DataType)
					}
					if f9iter.DataSource.S3DataSource.S3URI != nil {
						f9elemf3f1.SetS3Uri(*f9iter.DataSource.S3DataSource.S3URI)
					}
					f9elemf3.SetS3DataSource(f9elemf3f1)
				}
				f9elem.SetDataSource(f9elemf3)
			}
			if f9iter.InputMode != nil {
				f9elem.SetInputMode(*f9iter.InputMode)
			}
			if f9iter.RecordWrapperType != nil {
				f9elem.SetRecordWrapperType(*f9iter.RecordWrapperType)
			}
			if f9iter.ShuffleConfig != nil {
				f9elemf6 := &svcsdk.ShuffleConfig{}
				if f9iter.ShuffleConfig.Seed != nil {
					f9elemf6.SetSeed(*f9iter.ShuffleConfig.Seed)
				}
				f9elem.SetShuffleConfig(f9elemf6)
			}
			f9 = append(f9, f9elem)
		}
		res.SetInputDataConfig(f9)
	}
	if r.ko.Spec.OutputDataConfig != nil {
		f10 := &svcsdk.OutputDataConfig{}
		if r.ko.Spec.OutputDataConfig.KMSKeyID != nil {
			f10.SetKmsKeyId(*r.ko.Spec.OutputDataConfig.KMSKeyID)
		}
		if r.ko.Spec.OutputDataConfig.S3OutputPath != nil {
			f10.SetS3OutputPath(*r.ko.Spec.OutputDataConfig.S3OutputPath)
		}
		res.SetOutputDataConfig(f10)
	}
	if r.ko.Spec.ResourceConfig != nil {
		f11 := &svcsdk.ResourceConfig{}
		if r.ko.Spec.ResourceConfig.InstanceCount != nil {
			f11.SetInstanceCount(*r.ko.Spec.ResourceConfig.InstanceCount)
		}
		if r.ko.Spec.ResourceConfig.InstanceType != nil {
			f11.SetInstanceType(*r.ko.Spec.ResourceConfig.InstanceType)
		}
		if r.ko.Spec.ResourceConfig.VolumeKMSKeyID != nil {
			f11.SetVolumeKmsKeyId(*r.ko.Spec.ResourceConfig.VolumeKMSKeyID)
		}
		if r.ko.Spec.ResourceConfig.VolumeSizeInGB != nil {
			f11.SetVolumeSizeInGB(*r.ko.Spec.ResourceConfig.VolumeSizeInGB)
		}
		res.SetResourceConfig(f11)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.StoppingCondition != nil {
		f13 := &svcsdk.StoppingCondition{}
		if r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds != nil {
			f13.SetMaxRuntimeInSeconds(*r.ko.Spec.StoppingCondition.MaxRuntimeInSeconds)
		}
		if r.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds != nil {
			f13.SetMaxWaitTimeInSeconds(*r.ko.Spec.StoppingCondition.MaxWaitTimeInSeconds)
		}
		res.SetStoppingCondition(f13)
	}
	if r.ko.Spec.Tags != nil {
		f14 := []*svcsdk.Tag{}
		for _, f14iter := range r.ko.Spec.Tags {
			f14elem := &svcsdk.Tag{}
			if f14iter.Key != nil {
				f14elem.SetKey(*f14iter.Key)
			}
			if f14iter.Value != nil {
				f14elem.SetValue(*f14iter.Value)
			}
			f14 = append(f14, f14elem)
		}
		res.SetTags(f14)
	}
	if r.ko.Spec.TensorBoardOutputConfig != nil {
		f15 := &svcsdk.TensorBoardOutputConfig{}
		if r.ko.Spec.TensorBoardOutputConfig.LocalPath != nil {
			f15.SetLocalPath(*r.ko.Spec.TensorBoardOutputConfig.LocalPath)
		}
		if r.ko.Spec.TensorBoardOutputConfig.S3OutputPath != nil {
			f15.SetS3OutputPath(*r.ko.Spec.TensorBoardOutputConfig.S3OutputPath)
		}
		res.SetTensorBoardOutputConfig(f15)
	}
	if r.ko.Spec.TrainingJobName != nil {
		res.SetTrainingJobName(*r.ko.Spec.TrainingJobName)
	}
	if r.ko.Spec.VPCConfig != nil {
		f17 := &svcsdk.VpcConfig{}
		if r.ko.Spec.VPCConfig.SecurityGroupIDs != nil {
			f17f0 := []*string{}
			for _, f17f0iter := range r.ko.Spec.VPCConfig.SecurityGroupIDs {
				var f17f0elem string
				f17f0elem = *f17f0iter
				f17f0 = append(f17f0, &f17f0elem)
			}
			f17.SetSecurityGroupIds(f17f0)
		}
		if r.ko.Spec.VPCConfig.Subnets != nil {
			f17f1 := []*string{}
			for _, f17f1iter := range r.ko.Spec.VPCConfig.Subnets {
				var f17f1elem string
				f17f1elem = *f17f1iter
				f17f1 = append(f17f1, &f17f1elem)
			}
			f17.SetSubnets(f17f1)
		}
		res.SetVpcConfig(f17)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.StopTrainingJobWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "StopTrainingJob", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.StopTrainingJobInput, error) {
	res := &svcsdk.StopTrainingJobInput{}

	if r.ko.Spec.TrainingJobName != nil {
		res.SetTrainingJobName(*r.ko.Spec.TrainingJobName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.TrainingJob,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
