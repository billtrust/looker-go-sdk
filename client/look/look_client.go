// Code generated by go-swagger; DO NOT EDIT.

package look

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new look API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for look API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AllLooks(params *AllLooksParams, opts ...ClientOption) (*AllLooksOK, error)

	CreateLook(params *CreateLookParams, opts ...ClientOption) (*CreateLookOK, error)

	DeleteLook(params *DeleteLookParams, opts ...ClientOption) (*DeleteLookNoContent, error)

	Look(params *LookParams, opts ...ClientOption) (*LookOK, error)

	RunLook(params *RunLookParams, opts ...ClientOption) (*RunLookOK, error)

	SearchLooks(params *SearchLooksParams, opts ...ClientOption) (*SearchLooksOK, error)

	UpdateLook(params *UpdateLookParams, opts ...ClientOption) (*UpdateLookOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AllLooks gets all looks

  ### Get information about all active Looks

Returns an array of **abbreviated Look objects** describing all the looks that the caller has access to. Soft-deleted Looks are **not** included.

Get the **full details** of a specific look by id with [look(id)](#!/Look/look)

Find **soft-deleted looks** with [search_looks()](#!/Look/search_looks)

*/
func (a *Client) AllLooks(params *AllLooksParams, opts ...ClientOption) (*AllLooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllLooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "all_looks",
		Method:             "GET",
		PathPattern:        "/looks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AllLooksReader{formats: a.formats},
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
	success, ok := result.(*AllLooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for all_looks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateLook creates look

  ### Create a Look

To create a look to display query data, first create the query with [create_query()](#!/Query/create_query)
then assign the query's id to the `query_id` property in the call to `create_look()`.

To place the look into a particular space, assign the space's id to the `space_id` property
in the call to `create_look()`.

*/
func (a *Client) CreateLook(params *CreateLookParams, opts ...ClientOption) (*CreateLookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_look",
		Method:             "POST",
		PathPattern:        "/looks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLookReader{formats: a.formats},
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
	success, ok := result.(*CreateLookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create_look: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLook deletes look

  ### Permanently Delete a Look

This operation **permanently** removes a look from the Looker database.

NOTE: There is no "undo" for this kind of delete.

For information about soft-delete (which can be undone) see [update_look()](#!/Look/update_look).

*/
func (a *Client) DeleteLook(params *DeleteLookParams, opts ...ClientOption) (*DeleteLookNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_look",
		Method:             "DELETE",
		PathPattern:        "/looks/{look_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLookReader{formats: a.formats},
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
	success, ok := result.(*DeleteLookNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_look: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Look gets look

  ### Get a Look.

Returns detailed information about a Look and its associated Query.


*/
func (a *Client) Look(params *LookParams, opts ...ClientOption) (*LookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "look",
		Method:             "GET",
		PathPattern:        "/looks/{look_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LookReader{formats: a.formats},
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
	success, ok := result.(*LookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for look: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RunLook runs look

  ### Run a Look

Runs a given look's query and returns the results in the requested format.

Supported formats:

| result_format | Description
| :-----------: | :--- |
| json | Plain json
| json_detail | Row data plus metadata describing the fields, pivots, table calcs, and other aspects of the query
| csv | Comma separated values with a header
| txt | Tab separated values with a header
| html | Simple html
| md | Simple markdown
| xlsx | MS Excel spreadsheet
| sql | Returns the generated SQL rather than running the query
| png | A PNG image of the visualization of the query
| jpg | A JPG image of the visualization of the query



*/
func (a *Client) RunLook(params *RunLookParams, opts ...ClientOption) (*RunLookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunLookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "run_look",
		Method:             "GET",
		PathPattern:        "/looks/{look_id}/run/{result_format}",
		ProducesMediaTypes: []string{"application/json", "image/jpeg", "image/png", "text"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RunLookReader{formats: a.formats},
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
	success, ok := result.(*RunLookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for run_look: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchLooks searches looks

  ### Search Looks

Returns an **array of Look objects** that match the specified search criteria.

If multiple search params are given and `filter_or` is FALSE or not specified,
search params are combined in a logical AND operation.
Only rows that match *all* search param criteria will be returned.

If `filter_or` is TRUE, multiple search params are combined in a logical OR operation.
Results will include rows that match **any** of the search criteria.

String search params use case-insensitive matching.
String search params can contain `%` and '_' as SQL LIKE pattern match wildcard expressions.
example="dan%" will match "danger" and "Danzig" but not "David"
example="D_m%" will match "Damage" and "dump"

Integer search params can accept a single value or a comma separated list of values. The multiple
values will be combined under a logical OR operation - results will match at least one of
the given values.

Most search params can accept "IS NULL" and "NOT NULL" as special expressions to match
or exclude (respectively) rows where the column is null.

Boolean search params accept only "true" and "false" as values.


Get a **single look** by id with [look(id)](#!/Look/look)

*/
func (a *Client) SearchLooks(params *SearchLooksParams, opts ...ClientOption) (*SearchLooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchLooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "search_looks",
		Method:             "GET",
		PathPattern:        "/looks/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchLooksReader{formats: a.formats},
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
	success, ok := result.(*SearchLooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for search_looks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLook updates look

  ### Modify a Look

Use this function to modify parts of a look. Property values given in a call to `update_look` are
applied to the existing look, so there's no need to include properties whose values are not changing.
It's best to specify only the properties you want to change and leave everything else out
of your `update_look` call. **Look properties marked 'read-only' will be ignored.**

When a user deletes a look in the Looker UI, the look data remains in the database but is
marked with a deleted flag ("soft-deleted"). Soft-deleted looks can be undeleted (by an admin)
if the delete was in error.

To soft-delete a look via the API, use [update_look()](#!/Look/update_look) to change the look's `deleted` property to `true`.
You can undelete a look by calling `update_look` to change the look's `deleted` property to `false`.

Soft-deleted looks are excluded from the results of [all_looks()](#!/Look/all_looks) and [search_looks()](#!/Look/search_looks), so they
essentially disappear from view even though they still reside in the db.
In API 3.1 and later, you can pass `deleted: true` as a parameter to [search_looks()](#!/3.1/Look/search_looks) to list soft-deleted looks.

NOTE: [delete_look()](#!/Look/delete_look) performs a "hard delete" - the look data is removed from the Looker
database and destroyed. There is no "undo" for `delete_look()`.

*/
func (a *Client) UpdateLook(params *UpdateLookParams, opts ...ClientOption) (*UpdateLookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_look",
		Method:             "PATCH",
		PathPattern:        "/looks/{look_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLookReader{formats: a.formats},
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
	success, ok := result.(*UpdateLookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update_look: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
