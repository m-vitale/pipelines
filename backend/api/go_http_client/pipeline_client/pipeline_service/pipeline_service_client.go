// Copyright 2019 Google LLC
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

// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pipeline service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pipeline service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreatePipeline create pipeline API
*/
func (a *Client) CreatePipeline(params *CreatePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePipeline",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/pipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePipelineOK), nil

}

/*
CreatePipelineVersion create pipeline version API
*/
func (a *Client) CreatePipelineVersion(params *CreatePipelineVersionParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePipelineVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePipelineVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePipelineVersion",
		Method:             "POST",
		PathPattern:        "/apis/v1beta1/pipeline_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreatePipelineVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePipelineVersionOK), nil

}

/*
DeletePipeline delete pipeline API
*/
func (a *Client) DeletePipeline(params *DeletePipelineParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePipeline",
		Method:             "DELETE",
		PathPattern:        "/apis/v1beta1/pipelines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePipelineOK), nil

}

/*
DeletePipelineVersion delete pipeline version API
*/
func (a *Client) DeletePipelineVersion(params *DeletePipelineVersionParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePipelineVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePipelineVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePipelineVersion",
		Method:             "DELETE",
		PathPattern:        "/apis/v1beta1/pipeline_versions/{version_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePipelineVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePipelineVersionOK), nil

}

/*
GetPipeline get pipeline API
*/
func (a *Client) GetPipeline(params *GetPipelineParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPipeline",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/pipelines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPipelineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPipelineOK), nil

}

/*
GetPipelineVersion get pipeline version API
*/
func (a *Client) GetPipelineVersion(params *GetPipelineVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelineVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPipelineVersion",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/pipeline_versions/{version_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPipelineVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPipelineVersionOK), nil

}

/*
GetPipelineVersionTemplate get pipeline version template API
*/
func (a *Client) GetPipelineVersionTemplate(params *GetPipelineVersionTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*GetPipelineVersionTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineVersionTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPipelineVersionTemplate",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/pipeline_versions/{version_id}/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPipelineVersionTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPipelineVersionTemplateOK), nil

}

/*
GetTemplate get template API
*/
func (a *Client) GetTemplate(params *GetTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*GetTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTemplate",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/pipelines/{id}/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTemplateOK), nil

}

/*
ListPipelineVersions list pipeline versions API
*/
func (a *Client) ListPipelineVersions(params *ListPipelineVersionsParams, authInfo runtime.ClientAuthInfoWriter) (*ListPipelineVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPipelineVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPipelineVersions",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/pipeline_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPipelineVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPipelineVersionsOK), nil

}

/*
ListPipelines list pipelines API
*/
func (a *Client) ListPipelines(params *ListPipelinesParams, authInfo runtime.ClientAuthInfoWriter) (*ListPipelinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPipelinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPipelines",
		Method:             "GET",
		PathPattern:        "/apis/v1beta1/pipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPipelinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPipelinesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
