// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// IntegrationReader is a Reader for the Integration structure.
type IntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationOK creates a IntegrationOK with default headers values
func NewIntegrationOK() *IntegrationOK {
	return &IntegrationOK{}
}

/*IntegrationOK handles this case with default header values.

Integration
*/
type IntegrationOK struct {
	Payload *models.Integration
}

func (o *IntegrationOK) Error() string {
	return fmt.Sprintf("[GET /integrations/{integration_id}][%d] integrationOK  %+v", 200, o.Payload)
}

func (o *IntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationBadRequest creates a IntegrationBadRequest with default headers values
func NewIntegrationBadRequest() *IntegrationBadRequest {
	return &IntegrationBadRequest{}
}

/*IntegrationBadRequest handles this case with default header values.

Bad Request
*/
type IntegrationBadRequest struct {
	Payload *models.Error
}

func (o *IntegrationBadRequest) Error() string {
	return fmt.Sprintf("[GET /integrations/{integration_id}][%d] integrationBadRequest  %+v", 400, o.Payload)
}

func (o *IntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationNotFound creates a IntegrationNotFound with default headers values
func NewIntegrationNotFound() *IntegrationNotFound {
	return &IntegrationNotFound{}
}

/*IntegrationNotFound handles this case with default header values.

Not Found
*/
type IntegrationNotFound struct {
	Payload *models.Error
}

func (o *IntegrationNotFound) Error() string {
	return fmt.Sprintf("[GET /integrations/{integration_id}][%d] integrationNotFound  %+v", 404, o.Payload)
}

func (o *IntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
