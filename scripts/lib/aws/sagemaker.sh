#!/usr/bin/env bash



#!/usr/bin/env bash

THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="$THIS_DIR/../../.."
SCRIPTS_DIR="$ROOT_DIR/scripts"

. $SCRIPTS_DIR/lib/common.sh
. $SCRIPTS_DIR/lib/testutil.sh
. $SCRIPTS_DIR/lib/aws.sh

service_name="sagemaker"

#################################################
# functions for tests
#################################################

# print_k8s_ack_controller_pod_logs prints kubernetes ack controller pod logs
# this function depends upon testutil.sh
print_k8s_ack_controller_pod_logs() {
  local ack_ctrl_pod_id=$( controller_pod_id )
  #kubectl logs -n ack-system "$ack_ctrl_pod_id"
}

# TODO: Read Env variables ?
setup_sagemaker_common_test_inputs() {
  # uses non local variable for later use in tests
  sagemaker_data_bucket="REPLACE"
  sagemaker_execution_role_arn="REPLACE"
  sagemaker_region="us-west-2"
  sagemaker_status_failed="Failed"
  sagemaker_status_inprogress="InProgress"
  sagemaker_status_completed="Completed"
  sagemaker_status_stopped="Stopped"
  sagemaker_status_stopping="Stopping"


}

setup_sagemaker_common_test_inputs

assert_aws_sagemaker_model_created() {
  local model_name="$1"
  local model_arn="$(daws sagemaker describe-model --model-name "$model_name" | jq .ModelArn)"
  if [ -z "$model_arn" ]; then
    echo "ERROR: ModelArn not found for $model_name"
    return 1
  fi
  return 0
}

# is_aws_sagemaker_trainingjob_exists() returns 0 if an TrainingJob with the supplied name
# exists, 1 otherwise.
#
# training_job_created TRAINING_JOB_NAME
# Arguments:
#
#   TRAINING_JOB_NAME  required string for the name of the bucket to check
#
# Usage:
#
#   if ! training_job_created "$training_job_name"; then
#       echo "Training Job $training_job_name does not exist!"
#   fi
is_aws_sagemaker_trainingjob_exists() {
    local __training_job_name="$1"

    # TODO: This is a simple test right now, should ideally check for required status
    local __training_job_status="$(daws sagemaker describe-training-job --region $sagemaker_region --training-job-name $__training_job_name --output json | jq .TrainingJobStatus)"

    if [[ $? -eq 0 ]]; then
        echo "$__training_job_name found!"
        return 0
    else
        echo "$__training_job_name not found!"
        print_k8s_ack_controller_pod_logs
        return 1
    fi
}

# get_aws_sagemaker_trainingjob_status() prints the SageMaker status of the TrainingJob with the supplied name.
# If the job is not found, it returns a 1.
#
# get_aws_sagemaker_trainingjob_status TRAINING_JOB_NAME
# Arguments:
#
#   TRAINING_JOB_NAME  required string for the name of the trainingJob to check
#
# Usage:
#
#   local training_job_status=$(get_aws_sagemaker_trainingjob_status "${__training_job_name}")
get_aws_sagemaker_trainingjob_status() {
    local __training_job_name="$1"
    local __training_job_status="$(daws sagemaker describe-training-job --region $sagemaker_region --training-job-name $__training_job_name --output json | jq .TrainingJobStatus)"

    if [ -z "${__training_job_status}" ]; then
    echo "ERROR: trainingJob $__training_job_status not found"
    print_k8s_ack_controller_pod_logs
    exit 1
    else
        echo "${__training_job_status}"
        return 0
    fi
}

# assert_aws_sagemaker_trainingjob_created() checks if the SageMaker status of the TrainingJob with the supplied name is in creating or created.
# If not, it returns a 1.
#
# assert_aws_sagemaker_trainingjob_created TRAINING_JOB_NAME
# Arguments:
#
#   TRAINING_JOB_NAME  required string for the name of the trainingJob to check
#
# Usage:
#
#   if ! assert_aws_sagemaker_trainingjob_created "$training_job_name"; then
#       echo "Training Job $training_job_name was not created, Check logs!"
#   fi
assert_aws_sagemaker_trainingjob_created() {
  local __training_job_name="$1"
  local __training_job_status=$(get_aws_sagemaker_trainingjob_status "${__training_job_name}") 
  if [[ "$__training_job_status" = "$sagemaker_status_inprogress" ] || [ "$__training_job_status" = "$sagemaker_status_completed" ]]; then
    echo "SUCCESS: TrainingJob $__training_job_name created and has status ${__training_job_status}"
    return 0
  else 
    echo "ERROR: TrainingJob $__training_job_name was not created and has status ${__training_job_status}"
    print_k8s_ack_controller_pod_logs
    return 1
  fi
}

# assert_aws_sagemaker_trainingjob_stopped() checks if the SageMaker status of the TrainingJob with the supplied name is in creating or created.
# If not, it returns a 1.
#
# assert_aws_sagemaker_trainingjob_stopped TRAINING_JOB_NAME
# Arguments:
#
#   TRAINING_JOB_NAME  required string for the name of the trainingJob to check
#
# Usage:
#
#   if ! assert_aws_sagemaker_trainingjob_stopped "$training_job_name"; then
#       echo "Training Job $training_job_name was not stopped, Check logs!"
#   fi
assert_aws_sagemaker_trainingjob_stopped() {
  local __training_job_name="$1"
  local __training_job_status=$(get_aws_sagemaker_trainingjob_status "${__training_job_name}")  
  
  # TODO: should the first condition be considered a failure ? 
  if [[ "$__training_job_status" = "$sagemaker_status_completed" ]]; then
    echo "WARNING: TrainingJob $__training_job_status has ${__training_job_status} status, cannot stop now."
    return 0
  else if [[ "$__training_job_status" = "$sagemaker_status_stopping" ] || [ $__training_job_status = "$sagemaker_status_stopped" ]]; then
    echo "SUCCESS: TrainingJob $__training_job_name stopped and has status ${__training_job_status}"
    return 0
  else 
    echo "ERROR: TrainingJob $__training_job_name was not stopped and has status ${__training_job_status}"
    print_k8s_ack_controller_pod_logs 
    return 1
  fi
}
