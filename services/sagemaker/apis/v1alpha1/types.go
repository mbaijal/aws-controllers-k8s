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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AlgorithmSpecification struct {
	AlgorithmName                    *string             `json:"algorithmName,omitempty"`
	EnableSageMakerMetricsTimeSeries *bool               `json:"enableSageMakerMetricsTimeSeries,omitempty"`
	MetricDefinitions                []*MetricDefinition `json:"metricDefinitions,omitempty"`
	TrainingImage                    *string             `json:"trainingImage,omitempty"`
	TrainingInputMode                *string             `json:"trainingInputMode,omitempty"`
}

type AlgorithmValidationSpecification struct {
	ValidationRole *string `json:"validationRole,omitempty"`
}

type AutoMLCandidate struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	EndTime          *metav1.Time `json:"endTime,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type AutoMLChannel struct {
	CompressionType *string `json:"compressionType,omitempty"`
}

type AutoMLContainerDefinition struct {
	Environment  map[string]*string `json:"environment,omitempty"`
	Image        *string            `json:"image,omitempty"`
	ModelDataURL *string            `json:"modelDataURL,omitempty"`
}

type AutoMLJobSummary struct {
	AutoMLJobARN     *string      `json:"autoMLJobARN,omitempty"`
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	EndTime          *metav1.Time `json:"endTime,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type AutoMLOutputDataConfig struct {
	KMSKeyID     *string `json:"kmsKeyID,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type AutoMLS3DataSource struct {
	S3URI *string `json:"s3URI,omitempty"`
}

type AutoMLSecurityConfig struct {
	EnableInterContainerTrafficEncryption *bool      `json:"enableInterContainerTrafficEncryption,omitempty"`
	VolumeKMSKeyID                        *string    `json:"volumeKMSKeyID,omitempty"`
	VPCConfig                             *VPCConfig `json:"vpcConfig,omitempty"`
}

type Channel struct {
	ChannelName       *string        `json:"channelName,omitempty"`
	CompressionType   *string        `json:"compressionType,omitempty"`
	ContentType       *string        `json:"contentType,omitempty"`
	DataSource        *DataSource    `json:"dataSource,omitempty"`
	InputMode         *string        `json:"inputMode,omitempty"`
	RecordWrapperType *string        `json:"recordWrapperType,omitempty"`
	ShuffleConfig     *ShuffleConfig `json:"shuffleConfig,omitempty"`
}

type ChannelSpecification struct {
	IsRequired *bool   `json:"isRequired,omitempty"`
	Name       *string `json:"name,omitempty"`
}

type CheckpointConfig struct {
	LocalPath *string `json:"localPath,omitempty"`
	S3URI     *string `json:"s3URI,omitempty"`
}

type CollectionConfiguration struct {
	CollectionName       *string            `json:"collectionName,omitempty"`
	CollectionParameters map[string]*string `json:"collectionParameters,omitempty"`
}

type CompilationJobSummary struct {
	CompilationEndTime   *metav1.Time `json:"compilationEndTime,omitempty"`
	CompilationStartTime *metav1.Time `json:"compilationStartTime,omitempty"`
}

type ContainerDefinition struct {
	ContainerHostname *string            `json:"containerHostname,omitempty"`
	Environment       map[string]*string `json:"environment,omitempty"`
	Image             *string            `json:"image,omitempty"`
	ImageConfig       *ImageConfig       `json:"imageConfig,omitempty"`
	Mode              *string            `json:"mode,omitempty"`
	ModelDataURL      *string            `json:"modelDataURL,omitempty"`
	ModelPackageName  *string            `json:"modelPackageName,omitempty"`
}

type ContinuousParameterRange struct {
	MaxValue *string `json:"maxValue,omitempty"`
	MinValue *string `json:"minValue,omitempty"`
}

type ContinuousParameterRangeSpecification struct {
	MaxValue *string `json:"maxValue,omitempty"`
	MinValue *string `json:"minValue,omitempty"`
}

type DataCaptureConfig struct {
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
}

type DataCaptureConfigSummary struct {
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
}

type DataSource struct {
	FileSystemDataSource *FileSystemDataSource `json:"fileSystemDataSource,omitempty"`
	S3DataSource         *S3DataSource         `json:"s3DataSource,omitempty"`
}

type DebugHookConfig struct {
	CollectionConfigurations []*CollectionConfiguration `json:"collectionConfigurations,omitempty"`
	HookParameters           map[string]*string         `json:"hookParameters,omitempty"`
	LocalPath                *string                    `json:"localPath,omitempty"`
	S3OutputPath             *string                    `json:"s3OutputPath,omitempty"`
}

type DebugRuleConfiguration struct {
	InstanceType          *string            `json:"instanceType,omitempty"`
	LocalPath             *string            `json:"localPath,omitempty"`
	RuleConfigurationName *string            `json:"ruleConfigurationName,omitempty"`
	RuleEvaluatorImage    *string            `json:"ruleEvaluatorImage,omitempty"`
	RuleParameters        map[string]*string `json:"ruleParameters,omitempty"`
	S3OutputPath          *string            `json:"s3OutputPath,omitempty"`
	VolumeSizeInGB        *int64             `json:"volumeSizeInGB,omitempty"`
}

type DebugRuleEvaluationStatus struct {
	LastModifiedTime      *metav1.Time `json:"lastModifiedTime,omitempty"`
	RuleConfigurationName *string      `json:"ruleConfigurationName,omitempty"`
	RuleEvaluationJobARN  *string      `json:"ruleEvaluationJobARN,omitempty"`
	RuleEvaluationStatus  *string      `json:"ruleEvaluationStatus,omitempty"`
	StatusDetails         *string      `json:"statusDetails,omitempty"`
}

type DeployedImage struct {
	ResolutionTime *metav1.Time `json:"resolutionTime,omitempty"`
	ResolvedImage  *string      `json:"resolvedImage,omitempty"`
	SpecifiedImage *string      `json:"specifiedImage,omitempty"`
}

type EndpointConfigSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
}

type EndpointSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type Experiment struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	DisplayName      *string      `json:"displayName,omitempty"`
	ExperimentName   *string      `json:"experimentName,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	Tags             []*Tag       `json:"tags,omitempty"`
}

