// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateFolderReader is a Reader for the UpdateFolder structure.
type UpdateFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateFolderUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateFolderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFolderOK creates a UpdateFolderOK with default headers values
func NewUpdateFolderOK() *UpdateFolderOK {
	return &UpdateFolderOK{}
}

/* UpdateFolderOK describes a response with status code 200, with default header values.

Folder
*/
type UpdateFolderOK struct {
	Payload *models.Folder
}

func (o *UpdateFolderOK) Error() string {
	return fmt.Sprintf("[PATCH /folders/{folder_id}][%d] updateFolderOK  %+v", 200, o.Payload)
}
func (o *UpdateFolderOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *UpdateFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderBadRequest creates a UpdateFolderBadRequest with default headers values
func NewUpdateFolderBadRequest() *UpdateFolderBadRequest {
	return &UpdateFolderBadRequest{}
}

/* UpdateFolderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateFolderBadRequest struct {
	Payload *models.Error
}

func (o *UpdateFolderBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /folders/{folder_id}][%d] updateFolderBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateFolderBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderNotFound creates a UpdateFolderNotFound with default headers values
func NewUpdateFolderNotFound() *UpdateFolderNotFound {
	return &UpdateFolderNotFound{}
}

/* UpdateFolderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateFolderNotFound struct {
	Payload *models.Error
}

func (o *UpdateFolderNotFound) Error() string {
	return fmt.Sprintf("[PATCH /folders/{folder_id}][%d] updateFolderNotFound  %+v", 404, o.Payload)
}
func (o *UpdateFolderNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderUnprocessableEntity creates a UpdateFolderUnprocessableEntity with default headers values
func NewUpdateFolderUnprocessableEntity() *UpdateFolderUnprocessableEntity {
	return &UpdateFolderUnprocessableEntity{}
}

/* UpdateFolderUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateFolderUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateFolderUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /folders/{folder_id}][%d] updateFolderUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateFolderUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateFolderUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFolderTooManyRequests creates a UpdateFolderTooManyRequests with default headers values
func NewUpdateFolderTooManyRequests() *UpdateFolderTooManyRequests {
	return &UpdateFolderTooManyRequests{}
}

/* UpdateFolderTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateFolderTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateFolderTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /folders/{folder_id}][%d] updateFolderTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateFolderTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateFolderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
