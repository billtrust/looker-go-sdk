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

	"your-damain.com/swagger/looker-api-golang/models"
)

// NewUpdateConnectionParams creates a new UpdateConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConnectionParams() *UpdateConnectionParams {
	return &UpdateConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConnectionParamsWithTimeout creates a new UpdateConnectionParams object
// with the ability to set a timeout on a request.
func NewUpdateConnectionParamsWithTimeout(timeout time.Duration) *UpdateConnectionParams {
	return &UpdateConnectionParams{
		timeout: timeout,
	}
}

// NewUpdateConnectionParamsWithContext creates a new UpdateConnectionParams object
// with the ability to set a context for a request.
func NewUpdateConnectionParamsWithContext(ctx context.Context) *UpdateConnectionParams {
	return &UpdateConnectionParams{
		Context: ctx,
	}
}

// NewUpdateConnectionParamsWithHTTPClient creates a new UpdateConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConnectionParamsWithHTTPClient(client *http.Client) *UpdateConnectionParams {
	return &UpdateConnectionParams{
		HTTPClient: client,
	}
}

/* UpdateConnectionParams contains all the parameters to send to the API endpoint
   for the update connection operation.

   Typically these are written to a http.Request.
*/
type UpdateConnectionParams struct {

	/* Body.

	   Connection
	*/
	Body *models.DBConnection

	/* ConnectionName.

	   Name of connection
	*/
	ConnectionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConnectionParams) WithDefaults() *UpdateConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update connection params
func (o *UpdateConnectionParams) WithTimeout(timeout time.Duration) *UpdateConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update connection params
func (o *UpdateConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update connection params
func (o *UpdateConnectionParams) WithContext(ctx context.Context) *UpdateConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update connection params
func (o *UpdateConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update connection params
func (o *UpdateConnectionParams) WithHTTPClient(client *http.Client) *UpdateConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update connection params
func (o *UpdateConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update connection params
func (o *UpdateConnectionParams) WithBody(body *models.DBConnection) *UpdateConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update connection params
func (o *UpdateConnectionParams) SetBody(body *models.DBConnection) {
	o.Body = body
}

// WithConnectionName adds the connectionName to the update connection params
func (o *UpdateConnectionParams) WithConnectionName(connectionName string) *UpdateConnectionParams {
	o.SetConnectionName(connectionName)
	return o
}

// SetConnectionName adds the connectionName to the update connection params
func (o *UpdateConnectionParams) SetConnectionName(connectionName string) {
	o.ConnectionName = connectionName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param connection_name
	if err := r.SetPathParam("connection_name", o.ConnectionName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
