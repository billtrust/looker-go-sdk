// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardFilter dashboard filter
// swagger:model DashboardFilter
type DashboardFilter struct {

	// Whether the filter allows multiple filter values
	AllowMultipleValues bool `json:"allow_multiple_values,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Id of Dashboard
	// Read Only: true
	DashboardID string `json:"dashboard_id,omitempty"`

	// Default value of filter
	DefaultValue string `json:"default_value,omitempty"`

	// Dimension of filter (required if type = field)
	Dimension string `json:"dimension,omitempty"`

	// Explore of filter (required if type = field)
	Explore string `json:"explore,omitempty"`

	// Field information
	// Read Only: true
	Field map[string]string `json:"field,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Array of listeners for faceted filters
	ListensToFilters []string `json:"listens_to_filters"`

	// Model of filter (required if type = field)
	Model string `json:"model,omitempty"`

	// Name of filter
	// Required: true
	Name *string `json:"name"`

	// Whether the filter requires a value to run the dashboard
	Required bool `json:"required,omitempty"`

	// Display order of this filter relative to other filters
	Row int64 `json:"row,omitempty"`

	// Title of filter
	// Required: true
	Title *string `json:"title"`

	// Type of filter: one of date, number, string, or field
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this dashboard filter
func (m *DashboardFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardFilter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DashboardFilter) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *DashboardFilter) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardFilter) UnmarshalBinary(b []byte) error {
	var res DashboardFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
