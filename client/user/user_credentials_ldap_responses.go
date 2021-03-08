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

// UserCredentialsLdapReader is a Reader for the UserCredentialsLdap structure.
type UserCredentialsLdapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCredentialsLdapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserCredentialsLdapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserCredentialsLdapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserCredentialsLdapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserCredentialsLdapOK creates a UserCredentialsLdapOK with default headers values
func NewUserCredentialsLdapOK() *UserCredentialsLdapOK {
	return &UserCredentialsLdapOK{}
}

/* UserCredentialsLdapOK describes a response with status code 200, with default header values.

LDAP Credential
*/
type UserCredentialsLdapOK struct {
	Payload *models.CredentialsLDAP
}

func (o *UserCredentialsLdapOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_ldap][%d] userCredentialsLdapOK  %+v", 200, o.Payload)
}
func (o *UserCredentialsLdapOK) GetPayload() *models.CredentialsLDAP {
	return o.Payload
}

func (o *UserCredentialsLdapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsLDAP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsLdapBadRequest creates a UserCredentialsLdapBadRequest with default headers values
func NewUserCredentialsLdapBadRequest() *UserCredentialsLdapBadRequest {
	return &UserCredentialsLdapBadRequest{}
}

/* UserCredentialsLdapBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserCredentialsLdapBadRequest struct {
	Payload *models.Error
}

func (o *UserCredentialsLdapBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_ldap][%d] userCredentialsLdapBadRequest  %+v", 400, o.Payload)
}
func (o *UserCredentialsLdapBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserCredentialsLdapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsLdapNotFound creates a UserCredentialsLdapNotFound with default headers values
func NewUserCredentialsLdapNotFound() *UserCredentialsLdapNotFound {
	return &UserCredentialsLdapNotFound{}
}

/* UserCredentialsLdapNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserCredentialsLdapNotFound struct {
	Payload *models.Error
}

func (o *UserCredentialsLdapNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_ldap][%d] userCredentialsLdapNotFound  %+v", 404, o.Payload)
}
func (o *UserCredentialsLdapNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserCredentialsLdapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
