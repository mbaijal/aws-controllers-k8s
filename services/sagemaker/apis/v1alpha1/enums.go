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

type AWSManagedHumanLoopRequestSource string

const (
	AWSManagedHumanLoopRequestSource_AWS_Rekognition_DetectModerationLabels_Image_V3 AWSManagedHumanLoopRequestSource = "AWS/Rekognition/DetectModerationLabels/Image/V3"
	AWSManagedHumanLoopRequestSource_AWS_Textract_AnalyzeDocument_Forms_V1           AWSManagedHumanLoopRequestSource = "AWS/Textract/AnalyzeDocument/Forms/V1"
)

type AlgorithmSortBy string

const (
	AlgorithmSortBy_Name         AlgorithmSortBy = "Name"
	AlgorithmSortBy_CreationTime AlgorithmSortBy = "CreationTime"
)

type AlgorithmStatus string

const (
	AlgorithmStatus_Pending    AlgorithmStatus = "Pending"
	AlgorithmStatus_InProgress AlgorithmStatus = "InProgress"
	AlgorithmStatus_Completed  AlgorithmStatus = "Completed"
	AlgorithmStatus_Failed     AlgorithmStatus = "Failed"
	AlgorithmStatus_Deleting   AlgorithmStatus = "Deleting"
)

type AppInstanceType string

const (
	AppInstanceType_system           AppInstanceType = "system"
	AppInstanceType_ml_t3_micro      AppInstanceType = "ml.t3.micro"
	AppInstanceType_ml_t3_small      AppInstanceType = "ml.t3.small"
	AppInstanceType_ml_t3_medium     AppInstanceType = "ml.t3.medium"
	AppInstanceType_ml_t3_large      AppInstanceType = "ml.t3.large"
	AppInstanceType_ml_t3_xlarge     AppInstanceType = "ml.t3.xlarge"
	AppInstanceType_ml_t3_2xlarge    AppInstanceType = "ml.t3.2xlarge"
	AppInstanceType_ml_m5_large      AppInstanceType = "ml.m5.large"
	AppInstanceType_ml_m5_xlarge     AppInstanceType = "ml.m5.xlarge"
	AppInstanceType_ml_m5_2xlarge    AppInstanceType = "ml.m5.2xlarge"
	AppInstanceType_ml_m5_4xlarge    AppInstanceType = "ml.m5.4xlarge"
	AppInstanceType_ml_m5_8xlarge    AppInstanceType = "ml.m5.8xlarge"
	AppInstanceType_ml_m5_12xlarge   AppInstanceType = "ml.m5.12xlarge"
	AppInstanceType_ml_m5_16xlarge   AppInstanceType = "ml.m5.16xlarge"
	AppInstanceType_ml_m5_24xlarge   AppInstanceType = "ml.m5.24xlarge"
	AppInstanceType_ml_c5_large      AppInstanceType = "ml.c5.large"
	AppInstanceType_ml_c5_xlarge     AppInstanceType = "ml.c5.xlarge"
	AppInstanceType_ml_c5_2xlarge    AppInstanceType = "ml.c5.2xlarge"
	AppInstanceType_ml_c5_4xlarge    AppInstanceType = "ml.c5.4xlarge"
	AppInstanceType_ml_c5_9xlarge    AppInstanceType = "ml.c5.9xlarge"
	AppInstanceType_ml_c5_12xlarge   AppInstanceType = "ml.c5.12xlarge"
	AppInstanceType_ml_c5_18xlarge   AppInstanceType = "ml.c5.18xlarge"
	AppInstanceType_ml_c5_24xlarge   AppInstanceType = "ml.c5.24xlarge"
	AppInstanceType_ml_p3_2xlarge    AppInstanceType = "ml.p3.2xlarge"
	AppInstanceType_ml_p3_8xlarge    AppInstanceType = "ml.p3.8xlarge"
	AppInstanceType_ml_p3_16xlarge   AppInstanceType = "ml.p3.16xlarge"
	AppInstanceType_ml_g4dn_xlarge   AppInstanceType = "ml.g4dn.xlarge"
	AppInstanceType_ml_g4dn_2xlarge  AppInstanceType = "ml.g4dn.2xlarge"
	AppInstanceType_ml_g4dn_4xlarge  AppInstanceType = "ml.g4dn.4xlarge"
	AppInstanceType_ml_g4dn_8xlarge  AppInstanceType = "ml.g4dn.8xlarge"
	AppInstanceType_ml_g4dn_12xlarge AppInstanceType = "ml.g4dn.12xlarge"
	AppInstanceType_ml_g4dn_16xlarge AppInstanceType = "ml.g4dn.16xlarge"
)

type AppSortKey string

const (
	AppSortKey_CreationTime AppSortKey = "CreationTime"
)

type AppStatus string

const (
	AppStatus_Deleted   AppStatus = "Deleted"
	AppStatus_Deleting  AppStatus = "Deleting"
	AppStatus_Failed    AppStatus = "Failed"
	AppStatus_InService AppStatus = "InService"
	AppStatus_Pending   AppStatus = "Pending"
)

type AppType string

const (
	AppType_JupyterServer AppType = "JupyterServer"
	AppType_KernelGateway AppType = "KernelGateway"
	AppType_TensorBoard   AppType = "TensorBoard"
)

type AssemblyType string

const (
	AssemblyType_None AssemblyType = "None"
	AssemblyType_Line AssemblyType = "Line"
)

type AuthMode string

const (
	AuthMode_SSO AuthMode = "SSO"
	AuthMode_IAM AuthMode = "IAM"
)

type AutoMLJobObjectiveType string

const (
	AutoMLJobObjectiveType_Maximize AutoMLJobObjectiveType = "Maximize"
	AutoMLJobObjectiveType_Minimize AutoMLJobObjectiveType = "Minimize"
)

type AutoMLJobSecondaryStatus string

const (
	AutoMLJobSecondaryStatus_Starting                      AutoMLJobSecondaryStatus = "Starting"
	AutoMLJobSecondaryStatus_AnalyzingData                 AutoMLJobSecondaryStatus = "AnalyzingData"
	AutoMLJobSecondaryStatus_FeatureEngineering            AutoMLJobSecondaryStatus = "FeatureEngineering"
	AutoMLJobSecondaryStatus_ModelTuning                   AutoMLJobSecondaryStatus = "ModelTuning"
	AutoMLJobSecondaryStatus_MaxCandidatesReached          AutoMLJobSecondaryStatus = "MaxCandidatesReached"
	AutoMLJobSecondaryStatus_Failed                        AutoMLJobSecondaryStatus = "Failed"
	AutoMLJobSecondaryStatus_Stopped                       AutoMLJobSecondaryStatus = "Stopped"
	AutoMLJobSecondaryStatus_MaxAutoMLJobRuntimeReached    AutoMLJobSecondaryStatus = "MaxAutoMLJobRuntimeReached"
	AutoMLJobSecondaryStatus_Stopping                      AutoMLJobSecondaryStatus = "Stopping"
	AutoMLJobSecondaryStatus_CandidateDefinitionsGenerated AutoMLJobSecondaryStatus = "CandidateDefinitionsGenerated"
)

type AutoMLJobStatus string

const (
	AutoMLJobStatus_Completed  AutoMLJobStatus = "Completed"
	AutoMLJobStatus_InProgress AutoMLJobStatus = "InProgress"
	AutoMLJobStatus_Failed     AutoMLJobStatus = "Failed"
	AutoMLJobStatus_Stopped    AutoMLJobStatus = "Stopped"
	AutoMLJobStatus_Stopping   AutoMLJobStatus = "Stopping"
)

type AutoMLMetricEnum string

const (
	AutoMLMetricEnum_Accuracy AutoMLMetricEnum = "Accuracy"
	AutoMLMetricEnum_MSE      AutoMLMetricEnum = "MSE"
	AutoMLMetricEnum_F1       AutoMLMetricEnum = "F1"
	AutoMLMetricEnum_F1macro  AutoMLMetricEnum = "F1macro"
	AutoMLMetricEnum_AUC      AutoMLMetricEnum = "AUC"
)

type AutoMLS3DataType string

