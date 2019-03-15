# Copyright 2018 Google LLC
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

import re
from typing import Union, List, Any, Callable, TypeVar, Dict

from ._k8s_helper import K8sHelper
from .. import dsl

# generics
T = TypeVar('T')


def _get_pipelineparam(payload: str) -> List[str]:
    """Get a list of `PipelineParam` from a string.
    
    Args:
        payload {str}: string
    """

    matches = dsl._match_serialized_pipelineparam(payload)
    return [dsl.PipelineParam(x[1], x[0], x[2]) for x in list(set(matches))]


def _sanitize_pipelineparam(param: dsl.PipelineParam, in_place=True):
    """Sanitize the name and op_name of a PipelineParam.
  
    Args:
      params: a PipelineParam to sanitize
      in_place: if set, do an in-place update to PipelineParm, otherwise return a 
                new instance of PipelineParam.
    """
    if in_place:
        param.name = K8sHelper.sanitize_k8s_name(param.name)
        param.op_name = K8sHelper.sanitize_k8s_name(
            param.op_name) if param.op_name else param.op_name
        return param

    return dsl.PipelineParam(
        K8sHelper.sanitize_k8s_name(param.name),
        K8sHelper.sanitize_k8s_name(param.op_name), param.value)


def _sanitize_pipelineparams(
        params: Union[dsl.PipelineParam, List[dsl.PipelineParam]],
        in_place=True):
    """Sanitize the name(s) of a PipelineParam (or a list of PipelineParam) and
    return a list of sanitized PipelineParam.
  
    Args:
      params: a PipelineParam or a list of PipelineParam to sanitize
      in_place: if set, do an in-place update to the PipelineParm, otherwise return 
                new instances of PipelineParam.
    """
    params = params if isinstance(params, list) else [params]
    return [_sanitize_pipelineparam(param, in_place) for param in params]


def _process_obj(obj: Any, map_to_tmpl_var: dict):
    """Recursively sanitize and replace any PipelineParam (instances and serialized strings)
    in the object with the corresponding template variables
    (i.e. '{{inputs.parameters.<PipelineParam.full_name>}}').
    
    Args:
      obj: any obj that may have PipelineParam
      map_to_tmpl_var: a dict that maps an unsanitized pipeline
                       params signature into a template var
    """
    # serialized str might be unsanitized
    if isinstance(obj, str):
        # get signature
        pipeline_params = _get_pipelineparam(obj)
        if not pipeline_params:
            return obj
        # replace all unsanitized signature with template var
        for param in pipeline_params:
            pattern = str(param)
            sanitized = str(_sanitize_pipelineparam(param))
            obj = re.sub(pattern, map_to_tmpl_var[sanitized], obj)

    # list
    if isinstance(obj, list):
        return [_process_obj(item, map_to_tmpl_var) for item in obj]

    # tuple
    if isinstance(obj, tuple):
        return tuple((_process_obj(item, map_to_tmpl_var) for item in obj))

    # dict
    if isinstance(obj, dict):
        return {
            key: _process_obj(value, map_to_tmpl_var)
            for key, value in obj.items()
        }

    # pipelineparam
    if isinstance(obj, dsl.PipelineParam):
        # if not found in unsanitized map, then likely to be sanitized
        return map_to_tmpl_var.get(
            str(obj), '{{inputs.parameters.%s}}' % obj.full_name)

    # k8s_obj
    if hasattr(obj, 'swagger_types') and isinstance(obj.swagger_types, dict):
        # process everything inside recursively
        for key in obj.swagger_types.keys():
            setattr(obj, key, _process_obj(getattr(obj, key), map_to_tmpl_var))
        # return json representation of the k8s obj
        return K8sHelper.convert_k8s_obj_to_json(obj)

    # do nothing
    return obj


