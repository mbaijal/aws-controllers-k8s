# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.
"""Integration tests for the SageMaker TransformJob API.
"""

import boto3
import pytest
import logging
from typing import Dict
import time

from sagemaker import SERVICE_NAME, service_marker, CRD_GROUP, CRD_VERSION
from sagemaker.replacement_values import REPLACEMENT_VALUES
from common.resources import load_resource_file, random_suffix_name
from common import k8s
from time import sleep

RESOURCE_PLURAL = 'transformjobs'
TRANSFORM_JOB_STATUS_CREATED = ("InProgress", "Completed")
TRANSFORM_JOB_STATUS_STOPPED = ("Stopped", "Stopping")
TRANSFORM_JOB_STATUS_SUCCESSFUL = "Completed"
TRANSFORM_JOB_STATUS_INPROGRESS = "InProgress"

def _sagemaker_client():
    return boto3.client('sagemaker')

def _make_model():
    #Create model using boto3 for TransformJob
    transform_model_file = "s3://{d}/sagemaker/batch/model.tar.gz".format(
            d=REPLACEMENT_VALUES["SAGEMAKER_DATA_BUCKET"])
    model_name = random_suffix_name("xgboost-model", 32)

    create_response = _sagemaker_client().create_model(ModelName=model_name,
                        PrimaryContainer={
                                        'Image': REPLACEMENT_VALUES["XGBOOST_IMAGE_URI"],
                                        'ModelDataUrl': transform_model_file, 
                                        'Environment': {}
                                    },
                        ExecutionRoleArn=REPLACEMENT_VALUES["SAGEMAKER_EXECUTION_ROLE_ARN"]
                       )
    logging.debug(create_response)

    #Check if the model is created successfully
    describe_model_response = _sagemaker_client().describe_model(ModelName=model_name)
    assert describe_model_response["ModelName"] is not None

    return model_name

def _make_transformjob():
    model_name = _make_model()
    resource_name = random_suffix_name("xgboost-transformjob", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["MODEL_NAME"] = model_name
    replacements["TRANSFORM_JOB_NAME"] = resource_name

    data = load_resource_file(
        SERVICE_NAME, "xgboost_transformjob", additional_replacements=replacements
    )
    logging.debug(data)

    reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL, resource_name, namespace="default"
    )

    return reference, data

@pytest.fixture(scope="function")
def xgboost_transformjob():
    transform_job, data = _make_transformjob()
    resource = k8s.create_custom_resource(transform_job, data)
    resource = k8s.wait_resource_consumed_by_controller(transform_job)

    yield (transform_job, resource) 

    if k8s.get_resource_exists(transform_job):
        k8s.delete_custom_resource(transform_job)

def get_sagemaker_transform_job(transform_job_name: str):
    try:
        transform_desc = _sagemaker_client().describe_transform_job(
            TransformJobName=transform_job_name
        )
        return transform_desc
    except BaseException:
        logging.error(
            f"SageMaker could not find a transform job with the name {transform_job_name}"
        )
        return None

@service_marker
@pytest.mark.canary
class TestTransformJob:
    def _wait_resource_transform_status(
        self,
        reference: k8s.CustomResourceReference,
        expected_status: str,
        wait_periods: int = 30,
    ):
        resource_status = None
        for _ in range(wait_periods):
            time.sleep(30)
            resource = k8s.get_resource(reference)
            resource_status = resource["status"]["transformJobStatus"]
            if resource_status == expected_status:
                break
        else:
            logging.error(
                f"Wait for transformJobStatus: {expected_status} timed out. Actual status: {resource_status}"
            )

        return resource_status

    def _wait_sagemaker_transform_status(
        self,
        transform_job_name,
        expected_status: str,
        wait_periods: int = 30,
    ):
        actual_status = None
        for _ in range(wait_periods):
            time.sleep(30)
            transform_sm_desc = get_sagemaker_transform_job(transform_job_name)
            actual_status = transform_sm_desc["TransformJobStatus"]
            if actual_status == expected_status:
                break
        else:
            logging.error(
                f"Wait for sagemaker transform status: {expected_status} timed out. Actual status: {actual_status}"
            )

        return actual_status

    def _assert_transform_status_in_sync(
        self, transform_job_name, reference, expected_status
    ):
        assert (
            self._wait_sagemaker_transform_status(
                transform_job_name, expected_status
            )
            == self._wait_resource_transform_status(reference, expected_status)
            == expected_status
        )

    def test_transform(self, xgboost_transformjob):
        (reference, resource) = xgboost_transformjob
        assert k8s.get_resource_exists(reference)
    
        transform_job_name = resource["spec"].get("transformJobName", None)
        assert transform_job_name is not None

        transform_sm_desc = get_sagemaker_transform_job(transform_job_name)
        assert k8s.get_resource_arn(resource) == transform_sm_desc["TransformJobArn"]
        assert transform_sm_desc["TransformJobStatus"] in TRANSFORM_JOB_STATUS_CREATED

        self._assert_transform_status_in_sync(
            transform_job_name, reference, TRANSFORM_JOB_STATUS_INPROGRESS
        )

        # Delete the k8s resource.
        _, deleted = k8s.delete_custom_resource(reference)
        assert deleted is True

        transform_sm_desc = get_sagemaker_transform_job(transform_job_name)
        assert transform_sm_desc["TransformJobStatus"] in TRANSFORM_JOB_STATUS_STOPPED

    def test_completed_transform(self, xgboost_transformjob):
        (reference, resource) = xgboost_transformjob
        assert k8s.get_resource_exists(reference)

        transform_job_name = resource["spec"].get("transformJobName", None)
        assert transform_job_name is not None

        transform_sm_desc = get_sagemaker_transform_job(transform_job_name)
        assert (
            k8s.get_resource_arn(resource) == transform_sm_desc["TransformJobArn"]
        )
        assert transform_sm_desc["TransformJobStatus"] in TRANSFORM_JOB_STATUS_CREATED
        
        self._assert_transform_status_in_sync(
            transform_job_name, reference, TRANSFORM_JOB_STATUS_SUCCESSFUL
        )

        # Check that you can delete a completed resource from k8s
        # TODO: uncomment after fix is merged
        # _, deleted = k8s.delete_custom_resource(reference)
        # assert deleted is True
