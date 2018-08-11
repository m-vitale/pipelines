// Code generated by go-swagger; DO NOT EDIT.

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new job service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateJob create job API
*/
func (a *Client) CreateJob(params *CreateJobParams) (*CreateJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateJob",
		Method:             "POST",
		PathPattern:        "/apis/v1alpha2/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateJobOK), nil

}

/*
DeleteJob delete job API
*/
func (a *Client) DeleteJob(params *DeleteJobParams) (*DeleteJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteJob",
		Method:             "DELETE",
		PathPattern:        "/apis/v1alpha2/jobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteJobOK), nil

}

/*
DisableJob disable job API
*/
func (a *Client) DisableJob(params *DisableJobParams) (*DisableJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DisableJob",
		Method:             "POST",
		PathPattern:        "/apis/v1alpha2/jobs/{id}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisableJobOK), nil

}

/*
EnableJob enable job API
*/
func (a *Client) EnableJob(params *EnableJobParams) (*EnableJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EnableJob",
		Method:             "POST",
		PathPattern:        "/apis/v1alpha2/jobs/{id}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnableJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EnableJobOK), nil

}

/*
GetJob get job API
*/
func (a *Client) GetJob(params *GetJobParams) (*GetJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetJob",
		Method:             "GET",
		PathPattern:        "/apis/v1alpha2/jobs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJobOK), nil

}

/*
ListJobs list jobs API
*/
func (a *Client) ListJobs(params *ListJobsParams) (*ListJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListJobsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListJobs",
		Method:             "GET",
		PathPattern:        "/apis/v1alpha2/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListJobsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListJobsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