const (
	AutoMLS3DataType_ManifestFile AutoMLS3DataType = "ManifestFile"
	AutoMLS3DataType_S3Prefix     AutoMLS3DataType = "S3Prefix"
)

type AutoMLSortBy string

const (
	AutoMLSortBy_Name         AutoMLSortBy = "Name"
	AutoMLSortBy_CreationTime AutoMLSortBy = "CreationTime"
	AutoMLSortBy_Status       AutoMLSortBy = "Status"
)

type AutoMLSortOrder string

const (
	AutoMLSortOrder_Ascending  AutoMLSortOrder = "Ascending"
	AutoMLSortOrder_Descending AutoMLSortOrder = "Descending"
)

type BatchStrategy string

const (
	BatchStrategy_MultiRecord  BatchStrategy = "MultiRecord"
	BatchStrategy_SingleRecord BatchStrategy = "SingleRecord"
)

type BooleanOperator string

const (
	BooleanOperator_And BooleanOperator = "And"
	BooleanOperator_Or  BooleanOperator = "Or"
)

type CandidateSortBy string

const (
	CandidateSortBy_CreationTime              CandidateSortBy = "CreationTime"
	CandidateSortBy_Status                    CandidateSortBy = "Status"
	CandidateSortBy_FinalObjectiveMetricValue CandidateSortBy = "FinalObjectiveMetricValue"
)

type CandidateStatus string

const (
	CandidateStatus_Completed  CandidateStatus = "Completed"
	CandidateStatus_InProgress CandidateStatus = "InProgress"
	CandidateStatus_Failed     CandidateStatus = "Failed"
	CandidateStatus_Stopped    CandidateStatus = "Stopped"
	CandidateStatus_Stopping   CandidateStatus = "Stopping"
)

type CandidateStepType string

const (
	CandidateStepType_AWS__SageMaker__TrainingJob   CandidateStepType = "AWS::SageMaker::TrainingJob"
	CandidateStepType_AWS__SageMaker__TransformJob  CandidateStepType = "AWS::SageMaker::TransformJob"
	CandidateStepType_AWS__SageMaker__ProcessingJob CandidateStepType = "AWS::SageMaker::ProcessingJob"
)

type CaptureMode string

const (
	CaptureMode_Input  CaptureMode = "Input"
	CaptureMode_Output CaptureMode = "Output"
)

type CaptureStatus string

const (
	CaptureStatus_Started CaptureStatus = "Started"
	CaptureStatus_Stopped CaptureStatus = "Stopped"
)

type CodeRepositorySortBy string

const (
	CodeRepositorySortBy_Name             CodeRepositorySortBy = "Name"
	CodeRepositorySortBy_CreationTime     CodeRepositorySortBy = "CreationTime"
	CodeRepositorySortBy_LastModifiedTime CodeRepositorySortBy = "LastModifiedTime"
)

type CodeRepositorySortOrder string

const (
	CodeRepositorySortOrder_Ascending  CodeRepositorySortOrder = "Ascending"
	CodeRepositorySortOrder_Descending CodeRepositorySortOrder = "Descending"
)

type CompilationJobStatus string

const (
	CompilationJobStatus_INPROGRESS CompilationJobStatus = "INPROGRESS"
	CompilationJobStatus_COMPLETED  CompilationJobStatus = "COMPLETED"
	CompilationJobStatus_FAILED     CompilationJobStatus = "FAILED"
	CompilationJobStatus_STARTING   CompilationJobStatus = "STARTING"
	CompilationJobStatus_STOPPING   CompilationJobStatus = "STOPPING"
	CompilationJobStatus_STOPPED    CompilationJobStatus = "STOPPED"
)

type CompressionType string

const (
	CompressionType_None CompressionType = "None"
	CompressionType_Gzip CompressionType = "Gzip"
)

type ContainerMode string

const (
	ContainerMode_SingleModel ContainerMode = "SingleModel"
	ContainerMode_MultiModel  ContainerMode = "MultiModel"
)

type ContentClassifier string

const (
	ContentClassifier_FreeOfPersonallyIdentifiableInformation ContentClassifier = "FreeOfPersonallyIdentifiableInformation"
	ContentClassifier_FreeOfAdultContent                      ContentClassifier = "FreeOfAdultContent"
)

type DetailedAlgorithmStatus string

const (
	DetailedAlgorithmStatus_NotStarted DetailedAlgorithmStatus = "NotStarted"
	DetailedAlgorithmStatus_InProgress DetailedAlgorithmStatus = "InProgress"
	DetailedAlgorithmStatus_Completed  DetailedAlgorithmStatus = "Completed"
	DetailedAlgorithmStatus_Failed     DetailedAlgorithmStatus = "Failed"
)

type DetailedModelPackageStatus string

const (
	DetailedModelPackageStatus_NotStarted DetailedModelPackageStatus = "NotStarted"
	DetailedModelPackageStatus_InProgress DetailedModelPackageStatus = "InProgress"
	DetailedModelPackageStatus_Completed  DetailedModelPackageStatus = "Completed"
	DetailedModelPackageStatus_Failed     DetailedModelPackageStatus = "Failed"
)

type DirectInternetAccess string

const (
	DirectInternetAccess_Enabled  DirectInternetAccess = "Enabled"
	DirectInternetAccess_Disabled DirectInternetAccess = "Disabled"
)

type DomainStatus string

const (
	DomainStatus_Deleting  DomainStatus = "Deleting"
	DomainStatus_Failed    DomainStatus = "Failed"
	DomainStatus_InService DomainStatus = "InService"
	DomainStatus_Pending   DomainStatus = "Pending"
)

type EndpointConfigSortKey string

const (
	EndpointConfigSortKey_Name         EndpointConfigSortKey = "Name"
	EndpointConfigSortKey_CreationTime EndpointConfigSortKey = "CreationTime"
)

type EndpointSortKey string

const (
	EndpointSortKey_Name         EndpointSortKey = "Name"
	EndpointSortKey_CreationTime EndpointSortKey = "CreationTime"
	EndpointSortKey_Status       EndpointSortKey = "Status"
)

type EndpointStatus string

const (
	EndpointStatus_OutOfService   EndpointStatus = "OutOfService"
	EndpointStatus_Creating       EndpointStatus = "Creating"
	EndpointStatus_Updating       EndpointStatus = "Updating"
	EndpointStatus_SystemUpdating EndpointStatus = "SystemUpdating"
	EndpointStatus_RollingBack    EndpointStatus = "RollingBack"
	EndpointStatus_InService      EndpointStatus = "InService"
	EndpointStatus_Deleting       EndpointStatus = "Deleting"
	EndpointStatus_Failed         EndpointStatus = "Failed"
)

type ExecutionStatus string

const (
	ExecutionStatus_Pending                 ExecutionStatus = "Pending"
	ExecutionStatus_Completed               ExecutionStatus = "Completed"
	ExecutionStatus_CompletedWithViolations ExecutionStatus = "CompletedWithViolations"
	ExecutionStatus_InProgress              ExecutionStatus = "InProgress"
	ExecutionStatus_Failed                  ExecutionStatus = "Failed"
	ExecutionStatus_Stopping                ExecutionStatus = "Stopping"
	ExecutionStatus_Stopped                 ExecutionStatus = "Stopped"
)

type FileSystemAccessMode string

const (
	FileSystemAccessMode_rw FileSystemAccessMode = "rw"
	FileSystemAccessMode_ro FileSystemAccessMode = "ro"
)

type FileSystemType string

const (
	FileSystemType_EFS       FileSystemType = "EFS"
	FileSystemType_FSxLustre FileSystemType = "FSxLustre"
)

type FlowDefinitionStatus string

const (
	FlowDefinitionStatus_Initializing FlowDefinitionStatus = "Initializing"
	FlowDefinitionStatus_Active       FlowDefinitionStatus = "Active"
	FlowDefinitionStatus_Failed       FlowDefinitionStatus = "Failed"
	FlowDefinitionStatus_Deleting     FlowDefinitionStatus = "Deleting"
)

type Framework string

