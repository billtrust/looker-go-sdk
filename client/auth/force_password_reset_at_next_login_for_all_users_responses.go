// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// ForcePasswordResetAtNextLoginForAllUsersReader is a Reader for the ForcePasswordResetAtNextLoginForAllUsers structure.
type ForcePasswordResetAtNextLoginForAllUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForcePasswordResetAtNextLoginForAllUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewForcePasswordResetAtNextLoginForAllUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewForcePasswordResetAtNextLoginForAllUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewForcePasswordResetAtNextLoginForAllUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewForcePasswordResetAtNextLoginForAllUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewForcePasswordResetAtNextLoginForAllUsersOK creates a ForcePasswordResetAtNextLoginForAllUsersOK with default headers values
func NewForcePasswordResetAtNextLoginForAllUsersOK() *ForcePasswordResetAtNextLoginForAllUsersOK {
	return &ForcePasswordResetAtNextLoginForAllUsersOK{}
}

/* ForcePasswordResetAtNextLoginForAllUsersOK describes a response with status code 200, with default header values.

Password Config
*/
type ForcePasswordResetAtNextLoginForAllUsersOK struct {
	Payload string
}

func (o *ForcePasswordResetAtNextLoginForAllUsersOK) Error() string {
	return fmt.Sprintf("[PUT /password_config/force_password_reset_at_next_login_for_all_users][%d] forcePasswordResetAtNextLoginForAllUsersOK  %+v", 200, o.Payload)
}
func (o *ForcePasswordResetAtNextLoginForAllUsersOK) GetPayload() string {
	return o.Payload
}

func (o *ForcePasswordResetAtNextLoginForAllUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForcePasswordResetAtNextLoginForAllUsersBadRequest creates a ForcePasswordResetAtNextLoginForAllUsersBadRequest with default headers values
func NewForcePasswordResetAtNextLoginForAllUsersBadRequest() *ForcePasswordResetAtNextLoginForAllUsersBadRequest {
	return &ForcePasswordResetAtNextLoginForAllUsersBadRequest{}
}

/* ForcePasswordResetAtNextLoginForAllUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ForcePasswordResetAtNextLoginForAllUsersBadRequest struct {
	Payload *models.Error
}

func (o *ForcePasswordResetAtNextLoginForAllUsersBadRequest) Error() string {
	return fmt.Sprintf("[PUT /password_config/force_password_reset_at_next_login_for_all_users][%d] forcePasswordResetAtNextLoginForAllUsersBadRequest  %+v", 400, o.Payload)
}
func (o *ForcePasswordResetAtNextLoginForAllUsersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ForcePasswordResetAtNextLoginForAllUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForcePasswordResetAtNextLoginForAllUsersNotFound creates a ForcePasswordResetAtNextLoginForAllUsersNotFound with default headers values
func NewForcePasswordResetAtNextLoginForAllUsersNotFound() *ForcePasswordResetAtNextLoginForAllUsersNotFound {
	return &ForcePasswordResetAtNextLoginForAllUsersNotFound{}
}

/* ForcePasswordResetAtNextLoginForAllUsersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ForcePasswordResetAtNextLoginForAllUsersNotFound struct {
	Payload *models.Error
}

func (o *ForcePasswordResetAtNextLoginForAllUsersNotFound) Error() string {
	return fmt.Sprintf("[PUT /password_config/force_password_reset_at_next_login_for_all_users][%d] forcePasswordResetAtNextLoginForAllUsersNotFound  %+v", 404, o.Payload)
}
func (o *ForcePasswordResetAtNextLoginForAllUsersNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ForcePasswordResetAtNextLoginForAllUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity creates a ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity with default headers values
func NewForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity() *ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity {
	return &ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity{}
}

/* ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /password_config/force_password_reset_at_next_login_for_all_users][%d] forcePasswordResetAtNextLoginForAllUsersUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *ForcePasswordResetAtNextLoginForAllUsersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForcePasswordResetAtNextLoginForAllUsersTooManyRequests creates a ForcePasswordResetAtNextLoginForAllUsersTooManyRequests with default headers values
func NewForcePasswordResetAtNextLoginForAllUsersTooManyRequests() *ForcePasswordResetAtNextLoginForAllUsersTooManyRequests {
	return &ForcePasswordResetAtNextLoginForAllUsersTooManyRequests{}
}

/* ForcePasswordResetAtNextLoginForAllUsersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ForcePasswordResetAtNextLoginForAllUsersTooManyRequests struct {
	Payload *models.Error
}

func (o *ForcePasswordResetAtNextLoginForAllUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /password_config/force_password_reset_at_next_login_for_all_users][%d] forcePasswordResetAtNextLoginForAllUsersTooManyRequests  %+v", 429, o.Payload)
}
func (o *ForcePasswordResetAtNextLoginForAllUsersTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ForcePasswordResetAtNextLoginForAllUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
