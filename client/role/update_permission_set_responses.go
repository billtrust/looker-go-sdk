// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdatePermissionSetReader is a Reader for the UpdatePermissionSet structure.
type UpdatePermissionSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePermissionSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePermissionSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePermissionSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePermissionSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdatePermissionSetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdatePermissionSetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdatePermissionSetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePermissionSetOK creates a UpdatePermissionSetOK with default headers values
func NewUpdatePermissionSetOK() *UpdatePermissionSetOK {
	return &UpdatePermissionSetOK{}
}

/* UpdatePermissionSetOK describes a response with status code 200, with default header values.

Permission Set
*/
type UpdatePermissionSetOK struct {
	Payload *models.PermissionSet
}

func (o *UpdatePermissionSetOK) Error() string {
	return fmt.Sprintf("[PATCH /permission_sets/{permission_set_id}][%d] updatePermissionSetOK  %+v", 200, o.Payload)
}
func (o *UpdatePermissionSetOK) GetPayload() *models.PermissionSet {
	return o.Payload
}

func (o *UpdatePermissionSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePermissionSetBadRequest creates a UpdatePermissionSetBadRequest with default headers values
func NewUpdatePermissionSetBadRequest() *UpdatePermissionSetBadRequest {
	return &UpdatePermissionSetBadRequest{}
}

/* UpdatePermissionSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdatePermissionSetBadRequest struct {
	Payload *models.Error
}

func (o *UpdatePermissionSetBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /permission_sets/{permission_set_id}][%d] updatePermissionSetBadRequest  %+v", 400, o.Payload)
}
func (o *UpdatePermissionSetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePermissionSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePermissionSetNotFound creates a UpdatePermissionSetNotFound with default headers values
func NewUpdatePermissionSetNotFound() *UpdatePermissionSetNotFound {
	return &UpdatePermissionSetNotFound{}
}

/* UpdatePermissionSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdatePermissionSetNotFound struct {
	Payload *models.Error
}

func (o *UpdatePermissionSetNotFound) Error() string {
	return fmt.Sprintf("[PATCH /permission_sets/{permission_set_id}][%d] updatePermissionSetNotFound  %+v", 404, o.Payload)
}
func (o *UpdatePermissionSetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePermissionSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePermissionSetMethodNotAllowed creates a UpdatePermissionSetMethodNotAllowed with default headers values
func NewUpdatePermissionSetMethodNotAllowed() *UpdatePermissionSetMethodNotAllowed {
	return &UpdatePermissionSetMethodNotAllowed{}
}

/* UpdatePermissionSetMethodNotAllowed describes a response with status code 405, with default header values.

Resource Can't Be Modified
*/
type UpdatePermissionSetMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdatePermissionSetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /permission_sets/{permission_set_id}][%d] updatePermissionSetMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *UpdatePermissionSetMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePermissionSetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePermissionSetUnprocessableEntity creates a UpdatePermissionSetUnprocessableEntity with default headers values
func NewUpdatePermissionSetUnprocessableEntity() *UpdatePermissionSetUnprocessableEntity {
	return &UpdatePermissionSetUnprocessableEntity{}
}

/* UpdatePermissionSetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdatePermissionSetUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdatePermissionSetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /permission_sets/{permission_set_id}][%d] updatePermissionSetUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdatePermissionSetUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdatePermissionSetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePermissionSetTooManyRequests creates a UpdatePermissionSetTooManyRequests with default headers values
func NewUpdatePermissionSetTooManyRequests() *UpdatePermissionSetTooManyRequests {
	return &UpdatePermissionSetTooManyRequests{}
}

/* UpdatePermissionSetTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdatePermissionSetTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdatePermissionSetTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /permission_sets/{permission_set_id}][%d] updatePermissionSetTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdatePermissionSetTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePermissionSetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
