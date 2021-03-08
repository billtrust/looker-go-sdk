// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PasswordConfig password config
//
// swagger:model PasswordConfig
type PasswordConfig struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Minimum number of characters required for a new password.  Must be between 7 and 100
	MinLength int64 `json:"min_length,omitempty"`

	// Require at least one numeric character
	RequireNumeric bool `json:"require_numeric,omitempty"`

	// Require at least one special character
	RequireSpecial bool `json:"require_special,omitempty"`

	// Require at least one uppercase and one lowercase letter
	RequireUpperlower bool `json:"require_upperlower,omitempty"`
}

// Validate validates this password config
func (m *PasswordConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this password config based on the context it is used
func (m *PasswordConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PasswordConfig) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PasswordConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordConfig) UnmarshalBinary(b []byte) error {
	var res PasswordConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}