const (
	Framework_TENSORFLOW Framework = "TENSORFLOW"
	Framework_KERAS      Framework = "KERAS"
	Framework_MXNET      Framework = "MXNET"
	Framework_ONNX       Framework = "ONNX"
	Framework_PYTORCH    Framework = "PYTORCH"
	Framework_XGBOOST    Framework = "XGBOOST"
	Framework_TFLITE     Framework = "TFLITE"
)

type HumanTaskUiStatus string

const (
	HumanTaskUiStatus_Active   HumanTaskUiStatus = "Active"
	HumanTaskUiStatus_Deleting HumanTaskUiStatus = "Deleting"
)

type HyperParameterScalingType string

const (
	HyperParameterScalingType_Auto               HyperParameterScalingType = "Auto"
	HyperParameterScalingType_Linear             HyperParameterScalingType = "Linear"
	HyperParameterScalingType_Logarithmic        HyperParameterScalingType = "Logarithmic"
	HyperParameterScalingType_ReverseLogarithmic HyperParameterScalingType = "ReverseLogarithmic"
)

type HyperParameterTuningJobObjectiveType string

const (
	HyperParameterTuningJobObjectiveType_Maximize HyperParameterTuningJobObjectiveType = "Maximize"
	HyperParameterTuningJobObjectiveType_Minimize HyperParameterTuningJobObjectiveType = "Minimize"
)

type HyperParameterTuningJobSortByOptions string

const (
	HyperParameterTuningJobSortByOptions_Name         HyperParameterTuningJobSortByOptions = "Name"
	HyperParameterTuningJobSortByOptions_Status       HyperParameterTuningJobSortByOptions = "Status"
	HyperParameterTuningJobSortByOptions_CreationTime HyperParameterTuningJobSortByOptions = "CreationTime"
)

type HyperParameterTuningJobStatus string

const (
	HyperParameterTuningJobStatus_Completed  HyperParameterTuningJobStatus = "Completed"
	HyperParameterTuningJobStatus_InProgress HyperParameterTuningJobStatus = "InProgress"
	HyperParameterTuningJobStatus_Failed     HyperParameterTuningJobStatus = "Failed"
	HyperParameterTuningJobStatus_Stopped    HyperParameterTuningJobStatus = "Stopped"
	HyperParameterTuningJobStatus_Stopping   HyperParameterTuningJobStatus = "Stopping"
)

type HyperParameterTuningJobStrategyType string

const (
	HyperParameterTuningJobStrategyType_Bayesian HyperParameterTuningJobStrategyType = "Bayesian"
	HyperParameterTuningJobStrategyType_Random   HyperParameterTuningJobStrategyType = "Random"
)

type HyperParameterTuningJobWarmStartType string

const (
	HyperParameterTuningJobWarmStartType_IdenticalDataAndAlgorithm HyperParameterTuningJobWarmStartType = "IdenticalDataAndAlgorithm"
	HyperParameterTuningJobWarmStartType_TransferLearning          HyperParameterTuningJobWarmStartType = "TransferLearning"
)

type InstanceType string

const (
	InstanceType_ml_t2_medium    InstanceType = "ml.t2.medium"
	InstanceType_ml_t2_large     InstanceType = "ml.t2.large"
	InstanceType_ml_t2_xlarge    InstanceType = "ml.t2.xlarge"
	InstanceType_ml_t2_2xlarge   InstanceType = "ml.t2.2xlarge"
	InstanceType_ml_t3_medium    InstanceType = "ml.t3.medium"
	InstanceType_ml_t3_large     InstanceType = "ml.t3.large"
	InstanceType_ml_t3_xlarge    InstanceType = "ml.t3.xlarge"
	InstanceType_ml_t3_2xlarge   InstanceType = "ml.t3.2xlarge"
	InstanceType_ml_m4_xlarge    InstanceType = "ml.m4.xlarge"
	InstanceType_ml_m4_2xlarge   InstanceType = "ml.m4.2xlarge"
	InstanceType_ml_m4_4xlarge   InstanceType = "ml.m4.4xlarge"
	InstanceType_ml_m4_10xlarge  InstanceType = "ml.m4.10xlarge"
	InstanceType_ml_m4_16xlarge  InstanceType = "ml.m4.16xlarge"
	InstanceType_ml_m5_xlarge    InstanceType = "ml.m5.xlarge"
	InstanceType_ml_m5_2xlarge   InstanceType = "ml.m5.2xlarge"
	InstanceType_ml_m5_4xlarge   InstanceType = "ml.m5.4xlarge"
	InstanceType_ml_m5_12xlarge  InstanceType = "ml.m5.12xlarge"
	InstanceType_ml_m5_24xlarge  InstanceType = "ml.m5.24xlarge"
	InstanceType_ml_c4_xlarge    InstanceType = "ml.c4.xlarge"
	InstanceType_ml_c4_2xlarge   InstanceType = "ml.c4.2xlarge"
	InstanceType_ml_c4_4xlarge   InstanceType = "ml.c4.4xlarge"
	InstanceType_ml_c4_8xlarge   InstanceType = "ml.c4.8xlarge"
	InstanceType_ml_c5_xlarge    InstanceType = "ml.c5.xlarge"
	InstanceType_ml_c5_2xlarge   InstanceType = "ml.c5.2xlarge"
	InstanceType_ml_c5_4xlarge   InstanceType = "ml.c5.4xlarge"
	InstanceType_ml_c5_9xlarge   InstanceType = "ml.c5.9xlarge"
	InstanceType_ml_c5_18xlarge  InstanceType = "ml.c5.18xlarge"
	InstanceType_ml_c5d_xlarge   InstanceType = "ml.c5d.xlarge"
	InstanceType_ml_c5d_2xlarge  InstanceType = "ml.c5d.2xlarge"
	InstanceType_ml_c5d_4xlarge  InstanceType = "ml.c5d.4xlarge"
	InstanceType_ml_c5d_9xlarge  InstanceType = "ml.c5d.9xlarge"
	InstanceType_ml_c5d_18xlarge InstanceType = "ml.c5d.18xlarge"
	InstanceType_ml_p2_xlarge    InstanceType = "ml.p2.xlarge"
	InstanceType_ml_p2_8xlarge   InstanceType = "ml.p2.8xlarge"
	InstanceType_ml_p2_16xlarge  InstanceType = "ml.p2.16xlarge"
	InstanceType_ml_p3_2xlarge   InstanceType = "ml.p3.2xlarge"
	InstanceType_ml_p3_8xlarge   InstanceType = "ml.p3.8xlarge"
	InstanceType_ml_p3_16xlarge  InstanceType = "ml.p3.16xlarge"
)

type JoinSource string

const (
	JoinSource_Input JoinSource = "Input"
	JoinSource_None  JoinSource = "None"
)

type LabelingJobStatus string

const (
	LabelingJobStatus_Initializing LabelingJobStatus = "Initializing"
	LabelingJobStatus_InProgress   LabelingJobStatus = "InProgress"
	LabelingJobStatus_Completed    LabelingJobStatus = "Completed"
	LabelingJobStatus_Failed       LabelingJobStatus = "Failed"
	LabelingJobStatus_Stopping     LabelingJobStatus = "Stopping"
	LabelingJobStatus_Stopped      LabelingJobStatus = "Stopped"
)

type ListCompilationJobsSortBy string

const (
	ListCompilationJobsSortBy_Name         ListCompilationJobsSortBy = "Name"
	ListCompilationJobsSortBy_CreationTime ListCompilationJobsSortBy = "CreationTime"
	ListCompilationJobsSortBy_Status       ListCompilationJobsSortBy = "Status"
)

type ListLabelingJobsForWorkteamSortByOptions string

const (
	ListLabelingJobsForWorkteamSortByOptions_CreationTime ListLabelingJobsForWorkteamSortByOptions = "CreationTime"
)

type ListWorkforcesSortByOptions string

const (
	ListWorkforcesSortByOptions_Name       ListWorkforcesSortByOptions = "Name"
	ListWorkforcesSortByOptions_CreateDate ListWorkforcesSortByOptions = "CreateDate"
)

type ListWorkteamsSortByOptions string

const (
	ListWorkteamsSortByOptions_Name       ListWorkteamsSortByOptions = "Name"
	ListWorkteamsSortByOptions_CreateDate ListWorkteamsSortByOptions = "CreateDate"
)

