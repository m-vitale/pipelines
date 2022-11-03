// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/pipeline_model"
)

// GetPipelineByNameReader is a Reader for the GetPipelineByName structure.
type GetPipelineByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPipelineByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPipelineByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPipelineByNameOK creates a GetPipelineByNameOK with default headers values
func NewGetPipelineByNameOK() *GetPipelineByNameOK {
	return &GetPipelineByNameOK{}
}

/*GetPipelineByNameOK handles this case with default header values.

A successful response.
*/
type GetPipelineByNameOK struct {
	Payload *pipeline_model.APIPipeline
}

func (o *GetPipelineByNameOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/namespaces/{namespace}/pipelines/{name}][%d] getPipelineByNameOK  %+v", 200, o.Payload)
}

func (o *GetPipelineByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.APIPipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineByNameDefault creates a GetPipelineByNameDefault with default headers values
func NewGetPipelineByNameDefault(code int) *GetPipelineByNameDefault {
	return &GetPipelineByNameDefault{
		_statusCode: code,
	}
}

/*GetPipelineByNameDefault handles this case with default header values.

GetPipelineByNameDefault get pipeline by name default
*/
type GetPipelineByNameDefault struct {
	_statusCode int

	Payload *pipeline_model.APIStatus
}

// Code gets the status code for the get pipeline by name default response
func (o *GetPipelineByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetPipelineByNameDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1beta1/namespaces/{namespace}/pipelines/{name}][%d] GetPipelineByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetPipelineByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