def _process_container_ops(op: dsl.ContainerOp):
    """Recursively go through the attrs listed in `attrs_with_pipelineparams` 
    and sanitize and replace pipeline params with template var string.   
    
    Returns a processed `ContainerOp`.

    NOTE this is an in-place update to `ContainerOp`'s attributes (i.e. other than
    `file_outputs`, and `outputs`, all `PipelineParam` are replaced with the 
    corresponding template variable strings).

    Args:
        op {dsl.ContainerOp}: class that inherits from ds.ContainerOp
    
    Returns:
        dsl.ContainerOp
    """

    # tmp map: unsanitized rpr -> sanitized PipelineParam
    # in-place sanitize of all PipelineParam (except outputs and file_outputs)
    _map = {
        str(param): _sanitize_pipelineparam(param, in_place=True)
        for param in op.inputs
    }

    # map: unsanitized pipeline param rpr -> template var string
    # used to replace unsanitized pipeline param strings with the corresponding
    # template var strings
    map_to_tmpl_var = {
        key: '{{inputs.parameters.%s}}' % param.full_name
        for key, param in _map.items()
    }

    # process all attr with pipelineParams except inputs and outputs parameters
    for key in op.attrs_with_pipelineparams:
        setattr(op, key, _process_obj(getattr(op, key), map_to_tmpl_var))

    return op


def _parameters_to_json(params: List[dsl.PipelineParam]):
    """Converts a list of PipelineParam into an argo `parameter` JSON obj."""
    _to_json = (lambda param: dict(name=param.full_name, value=param.value)
                if param.value else dict(name=param.full_name))
    params = [_to_json(param) for param in params]
    # Sort to make the results deterministic.
    params.sort(key=lambda x: x['name'])
    return params


# TODO: artifacts?
def _inputs_to_json(inputs_params: List[dsl.PipelineParam], _artifacts=None):
    """Converts a list of PipelineParam into an argo `inputs` JSON obj."""
    parameters = _parameters_to_json(inputs_params)
    return {'parameters': parameters} if parameters else None


def _outputs_to_json(outputs: Dict[str, dsl.PipelineParam],
                     file_outputs: Dict[str, str],
                     output_artifacts: List[dict]):
    """Creates an argo `outputs` JSON obj."""
    output_parameters = []
    for param in outputs.values():
        output_parameters.append({
            'name': param.full_name,
            'valueFrom': {
                'path': file_outputs[param.name]
            }
        })
    output_parameters.sort(key=lambda x: x['name'])
    ret = {}
    if output_parameters:
        ret['parameters'] = output_parameters
    if output_artifacts:
        ret['artifacts'] = output_artifacts

    return ret


def _build_conventional_artifact(name):
    return {
        'name': name,
        'path': '/' + name + '.json',
        's3': {
            # TODO: parameterize namespace for minio service
            'endpoint': 'minio-service.kubeflow:9000',
            'bucket': 'mlpipeline',
            'key': 'runs/{{workflow.uid}}/{{pod.name}}/' + name + '.tgz',
            'insecure': True,
            'accessKeySecret': {
                'name': 'mlpipeline-minio-artifact',
                'key': 'accesskey',
            },
            'secretKeySecret': {
                'name': 'mlpipeline-minio-artifact',
                'key': 'secretkey'
            }
        },
    }


# TODO: generate argo python classes from swagger and use convert_k8s_obj_to_json??
def _op_to_template(op: dsl.ContainerOp):
    """Generate template given an operator inherited from dsl.ContainerOp."""

    # NOTE in-place update to ContainerOp
    # replace all PipelineParams (except in `file_outputs`, `outputs`, `inputs`)
    # with template var strings
    processed_op = _process_container_ops(op)

    # default output artifacts
    output_artifacts = [
        _build_conventional_artifact(name)
        for name in ['mlpipeline-ui-metadata', 'mlpipeline-metrics']
    ]

    # workflow template
    template = {
        'name': op.name,
        'container': K8sHelper.convert_k8s_obj_to_json(op.container)
    }

    # inputs
    inputs = _inputs_to_json(processed_op.inputs)
    if inputs:
        template['inputs'] = inputs

    # outputs
    template['outputs'] = _outputs_to_json(op.outputs, op.file_outputs,
                                           output_artifacts)

    # node selector
    if processed_op.node_selector:
        template['nodeSelector'] = processed_op.node_selector

    # metadata
    if processed_op.pod_annotations or processed_op.pod_labels:
        template['metadata'] = {}
        if processed_op.pod_annotations:
            template['metadata']['annotations'] = processed_op.pod_annotations
        if processed_op.pod_labels:
            template['metadata']['labels'] = processed_op.pod_labels
    # retries
    if processed_op.num_retries:
        template['retryStrategy'] = {'limit': processed_op.num_retries}

    # sidecars
    if processed_op.sidecars:
        template['sidecars'] = processed_op.sidecars

    return template
