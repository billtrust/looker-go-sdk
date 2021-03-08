// Code generated by go-swagger; DO NOT EDIT.

package connection

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

// NewTestConnectionParams creates a new TestConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestConnectionParams() *TestConnectionParams {
	return &TestConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestConnectionParamsWithTimeout creates a new TestConnectionParams object
// with the ability to set a timeout on a request.
func NewTestConnectionParamsWithTimeout(timeout time.Duration) *TestConnectionParams {
	return &TestConnectionParams{
		timeout: timeout,
	}
}

// NewTestConnectionParamsWithContext creates a new TestConnectionParams object
// with the ability to set a context for a request.
func NewTestConnectionParamsWithContext(ctx context.Context) *TestConnectionParams {
	return &TestConnectionParams{
		Context: ctx,
	}
}

// NewTestConnectionParamsWithHTTPClient creates a new TestConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestConnectionParamsWithHTTPClient(client *http.Client) *TestConnectionParams {
	return &TestConnectionParams{
		HTTPClient: client,
	}
}

/* TestConnectionParams contains all the parameters to send to the API endpoint
   for the test connection operation.

   Typically these are written to a http.Request.
*/
type TestConnectionParams struct {

	/* ConnectionName.

	   Name of connection
	*/
	ConnectionName string

	/* Tests.

	   Array of names of tests to run
	*/
	Tests []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestConnectionParams) WithDefaults() *TestConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestConnectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test connection params
func (o *TestConnectionParams) WithTimeout(timeout time.Duration) *TestConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test connection params
func (o *TestConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test connection params
func (o *TestConnectionParams) WithContext(ctx context.Context) *TestConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test connection params
func (o *TestConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test connection params
func (o *TestConnectionParams) WithHTTPClient(client *http.Client) *TestConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test connection params
func (o *TestConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionName adds the connectionName to the test connection params
func (o *TestConnectionParams) WithConnectionName(connectionName string) *TestConnectionParams {
	o.SetConnectionName(connectionName)
	return o
}

// SetConnectionName adds the connectionName to the test connection params
func (o *TestConnectionParams) SetConnectionName(connectionName string) {
	o.ConnectionName = connectionName
}

// WithTests adds the tests to the test connection params
func (o *TestConnectionParams) WithTests(tests []string) *TestConnectionParams {
	o.SetTests(tests)
	return o
}

// SetTests adds the tests to the test connection params
func (o *TestConnectionParams) SetTests(tests []string) {
	o.Tests = tests
}

// WriteToRequest writes these params to a swagger request
func (o *TestConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_name
	if err := r.SetPathParam("connection_name", o.ConnectionName); err != nil {
		return err
	}

	if o.Tests != nil {

		// binding items for tests
		joinedTests := o.bindParamTests(reg)

		// query array param tests
		if err := r.SetQueryParam("tests", joinedTests...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTestConnection binds the parameter tests
func (o *TestConnectionParams) bindParamTests(formats strfmt.Registry) []string {
	testsIR := o.Tests

	var testsIC []string
	for _, testsIIR := range testsIR { // explode []string

		testsIIV := testsIIR // string as string
		testsIC = append(testsIC, testsIIV)
	}

	// items.CollectionFormat: "csv"
	testsIS := swag.JoinByFormat(testsIC, "csv")

	return testsIS
}