type ExperimentConfig struct {
	ExperimentName            *string `json:"experimentName,omitempty"`
	TrialComponentDisplayName *string `json:"trialComponentDisplayName,omitempty"`
	TrialName                 *string `json:"trialName,omitempty"`
}

type ExperimentSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	DisplayName      *string      `json:"displayName,omitempty"`
	ExperimentName   *string      `json:"experimentName,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type FileSystemDataSource struct {
	DirectoryPath        *string `json:"directoryPath,omitempty"`
	FileSystemAccessMode *string `json:"fileSystemAccessMode,omitempty"`
	FileSystemID         *string `json:"fileSystemID,omitempty"`
	FileSystemType       *string `json:"fileSystemType,omitempty"`
}

type FinalHyperParameterTuningJobObjectiveMetric struct {
	MetricName *string `json:"metricName,omitempty"`
}

type FlowDefinitionOutputConfig struct {
	KMSKeyID     *string `json:"kmsKeyID,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type FlowDefinitionSummary struct {
	CreationTime  *metav1.Time `json:"creationTime,omitempty"`
	FailureReason *string      `json:"failureReason,omitempty"`
}

type HumanTaskUiSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
}

type HyperParameterAlgorithmSpecification struct {
	AlgorithmName     *string             `json:"algorithmName,omitempty"`
	MetricDefinitions []*MetricDefinition `json:"metricDefinitions,omitempty"`
	TrainingImage     *string             `json:"trainingImage,omitempty"`
	TrainingInputMode *string             `json:"trainingInputMode,omitempty"`
}

type HyperParameterSpecification struct {
	DefaultValue *string `json:"defaultValue,omitempty"`
	IsRequired   *bool   `json:"isRequired,omitempty"`
	IsTunable    *bool   `json:"isTunable,omitempty"`
}

type HyperParameterTrainingJobDefinition struct {
	CheckpointConfig                      *CheckpointConfig  `json:"checkpointConfig,omitempty"`
	EnableInterContainerTrafficEncryption *bool              `json:"enableInterContainerTrafficEncryption,omitempty"`
	EnableManagedSpotTraining             *bool              `json:"enableManagedSpotTraining,omitempty"`
	EnableNetworkIsolation                *bool              `json:"enableNetworkIsolation,omitempty"`
	InputDataConfig                       []*Channel         `json:"inputDataConfig,omitempty"`
	OutputDataConfig                      *OutputDataConfig  `json:"outputDataConfig,omitempty"`
	ResourceConfig                        *ResourceConfig    `json:"resourceConfig,omitempty"`
	RoleARN                               *string            `json:"roleARN,omitempty"`
	StaticHyperParameters                 map[string]*string `json:"staticHyperParameters,omitempty"`
	StoppingCondition                     *StoppingCondition `json:"stoppingCondition,omitempty"`
	VPCConfig                             *VPCConfig         `json:"vpcConfig,omitempty"`
}

