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

// NewCreateSamlTestConfigParams creates a new CreateSamlTestConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSamlTestConfigParams() *CreateSamlTestConfigParams {
	return &CreateSamlTestConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSamlTestConfigParamsWithTimeout creates a new CreateSamlTestConfigParams object
// with the ability to set a timeout on a request.
func NewCreateSamlTestConfigParamsWithTimeout(timeout time.Duration) *CreateSamlTestConfigParams {
	return &CreateSamlTestConfigParams{
		timeout: timeout,
	}
}

// NewCreateSamlTestConfigParamsWithContext creates a new CreateSamlTestConfigParams object
// with the ability to set a context for a request.
func NewCreateSamlTestConfigParamsWithContext(ctx context.Context) *CreateSamlTestConfigParams {
	return &CreateSamlTestConfigParams{
		Context: ctx,
	}
}

// NewCreateSamlTestConfigParamsWithHTTPClient creates a new CreateSamlTestConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSamlTestConfigParamsWithHTTPClient(client *http.Client) *CreateSamlTestConfigParams {
	return &CreateSamlTestConfigParams{
		HTTPClient: client,
	}
}

/* CreateSamlTestConfigParams contains all the parameters to send to the API endpoint
   for the create saml test config operation.

   Typically these are written to a http.Request.
*/
type CreateSamlTestConfigParams struct {

	/* Body.

	   SAML test config
	*/
	Body *models.SamlConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create saml test config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSamlTestConfigParams) WithDefaults() *CreateSamlTestConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create saml test config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSamlTestConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create saml test config params
func (o *CreateSamlTestConfigParams) WithTimeout(timeout time.Duration) *CreateSamlTestConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create saml test config params
func (o *CreateSamlTestConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create saml test config params
func (o *CreateSamlTestConfigParams) WithContext(ctx context.Context) *CreateSamlTestConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create saml test config params
func (o *CreateSamlTestConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create saml test config params
func (o *CreateSamlTestConfigParams) WithHTTPClient(client *http.Client) *CreateSamlTestConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create saml test config params
func (o *CreateSamlTestConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create saml test config params
func (o *CreateSamlTestConfigParams) WithBody(body *models.SamlConfig) *CreateSamlTestConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create saml test config params
func (o *CreateSamlTestConfigParams) SetBody(body *models.SamlConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSamlTestConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
