// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataActionFormField data action form field
//
// swagger:model DataActionFormField
type DataActionFormField struct {

	// Default value of the field.
	// Read Only: true
	Default string `json:"default,omitempty"`

	// Description of field
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Whether or not a field supports interactive forms.
	// Read Only: true
	Interactive *bool `json:"interactive,omitempty"`

	// Human-readable label
	// Read Only: true
	Label string `json:"label,omitempty"`

	// Name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The URL for an oauth link, if type is 'oauth_link'.
	// Read Only: true
	OauthURL string `json:"oauth_url,omitempty"`

	// If the form type is 'select', a list of options to be selected from.
	// Read Only: true
	Options []*DataActionFormSelectOption `json:"options"`

	// Whether or not the field is required. This is a user-interface hint. A user interface displaying this form should not submit it without a value for this field. The action server must also perform this validation.
	// Read Only: true
	Required *bool `json:"required,omitempty"`

	// Type of field.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this data action form field
func (m *DataActionFormField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataActionFormField) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {
		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {
			if err := m.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this data action form field based on the context it is used
func (m *DataActionFormField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInteractive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOauthURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequired(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataActionFormField) contextValidateDefault(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "default", "body", string(m.Default)); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateInteractive(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interactive", "body", m.Interactive); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateLabel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "label", "body", string(m.Label)); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateOauthURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "oauth_url", "body", string(m.OauthURL)); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "options", "body", []*DataActionFormSelectOption(m.Options)); err != nil {
		return err
	}

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {
			if err := m.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataActionFormField) contextValidateRequired(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

func (m *DataActionFormField) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataActionFormField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataActionFormField) UnmarshalBinary(b []byte) error {
	var res DataActionFormField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
