// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewCreateModelSetParams creates a new CreateModelSetParams object
// with the default values initialized.
func NewCreateModelSetParams() *CreateModelSetParams {
	var ()
	return &CreateModelSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateModelSetParamsWithTimeout creates a new CreateModelSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateModelSetParamsWithTimeout(timeout time.Duration) *CreateModelSetParams {
	var ()
	return &CreateModelSetParams{

		timeout: timeout,
	}
}

// NewCreateModelSetParamsWithContext creates a new CreateModelSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateModelSetParamsWithContext(ctx context.Context) *CreateModelSetParams {
	var ()
	return &CreateModelSetParams{

		Context: ctx,
	}
}

// NewCreateModelSetParamsWithHTTPClient creates a new CreateModelSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateModelSetParamsWithHTTPClient(client *http.Client) *CreateModelSetParams {
	var ()
	return &CreateModelSetParams{
		HTTPClient: client,
	}
}

/*CreateModelSetParams contains all the parameters to send to the API endpoint
for the create model set operation typically these are written to a http.Request
*/
type CreateModelSetParams struct {

	/*Body
	  ModelSet

	*/
	Body *models.ModelSet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create model set params
func (o *CreateModelSetParams) WithTimeout(timeout time.Duration) *CreateModelSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create model set params
func (o *CreateModelSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create model set params
func (o *CreateModelSetParams) WithContext(ctx context.Context) *CreateModelSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create model set params
func (o *CreateModelSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create model set params
func (o *CreateModelSetParams) WithHTTPClient(client *http.Client) *CreateModelSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create model set params
func (o *CreateModelSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create model set params
func (o *CreateModelSetParams) WithBody(body *models.ModelSet) *CreateModelSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create model set params
func (o *CreateModelSetParams) SetBody(body *models.ModelSet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateModelSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
