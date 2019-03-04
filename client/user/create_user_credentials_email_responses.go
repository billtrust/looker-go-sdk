// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// CreateUserCredentialsEmailReader is a Reader for the CreateUserCredentialsEmail structure.
type CreateUserCredentialsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserCredentialsEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUserCredentialsEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateUserCredentialsEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateUserCredentialsEmailConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCreateUserCredentialsEmailUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCredentialsEmailOK creates a CreateUserCredentialsEmailOK with default headers values
func NewCreateUserCredentialsEmailOK() *CreateUserCredentialsEmailOK {
	return &CreateUserCredentialsEmailOK{}
}

/*CreateUserCredentialsEmailOK handles this case with default header values.

Email/Password Credential
*/
type CreateUserCredentialsEmailOK struct {
	Payload *models.CredentialsEmail
}

func (o *CreateUserCredentialsEmailOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email][%d] createUserCredentialsEmailOK  %+v", 200, o.Payload)
}

func (o *CreateUserCredentialsEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsEmail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsEmailBadRequest creates a CreateUserCredentialsEmailBadRequest with default headers values
func NewCreateUserCredentialsEmailBadRequest() *CreateUserCredentialsEmailBadRequest {
	return &CreateUserCredentialsEmailBadRequest{}
}

/*CreateUserCredentialsEmailBadRequest handles this case with default header values.

Bad Request
*/
type CreateUserCredentialsEmailBadRequest struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email][%d] createUserCredentialsEmailBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserCredentialsEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsEmailNotFound creates a CreateUserCredentialsEmailNotFound with default headers values
func NewCreateUserCredentialsEmailNotFound() *CreateUserCredentialsEmailNotFound {
	return &CreateUserCredentialsEmailNotFound{}
}

/*CreateUserCredentialsEmailNotFound handles this case with default header values.

Not Found
*/
type CreateUserCredentialsEmailNotFound struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email][%d] createUserCredentialsEmailNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserCredentialsEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsEmailConflict creates a CreateUserCredentialsEmailConflict with default headers values
func NewCreateUserCredentialsEmailConflict() *CreateUserCredentialsEmailConflict {
	return &CreateUserCredentialsEmailConflict{}
}

/*CreateUserCredentialsEmailConflict handles this case with default header values.

Resource Already Exists
*/
type CreateUserCredentialsEmailConflict struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsEmailConflict) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email][%d] createUserCredentialsEmailConflict  %+v", 409, o.Payload)
}

func (o *CreateUserCredentialsEmailConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsEmailUnprocessableEntity creates a CreateUserCredentialsEmailUnprocessableEntity with default headers values
func NewCreateUserCredentialsEmailUnprocessableEntity() *CreateUserCredentialsEmailUnprocessableEntity {
	return &CreateUserCredentialsEmailUnprocessableEntity{}
}

/*CreateUserCredentialsEmailUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateUserCredentialsEmailUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateUserCredentialsEmailUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_email][%d] createUserCredentialsEmailUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUserCredentialsEmailUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
