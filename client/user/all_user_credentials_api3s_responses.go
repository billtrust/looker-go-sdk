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

// AllUserCredentialsApi3sReader is a Reader for the AllUserCredentialsApi3s structure.
type AllUserCredentialsApi3sReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllUserCredentialsApi3sReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllUserCredentialsApi3sOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllUserCredentialsApi3sBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllUserCredentialsApi3sNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllUserCredentialsApi3sOK creates a AllUserCredentialsApi3sOK with default headers values
func NewAllUserCredentialsApi3sOK() *AllUserCredentialsApi3sOK {
	return &AllUserCredentialsApi3sOK{}
}

/* AllUserCredentialsApi3sOK describes a response with status code 200, with default header values.

API 3 Credential
*/
type AllUserCredentialsApi3sOK struct {
	Payload []*models.CredentialsApi3
}

func (o *AllUserCredentialsApi3sOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api3][%d] allUserCredentialsApi3sOK  %+v", 200, o.Payload)
}
func (o *AllUserCredentialsApi3sOK) GetPayload() []*models.CredentialsApi3 {
	return o.Payload
}

func (o *AllUserCredentialsApi3sOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserCredentialsApi3sBadRequest creates a AllUserCredentialsApi3sBadRequest with default headers values
func NewAllUserCredentialsApi3sBadRequest() *AllUserCredentialsApi3sBadRequest {
	return &AllUserCredentialsApi3sBadRequest{}
}

/* AllUserCredentialsApi3sBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllUserCredentialsApi3sBadRequest struct {
	Payload *models.Error
}

func (o *AllUserCredentialsApi3sBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api3][%d] allUserCredentialsApi3sBadRequest  %+v", 400, o.Payload)
}
func (o *AllUserCredentialsApi3sBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllUserCredentialsApi3sBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllUserCredentialsApi3sNotFound creates a AllUserCredentialsApi3sNotFound with default headers values
func NewAllUserCredentialsApi3sNotFound() *AllUserCredentialsApi3sNotFound {
	return &AllUserCredentialsApi3sNotFound{}
}

/* AllUserCredentialsApi3sNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllUserCredentialsApi3sNotFound struct {
	Payload *models.Error
}

func (o *AllUserCredentialsApi3sNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_api3][%d] allUserCredentialsApi3sNotFound  %+v", 404, o.Payload)
}
func (o *AllUserCredentialsApi3sNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllUserCredentialsApi3sNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
