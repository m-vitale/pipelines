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

import json
import os
import shutil
import subprocess
import tempfile
import unittest
import kfp


class TestV2Compiler(unittest.TestCase):
  """Unit tests for KFP DSL v2 compiler."""

  def _test_compilation_to_json(self, file_base_name):
    test_data_dir = os.path.dirname(__file__)
    py_file = os.path.join(test_data_dir, file_base_name + '.py')
    tmpdir = tempfile.mkdtemp()
    try:
      target_json = os.path.join(tmpdir, file_base_name + '-pipeline.json')
      subprocess.check_call(
          ['dsl-compile-v2', '--py', py_file, '--output', target_json])
      with open(os.path.join(test_data_dir, file_base_name + '.json'),
                'r') as f:
        golden = json.load(f)
        # Correct the sdkVersion
        golden['sdkVersion'] = kfp.__version__

      with open(os.path.join(test_data_dir, target_json), 'r') as f:
        compiled = json.load(f)

      self.maxDiff = None
      self.assertEqual(sorted(golden.items()), sorted(compiled.items()))
    finally:
      shutil.rmtree(tmpdir)

  def test_two_step_pipeline_with_importer(self):
    """Test two steps linear pipeline with importer node."""
    self._test_compilation_to_json('two_step_pipeline_with_importer')
