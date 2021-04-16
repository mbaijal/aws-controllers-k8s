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
"""Integration tests for the SageMaker HyperParameterTuning API.
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

RESOURCE_PLURAL = "hyperparametertuningjobs"
HPO_JOB_STATUS_CREATED = ("InProgress", "Completed")
HPO_JOB_STATUS_STOPPED = ("Stopped", "Stopping")
HPO_JOB_STATUS_SUCCESSFUL = "Completed"
HPO_JOB_STATUS_INPROGRESS = "InProgress"

def _sagemaker_client():
    return boto3.client('sagemaker')

def _make_hpojob():
    resource_name = random_suffix_name("xgboost-hpojob", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["HPO_JOB_NAME"] = resource_name

    data = load_resource_file(
        SERVICE_NAME, "xgboost_hpojob", additional_replacements=replacements
    )
    logging.debug(data)

    reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL, resource_name, namespace="default"
    )

    return reference, data

@pytest.fixture(scope="function")
def xgboost_hpojob():
    hpo_job, data = _make_hpojob()
    resource = k8s.create_custom_resource(hpo_job, data)
    resource = k8s.wait_resource_consumed_by_controller(hpo_job)

    yield (hpo_job, resource) 

    if k8s.get_resource_exists(hpo_job):
        k8s.delete_custom_resource(hpo_job)

def get_sagemaker_hpo_job(hpo_job_name: str):
    try:
        hpo_desc = _sagemaker_client().describe_hyper_parameter_tuning_job(
            HyperParameterTuningJobName=hpo_job_name
        )
        return hpo_desc
    except BaseException:
        logging.error(
            f"SageMaker could not find an hpo job with the name {hpo_job_name}"
        )
        return None

@service_marker
@pytest.mark.canary
class TestHPO:
    def _wait_resource_hpo_status(
        self,
        reference: k8s.CustomResourceReference,
        expected_status: str,
        wait_periods: int = 30,
    ):
        resource_status = None
        for _ in range(wait_periods):
            time.sleep(30)
            resource = k8s.get_resource(reference)
            resource_status = resource["status"]["hyperParameterTuningJobStatus"]
            if resource_status == expected_status:
                break
        else:
            logging.error(
                f"Wait for hyperparameterTuningJobStatus: {expected_status} timed out. Actual status: {resource_status}"
            )

        return resource_status

    def _wait_sagemaker_hpo_status(
        self,
        hpo_job_name,
        expected_status: str,
        wait_periods: int = 30,
    ):
        actual_status = None
        for _ in range(wait_periods):
            time.sleep(30)
            hpo_sm_desc = get_sagemaker_hpo_job(hpo_job_name)
            actual_status = hpo_sm_desc["HyperParameterTuningJobStatus"]
            if actual_status == expected_status:
                break
        else:
            logging.error(
                f"Wait for sagemaker hpo status: {expected_status} timed out. Actual status: {actual_status}"
            )

        return actual_status

    def _assert_hpo_status_in_sync(
        self, hpo_job_name, reference, expected_status
    ):
        assert (
            self._wait_sagemaker_hpo_status(
                hpo_job_name, expected_status
            )
            == self._wait_resource_hpo_status(reference, expected_status)
            == expected_status
        )

    def test_hpo(self, xgboost_hpojob):
        (reference, resource) = xgboost_hpojob
        assert k8s.get_resource_exists(reference)
    
        hpo_job_name = resource["spec"].get("hyperParameterTuningJobName", None)
        assert hpo_job_name is not None

        hpo_sm_desc = get_sagemaker_hpo_job(hpo_job_name)
        assert k8s.get_resource_arn(resource) == hpo_sm_desc["HyperParameterTuningJobArn"]
        assert hpo_sm_desc["HyperParameterTuningJobStatus"] in HPO_JOB_STATUS_CREATED

        self._assert_hpo_status_in_sync(
            hpo_job_name, reference, HPO_JOB_STATUS_INPROGRESS
        )

        # Delete the k8s resource.
        _, deleted = k8s.delete_custom_resource(reference)
        assert deleted is True

        hpo_sm_desc = get_sagemaker_hpo_job(hpo_job_name)
        assert hpo_sm_desc["HyperParameterTuningJobStatus"] in HPO_JOB_STATUS_STOPPED

    def test_completed_hpo(self, xgboost_hpojob):
        (reference, resource) = xgboost_hpojob
        assert k8s.get_resource_exists(reference)

        hpo_job_name = resource["spec"].get("hyperParameterTuningJobName", None)
        assert hpo_job_name is not None

        hpo_sm_desc = get_sagemaker_hpo_job(hpo_job_name)
        assert (
            k8s.get_resource_arn(resource) == hpo_sm_desc["HyperParameterTuningJobArn"]
        )
        assert hpo_sm_desc["HyperParameterTuningJobStatus"] in HPO_JOB_STATUS_CREATED
        
        self._assert_hpo_status_in_sync(
            hpo_job_name, reference, HPO_JOB_STATUS_SUCCESSFUL
        )

        # Check that you can delete a completed resource from k8s
        # TODO: uncomment after fix is merged
        # _, deleted = k8s.delete_custom_resource(reference)
        # assert deleted is True
