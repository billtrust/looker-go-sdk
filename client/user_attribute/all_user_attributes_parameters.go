// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

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

// NewAllUserAttributesParams creates a new AllUserAttributesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllUserAttributesParams() *AllUserAttributesParams {
	return &AllUserAttributesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllUserAttributesParamsWithTimeout creates a new AllUserAttributesParams object
// with the ability to set a timeout on a request.
func NewAllUserAttributesParamsWithTimeout(timeout time.Duration) *AllUserAttributesParams {
	return &AllUserAttributesParams{
		timeout: timeout,
	}
}

// NewAllUserAttributesParamsWithContext creates a new AllUserAttributesParams object
// with the ability to set a context for a request.
func NewAllUserAttributesParamsWithContext(ctx context.Context) *AllUserAttributesParams {
	return &AllUserAttributesParams{
		Context: ctx,
	}
}

// NewAllUserAttributesParamsWithHTTPClient creates a new AllUserAttributesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllUserAttributesParamsWithHTTPClient(client *http.Client) *AllUserAttributesParams {
	return &AllUserAttributesParams{
		HTTPClient: client,
	}
}

/* AllUserAttributesParams contains all the parameters to send to the API endpoint
   for the all user attributes operation.

   Typically these are written to a http.Request.
*/
type AllUserAttributesParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* Sorts.

	   Fields to order the results by. Sortable fields include: name, label
	*/
	Sorts *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all user attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllUserAttributesParams) WithDefaults() *AllUserAttributesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all user attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllUserAttributesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all user attributes params
func (o *AllUserAttributesParams) WithTimeout(timeout time.Duration) *AllUserAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all user attributes params
func (o *AllUserAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all user attributes params
func (o *AllUserAttributesParams) WithContext(ctx context.Context) *AllUserAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all user attributes params
func (o *AllUserAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all user attributes params
func (o *AllUserAttributesParams) WithHTTPClient(client *http.Client) *AllUserAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all user attributes params
func (o *AllUserAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all user attributes params
func (o *AllUserAttributesParams) WithFields(fields *string) *AllUserAttributesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all user attributes params
func (o *AllUserAttributesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSorts adds the sorts to the all user attributes params
func (o *AllUserAttributesParams) WithSorts(sorts *string) *AllUserAttributesParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the all user attributes params
func (o *AllUserAttributesParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WriteToRequest writes these params to a swagger request
func (o *AllUserAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string

		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {

			if err := r.SetQueryParam("sorts", qSorts); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
