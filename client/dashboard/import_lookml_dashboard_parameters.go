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
	"github.com/go-openapi/swag"

	"your-damain.com/swagger/looker-api-golang/models"
)

// NewImportLookmlDashboardParams creates a new ImportLookmlDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportLookmlDashboardParams() *ImportLookmlDashboardParams {
	return &ImportLookmlDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportLookmlDashboardParamsWithTimeout creates a new ImportLookmlDashboardParams object
// with the ability to set a timeout on a request.
func NewImportLookmlDashboardParamsWithTimeout(timeout time.Duration) *ImportLookmlDashboardParams {
	return &ImportLookmlDashboardParams{
		timeout: timeout,
	}
}

// NewImportLookmlDashboardParamsWithContext creates a new ImportLookmlDashboardParams object
// with the ability to set a context for a request.
func NewImportLookmlDashboardParamsWithContext(ctx context.Context) *ImportLookmlDashboardParams {
	return &ImportLookmlDashboardParams{
		Context: ctx,
	}
}

// NewImportLookmlDashboardParamsWithHTTPClient creates a new ImportLookmlDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportLookmlDashboardParamsWithHTTPClient(client *http.Client) *ImportLookmlDashboardParams {
	return &ImportLookmlDashboardParams{
		HTTPClient: client,
	}
}

/* ImportLookmlDashboardParams contains all the parameters to send to the API endpoint
   for the import lookml dashboard operation.

   Typically these are written to a http.Request.
*/
type ImportLookmlDashboardParams struct {

	/* Body.

	   Dashboard
	*/
	Body *models.Dashboard

	/* LookmlDashboardID.

	   Id of LookML dashboard
	*/
	LookmlDashboardID string

	/* RawLocale.

	   If true, and this dashboard is localized, export it with the raw keys, not localized.
	*/
	RawLocale *bool

	/* SpaceID.

	   Id of space to import the dashboard to
	*/
	SpaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import lookml dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportLookmlDashboardParams) WithDefaults() *ImportLookmlDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import lookml dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportLookmlDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithTimeout(timeout time.Duration) *ImportLookmlDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithContext(ctx context.Context) *ImportLookmlDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithHTTPClient(client *http.Client) *ImportLookmlDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithBody(body *models.Dashboard) *ImportLookmlDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetBody(body *models.Dashboard) {
	o.Body = body
}

// WithLookmlDashboardID adds the lookmlDashboardID to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithLookmlDashboardID(lookmlDashboardID string) *ImportLookmlDashboardParams {
	o.SetLookmlDashboardID(lookmlDashboardID)
	return o
}

// SetLookmlDashboardID adds the lookmlDashboardId to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetLookmlDashboardID(lookmlDashboardID string) {
	o.LookmlDashboardID = lookmlDashboardID
}

// WithRawLocale adds the rawLocale to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithRawLocale(rawLocale *bool) *ImportLookmlDashboardParams {
	o.SetRawLocale(rawLocale)
	return o
}

// SetRawLocale adds the rawLocale to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetRawLocale(rawLocale *bool) {
	o.RawLocale = rawLocale
}

// WithSpaceID adds the spaceID to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) WithSpaceID(spaceID string) *ImportLookmlDashboardParams {
	o.SetSpaceID(spaceID)
	return o
}

// SetSpaceID adds the spaceId to the import lookml dashboard params
func (o *ImportLookmlDashboardParams) SetSpaceID(spaceID string) {
	o.SpaceID = spaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ImportLookmlDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param lookml_dashboard_id
	if err := r.SetPathParam("lookml_dashboard_id", o.LookmlDashboardID); err != nil {
		return err
	}

	if o.RawLocale != nil {

		// query param raw_locale
		var qrRawLocale bool

		if o.RawLocale != nil {
			qrRawLocale = *o.RawLocale
		}
		qRawLocale := swag.FormatBool(qrRawLocale)
		if qRawLocale != "" {

			if err := r.SetQueryParam("raw_locale", qRawLocale); err != nil {
				return err
			}
		}
	}

	// path param space_id
	if err := r.SetPathParam("space_id", o.SpaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
