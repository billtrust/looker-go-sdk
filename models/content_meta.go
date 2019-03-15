// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ContentMeta content meta
// swagger:model ContentMeta
type ContentMeta struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Content Type ("dashboard", "look", or "space")
	// Read Only: true
	ContentType string `json:"content_type,omitempty"`

	// Id of associated dashboard when content_type is "dashboard"
	// Read Only: true
	DashboardID string `json:"dashboard_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Id of Inherited Content
	// Read Only: true
	InheritingID int64 `json:"inheriting_id,omitempty"`

	// Whether content inherits its access levels from parent
	// bmccarthy need to change to pointer so it can be set to false correctly
	Inherits *bool `json:"inherits,omitempty"`

	// Id of associated look when content_type is "look"
	// Read Only: true
	LookID int64 `json:"look_id,omitempty"`

	// Name or title of underlying content
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Id of Parent Content
	// Read Only: true
	ParentID int64 `json:"parent_id,omitempty"`

	// Content Slug
	// Read Only: true
	Slug string `json:"slug,omitempty"`

	// Id of associated space when content_type is "space"
	// Read Only: true
	// bmccarthy API actually returns int64 and not string
	SpaceID int64 `json:"space_id,omitempty"`
}

// Validate validates this content meta
func (m *ContentMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentMeta) UnmarshalBinary(b []byte) error {
	var res ContentMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
