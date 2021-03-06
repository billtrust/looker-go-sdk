// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DBConnectionTestResult d b connection test result
// swagger:model DBConnectionTestResult
type DBConnectionTestResult struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// JDBC connection string. (only populated in the 'connect' test)
	// Read Only: true
	ConnectionString string `json:"connection_string,omitempty"`

	// Result message of test
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Name of test
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Result code of test
	// Read Only: true
	Status string `json:"status,omitempty"`
}

// Validate validates this d b connection test result
func (m *DBConnectionTestResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DBConnectionTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DBConnectionTestResult) UnmarshalBinary(b []byte) error {
	var res DBConnectionTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
