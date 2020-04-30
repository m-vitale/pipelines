#!/usr/bin/env bash

set -e

REMOTE_REPOSITORY="amazon/aws-sagemaker-kfp-components"
DRYRUN=true

while getopts ":d:" opt; do
	case ${opt} in
		d)
			if [[ "${OPTARG}" = "false" ]]; then
				DRYRUN=false
			else
				DRYRUN=true
			fi
			;;
	esac
done

function docker_tag_exists() {
    curl --silent -f -lSL https://index.docker.io/v1/repositories/$1/tags/$2 > /dev/null 2> /dev/null
}

# Check version does not already exist
VERSION_LICENSE_FILE="THIRD-PARTY-LICENSES.txt"
FULL_VERSION_TAG="$(cat ${VERSION_LICENSE_FILE} | head -n1 | grep -Po '(?<=version )\d.\d.\d')"

if [ -z "$FULL_VERSION_TAG" ]; then
  >&2 echo "Could not find version inside ${VERSION_LICENSE_FILE} file."
  exit 1
fi

echo "Deploying version ${FULL_VERSION_TAG}"

if docker_tag_exists "$REMOTE_REPOSITORY" "$FULL_VERSION_TAG"; then
  >&2 echo "Tag ${REMOTE_REPOSITORY}:${FULL_VERSION_TAG} already exists. Cannot overwrite an existing image."
  exit 1
fi

# Build the image
FULL_VERSION_IMAGE="${REMOTE_REPOSITORY}:${FULL_VERSION_TAG}"
docker build . -f Dockerfile -t "${FULL_VERSION_IMAGE}"

# Get the minor and major versions
[[ $FULL_VERSION_TAG =~ ^[0-9]+\.[0-9]+ ]] && 	MINOR_VERSION_IMAGE="${REMOTE_REPOSITORY}:${BASH_REMATCH[0]}"
[[ $FULL_VERSION_TAG =~ ^[0-9]+ ]] && 					MAJOR_VERSION_IMAGE="${REMOTE_REPOSITORY}:${BASH_REMATCH[0]}"

# Re-tag the image with major and minor versions
docker tag "${FULL_VERSION_IMAGE}" "${MINOR_VERSION_IMAGE}"
echo "Tagged image with ${MINOR_VERSION_IMAGE}"
docker tag "${FULL_VERSION_IMAGE}" "${MAJOR_VERSION_IMAGE}"
echo "Tagged image with ${MAJOR_VERSION_IMAGE}"

# Push to the remote repository
if [ "${DRYRUN}" == "false" ]; then
  docker push "${FULL_VERSION_IMAGE}"
  echo "Successfully pushed tag ${FULL_VERSION_IMAGE} to Docker Hub"

	docker push "${MINOR_VERSION_IMAGE}"
  echo "Successfully pushed tag ${MINOR_VERSION_IMAGE} to Docker Hub"

	docker push "${MAJOR_VERSION_IMAGE}"
  echo "Successfully pushed tag ${MAJOR_VERSION_IMAGE} to Docker Hub"
else
  echo "Dry run detected. Not pushing images."
fi