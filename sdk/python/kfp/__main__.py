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

from .cli.cli import main

# TODO(hongyes): add more commands:
# kfp compile (migrate from dsl-compile)
# kfp experiment (manage experiments)

# TODO: kfp pipeline delete
# Calling client._pipelines_api.delete_pipeline(id=pipeline_id) throws
# exception: module 'kfp_server_api.models' has no attribute 'ERRORUNKNOWN'

if __name__ == '__main__':
    main()
