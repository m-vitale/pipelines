// Code generated by go-swagger; DO NOT EDIT.

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	job_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/job_model"
)

// ListJobsReader is a Reader for the ListJobs structure.
type ListJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListJobsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListJobsOK creates a ListJobsOK with default headers values
func NewListJobsOK() *ListJobsOK {
	return &ListJobsOK{}
}

/*ListJobsOK handles this case with default header values.

A successful response.
*/
type ListJobsOK struct {
	Payload *job_model.V2beta1ListJobsResponse
}

func (o *ListJobsOK) Error() string {
	return fmt.Sprintf("[GET /apis/v2beta1/jobs][%d] listJobsOK  %+v", 200, o.Payload)
}

func (o *ListJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.V2beta1ListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobsDefault creates a ListJobsDefault with default headers values
func NewListJobsDefault(code int) *ListJobsDefault {
	return &ListJobsDefault{
		_statusCode: code,
	}
}

/*ListJobsDefault handles this case with default header values.

ListJobsDefault list jobs default
*/
type ListJobsDefault struct {
	_statusCode int

	Payload *job_model.V2beta1Status
}

// Code gets the status code for the list jobs default response
func (o *ListJobsDefault) Code() int {
	return o._statusCode
}

func (o *ListJobsDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v2beta1/jobs][%d] ListJobs default  %+v", o._statusCode, o.Payload)
}

func (o *ListJobsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.V2beta1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
