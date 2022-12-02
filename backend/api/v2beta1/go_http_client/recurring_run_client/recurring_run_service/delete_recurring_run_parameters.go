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

// NewDeleteRecurringRunParams creates a new DeleteRecurringRunParams object
// with the default values initialized.
func NewDeleteRecurringRunParams() *DeleteRecurringRunParams {
	var ()
	return &DeleteRecurringRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecurringRunParamsWithTimeout creates a new DeleteRecurringRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRecurringRunParamsWithTimeout(timeout time.Duration) *DeleteRecurringRunParams {
	var ()
	return &DeleteRecurringRunParams{

		timeout: timeout,
	}
}

// NewDeleteRecurringRunParamsWithContext creates a new DeleteRecurringRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRecurringRunParamsWithContext(ctx context.Context) *DeleteRecurringRunParams {
	var ()
	return &DeleteRecurringRunParams{

		Context: ctx,
	}
}

// NewDeleteRecurringRunParamsWithHTTPClient creates a new DeleteRecurringRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRecurringRunParamsWithHTTPClient(client *http.Client) *DeleteRecurringRunParams {
	var ()
	return &DeleteRecurringRunParams{
		HTTPClient: client,
	}
}

/*DeleteRecurringRunParams contains all the parameters to send to the API endpoint
for the delete recurring run operation typically these are written to a http.Request
*/
type DeleteRecurringRunParams struct {

	/*RecRunID
	  The ID of the recurring run to be deleted.

	*/
	RecRunID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete recurring run params
func (o *DeleteRecurringRunParams) WithTimeout(timeout time.Duration) *DeleteRecurringRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recurring run params
func (o *DeleteRecurringRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recurring run params
func (o *DeleteRecurringRunParams) WithContext(ctx context.Context) *DeleteRecurringRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recurring run params
func (o *DeleteRecurringRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recurring run params
func (o *DeleteRecurringRunParams) WithHTTPClient(client *http.Client) *DeleteRecurringRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recurring run params
func (o *DeleteRecurringRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecRunID adds the recRunID to the delete recurring run params
func (o *DeleteRecurringRunParams) WithRecRunID(recRunID string) *DeleteRecurringRunParams {
	o.SetRecRunID(recRunID)
	return o
}

// SetRecRunID adds the recRunId to the delete recurring run params
func (o *DeleteRecurringRunParams) SetRecRunID(recRunID string) {
	o.RecRunID = recRunID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecurringRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
