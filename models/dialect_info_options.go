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

// DialectInfoOptions dialect info options
//
// swagger:model DialectInfoOptions
type DialectInfoOptions struct {

	// Has additional params support
	// Read Only: true
	AdditionalParams *bool `json:"additional_params,omitempty"`

	// Has auth support
	// Read Only: true
	Auth *bool `json:"auth,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Has host support
	// Read Only: true
	Host *bool `json:"host,omitempty"`

	// Has support for a service account
	// Read Only: true
	OauthCredentials *bool `json:"oauth_credentials,omitempty"`

	// Has project name support
	// Read Only: true
	ProjectName *bool `json:"project_name,omitempty"`

	// Has schema support
	// Read Only: true
	Schema *bool `json:"schema,omitempty"`

	// Has SSL support
	// Read Only: true
	Ssl *bool `json:"ssl,omitempty"`

	// Has timezone support
	// Read Only: true
	Timezone *bool `json:"timezone,omitempty"`

	// Has tmp table support
	// Read Only: true
	TmpTable *bool `json:"tmp_table,omitempty"`

	// Username is required
	// Read Only: true
	UsernameRequired *bool `json:"username_required,omitempty"`
}

// Validate validates this dialect info options
func (m *DialectInfoOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this dialect info options based on the context it is used
func (m *DialectInfoOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOauthCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSsl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimezone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTmpTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsernameRequired(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DialectInfoOptions) contextValidateAdditionalParams(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "additional_params", "body", m.AdditionalParams); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateAuth(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "auth", "body", m.Auth); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *DialectInfoOptions) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateOauthCredentials(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "oauth_credentials", "body", m.OauthCredentials); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateProjectName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "project_name", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "schema", "body", m.Schema); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateSsl(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ssl", "body", m.Ssl); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateTimezone(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timezone", "body", m.Timezone); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateTmpTable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tmp_table", "body", m.TmpTable); err != nil {
		return err
	}

	return nil
}

func (m *DialectInfoOptions) contextValidateUsernameRequired(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "username_required", "body", m.UsernameRequired); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DialectInfoOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DialectInfoOptions) UnmarshalBinary(b []byte) error {
	var res DialectInfoOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
