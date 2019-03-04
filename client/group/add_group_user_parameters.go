// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// NewAddGroupUserParams creates a new AddGroupUserParams object
// with the default values initialized.
func NewAddGroupUserParams() *AddGroupUserParams {
	var ()
	return &AddGroupUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddGroupUserParamsWithTimeout creates a new AddGroupUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddGroupUserParamsWithTimeout(timeout time.Duration) *AddGroupUserParams {
	var ()
	return &AddGroupUserParams{

		timeout: timeout,
	}
}

// NewAddGroupUserParamsWithContext creates a new AddGroupUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddGroupUserParamsWithContext(ctx context.Context) *AddGroupUserParams {
	var ()
	return &AddGroupUserParams{

		Context: ctx,
	}
}

// NewAddGroupUserParamsWithHTTPClient creates a new AddGroupUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddGroupUserParamsWithHTTPClient(client *http.Client) *AddGroupUserParams {
	var ()
	return &AddGroupUserParams{
		HTTPClient: client,
	}
}

/*AddGroupUserParams contains all the parameters to send to the API endpoint
for the add group user operation typically these are written to a http.Request
*/
type AddGroupUserParams struct {

	/*Body
	  User id to add

	*/
	Body *models.GroupIDForGroupUserInclusion
	/*GroupID
	  Id of group

	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add group user params
func (o *AddGroupUserParams) WithTimeout(timeout time.Duration) *AddGroupUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add group user params
func (o *AddGroupUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add group user params
func (o *AddGroupUserParams) WithContext(ctx context.Context) *AddGroupUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add group user params
func (o *AddGroupUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add group user params
func (o *AddGroupUserParams) WithHTTPClient(client *http.Client) *AddGroupUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add group user params
func (o *AddGroupUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add group user params
func (o *AddGroupUserParams) WithBody(body *models.GroupIDForGroupUserInclusion) *AddGroupUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add group user params
func (o *AddGroupUserParams) SetBody(body *models.GroupIDForGroupUserInclusion) {
	o.Body = body
}

// WithGroupID adds the groupID to the add group user params
func (o *AddGroupUserParams) WithGroupID(groupID int64) *AddGroupUserParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the add group user params
func (o *AddGroupUserParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *AddGroupUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
