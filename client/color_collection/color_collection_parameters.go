// Code generated by go-swagger; DO NOT EDIT.

package color_collection

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

// NewColorCollectionParams creates a new ColorCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewColorCollectionParams() *ColorCollectionParams {
	return &ColorCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewColorCollectionParamsWithTimeout creates a new ColorCollectionParams object
// with the ability to set a timeout on a request.
func NewColorCollectionParamsWithTimeout(timeout time.Duration) *ColorCollectionParams {
	return &ColorCollectionParams{
		timeout: timeout,
	}
}

// NewColorCollectionParamsWithContext creates a new ColorCollectionParams object
// with the ability to set a context for a request.
func NewColorCollectionParamsWithContext(ctx context.Context) *ColorCollectionParams {
	return &ColorCollectionParams{
		Context: ctx,
	}
}

// NewColorCollectionParamsWithHTTPClient creates a new ColorCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewColorCollectionParamsWithHTTPClient(client *http.Client) *ColorCollectionParams {
	return &ColorCollectionParams{
		HTTPClient: client,
	}
}

/* ColorCollectionParams contains all the parameters to send to the API endpoint
   for the color collection operation.

   Typically these are written to a http.Request.
*/
type ColorCollectionParams struct {

	/* CollectionID.

	   Id of Color Collection
	*/
	CollectionID string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the color collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColorCollectionParams) WithDefaults() *ColorCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the color collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ColorCollectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the color collection params
func (o *ColorCollectionParams) WithTimeout(timeout time.Duration) *ColorCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the color collection params
func (o *ColorCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the color collection params
func (o *ColorCollectionParams) WithContext(ctx context.Context) *ColorCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the color collection params
func (o *ColorCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the color collection params
func (o *ColorCollectionParams) WithHTTPClient(client *http.Client) *ColorCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the color collection params
func (o *ColorCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionID adds the collectionID to the color collection params
func (o *ColorCollectionParams) WithCollectionID(collectionID string) *ColorCollectionParams {
	o.SetCollectionID(collectionID)
	return o
}

// SetCollectionID adds the collectionId to the color collection params
func (o *ColorCollectionParams) SetCollectionID(collectionID string) {
	o.CollectionID = collectionID
}

// WithFields adds the fields to the color collection params
func (o *ColorCollectionParams) WithFields(fields *string) *ColorCollectionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the color collection params
func (o *ColorCollectionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ColorCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_id
	if err := r.SetPathParam("collection_id", o.CollectionID); err != nil {
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