type HyperParameterTrainingJobSummary struct {
	CreationTime         *metav1.Time       `json:"creationTime,omitempty"`
	FailureReason        *string            `json:"failureReason,omitempty"`
	TrainingEndTime      *metav1.Time       `json:"trainingEndTime,omitempty"`
	TrainingJobARN       *string            `json:"trainingJobARN,omitempty"`
	TrainingJobName      *string            `json:"trainingJobName,omitempty"`
	TrainingJobStatus    *string            `json:"trainingJobStatus,omitempty"`
	TrainingStartTime    *metav1.Time       `json:"trainingStartTime,omitempty"`
	TunedHyperParameters map[string]*string `json:"tunedHyperParameters,omitempty"`
}

type HyperParameterTuningJobObjective struct {
	MetricName *string `json:"metricName,omitempty"`
}

type HyperParameterTuningJobSummary struct {
	CreationTime                *metav1.Time `json:"creationTime,omitempty"`
	HyperParameterTuningEndTime *metav1.Time `json:"hyperParameterTuningEndTime,omitempty"`
	HyperParameterTuningJobARN  *string      `json:"hyperParameterTuningJobARN,omitempty"`
	LastModifiedTime            *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type ImageConfig struct {
	RepositoryAccessMode *string `json:"repositoryAccessMode,omitempty"`
}

type InputConfig struct {
	S3URI *string `json:"s3URI,omitempty"`
}

type IntegerParameterRange struct {
	MaxValue *string `json:"maxValue,omitempty"`
	MinValue *string `json:"minValue,omitempty"`
}

type IntegerParameterRangeSpecification struct {
	MaxValue *string `json:"maxValue,omitempty"`
	MinValue *string `json:"minValue,omitempty"`
}

type LabelingJobAlgorithmsConfig struct {
	InitialActiveLearningModelARN *string `json:"initialActiveLearningModelARN,omitempty"`
}

type LabelingJobForWorkteamSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
}

type LabelingJobOutput struct {
	FinalActiveLearningModelARN *string `json:"finalActiveLearningModelARN,omitempty"`
	OutputDatasetS3URI          *string `json:"outputDatasetS3URI,omitempty"`
}

type LabelingJobOutputConfig struct {
	KMSKeyID     *string `json:"kmsKeyID,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type LabelingJobResourceConfig struct {
	VolumeKMSKeyID *string `json:"volumeKMSKeyID,omitempty"`
}

type LabelingJobS3DataSource struct {
	ManifestS3URI *string `json:"manifestS3URI,omitempty"`
}

type LabelingJobSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	FailureReason    *string      `json:"failureReason,omitempty"`
	LabelingJobARN   *string      `json:"labelingJobARN,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type MetricData struct {
	MetricName *string      `json:"metricName,omitempty"`
	Timestamp  *metav1.Time `json:"timestamp,omitempty"`
	Value      *float64     `json:"value,omitempty"`
}

type MetricDefinition struct {
	Name  *string `json:"name,omitempty"`
	Regex *string `json:"regex,omitempty"`
}

type ModelArtifacts struct {
	S3ModelArtifacts *string `json:"s3ModelArtifacts,omitempty"`
}

type ModelPackageContainerDefinition struct {
	ContainerHostname *string `json:"containerHostname,omitempty"`
	Image             *string `json:"image,omitempty"`
	ModelDataURL      *string `json:"modelDataURL,omitempty"`
}

type ModelPackageValidationSpecification struct {
	ValidationRole *string `json:"validationRole,omitempty"`
}

type ModelSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	ModelARN     *string      `json:"modelARN,omitempty"`
	ModelName    *string      `json:"modelName,omitempty"`
}

type MonitoringAppSpecification struct {
	PostAnalyticsProcessorSourceURI *string `json:"postAnalyticsProcessorSourceURI,omitempty"`
	RecordPreprocessorSourceURI     *string `json:"recordPreprocessorSourceURI,omitempty"`
}

type MonitoringClusterConfig struct {
	InstanceType   *string `json:"instanceType,omitempty"`
	VolumeKMSKeyID *string `json:"volumeKMSKeyID,omitempty"`
}

type MonitoringConstraintsResource struct {
	S3URI *string `json:"s3URI,omitempty"`
}

type MonitoringExecutionSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	FailureReason    *string      `json:"failureReason,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	ProcessingJobARN *string      `json:"processingJobARN,omitempty"`
	ScheduledTime    *metav1.Time `json:"scheduledTime,omitempty"`
}

type MonitoringJobDefinition struct {
	RoleARN *string `json:"roleARN,omitempty"`
}

type MonitoringOutputConfig struct {
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
}

type MonitoringScheduleSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

type MonitoringStatisticsResource struct {
	S3URI *string `json:"s3URI,omitempty"`
}

type NetworkConfig struct {
	EnableInterContainerTrafficEncryption *bool      `json:"enableInterContainerTrafficEncryption,omitempty"`
	EnableNetworkIsolation                *bool      `json:"enableNetworkIsolation,omitempty"`
	VPCConfig                             *VPCConfig `json:"vpcConfig,omitempty"`
}

type OutputConfig struct {
	S3OutputLocation *string `json:"s3OutputLocation,omitempty"`
}

type OutputDataConfig struct {
	KMSKeyID     *string `json:"kmsKeyID,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type Parent struct {
	ExperimentName *string `json:"experimentName,omitempty"`
	TrialName      *string `json:"trialName,omitempty"`
}

type ProcessingClusterConfig struct {
	InstanceType   *string `json:"instanceType,omitempty"`
	VolumeKMSKeyID *string `json:"volumeKMSKeyID,omitempty"`
}

type ProcessingJob struct {
	AutoMLJobARN        *string           `json:"autoMLJobARN,omitempty"`
	CreationTime        *metav1.Time      `json:"creationTime,omitempty"`
	ExperimentConfig    *ExperimentConfig `json:"experimentConfig,omitempty"`
	FailureReason       *string           `json:"failureReason,omitempty"`
	LastModifiedTime    *metav1.Time      `json:"lastModifiedTime,omitempty"`
	ProcessingEndTime   *metav1.Time      `json:"processingEndTime,omitempty"`
	ProcessingJobARN    *string           `json:"processingJobARN,omitempty"`
	ProcessingStartTime *metav1.Time      `json:"processingStartTime,omitempty"`
	RoleARN             *string           `json:"roleARN,omitempty"`
	Tags                []*Tag            `json:"tags,omitempty"`
	TrainingJobARN      *string           `json:"trainingJobARN,omitempty"`
}

type ProcessingJobSummary struct {
	CreationTime      *metav1.Time `json:"creationTime,omitempty"`
	FailureReason     *string      `json:"failureReason,omitempty"`
	LastModifiedTime  *metav1.Time `json:"lastModifiedTime,omitempty"`
	ProcessingEndTime *metav1.Time `json:"processingEndTime,omitempty"`
	ProcessingJobARN  *string      `json:"processingJobARN,omitempty"`
}

type ProcessingOutputConfig struct {
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
}

type ProcessingS3Input struct {
	S3URI *string `json:"s3URI,omitempty"`
}

type ProcessingS3Output struct {
	S3URI *string `json:"s3URI,omitempty"`
}

type ProductionVariant struct {
	ModelName *string `json:"modelName,omitempty"`
}

type ResourceConfig struct {
	InstanceCount  *int64  `json:"instanceCount,omitempty"`
	InstanceType   *string `json:"instanceType,omitempty"`
	VolumeKMSKeyID *string `json:"volumeKMSKeyID,omitempty"`
	VolumeSizeInGB *int64  `json:"volumeSizeInGB,omitempty"`
}

type S3DataSource struct {
	AttributeNames         []*string `json:"attributeNames,omitempty"`
	S3DataDistributionType *string   `json:"s3DataDistributionType,omitempty"`
	S3DataType             *string   `json:"s3DataType,omitempty"`
	S3URI                  *string   `json:"s3URI,omitempty"`
}

type SecondaryStatusTransition struct {
	EndTime       *metav1.Time `json:"endTime,omitempty"`
	StartTime     *metav1.Time `json:"startTime,omitempty"`
	Status        *string      `json:"status,omitempty"`
	StatusMessage *string      `json:"statusMessage,omitempty"`
}

