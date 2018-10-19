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

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIRun api run
// swagger:model apiRun
type APIRun struct {

	// Output. The time that the run created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional input field. Describing the purpose of the run
	Description string `json:"description,omitempty"`

	// In case any error happens retrieving a run field, only run ID
	// and the error message is returned. Client has the flexibility of choosing
	// how to handle error. This is especially useful during listing call.
	Error string `json:"error,omitempty"`

	// Output. Unique run ID. Generated by API server.
	ID string `json:"id,omitempty"`

	// TODO(yangpa): Following will be deprecated in v1beta1
	JobID string `json:"job_id,omitempty"`

	// Output. The metrics of the run. The metrics are reported by ReportMetrics
	// API.
	Metrics []*APIRunMetric `json:"metrics"`

	// Required input field. Name provided by user,
	// or auto generated if run is created by scheduled job. Not unique.
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// Required input field.
	// Describing what the pipeline manifest and parameters to use for the run.
	PipelineSpec *APIPipelineSpec `json:"pipeline_spec,omitempty"`

	// Optional input field. Specify which resource this run belongs to.
	ResourceReferences []*APIResourceReference `json:"resource_references"`

	// Output. When this run is scheduled to run. This could be different from
	// created_at. For example, if a run is from a backfilling job that was
	// supposed to run 2 month ago, the scheduled_at is 2 month ago,
	// v.s. created_at is the current time.
	// Format: date-time
	ScheduledAt strfmt.DateTime `json:"scheduled_at,omitempty"`

	// Output. The status of the run.
	// One of [Pending, Running, Succeeded, Skipped, Failed, Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this api run
func (m *APIRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipelineSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRun) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIRun) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {
		if swag.IsZero(m.Metrics[i]) { // not required
			continue
		}

		if m.Metrics[i] != nil {
			if err := m.Metrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRun) validatePipelineSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.PipelineSpec) { // not required
		return nil
	}

	if m.PipelineSpec != nil {
		if err := m.PipelineSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline_spec")
			}
			return err
		}
	}

	return nil
}

func (m *APIRun) validateResourceReferences(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceReferences) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceReferences); i++ {
		if swag.IsZero(m.ResourceReferences[i]) { // not required
			continue
		}

		if m.ResourceReferences[i] != nil {
			if err := m.ResourceReferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resource_references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRun) validateScheduledAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledAt) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduled_at", "body", "date-time", m.ScheduledAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRun) UnmarshalBinary(b []byte) error {
	var res APIRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
