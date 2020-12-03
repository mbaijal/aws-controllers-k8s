#!/usr/bin/env bash

##############################################
# Tests for AWS SageMaker Training
##############################################

set -u
set -x

THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$THIS_DIR/../../../../.."
SCRIPTS_DIR="$ROOT_DIR/scripts"

source "$SCRIPTS_DIR/lib/common.sh"
source "$SCRIPTS_DIR/lib/k8s.sh"
source "$SCRIPTS_DIR/lib/testutil.sh"
source "$SCRIPTS_DIR/lib/aws/sagemaker.sh"

# Create a random trainingjob name
# Expects RANDOM_ID env var to be set
setup_sagemaker_training_test_inputs() {
  if [[ -z ${RANDOM_ID+x} || -z ${AWS_REGION+x} || -z ${SAGEMAKER_DATA_BUCKET+x} ]]; then
    echo "[FAIL] Usage setup_sagemaker_training_test_inputs. Expects RANDOM_ID, SAGEMAKER_DATA_BUCKET and AWS_REGION to be set"
    exit 1
  fi

  TRAINING_JOB_NAME="ack-trainingjob-${RANDOM_ID}"
  TRAINING_IMAGE_REGISTRY=`echo $(get_xgboost_registry $AWS_REGION)`
  echo $TRAINING_IMAGE_REGISTRY
  # TODO: Debug why line 28 returns null. The registry is hardcoded to us-west-2 for now. 
  TRAINING_IMAGE_USWEST2="433757028032.dkr.ecr.$AWS_REGION.amazonaws.com/xgboost:1"
}

get_xgboost_training_yaml() {
  cat <<EOF
apiVersion: sagemaker.services.k8s.aws/v1alpha1
kind: TrainingJob
metadata:
  name: $TRAINING_JOB_NAME
spec:
  trainingJobName: $TRAINING_JOB_NAME
  roleARN: $SAGEMAKER_EXECUTION_ROLE_ARN
  hyperParameters:
    max_depth: "5"
    eta: "0.2"
    gamma: "4"
    min_child_weight: "6"
  algorithmSpecification:
    trainingImage: $TRAINING_IMAGE_USWEST2
    trainingInputMode: File
  outputDataConfig:
    s3OutputPath: s3://$SAGEMAKER_DATA_BUCKET/xgboost/
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
          s3URI: s3://$SAGEMAKER_DATA_BUCKET/xgboost/
          s3DataDistributionType: FullyReplicated
      contentType: text/csv
      compressionType: None
    - channelName: validation
      dataSource:
        s3DataSource:
          s3DataType: S3Prefix
          s3URI: s3://$SAGEMAKER_DATA_BUCKET/validation/
          s3DataDistributionType: FullyReplicated
      contentType: text/csv
      compressionType: None
  tags:
    - key: key
      value: value
EOF
}

#################################################
# PRE-CHECKS
#################################################
sagemaker_trainingjob_prechecks() {
    setup_sagemaker_training_test_inputs
    if training_job_created "$TRAINING_JOB_NAME"; then
        echo "FAIL: expected $TRAINING_JOB_NAME to not exist in SageMaker. Did previous test run cleanup?"
        exit 1
    fi

    if k8s_resource_exists "$TRAINING_JOB_NAME"; then
        echo "FAIL: expected $TRAINING_JOB_NAME to not exist. Did previous test run cleanup?"
        exit 1
    fi
}

#################################################
# Create training job
#################################################
sagemaker_test_create_trainingjob() {
    setup_sagemaker_training_test_inputs
    debug_msg "Testing create trainingJob: $TRAINING_JOB_NAME."
    __yaml="$(get_xgboost_training_yaml)"
    echo "$__yaml" | kubectl apply -f -
    debug_msg "checking trainingJob $TRAINING_JOB_NAME created in SageMaker"
    assert_aws_sagemaker_trainingjob_created "$TRAINING_JOB_NAME" || exit 1
}

#################################################
# Stop (Delete) training job
#################################################
sagemaker_test_delete_trainingjob() {
    setup_sagemaker_training_test_inputs
    debug_msg "Testing Stopping TrainingJob: $TRAINING_JOB_NAME."
    kubectl delete TrainingJob/"$TRAINING_JOB_NAME" 2>/dev/null
    assert_equal "0" "$?" "Expected success from kubectl delete but got $?" || exit 1
    assert_aws_sagemaker_trainingjob_stopped "$TRAINING_JOB_NAME" || exit 1
}

#################################################
# Post-Checks
#################################################
sagemaker_trainingjob_postchecks() {
    debug_msg "Check that the controller pod was not restarted."
    assert_pod_not_restarted $ack_ctrl_pod_id
}
