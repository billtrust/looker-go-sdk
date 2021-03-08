// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewSessionConfigParams creates a new SessionConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSessionConfigParams() *SessionConfigParams {
	return &SessionConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSessionConfigParamsWithTimeout creates a new SessionConfigParams object
// with the ability to set a timeout on a request.
func NewSessionConfigParamsWithTimeout(timeout time.Duration) *SessionConfigParams {
	return &SessionConfigParams{
		timeout: timeout,
	}
}

// NewSessionConfigParamsWithContext creates a new SessionConfigParams object
// with the ability to set a context for a request.
func NewSessionConfigParamsWithContext(ctx context.Context) *SessionConfigParams {
	return &SessionConfigParams{
		Context: ctx,
	}
}

// NewSessionConfigParamsWithHTTPClient creates a new SessionConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewSessionConfigParamsWithHTTPClient(client *http.Client) *SessionConfigParams {
	return &SessionConfigParams{
		HTTPClient: client,
	}
}

/* SessionConfigParams contains all the parameters to send to the API endpoint
   for the session config operation.

   Typically these are written to a http.Request.
*/
type SessionConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the session config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionConfigParams) WithDefaults() *SessionConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the session config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SessionConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the session config params
func (o *SessionConfigParams) WithTimeout(timeout time.Duration) *SessionConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session config params
func (o *SessionConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session config params
func (o *SessionConfigParams) WithContext(ctx context.Context) *SessionConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session config params
func (o *SessionConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session config params
func (o *SessionConfigParams) WithHTTPClient(client *http.Client) *SessionConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session config params
func (o *SessionConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SessionConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
