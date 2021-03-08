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

// NewUpdateDashboardElementParams creates a new UpdateDashboardElementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDashboardElementParams() *UpdateDashboardElementParams {
	return &UpdateDashboardElementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDashboardElementParamsWithTimeout creates a new UpdateDashboardElementParams object
// with the ability to set a timeout on a request.
func NewUpdateDashboardElementParamsWithTimeout(timeout time.Duration) *UpdateDashboardElementParams {
	return &UpdateDashboardElementParams{
		timeout: timeout,
	}
}

// NewUpdateDashboardElementParamsWithContext creates a new UpdateDashboardElementParams object
// with the ability to set a context for a request.
func NewUpdateDashboardElementParamsWithContext(ctx context.Context) *UpdateDashboardElementParams {
	return &UpdateDashboardElementParams{
		Context: ctx,
	}
}

// NewUpdateDashboardElementParamsWithHTTPClient creates a new UpdateDashboardElementParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDashboardElementParamsWithHTTPClient(client *http.Client) *UpdateDashboardElementParams {
	return &UpdateDashboardElementParams{
		HTTPClient: client,
	}
}

/* UpdateDashboardElementParams contains all the parameters to send to the API endpoint
   for the update dashboard element operation.

   Typically these are written to a http.Request.
*/
type UpdateDashboardElementParams struct {

	/* Body.

	   DashboardElement
	*/
	Body *models.DashboardElement

	/* DashboardElementID.

	   Id of dashboard element
	*/
	DashboardElementID string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update dashboard element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardElementParams) WithDefaults() *UpdateDashboardElementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update dashboard element params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDashboardElementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update dashboard element params
func (o *UpdateDashboardElementParams) WithTimeout(timeout time.Duration) *UpdateDashboardElementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dashboard element params
func (o *UpdateDashboardElementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dashboard element params
func (o *UpdateDashboardElementParams) WithContext(ctx context.Context) *UpdateDashboardElementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dashboard element params
func (o *UpdateDashboardElementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dashboard element params
func (o *UpdateDashboardElementParams) WithHTTPClient(client *http.Client) *UpdateDashboardElementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dashboard element params
func (o *UpdateDashboardElementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dashboard element params
func (o *UpdateDashboardElementParams) WithBody(body *models.DashboardElement) *UpdateDashboardElementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dashboard element params
func (o *UpdateDashboardElementParams) SetBody(body *models.DashboardElement) {
	o.Body = body
}

// WithDashboardElementID adds the dashboardElementID to the update dashboard element params
func (o *UpdateDashboardElementParams) WithDashboardElementID(dashboardElementID string) *UpdateDashboardElementParams {
	o.SetDashboardElementID(dashboardElementID)
	return o
}

// SetDashboardElementID adds the dashboardElementId to the update dashboard element params
func (o *UpdateDashboardElementParams) SetDashboardElementID(dashboardElementID string) {
	o.DashboardElementID = dashboardElementID
}

// WithFields adds the fields to the update dashboard element params
func (o *UpdateDashboardElementParams) WithFields(fields *string) *UpdateDashboardElementParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the update dashboard element params
func (o *UpdateDashboardElementParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDashboardElementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dashboard_element_id
	if err := r.SetPathParam("dashboard_element_id", o.DashboardElementID); err != nil {
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
