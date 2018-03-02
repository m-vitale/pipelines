#!/bin/bash -e
# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#  http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


export PROJECT_ID=$(gcloud config config-helper --format "value(configuration.properties.core.project)")
mkdir -p ./build
rsync -arvp "../xgboost"/ ./build/

docker pull ubuntu:16.04
docker build -t ml-pipeline-dataproc-xgboost .
rm -rf ./build

docker tag ml-pipeline-dataproc-xgboost gcr.io/${PROJECT_ID}/ml-pipeline-dataproc-xgboost
gcloud docker -- push gcr.io/${PROJECT_ID}/ml-pipeline-dataproc-xgboost
