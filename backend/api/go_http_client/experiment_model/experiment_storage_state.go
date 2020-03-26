// Copyright 2020 Google LLC
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

package experiment_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ExperimentStorageState experiment storage state
// swagger:model ExperimentStorageState
type ExperimentStorageState string

const (

	// ExperimentStorageStateRESERVED captures enum value "RESERVED"
	ExperimentStorageStateRESERVED ExperimentStorageState = "RESERVED"

	// ExperimentStorageStateSTORAGESTATEARCHIVED captures enum value "STORAGESTATE_ARCHIVED"
	ExperimentStorageStateSTORAGESTATEARCHIVED ExperimentStorageState = "STORAGESTATE_ARCHIVED"

	// ExperimentStorageStateSTORAGESTATEAVAILABLE captures enum value "STORAGESTATE_AVAILABLE"
	ExperimentStorageStateSTORAGESTATEAVAILABLE ExperimentStorageState = "STORAGESTATE_AVAILABLE"
)

// for schema
var experimentStorageStateEnum []interface{}

func init() {
	var res []ExperimentStorageState
	if err := json.Unmarshal([]byte(`["RESERVED","STORAGESTATE_ARCHIVED","STORAGESTATE_AVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		experimentStorageStateEnum = append(experimentStorageStateEnum, v)
	}
}

func (m ExperimentStorageState) validateExperimentStorageStateEnum(path, location string, value ExperimentStorageState) error {
	if err := validate.Enum(path, location, value, experimentStorageStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this experiment storage state
func (m ExperimentStorageState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExperimentStorageStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
