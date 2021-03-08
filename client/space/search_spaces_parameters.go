// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSearchSpacesParams creates a new SearchSpacesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchSpacesParams() *SearchSpacesParams {
	return &SearchSpacesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchSpacesParamsWithTimeout creates a new SearchSpacesParams object
// with the ability to set a timeout on a request.
func NewSearchSpacesParamsWithTimeout(timeout time.Duration) *SearchSpacesParams {
	return &SearchSpacesParams{
		timeout: timeout,
	}
}

// NewSearchSpacesParamsWithContext creates a new SearchSpacesParams object
// with the ability to set a context for a request.
func NewSearchSpacesParamsWithContext(ctx context.Context) *SearchSpacesParams {
	return &SearchSpacesParams{
		Context: ctx,
	}
}

// NewSearchSpacesParamsWithHTTPClient creates a new SearchSpacesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchSpacesParamsWithHTTPClient(client *http.Client) *SearchSpacesParams {
	return &SearchSpacesParams{
		HTTPClient: client,
	}
}

/* SearchSpacesParams contains all the parameters to send to the API endpoint
   for the search spaces operation.

   Typically these are written to a http.Request.
*/
type SearchSpacesParams struct {

	/* CreatorID.

	   Filter on spaces created by a particular user.
	*/
	CreatorID *string

	/* Fields.

	   Requested fields.
	*/
	Fields *string

	/* FilterOr.

	   Combine given search criteria in a boolean OR expression
	*/
	FilterOr *bool

	/* ID.

	   Match Space id

	   Format: int64
	*/
	ID *int64

	/* Limit.

	   Number of results to return. (used with offset and takes priority over page and per_page)

	   Format: int64
	*/
	Limit *int64

	/* Name.

	   Match Space title.
	*/
	Name *string

	/* Offset.

	   Number of results to skip before returning any. (used with limit and takes priority over page and per_page)

	   Format: int64
	*/
	Offset *int64

	/* Page.

	   Requested page.

	   Format: int64
	*/
	Page *int64

	/* ParentID.

	   Filter on a children of a particular space.
	*/
	ParentID *string

	/* PerPage.

	   Results per page.

	   Format: int64
	*/
	PerPage *int64

	/* Sorts.

	   Fields to sort by.
	*/
	Sorts *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search spaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpacesParams) WithDefaults() *SearchSpacesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search spaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpacesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search spaces params
func (o *SearchSpacesParams) WithTimeout(timeout time.Duration) *SearchSpacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search spaces params
func (o *SearchSpacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search spaces params
func (o *SearchSpacesParams) WithContext(ctx context.Context) *SearchSpacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search spaces params
func (o *SearchSpacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search spaces params
func (o *SearchSpacesParams) WithHTTPClient(client *http.Client) *SearchSpacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search spaces params
func (o *SearchSpacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatorID adds the creatorID to the search spaces params
func (o *SearchSpacesParams) WithCreatorID(creatorID *string) *SearchSpacesParams {
	o.SetCreatorID(creatorID)
	return o
}

// SetCreatorID adds the creatorId to the search spaces params
func (o *SearchSpacesParams) SetCreatorID(creatorID *string) {
	o.CreatorID = creatorID
}

// WithFields adds the fields to the search spaces params
func (o *SearchSpacesParams) WithFields(fields *string) *SearchSpacesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search spaces params
func (o *SearchSpacesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilterOr adds the filterOr to the search spaces params
func (o *SearchSpacesParams) WithFilterOr(filterOr *bool) *SearchSpacesParams {
	o.SetFilterOr(filterOr)
	return o
}

// SetFilterOr adds the filterOr to the search spaces params
func (o *SearchSpacesParams) SetFilterOr(filterOr *bool) {
	o.FilterOr = filterOr
}

// WithID adds the id to the search spaces params
func (o *SearchSpacesParams) WithID(id *int64) *SearchSpacesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search spaces params
func (o *SearchSpacesParams) SetID(id *int64) {
	o.ID = id
}

// WithLimit adds the limit to the search spaces params
func (o *SearchSpacesParams) WithLimit(limit *int64) *SearchSpacesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search spaces params
func (o *SearchSpacesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the search spaces params
func (o *SearchSpacesParams) WithName(name *string) *SearchSpacesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the search spaces params
func (o *SearchSpacesParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the search spaces params
func (o *SearchSpacesParams) WithOffset(offset *int64) *SearchSpacesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search spaces params
func (o *SearchSpacesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the search spaces params
func (o *SearchSpacesParams) WithPage(page *int64) *SearchSpacesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search spaces params
func (o *SearchSpacesParams) SetPage(page *int64) {
	o.Page = page
}

// WithParentID adds the parentID to the search spaces params
func (o *SearchSpacesParams) WithParentID(parentID *string) *SearchSpacesParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the search spaces params
func (o *SearchSpacesParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WithPerPage adds the perPage to the search spaces params
func (o *SearchSpacesParams) WithPerPage(perPage *int64) *SearchSpacesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the search spaces params
func (o *SearchSpacesParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSorts adds the sorts to the search spaces params
func (o *SearchSpacesParams) WithSorts(sorts *string) *SearchSpacesParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the search spaces params
func (o *SearchSpacesParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WriteToRequest writes these params to a swagger request
func (o *SearchSpacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatorID != nil {

		// query param creator_id
		var qrCreatorID string

		if o.CreatorID != nil {
			qrCreatorID = *o.CreatorID
		}
		qCreatorID := qrCreatorID
		if qCreatorID != "" {

			if err := r.SetQueryParam("creator_id", qCreatorID); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if o.FilterOr != nil {

		// query param filter_or
		var qrFilterOr bool

		if o.FilterOr != nil {
			qrFilterOr = *o.FilterOr
		}
		qFilterOr := swag.FormatBool(qrFilterOr)
		if qFilterOr != "" {

			if err := r.SetQueryParam("filter_or", qFilterOr); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID int64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.ParentID != nil {

		// query param parent_id
		var qrParentID string

		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {

			if err := r.SetQueryParam("parent_id", qParentID); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string

		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {

			if err := r.SetQueryParam("sorts", qSorts); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
