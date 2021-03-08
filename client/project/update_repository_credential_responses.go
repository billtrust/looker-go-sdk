// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateRepositoryCredentialReader is a Reader for the UpdateRepositoryCredential structure.
type UpdateRepositoryCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRepositoryCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepositoryCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepositoryCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRepositoryCredentialConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateRepositoryCredentialUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRepositoryCredentialTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryCredentialOK creates a UpdateRepositoryCredentialOK with default headers values
func NewUpdateRepositoryCredentialOK() *UpdateRepositoryCredentialOK {
	return &UpdateRepositoryCredentialOK{}
}

/* UpdateRepositoryCredentialOK describes a response with status code 200, with default header values.

Repository Credential
*/
type UpdateRepositoryCredentialOK struct {
	Payload *models.RepositoryCredential
}

func (o *UpdateRepositoryCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{root_project_id}/credential/{credential_id}][%d] updateRepositoryCredentialOK  %+v", 200, o.Payload)
}
func (o *UpdateRepositoryCredentialOK) GetPayload() *models.RepositoryCredential {
	return o.Payload
}

func (o *UpdateRepositoryCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepositoryCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryCredentialBadRequest creates a UpdateRepositoryCredentialBadRequest with default headers values
func NewUpdateRepositoryCredentialBadRequest() *UpdateRepositoryCredentialBadRequest {
	return &UpdateRepositoryCredentialBadRequest{}
}

/* UpdateRepositoryCredentialBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateRepositoryCredentialBadRequest struct {
	Payload *models.Error
}

func (o *UpdateRepositoryCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{root_project_id}/credential/{credential_id}][%d] updateRepositoryCredentialBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRepositoryCredentialBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryCredentialNotFound creates a UpdateRepositoryCredentialNotFound with default headers values
func NewUpdateRepositoryCredentialNotFound() *UpdateRepositoryCredentialNotFound {
	return &UpdateRepositoryCredentialNotFound{}
}

/* UpdateRepositoryCredentialNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRepositoryCredentialNotFound struct {
	Payload *models.Error
}

func (o *UpdateRepositoryCredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{root_project_id}/credential/{credential_id}][%d] updateRepositoryCredentialNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRepositoryCredentialNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryCredentialConflict creates a UpdateRepositoryCredentialConflict with default headers values
func NewUpdateRepositoryCredentialConflict() *UpdateRepositoryCredentialConflict {
	return &UpdateRepositoryCredentialConflict{}
}

/* UpdateRepositoryCredentialConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type UpdateRepositoryCredentialConflict struct {
	Payload *models.Error
}

func (o *UpdateRepositoryCredentialConflict) Error() string {
	return fmt.Sprintf("[PUT /projects/{root_project_id}/credential/{credential_id}][%d] updateRepositoryCredentialConflict  %+v", 409, o.Payload)
}
func (o *UpdateRepositoryCredentialConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryCredentialConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryCredentialUnprocessableEntity creates a UpdateRepositoryCredentialUnprocessableEntity with default headers values
func NewUpdateRepositoryCredentialUnprocessableEntity() *UpdateRepositoryCredentialUnprocessableEntity {
	return &UpdateRepositoryCredentialUnprocessableEntity{}
}

/* UpdateRepositoryCredentialUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateRepositoryCredentialUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateRepositoryCredentialUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /projects/{root_project_id}/credential/{credential_id}][%d] updateRepositoryCredentialUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateRepositoryCredentialUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateRepositoryCredentialUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRepositoryCredentialTooManyRequests creates a UpdateRepositoryCredentialTooManyRequests with default headers values
func NewUpdateRepositoryCredentialTooManyRequests() *UpdateRepositoryCredentialTooManyRequests {
	return &UpdateRepositoryCredentialTooManyRequests{}
}

/* UpdateRepositoryCredentialTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRepositoryCredentialTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateRepositoryCredentialTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /projects/{root_project_id}/credential/{credential_id}][%d] updateRepositoryCredentialTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateRepositoryCredentialTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateRepositoryCredentialTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