type ModelPackageSortBy string

const (
	ModelPackageSortBy_Name         ModelPackageSortBy = "Name"
	ModelPackageSortBy_CreationTime ModelPackageSortBy = "CreationTime"
)

type ModelPackageStatus string

const (
	ModelPackageStatus_Pending    ModelPackageStatus = "Pending"
	ModelPackageStatus_InProgress ModelPackageStatus = "InProgress"
	ModelPackageStatus_Completed  ModelPackageStatus = "Completed"
	ModelPackageStatus_Failed     ModelPackageStatus = "Failed"
	ModelPackageStatus_Deleting   ModelPackageStatus = "Deleting"
)

type ModelSortKey string

const (
	ModelSortKey_Name         ModelSortKey = "Name"
	ModelSortKey_CreationTime ModelSortKey = "CreationTime"
)

type MonitoringExecutionSortKey string

const (
	MonitoringExecutionSortKey_CreationTime  MonitoringExecutionSortKey = "CreationTime"
	MonitoringExecutionSortKey_ScheduledTime MonitoringExecutionSortKey = "ScheduledTime"
	MonitoringExecutionSortKey_Status        MonitoringExecutionSortKey = "Status"
)

type MonitoringScheduleSortKey string

const (
	MonitoringScheduleSortKey_Name         MonitoringScheduleSortKey = "Name"
	MonitoringScheduleSortKey_CreationTime MonitoringScheduleSortKey = "CreationTime"
	MonitoringScheduleSortKey_Status       MonitoringScheduleSortKey = "Status"
)

type NotebookInstanceAcceleratorType string

const (
	NotebookInstanceAcceleratorType_ml_eia1_medium NotebookInstanceAcceleratorType = "ml.eia1.medium"
	NotebookInstanceAcceleratorType_ml_eia1_large  NotebookInstanceAcceleratorType = "ml.eia1.large"
	NotebookInstanceAcceleratorType_ml_eia1_xlarge NotebookInstanceAcceleratorType = "ml.eia1.xlarge"
	NotebookInstanceAcceleratorType_ml_eia2_medium NotebookInstanceAcceleratorType = "ml.eia2.medium"
	NotebookInstanceAcceleratorType_ml_eia2_large  NotebookInstanceAcceleratorType = "ml.eia2.large"
	NotebookInstanceAcceleratorType_ml_eia2_xlarge NotebookInstanceAcceleratorType = "ml.eia2.xlarge"
)

type NotebookInstanceLifecycleConfigSortKey string

const (
	NotebookInstanceLifecycleConfigSortKey_Name             NotebookInstanceLifecycleConfigSortKey = "Name"
	NotebookInstanceLifecycleConfigSortKey_CreationTime     NotebookInstanceLifecycleConfigSortKey = "CreationTime"
	NotebookInstanceLifecycleConfigSortKey_LastModifiedTime NotebookInstanceLifecycleConfigSortKey = "LastModifiedTime"
)

type NotebookInstanceLifecycleConfigSortOrder string

const (
	NotebookInstanceLifecycleConfigSortOrder_Ascending  NotebookInstanceLifecycleConfigSortOrder = "Ascending"
	NotebookInstanceLifecycleConfigSortOrder_Descending NotebookInstanceLifecycleConfigSortOrder = "Descending"
)

type NotebookInstanceSortKey string

const (
	NotebookInstanceSortKey_Name         NotebookInstanceSortKey = "Name"
	NotebookInstanceSortKey_CreationTime NotebookInstanceSortKey = "CreationTime"
	NotebookInstanceSortKey_Status       NotebookInstanceSortKey = "Status"
)

type NotebookInstanceSortOrder string

const (
	NotebookInstanceSortOrder_Ascending  NotebookInstanceSortOrder = "Ascending"
	NotebookInstanceSortOrder_Descending NotebookInstanceSortOrder = "Descending"
)

type NotebookInstanceStatus string

const (
	NotebookInstanceStatus_Pending   NotebookInstanceStatus = "Pending"
	NotebookInstanceStatus_InService NotebookInstanceStatus = "InService"
	NotebookInstanceStatus_Stopping  NotebookInstanceStatus = "Stopping"
	NotebookInstanceStatus_Stopped   NotebookInstanceStatus = "Stopped"
	NotebookInstanceStatus_Failed    NotebookInstanceStatus = "Failed"
	NotebookInstanceStatus_Deleting  NotebookInstanceStatus = "Deleting"
	NotebookInstanceStatus_Updating  NotebookInstanceStatus = "Updating"
)

type NotebookOutputOption string

const (
	NotebookOutputOption_Allowed  NotebookOutputOption = "Allowed"
	NotebookOutputOption_Disabled NotebookOutputOption = "Disabled"
)

type ObjectiveStatus string

const (
	ObjectiveStatus_Succeeded ObjectiveStatus = "Succeeded"
	ObjectiveStatus_Pending   ObjectiveStatus = "Pending"
	ObjectiveStatus_Failed    ObjectiveStatus = "Failed"
)

type Operator string

const (
	Operator_Equals               Operator = "Equals"
	Operator_NotEquals            Operator = "NotEquals"
	Operator_GreaterThan          Operator = "GreaterThan"
	Operator_GreaterThanOrEqualTo Operator = "GreaterThanOrEqualTo"
	Operator_LessThan             Operator = "LessThan"
	Operator_LessThanOrEqualTo    Operator = "LessThanOrEqualTo"
	Operator_Contains             Operator = "Contains"
	Operator_Exists               Operator = "Exists"
	Operator_NotExists            Operator = "NotExists"
	Operator_In                   Operator = "In"
)

type OrderKey string

const (
	OrderKey_Ascending  OrderKey = "Ascending"
	OrderKey_Descending OrderKey = "Descending"
)

type ParameterType string

const (
	ParameterType_Integer     ParameterType = "Integer"
	ParameterType_Continuous  ParameterType = "Continuous"
	ParameterType_Categorical ParameterType = "Categorical"
	ParameterType_FreeText    ParameterType = "FreeText"
)

type ProblemType string

const (
	ProblemType_BinaryClassification     ProblemType = "BinaryClassification"
	ProblemType_MulticlassClassification ProblemType = "MulticlassClassification"
	ProblemType_Regression               ProblemType = "Regression"
)

type ProcessingInstanceType string

