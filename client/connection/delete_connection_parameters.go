// Code generated by go-swagger; DO NOT EDIT.

package connection

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

// NewDeleteConnectionParams creates a new DeleteConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConnectionParams() *DeleteConnectionParams {
	return &DeleteConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConnectionParamsWithTimeout creates a new DeleteConnectionParams object
// with the ability to set a timeout on a request.
func NewDeleteConnectionParamsWithTimeout(timeout time.Duration) *DeleteConnectionParams {
	return &DeleteConnectionParams{
		timeout: timeout,
	}
}

// NewDeleteConnectionParamsWithContext creates a new DeleteConnectionParams object
// with the ability to set a context for a request.
func NewDeleteConnectionParamsWithContext(ctx context.Context) *DeleteConnectionParams {
	return &DeleteConnectionParams{
		Context: ctx,
	}
}

// NewDeleteConnectionParamsWithHTTPClient creates a new DeleteConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConnectionParamsWithHTTPClient(client *http.Client) *DeleteConnectionParams {
	return &DeleteConnectionParams{
		HTTPClient: client,
	}
}

/* DeleteConnectionParams contains all the parameters to send to the API endpoint
   for the delete connection operation.

   Typically these are written to a http.Request.
*/
type DeleteConnectionParams struct {

	/* ConnectionName.

	   Name of connection
	*/
	ConnectionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConnectionParams) WithDefaults() *DeleteConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete connection params
func (o *DeleteConnectionParams) WithTimeout(timeout time.Duration) *DeleteConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete connection params
func (o *DeleteConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete connection params
func (o *DeleteConnectionParams) WithContext(ctx context.Context) *DeleteConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete connection params
func (o *DeleteConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete connection params
func (o *DeleteConnectionParams) WithHTTPClient(client *http.Client) *DeleteConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete connection params
func (o *DeleteConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionName adds the connectionName to the delete connection params
func (o *DeleteConnectionParams) WithConnectionName(connectionName string) *DeleteConnectionParams {
	o.SetConnectionName(connectionName)
	return o
}

// SetConnectionName adds the connectionName to the delete connection params
func (o *DeleteConnectionParams) SetConnectionName(connectionName string) {
	o.ConnectionName = connectionName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_name
	if err := r.SetPathParam("connection_name", o.ConnectionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
