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
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// NewTestLdapConfigAuthParams creates a new TestLdapConfigAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestLdapConfigAuthParams() *TestLdapConfigAuthParams {
	return &TestLdapConfigAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestLdapConfigAuthParamsWithTimeout creates a new TestLdapConfigAuthParams object
// with the ability to set a timeout on a request.
func NewTestLdapConfigAuthParamsWithTimeout(timeout time.Duration) *TestLdapConfigAuthParams {
	return &TestLdapConfigAuthParams{
		timeout: timeout,
	}
}

// NewTestLdapConfigAuthParamsWithContext creates a new TestLdapConfigAuthParams object
// with the ability to set a context for a request.
func NewTestLdapConfigAuthParamsWithContext(ctx context.Context) *TestLdapConfigAuthParams {
	return &TestLdapConfigAuthParams{
		Context: ctx,
	}
}

// NewTestLdapConfigAuthParamsWithHTTPClient creates a new TestLdapConfigAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestLdapConfigAuthParamsWithHTTPClient(client *http.Client) *TestLdapConfigAuthParams {
	return &TestLdapConfigAuthParams{
		HTTPClient: client,
	}
}

/* TestLdapConfigAuthParams contains all the parameters to send to the API endpoint
   for the test ldap config auth operation.

   Typically these are written to a http.Request.
*/
type TestLdapConfigAuthParams struct {

	/* Body.

	   LDAP Config
	*/
	Body *models.LDAPConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test ldap config auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestLdapConfigAuthParams) WithDefaults() *TestLdapConfigAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test ldap config auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestLdapConfigAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test ldap config auth params
func (o *TestLdapConfigAuthParams) WithTimeout(timeout time.Duration) *TestLdapConfigAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test ldap config auth params
func (o *TestLdapConfigAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test ldap config auth params
func (o *TestLdapConfigAuthParams) WithContext(ctx context.Context) *TestLdapConfigAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test ldap config auth params
func (o *TestLdapConfigAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test ldap config auth params
func (o *TestLdapConfigAuthParams) WithHTTPClient(client *http.Client) *TestLdapConfigAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test ldap config auth params
func (o *TestLdapConfigAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test ldap config auth params
func (o *TestLdapConfigAuthParams) WithBody(body *models.LDAPConfig) *TestLdapConfigAuthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test ldap config auth params
func (o *TestLdapConfigAuthParams) SetBody(body *models.LDAPConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConfigAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
