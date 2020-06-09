# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kfp_server_api.configuration import Configuration


class ApiPipelineRuntime(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'pipeline_manifest': 'str',
        'workflow_manifest': 'str'
    }

    attribute_map = {
        'pipeline_manifest': 'pipeline_manifest',
        'workflow_manifest': 'workflow_manifest'
    }

    def __init__(self, pipeline_manifest=None, workflow_manifest=None, local_vars_configuration=None):  # noqa: E501
        """ApiPipelineRuntime - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._pipeline_manifest = None
        self._workflow_manifest = None
        self.discriminator = None

        if pipeline_manifest is not None:
            self.pipeline_manifest = pipeline_manifest
        if workflow_manifest is not None:
            self.workflow_manifest = workflow_manifest

    @property
    def pipeline_manifest(self):
        """Gets the pipeline_manifest of this ApiPipelineRuntime.  # noqa: E501

        Output. The runtime JSON manifest of the pipeline, including the status of pipeline steps and fields need for UI visualization etc.  # noqa: E501

        :return: The pipeline_manifest of this ApiPipelineRuntime.  # noqa: E501
        :rtype: str
        """
        return self._pipeline_manifest

    @pipeline_manifest.setter
    def pipeline_manifest(self, pipeline_manifest):
        """Sets the pipeline_manifest of this ApiPipelineRuntime.

        Output. The runtime JSON manifest of the pipeline, including the status of pipeline steps and fields need for UI visualization etc.  # noqa: E501

        :param pipeline_manifest: The pipeline_manifest of this ApiPipelineRuntime.  # noqa: E501
        :type: str
        """

        self._pipeline_manifest = pipeline_manifest

    @property
    def workflow_manifest(self):
        """Gets the workflow_manifest of this ApiPipelineRuntime.  # noqa: E501

        Output. The runtime JSON manifest of the argo workflow. This is deprecated after pipeline_runtime_manifest is in use.  # noqa: E501

        :return: The workflow_manifest of this ApiPipelineRuntime.  # noqa: E501
        :rtype: str
        """
        return self._workflow_manifest

    @workflow_manifest.setter
    def workflow_manifest(self, workflow_manifest):
        """Sets the workflow_manifest of this ApiPipelineRuntime.

        Output. The runtime JSON manifest of the argo workflow. This is deprecated after pipeline_runtime_manifest is in use.  # noqa: E501

        :param workflow_manifest: The workflow_manifest of this ApiPipelineRuntime.  # noqa: E501
        :type: str
        """

        self._workflow_manifest = workflow_manifest

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ApiPipelineRuntime):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ApiPipelineRuntime):
            return True

        return self.to_dict() != other.to_dict()
