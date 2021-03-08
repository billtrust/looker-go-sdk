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

// ContentValidationError content validation error
//
// swagger:model ContentValidationError
type ContentValidationError struct {

	// Name of the explore involved in the error
	// Read Only: true
	ExploreName string `json:"explore_name,omitempty"`

	// Name of the field involved in the error
	// Read Only: true
	FieldName string `json:"field_name,omitempty"`

	// Error message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Name of the model involved in the error
	// Read Only: true
	ModelName string `json:"model_name,omitempty"`

	// Whether this validation error is removable
	// Read Only: true
	Removable *bool `json:"removable,omitempty"`
}

// Validate validates this content validation error
func (m *ContentValidationError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this content validation error based on the context it is used
func (m *ContentValidationError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExploreName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFieldName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModelName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemovable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidationError) contextValidateExploreName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "explore_name", "body", string(m.ExploreName)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidationError) contextValidateFieldName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "field_name", "body", string(m.FieldName)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidationError) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidationError) contextValidateModelName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "model_name", "body", string(m.ModelName)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidationError) contextValidateRemovable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "removable", "body", m.Removable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentValidationError) UnmarshalBinary(b []byte) error {
	var res ContentValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
