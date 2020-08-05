"""SageMaker component for deploy."""
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from dataclasses import dataclass
import logging
from typing import Dict

from typing_extensions import TypedDict

from deploy.src.sagemaker_deploy_spec import (
    SageMakerDeploySpec,
    SageMakerDeployInputs,
    SageMakerDeployOutputs,
)
from common.sagemaker_component import (
    SageMakerComponent,
    ComponentMetadata,
    SageMakerJobStatus,
)

@dataclass(frozen=True)
class EndpointRequests:
    """Holds the request types for creating each of the deploy steps."""

    config_request: Dict
    endpoint_request: Dict


@dataclass(frozen=True)
class EndpointResponses:
    """Holds the responses for creating each of the deploy step."""

    config_response: Dict
    endpoint_response: Dict


@ComponentMetadata(
    name="SageMaker - Deploy Model",
    description="Deploy Machine Learning Model Endpoint in SageMaker",
    spec=SageMakerDeploySpec,
)
class SageMakerDeployComponent(SageMakerComponent):
    """SageMaker component for deploy."""

    def Do(self, spec: SageMakerDeploySpec):
        self._endpoint_config_name = (
            spec.inputs.endpoint_config_name
            if spec.inputs.endpoint_config_name
            else "EndpointConfig" + spec.inputs.model_name_1[spec.inputs.model_name_1.index("-") :]
        )

        self._endpoint_name = (
            spec.inputs.endpoint_name
            if spec.inputs.endpoint_name
            else "Endpoint" + self._endpoint_config_name[self._endpoint_config_name.index("-") :]
        )
        super().Do(spec.inputs, spec.outputs, spec.output_paths)

    def _get_job_status(self) -> SageMakerJobStatus:
        # Job will complete instantly, unless there is a Create* exception
        return SageMakerJobStatus(is_completed=True, raw_status="")

    def _after_job_complete(
        self,
        job: object,
        request: EndpointRequests,
        inputs: SageMakerDeployInputs,
        outputs: SageMakerDeployOutputs,
    ):
        outputs.endpoint_name = self._endpoint_name

    def _create_endpoint_config_request(
        self, inputs: SageMakerDeployInputs, outputs: SageMakerDeployOutputs
    ):
        ### Documentation: https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/sagemaker.html#SageMaker.Client.create_endpoint_config
        request = self._get_request_template("endpoint_config")

        if not inputs.model_name_1:
            logging.error("Must specify at least one model (model name) to host.")
            raise Exception("Could not create endpoint config.")

        request["EndpointConfigName"] = self._endpoint_config_name

        if inputs.resource_encryption_key:
            request["KmsKeyId"] = inputs.resource_encryption_key
        else:
            request.pop("KmsKeyId")

        for i in range(len(request["ProductionVariants"]), 0, -1):
            if inputs.__dict__["model_name_" + str(i)]:
                request["ProductionVariants"][i - 1]["ModelName"] = inputs.__dict__[
                    "model_name_" + str(i)
                ]
                if inputs.__dict__["variant_name_" + str(i)]:
                    request["ProductionVariants"][i - 1][
                        "VariantName"
                    ] = inputs.__dict__["variant_name_" + str(i)]
                if inputs.__dict__["initial_instance_count_" + str(i)]:
                    request["ProductionVariants"][i - 1][
                        "InitialInstanceCount"
                    ] = inputs.__dict__["initial_instance_count_" + str(i)]
                if inputs.__dict__["instance_type_" + str(i)]:
                    request["ProductionVariants"][i - 1][
                        "InstanceType"
                    ] = inputs.__dict__["instance_type_" + str(i)]
                if inputs.__dict__["initial_variant_weight_" + str(i)]:
                    request["ProductionVariants"][i - 1][
                        "InitialVariantWeight"
                    ] = inputs.__dict__["initial_variant_weight_" + str(i)]
                if inputs.__dict__["accelerator_type_" + str(i)]:
                    request["ProductionVariants"][i - 1][
                        "AcceleratorType"
                    ] = inputs.__dict__["accelerator_type_" + str(i)]
                else:
                    request["ProductionVariants"][i - 1].pop("AcceleratorType")
            else:
                request["ProductionVariants"].pop(i - 1)

        ### Update tags
        for key, val in inputs.endpoint_config_tags.items():
            request["Tags"].append({"Key": key, "Value": val})

        return request

    def _create_endpoint_request(
        self, inputs: SageMakerDeployInputs, outputs: SageMakerDeployOutputs
    ):
        ### Documentation: https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/sagemaker.html#SageMaker.Client.create_endpoint
        request = {}
        
        request["EndpointName"] = self._endpoint_name
        request["EndpointConfigName"] = self._endpoint_config_name

        self._enable_tag_support(request, inputs)

        return request

    def _create_job_request(
        self, inputs: SageMakerDeployInputs, outputs: SageMakerDeployOutputs,
    ) -> EndpointRequests:
        return EndpointRequests(
            config_request=self._create_endpoint_config_request(inputs, outputs),
            endpoint_request=self._create_endpoint_request(inputs, outputs)
        )

    def _submit_job_request(self, request: EndpointRequests) -> EndpointResponses:
        config_response = self._sm_client.create_endpoint_config(**request.config_request)
        endpoint_response = self._sm_client.create_endpoint(**request.endpoint_request)

        return EndpointResponses(config_response=config_response, endpoint_response=endpoint_response)

    def _after_submit_job_request(
        self,
        job: EndpointResponses,
        request: EndpointRequests,
        inputs: SageMakerDeployInputs,
        outputs: SageMakerDeployOutputs,
    ):
        region = inputs.region
        logging.info(
            "Endpoint configuration in SageMaker: https://{}.console.aws.amazon.com/sagemaker/home?region={}#/endpointConfig/{}".format(
                region, region, request.config_request["EndpointConfigName"]
            )
        )
        logging.info(
            "Endpoint Config Arn: "
            + job.config_response["EndpointConfigArn"]
        )
        logging.info("Created endpoint with name: " + self._endpoint_name)
        logging.info(
            "Endpoint in SageMaker: https://{}.console.aws.amazon.com/sagemaker/home?region={}#/endpoints/{}".format(
                region, region, self._endpoint_name
            )
        )
        logging.info(
            "CloudWatch logs: https://{}.console.aws.amazon.com/cloudwatch/home?region={}#logStream:group=/aws/sagemaker/Endpoints/{};streamFilter=typeLogStreamPrefix".format(
                region, region, self._endpoint_name
            )
        )


if __name__ == "__main__":
    import sys

    spec = SageMakerDeploySpec(sys.argv[1:])

    component = SageMakerDeployComponent()
    component.Do(spec)