const (
	ProcessingInstanceType_ml_t3_medium   ProcessingInstanceType = "ml.t3.medium"
	ProcessingInstanceType_ml_t3_large    ProcessingInstanceType = "ml.t3.large"
	ProcessingInstanceType_ml_t3_xlarge   ProcessingInstanceType = "ml.t3.xlarge"
	ProcessingInstanceType_ml_t3_2xlarge  ProcessingInstanceType = "ml.t3.2xlarge"
	ProcessingInstanceType_ml_m4_xlarge   ProcessingInstanceType = "ml.m4.xlarge"
	ProcessingInstanceType_ml_m4_2xlarge  ProcessingInstanceType = "ml.m4.2xlarge"
	ProcessingInstanceType_ml_m4_4xlarge  ProcessingInstanceType = "ml.m4.4xlarge"
	ProcessingInstanceType_ml_m4_10xlarge ProcessingInstanceType = "ml.m4.10xlarge"
	ProcessingInstanceType_ml_m4_16xlarge ProcessingInstanceType = "ml.m4.16xlarge"
	ProcessingInstanceType_ml_c4_xlarge   ProcessingInstanceType = "ml.c4.xlarge"
	ProcessingInstanceType_ml_c4_2xlarge  ProcessingInstanceType = "ml.c4.2xlarge"
	ProcessingInstanceType_ml_c4_4xlarge  ProcessingInstanceType = "ml.c4.4xlarge"
	ProcessingInstanceType_ml_c4_8xlarge  ProcessingInstanceType = "ml.c4.8xlarge"
	ProcessingInstanceType_ml_p2_xlarge   ProcessingInstanceType = "ml.p2.xlarge"
	ProcessingInstanceType_ml_p2_8xlarge  ProcessingInstanceType = "ml.p2.8xlarge"
	ProcessingInstanceType_ml_p2_16xlarge ProcessingInstanceType = "ml.p2.16xlarge"
	ProcessingInstanceType_ml_p3_2xlarge  ProcessingInstanceType = "ml.p3.2xlarge"
	ProcessingInstanceType_ml_p3_8xlarge  ProcessingInstanceType = "ml.p3.8xlarge"
	ProcessingInstanceType_ml_p3_16xlarge ProcessingInstanceType = "ml.p3.16xlarge"
	ProcessingInstanceType_ml_c5_xlarge   ProcessingInstanceType = "ml.c5.xlarge"
	ProcessingInstanceType_ml_c5_2xlarge  ProcessingInstanceType = "ml.c5.2xlarge"
	ProcessingInstanceType_ml_c5_4xlarge  ProcessingInstanceType = "ml.c5.4xlarge"
	ProcessingInstanceType_ml_c5_9xlarge  ProcessingInstanceType = "ml.c5.9xlarge"
	ProcessingInstanceType_ml_c5_18xlarge ProcessingInstanceType = "ml.c5.18xlarge"
	ProcessingInstanceType_ml_m5_large    ProcessingInstanceType = "ml.m5.large"
	ProcessingInstanceType_ml_m5_xlarge   ProcessingInstanceType = "ml.m5.xlarge"
	ProcessingInstanceType_ml_m5_2xlarge  ProcessingInstanceType = "ml.m5.2xlarge"
	ProcessingInstanceType_ml_m5_4xlarge  ProcessingInstanceType = "ml.m5.4xlarge"
	ProcessingInstanceType_ml_m5_12xlarge ProcessingInstanceType = "ml.m5.12xlarge"
	ProcessingInstanceType_ml_m5_24xlarge ProcessingInstanceType = "ml.m5.24xlarge"
	ProcessingInstanceType_ml_r5_large    ProcessingInstanceType = "ml.r5.large"
	ProcessingInstanceType_ml_r5_xlarge   ProcessingInstanceType = "ml.r5.xlarge"
	ProcessingInstanceType_ml_r5_2xlarge  ProcessingInstanceType = "ml.r5.2xlarge"
	ProcessingInstanceType_ml_r5_4xlarge  ProcessingInstanceType = "ml.r5.4xlarge"
	ProcessingInstanceType_ml_r5_8xlarge  ProcessingInstanceType = "ml.r5.8xlarge"
	ProcessingInstanceType_ml_r5_12xlarge ProcessingInstanceType = "ml.r5.12xlarge"
	ProcessingInstanceType_ml_r5_16xlarge ProcessingInstanceType = "ml.r5.16xlarge"
	ProcessingInstanceType_ml_r5_24xlarge ProcessingInstanceType = "ml.r5.24xlarge"
)

type ProcessingJobStatus string

const (
	ProcessingJobStatus_InProgress ProcessingJobStatus = "InProgress"
	ProcessingJobStatus_Completed  ProcessingJobStatus = "Completed"
	ProcessingJobStatus_Failed     ProcessingJobStatus = "Failed"
	ProcessingJobStatus_Stopping   ProcessingJobStatus = "Stopping"
	ProcessingJobStatus_Stopped    ProcessingJobStatus = "Stopped"
)

type ProcessingS3CompressionType string

const (
	ProcessingS3CompressionType_None ProcessingS3CompressionType = "None"
	ProcessingS3CompressionType_Gzip ProcessingS3CompressionType = "Gzip"
)

type ProcessingS3DataDistributionType string

const (
	ProcessingS3DataDistributionType_FullyReplicated ProcessingS3DataDistributionType = "FullyReplicated"
	ProcessingS3DataDistributionType_ShardedByS3Key  ProcessingS3DataDistributionType = "ShardedByS3Key"
)

type ProcessingS3DataType string

const (
	ProcessingS3DataType_ManifestFile ProcessingS3DataType = "ManifestFile"
	ProcessingS3DataType_S3Prefix     ProcessingS3DataType = "S3Prefix"
)

type ProcessingS3InputMode string

const (
	ProcessingS3InputMode_Pipe ProcessingS3InputMode = "Pipe"
	ProcessingS3InputMode_File ProcessingS3InputMode = "File"
)

type ProcessingS3UploadMode string

const (
	ProcessingS3UploadMode_Continuous ProcessingS3UploadMode = "Continuous"
	ProcessingS3UploadMode_EndOfJob   ProcessingS3UploadMode = "EndOfJob"
)

type ProductionVariantAcceleratorType string

const (
	ProductionVariantAcceleratorType_ml_eia1_medium ProductionVariantAcceleratorType = "ml.eia1.medium"
	ProductionVariantAcceleratorType_ml_eia1_large  ProductionVariantAcceleratorType = "ml.eia1.large"
	ProductionVariantAcceleratorType_ml_eia1_xlarge ProductionVariantAcceleratorType = "ml.eia1.xlarge"
	ProductionVariantAcceleratorType_ml_eia2_medium ProductionVariantAcceleratorType = "ml.eia2.medium"
	ProductionVariantAcceleratorType_ml_eia2_large  ProductionVariantAcceleratorType = "ml.eia2.large"
	ProductionVariantAcceleratorType_ml_eia2_xlarge ProductionVariantAcceleratorType = "ml.eia2.xlarge"
)

type ProductionVariantInstanceType string

