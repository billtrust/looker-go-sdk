// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateFolder update folder
//
// swagger:model UpdateFolder
type UpdateFolder struct {

	// Unique Name
	Name string `json:"name,omitempty"`

	// Id of Parent. If the parent id is null, this is a root-level entry
	ParentID string `json:"parent_id,omitempty"`
}

// Validate validates this update folder
func (m *UpdateFolder) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update folder based on context it is used
func (m *UpdateFolder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateFolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateFolder) UnmarshalBinary(b []byte) error {
	var res UpdateFolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}