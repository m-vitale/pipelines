// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	experiment_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/experiment_model"
)

// ArchiveExperimentReader is a Reader for the ArchiveExperiment structure.
type ArchiveExperimentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveExperimentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewArchiveExperimentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewArchiveExperimentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveExperimentOK creates a ArchiveExperimentOK with default headers values
func NewArchiveExperimentOK() *ArchiveExperimentOK {
	return &ArchiveExperimentOK{}
}

/*ArchiveExperimentOK handles this case with default header values.

A successful response.
*/
type ArchiveExperimentOK struct {
	Payload interface{}
}

func (o *ArchiveExperimentOK) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/experiments/{id}:archive][%d] archiveExperimentOK  %+v", 200, o.Payload)
}

func (o *ArchiveExperimentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveExperimentDefault creates a ArchiveExperimentDefault with default headers values
func NewArchiveExperimentDefault(code int) *ArchiveExperimentDefault {
	return &ArchiveExperimentDefault{
		_statusCode: code,
	}
}

/*ArchiveExperimentDefault handles this case with default header values.

ArchiveExperimentDefault archive experiment default
*/
type ArchiveExperimentDefault struct {
	_statusCode int

	Payload *experiment_model.APIStatus
}

// Code gets the status code for the archive experiment default response
func (o *ArchiveExperimentDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveExperimentDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/experiments/{id}:archive][%d] ArchiveExperiment default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveExperimentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(experiment_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
