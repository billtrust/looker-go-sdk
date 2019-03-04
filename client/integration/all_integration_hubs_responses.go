// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// AllIntegrationHubsReader is a Reader for the AllIntegrationHubs structure.
type AllIntegrationHubsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllIntegrationHubsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllIntegrationHubsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllIntegrationHubsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllIntegrationHubsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllIntegrationHubsOK creates a AllIntegrationHubsOK with default headers values
func NewAllIntegrationHubsOK() *AllIntegrationHubsOK {
	return &AllIntegrationHubsOK{}
}

/*AllIntegrationHubsOK handles this case with default header values.

Integration Hub
*/
type AllIntegrationHubsOK struct {
	Payload []*models.IntegrationHub
}

func (o *AllIntegrationHubsOK) Error() string {
	return fmt.Sprintf("[GET /integration_hubs][%d] allIntegrationHubsOK  %+v", 200, o.Payload)
}

func (o *AllIntegrationHubsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllIntegrationHubsBadRequest creates a AllIntegrationHubsBadRequest with default headers values
func NewAllIntegrationHubsBadRequest() *AllIntegrationHubsBadRequest {
	return &AllIntegrationHubsBadRequest{}
}

/*AllIntegrationHubsBadRequest handles this case with default header values.

Bad Request
*/
type AllIntegrationHubsBadRequest struct {
	Payload *models.Error
}

func (o *AllIntegrationHubsBadRequest) Error() string {
	return fmt.Sprintf("[GET /integration_hubs][%d] allIntegrationHubsBadRequest  %+v", 400, o.Payload)
}

func (o *AllIntegrationHubsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllIntegrationHubsNotFound creates a AllIntegrationHubsNotFound with default headers values
func NewAllIntegrationHubsNotFound() *AllIntegrationHubsNotFound {
	return &AllIntegrationHubsNotFound{}
}

/*AllIntegrationHubsNotFound handles this case with default header values.

Not Found
*/
type AllIntegrationHubsNotFound struct {
	Payload *models.Error
}

func (o *AllIntegrationHubsNotFound) Error() string {
	return fmt.Sprintf("[GET /integration_hubs][%d] allIntegrationHubsNotFound  %+v", 404, o.Payload)
}

func (o *AllIntegrationHubsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
