// Code generated by go-swagger; DO NOT EDIT.

package pipeline_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V2beta1ListPipelinesResponse v2beta1 list pipelines response
// swagger:model v2beta1ListPipelinesResponse
type V2beta1ListPipelinesResponse struct {

	// The token to list the next page of pipelines.
	NextPageToken string `json:"next_page_token,omitempty"`

	// pipelines
	Pipelines []*V2beta1Pipeline `json:"pipelines"`

	// The total number of pipelines for the given query.
	TotalSize int32 `json:"total_size,omitempty"`
}

// Validate validates this v2beta1 list pipelines response
func (m *V2beta1ListPipelinesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePipelines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V2beta1ListPipelinesResponse) validatePipelines(formats strfmt.Registry) error {

	if swag.IsZero(m.Pipelines) { // not required
		return nil
	}

	for i := 0; i < len(m.Pipelines); i++ {
		if swag.IsZero(m.Pipelines[i]) { // not required
			continue
		}

		if m.Pipelines[i] != nil {
			if err := m.Pipelines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pipelines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V2beta1ListPipelinesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2beta1ListPipelinesResponse) UnmarshalBinary(b []byte) error {
	var res V2beta1ListPipelinesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
