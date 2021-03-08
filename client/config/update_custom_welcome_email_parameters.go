// Code generated by go-swagger; DO NOT EDIT.

package config

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

	"your-damain.com/swagger/looker-api-golang/models"
)

// NewUpdateCustomWelcomeEmailParams creates a new UpdateCustomWelcomeEmailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCustomWelcomeEmailParams() *UpdateCustomWelcomeEmailParams {
	return &UpdateCustomWelcomeEmailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCustomWelcomeEmailParamsWithTimeout creates a new UpdateCustomWelcomeEmailParams object
// with the ability to set a timeout on a request.
func NewUpdateCustomWelcomeEmailParamsWithTimeout(timeout time.Duration) *UpdateCustomWelcomeEmailParams {
	return &UpdateCustomWelcomeEmailParams{
		timeout: timeout,
	}
}

// NewUpdateCustomWelcomeEmailParamsWithContext creates a new UpdateCustomWelcomeEmailParams object
// with the ability to set a context for a request.
func NewUpdateCustomWelcomeEmailParamsWithContext(ctx context.Context) *UpdateCustomWelcomeEmailParams {
	return &UpdateCustomWelcomeEmailParams{
		Context: ctx,
	}
}

// NewUpdateCustomWelcomeEmailParamsWithHTTPClient creates a new UpdateCustomWelcomeEmailParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCustomWelcomeEmailParamsWithHTTPClient(client *http.Client) *UpdateCustomWelcomeEmailParams {
	return &UpdateCustomWelcomeEmailParams{
		HTTPClient: client,
	}
}

/* UpdateCustomWelcomeEmailParams contains all the parameters to send to the API endpoint
   for the update custom welcome email operation.

   Typically these are written to a http.Request.
*/
type UpdateCustomWelcomeEmailParams struct {

	/* Body.

	   Custom Welcome Email setting and value to save
	*/
	Body *models.CustomWelcomeEmail

	/* SendTestWelcomeEmail.

	   If true a test email with the content from the request will be sent to the current user after saving
	*/
	SendTestWelcomeEmail *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update custom welcome email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomWelcomeEmailParams) WithDefaults() *UpdateCustomWelcomeEmailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update custom welcome email params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCustomWelcomeEmailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) WithTimeout(timeout time.Duration) *UpdateCustomWelcomeEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) WithContext(ctx context.Context) *UpdateCustomWelcomeEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) WithHTTPClient(client *http.Client) *UpdateCustomWelcomeEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) WithBody(body *models.CustomWelcomeEmail) *UpdateCustomWelcomeEmailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) SetBody(body *models.CustomWelcomeEmail) {
	o.Body = body
}

// WithSendTestWelcomeEmail adds the sendTestWelcomeEmail to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) WithSendTestWelcomeEmail(sendTestWelcomeEmail *bool) *UpdateCustomWelcomeEmailParams {
	o.SetSendTestWelcomeEmail(sendTestWelcomeEmail)
	return o
}

// SetSendTestWelcomeEmail adds the sendTestWelcomeEmail to the update custom welcome email params
func (o *UpdateCustomWelcomeEmailParams) SetSendTestWelcomeEmail(sendTestWelcomeEmail *bool) {
	o.SendTestWelcomeEmail = sendTestWelcomeEmail
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCustomWelcomeEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.SendTestWelcomeEmail != nil {

		// query param send_test_welcome_email
		var qrSendTestWelcomeEmail bool

		if o.SendTestWelcomeEmail != nil {
			qrSendTestWelcomeEmail = *o.SendTestWelcomeEmail
		}
		qSendTestWelcomeEmail := swag.FormatBool(qrSendTestWelcomeEmail)
		if qSendTestWelcomeEmail != "" {

			if err := r.SetQueryParam("send_test_welcome_email", qSendTestWelcomeEmail); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
