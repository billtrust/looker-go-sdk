// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIVersionElement Api version element
//
// swagger:model ApiVersionElement
type APIVersionElement struct {

	// Full version number including minor version
	// Read Only: true
	FullVersion string `json:"full_version,omitempty"`

	// Status of this version
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Url for swagger.json for this version
	// Read Only: true
	// Format: uri
	SwaggerURL strfmt.URI `json:"swagger_url,omitempty"`

	// Version number as it appears in '/api/xxx/' urls
	// Read Only: true
	Version string `json:"version,omitempty"`
}

// Validate validates this Api version element
func (m *APIVersionElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSwaggerURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIVersionElement) validateSwaggerURL(formats strfmt.Registry) error {
	if swag.IsZero(m.SwaggerURL) { // not required
		return nil
	}

	if err := validate.FormatOf("swagger_url", "body", "uri", m.SwaggerURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Api version element based on the context it is used
func (m *APIVersionElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFullVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwaggerURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIVersionElement) contextValidateFullVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "full_version", "body", string(m.FullVersion)); err != nil {
		return err
	}

	return nil
}

func (m *APIVersionElement) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *APIVersionElement) contextValidateSwaggerURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "swagger_url", "body", strfmt.URI(m.SwaggerURL)); err != nil {
		return err
	}

	return nil
}

func (m *APIVersionElement) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", string(m.Version)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIVersionElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIVersionElement) UnmarshalBinary(b []byte) error {
	var res APIVersionElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
