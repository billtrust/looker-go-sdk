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

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// NewTestLdapConfigConnectionParams creates a new TestLdapConfigConnectionParams object
// with the default values initialized.
func NewTestLdapConfigConnectionParams() *TestLdapConfigConnectionParams {
	var ()
	return &TestLdapConfigConnectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestLdapConfigConnectionParamsWithTimeout creates a new TestLdapConfigConnectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestLdapConfigConnectionParamsWithTimeout(timeout time.Duration) *TestLdapConfigConnectionParams {
	var ()
	return &TestLdapConfigConnectionParams{

		timeout: timeout,
	}
}

// NewTestLdapConfigConnectionParamsWithContext creates a new TestLdapConfigConnectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestLdapConfigConnectionParamsWithContext(ctx context.Context) *TestLdapConfigConnectionParams {
	var ()
	return &TestLdapConfigConnectionParams{

		Context: ctx,
	}
}

// NewTestLdapConfigConnectionParamsWithHTTPClient creates a new TestLdapConfigConnectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestLdapConfigConnectionParamsWithHTTPClient(client *http.Client) *TestLdapConfigConnectionParams {
	var ()
	return &TestLdapConfigConnectionParams{
		HTTPClient: client,
	}
}

/*TestLdapConfigConnectionParams contains all the parameters to send to the API endpoint
for the test ldap config connection operation typically these are written to a http.Request
*/
type TestLdapConfigConnectionParams struct {

	/*Body
	  LDAP Config

	*/
	Body *models.LDAPConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) WithTimeout(timeout time.Duration) *TestLdapConfigConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) WithContext(ctx context.Context) *TestLdapConfigConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) WithHTTPClient(client *http.Client) *TestLdapConfigConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) WithBody(body *models.LDAPConfig) *TestLdapConfigConnectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test ldap config connection params
func (o *TestLdapConfigConnectionParams) SetBody(body *models.LDAPConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConfigConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
