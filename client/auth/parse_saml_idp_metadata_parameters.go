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
)

// NewParseSamlIdpMetadataParams creates a new ParseSamlIdpMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewParseSamlIdpMetadataParams() *ParseSamlIdpMetadataParams {
	return &ParseSamlIdpMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewParseSamlIdpMetadataParamsWithTimeout creates a new ParseSamlIdpMetadataParams object
// with the ability to set a timeout on a request.
func NewParseSamlIdpMetadataParamsWithTimeout(timeout time.Duration) *ParseSamlIdpMetadataParams {
	return &ParseSamlIdpMetadataParams{
		timeout: timeout,
	}
}

// NewParseSamlIdpMetadataParamsWithContext creates a new ParseSamlIdpMetadataParams object
// with the ability to set a context for a request.
func NewParseSamlIdpMetadataParamsWithContext(ctx context.Context) *ParseSamlIdpMetadataParams {
	return &ParseSamlIdpMetadataParams{
		Context: ctx,
	}
}

// NewParseSamlIdpMetadataParamsWithHTTPClient creates a new ParseSamlIdpMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewParseSamlIdpMetadataParamsWithHTTPClient(client *http.Client) *ParseSamlIdpMetadataParams {
	return &ParseSamlIdpMetadataParams{
		HTTPClient: client,
	}
}

/* ParseSamlIdpMetadataParams contains all the parameters to send to the API endpoint
   for the parse saml idp metadata operation.

   Typically these are written to a http.Request.
*/
type ParseSamlIdpMetadataParams struct {

	/* Body.

	   SAML IdP metadata xml
	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the parse saml idp metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ParseSamlIdpMetadataParams) WithDefaults() *ParseSamlIdpMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the parse saml idp metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ParseSamlIdpMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) WithTimeout(timeout time.Duration) *ParseSamlIdpMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) WithContext(ctx context.Context) *ParseSamlIdpMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) WithHTTPClient(client *http.Client) *ParseSamlIdpMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) WithBody(body string) *ParseSamlIdpMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the parse saml idp metadata params
func (o *ParseSamlIdpMetadataParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ParseSamlIdpMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
