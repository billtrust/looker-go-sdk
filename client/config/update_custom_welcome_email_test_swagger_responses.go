// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// UpdateCustomWelcomeEmailTestReader is a Reader for the UpdateCustomWelcomeEmailTest structure.
type UpdateCustomWelcomeEmailTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomWelcomeEmailTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomWelcomeEmailTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomWelcomeEmailTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCustomWelcomeEmailTestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateCustomWelcomeEmailTestUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCustomWelcomeEmailTestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomWelcomeEmailTestOK creates a UpdateCustomWelcomeEmailTestOK with default headers values
func NewUpdateCustomWelcomeEmailTestOK() *UpdateCustomWelcomeEmailTestOK {
	return &UpdateCustomWelcomeEmailTestOK{}
}

/* UpdateCustomWelcomeEmailTestOK describes a response with status code 200, with default header values.

Send Test Welcome Email
*/
type UpdateCustomWelcomeEmailTestOK struct {
	Payload *models.WelcomeEmailTest
}

func (o *UpdateCustomWelcomeEmailTestOK) Error() string {
	return fmt.Sprintf("[PUT /custom_welcome_email_test][%d] updateCustomWelcomeEmailTestOK  %+v", 200, o.Payload)
}
func (o *UpdateCustomWelcomeEmailTestOK) GetPayload() *models.WelcomeEmailTest {
	return o.Payload
}

func (o *UpdateCustomWelcomeEmailTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WelcomeEmailTest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomWelcomeEmailTestBadRequest creates a UpdateCustomWelcomeEmailTestBadRequest with default headers values
func NewUpdateCustomWelcomeEmailTestBadRequest() *UpdateCustomWelcomeEmailTestBadRequest {
	return &UpdateCustomWelcomeEmailTestBadRequest{}
}

/* UpdateCustomWelcomeEmailTestBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCustomWelcomeEmailTestBadRequest struct {
	Payload *models.Error
}

func (o *UpdateCustomWelcomeEmailTestBadRequest) Error() string {
	return fmt.Sprintf("[PUT /custom_welcome_email_test][%d] updateCustomWelcomeEmailTestBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCustomWelcomeEmailTestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomWelcomeEmailTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomWelcomeEmailTestNotFound creates a UpdateCustomWelcomeEmailTestNotFound with default headers values
func NewUpdateCustomWelcomeEmailTestNotFound() *UpdateCustomWelcomeEmailTestNotFound {
	return &UpdateCustomWelcomeEmailTestNotFound{}
}

/* UpdateCustomWelcomeEmailTestNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateCustomWelcomeEmailTestNotFound struct {
	Payload *models.Error
}

func (o *UpdateCustomWelcomeEmailTestNotFound) Error() string {
	return fmt.Sprintf("[PUT /custom_welcome_email_test][%d] updateCustomWelcomeEmailTestNotFound  %+v", 404, o.Payload)
}
func (o *UpdateCustomWelcomeEmailTestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomWelcomeEmailTestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomWelcomeEmailTestUnprocessableEntity creates a UpdateCustomWelcomeEmailTestUnprocessableEntity with default headers values
func NewUpdateCustomWelcomeEmailTestUnprocessableEntity() *UpdateCustomWelcomeEmailTestUnprocessableEntity {
	return &UpdateCustomWelcomeEmailTestUnprocessableEntity{}
}

/* UpdateCustomWelcomeEmailTestUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type UpdateCustomWelcomeEmailTestUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateCustomWelcomeEmailTestUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /custom_welcome_email_test][%d] updateCustomWelcomeEmailTestUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateCustomWelcomeEmailTestUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateCustomWelcomeEmailTestUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomWelcomeEmailTestTooManyRequests creates a UpdateCustomWelcomeEmailTestTooManyRequests with default headers values
func NewUpdateCustomWelcomeEmailTestTooManyRequests() *UpdateCustomWelcomeEmailTestTooManyRequests {
	return &UpdateCustomWelcomeEmailTestTooManyRequests{}
}

/* UpdateCustomWelcomeEmailTestTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCustomWelcomeEmailTestTooManyRequests struct {
	Payload *models.Error
}

func (o *UpdateCustomWelcomeEmailTestTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /custom_welcome_email_test][%d] updateCustomWelcomeEmailTestTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateCustomWelcomeEmailTestTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomWelcomeEmailTestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
