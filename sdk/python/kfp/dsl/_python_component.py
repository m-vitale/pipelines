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
  """Decorator for component functions that use ContainerOp.
  This is useful to enable type checking in the DSL compiler

  Usage:
  ```python
  @dsl.component
  def foobar(model: TFModel(), step: MLStep()):
    return dsl.ContainerOp()
  """
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
      arg_default = arg_defaults[arg] if arg in arg_defaults else ''
      if arg in annotations:
        arg_type = _annotation_to_typemeta(annotations[arg])
      component_meta.inputs.append(ParameterMeta(name=arg, description='', param_type=arg_type, default=arg_default))
    # Outputs
    for output in annotations['return']:
      arg_type = _annotation_to_typemeta(annotations['return'][output])
      component_meta.outputs.append(ParameterMeta(name=output, description='', param_type=arg_type))

    #TODO: add descriptions to the metadata
    #docstring parser:
    #  https://github.com/rr-/docstring_parser
    #  https://github.com/terrencepreilly/darglint/blob/master/darglint/parse.py

    print(component_meta.serialize())
    #TODO: parse the metadata to the ContainerOp.
    container_op = func(*args, **kargs)
    return container_op

  return _component