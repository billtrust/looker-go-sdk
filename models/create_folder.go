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

// CreateFolder create folder
//
// swagger:model CreateFolder
type CreateFolder struct {

	// Unique Name
	// Required: true
	Name *string `json:"name"`

	// Id of Parent. If the parent id is null, this is a root-level entry
	// Required: true
	ParentID *string `json:"parent_id"`
}

// Validate validates this create folder
func (m *CreateFolder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFolder) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateFolder) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create folder based on context it is used
func (m *CreateFolder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateFolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFolder) UnmarshalBinary(b []byte) error {
	var res CreateFolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}