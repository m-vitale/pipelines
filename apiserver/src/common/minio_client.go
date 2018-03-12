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

package common

import (
	"io"

	"github.com/minio/minio-go"
)

// Create interface for minio client struct, making it more unit testable.
type MinioClientInterface interface {
	PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (n int64, err error)
	GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (io.Reader, error)
}

type MinioClient struct {
	Client *minio.Client
}

func (c *MinioClient) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (n int64, err error) {
	return c.Client.PutObject(bucketName, objectName, reader, objectSize, opts)
}

func (c *MinioClient) GetObject(bucketName, objectName string, opts minio.GetObjectOptions) (io.Reader, error) {
	return c.Client.GetObject(bucketName, objectName, opts)
}
