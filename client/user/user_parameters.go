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

// NewUserParams creates a new UserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserParams() *UserParams {
	return &UserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserParamsWithTimeout creates a new UserParams object
// with the ability to set a timeout on a request.
func NewUserParamsWithTimeout(timeout time.Duration) *UserParams {
	return &UserParams{
		timeout: timeout,
	}
}

// NewUserParamsWithContext creates a new UserParams object
// with the ability to set a context for a request.
func NewUserParamsWithContext(ctx context.Context) *UserParams {
	return &UserParams{
		Context: ctx,
	}
}

// NewUserParamsWithHTTPClient creates a new UserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserParamsWithHTTPClient(client *http.Client) *UserParams {
	return &UserParams{
		HTTPClient: client,
	}
}

/* UserParams contains all the parameters to send to the API endpoint
   for the user operation.

   Typically these are written to a http.Request.
*/
type UserParams struct {

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* UserID.

	   Id of user

	   Format: int64
	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserParams) WithDefaults() *UserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user params
func (o *UserParams) WithTimeout(timeout time.Duration) *UserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user params
func (o *UserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user params
func (o *UserParams) WithContext(ctx context.Context) *UserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user params
func (o *UserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user params
func (o *UserParams) WithHTTPClient(client *http.Client) *UserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user params
func (o *UserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the user params
func (o *UserParams) WithFields(fields *string) *UserParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user params
func (o *UserParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the user params
func (o *UserParams) WithUserID(userID int64) *UserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the user params
func (o *UserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
