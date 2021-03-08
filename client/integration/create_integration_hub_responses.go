// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// CreateIntegrationHubReader is a Reader for the CreateIntegrationHub structure.
type CreateIntegrationHubReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntegrationHubReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIntegrationHubOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIntegrationHubBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateIntegrationHubNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateIntegrationHubConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateIntegrationHubUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateIntegrationHubTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateIntegrationHubOK creates a CreateIntegrationHubOK with default headers values
func NewCreateIntegrationHubOK() *CreateIntegrationHubOK {
	return &CreateIntegrationHubOK{}
}

/* CreateIntegrationHubOK describes a response with status code 200, with default header values.

Integration Hub
*/
type CreateIntegrationHubOK struct {
	Payload *models.IntegrationHub
}

func (o *CreateIntegrationHubOK) Error() string {
	return fmt.Sprintf("[POST /integration_hubs][%d] createIntegrationHubOK  %+v", 200, o.Payload)
}
func (o *CreateIntegrationHubOK) GetPayload() *models.IntegrationHub {
	return o.Payload
}

func (o *CreateIntegrationHubOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationHub)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationHubBadRequest creates a CreateIntegrationHubBadRequest with default headers values
func NewCreateIntegrationHubBadRequest() *CreateIntegrationHubBadRequest {
	return &CreateIntegrationHubBadRequest{}
}

/* CreateIntegrationHubBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateIntegrationHubBadRequest struct {
	Payload *models.Error
}

func (o *CreateIntegrationHubBadRequest) Error() string {
	return fmt.Sprintf("[POST /integration_hubs][%d] createIntegrationHubBadRequest  %+v", 400, o.Payload)
}
func (o *CreateIntegrationHubBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntegrationHubBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationHubNotFound creates a CreateIntegrationHubNotFound with default headers values
func NewCreateIntegrationHubNotFound() *CreateIntegrationHubNotFound {
	return &CreateIntegrationHubNotFound{}
}

/* CreateIntegrationHubNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateIntegrationHubNotFound struct {
	Payload *models.Error
}

func (o *CreateIntegrationHubNotFound) Error() string {
	return fmt.Sprintf("[POST /integration_hubs][%d] createIntegrationHubNotFound  %+v", 404, o.Payload)
}
func (o *CreateIntegrationHubNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntegrationHubNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationHubConflict creates a CreateIntegrationHubConflict with default headers values
func NewCreateIntegrationHubConflict() *CreateIntegrationHubConflict {
	return &CreateIntegrationHubConflict{}
}

/* CreateIntegrationHubConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateIntegrationHubConflict struct {
	Payload *models.Error
}

func (o *CreateIntegrationHubConflict) Error() string {
	return fmt.Sprintf("[POST /integration_hubs][%d] createIntegrationHubConflict  %+v", 409, o.Payload)
}
func (o *CreateIntegrationHubConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntegrationHubConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationHubUnprocessableEntity creates a CreateIntegrationHubUnprocessableEntity with default headers values
func NewCreateIntegrationHubUnprocessableEntity() *CreateIntegrationHubUnprocessableEntity {
	return &CreateIntegrationHubUnprocessableEntity{}
}

/* CreateIntegrationHubUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateIntegrationHubUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateIntegrationHubUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /integration_hubs][%d] createIntegrationHubUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateIntegrationHubUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateIntegrationHubUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationHubTooManyRequests creates a CreateIntegrationHubTooManyRequests with default headers values
func NewCreateIntegrationHubTooManyRequests() *CreateIntegrationHubTooManyRequests {
	return &CreateIntegrationHubTooManyRequests{}
}

/* CreateIntegrationHubTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateIntegrationHubTooManyRequests struct {
	Payload *models.Error
}

func (o *CreateIntegrationHubTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /integration_hubs][%d] createIntegrationHubTooManyRequests  %+v", 429, o.Payload)
}
func (o *CreateIntegrationHubTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntegrationHubTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
