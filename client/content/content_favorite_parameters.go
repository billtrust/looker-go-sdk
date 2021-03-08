// Code generated by go-swagger; DO NOT EDIT.

package content

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
)

// NewContentFavoriteParams creates a new ContentFavoriteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContentFavoriteParams() *ContentFavoriteParams {
	return &ContentFavoriteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContentFavoriteParamsWithTimeout creates a new ContentFavoriteParams object
// with the ability to set a timeout on a request.
func NewContentFavoriteParamsWithTimeout(timeout time.Duration) *ContentFavoriteParams {
	return &ContentFavoriteParams{
		timeout: timeout,
	}
}

// NewContentFavoriteParamsWithContext creates a new ContentFavoriteParams object
// with the ability to set a context for a request.
func NewContentFavoriteParamsWithContext(ctx context.Context) *ContentFavoriteParams {
	return &ContentFavoriteParams{
		Context: ctx,
	}
}

// NewContentFavoriteParamsWithHTTPClient creates a new ContentFavoriteParams object
// with the ability to set a custom HTTPClient for a request.
func NewContentFavoriteParamsWithHTTPClient(client *http.Client) *ContentFavoriteParams {
	return &ContentFavoriteParams{
		HTTPClient: client,
	}
}

/* ContentFavoriteParams contains all the parameters to send to the API endpoint
   for the content favorite operation.

   Typically these are written to a http.Request.
*/
type ContentFavoriteParams struct {

	/* ContentFavoriteID.

	   Id of favorite content

	   Format: int64
	*/
	ContentFavoriteID int64

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the content favorite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContentFavoriteParams) WithDefaults() *ContentFavoriteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the content favorite params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContentFavoriteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the content favorite params
func (o *ContentFavoriteParams) WithTimeout(timeout time.Duration) *ContentFavoriteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the content favorite params
func (o *ContentFavoriteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the content favorite params
func (o *ContentFavoriteParams) WithContext(ctx context.Context) *ContentFavoriteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the content favorite params
func (o *ContentFavoriteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the content favorite params
func (o *ContentFavoriteParams) WithHTTPClient(client *http.Client) *ContentFavoriteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the content favorite params
func (o *ContentFavoriteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentFavoriteID adds the contentFavoriteID to the content favorite params
func (o *ContentFavoriteParams) WithContentFavoriteID(contentFavoriteID int64) *ContentFavoriteParams {
	o.SetContentFavoriteID(contentFavoriteID)
	return o
}

// SetContentFavoriteID adds the contentFavoriteId to the content favorite params
func (o *ContentFavoriteParams) SetContentFavoriteID(contentFavoriteID int64) {
	o.ContentFavoriteID = contentFavoriteID
}

// WithFields adds the fields to the content favorite params
func (o *ContentFavoriteParams) WithFields(fields *string) *ContentFavoriteParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the content favorite params
func (o *ContentFavoriteParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ContentFavoriteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param content_favorite_id
	if err := r.SetPathParam("content_favorite_id", swag.FormatInt64(o.ContentFavoriteID)); err != nil {
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