const (
	ProductionVariantInstanceType_ml_t2_medium     ProductionVariantInstanceType = "ml.t2.medium"
	ProductionVariantInstanceType_ml_t2_large      ProductionVariantInstanceType = "ml.t2.large"
	ProductionVariantInstanceType_ml_t2_xlarge     ProductionVariantInstanceType = "ml.t2.xlarge"
	ProductionVariantInstanceType_ml_t2_2xlarge    ProductionVariantInstanceType = "ml.t2.2xlarge"
	ProductionVariantInstanceType_ml_m4_xlarge     ProductionVariantInstanceType = "ml.m4.xlarge"
	ProductionVariantInstanceType_ml_m4_2xlarge    ProductionVariantInstanceType = "ml.m4.2xlarge"
	ProductionVariantInstanceType_ml_m4_4xlarge    ProductionVariantInstanceType = "ml.m4.4xlarge"
	ProductionVariantInstanceType_ml_m4_10xlarge   ProductionVariantInstanceType = "ml.m4.10xlarge"
	ProductionVariantInstanceType_ml_m4_16xlarge   ProductionVariantInstanceType = "ml.m4.16xlarge"
	ProductionVariantInstanceType_ml_m5_large      ProductionVariantInstanceType = "ml.m5.large"
	ProductionVariantInstanceType_ml_m5_xlarge     ProductionVariantInstanceType = "ml.m5.xlarge"
	ProductionVariantInstanceType_ml_m5_2xlarge    ProductionVariantInstanceType = "ml.m5.2xlarge"
	ProductionVariantInstanceType_ml_m5_4xlarge    ProductionVariantInstanceType = "ml.m5.4xlarge"
	ProductionVariantInstanceType_ml_m5_12xlarge   ProductionVariantInstanceType = "ml.m5.12xlarge"
	ProductionVariantInstanceType_ml_m5_24xlarge   ProductionVariantInstanceType = "ml.m5.24xlarge"
	ProductionVariantInstanceType_ml_m5d_large     ProductionVariantInstanceType = "ml.m5d.large"
	ProductionVariantInstanceType_ml_m5d_xlarge    ProductionVariantInstanceType = "ml.m5d.xlarge"
	ProductionVariantInstanceType_ml_m5d_2xlarge   ProductionVariantInstanceType = "ml.m5d.2xlarge"
	ProductionVariantInstanceType_ml_m5d_4xlarge   ProductionVariantInstanceType = "ml.m5d.4xlarge"
	ProductionVariantInstanceType_ml_m5d_12xlarge  ProductionVariantInstanceType = "ml.m5d.12xlarge"
	ProductionVariantInstanceType_ml_m5d_24xlarge  ProductionVariantInstanceType = "ml.m5d.24xlarge"
	ProductionVariantInstanceType_ml_c4_large      ProductionVariantInstanceType = "ml.c4.large"
	ProductionVariantInstanceType_ml_c4_xlarge     ProductionVariantInstanceType = "ml.c4.xlarge"
	ProductionVariantInstanceType_ml_c4_2xlarge    ProductionVariantInstanceType = "ml.c4.2xlarge"
	ProductionVariantInstanceType_ml_c4_4xlarge    ProductionVariantInstanceType = "ml.c4.4xlarge"
	ProductionVariantInstanceType_ml_c4_8xlarge    ProductionVariantInstanceType = "ml.c4.8xlarge"
	ProductionVariantInstanceType_ml_p2_xlarge     ProductionVariantInstanceType = "ml.p2.xlarge"
	ProductionVariantInstanceType_ml_p2_8xlarge    ProductionVariantInstanceType = "ml.p2.8xlarge"
	ProductionVariantInstanceType_ml_p2_16xlarge   ProductionVariantInstanceType = "ml.p2.16xlarge"
	ProductionVariantInstanceType_ml_p3_2xlarge    ProductionVariantInstanceType = "ml.p3.2xlarge"
	ProductionVariantInstanceType_ml_p3_8xlarge    ProductionVariantInstanceType = "ml.p3.8xlarge"
	ProductionVariantInstanceType_ml_p3_16xlarge   ProductionVariantInstanceType = "ml.p3.16xlarge"
	ProductionVariantInstanceType_ml_c5_large      ProductionVariantInstanceType = "ml.c5.large"
	ProductionVariantInstanceType_ml_c5_xlarge     ProductionVariantInstanceType = "ml.c5.xlarge"
	ProductionVariantInstanceType_ml_c5_2xlarge    ProductionVariantInstanceType = "ml.c5.2xlarge"
	ProductionVariantInstanceType_ml_c5_4xlarge    ProductionVariantInstanceType = "ml.c5.4xlarge"
	ProductionVariantInstanceType_ml_c5_9xlarge    ProductionVariantInstanceType = "ml.c5.9xlarge"
	ProductionVariantInstanceType_ml_c5_18xlarge   ProductionVariantInstanceType = "ml.c5.18xlarge"
	ProductionVariantInstanceType_ml_c5d_large     ProductionVariantInstanceType = "ml.c5d.large"
	ProductionVariantInstanceType_ml_c5d_xlarge    ProductionVariantInstanceType = "ml.c5d.xlarge"
	ProductionVariantInstanceType_ml_c5d_2xlarge   ProductionVariantInstanceType = "ml.c5d.2xlarge"
	ProductionVariantInstanceType_ml_c5d_4xlarge   ProductionVariantInstanceType = "ml.c5d.4xlarge"
	ProductionVariantInstanceType_ml_c5d_9xlarge   ProductionVariantInstanceType = "ml.c5d.9xlarge"
	ProductionVariantInstanceType_ml_c5d_18xlarge  ProductionVariantInstanceType = "ml.c5d.18xlarge"
	ProductionVariantInstanceType_ml_g4dn_xlarge   ProductionVariantInstanceType = "ml.g4dn.xlarge"
	ProductionVariantInstanceType_ml_g4dn_2xlarge  ProductionVariantInstanceType = "ml.g4dn.2xlarge"
	ProductionVariantInstanceType_ml_g4dn_4xlarge  ProductionVariantInstanceType = "ml.g4dn.4xlarge"
	ProductionVariantInstanceType_ml_g4dn_8xlarge  ProductionVariantInstanceType = "ml.g4dn.8xlarge"
	ProductionVariantInstanceType_ml_g4dn_12xlarge ProductionVariantInstanceType = "ml.g4dn.12xlarge"
	ProductionVariantInstanceType_ml_g4dn_16xlarge ProductionVariantInstanceType = "ml.g4dn.16xlarge"
	ProductionVariantInstanceType_ml_r5_large      ProductionVariantInstanceType = "ml.r5.large"
	ProductionVariantInstanceType_ml_r5_xlarge     ProductionVariantInstanceType = "ml.r5.xlarge"
	ProductionVariantInstanceType_ml_r5_2xlarge    ProductionVariantInstanceType = "ml.r5.2xlarge"
	ProductionVariantInstanceType_ml_r5_4xlarge    ProductionVariantInstanceType = "ml.r5.4xlarge"
	ProductionVariantInstanceType_ml_r5_12xlarge   ProductionVariantInstanceType = "ml.r5.12xlarge"
	ProductionVariantInstanceType_ml_r5_24xlarge   ProductionVariantInstanceType = "ml.r5.24xlarge"
	ProductionVariantInstanceType_ml_r5d_large     ProductionVariantInstanceType = "ml.r5d.large"
	ProductionVariantInstanceType_ml_r5d_xlarge    ProductionVariantInstanceType = "ml.r5d.xlarge"
	ProductionVariantInstanceType_ml_r5d_2xlarge   ProductionVariantInstanceType = "ml.r5d.2xlarge"
	ProductionVariantInstanceType_ml_r5d_4xlarge   ProductionVariantInstanceType = "ml.r5d.4xlarge"
	ProductionVariantInstanceType_ml_r5d_12xlarge  ProductionVariantInstanceType = "ml.r5d.12xlarge"
	ProductionVariantInstanceType_ml_r5d_24xlarge  ProductionVariantInstanceType = "ml.r5d.24xlarge"
	ProductionVariantInstanceType_ml_inf1_xlarge   ProductionVariantInstanceType = "ml.inf1.xlarge"
	ProductionVariantInstanceType_ml_inf1_2xlarge  ProductionVariantInstanceType = "ml.inf1.2xlarge"
	ProductionVariantInstanceType_ml_inf1_6xlarge  ProductionVariantInstanceType = "ml.inf1.6xlarge"
	ProductionVariantInstanceType_ml_inf1_24xlarge ProductionVariantInstanceType = "ml.inf1.24xlarge"
)

type RecordWrapper string

const (
	RecordWrapper_None     RecordWrapper = "None"
	RecordWrapper_RecordIO RecordWrapper = "RecordIO"
)

type RepositoryAccessMode string

const (
	RepositoryAccessMode_Platform RepositoryAccessMode = "Platform"
	RepositoryAccessMode_Vpc      RepositoryAccessMode = "Vpc"
)

type ResourceType string

const (
	ResourceType_TrainingJob              ResourceType = "TrainingJob"
	ResourceType_Experiment               ResourceType = "Experiment"
	ResourceType_ExperimentTrial          ResourceType = "ExperimentTrial"
	ResourceType_ExperimentTrialComponent ResourceType = "ExperimentTrialComponent"
)

type RetentionType string

const (
	RetentionType_Retain RetentionType = "Retain"
	RetentionType_Delete RetentionType = "Delete"
)

type RootAccess string

const (
	RootAccess_Enabled  RootAccess = "Enabled"
	RootAccess_Disabled RootAccess = "Disabled"
)

type RuleEvaluationStatus string

const (
	RuleEvaluationStatus_InProgress    RuleEvaluationStatus = "InProgress"
	RuleEvaluationStatus_NoIssuesFound RuleEvaluationStatus = "NoIssuesFound"
	RuleEvaluationStatus_IssuesFound   RuleEvaluationStatus = "IssuesFound"
	RuleEvaluationStatus_Error         RuleEvaluationStatus = "Error"
	RuleEvaluationStatus_Stopping      RuleEvaluationStatus = "Stopping"
	RuleEvaluationStatus_Stopped       RuleEvaluationStatus = "Stopped"
)

type S3DataDistribution string

const (
	S3DataDistribution_FullyReplicated S3DataDistribution = "FullyReplicated"
	S3DataDistribution_ShardedByS3Key  S3DataDistribution = "ShardedByS3Key"
)

type S3DataType string

const (
	S3DataType_ManifestFile          S3DataType = "ManifestFile"
	S3DataType_S3Prefix              S3DataType = "S3Prefix"
	S3DataType_AugmentedManifestFile S3DataType = "AugmentedManifestFile"
)

type ScheduleStatus string

const (
	ScheduleStatus_Pending   ScheduleStatus = "Pending"
	ScheduleStatus_Failed    ScheduleStatus = "Failed"
	ScheduleStatus_Scheduled ScheduleStatus = "Scheduled"
	ScheduleStatus_Stopped   ScheduleStatus = "Stopped"
)

type SearchSortOrder string

