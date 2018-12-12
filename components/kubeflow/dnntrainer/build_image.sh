#!/bin/bash -e
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

while getopts ":hp:t:i:m:" opt; do
  case "${opt}" in
    h) echo "-p: project name"
        echo "-t: tag name"
        echo "-i: image name. If provided, project name and tag name are not necessary"
        echo "-m: mode. Optional. The value can be 'cpu' or 'gpu'. Defaults to 'cpu'."
        exit
      ;;
    p) PROJECT_ID=${OPTARG}
      ;;
    t) TAG_NAME=${OPTARG}
      ;;
    i) IMAGE_NAME=${OPTARG}
      ;;
    m) MODE=${OPTARG}
      if [ "${MODE}" != "cpu" ] && [ "${MODE}" != "gpu" ]; then
        echo "mode value can only be 'cpu' or 'gpu'."
        exit
      fi
      ;;
    \? ) echo "Usage: cmd [-p] project [-t] tag [-i] image [-m] mode"
      exit
      ;;
  esac
done

set -x
LOCAL_IMAGE_NAME=ml-pipeline-kubeflow-tf-trainer

if [ -z "${PROJECT_ID}" ]; then
  PROJECT_ID=$(gcloud config config-helper --format "value(configuration.properties.core.project)")
fi

if [ -z "${TAG_NAME}" ]; then
  TAG_NAME=$(date +v%Y%m%d)-$(git describe --tags --always --dirty)-$(git diff | shasum -a256 | cut -c -6)
fi

TF_VERSION=1.6.0
if [ "${MODE}" = "gpu" ]; then
  TF_VERSION="${TF_VERSION}-gpu"
  LOCAL_IMAGE_NAME="${LOCAL_IMAGE_NAME}-gpu"
fi

mkdir -p ./build
rsync -arvp ./src/ ./build/
cp ../../license.sh ./build
cp ../../third_party_licenses.csv ./build

docker build --build-arg TF_VERSION=${TF_VERSION} -t ${LOCAL_IMAGE_NAME} .
if [ -z "${IMAGE_NAME}" ]; then
  docker tag ${LOCAL_IMAGE_NAME} gcr.io/${PROJECT_ID}/${LOCAL_IMAGE_NAME}:${TAG_NAME}
  docker push gcr.io/${PROJECT_ID}/${LOCAL_IMAGE_NAME}:${TAG_NAME}
else
  docker tag ${LOCAL_IMAGE_NAME} "${IMAGE_NAME}"
  docker push "${IMAGE_NAME}"
fi

rm -rf ./build