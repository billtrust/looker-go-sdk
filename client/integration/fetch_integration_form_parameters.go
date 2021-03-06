// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFetchIntegrationFormParams creates a new FetchIntegrationFormParams object
// with the default values initialized.
func NewFetchIntegrationFormParams() *FetchIntegrationFormParams {
	var ()
	return &FetchIntegrationFormParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFetchIntegrationFormParamsWithTimeout creates a new FetchIntegrationFormParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFetchIntegrationFormParamsWithTimeout(timeout time.Duration) *FetchIntegrationFormParams {
	var ()
	return &FetchIntegrationFormParams{

		timeout: timeout,
	}
}

// NewFetchIntegrationFormParamsWithContext creates a new FetchIntegrationFormParams object
// with the default values initialized, and the ability to set a context for a request
func NewFetchIntegrationFormParamsWithContext(ctx context.Context) *FetchIntegrationFormParams {
	var ()
	return &FetchIntegrationFormParams{

		Context: ctx,
	}
}

// NewFetchIntegrationFormParamsWithHTTPClient creates a new FetchIntegrationFormParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFetchIntegrationFormParamsWithHTTPClient(client *http.Client) *FetchIntegrationFormParams {
	var ()
	return &FetchIntegrationFormParams{
		HTTPClient: client,
	}
}

/*FetchIntegrationFormParams contains all the parameters to send to the API endpoint
for the fetch integration form operation typically these are written to a http.Request
*/
type FetchIntegrationFormParams struct {

	/*IntegrationID
	  Id of Integration

	*/
	IntegrationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the fetch integration form params
func (o *FetchIntegrationFormParams) WithTimeout(timeout time.Duration) *FetchIntegrationFormParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fetch integration form params
func (o *FetchIntegrationFormParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fetch integration form params
func (o *FetchIntegrationFormParams) WithContext(ctx context.Context) *FetchIntegrationFormParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fetch integration form params
func (o *FetchIntegrationFormParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fetch integration form params
func (o *FetchIntegrationFormParams) WithHTTPClient(client *http.Client) *FetchIntegrationFormParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fetch integration form params
func (o *FetchIntegrationFormParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the fetch integration form params
func (o *FetchIntegrationFormParams) WithIntegrationID(integrationID int64) *FetchIntegrationFormParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the fetch integration form params
func (o *FetchIntegrationFormParams) SetIntegrationID(integrationID int64) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *FetchIntegrationFormParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param integration_id
	if err := r.SetPathParam("integration_id", swag.FormatInt64(o.IntegrationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
