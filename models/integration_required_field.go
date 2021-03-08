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

// IntegrationRequiredField integration required field
//
// swagger:model IntegrationRequiredField
type IntegrationRequiredField struct {

	// If present, supercedes 'tag' and matches a field that has all of the provided tags.
	// Read Only: true
	AllTags []string `json:"all_tags"`

	// If present, supercedes 'tag' and matches a field that has any of the provided tags.
	// Read Only: true
	AnyTag []string `json:"any_tag"`

	// Matches a field that has this tag.
	// Read Only: true
	Tag string `json:"tag,omitempty"`
}

// Validate validates this integration required field
func (m *IntegrationRequiredField) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this integration required field based on the context it is used
func (m *IntegrationRequiredField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnyTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationRequiredField) contextValidateAllTags(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "all_tags", "body", []string(m.AllTags)); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationRequiredField) contextValidateAnyTag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "any_tag", "body", []string(m.AnyTag)); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationRequiredField) contextValidateTag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tag", "body", string(m.Tag)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationRequiredField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationRequiredField) UnmarshalBinary(b []byte) error {
	var res IntegrationRequiredField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
