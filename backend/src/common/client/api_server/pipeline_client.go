package api_server

import (
	"fmt"

	workflowapi "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/ghodss/yaml"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/googleprivate/ml/backend/api/pipeline_client"
	params "github.com/googleprivate/ml/backend/api/pipeline_client/pipeline_service"
	model "github.com/googleprivate/ml/backend/api/pipeline_model"
	"github.com/googleprivate/ml/backend/src/common/util"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type PipelineInterface interface {
	Get(params *params.GetPipelineParams) (*model.APIPipeline, error)
	Delete(params *params.DeletePipelineParams) error
	GetTemplate(params *params.GetTemplateParams) (*workflowapi.Workflow, error)
	List(params *params.ListPipelinesParams) ([]*model.APIPipeline, string, error)
	ListAll(params *params.ListPipelinesParams, maxResultSize int) (
		[]*model.APIPipeline, error)
}

type PipelineClient struct {
	apiClient *apiclient.Pipeline
}

func NewPipelineClient(clientConfig clientcmd.ClientConfig) (*PipelineClient, error) {
	// Creating k8 client
	k8Client, config, namespace, err := util.GetKubernetesClientFromClientConfig(clientConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "Error while creating K8 client")
	}

	// Create API client
	httpClient := k8Client.RESTClient().(*rest.RESTClient).Client
	masterIPAndPort := util.ExtractMasterIPAndPort(config)
	apiClient := apiclient.New(httptransport.NewWithClient(masterIPAndPort,
		fmt.Sprintf(apiServerBasePath, namespace), nil, httpClient), strfmt.Default)

	// Creating upload client
	return &PipelineClient{
		apiClient: apiClient,
	}, nil
}

func (c *PipelineClient) Get(parameters *params.GetPipelineParams) (*model.APIPipeline,
	error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service all
	parameters.Context = ctx
	response, err := c.apiClient.PipelineService.GetPipeline(parameters)
	if err != nil {
		if defaultError, ok := err.(*params.GetPipelineDefault); ok {
			err = fmt.Errorf(defaultError.Payload.Error)
		}

		return nil, util.NewUserError(err,
			fmt.Sprintf("Failed to get pipeline. Params: %v", parameters),
			fmt.Sprintf("Failed to get pipeline %v", parameters.ID))
	}

	return response.Payload, nil
}

func (c *PipelineClient) Delete(parameters *params.DeletePipelineParams) error {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service all
	parameters.Context = ctx
	_, err := c.apiClient.PipelineService.DeletePipeline(parameters)
	if err != nil {
		if defaultError, ok := err.(*params.DeletePipelineDefault); ok {
			err = fmt.Errorf(defaultError.Payload.Error)
		}

		return util.NewUserError(err,
			fmt.Sprintf("Failed to delete pipeline. Params: %+v", parameters),
			fmt.Sprintf("Failed to delete pipeline %v", parameters.ID))
	}

	return nil
}

func (c *PipelineClient) GetTemplate(parameters *params.GetTemplateParams) (
	*workflowapi.Workflow, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service all
	parameters.Context = ctx
	response, err := c.apiClient.PipelineService.GetTemplate(parameters)
	if err != nil {
		if defaultError, ok := err.(*params.GetTemplateDefault); ok {
			err = fmt.Errorf(defaultError.Payload.Error)
		}

		return nil, util.NewUserError(err,
			fmt.Sprintf("Failed to get template. Params: %+v", parameters),
			fmt.Sprintf("Failed to get template for pipeline %v", parameters.ID))
	}

	// Unmarshal response
	var workflow workflowapi.Workflow
	err = yaml.Unmarshal([]byte(response.Payload.Template), &workflow)
	if err != nil {
		return nil, util.NewUserError(err,
			fmt.Sprintf("Failed to unmarshal reponse. Params: %+v. Response: %s", parameters,
				response.Payload.Template),
			fmt.Sprintf("Failed to unmarshal reponse"))
	}

	return &workflow, nil
}

func (c *PipelineClient) List(parameters *params.ListPipelinesParams) (
	[]*model.APIPipeline, string, error) {
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), apiServerDefaultTimeout)
	defer cancel()

	// Make service all
	parameters.Context = ctx
	response, err := c.apiClient.PipelineService.ListPipelines(parameters)
	if err != nil {
		if defaultError, ok := err.(*params.ListPipelinesDefault); ok {
			err = fmt.Errorf(defaultError.Payload.Error)
		}

		return nil, "", util.NewUserError(err,
			fmt.Sprintf("Failed to list pipelines. Params: %+v", parameters),
			fmt.Sprintf("Failed to list pipelines"))
	}

	return response.Payload.Pipelines, response.Payload.NextPageToken, nil
}

func (c *PipelineClient) ListAll(parameters *params.ListPipelinesParams, maxResultSize int) (
	[]*model.APIPipeline, error) {
	return listAllForPipeline(c, parameters, maxResultSize)
}

func listAllForPipeline(client PipelineInterface, parameters *params.ListPipelinesParams,
	maxResultSize int) ([]*model.APIPipeline, error) {
	if maxResultSize < 0 {
		maxResultSize = 0
	}

	allResults := make([]*model.APIPipeline, 0)
	firstCall := true
	for (firstCall || (parameters.PageToken != nil && *parameters.PageToken != "")) &&
		(len(allResults) < maxResultSize) {
		results, pageToken, err := client.List(parameters)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
		parameters.PageToken = util.StringPointer(pageToken)
		firstCall = false
	}
	if len(allResults) > maxResultSize {
		allResults = allResults[0:maxResultSize]
	}

	return allResults, nil
}
