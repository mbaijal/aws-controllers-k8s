# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.
"""Integration tests for the SageMaker ProcessingJob API.
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

RESOURCE_PLURAL = "processingjobs"
PROCESSING_JOB_STATUS_CREATED = ("InProgress", "Completed")
PROCESSING_JOB_STATUS_STOPPED = ("Stopped", "Stopping")
PROCESSING_JOB_STATUS_INPROGRESS = "InProgress"
PROCESSING_JOB_STATUS_SUCCESSFUL = "Completed"

def _sagemaker_client():
    return boto3.client("sagemaker")

def _make_processing_job():
    resource_name = random_suffix_name("kmeans-processingjob", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["PROCESSING_JOB_NAME"] = resource_name

    data = load_resource_file(
        SERVICE_NAME, "kmeans_processingjob", additional_replacements=replacements
    )
    logging.debug(data)

    # Create the k8s resource
    reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL, resource_name, namespace="default"
    )

    return reference, data

@pytest.fixture(scope="function")
def kmeans_processing_job():
    (processing_job, data) = _make_processing_job()
    resource = k8s.create_custom_resource(processing_job, data)
    resource = k8s.wait_resource_consumed_by_controller(processing_job)

    yield (processing_job, resource) 

    if k8s.get_resource_exists(processing_job):
        k8s.delete_custom_resource(processing_job)

def get_sagemaker_processing_job(processing_job_name: str):
    try:
        processing_job = _sagemaker_client().describe_processing_job(
            ProcessingJobName=processing_job_name
        )
        return processing_job
    except BaseException:
        logging.error(
            f"SageMaker could not find a processing job with the name {processing_job_name}"
        )
        return None

@service_marker
@pytest.mark.canary
class TestProcessingJob:
    def _wait_resource_processing_status(
        self,
        reference: k8s.CustomResourceReference,
        expected_status: str,
        wait_periods: int = 30,
    ):
        resource_status = None
        for _ in range(wait_periods):
            time.sleep(30)
            resource = k8s.get_resource(reference)
            resource_status = resource["status"]["processingJobStatus"]
            if resource_status == expected_status:
                break
        else:
            logging.error(
                f"Wait for ProcessingJobStatus: {expected_status} timed out. Actual status: {resource_status}"
            )

        return resource_status

    def _wait_sagemaker_processing_status(
        self,
        processing_job_name,
        expected_status: str,
        wait_periods: int = 30,
    ):
        actual_status = None
        for _ in range(wait_periods):
            time.sleep(30)
            processing_sm_desc = get_sagemaker_processing_job(processing_job_name)
            actual_status = processing_sm_desc["ProcessingJobStatus"]
            if actual_status == expected_status:
                break
        else:
            logging.error(
                f"Wait for sagemaker processing status: {expected_status} timed out. Actual status: {actual_status}"
            )

        return actual_status

    def _assert_processing_status_in_sync(
        self, processing_job_name, reference, expected_status
    ):
        assert (
            self._wait_sagemaker_processing_status(
                processing_job_name, expected_status
            )
            == self._wait_resource_processing_status(reference, expected_status)
            == expected_status
        )

    def test_processing_job(self, kmeans_processing_job):
        (reference, resource) = kmeans_processing_job
        assert k8s.get_resource_exists(reference)

        processing_job_name = resource["spec"].get("processingJobName", None)
        assert processing_job_name is not None

        processing_job_desc = get_sagemaker_processing_job(processing_job_name)

        assert k8s.get_resource_arn(resource) == processing_job_desc["ProcessingJobArn"]
        assert processing_job_desc["ProcessingJobStatus"] in PROCESSING_JOB_STATUS_CREATED

        self._assert_processing_status_in_sync(
            processing_job_name, reference, PROCESSING_JOB_STATUS_INPROGRESS
        )

        # Delete the k8s resource.
        _, deleted = k8s.delete_custom_resource(reference)
        assert deleted is True

        processing_job_desc = get_sagemaker_processing_job(processing_job_name)
        assert processing_job_desc["ProcessingJobStatus"] in PROCESSING_JOB_STATUS_STOPPED

    def test_completed_processing_job(self, kmeans_processing_job):
        (reference, resource) = kmeans_processing_job
        assert k8s.get_resource_exists(reference)

        processing_job_name = resource["spec"].get("processingJobName", None)
        assert processing_job_name is not None

        processing_job_desc = get_sagemaker_processing_job(processing_job_name)

        assert k8s.get_resource_arn(resource) == processing_job_desc["ProcessingJobArn"]
        assert processing_job_desc["ProcessingJobStatus"] in PROCESSING_JOB_STATUS_CREATED

        self._assert_processing_status_in_sync(
            processing_job_name, reference, PROCESSING_JOB_STATUS_SUCCESSFUL
        )

        # Check that you can delete a completed resource from k8s
        # TODO: uncomment after fix is merged
        # _, deleted = k8s.delete_custom_resource(reference)
        # assert deleted is True
