// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"path"

	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
)

var pipelineFolder string = common.GetStringConfigWithDefault("ObjectStoreConfig.PipelineFolder", "pipelines")

// CreatePipelinePath creates object store path to a pipeline spec.
func CreatePipelinePath(pipelineID string) string {
	return path.Join(pipelineFolder, pipelineID)
}
