// Code generated by go-swagger; DO NOT EDIT.

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	fmt "fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	job_model "github.com/googleprivate/ml/backend/api/job_model"
)

// ListRunsReader is a Reader for the ListRuns structure.
type ListRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRunsOK creates a ListRunsOK with default headers values
func NewListRunsOK() *ListRunsOK {
	return &ListRunsOK{}
}

/*ListRunsOK handles this case with default header values.

ListRunsOK list runs o k
*/
type ListRunsOK struct {
	Payload *job_model.APIListRunsResponse
}

func (o *ListRunsOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1alpha2/jobs/{job_id}/runs][%d] listRunsOK  %+v", 200, o.Payload)
}

func (o *ListRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIListRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRunsDefault creates a ListRunsDefault with default headers values
func NewListRunsDefault(code int) *ListRunsDefault {
	return &ListRunsDefault{
		_statusCode: code,
	}
}

/*ListRunsDefault handles this case with default header values.

ListRunsDefault list runs default
*/
type ListRunsDefault struct {
	_statusCode int

	Payload *job_model.APIStatus
}

// Code gets the status code for the list runs default response
func (o *ListRunsDefault) Code() int {
	return o._statusCode
}

func (o *ListRunsDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1alpha2/jobs/{job_id}/runs][%d] ListRuns default  %+v", o._statusCode, o.Payload)
}

func (o *ListRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
