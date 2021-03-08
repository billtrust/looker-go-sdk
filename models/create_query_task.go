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

// CreateQueryTask create query task
//
// swagger:model CreateQueryTask
type CreateQueryTask struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Id of dashboard associated with query.
	DashboardID string `json:"dashboard_id,omitempty"`

	// Create the task but defer execution
	Deferred bool `json:"deferred,omitempty"`

	// Id of look associated with query.
	LookID int64 `json:"look_id,omitempty"`

	// Id of query to run
	// Required: true
	QueryID *int64 `json:"query_id"`

	// Desired async query result format. Valid values are: "inline_json", "json", "json_detail", "json_fe", "csv", "html", "md", "txt", "xlsx", "gsxml".
	// Required: true
	ResultFormat *string `json:"result_format"`

	// Source of query task
	Source string `json:"source,omitempty"`
}

// Validate validates this create query task
func (m *CreateQueryTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateQueryTask) validateQueryID(formats strfmt.Registry) error {

	if err := validate.Required("query_id", "body", m.QueryID); err != nil {
		return err
	}

	return nil
}

func (m *CreateQueryTask) validateResultFormat(formats strfmt.Registry) error {

	if err := validate.Required("result_format", "body", m.ResultFormat); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create query task based on the context it is used
func (m *CreateQueryTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateQueryTask) contextValidateCan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CreateQueryTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateQueryTask) UnmarshalBinary(b []byte) error {
	var res CreateQueryTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
