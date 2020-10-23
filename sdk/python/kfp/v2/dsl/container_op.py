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
"""IR-based ContainerOp."""

from typing import Callable, Text

from kfp import dsl
from kfp.v2.proto import pipeline_spec_pb2

# Unit constants for k8s size string.
_E = 10 ** 18  # Exa
_EI = 1 << 60  # Exa: power-of-two approximate
_P = 10 ** 15  # Peta
_PI = 1 << 50  # Peta: power-of-two approximate
# noinspection PyShadowingBuiltins
_T = 10 ** 12  # Tera
_TI = 1 << 40  # Tera: power-of-two approximate
_G = 10 ** 9  # Giga
_GI = 1 << 30  # Giga: power-of-two approximate
_M = 10 ** 6  # Mega
_MI = 1 << 20  # Mega: power-of-two approximate
_K = 10 ** 3  # Kilo
_KI = 1 << 10  # Kilo: power-of-two approximate


def resource_setter(func: Callable):
  """Function decorator for common validation before setting resource spec."""

  def resource_setter_wrapper(container_op: 'ContainerOp', *args, **kwargs):
    # Validate the container_op has right format of container_spec set.
    if not hasattr(container_op, 'container_spec'):
      raise ValueError('Expecting container_spec attribute of the container_op:'
                       ' {}'.format(container_op))
    if not isinstance(
        container_op.container_spec,
        pipeline_spec_pb2.PipelineDeploymentConfig.PipelineContainerSpec):
      raise TypeError('ContainerOp.container_spec is expected to be a '
                      'PipelineContainerSpec proto. Got: {} for {}'.format(
          type(container_op.container_spec), container_op.container_spec))
    # Run the resource setter function
    func(container_op, *args, **kwargs)

  return resource_setter_wrapper


def _get_cpu_number(cpu_string: Text) -> float:
  """Converts the cpu string to number of vCPU core."""
  # dsl.ContainerOp._validate_cpu_string guaranteed that cpu_string is either
  # 1) a string can be converted to a float; or
  # 2) a string followed by 'm', and it can be converted to a float.
  if cpu_string.endswith('m'):
    return float(cpu_string[:-1]) / 1000
  else:
    return float(cpu_string)


def _get_memory_number(memory_string: Text) -> float:
  """Converts the memory string to number of memory in GBi."""
  # dsl.ContainerOp._validate_size_string guaranteed that memory_string
  # represents an integer, optionally followed by one of (E, Ei, P, Pi, T, Ti,
  # G, Gi, M, Mi, K, Ki).
  # See the meaning of different suffix at
  # https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-memory
  # Also, ResourceSpec in pipeline IR expects a number in GB.
  if memory_string.endswith('E'):
    return float(memory_string) * _E / _G
  elif memory_string.endswith('Ei'):
    return float(memory_string) * _EI / _G
  elif memory_string.endswith('P'):
    return float(memory_string) * _P / _G
  elif memory_string.endswith('Pi'):
    return float(memory_string) * _PI / _G
  elif memory_string.endswith('T'):
    return float(memory_string) * _T / _G
  elif memory_string.endswith('Ti'):
    return float(memory_string) * _TI / _G
  elif memory_string.endswith('G'):
    return float(memory_string)
  elif memory_string.endswith('Gi'):
    return float(memory_string) * _GI / _G
  elif memory_string.endswith('M'):
    return float(memory_string) * _M / _G
  elif memory_string.endswith('Mi'):
    return float(memory_string) * _MI / _G
  elif memory_string.endswith('K'):
    return float(memory_string) * _K / _G
  elif memory_string.endswith('Ki'):
    return float(memory_string) * _KI / _G
  else:
    # By default interpret as a plain integer, in the unit of Bytes.
    return float(memory_string) / _G


def _sanitize_gpu_type(gpu_type: Text) -> Text:
  """Converts the GPU type to conform the enum style."""
  return gpu_type.replace('-', '_').upper()


class ContainerOp(dsl.ContainerOp):
  """V2 ContainerOp class.

  This class inherits an almost identical behavior as the previous ContainerOp
  class. The diffs are in two aspects:
  - The source of truth is migrating to the PipelineContainerSpec proto.
  - The implementation (and impact) of several APIs are different. For example,
    resource spec will be set in the pipeline IR proto instead of using k8s API.
  """

  def __init__(self, **kwargs):
    super(ContainerOp, self).__init__(**kwargs)

  # Override resource specification calls.
  @resource_setter
  def set_cpu_limit(self, cpu: Text) -> 'ContainerOp':
    """Sets the cpu provisioned for this task.
    
    Args:
      cpu: a string indicating the amount of vCPU required by this task.
        Please refer to dsl.ContainerOp._validate_cpu_string regarding
        its format.
    Returns:
      self return to allow chained call with other resource specification.
    """
    self.container._validate_cpu_string(cpu)
    self.container_spec.resources.cpu_limit = _get_cpu_number(cpu)
    return self

  @resource_setter
  def set_memory_limit(self, memory: Text) -> 'ContainerOp':
    """Sets the memory provisioned for this task.
    
    Args:
      memory: a string described the amount of memory required by this
        task. Please refer to dsl.ContainerOp._validate_size_string
        regarding its format.
    Returns:
      self return to allow chained call with other resource specification.
    """
    self.container._validate_size_string(memory)
    self.container_spec.resources.memory_limit = _get_memory_number(memory)
    return self

  @resource_setter
  def add_node_selector_constraint(
      self, label_name: Text, value: Text
  ) -> 'ContainerOp':
    """Sets accelerator type requirement for this task.
    
    This function is designed to enable users to specify accelerator using
    a similar DSL syntax as KFP V1. Under the hood, it will directly specify
    the accelerator required in the IR proto, instead of relying on the
    k8s node selector API.

    This function can be optionally used with set_gpu_limit to set the number
    of accelerator required. Otherwise, by default the number requested will be
    1.
    
    Args:
      label_name: only support 'cloud.google.com/gke-accelerator' now.
      value: name of the accelerator. For example, 'nvidia-tesla-k80', or
        'tpu-v3'.
    Returns:
      self return to allow chained call with other resource specification.
    """
    accelerator_cnt = 1
    if self.container_spec.resources.accelerator.count > 1:
      # Reserve the number of already set.
      accelerator_cnt = self.container_spec.resources.accelerator.count

    accelerator_config = pipeline_spec_pb2.PipelineDeploymentConfig.PipelineContainerSpec.ResourceSpec.AcceleratorConfig(
        type=_sanitize_gpu_type(value),
        count=accelerator_cnt
    )
    self.container_spec.resources.accelerator.CopyFrom(accelerator_config)
    return self

  @resource_setter
  def set_gpu_limit(self, count: int) -> 'ContainerOp':
    """Sets the number of accelerator needed for this task."""
    if count < 1:
      raise ValueError('Accelerator count needs to be positive: Got: '
                       '{}'.format(count))
    self.container_spec.resources.accelerator.count = count
    return self
