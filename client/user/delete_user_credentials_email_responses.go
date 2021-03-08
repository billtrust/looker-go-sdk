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

// DeleteUserCredentialsEmailReader is a Reader for the DeleteUserCredentialsEmail structure.
type DeleteUserCredentialsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserCredentialsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserCredentialsEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserCredentialsEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserCredentialsEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserCredentialsEmailTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserCredentialsEmailNoContent creates a DeleteUserCredentialsEmailNoContent with default headers values
func NewDeleteUserCredentialsEmailNoContent() *DeleteUserCredentialsEmailNoContent {
	return &DeleteUserCredentialsEmailNoContent{}
}

/* DeleteUserCredentialsEmailNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteUserCredentialsEmailNoContent struct {
	Payload string
}

func (o *DeleteUserCredentialsEmailNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_email][%d] deleteUserCredentialsEmailNoContent  %+v", 204, o.Payload)
}
func (o *DeleteUserCredentialsEmailNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteUserCredentialsEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsEmailBadRequest creates a DeleteUserCredentialsEmailBadRequest with default headers values
func NewDeleteUserCredentialsEmailBadRequest() *DeleteUserCredentialsEmailBadRequest {
	return &DeleteUserCredentialsEmailBadRequest{}
}

/* DeleteUserCredentialsEmailBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserCredentialsEmailBadRequest struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsEmailBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_email][%d] deleteUserCredentialsEmailBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserCredentialsEmailBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsEmailNotFound creates a DeleteUserCredentialsEmailNotFound with default headers values
func NewDeleteUserCredentialsEmailNotFound() *DeleteUserCredentialsEmailNotFound {
	return &DeleteUserCredentialsEmailNotFound{}
}

/* DeleteUserCredentialsEmailNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserCredentialsEmailNotFound struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsEmailNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_email][%d] deleteUserCredentialsEmailNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserCredentialsEmailNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserCredentialsEmailTooManyRequests creates a DeleteUserCredentialsEmailTooManyRequests with default headers values
func NewDeleteUserCredentialsEmailTooManyRequests() *DeleteUserCredentialsEmailTooManyRequests {
	return &DeleteUserCredentialsEmailTooManyRequests{}
}

/* DeleteUserCredentialsEmailTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteUserCredentialsEmailTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteUserCredentialsEmailTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/credentials_email][%d] deleteUserCredentialsEmailTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteUserCredentialsEmailTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUserCredentialsEmailTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
