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

import argparse
import importlib
from pathlib import Path
from typing import Text
import shlex

from nbformat import NotebookNode
from nbformat.v4 import new_notebook, new_code_cell
import tornado.ioloop
import tornado.web

exporter = importlib.import_module("exporter")

parser = argparse.ArgumentParser(description="Server Arguments")
parser.add_argument(
    "--timeout",
    type=int,
    default=100,
    help="Amount of time in seconds that a visualization can run for before " +
         "being stopped."
)


class VisualizationHandler(tornado.web.RequestHandler):

    def initialize(self, timeout: int):
        """Initializes custom RequestHandler that generates visualizations.

        The initialize function is used rather than the __init__ function due to
        tornado specs for RequestHandler.

        Args:
            timeout (int): Amount of time in seconds that a visualization can
            run for before being stopped.
        """
        # All necessary arguments required to generate visualizations.
        self.requestParser = argparse.ArgumentParser(
            description="Visualization Generator"
        )
        # Type of visualization to be generated.
        self.requestParser.add_argument(
            "--type",
            type=str,
            help="Type of visualization to be generated."
        )
        # Path of data to be used to generate visualization.
        self.requestParser.add_argument(
            "--input_path",
            type=str,
            help="Path of data to be used for generating visualization."
        )
        # Additional arguments to be used when generating a visualization
        # (provided as a string representation of JSON).
        self.requestParser.add_argument(
            "--arguments",
            type=str,
            default="{}",
            help="JSON string of arguments to be provided to visualizations."
        )
        self.exporter = exporter.Exporter(timeout)

    def get_arguments_from_body(self) -> argparse.Namespace:
        """Converts arguments from post request to argparser.Namespace format.

        This is done because arguments, by default are provided in the
        x-www-form-urlencoded format. This format is difficult to parse compared
        to argparser.Namespace, which is a dict.

        Returns:
            Arguments provided from post request as arparser.Namespace object.
        """
        split_arguments = shlex.split(self.get_body_argument("arguments"))
        return self.requestParser.parse_args(split_arguments)

    def validate_request_arguments(self, arguments: argparse.Namespace):
        """Validates arguments from post request and sends error if invalid.

        Args:
            arguments: x-www-form-urlencoded formatted arguments.
        """
        if arguments.type is None:
            return self.send_error(400, reason="No type specified.")
        if arguments.input_path is None:
            return self.send_error(400, reason="No input_path specified.")

    def generate_notebook_from_arguments(
        self,
        arguments: argparse.Namespace,
        input_path: Text,
        visualization_type: Text
    ) -> NotebookNode:
        """Generates a NotebookNode from provided arguments.

        Args:
            arguments: x-www-form-urlencoded formatted arguments.
            input_path: Path or path pattern to be used as data reference for
            visualization.
            visualization_type: Name of visualization to be generated.

        Returns:
                NotebookNode that contains all parameters from a post request.
        """
        nb = new_notebook()
        nb.cells.append(self.exporter.create_cell_from_args(arguments))
        nb.cells.append(new_code_cell('input_path = "{}"'.format(input_path)))
        visualization_file = str(Path.cwd() / "{}.py".format(visualization_type))
        nb.cells.append(self.exporter.create_cell_from_file(visualization_file))
        return nb

    def get(self):
        """Health check.
        """
        self.write("alive")

    def post(self):
        """Generates visualization based on provided arguments.
        """
        # Parse arguments from request.
        request_arguments = self.get_arguments_from_body()
        # Validate arguments from request.
        self.validate_request_arguments(request_arguments)
        # Create notebook with arguments from request.
        nb = self.generate_notebook_from_arguments(
            request_arguments.arguments,
            request_arguments.input_path,
            request_arguments.type
        )
        # Generate visualization (output for notebook).
        html = self.exporter.generate_html_from_notebook(nb)
        self.write(html)


if __name__ == "__main__":
    args = parser.parse_args()
    application = tornado.web.Application([
        (r"/", VisualizationHandler, dict(timeout=args.timeout)),
    ])
    application.listen(8888)
    tornado.ioloop.IOLoop.current().start()
