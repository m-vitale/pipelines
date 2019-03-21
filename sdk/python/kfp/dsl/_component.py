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

from ._metadata import ComponentMeta, ParameterMeta, TypeMeta, _annotation_to_typemeta
from ._pipeline_param import PipelineParam
from ._types import check_types, InconsistentTypeException
from ._ops_group import OpsGroup, Graph
from . import _pipeline
import kfp

def python_component(name, description=None, base_image=None, target_component_file: str = None):
  """Decorator for Python component functions.
  This decorator adds the metadata to the function object itself.

  Args:
      name: Human-readable name of the component
      description: Optional. Description of the component
      base_image: Optional. Docker container image to use as the base of the component. Needs to have Python 3.5+ installed.
      target_component_file: Optional. Local file to store the component definition. The file can then be used for sharing.

  Returns:
      The same function (with some metadata fields set).

  Usage:
  ```python
  @dsl.python_component(
    name='my awesome component',
    description='Come, Let's play',
    base_image='tensorflow/tensorflow:1.11.0-py3',
  )
  def my_component(a: str, b: int) -> str:
    ...
  ```
  """
  def _python_component(func):
    func._component_human_name = name
    if description:
      func._component_description = description
    if base_image:
      func._component_base_image = base_image
    if target_component_file:
      func._component_target_component_file = target_component_file
    return func

  return _python_component

def component(func):
  """Decorator for component functions that returns a ContainerOp.
  This is useful to enable type checking in the DSL compiler

  Usage:
  ```python
  @dsl.component
  def foobar(model: TFModel(), step: MLStep()):
    return dsl.ContainerOp()
  """
  from functools import wraps
  @wraps(func)
  def _component(*args, **kargs):
    import inspect
    fullargspec = inspect.getfullargspec(func)
    annotations = fullargspec.annotations

    # defaults
    arg_defaults = {}
    if fullargspec.defaults:
      for arg, default in zip(reversed(fullargspec.args), reversed(fullargspec.defaults)):
        arg_defaults[arg] = default

    # Construct the ComponentMeta
    component_meta = ComponentMeta(name=func.__name__, description='')
    # Inputs
    for arg in fullargspec.args:
      arg_type = TypeMeta()
      arg_default = arg_defaults[arg] if arg in arg_defaults else None
      if arg in annotations:
        arg_type = _annotation_to_typemeta(annotations[arg])
      component_meta.inputs.append(ParameterMeta(name=arg, description='', param_type=arg_type, default=arg_default))
    # Outputs
    if 'return' in annotations:
      for output in annotations['return']:
        arg_type = _annotation_to_typemeta(annotations['return'][output])
        component_meta.outputs.append(ParameterMeta(name=output, description='', param_type=arg_type))

    #TODO: add descriptions to the metadata
    #docstring parser:
    #  https://github.com/rr-/docstring_parser
    #  https://github.com/terrencepreilly/darglint/blob/master/darglint/parse.py

    if kfp.TYPE_CHECK:
      arg_index = 0
      for arg in args:
        if isinstance(arg, PipelineParam) and not check_types(arg.param_type.to_dict_or_str(), component_meta.inputs[arg_index].param_type.to_dict_or_str()):
          raise InconsistentTypeException('Component "' + component_meta.name + '" is expecting ' + component_meta.inputs[arg_index].name +
                                          ' to be type(' + component_meta.inputs[arg_index].param_type.serialize() +
                                          '), but the passed argument is type(' + arg.param_type.serialize() + ')')
        arg_index += 1
      if kargs is not None:
        for key in kargs:
          if isinstance(kargs[key], PipelineParam):
            for input_spec in component_meta.inputs:
              if input_spec.name == key and not check_types(kargs[key].param_type.to_dict_or_str(), input_spec.param_type.to_dict_or_str()):
                raise InconsistentTypeException('Component "' + component_meta.name + '" is expecting ' + input_spec.name +
                                                ' to be type(' + input_spec.param_type.serialize() +
                                                '), but the passed argument is type(' + kargs[key].param_type.serialize() + ')')

    container_op = func(*args, **kargs)
    container_op._set_metadata(component_meta)
    return container_op

  return _component

def graph_component(func):
  """Decorator for graph component functions.
  This decorator returns an ops_group """
  #TODO: add usage
  from functools import wraps
  @wraps(func)
  def _graph_component(*args, **kargs):
    graph_ops_group = Graph(func.__name__)
    graph_ops_group.inputs = list(args) + list(kargs.values())
    #TODO: check if the inputs is a list of pipelineparams

    graph_ops_group.resolve_recursion()

    # Entering Graph Context
    _pipeline.Pipeline.get_default_pipeline().push_ops_group(graph_ops_group)

    # Call the function
    if not graph_ops_group.is_recursive:
      graph_ops_group._make_name_unique()
      graph_ops_group.outputs = func(*args, **kargs)
      #TODO: check if the outputs is a dictionary of str to pipelineparams

    # Exiting Graph Context
    _pipeline.Pipeline.get_default_pipeline().pop_ops_group()

    return graph_ops_group
  return _graph_component