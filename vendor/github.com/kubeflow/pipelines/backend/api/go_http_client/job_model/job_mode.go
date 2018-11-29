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

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobMode Required input.
//
//  - DISABLED: The job won't schedule any run if disabled.
// swagger:model JobMode
type JobMode string

const (

	// JobModeUNKNOWNMODE captures enum value "UNKNOWN_MODE"
	JobModeUNKNOWNMODE JobMode = "UNKNOWN_MODE"

	// JobModeENABLED captures enum value "ENABLED"
	JobModeENABLED JobMode = "ENABLED"

	// JobModeDISABLED captures enum value "DISABLED"
	JobModeDISABLED JobMode = "DISABLED"
)

// for schema
var jobModeEnum []interface{}

func init() {
	var res []JobMode
	if err := json.Unmarshal([]byte(`["UNKNOWN_MODE","ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobModeEnum = append(jobModeEnum, v)
	}
}

func (m JobMode) validateJobModeEnum(path, location string, value JobMode) error {
	if err := validate.Enum(path, location, value, jobModeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this job mode
func (m JobMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
