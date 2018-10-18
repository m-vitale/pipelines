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

// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/googleprivate/ml/backend/api/go_http_client/pipeline_model"
)

// CreatePipelineReader is a Reader for the CreatePipeline structure.
type CreatePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreatePipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePipelineOK creates a CreatePipelineOK with default headers values
func NewCreatePipelineOK() *CreatePipelineOK {
	return &CreatePipelineOK{}
}

/*CreatePipelineOK handles this case with default header values.

A successful response.
*/
type CreatePipelineOK struct {
	Payload *pipeline_model.APIPipeline
}

func (o *CreatePipelineOK) Error() string {
	return fmt.Sprintf("[POST /apis/v1alpha2/pipelines][%d] createPipelineOK  %+v", 200, o.Payload)
}

func (o *CreatePipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.APIPipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePipelineDefault creates a CreatePipelineDefault with default headers values
func NewCreatePipelineDefault(code int) *CreatePipelineDefault {
	return &CreatePipelineDefault{
		_statusCode: code,
	}
}

/*CreatePipelineDefault handles this case with default header values.

CreatePipelineDefault create pipeline default
*/
type CreatePipelineDefault struct {
	_statusCode int

	Payload *pipeline_model.APIStatus
}

// Code gets the status code for the create pipeline default response
func (o *CreatePipelineDefault) Code() int {
	return o._statusCode
}

func (o *CreatePipelineDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v1alpha2/pipelines][%d] CreatePipeline default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
