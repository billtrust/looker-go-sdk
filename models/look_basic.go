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

// LookBasic look basic
//
// swagger:model LookBasic
type LookBasic struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Look Title
	// Read Only: true
	Title string `json:"title,omitempty"`
}

// Validate validates this look basic
func (m *LookBasic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this look basic based on the context it is used
func (m *LookBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentMetadataID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookBasic) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *LookBasic) contextValidateContentMetadataID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "content_metadata_id", "body", int64(m.ContentMetadataID)); err != nil {
		return err
	}

	return nil
}

func (m *LookBasic) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *LookBasic) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "title", "body", string(m.Title)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookBasic) UnmarshalBinary(b []byte) error {
	var res LookBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
