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

// LookmlModelExploreAccessFilter lookml model explore access filter
//
// swagger:model LookmlModelExploreAccessFilter
type LookmlModelExploreAccessFilter struct {

	// Field to be filtered
	// Read Only: true
	Field string `json:"field,omitempty"`

	// User attribute name
	// Read Only: true
	UserAttribute string `json:"user_attribute,omitempty"`
}

// Validate validates this lookml model explore access filter
func (m *LookmlModelExploreAccessFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this lookml model explore access filter based on the context it is used
func (m *LookmlModelExploreAccessFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateField(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserAttribute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookmlModelExploreAccessFilter) contextValidateField(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "field", "body", string(m.Field)); err != nil {
		return err
	}

	return nil
}

func (m *LookmlModelExploreAccessFilter) contextValidateUserAttribute(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_attribute", "body", string(m.UserAttribute)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExploreAccessFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExploreAccessFilter) UnmarshalBinary(b []byte) error {
	var res LookmlModelExploreAccessFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
