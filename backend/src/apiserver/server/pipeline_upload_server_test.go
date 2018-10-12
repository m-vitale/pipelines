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

package server

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"io"

	"os"

	"github.com/googleprivate/ml/backend/src/apiserver/common"
	"github.com/googleprivate/ml/backend/src/apiserver/model"
	"github.com/googleprivate/ml/backend/src/apiserver/resource"
	"github.com/googleprivate/ml/backend/src/apiserver/storage"
	"github.com/googleprivate/ml/backend/src/common/util"
	"github.com/stretchr/testify/assert"
)

func TestUploadPipeline_YAML(t *testing.T) {
	clientManager := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	resourceManager := resource.NewResourceManager(clientManager)
	server := PipelineUploadServer{resourceManager: resourceManager}
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	part, _ := w.CreateFormFile("uploadfile", "hello-world.yaml")
	io.Copy(part, bytes.NewBufferString("apiVersion: argoproj.io/v1alpha1\nkind: Workflow"))
	w.Close()
	req, _ := http.NewRequest("POST", "/apis/v1alpha2/pipelines/upload", bytes.NewReader(b.Bytes()))
	req.Header.Set("Content-Type", w.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UploadPipeline)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	// Verify time format is RFC3339
	assert.Contains(t, rr.Body.String(), `"created_at":"1970-01-01T00:00:01Z"`)

	// Verify stored in object store
	template, err := clientManager.ObjectStore().GetFile(storage.PipelineFolder, resource.DefaultFakeUUID)
	assert.Nil(t, err)
	assert.NotNil(t, template)

	// Verify metadata in db
	pkgsExpect := []model.Pipeline{
		{
			UUID:           resource.DefaultFakeUUID,
			CreatedAtInSec: 1,
			Name:           "hello-world.yaml",
			Parameters:     "[]",
			Status:         model.PipelineReady}}
	pkg, str, err := clientManager.PipelineStore().ListPipelines(&common.PaginationContext{
		PageSize:        2,
		KeyFieldName:    model.GetPipelineTablePrimaryKeyColumn(),
		SortByFieldName: model.GetPipelineTablePrimaryKeyColumn(),
		IsDesc:          false,
	})
	assert.Nil(t, err)
	assert.Equal(t, str, "")
	assert.Equal(t, pkgsExpect, pkg)
}

func TestUploadPipeline_Tarball(t *testing.T) {
	clientManager := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	resourceManager := resource.NewResourceManager(clientManager)
	server := PipelineUploadServer{resourceManager: resourceManager}
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	part, _ := w.CreateFormFile("uploadfile", "arguments.tar.gz")
	fileReader, _ := os.Open("test/arguments_tarball/arguments.tar.gz")
	io.Copy(part, fileReader)
	w.Close()
	req, _ := http.NewRequest("POST", "/apis/v1alpha2/pipelines/upload", bytes.NewReader(b.Bytes()))
	req.Header.Set("Content-Type", w.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UploadPipeline)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	// Verify time format is RFC3339
	assert.Contains(t, rr.Body.String(), `"created_at":"1970-01-01T00:00:01Z"`)

	// Verify stored in object store
	template, err := clientManager.ObjectStore().GetFile(storage.PipelineFolder, resource.DefaultFakeUUID)
	assert.Nil(t, err)
	assert.NotNil(t, template)

	// Verify metadata in db
	pkgsExpect := []model.Pipeline{
		{
			UUID:           resource.DefaultFakeUUID,
			CreatedAtInSec: 1,
			Name:           "arguments.tar.gz",
			Parameters:     "[{\"name\":\"param1\",\"value\":\"hello\"},{\"name\":\"param2\"}]",
			Status:         model.PipelineReady}}
	pkg, str, err := clientManager.PipelineStore().ListPipelines(&common.PaginationContext{
		PageSize:        2,
		KeyFieldName:    model.GetPipelineTablePrimaryKeyColumn(),
		SortByFieldName: model.GetPipelineTablePrimaryKeyColumn(),
		IsDesc:          false,
	})
	assert.Nil(t, err)
	assert.Equal(t, str, "")
	assert.Equal(t, pkgsExpect, pkg)
}

func TestUploadPipeline_GetFormFileError(t *testing.T) {
	clientManager := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	resourceManager := resource.NewResourceManager(clientManager)

	server := PipelineUploadServer{resourceManager: resourceManager}
	var b bytes.Buffer
	b.WriteString("I am invalid file")
	w := multipart.NewWriter(&b)
	w.CreateFormFile("uploadfile", "hello-world.yaml")
	w.Close()
	req, _ := http.NewRequest("POST", "/apis/v1alpha2/pipeline/upload", bytes.NewReader(b.Bytes()))
	req.Header.Set("Content-Type", w.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UploadPipeline)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code)
	assert.Contains(t, string(rr.Body.Bytes()), "Failed to read pipeline")
}

func TestUploadPipeline_SpecifyFileName(t *testing.T) {
	clientManager := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	resourceManager := resource.NewResourceManager(clientManager)
	server := PipelineUploadServer{resourceManager: resourceManager}
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	part, _ := w.CreateFormFile("uploadfile", "hello-world.yaml")
	io.Copy(part, bytes.NewBufferString("apiVersion: argoproj.io/v1alpha1\nkind: Workflow"))
	w.Close()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/apis/v1alpha2/pipelines/upload?name=%s", url.PathEscape("foo bar")), bytes.NewReader(b.Bytes()))
	req.Header.Set("Content-Type", w.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UploadPipeline)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	// Verify stored in object store
	template, err := clientManager.ObjectStore().GetFile(storage.PipelineFolder, resource.DefaultFakeUUID)
	assert.Nil(t, err)
	assert.NotNil(t, template)

	// Verify metadata in db
	pkgsExpect := []model.Pipeline{
		{
			UUID:           resource.DefaultFakeUUID,
			CreatedAtInSec: 1,
			Name:           "foo bar",
			Parameters:     "[]",
			Status:         model.PipelineReady}}
	pkg, str, err := clientManager.PipelineStore().ListPipelines(&common.PaginationContext{
		PageSize:        2,
		KeyFieldName:    model.GetPipelineTablePrimaryKeyColumn(),
		SortByFieldName: model.GetPipelineTablePrimaryKeyColumn(),
		IsDesc:          false,
	})
	assert.Nil(t, err)
	assert.Equal(t, str, "")
	assert.Equal(t, pkgsExpect, pkg)
}

func TestUploadPipeline_FileNameTooLong(t *testing.T) {
	clientManager := resource.NewFakeClientManagerOrFatal(util.NewFakeTimeForEpoch())
	resourceManager := resource.NewResourceManager(clientManager)
	server := PipelineUploadServer{resourceManager: resourceManager}
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	part, _ := w.CreateFormFile("uploadfile", "hello-world.yaml")
	io.Copy(part, bytes.NewBufferString("apiVersion: argoproj.io/v1alpha1\nkind: Workflow"))
	w.Close()
	encodedName := url.PathEscape(
		"this is a loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog name")
	req, _ := http.NewRequest("POST", fmt.Sprintf("/apis/v1alpha2/pipelines/upload?name=%s", encodedName), bytes.NewReader(b.Bytes()))
	req.Header.Set("Content-Type", w.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UploadPipeline)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code)
	assert.Contains(t, string(rr.Body.Bytes()), "Pipeline name too long")
}
