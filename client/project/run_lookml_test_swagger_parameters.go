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

// NewRunLookmlTestParams creates a new RunLookmlTestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunLookmlTestParams() *RunLookmlTestParams {
	return &RunLookmlTestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunLookmlTestParamsWithTimeout creates a new RunLookmlTestParams object
// with the ability to set a timeout on a request.
func NewRunLookmlTestParamsWithTimeout(timeout time.Duration) *RunLookmlTestParams {
	return &RunLookmlTestParams{
		timeout: timeout,
	}
}

// NewRunLookmlTestParamsWithContext creates a new RunLookmlTestParams object
// with the ability to set a context for a request.
func NewRunLookmlTestParamsWithContext(ctx context.Context) *RunLookmlTestParams {
	return &RunLookmlTestParams{
		Context: ctx,
	}
}

// NewRunLookmlTestParamsWithHTTPClient creates a new RunLookmlTestParams object
// with the ability to set a custom HTTPClient for a request.
func NewRunLookmlTestParamsWithHTTPClient(client *http.Client) *RunLookmlTestParams {
	return &RunLookmlTestParams{
		HTTPClient: client,
	}
}

/* RunLookmlTestParams contains all the parameters to send to the API endpoint
   for the run lookml test operation.

   Typically these are written to a http.Request.
*/
type RunLookmlTestParams struct {

	/* FileID.

	   File Name
	*/
	FileID *string

	/* Model.

	   Model Name
	*/
	Model *string

	/* ProjectID.

	   Project Id
	*/
	ProjectID string

	/* Test.

	   Test Name
	*/
	Test *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the run lookml test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunLookmlTestParams) WithDefaults() *RunLookmlTestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the run lookml test params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunLookmlTestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the run lookml test params
func (o *RunLookmlTestParams) WithTimeout(timeout time.Duration) *RunLookmlTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the run lookml test params
func (o *RunLookmlTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the run lookml test params
func (o *RunLookmlTestParams) WithContext(ctx context.Context) *RunLookmlTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the run lookml test params
func (o *RunLookmlTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the run lookml test params
func (o *RunLookmlTestParams) WithHTTPClient(client *http.Client) *RunLookmlTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the run lookml test params
func (o *RunLookmlTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileID adds the fileID to the run lookml test params
func (o *RunLookmlTestParams) WithFileID(fileID *string) *RunLookmlTestParams {
	o.SetFileID(fileID)
	return o
}

// SetFileID adds the fileId to the run lookml test params
func (o *RunLookmlTestParams) SetFileID(fileID *string) {
	o.FileID = fileID
}

// WithModel adds the model to the run lookml test params
func (o *RunLookmlTestParams) WithModel(model *string) *RunLookmlTestParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the run lookml test params
func (o *RunLookmlTestParams) SetModel(model *string) {
	o.Model = model
}

// WithProjectID adds the projectID to the run lookml test params
func (o *RunLookmlTestParams) WithProjectID(projectID string) *RunLookmlTestParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the run lookml test params
func (o *RunLookmlTestParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithTest adds the test to the run lookml test params
func (o *RunLookmlTestParams) WithTest(test *string) *RunLookmlTestParams {
	o.SetTest(test)
	return o
}

// SetTest adds the test to the run lookml test params
func (o *RunLookmlTestParams) SetTest(test *string) {
	o.Test = test
}

// WriteToRequest writes these params to a swagger request
func (o *RunLookmlTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileID != nil {

		// query param file_id
		var qrFileID string

		if o.FileID != nil {
			qrFileID = *o.FileID
		}
		qFileID := qrFileID
		if qFileID != "" {

			if err := r.SetQueryParam("file_id", qFileID); err != nil {
				return err
			}
		}
	}

	if o.Model != nil {

		// query param model
		var qrModel string

		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {

			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if o.Test != nil {

		// query param test
		var qrTest string

		if o.Test != nil {
			qrTest = *o.Test
		}
		qTest := qrTest
		if qTest != "" {

			if err := r.SetQueryParam("test", qTest); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
