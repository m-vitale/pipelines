# Copyright 2019 Google LLC
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

import unittest
import snapshottest
from nbformat.v4 import new_code_cell
from nbformat.v4 import new_notebook
import exporter


class TestExporterMethods(snapshottest.TestCase):

    def test_create_cell_from_args_with_no_args(self):
        self.maxDiff = None
        args = "{}"
        cell = exporter.create_cell_from_args(args)
        self.assertEqual(cell.source, "")

    def test_create_cell_from_args_with_one_arg(self):
        self.maxDiff = None
        args = "{\"input_path\": \"gs://ml-pipeline/data.csv\"}"
        cell = exporter.create_cell_from_args(args)
        self.assertMatchSnapshot(cell.source)

    def test_create_cell_from_args_with_multiple_args(self):
        self.maxDiff = None
        args = "{\"input_path\": \"gs://ml-pipeline/data.csv\", " \
               "\"target_lambda\": \"lambda x: (x['target'] > x['fare'] * " \
               "0.2)\"} "
        cell = exporter.create_cell_from_args(args)
        self.assertMatchSnapshot(cell.source)

    def test_create_cell_from_file(self):
        self.maxDiff = None
        cell = exporter.create_cell_from_file("tfdv.py")
        self.assertMatchSnapshot(cell.source)

    def test_generate_html_from_notebook(self):
        self.maxDiff = None
        nb = new_notebook()
        args = "{\"x\": 2}"
        nb.cells.append(exporter.create_cell_from_args(args))
        nb.cells.append(new_code_cell("print(x)"))
        html = exporter.generate_html_from_notebook(nb, template_type='basic')
        self.assertMatchSnapshot(html)


if __name__ == "__main__":
    unittest.main()
    exporter.shutdown_kernel()