type SharingSettings struct {
	S3KMSKeyID   *string `json:"s3KMSKeyID,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type ShuffleConfig struct {
	Seed *int64 `json:"seed,omitempty"`
}

type SourceAlgorithm struct {
	AlgorithmName *string `json:"algorithmName,omitempty"`
	ModelDataURL  *string `json:"modelDataURL,omitempty"`
}

type StoppingCondition struct {
	MaxRuntimeInSeconds  *int64 `json:"maxRuntimeInSeconds,omitempty"`
	MaxWaitTimeInSeconds *int64 `json:"maxWaitTimeInSeconds,omitempty"`
}

type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type TensorBoardOutputConfig struct {
	LocalPath    *string `json:"localPath,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type TrainingJobDefinition struct {
	HyperParameters   map[string]*string `json:"hyperParameters,omitempty"`
	InputDataConfig   []*Channel         `json:"inputDataConfig,omitempty"`
	OutputDataConfig  *OutputDataConfig  `json:"outputDataConfig,omitempty"`
	ResourceConfig    *ResourceConfig    `json:"resourceConfig,omitempty"`
	StoppingCondition *StoppingCondition `json:"stoppingCondition,omitempty"`
	TrainingInputMode *string            `json:"trainingInputMode,omitempty"`
}

type TrainingJobSummary struct {
	CreationTime      *metav1.Time `json:"creationTime,omitempty"`
	LastModifiedTime  *metav1.Time `json:"lastModifiedTime,omitempty"`
	TrainingEndTime   *metav1.Time `json:"trainingEndTime,omitempty"`
	TrainingJobARN    *string      `json:"trainingJobARN,omitempty"`
	TrainingJobName   *string      `json:"trainingJobName,omitempty"`
	TrainingJobStatus *string      `json:"trainingJobStatus,omitempty"`
}

type TrainingJob_SDK struct {
	AlgorithmSpecification                *AlgorithmSpecification      `json:"algorithmSpecification,omitempty"`
	AutoMLJobARN                          *string                      `json:"autoMLJobARN,omitempty"`
	BillableTimeInSeconds                 *int64                       `json:"billableTimeInSeconds,omitempty"`
	CheckpointConfig                      *CheckpointConfig            `json:"checkpointConfig,omitempty"`
	CreationTime                          *metav1.Time                 `json:"creationTime,omitempty"`
	DebugHookConfig                       *DebugHookConfig             `json:"debugHookConfig,omitempty"`
	DebugRuleConfigurations               []*DebugRuleConfiguration    `json:"debugRuleConfigurations,omitempty"`
	DebugRuleEvaluationStatuses           []*DebugRuleEvaluationStatus `json:"debugRuleEvaluationStatuses,omitempty"`
	EnableInterContainerTrafficEncryption *bool                        `json:"enableInterContainerTrafficEncryption,omitempty"`
	EnableManagedSpotTraining             *bool                        `json:"enableManagedSpotTraining,omitempty"`
	EnableNetworkIsolation                *bool                        `json:"enableNetworkIsolation,omitempty"`
	ExperimentConfig                      *ExperimentConfig            `json:"experimentConfig,omitempty"`
	FailureReason                         *string                      `json:"failureReason,omitempty"`
	FinalMetricDataList                   []*MetricData                `json:"finalMetricDataList,omitempty"`
	HyperParameters                       map[string]*string           `json:"hyperParameters,omitempty"`
	InputDataConfig                       []*Channel                   `json:"inputDataConfig,omitempty"`
	LabelingJobARN                        *string                      `json:"labelingJobARN,omitempty"`
	LastModifiedTime                      *metav1.Time                 `json:"lastModifiedTime,omitempty"`
	ModelArtifacts                        *ModelArtifacts              `json:"modelArtifacts,omitempty"`
	OutputDataConfig                      *OutputDataConfig            `json:"outputDataConfig,omitempty"`
	ResourceConfig                        *ResourceConfig              `json:"resourceConfig,omitempty"`
	RoleARN                               *string                      `json:"roleARN,omitempty"`
	SecondaryStatus                       *string                      `json:"secondaryStatus,omitempty"`
	SecondaryStatusTransitions            []*SecondaryStatusTransition `json:"secondaryStatusTransitions,omitempty"`
	StoppingCondition                     *StoppingCondition           `json:"stoppingCondition,omitempty"`
	Tags                                  []*Tag                       `json:"tags,omitempty"`
	TensorBoardOutputConfig               *TensorBoardOutputConfig     `json:"tensorBoardOutputConfig,omitempty"`
	TrainingEndTime                       *metav1.Time                 `json:"trainingEndTime,omitempty"`
	TrainingJobARN                        *string                      `json:"trainingJobARN,omitempty"`
	TrainingJobName                       *string                      `json:"trainingJobName,omitempty"`
	TrainingJobStatus                     *string                      `json:"trainingJobStatus,omitempty"`
	TrainingStartTime                     *metav1.Time                 `json:"trainingStartTime,omitempty"`
	TrainingTimeInSeconds                 *int64                       `json:"trainingTimeInSeconds,omitempty"`
	TuningJobARN                          *string                      `json:"tuningJobARN,omitempty"`
	VPCConfig                             *VPCConfig                   `json:"vpcConfig,omitempty"`
}

