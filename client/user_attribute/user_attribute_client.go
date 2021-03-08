// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user attribute API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user attribute API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllUserAttributeGroupValues(params *AllUserAttributeGroupValuesParams, opts ...ClientOption) (*AllUserAttributeGroupValuesOK, error)

	AllUserAttributes(params *AllUserAttributesParams, opts ...ClientOption) (*AllUserAttributesOK, error)

	CreateUserAttribute(params *CreateUserAttributeParams, opts ...ClientOption) (*CreateUserAttributeOK, error)

	DeleteUserAttribute(params *DeleteUserAttributeParams, opts ...ClientOption) (*DeleteUserAttributeNoContent, error)

	SetUserAttributeGroupValues(params *SetUserAttributeGroupValuesParams, opts ...ClientOption) (*SetUserAttributeGroupValuesOK, error)

	UpdateUserAttribute(params *UpdateUserAttributeParams, opts ...ClientOption) (*UpdateUserAttributeOK, error)

	UserAttribute(params *UserAttributeParams, opts ...ClientOption) (*UserAttributeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AllUserAttributeGroupValues gets user attribute group values

  ### Returns all values of a user attribute defined by user groups, in precedence order.

A user may be a member of multiple groups which define different values for a given user attribute.
The order of group-values in the response determines precedence for selecting which group-value applies
to a given user.  For more information, see [Set User Attribute Group Values](#!/UserAttribute/set_user_attribute_group_values).

Results will only include groups that the caller's user account has permission to see.

*/
func (a *Client) AllUserAttributeGroupValues(params *AllUserAttributeGroupValuesParams, opts ...ClientOption) (*AllUserAttributeGroupValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllUserAttributeGroupValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_user_attribute_group_values",
		Method:             "GET",
		PathPattern:        "/user_attributes/{user_attribute_id}/group_values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllUserAttributeGroupValuesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllUserAttributeGroupValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_user_attribute_group_values: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AllUserAttributes gets all user attributes

  ### Get information about all user attributes.

*/
func (a *Client) AllUserAttributes(params *AllUserAttributesParams, opts ...ClientOption) (*AllUserAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllUserAttributesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_user_attributes",
		Method:             "GET",
		PathPattern:        "/user_attributes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllUserAttributesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AllUserAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_user_attributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateUserAttribute creates user attribute

  ### Create a new user attribute

Permission information for a user attribute is conveyed through the `can` and `user_can_edit` fields.
The `user_can_edit` field indicates whether an attribute is user-editable _anywhere_ in the application.
The `can` field gives more granular access information, with the `set_value` child field indicating whether
an attribute's value can be set by [Setting the User Attribute User Value](#!/User/set_user_attribute_user_value).

Note: `name` and `label` fields must be unique across all user attributes in the Looker instance.
Attempting to create a new user attribute with a name or label that duplicates an existing
user attribute will fail with a 422 error.

*/
func (a *Client) CreateUserAttribute(params *CreateUserAttributeParams, opts ...ClientOption) (*CreateUserAttributeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserAttributeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_user_attribute",
		Method:             "POST",
		PathPattern:        "/user_attributes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserAttributeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUserAttributeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_user_attribute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserAttribute deletes user attribute

  ### Delete a user attribute (admin only).

*/
func (a *Client) DeleteUserAttribute(params *DeleteUserAttributeParams, opts ...ClientOption) (*DeleteUserAttributeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserAttributeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_user_attribute",
		Method:             "DELETE",
		PathPattern:        "/user_attributes/{user_attribute_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserAttributeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserAttributeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_user_attribute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetUserAttributeGroupValues sets user attribute group values

  ### Define values for a user attribute across a set of groups, in priority order.

This function defines all values for a user attribute defined by user groups. This is a global setting, potentially affecting
all users in the system. This function replaces any existing group value definitions for the indicated user attribute.

The value of a user attribute for a given user is determined by searching the following locations, in this order:

1. the user's account settings
2. the groups that the user is a member of
3. the default value of the user attribute, if any

The user may be a member of multiple groups which define different values for that user attribute. The order of items in the group_values parameter
determines which group takes priority for that user. Lowest array index wins.

An alternate method to indicate the selection precedence of group-values is to assign numbers to the 'rank' property of each
group-value object in the array. Lowest 'rank' value wins. If you use this technique, you must assign a
rank value to every group-value object in the array.

  To set a user attribute value for a single user, see [Set User Attribute User Value](#!/User/set_user_attribute_user_value).
To set a user attribute value for all members of a group, see [Set User Attribute Group Value](#!/Group/update_user_attribute_group_value).

*/
func (a *Client) SetUserAttributeGroupValues(params *SetUserAttributeGroupValuesParams, opts ...ClientOption) (*SetUserAttributeGroupValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetUserAttributeGroupValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set_user_attribute_group_values",
		Method:             "POST",
		PathPattern:        "/user_attributes/{user_attribute_id}/group_values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetUserAttributeGroupValuesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetUserAttributeGroupValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set_user_attribute_group_values: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUserAttribute updates user attribute

  ### Update a user attribute definition.

*/
func (a *Client) UpdateUserAttribute(params *UpdateUserAttributeParams, opts ...ClientOption) (*UpdateUserAttributeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserAttributeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_user_attribute",
		Method:             "PATCH",
		PathPattern:        "/user_attributes/{user_attribute_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserAttributeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserAttributeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_user_attribute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserAttribute gets user attribute

  ### Get information about a user attribute.

*/
func (a *Client) UserAttribute(params *UserAttributeParams, opts ...ClientOption) (*UserAttributeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAttributeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "user_attribute",
		Method:             "GET",
		PathPattern:        "/user_attributes/{user_attribute_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserAttributeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserAttributeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user_attribute: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
