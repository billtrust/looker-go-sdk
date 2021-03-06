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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// NewUpdateModelSetParams creates a new UpdateModelSetParams object
// with the default values initialized.
func NewUpdateModelSetParams() *UpdateModelSetParams {
	var ()
	return &UpdateModelSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateModelSetParamsWithTimeout creates a new UpdateModelSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateModelSetParamsWithTimeout(timeout time.Duration) *UpdateModelSetParams {
	var ()
	return &UpdateModelSetParams{

		timeout: timeout,
	}
}

// NewUpdateModelSetParamsWithContext creates a new UpdateModelSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateModelSetParamsWithContext(ctx context.Context) *UpdateModelSetParams {
	var ()
	return &UpdateModelSetParams{

		Context: ctx,
	}
}

// NewUpdateModelSetParamsWithHTTPClient creates a new UpdateModelSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateModelSetParamsWithHTTPClient(client *http.Client) *UpdateModelSetParams {
	var ()
	return &UpdateModelSetParams{
		HTTPClient: client,
	}
}

/*UpdateModelSetParams contains all the parameters to send to the API endpoint
for the update model set operation typically these are written to a http.Request
*/
type UpdateModelSetParams struct {

	/*Body
	  ModelSet

	*/
	Body *models.ModelSet
	/*ModelSetID
	  id of model set

	*/
	ModelSetID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update model set params
func (o *UpdateModelSetParams) WithTimeout(timeout time.Duration) *UpdateModelSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update model set params
func (o *UpdateModelSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update model set params
func (o *UpdateModelSetParams) WithContext(ctx context.Context) *UpdateModelSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update model set params
func (o *UpdateModelSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update model set params
func (o *UpdateModelSetParams) WithHTTPClient(client *http.Client) *UpdateModelSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update model set params
func (o *UpdateModelSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update model set params
func (o *UpdateModelSetParams) WithBody(body *models.ModelSet) *UpdateModelSetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update model set params
func (o *UpdateModelSetParams) SetBody(body *models.ModelSet) {
	o.Body = body
}

// WithModelSetID adds the modelSetID to the update model set params
func (o *UpdateModelSetParams) WithModelSetID(modelSetID int64) *UpdateModelSetParams {
	o.SetModelSetID(modelSetID)
	return o
}

// SetModelSetID adds the modelSetId to the update model set params
func (o *UpdateModelSetParams) SetModelSetID(modelSetID int64) {
	o.ModelSetID = modelSetID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateModelSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param model_set_id
	if err := r.SetPathParam("model_set_id", swag.FormatInt64(o.ModelSetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
