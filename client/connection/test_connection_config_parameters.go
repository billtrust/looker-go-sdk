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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// NewTestConnectionConfigParams creates a new TestConnectionConfigParams object
// with the default values initialized.
func NewTestConnectionConfigParams() *TestConnectionConfigParams {
	var ()
	return &TestConnectionConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestConnectionConfigParamsWithTimeout creates a new TestConnectionConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestConnectionConfigParamsWithTimeout(timeout time.Duration) *TestConnectionConfigParams {
	var ()
	return &TestConnectionConfigParams{

		timeout: timeout,
	}
}

// NewTestConnectionConfigParamsWithContext creates a new TestConnectionConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestConnectionConfigParamsWithContext(ctx context.Context) *TestConnectionConfigParams {
	var ()
	return &TestConnectionConfigParams{

		Context: ctx,
	}
}

// NewTestConnectionConfigParamsWithHTTPClient creates a new TestConnectionConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestConnectionConfigParamsWithHTTPClient(client *http.Client) *TestConnectionConfigParams {
	var ()
	return &TestConnectionConfigParams{
		HTTPClient: client,
	}
}

/*TestConnectionConfigParams contains all the parameters to send to the API endpoint
for the test connection config operation typically these are written to a http.Request
*/
type TestConnectionConfigParams struct {

	/*Body
	  Connection

	*/
	Body *models.DBConnection
	/*Tests
	  Array of names of tests to run

	*/
	Tests []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test connection config params
func (o *TestConnectionConfigParams) WithTimeout(timeout time.Duration) *TestConnectionConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test connection config params
func (o *TestConnectionConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test connection config params
func (o *TestConnectionConfigParams) WithContext(ctx context.Context) *TestConnectionConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test connection config params
func (o *TestConnectionConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test connection config params
func (o *TestConnectionConfigParams) WithHTTPClient(client *http.Client) *TestConnectionConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test connection config params
func (o *TestConnectionConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test connection config params
func (o *TestConnectionConfigParams) WithBody(body *models.DBConnection) *TestConnectionConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test connection config params
func (o *TestConnectionConfigParams) SetBody(body *models.DBConnection) {
	o.Body = body
}

// WithTests adds the tests to the test connection config params
func (o *TestConnectionConfigParams) WithTests(tests []string) *TestConnectionConfigParams {
	o.SetTests(tests)
	return o
}

// SetTests adds the tests to the test connection config params
func (o *TestConnectionConfigParams) SetTests(tests []string) {
	o.Tests = tests
}

// WriteToRequest writes these params to a swagger request
func (o *TestConnectionConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesTests := o.Tests

	joinedTests := swag.JoinByFormat(valuesTests, "csv")
	// query array param tests
	if err := r.SetQueryParam("tests", joinedTests...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
