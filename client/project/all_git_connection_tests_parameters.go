// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewAllGitConnectionTestsParams creates a new AllGitConnectionTestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllGitConnectionTestsParams() *AllGitConnectionTestsParams {
	return &AllGitConnectionTestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllGitConnectionTestsParamsWithTimeout creates a new AllGitConnectionTestsParams object
// with the ability to set a timeout on a request.
func NewAllGitConnectionTestsParamsWithTimeout(timeout time.Duration) *AllGitConnectionTestsParams {
	return &AllGitConnectionTestsParams{
		timeout: timeout,
	}
}

// NewAllGitConnectionTestsParamsWithContext creates a new AllGitConnectionTestsParams object
// with the ability to set a context for a request.
func NewAllGitConnectionTestsParamsWithContext(ctx context.Context) *AllGitConnectionTestsParams {
	return &AllGitConnectionTestsParams{
		Context: ctx,
	}
}

// NewAllGitConnectionTestsParamsWithHTTPClient creates a new AllGitConnectionTestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllGitConnectionTestsParamsWithHTTPClient(client *http.Client) *AllGitConnectionTestsParams {
	return &AllGitConnectionTestsParams{
		HTTPClient: client,
	}
}

/* AllGitConnectionTestsParams contains all the parameters to send to the API endpoint
   for the all git connection tests operation.

   Typically these are written to a http.Request.
*/
type AllGitConnectionTestsParams struct {

	/* ProjectID.

	   Project Id
	*/
	ProjectID string

	/* RemoteURL.

	   (Optional: leave blank for root project) The remote url for remote dependency to test.
	*/
	RemoteURL *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the all git connection tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllGitConnectionTestsParams) WithDefaults() *AllGitConnectionTestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the all git connection tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllGitConnectionTestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the all git connection tests params
func (o *AllGitConnectionTestsParams) WithTimeout(timeout time.Duration) *AllGitConnectionTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all git connection tests params
func (o *AllGitConnectionTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all git connection tests params
func (o *AllGitConnectionTestsParams) WithContext(ctx context.Context) *AllGitConnectionTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all git connection tests params
func (o *AllGitConnectionTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all git connection tests params
func (o *AllGitConnectionTestsParams) WithHTTPClient(client *http.Client) *AllGitConnectionTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all git connection tests params
func (o *AllGitConnectionTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the all git connection tests params
func (o *AllGitConnectionTestsParams) WithProjectID(projectID string) *AllGitConnectionTestsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the all git connection tests params
func (o *AllGitConnectionTestsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRemoteURL adds the remoteURL to the all git connection tests params
func (o *AllGitConnectionTestsParams) WithRemoteURL(remoteURL *string) *AllGitConnectionTestsParams {
	o.SetRemoteURL(remoteURL)
	return o
}

// SetRemoteURL adds the remoteUrl to the all git connection tests params
func (o *AllGitConnectionTestsParams) SetRemoteURL(remoteURL *string) {
	o.RemoteURL = remoteURL
}

// WriteToRequest writes these params to a swagger request
func (o *AllGitConnectionTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if o.RemoteURL != nil {

		// query param remote_url
		var qrRemoteURL string

		if o.RemoteURL != nil {
			qrRemoteURL = *o.RemoteURL
		}
		qRemoteURL := qrRemoteURL
		if qRemoteURL != "" {

			if err := r.SetQueryParam("remote_url", qRemoteURL); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
