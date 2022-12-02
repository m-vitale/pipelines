// Code generated by go-swagger; DO NOT EDIT.

package recurring_run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDisableRecurringRunParams creates a new DisableRecurringRunParams object
// with the default values initialized.
func NewDisableRecurringRunParams() *DisableRecurringRunParams {
	var ()
	return &DisableRecurringRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableRecurringRunParamsWithTimeout creates a new DisableRecurringRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableRecurringRunParamsWithTimeout(timeout time.Duration) *DisableRecurringRunParams {
	var ()
	return &DisableRecurringRunParams{

		timeout: timeout,
	}
}

// NewDisableRecurringRunParamsWithContext creates a new DisableRecurringRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableRecurringRunParamsWithContext(ctx context.Context) *DisableRecurringRunParams {
	var ()
	return &DisableRecurringRunParams{

		Context: ctx,
	}
}

// NewDisableRecurringRunParamsWithHTTPClient creates a new DisableRecurringRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableRecurringRunParamsWithHTTPClient(client *http.Client) *DisableRecurringRunParams {
	var ()
	return &DisableRecurringRunParams{
		HTTPClient: client,
	}
}

/*DisableRecurringRunParams contains all the parameters to send to the API endpoint
for the disable recurring run operation typically these are written to a http.Request
*/
type DisableRecurringRunParams struct {

	/*RecRunID
	  The ID of the recurring runs to be disabled.

	*/
	RecRunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the disable recurring run params
func (o *DisableRecurringRunParams) WithTimeout(timeout time.Duration) *DisableRecurringRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable recurring run params
func (o *DisableRecurringRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable recurring run params
func (o *DisableRecurringRunParams) WithContext(ctx context.Context) *DisableRecurringRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable recurring run params
func (o *DisableRecurringRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable recurring run params
func (o *DisableRecurringRunParams) WithHTTPClient(client *http.Client) *DisableRecurringRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable recurring run params
func (o *DisableRecurringRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecRunID adds the recRunID to the disable recurring run params
func (o *DisableRecurringRunParams) WithRecRunID(recRunID string) *DisableRecurringRunParams {
	o.SetRecRunID(recRunID)
	return o
}

// SetRecRunID adds the recRunId to the disable recurring run params
func (o *DisableRecurringRunParams) SetRecRunID(recRunID string) {
	o.RecRunID = recRunID
}

// WriteToRequest writes these params to a swagger request
func (o *DisableRecurringRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param recurring_run_id
	if err := r.SetPathParam("recurring_run_id", o.RecRunID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
