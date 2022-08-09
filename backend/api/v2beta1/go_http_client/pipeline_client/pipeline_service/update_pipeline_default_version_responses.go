// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v2beta1/go_http_client/pipeline_model"
)

// UpdatePipelineDefaultVersionReader is a Reader for the UpdatePipelineDefaultVersion structure.
type UpdatePipelineDefaultVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePipelineDefaultVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePipelineDefaultVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdatePipelineDefaultVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePipelineDefaultVersionOK creates a UpdatePipelineDefaultVersionOK with default headers values
func NewUpdatePipelineDefaultVersionOK() *UpdatePipelineDefaultVersionOK {
	return &UpdatePipelineDefaultVersionOK{}
}

/*UpdatePipelineDefaultVersionOK handles this case with default header values.

A successful response.
*/
type UpdatePipelineDefaultVersionOK struct {
	Payload interface{}
}

func (o *UpdatePipelineDefaultVersionOK) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/pipelines/{pipeline_id}/default_version/{version_id}][%d] updatePipelineDefaultVersionOK  %+v", 200, o.Payload)
}

func (o *UpdatePipelineDefaultVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePipelineDefaultVersionDefault creates a UpdatePipelineDefaultVersionDefault with default headers values
func NewUpdatePipelineDefaultVersionDefault(code int) *UpdatePipelineDefaultVersionDefault {
	return &UpdatePipelineDefaultVersionDefault{
		_statusCode: code,
	}
}

/*UpdatePipelineDefaultVersionDefault handles this case with default header values.

UpdatePipelineDefaultVersionDefault update pipeline default version default
*/
type UpdatePipelineDefaultVersionDefault struct {
	_statusCode int

	Payload *pipeline_model.V2beta1Status
}

// Code gets the status code for the update pipeline default version default response
func (o *UpdatePipelineDefaultVersionDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePipelineDefaultVersionDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v2beta1/pipelines/{pipeline_id}/default_version/{version_id}][%d] UpdatePipelineDefaultVersion default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePipelineDefaultVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V2beta1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
