// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateConnectionReader is a Reader for the UpdateConnection structure.
type UpdateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateConnectionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateConnectionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConnectionOK creates a UpdateConnectionOK with default headers values
func NewUpdateConnectionOK() *UpdateConnectionOK {
	return &UpdateConnectionOK{}
}

/* UpdateConnectionOK describes a response with status code 200, with default header values.

Connection
*/
type UpdateConnectionOK struct {
	Payload *models.DBConnection
}

func (o *UpdateConnectionOK) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionOK  %+v", 200, o.Payload)
}
func (o *UpdateConnectionOK) GetPayload() *models.DBConnection {
	return o.Payload
}

func (o *UpdateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DBConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionBadRequest creates a UpdateConnectionBadRequest with default headers values
func NewUpdateConnectionBadRequest() *UpdateConnectionBadRequest {
	return &UpdateConnectionBadRequest{}
}

/* UpdateConnectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateConnectionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateConnectionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateConnectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionNotFound creates a UpdateConnectionNotFound with default headers values
func NewUpdateConnectionNotFound() *UpdateConnectionNotFound {
	return &UpdateConnectionNotFound{}
}

/* UpdateConnectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateConnectionNotFound struct {
	Payload *models.Error
}

func (o *UpdateConnectionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionNotFound  %+v", 404, o.Payload)
}
func (o *UpdateConnectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionUnprocessableEntity creates a UpdateConnectionUnprocessableEntity with default headers values
func NewUpdateConnectionUnprocessableEntity() *UpdateConnectionUnprocessableEntity {
	return &UpdateConnectionUnprocessableEntity{}
}

/* UpdateConnectionUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateConnectionUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateConnectionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateConnectionUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateConnectionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionTooManyRequests creates a UpdateConnectionTooManyRequests with default headers values
func NewUpdateConnectionTooManyRequests() *UpdateConnectionTooManyRequests {
	return &UpdateConnectionTooManyRequests{}
}

/* UpdateConnectionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateConnectionTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateConnectionTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /connections/{connection_name}][%d] updateConnectionTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateConnectionTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateConnectionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
