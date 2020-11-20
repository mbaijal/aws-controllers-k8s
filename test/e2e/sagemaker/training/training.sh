#!/usr/bin/env bash

##############################################
# Tests for AWS SageMaker Training
##############################################

set -u
set -x

THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$THIS_DIR/../../../.."
SCRIPTS_DIR="$ROOT_DIR/scripts"

source "$SCRIPTS_DIR/lib/common.sh"
source "$SCRIPTS_DIR/lib/k8s.sh"
source "$SCRIPTS_DIR/lib/testutil.sh"
source "$SCRIPTS_DIR/lib/aws/sagemaker.sh"

service_name="sagemaker"
test_name="$( filenoext "${BASH_SOURCE[0]}" )"
ack_ctrl_pod_id=$( controller_pod_id )
debug_msg "executing test: $service_name/$test_name"

setup_sagemaker_training_test_inputs() {
  training_job_name="ack-training-job-"$(date '+%Y-%m-%d-%H-%M-%S')""
  training_image_us_west_2="433757028032.dkr.ecr.us-west-2.amazonaws.com/xgboost:1"
}

setup_sagemaker_training_test_inputs
assert_equal "$sagemaker_region" "$AWS_REGION" "Expected AWS_REGION ($AWS_REGION) to be the same as sagemaker_region ($sagemaker_region)" || exit 1

get_xgboost_training_yaml() {
  cat <<EOF
apiVersion: sagemaker.services.k8s.aws/v1alpha1
kind: TrainingJob
metadata:
  name: $training_job_name
spec:
  trainingJobName: $training_job_name
  roleARN: $sagemaker_execution_role_arn
  hyperParameters:
    max_depth: "5"
    eta: "0.2"
    gamma: "4"
    min_child_weight: "6"
  algorithmSpecification:
    trainingImage: $training_image_us_west_2
    trainingInputMode: File
  outputDataConfig:
    s3OutputPath: s3://$sagemaker_data_bucket/xgboost/
  resourceConfig:
    instanceCount: 1
    instanceType: ml.m4.xlarge
    volumeSizeInGB: 5
  stoppingCondition:
    maxRuntimeInSeconds: 86400
  inputDataConfig:
    - channelName: train
      dataSource:
        s3DataSource:
          s3DataType: S3Prefix
          s3URI: s3://$sagemaker_data_bucket/xgboost/
          s3DataDistributionType: FullyReplicated
      contentType: text/csv
      compressionType: None
    - channelName: validation
      dataSource:
        s3DataSource:
          s3DataType: S3Prefix
          s3URI: s3://$sagemaker_data_bucket/validation/
          s3DataDistributionType: FullyReplicated
      contentType: text/csv
      compressionType: None
  tags:
    - key: key
      value: value
EOF
}


get_xgboost_training_debugger_yaml() {
  cat <<EOF
apiVersion: sagemaker.services.k8s.aws/v1alpha1
kind: TrainingJob
metadata:
  name: $training_job_name
spec:
  trainingJobName: $training_job_name
  roleARN: $sagemaker_execution_role_arn
  hyperParameters:
    max_depth: "5"
    eta: "0.2"
    gamma: "4"
    min_child_weight: "6"
  algorithmSpecification:
    trainingImage: $training_image_us_west_2
    trainingInputMode: File
  outputDataConfig:
    s3OutputPath: s3://$sagemaker_data_bucket/xgboost/
  resourceConfig:
    instanceCount: 1
    instanceType: ml.m4.xlarge
    volumeSizeInGB: 5
  stoppingCondition:
    maxRuntimeInSeconds: 86400
  inputDataConfig:
    - channelName: train
      dataSource:
        s3DataSource:
          s3DataType: S3Prefix
          s3URI: s3://$sagemaker_data_bucket/xgboost/
          s3DataDistributionType: FullyReplicated
      contentType: text/csv
      compressionType: None
    - channelName: validation
      dataSource:
        s3DataSource:
          s3DataType: S3Prefix
          s3URI: s3://$sagemaker_data_bucket/validation/
          s3DataDistributionType: FullyReplicated
      contentType: text/csv
      compressionType: None
  debugHookConfig:
    s3OutputPath: s3://$sagemaker_data_bucket/xgboost-debugger/hookconfig
    collectionConfigurations:
      - collectionName: feature_importance
        collectionParameters:
          - name: save_interval
            value: "5"
      - collectionName: losses
        collectionParameters:
          - name: save_interval"
            value: "500" 
      - collectionName: average_shap
        collectionParameters:
          - name: save_interval
            value: "5" 
      - collectionName: metrics
        collectionParameters:
          - name: save_interval
            value: "5" 
  debugRuleConfigurations:
    - ruleConfigurationName: LossNotDecreasing
      ruleEvaluatorImage: 895741380848.dkr.ecr.us-west-2.amazonaws.com/sagemaker-debugger-rules:latest
      ruleParameters:
        - name: collection_names
          value: metrics
        - name: num_steps
          value: "10"
        - name: rule_to_invoke
          value: LossNotDecreasing
  tags:
    - key: key
      value: value
EOF
}

#################################################
# PRE-CHECKS
#################################################
if training_job_created "$training_job_name"; then
    echo "FAIL: expected $training_job_name to not exist in SageMaker. Did previous test run cleanup?"
    exit 1
fi

if k8s_resource_exists "$training_job_name"; then
    echo "FAIL: expected $training_job_name to not exist. Did previous test run cleanup?"
    exit 1
fi

#################################################
# Create training job
#################################################
k8s_create_trainingjob() {
  debug_msg "Testing create trainingJob: $training_job_name."
  __yaml="$(get_xgboost_training_yaml)"
  echo "$__yaml" | kubectl apply -f -
}
k8s_create_trainingjob

# TODO: remove sleep wait for job cr to have ARN
sleep 10
debug_msg "checking trainingJob $training_job_name created in SageMaker"
assert_aws_sagemaker_trainingjob_created "$training_job_name" || exit 1

#################################################
# Stop (Delete) training job
#################################################
debug_msg "Testing Stopping TrainingJob: $training_job_name."
kubectl delete TrainingJob/"$training_job_name" 2>/dev/null
assert_equal "0" "$?" "Expected success from kubectl delete but got $?" || exit 1

assert_aws_sagemaker_trainingjob_created "$training_job_name" || exit 1

debug_msg "Check that the controller pod was not restarted."
assert_pod_not_restarted $ack_ctrl_pod_id
