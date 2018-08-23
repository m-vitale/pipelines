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


import mlp


@mlp.pipeline(
  name='Immediate Value',
  description='A pipeline with parameter values hard coded'
)
def immediate_value_pipeline():
  # "url" is a pipeline parameter with value being hard coded.
  # It is useful in case for some component you want to hard code a parameter instead
  # of exposing it as a pipeline parameter.
  url=mlp.PipelineParam(name='url', value='gs://ml-pipeline/shakespeare1.txt')
  op1 = mlp.ContainerOp(
     name='download',
     image='google/cloud-sdk',
     command=['sh', '-c'],
     arguments=['gsutil cat %s | tee /tmp/results.txt' % url],
     file_outputs={'downloaded': '/tmp/results.txt'})
  op2 = mlp.ContainerOp(
     name='echo',
     image='library/bash',
     command=['sh', '-c'],
     arguments=['echo %s' % op1.output])
