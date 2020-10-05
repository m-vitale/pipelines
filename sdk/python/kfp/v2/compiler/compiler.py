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
"""KFP DSL v2 compiler.

This is an experimental implementation of KFP compiler that compiles KFP
pipeline into Pipeline IR:
https://docs.google.com/document/d/1PUDuSQ8vmeKSBloli53mp7GIvzekaY7sggg6ywy35Dk/
"""

import inspect
from typing import Any, Callable, List, Optional

import kfp
from kfp.compiler._k8s_helper import sanitize_k8s_name
from kfp.components import _python_op
from kfp.components import structures
from kfp.v2 import dsl
from kfp.v2.compiler import importer_node
from kfp.v2.dsl import type_utils
from kfp.v2.proto import pipeline_spec_pb2

from google.protobuf.json_format import MessageToJson


class Compiler(object):
  """Experimental DSL compiler that targets the PipelineSpec IR.

  It compiles pipeline function into PipelineSpec json string.
  PipelineSpec is the IR protobuf message that defines a pipeline:
  https://github.com/kubeflow/pipelines/blob/237795539f7b85bac77435e2464367226ee19391/api/v2alpha1/pipeline_spec.proto#L8
  In this initial implementation, we only support components authored through
  Component yaml spec. And we don't support advanced features like conditions,
  static and dynamic loops, etc.

  Example:
    How to use the compiler to construct pipeline_spec json:

      @dsl.pipeline(
        name='name',
        description='description'
      )
      def my_pipeline(a: int = 1, b: str = "default value"):
        ...

      kfp.v2.compiler.Compiler().compile(my_pipeline, 'path/to/pipeline.json')
  """

  def _create_pipeline_spec(
      self,
      args: List[dsl.PipelineParam],
      pipeline: dsl.Pipeline,
  ) -> pipeline_spec_pb2.PipelineSpec:
    """Creates the pipeline spec object.

    Args:
      args: The list of pipeline arguments.
      pipeline: The instantiated pipeline object.

    Returns:
      A PipelineSpec proto representing the compiled pipeline.

    Raises:
      NotImplementedError if the argument is of unsupported types.
    """
    pipeline_spec = pipeline_spec_pb2.PipelineSpec()
    if not pipeline.name:
      raise ValueError('Pipeline name is required.')
    pipeline_spec.pipeline_info.name = pipeline.name
    pipeline_spec.sdk_version = kfp.__version__
    pipeline_spec.schema_version = 'v2alpha1'

    # Pipeline Parameters
    for arg in args:
      if arg.value is not None:
        if isinstance(arg.value, int):
          pipeline_spec.runtime_parameters[
              arg.name].type = pipeline_spec_pb2.PrimitiveType.INT
          pipeline_spec.runtime_parameters[
              arg.name].default_value.int_value = arg.value
        elif isinstance(arg.value, float):
          pipeline_spec.runtime_parameters[
              arg.name].type = pipeline_spec_pb2.PrimitiveType.DOUBLE
          pipeline_spec.runtime_parameters[
              arg.name].default_value.double_value = arg.value
        elif isinstance(arg.value, str):
          pipeline_spec.runtime_parameters[
              arg.name].type = pipeline_spec_pb2.PrimitiveType.STRING
          pipeline_spec.runtime_parameters[
              arg.name].default_value.string_value = arg.value
        else:
          raise NotImplementedError(
              'Unexpected parameter type with: "{}".'.format(str(arg.value)))

    def _get_input_artifact_type(
        input_name: str,
        inputs: List[structures.InputSpec],
    ) -> str:
      """Find the input artifact type by input name.

      Args:
        input_name: The name of the component input.
        inputs: The list of InputSpec

      Returns:
        The input type name if found in inputs, or '' if not found.
      """
      for component_input in inputs:
        if component_input.name == input_name:
          return type_utils.get_artifact_type(component_input.type)
      return ''

    deployment_config = pipeline_spec_pb2.PipelineDeploymentConfig()
    importer_tasks = []

    for op in pipeline.ops.values():
      component_spec = op._metadata
      task = pipeline_spec.tasks.add()
      task.CopyFrom(op.task_spec)
      deployment_config.executors[task.executor_label].container.CopyFrom(
          op.container_spec)

      # Check if need to insert importer node
      for input_name in task.inputs.artifacts:
        if not task.inputs.artifacts[input_name].producer_task:
          artifact_type = _get_input_artifact_type(input_name,
                                                   component_spec.inputs)

          importer_task, importer_spec = importer_node.build_importer_spec(
              task, input_name, artifact_type)
          importer_tasks.append(importer_task)

          task.inputs.artifacts[
              input_name].producer_task = importer_task.task_info.name
          task.inputs.artifacts[
              input_name].output_artifact_key = importer_node.OUTPUT_KEY

          deployment_config.executors[
              importer_task.executor_label].importer.CopyFrom(importer_spec)

    pipeline_spec.deployment_config.Pack(deployment_config)
    pipeline_spec.tasks.extend(importer_tasks)

    return pipeline_spec

  def _create_pipeline(
      self,
      pipeline_func: Callable[..., Any],
      pipeline_name: Optional[str] = None,
      pipeline_description: Optional[str] = None,
      params_list: List[dsl.PipelineParam] = None,
  ) -> pipeline_spec_pb2.PipelineSpec:
    """Internal implementation of create_pipeline.

    Args:
      pipeline_func: Pipeline function with @dsl.pipeline decorator.
      pipeline_name: The name of the pipeline. Optional.
      pipeline_description: An optional description of the pipeline.
      params_list: A list of pipeline arguments.

    Returns:
      The IR representation (pipeline spec) of the pipeline.
    """
    params_list = params_list or []

    # Create the arg list with no default values and call pipeline function.
    # Assign type information to the PipelineParam
    pipeline_meta = _python_op._extract_component_interface(pipeline_func)
    pipeline_meta.name = pipeline_name or pipeline_meta.name
    pipeline_meta.description = pipeline_description or pipeline_meta.description
    pipeline_name = sanitize_k8s_name(pipeline_meta.name)

    # Need to first clear the default value of dsl.PipelineParams. Otherwise, it
    # will be resolved immediately in place when being to each component.
    default_param_values = {}
    for param in params_list:
      default_param_values[param.name] = param.value
      param.value = None

    # Currently only allow specifying pipeline params at one place.
    if params_list and pipeline_meta.inputs:
      raise ValueError(
          'Either specify pipeline params in the pipeline function, or in '
          '"params_list", but not both.')

    args_list = []
    signature = inspect.signature(pipeline_func)
    for arg_name in signature.parameters:
      arg_type = None
      for pipeline_input in pipeline_meta.inputs or []:
        if arg_name == pipeline_input.name:
          arg_type = pipeline_input.type
          break
      args_list.append(
          dsl.PipelineParam(
              sanitize_k8s_name(arg_name, True), param_type=arg_type))

    with dsl.Pipeline(pipeline_name) as dsl_pipeline:
      pipeline_func(*args_list)

    # Fill in the default values.
    args_list_with_defaults = []
    if pipeline_meta.inputs:
      args_list_with_defaults = [
          dsl.PipelineParam(
              sanitize_k8s_name(input_spec.name, True),
              value=input_spec.default) for input_spec in pipeline_meta.inputs
      ]
    elif params_list:
      # Or, if args are provided by params_list, fill in pipeline_meta.
      for param in params_list:
        param.value = default_param_values[param.name]

      args_list_with_defaults = params_list
      pipeline_meta.inputs = [
          structures.InputSpec(
              name=param.name, type=param.param_type, default=param.value)
          for param in params_list
      ]

    pipeline_spec = self._create_pipeline_spec(
        args_list_with_defaults,
        dsl_pipeline,
    )

    return pipeline_spec

  def compile(self,
              pipeline_func: Callable[..., Any],
              output_path: str,
              type_check: bool = True) -> None:
    """Compile the given pipeline function into workflow yaml.

    Args:
      pipeline_func: Pipeline function with @dsl.pipeline decorator.
      output_path: The output pipeline spec .json file path. for example,
        "~/a.json"
      type_check: Whether to enable the type check or not, default: True.
    """
    type_check_old_value = kfp.TYPE_CHECK
    try:
      kfp.TYPE_CHECK = type_check
      self._create_and_write_pipeline_spec(
          pipeline_func=pipeline_func, output_path=output_path)
    finally:
      kfp.TYPE_CHECK = type_check_old_value

  def _write_pipeline(self,
                      pipeline_spec: pipeline_spec_pb2.PipelineSpec,
                      output_path: str = None) -> Optional[str]:
    """Dump pipeline spec into json file.

    Args:
      pipeline_spec: IR pipeline spec.
      ouput_path: The file path to be written. If not specified, a json_text
        string will be returned.

    Returns:
      The json representation of pipeline_spec if no package_path specified.

    Raises:
      ValueError: if the specified output path doesn't end with the acceptable
      extentions.
    """
    json_text = MessageToJson(pipeline_spec)

    if output_path is None:
      return json_text

    if output_path.endswith('.json'):
      with open(output_path, 'w') as json_file:
        json_file.write(json_text)
    else:
      raise ValueError(
          'The output path {} should ends with ".json".'.format(output_path))

  def _create_and_write_pipeline_spec(self,
                                      pipeline_func: Callable[..., Any],
                                      pipeline_name: str = None,
                                      pipeline_description: str = None,
                                      params_list: List[
                                          dsl.PipelineParam] = None,
                                      output_path: str = None) -> None:
    """Compile the given pipeline function and dump it to specified file format.

    Args:
      pipeline_func: Pipeline function with @dsl.pipeline decorator.
      pipeline_name: The name of the pipeline. Optional.
      pipeline_description: An optional description of the pipeline.
      params_list: A list of pipeline arguments.
      output_path: The path of the compiled output.
    """
    pipeline = self._create_pipeline(pipeline_func, pipeline_name,
                                     pipeline_description, params_list)
    self._write_pipeline(pipeline, output_path)
