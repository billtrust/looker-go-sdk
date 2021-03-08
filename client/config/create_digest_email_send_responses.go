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

// CreateDigestEmailSendReader is a Reader for the CreateDigestEmailSend structure.
type CreateDigestEmailSendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDigestEmailSendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDigestEmailSendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDigestEmailSendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDigestEmailSendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDigestEmailSendTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDigestEmailSendOK creates a CreateDigestEmailSendOK with default headers values
func NewCreateDigestEmailSendOK() *CreateDigestEmailSendOK {
	return &CreateDigestEmailSendOK{}
}

/* CreateDigestEmailSendOK describes a response with status code 200, with default header values.

Status of generating and sending the data
*/
type CreateDigestEmailSendOK struct {
	Payload *models.DigestEmailSend
}

func (o *CreateDigestEmailSendOK) Error() string {
	return fmt.Sprintf("[POST /digest_email_send][%d] createDigestEmailSendOK  %+v", 200, o.Payload)
}
func (o *CreateDigestEmailSendOK) GetPayload() *models.DigestEmailSend {
	return o.Payload
}

func (o *CreateDigestEmailSendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DigestEmailSend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigestEmailSendBadRequest creates a CreateDigestEmailSendBadRequest with default headers values
func NewCreateDigestEmailSendBadRequest() *CreateDigestEmailSendBadRequest {
	return &CreateDigestEmailSendBadRequest{}
}

/* CreateDigestEmailSendBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateDigestEmailSendBadRequest struct {
	Payload *models.Error
}

func (o *CreateDigestEmailSendBadRequest) Error() string {
	return fmt.Sprintf("[POST /digest_email_send][%d] createDigestEmailSendBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDigestEmailSendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDigestEmailSendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigestEmailSendNotFound creates a CreateDigestEmailSendNotFound with default headers values
func NewCreateDigestEmailSendNotFound() *CreateDigestEmailSendNotFound {
	return &CreateDigestEmailSendNotFound{}
}

/* CreateDigestEmailSendNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateDigestEmailSendNotFound struct {
	Payload *models.Error
}

func (o *CreateDigestEmailSendNotFound) Error() string {
	return fmt.Sprintf("[POST /digest_email_send][%d] createDigestEmailSendNotFound  %+v", 404, o.Payload)
}
func (o *CreateDigestEmailSendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDigestEmailSendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDigestEmailSendTooManyRequests creates a CreateDigestEmailSendTooManyRequests with default headers values
func NewCreateDigestEmailSendTooManyRequests() *CreateDigestEmailSendTooManyRequests {
	return &CreateDigestEmailSendTooManyRequests{}
}

/* CreateDigestEmailSendTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateDigestEmailSendTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateDigestEmailSendTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /digest_email_send][%d] createDigestEmailSendTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateDigestEmailSendTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDigestEmailSendTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}