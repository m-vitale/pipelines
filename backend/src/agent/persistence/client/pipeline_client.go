// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"time"

	"github.com/googleprivate/ml/backend/api"
	"github.com/googleprivate/ml/backend/src/common/util"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PipelineClientInterface interface {
	ReportWorkflow(workflow *util.Workflow) error
	ReportScheduledWorkflow(swf *util.ScheduledWorkflow) error
}

type PipelineClient struct {
	namespace             string
	initializeTimeout     time.Duration
	timeout               time.Duration
	basePath              string
	mlPipelineServiceName string
	mlPipelineServicePort string
	reportServiceClient   api.ReportServiceClient
}

func NewPipelineClient(
	namespace string,
	initializeTimeout time.Duration,
	timeout time.Duration,
	basePath string,
	mlPipelineServiceName string,
	mlPipelineServicePort string,
	masterurl string,
	kubeconfig string) (*PipelineClient, error) {

	err := util.WaitForGrpcClientAvailable(namespace, initializeTimeout, basePath, masterurl,
		kubeconfig)
	if err != nil {
		return nil, errors.Wrapf(err,
			"Failed to initialize pipeline client. Error: %s", err.Error())
	}

	connection, err := util.GetRpcConnection(namespace,
		mlPipelineServiceName, mlPipelineServicePort, masterurl, kubeconfig)
	if err != nil {
		return nil, errors.Wrapf(err,
			"Failed to get RPC connection. Error: %s", err.Error())
	}

	return &PipelineClient{
		namespace:             namespace,
		initializeTimeout:     initializeTimeout,
		timeout:               timeout,
		basePath:              basePath,
		mlPipelineServiceName: mlPipelineServiceName,
		mlPipelineServicePort: mlPipelineServicePort,
		reportServiceClient:   api.NewReportServiceClient(connection),
	}, nil
}

func (p *PipelineClient) ReportWorkflow(workflow *util.Workflow) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err := p.reportServiceClient.ReportWorkflow(ctx, &api.ReportWorkflowRequest{
		Workflow: workflow.ToStringForStore(),
	})

	if err != nil {
		statusCode, _ := status.FromError(err)
		if statusCode.Code() == codes.InvalidArgument {
			// Do not retry if there is something wrong with the workflow
			return util.NewCustomError(err, util.CUSTOM_CODE_PERMANENT,
				"Error while reporting workflow resource (code: %v, message: %v): %v, %+v",
				statusCode.Code(),
				statusCode.Message(),
				err.Error(),
				workflow.Workflow)
		} else {
			// Retry otherwise
			return util.NewCustomError(err, util.CUSTOM_CODE_TRANSIENT,
				"Error while reporting workflow resource (code: %v, message: %v): %v, %+v",
				statusCode.Code(),
				statusCode.Message(),
				err.Error(),
				workflow.Workflow)
		}
	}
	return nil
}

func (p *PipelineClient) ReportScheduledWorkflow(swf *util.ScheduledWorkflow) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err := p.reportServiceClient.ReportScheduledWorkflow(ctx,
		&api.ReportScheduledWorkflowRequest{
			ScheduledWorkflow: swf.ToStringForStore(),
		})

	if err != nil {
		statusCode, _ := status.FromError(err)
		if statusCode.Code() == codes.InvalidArgument {
			// Do not retry if there is something wrong with the workflow
			return util.NewCustomError(err, util.CUSTOM_CODE_PERMANENT,
				"Error while reporting workflow resource (code: %v, message: %v): %v, %+v",
				statusCode.Code(),
				statusCode.Message(),
				err.Error(),
				swf.ScheduledWorkflow)
		} else {
			// Retry otherwise
			return util.NewCustomError(err, util.CUSTOM_CODE_TRANSIENT,
				"Error while reporting workflow resource (code: %v, message: %v): %v, %+v",
				statusCode.Code(),
				statusCode.Message(),
				err.Error(),
				swf.ScheduledWorkflow)
		}
	}
	return nil
}
