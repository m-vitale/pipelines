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

import os
import shutil
import tempfile
import unittest
from kfp import components
from kfp import dsl
from kfp.v2 import compiler


class CompilerTest(unittest.TestCase):

  def test_compile_simple_pipeline(self):

    tmpdir = tempfile.mkdtemp()
    try:
      producer_op = components.load_component_from_text("""
      name: producer
      inputs:
      - {name: input_param, type: String}
      outputs:
      - {name: output_model, type: Model}
      - {name: output_value, type: Integer}
      implementation:
        container:
          image: gcr.io/my-project/my-image:tag
          args:
          - {inputValue: input_param}
          - {outputPath: output_model}
          - {outputPath: output_value}
      """)

      consumer_op = components.load_component_from_text("""
      name: consumer
      inputs:
      - {name: input_model, type: Model}
      - {name: input_value, type: Integer}
      implementation:
        container:
          image: gcr.io/my-project/my-image:tag
          args:
          - {inputPath: input_model}
          - {inputValue: input_value}
      """)

      @dsl.pipeline(name='two-step-pipeline')
      def simple_pipeline(pipeline_input='Hello KFP!',):
        producer = producer_op(input_param=pipeline_input)
        consumer = consumer_op(
            input_model=producer.outputs['output_model'],
            input_value=producer.outputs['output_value'])

      target_json_file = os.path.join(tmpdir, 'result.json')
      compiler.Compiler().compile(simple_pipeline, target_json_file)

      self.assertTrue(os.path.exists(target_json_file))
    finally:
      shutil.rmtree(tmpdir)


if __name__ == '__main__':
  unittest.main()
