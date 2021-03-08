// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewCreateDashboardLayoutParams creates a new CreateDashboardLayoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDashboardLayoutParams() *CreateDashboardLayoutParams {
	return &CreateDashboardLayoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDashboardLayoutParamsWithTimeout creates a new CreateDashboardLayoutParams object
// with the ability to set a timeout on a request.
func NewCreateDashboardLayoutParamsWithTimeout(timeout time.Duration) *CreateDashboardLayoutParams {
	return &CreateDashboardLayoutParams{
		timeout: timeout,
	}
}

// NewCreateDashboardLayoutParamsWithContext creates a new CreateDashboardLayoutParams object
// with the ability to set a context for a request.
func NewCreateDashboardLayoutParamsWithContext(ctx context.Context) *CreateDashboardLayoutParams {
	return &CreateDashboardLayoutParams{
		Context: ctx,
	}
}

// NewCreateDashboardLayoutParamsWithHTTPClient creates a new CreateDashboardLayoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDashboardLayoutParamsWithHTTPClient(client *http.Client) *CreateDashboardLayoutParams {
	return &CreateDashboardLayoutParams{
		HTTPClient: client,
	}
}

/* CreateDashboardLayoutParams contains all the parameters to send to the API endpoint
   for the create dashboard layout operation.

   Typically these are written to a http.Request.
*/
type CreateDashboardLayoutParams struct {

	/* Body.

	   DashboardLayout
	*/
	Body *models.DashboardLayout

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create dashboard layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDashboardLayoutParams) WithDefaults() *CreateDashboardLayoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create dashboard layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDashboardLayoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create dashboard layout params
func (o *CreateDashboardLayoutParams) WithTimeout(timeout time.Duration) *CreateDashboardLayoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create dashboard layout params
func (o *CreateDashboardLayoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create dashboard layout params
func (o *CreateDashboardLayoutParams) WithContext(ctx context.Context) *CreateDashboardLayoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create dashboard layout params
func (o *CreateDashboardLayoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create dashboard layout params
func (o *CreateDashboardLayoutParams) WithHTTPClient(client *http.Client) *CreateDashboardLayoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create dashboard layout params
func (o *CreateDashboardLayoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create dashboard layout params
func (o *CreateDashboardLayoutParams) WithBody(body *models.DashboardLayout) *CreateDashboardLayoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create dashboard layout params
func (o *CreateDashboardLayoutParams) SetBody(body *models.DashboardLayout) {
	o.Body = body
}

// WithFields adds the fields to the create dashboard layout params
func (o *CreateDashboardLayoutParams) WithFields(fields *string) *CreateDashboardLayoutParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create dashboard layout params
func (o *CreateDashboardLayoutParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDashboardLayoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
