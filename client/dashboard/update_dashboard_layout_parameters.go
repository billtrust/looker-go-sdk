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

// NewUpdateDashboardLayoutParams creates a new UpdateDashboardLayoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardLayoutParams() *UpdateDashboardLayoutParams {
	return &UpdateDashboardLayoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardLayoutParamsWithTimeout creates a new UpdateDashboardLayoutParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardLayoutParamsWithTimeout(timeout time.Duration) *UpdateDashboardLayoutParams {
	return &UpdateDashboardLayoutParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardLayoutParamsWithContext creates a new UpdateDashboardLayoutParams object
// with the ability to set a context for a request.
func NewUpdateDashboardLayoutParamsWithContext(ctx context.Context) *UpdateDashboardLayoutParams {
	return &UpdateDashboardLayoutParams{
		Context: ctx,
	}
}

// NewUpdateDashboardLayoutParamsWithHTTPClient creates a new UpdateDashboardLayoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardLayoutParamsWithHTTPClient(client *http.Client) *UpdateDashboardLayoutParams {
	return &UpdateDashboardLayoutParams{
		HTTPClient: client,
	}
}

/* UpdateDashboardLayoutParams contains all the parameters to send to the API endpoint
   for the update dashboard layout operation.

   Typically these are written to a http.Request.
*/
type UpdateDashboardLayoutParams struct {

	/* Body.

	   DashboardLayout
	*/
	Body *models.DashboardLayout

	/* DashboardLayoutID.

	   Id of dashboard layout
	*/
	DashboardLayoutID string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardLayoutParams) WithDefaults() *UpdateDashboardLayoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardLayoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) WithTimeout(timeout time.Duration) *UpdateDashboardLayoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) WithContext(ctx context.Context) *UpdateDashboardLayoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) WithHTTPClient(client *http.Client) *UpdateDashboardLayoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) WithBody(body *models.DashboardLayout) *UpdateDashboardLayoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) SetBody(body *models.DashboardLayout) {
	o.Body = body
}

// WithDashboardLayoutID adds the dashboardLayoutID to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) WithDashboardLayoutID(dashboardLayoutID string) *UpdateDashboardLayoutParams {
	o.SetDashboardLayoutID(dashboardLayoutID)
	return o
}

// SetDashboardLayoutID adds the dashboardLayoutId to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) SetDashboardLayoutID(dashboardLayoutID string) {
	o.DashboardLayoutID = dashboardLayoutID
}

// WithFields adds the fields to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) WithFields(fields *string) *UpdateDashboardLayoutParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update dashboard layout params
func (o *UpdateDashboardLayoutParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardLayoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dashboard_layout_id
	if err := r.SetPathParam("dashboard_layout_id", o.DashboardLayoutID); err != nil {
		return err
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
