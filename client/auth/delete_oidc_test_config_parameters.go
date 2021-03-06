// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteOidcTestConfigParams creates a new DeleteOidcTestConfigParams object
// with the default values initialized.
func NewDeleteOidcTestConfigParams() *DeleteOidcTestConfigParams {
	var ()
	return &DeleteOidcTestConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOidcTestConfigParamsWithTimeout creates a new DeleteOidcTestConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOidcTestConfigParamsWithTimeout(timeout time.Duration) *DeleteOidcTestConfigParams {
	var ()
	return &DeleteOidcTestConfigParams{

		timeout: timeout,
	}
}

// NewDeleteOidcTestConfigParamsWithContext creates a new DeleteOidcTestConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOidcTestConfigParamsWithContext(ctx context.Context) *DeleteOidcTestConfigParams {
	var ()
	return &DeleteOidcTestConfigParams{

		Context: ctx,
	}
}

// NewDeleteOidcTestConfigParamsWithHTTPClient creates a new DeleteOidcTestConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOidcTestConfigParamsWithHTTPClient(client *http.Client) *DeleteOidcTestConfigParams {
	var ()
	return &DeleteOidcTestConfigParams{
		HTTPClient: client,
	}
}

/*DeleteOidcTestConfigParams contains all the parameters to send to the API endpoint
for the delete oidc test config operation typically these are written to a http.Request
*/
type DeleteOidcTestConfigParams struct {

	/*TestSlug
	  Slug of test config

	*/
	TestSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) WithTimeout(timeout time.Duration) *DeleteOidcTestConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) WithContext(ctx context.Context) *DeleteOidcTestConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) WithHTTPClient(client *http.Client) *DeleteOidcTestConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTestSlug adds the testSlug to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) WithTestSlug(testSlug string) *DeleteOidcTestConfigParams {
	o.SetTestSlug(testSlug)
	return o
}

// SetTestSlug adds the testSlug to the delete oidc test config params
func (o *DeleteOidcTestConfigParams) SetTestSlug(testSlug string) {
	o.TestSlug = testSlug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOidcTestConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param test_slug
	if err := r.SetPathParam("test_slug", o.TestSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