const (
	SearchSortOrder_Ascending  SearchSortOrder = "Ascending"
	SearchSortOrder_Descending SearchSortOrder = "Descending"
)

type SecondaryStatus string

const (
	SecondaryStatus_Starting                 SecondaryStatus = "Starting"
	SecondaryStatus_LaunchingMLInstances     SecondaryStatus = "LaunchingMLInstances"
	SecondaryStatus_PreparingTrainingStack   SecondaryStatus = "PreparingTrainingStack"
	SecondaryStatus_Downloading              SecondaryStatus = "Downloading"
	SecondaryStatus_DownloadingTrainingImage SecondaryStatus = "DownloadingTrainingImage"
	SecondaryStatus_Training                 SecondaryStatus = "Training"
	SecondaryStatus_Uploading                SecondaryStatus = "Uploading"
	SecondaryStatus_Stopping                 SecondaryStatus = "Stopping"
	SecondaryStatus_Stopped                  SecondaryStatus = "Stopped"
	SecondaryStatus_MaxRuntimeExceeded       SecondaryStatus = "MaxRuntimeExceeded"
	SecondaryStatus_Completed                SecondaryStatus = "Completed"
	SecondaryStatus_Failed                   SecondaryStatus = "Failed"
	SecondaryStatus_Interrupted              SecondaryStatus = "Interrupted"
	SecondaryStatus_MaxWaitTimeExceeded      SecondaryStatus = "MaxWaitTimeExceeded"
)

type SortBy string

const (
	SortBy_Name         SortBy = "Name"
	SortBy_CreationTime SortBy = "CreationTime"
	SortBy_Status       SortBy = "Status"
)

type SortExperimentsBy string

const (
	SortExperimentsBy_Name         SortExperimentsBy = "Name"
	SortExperimentsBy_CreationTime SortExperimentsBy = "CreationTime"
)

type SortOrder string

const (
	SortOrder_Ascending  SortOrder = "Ascending"
	SortOrder_Descending SortOrder = "Descending"
)

type SortTrialComponentsBy string

const (
	SortTrialComponentsBy_Name         SortTrialComponentsBy = "Name"
	SortTrialComponentsBy_CreationTime SortTrialComponentsBy = "CreationTime"
)

type SortTrialsBy string

const (
	SortTrialsBy_Name         SortTrialsBy = "Name"
	SortTrialsBy_CreationTime SortTrialsBy = "CreationTime"
)

type SplitType string

const (
	SplitType_None     SplitType = "None"
	SplitType_Line     SplitType = "Line"
	SplitType_RecordIO SplitType = "RecordIO"
	SplitType_TFRecord SplitType = "TFRecord"
)

type TargetDevice string

const (
	TargetDevice_lambda        TargetDevice = "lambda"
	TargetDevice_ml_m4         TargetDevice = "ml_m4"
	TargetDevice_ml_m5         TargetDevice = "ml_m5"
	TargetDevice_ml_c4         TargetDevice = "ml_c4"
	TargetDevice_ml_c5         TargetDevice = "ml_c5"
	TargetDevice_ml_p2         TargetDevice = "ml_p2"
	TargetDevice_ml_p3         TargetDevice = "ml_p3"
	TargetDevice_ml_g4dn       TargetDevice = "ml_g4dn"
	TargetDevice_ml_inf1       TargetDevice = "ml_inf1"
	TargetDevice_jetson_tx1    TargetDevice = "jetson_tx1"
	TargetDevice_jetson_tx2    TargetDevice = "jetson_tx2"
	TargetDevice_jetson_nano   TargetDevice = "jetson_nano"
	TargetDevice_jetson_xavier TargetDevice = "jetson_xavier"
	TargetDevice_rasp3b        TargetDevice = "rasp3b"
	TargetDevice_imx8qm        TargetDevice = "imx8qm"
	TargetDevice_deeplens      TargetDevice = "deeplens"
	TargetDevice_rk3399        TargetDevice = "rk3399"
	TargetDevice_rk3288        TargetDevice = "rk3288"
	TargetDevice_aisage        TargetDevice = "aisage"
	TargetDevice_sbe_c         TargetDevice = "sbe_c"
	TargetDevice_qcs605        TargetDevice = "qcs605"
	TargetDevice_qcs603        TargetDevice = "qcs603"
	TargetDevice_sitara_am57x  TargetDevice = "sitara_am57x"
	TargetDevice_amba_cv22     TargetDevice = "amba_cv22"
	TargetDevice_x86_win32     TargetDevice = "x86_win32"
	TargetDevice_x86_win64     TargetDevice = "x86_win64"
)

type TargetPlatformAccelerator string

const (
	TargetPlatformAccelerator_INTEL_GRAPHICS TargetPlatformAccelerator = "INTEL_GRAPHICS"
	TargetPlatformAccelerator_MALI           TargetPlatformAccelerator = "MALI"
	TargetPlatformAccelerator_NVIDIA         TargetPlatformAccelerator = "NVIDIA"
)

type TargetPlatformArch string

const (
	TargetPlatformArch_X86_64     TargetPlatformArch = "X86_64"
	TargetPlatformArch_X86        TargetPlatformArch = "X86"
	TargetPlatformArch_ARM64      TargetPlatformArch = "ARM64"
	TargetPlatformArch_ARM_EABI   TargetPlatformArch = "ARM_EABI"
	TargetPlatformArch_ARM_EABIHF TargetPlatformArch = "ARM_EABIHF"
)

type TargetPlatformOs string

const (
	TargetPlatformOs_ANDROID TargetPlatformOs = "ANDROID"
	TargetPlatformOs_LINUX   TargetPlatformOs = "LINUX"
)

type TrainingInputMode string

const (
	TrainingInputMode_Pipe TrainingInputMode = "Pipe"
	TrainingInputMode_File TrainingInputMode = "File"
)

type TrainingInstanceType string

const (
	TrainingInstanceType_ml_m4_xlarge     TrainingInstanceType = "ml.m4.xlarge"
	TrainingInstanceType_ml_m4_2xlarge    TrainingInstanceType = "ml.m4.2xlarge"
	TrainingInstanceType_ml_m4_4xlarge    TrainingInstanceType = "ml.m4.4xlarge"
	TrainingInstanceType_ml_m4_10xlarge   TrainingInstanceType = "ml.m4.10xlarge"
	TrainingInstanceType_ml_m4_16xlarge   TrainingInstanceType = "ml.m4.16xlarge"
	TrainingInstanceType_ml_g4dn_xlarge   TrainingInstanceType = "ml.g4dn.xlarge"
	TrainingInstanceType_ml_g4dn_2xlarge  TrainingInstanceType = "ml.g4dn.2xlarge"
	TrainingInstanceType_ml_g4dn_4xlarge  TrainingInstanceType = "ml.g4dn.4xlarge"
	TrainingInstanceType_ml_g4dn_8xlarge  TrainingInstanceType = "ml.g4dn.8xlarge"
	TrainingInstanceType_ml_g4dn_12xlarge TrainingInstanceType = "ml.g4dn.12xlarge"
	TrainingInstanceType_ml_g4dn_16xlarge TrainingInstanceType = "ml.g4dn.16xlarge"
	TrainingInstanceType_ml_m5_large      TrainingInstanceType = "ml.m5.large"
	TrainingInstanceType_ml_m5_xlarge     TrainingInstanceType = "ml.m5.xlarge"
	TrainingInstanceType_ml_m5_2xlarge    TrainingInstanceType = "ml.m5.2xlarge"
	TrainingInstanceType_ml_m5_4xlarge    TrainingInstanceType = "ml.m5.4xlarge"
	TrainingInstanceType_ml_m5_12xlarge   TrainingInstanceType = "ml.m5.12xlarge"
	TrainingInstanceType_ml_m5_24xlarge   TrainingInstanceType = "ml.m5.24xlarge"
	TrainingInstanceType_ml_c4_xlarge     TrainingInstanceType = "ml.c4.xlarge"
	TrainingInstanceType_ml_c4_2xlarge    TrainingInstanceType = "ml.c4.2xlarge"
	TrainingInstanceType_ml_c4_4xlarge    TrainingInstanceType = "ml.c4.4xlarge"
	TrainingInstanceType_ml_c4_8xlarge    TrainingInstanceType = "ml.c4.8xlarge"
	TrainingInstanceType_ml_p2_xlarge     TrainingInstanceType = "ml.p2.xlarge"
	TrainingInstanceType_ml_p2_8xlarge    TrainingInstanceType = "ml.p2.8xlarge"
	TrainingInstanceType_ml_p2_16xlarge   TrainingInstanceType = "ml.p2.16xlarge"
	TrainingInstanceType_ml_p3_2xlarge    TrainingInstanceType = "ml.p3.2xlarge"
	TrainingInstanceType_ml_p3_8xlarge    TrainingInstanceType = "ml.p3.8xlarge"
	TrainingInstanceType_ml_p3_16xlarge   TrainingInstanceType = "ml.p3.16xlarge"
	TrainingInstanceType_ml_p3dn_24xlarge TrainingInstanceType = "ml.p3dn.24xlarge"
	TrainingInstanceType_ml_c5_xlarge     TrainingInstanceType = "ml.c5.xlarge"
	TrainingInstanceType_ml_c5_2xlarge    TrainingInstanceType = "ml.c5.2xlarge"
	TrainingInstanceType_ml_c5_4xlarge    TrainingInstanceType = "ml.c5.4xlarge"
	TrainingInstanceType_ml_c5_9xlarge    TrainingInstanceType = "ml.c5.9xlarge"
	TrainingInstanceType_ml_c5_18xlarge   TrainingInstanceType = "ml.c5.18xlarge"
	TrainingInstanceType_ml_c5n_xlarge    TrainingInstanceType = "ml.c5n.xlarge"
	TrainingInstanceType_ml_c5n_2xlarge   TrainingInstanceType = "ml.c5n.2xlarge"
	TrainingInstanceType_ml_c5n_4xlarge   TrainingInstanceType = "ml.c5n.4xlarge"
	TrainingInstanceType_ml_c5n_9xlarge   TrainingInstanceType = "ml.c5n.9xlarge"
	TrainingInstanceType_ml_c5n_18xlarge  TrainingInstanceType = "ml.c5n.18xlarge"
)

