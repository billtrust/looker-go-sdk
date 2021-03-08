// Code generated by go-swagger; DO NOT EDIT.

package theme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateThemeReader is a Reader for the UpdateTheme structure.
type UpdateThemeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateThemeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateThemeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateThemeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateThemeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateThemeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateThemeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateThemeOK creates a UpdateThemeOK with default headers values
func NewUpdateThemeOK() *UpdateThemeOK {
	return &UpdateThemeOK{}
}

/* UpdateThemeOK describes a response with status code 200, with default header values.

Theme
*/
type UpdateThemeOK struct {
	Payload *models.Theme
}

func (o *UpdateThemeOK) Error() string {
	return fmt.Sprintf("[PATCH /themes/{theme_id}][%d] updateThemeOK  %+v", 200, o.Payload)
}
func (o *UpdateThemeOK) GetPayload() *models.Theme {
	return o.Payload
}

func (o *UpdateThemeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Theme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateThemeBadRequest creates a UpdateThemeBadRequest with default headers values
func NewUpdateThemeBadRequest() *UpdateThemeBadRequest {
	return &UpdateThemeBadRequest{}
}

/* UpdateThemeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateThemeBadRequest struct {
	Payload *models.Error
}

func (o *UpdateThemeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /themes/{theme_id}][%d] updateThemeBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateThemeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateThemeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateThemeNotFound creates a UpdateThemeNotFound with default headers values
func NewUpdateThemeNotFound() *UpdateThemeNotFound {
	return &UpdateThemeNotFound{}
}

/* UpdateThemeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateThemeNotFound struct {
	Payload *models.Error
}

func (o *UpdateThemeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /themes/{theme_id}][%d] updateThemeNotFound  %+v", 404, o.Payload)
}
func (o *UpdateThemeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateThemeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateThemeUnprocessableEntity creates a UpdateThemeUnprocessableEntity with default headers values
func NewUpdateThemeUnprocessableEntity() *UpdateThemeUnprocessableEntity {
	return &UpdateThemeUnprocessableEntity{}
}

/* UpdateThemeUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateThemeUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateThemeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /themes/{theme_id}][%d] updateThemeUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateThemeUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateThemeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateThemeTooManyRequests creates a UpdateThemeTooManyRequests with default headers values
func NewUpdateThemeTooManyRequests() *UpdateThemeTooManyRequests {
	return &UpdateThemeTooManyRequests{}
}

/* UpdateThemeTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateThemeTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateThemeTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /themes/{theme_id}][%d] updateThemeTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateThemeTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateThemeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