type TrainingSpecification struct {
	MetricDefinitions           []*MetricDefinition `json:"metricDefinitions,omitempty"`
	SupportsDistributedTraining *bool               `json:"supportsDistributedTraining,omitempty"`
	TrainingImage               *string             `json:"trainingImage,omitempty"`
}

type TransformInput struct {
	CompressionType *string `json:"compressionType,omitempty"`
	ContentType     *string `json:"contentType,omitempty"`
}

type TransformJob struct {
	AutoMLJobARN       *string           `json:"autoMLJobARN,omitempty"`
	CreationTime       *metav1.Time      `json:"creationTime,omitempty"`
	ExperimentConfig   *ExperimentConfig `json:"experimentConfig,omitempty"`
	FailureReason      *string           `json:"failureReason,omitempty"`
	LabelingJobARN     *string           `json:"labelingJobARN,omitempty"`
	ModelName          *string           `json:"modelName,omitempty"`
	Tags               []*Tag            `json:"tags,omitempty"`
	TransformEndTime   *metav1.Time      `json:"transformEndTime,omitempty"`
	TransformStartTime *metav1.Time      `json:"transformStartTime,omitempty"`
}

type TransformJobSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	FailureReason    *string      `json:"failureReason,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	TransformEndTime *metav1.Time `json:"transformEndTime,omitempty"`
}

type TransformOutput struct {
	KMSKeyID     *string `json:"kmsKeyID,omitempty"`
	S3OutputPath *string `json:"s3OutputPath,omitempty"`
}

type TransformResources struct {
	VolumeKMSKeyID *string `json:"volumeKMSKeyID,omitempty"`
}

type TransformS3DataSource struct {
	S3DataType *string `json:"s3DataType,omitempty"`
	S3URI      *string `json:"s3URI,omitempty"`
}

type Trial struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	DisplayName      *string      `json:"displayName,omitempty"`
	ExperimentName   *string      `json:"experimentName,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	Tags             []*Tag       `json:"tags,omitempty"`
	TrialName        *string      `json:"trialName,omitempty"`
}

type TrialComponent struct {
	CreationTime       *metav1.Time `json:"creationTime,omitempty"`
	DisplayName        *string      `json:"displayName,omitempty"`
	EndTime            *metav1.Time `json:"endTime,omitempty"`
	LastModifiedTime   *metav1.Time `json:"lastModifiedTime,omitempty"`
	StartTime          *metav1.Time `json:"startTime,omitempty"`
	Tags               []*Tag       `json:"tags,omitempty"`
	TrialComponentName *string      `json:"trialComponentName,omitempty"`
}

type TrialComponentMetricSummary struct {
	MetricName *string      `json:"metricName,omitempty"`
	TimeStamp  *metav1.Time `json:"timeStamp,omitempty"`
}

type TrialComponentSimpleSummary struct {
	CreationTime       *metav1.Time `json:"creationTime,omitempty"`
	TrialComponentName *string      `json:"trialComponentName,omitempty"`
}

type TrialComponentSummary struct {
	CreationTime       *metav1.Time `json:"creationTime,omitempty"`
	DisplayName        *string      `json:"displayName,omitempty"`
	EndTime            *metav1.Time `json:"endTime,omitempty"`
	LastModifiedTime   *metav1.Time `json:"lastModifiedTime,omitempty"`
	StartTime          *metav1.Time `json:"startTime,omitempty"`
	TrialComponentName *string      `json:"trialComponentName,omitempty"`
}

type TrialSummary struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	DisplayName      *string      `json:"displayName,omitempty"`
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
	TrialName        *string      `json:"trialName,omitempty"`
}

type UiConfig struct {
	UiTemplateS3URI *string `json:"uiTemplateS3URI,omitempty"`
}

type UserSettings struct {
	ExecutionRole *string `json:"executionRole,omitempty"`
}

type VPCConfig struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	Subnets          []*string `json:"subnets,omitempty"`
}

type Workforce struct {
	CreateDate      *metav1.Time `json:"createDate,omitempty"`
	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`
}

type Workteam struct {
	CreateDate      *metav1.Time `json:"createDate,omitempty"`
	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`
}