type TrainingJobEarlyStoppingType string

const (
	TrainingJobEarlyStoppingType_Off  TrainingJobEarlyStoppingType = "Off"
	TrainingJobEarlyStoppingType_Auto TrainingJobEarlyStoppingType = "Auto"
)

type TrainingJobSortByOptions string

const (
	TrainingJobSortByOptions_Name                      TrainingJobSortByOptions = "Name"
	TrainingJobSortByOptions_CreationTime              TrainingJobSortByOptions = "CreationTime"
	TrainingJobSortByOptions_Status                    TrainingJobSortByOptions = "Status"
	TrainingJobSortByOptions_FinalObjectiveMetricValue TrainingJobSortByOptions = "FinalObjectiveMetricValue"
)

type TrainingJobStatus_SDK string

const (
	TrainingJobStatus_SDK_InProgress TrainingJobStatus_SDK = "InProgress"
	TrainingJobStatus_SDK_Completed  TrainingJobStatus_SDK = "Completed"
	TrainingJobStatus_SDK_Failed     TrainingJobStatus_SDK = "Failed"
	TrainingJobStatus_SDK_Stopping   TrainingJobStatus_SDK = "Stopping"
	TrainingJobStatus_SDK_Stopped    TrainingJobStatus_SDK = "Stopped"
)

type TransformInstanceType string

const (
	TransformInstanceType_ml_m4_xlarge   TransformInstanceType = "ml.m4.xlarge"
	TransformInstanceType_ml_m4_2xlarge  TransformInstanceType = "ml.m4.2xlarge"
	TransformInstanceType_ml_m4_4xlarge  TransformInstanceType = "ml.m4.4xlarge"
	TransformInstanceType_ml_m4_10xlarge TransformInstanceType = "ml.m4.10xlarge"
	TransformInstanceType_ml_m4_16xlarge TransformInstanceType = "ml.m4.16xlarge"
	TransformInstanceType_ml_c4_xlarge   TransformInstanceType = "ml.c4.xlarge"
	TransformInstanceType_ml_c4_2xlarge  TransformInstanceType = "ml.c4.2xlarge"
	TransformInstanceType_ml_c4_4xlarge  TransformInstanceType = "ml.c4.4xlarge"
	TransformInstanceType_ml_c4_8xlarge  TransformInstanceType = "ml.c4.8xlarge"
	TransformInstanceType_ml_p2_xlarge   TransformInstanceType = "ml.p2.xlarge"
	TransformInstanceType_ml_p2_8xlarge  TransformInstanceType = "ml.p2.8xlarge"
	TransformInstanceType_ml_p2_16xlarge TransformInstanceType = "ml.p2.16xlarge"
	TransformInstanceType_ml_p3_2xlarge  TransformInstanceType = "ml.p3.2xlarge"
	TransformInstanceType_ml_p3_8xlarge  TransformInstanceType = "ml.p3.8xlarge"
	TransformInstanceType_ml_p3_16xlarge TransformInstanceType = "ml.p3.16xlarge"
	TransformInstanceType_ml_c5_xlarge   TransformInstanceType = "ml.c5.xlarge"
	TransformInstanceType_ml_c5_2xlarge  TransformInstanceType = "ml.c5.2xlarge"
	TransformInstanceType_ml_c5_4xlarge  TransformInstanceType = "ml.c5.4xlarge"
	TransformInstanceType_ml_c5_9xlarge  TransformInstanceType = "ml.c5.9xlarge"
	TransformInstanceType_ml_c5_18xlarge TransformInstanceType = "ml.c5.18xlarge"
	TransformInstanceType_ml_m5_large    TransformInstanceType = "ml.m5.large"
	TransformInstanceType_ml_m5_xlarge   TransformInstanceType = "ml.m5.xlarge"
	TransformInstanceType_ml_m5_2xlarge  TransformInstanceType = "ml.m5.2xlarge"
	TransformInstanceType_ml_m5_4xlarge  TransformInstanceType = "ml.m5.4xlarge"
	TransformInstanceType_ml_m5_12xlarge TransformInstanceType = "ml.m5.12xlarge"
	TransformInstanceType_ml_m5_24xlarge TransformInstanceType = "ml.m5.24xlarge"
)

type TransformJobStatus string

const (
	TransformJobStatus_InProgress TransformJobStatus = "InProgress"
	TransformJobStatus_Completed  TransformJobStatus = "Completed"
	TransformJobStatus_Failed     TransformJobStatus = "Failed"
	TransformJobStatus_Stopping   TransformJobStatus = "Stopping"
	TransformJobStatus_Stopped    TransformJobStatus = "Stopped"
)

type TrialComponentPrimaryStatus string

const (
	TrialComponentPrimaryStatus_InProgress TrialComponentPrimaryStatus = "InProgress"
	TrialComponentPrimaryStatus_Completed  TrialComponentPrimaryStatus = "Completed"
	TrialComponentPrimaryStatus_Failed     TrialComponentPrimaryStatus = "Failed"
	TrialComponentPrimaryStatus_Stopping   TrialComponentPrimaryStatus = "Stopping"
	TrialComponentPrimaryStatus_Stopped    TrialComponentPrimaryStatus = "Stopped"
)

type UserProfileSortKey string

const (
	UserProfileSortKey_CreationTime     UserProfileSortKey = "CreationTime"
	UserProfileSortKey_LastModifiedTime UserProfileSortKey = "LastModifiedTime"
)

type UserProfileStatus string

const (
	UserProfileStatus_Deleting  UserProfileStatus = "Deleting"
	UserProfileStatus_Failed    UserProfileStatus = "Failed"
	UserProfileStatus_InService UserProfileStatus = "InService"
	UserProfileStatus_Pending   UserProfileStatus = "Pending"
)

type VariantPropertyType string

const (
	VariantPropertyType_DesiredInstanceCount VariantPropertyType = "DesiredInstanceCount"
	VariantPropertyType_DesiredWeight        VariantPropertyType = "DesiredWeight"
	VariantPropertyType_DataCaptureConfig    VariantPropertyType = "DataCaptureConfig"
)
