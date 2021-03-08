// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// CreateUserCredentialsApi3Reader is a Reader for the CreateUserCredentialsApi3 structure.
type CreateUserCredentialsApi3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialsApi3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserCredentialsApi3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserCredentialsApi3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserCredentialsApi3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserCredentialsApi3Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateUserCredentialsApi3UnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateUserCredentialsApi3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserCredentialsApi3OK creates a CreateUserCredentialsApi3OK with default headers values
func NewCreateUserCredentialsApi3OK() *CreateUserCredentialsApi3OK {
	return &CreateUserCredentialsApi3OK{}
}

/* CreateUserCredentialsApi3OK describes a response with status code 200, with default header values.

API 3 Credential
*/
type CreateUserCredentialsApi3OK struct {
	Payload *models.CredentialsApi3
}

func (o *CreateUserCredentialsApi3OK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api3][%d] createUserCredentialsApi3OK  %+v", 200, o.Payload)
}
func (o *CreateUserCredentialsApi3OK) GetPayload() *models.CredentialsApi3 {
	return o.Payload
}

func (o *CreateUserCredentialsApi3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsApi3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsApi3BadRequest creates a CreateUserCredentialsApi3BadRequest with default headers values
func NewCreateUserCredentialsApi3BadRequest() *CreateUserCredentialsApi3BadRequest {
	return &CreateUserCredentialsApi3BadRequest{}
}

/* CreateUserCredentialsApi3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateUserCredentialsApi3BadRequest struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsApi3BadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api3][%d] createUserCredentialsApi3BadRequest  %+v", 400, o.Payload)
}
func (o *CreateUserCredentialsApi3BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsApi3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsApi3NotFound creates a CreateUserCredentialsApi3NotFound with default headers values
func NewCreateUserCredentialsApi3NotFound() *CreateUserCredentialsApi3NotFound {
	return &CreateUserCredentialsApi3NotFound{}
}

/* CreateUserCredentialsApi3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateUserCredentialsApi3NotFound struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsApi3NotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api3][%d] createUserCredentialsApi3NotFound  %+v", 404, o.Payload)
}
func (o *CreateUserCredentialsApi3NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsApi3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsApi3Conflict creates a CreateUserCredentialsApi3Conflict with default headers values
func NewCreateUserCredentialsApi3Conflict() *CreateUserCredentialsApi3Conflict {
	return &CreateUserCredentialsApi3Conflict{}
}

/* CreateUserCredentialsApi3Conflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateUserCredentialsApi3Conflict struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsApi3Conflict) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api3][%d] createUserCredentialsApi3Conflict  %+v", 409, o.Payload)
}
func (o *CreateUserCredentialsApi3Conflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsApi3Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsApi3UnprocessableEntity creates a CreateUserCredentialsApi3UnprocessableEntity with default headers values
func NewCreateUserCredentialsApi3UnprocessableEntity() *CreateUserCredentialsApi3UnprocessableEntity {
	return &CreateUserCredentialsApi3UnprocessableEntity{}
}

/* CreateUserCredentialsApi3UnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateUserCredentialsApi3UnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateUserCredentialsApi3UnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api3][%d] createUserCredentialsApi3UnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateUserCredentialsApi3UnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateUserCredentialsApi3UnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsApi3TooManyRequests creates a CreateUserCredentialsApi3TooManyRequests with default headers values
func NewCreateUserCredentialsApi3TooManyRequests() *CreateUserCredentialsApi3TooManyRequests {
	return &CreateUserCredentialsApi3TooManyRequests{}
}

/* CreateUserCredentialsApi3TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateUserCredentialsApi3TooManyRequests struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsApi3TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_api3][%d] createUserCredentialsApi3TooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateUserCredentialsApi3TooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsApi3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
