// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewAllUserCredentialsApi3sParams creates a new AllUserCredentialsApi3sParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllUserCredentialsApi3sParams() *AllUserCredentialsApi3sParams {
	return &AllUserCredentialsApi3sParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllUserCredentialsApi3sParamsWithTimeout creates a new AllUserCredentialsApi3sParams object
// with the ability to set a timeout on a request.
func NewAllUserCredentialsApi3sParamsWithTimeout(timeout time.Duration) *AllUserCredentialsApi3sParams {
	return &AllUserCredentialsApi3sParams{
		timeout: timeout,
	}
}

// NewAllUserCredentialsApi3sParamsWithContext creates a new AllUserCredentialsApi3sParams object
// with the ability to set a context for a request.
func NewAllUserCredentialsApi3sParamsWithContext(ctx context.Context) *AllUserCredentialsApi3sParams {
	return &AllUserCredentialsApi3sParams{
		Context: ctx,
	}
}

// NewAllUserCredentialsApi3sParamsWithHTTPClient creates a new AllUserCredentialsApi3sParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllUserCredentialsApi3sParamsWithHTTPClient(client *http.Client) *AllUserCredentialsApi3sParams {
	return &AllUserCredentialsApi3sParams{
		HTTPClient: client,
	}
}

/* AllUserCredentialsApi3sParams contains all the parameters to send to the API endpoint
   for the all user credentials api3s operation.

   Typically these are written to a http.Request.
*/
type AllUserCredentialsApi3sParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* UserID.

	   id of user

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all user credentials api3s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllUserCredentialsApi3sParams) WithDefaults() *AllUserCredentialsApi3sParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all user credentials api3s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllUserCredentialsApi3sParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) WithTimeout(timeout time.Duration) *AllUserCredentialsApi3sParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) WithContext(ctx context.Context) *AllUserCredentialsApi3sParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) WithHTTPClient(client *http.Client) *AllUserCredentialsApi3sParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) WithFields(fields *string) *AllUserCredentialsApi3sParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) WithUserID(userID int64) *AllUserCredentialsApi3sParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the all user credentials api3s params
func (o *AllUserCredentialsApi3sParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AllUserCredentialsApi3sParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
