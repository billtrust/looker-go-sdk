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

// ContentValidatorError content validator error
//
// swagger:model ContentValidatorError
type ContentValidatorError struct {

	// alert
	// Read Only: true
	Alert *ContentValidationAlert `json:"alert,omitempty"`

	// dashboard
	// Read Only: true
	Dashboard *ContentValidationDashboard `json:"dashboard,omitempty"`

	// dashboard element
	// Read Only: true
	DashboardElement *ContentValidationDashboardElement `json:"dashboard_element,omitempty"`

	// dashboard filter
	// Read Only: true
	DashboardFilter *ContentValidationDashboardFilter `json:"dashboard_filter,omitempty"`

	// A list of errors found for this piece of content
	// Read Only: true
	Errors []*ContentValidationError `json:"errors"`

	// An id unique to this piece of content for this validation run
	// Read Only: true
	ID string `json:"id,omitempty"`

	// look
	// Read Only: true
	Look *ContentValidationLook `json:"look,omitempty"`

	// lookml dashboard
	// Read Only: true
	LookmlDashboard *ContentValidationLookMLDashboard `json:"lookml_dashboard,omitempty"`

	// lookml dashboard element
	// Read Only: true
	LookmlDashboardElement *ContentValidationLookMLDashboardElement `json:"lookml_dashboard_element,omitempty"`

	// scheduled plan
	// Read Only: true
	ScheduledPlan *ContentValidationScheduledPlan `json:"scheduled_plan,omitempty"`
}

// Validate validates this content validator error
func (m *ContentValidatorError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardElement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLookmlDashboard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLookmlDashboardElement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledPlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidatorError) validateAlert(formats strfmt.Registry) error {
	if swag.IsZero(m.Alert) { // not required
		return nil
	}

	if m.Alert != nil {
		if err := m.Alert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboard) { // not required
		return nil
	}

	if m.Dashboard != nil {
		if err := m.Dashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateDashboardElement(formats strfmt.Registry) error {
	if swag.IsZero(m.DashboardElement) { // not required
		return nil
	}

	if m.DashboardElement != nil {
		if err := m.DashboardElement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard_element")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateDashboardFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.DashboardFilter) { // not required
		return nil
	}

	if m.DashboardFilter != nil {
		if err := m.DashboardFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard_filter")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentValidatorError) validateLook(formats strfmt.Registry) error {
	if swag.IsZero(m.Look) { // not required
		return nil
	}

	if m.Look != nil {
		if err := m.Look.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("look")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateLookmlDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.LookmlDashboard) { // not required
		return nil
	}

	if m.LookmlDashboard != nil {
		if err := m.LookmlDashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lookml_dashboard")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateLookmlDashboardElement(formats strfmt.Registry) error {
	if swag.IsZero(m.LookmlDashboardElement) { // not required
		return nil
	}

	if m.LookmlDashboardElement != nil {
		if err := m.LookmlDashboardElement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lookml_dashboard_element")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) validateScheduledPlan(formats strfmt.Registry) error {
	if swag.IsZero(m.ScheduledPlan) { // not required
		return nil
	}

	if m.ScheduledPlan != nil {
		if err := m.ScheduledPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduled_plan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this content validator error based on the context it is used
func (m *ContentValidatorError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDashboardElement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDashboardFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLook(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLookmlDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLookmlDashboardElement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheduledPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidatorError) contextValidateAlert(ctx context.Context, formats strfmt.Registry) error {

	if m.Alert != nil {
		if err := m.Alert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alert")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateDashboard(ctx context.Context, formats strfmt.Registry) error {

	if m.Dashboard != nil {
		if err := m.Dashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateDashboardElement(ctx context.Context, formats strfmt.Registry) error {

	if m.DashboardElement != nil {
		if err := m.DashboardElement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard_element")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateDashboardFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.DashboardFilter != nil {
		if err := m.DashboardFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dashboard_filter")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "errors", "body", []*ContentValidationError(m.Errors)); err != nil {
		return err
	}

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentValidatorError) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ContentValidatorError) contextValidateLook(ctx context.Context, formats strfmt.Registry) error {

	if m.Look != nil {
		if err := m.Look.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("look")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateLookmlDashboard(ctx context.Context, formats strfmt.Registry) error {

	if m.LookmlDashboard != nil {
		if err := m.LookmlDashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lookml_dashboard")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateLookmlDashboardElement(ctx context.Context, formats strfmt.Registry) error {

	if m.LookmlDashboardElement != nil {
		if err := m.LookmlDashboardElement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lookml_dashboard_element")
			}
			return err
		}
	}

	return nil
}

func (m *ContentValidatorError) contextValidateScheduledPlan(ctx context.Context, formats strfmt.Registry) error {

	if m.ScheduledPlan != nil {
		if err := m.ScheduledPlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduled_plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentValidatorError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentValidatorError) UnmarshalBinary(b []byte) error {
	var res ContentValidatorError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
