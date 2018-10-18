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

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	job_model "github.com/googleprivate/ml/backend/api/go_http_client/job_model"
)

// ListJobRunsReader is a Reader for the ListJobRuns structure.
type ListJobRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListJobRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListJobRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListJobRunsOK creates a ListJobRunsOK with default headers values
func NewListJobRunsOK() *ListJobRunsOK {
	return &ListJobRunsOK{}
}

/*ListJobRunsOK handles this case with default header values.

A successful response.
*/
type ListJobRunsOK struct {
	Payload *job_model.APIListJobRunsResponse
}

func (o *ListJobRunsOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1alpha2/jobs/{job_id}/runs][%d] listJobRunsOK  %+v", 200, o.Payload)
}

func (o *ListJobRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIListJobRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobRunsDefault creates a ListJobRunsDefault with default headers values
func NewListJobRunsDefault(code int) *ListJobRunsDefault {
	return &ListJobRunsDefault{
		_statusCode: code,
	}
}

/*ListJobRunsDefault handles this case with default header values.

ListJobRunsDefault list job runs default
*/
type ListJobRunsDefault struct {
	_statusCode int

	Payload *job_model.APIStatus
}

// Code gets the status code for the list job runs default response
func (o *ListJobRunsDefault) Code() int {
	return o._statusCode
}

func (o *ListJobRunsDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1alpha2/jobs/{job_id}/runs][%d] ListJobRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListJobRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
