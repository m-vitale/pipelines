// Copyright 2019 Google LLC
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

package client

import (
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)


func CreatePodClient(namespace string) (v1.PodInterface, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client.CoreV1().Pods(namespace), nil
}

func CreatePodClientOrFatal(namespace string, initConnectionTimeout time.Duration) v1.PodInterface {
	var podInterface v1.PodInterface
	var err error
	var operation = func() error {
		podInterface, err = CreatePodClient(namespace)
		if err != nil {
			return err
		}
		return nil
	}
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = initConnectionTimeout
	if err := backoff.Retry(operation, b); err != nil {
		glog.Fatalf("Failed to create Pod client. Error: %v", err)
	}
	return podInterface
}
