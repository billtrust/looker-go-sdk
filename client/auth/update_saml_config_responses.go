// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// UpdateSamlConfigReader is a Reader for the UpdateSamlConfig structure.
type UpdateSamlConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSamlConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSamlConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateSamlConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateSamlConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateSamlConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSamlConfigOK creates a UpdateSamlConfigOK with default headers values
func NewUpdateSamlConfigOK() *UpdateSamlConfigOK {
	return &UpdateSamlConfigOK{}
}

/*UpdateSamlConfigOK handles this case with default header values.

New state for SAML Configuration.
*/
type UpdateSamlConfigOK struct {
	Payload *models.SamlConfig
}

func (o *UpdateSamlConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /saml_config][%d] updateSamlConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateSamlConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SamlConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigBadRequest creates a UpdateSamlConfigBadRequest with default headers values
func NewUpdateSamlConfigBadRequest() *UpdateSamlConfigBadRequest {
	return &UpdateSamlConfigBadRequest{}
}

/*UpdateSamlConfigBadRequest handles this case with default header values.

Bad Request
*/
type UpdateSamlConfigBadRequest struct {
	Payload *models.Error
}

func (o *UpdateSamlConfigBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /saml_config][%d] updateSamlConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSamlConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigNotFound creates a UpdateSamlConfigNotFound with default headers values
func NewUpdateSamlConfigNotFound() *UpdateSamlConfigNotFound {
	return &UpdateSamlConfigNotFound{}
}

/*UpdateSamlConfigNotFound handles this case with default header values.

Not Found
*/
type UpdateSamlConfigNotFound struct {
	Payload *models.Error
}

func (o *UpdateSamlConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /saml_config][%d] updateSamlConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSamlConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSamlConfigUnprocessableEntity creates a UpdateSamlConfigUnprocessableEntity with default headers values
func NewUpdateSamlConfigUnprocessableEntity() *UpdateSamlConfigUnprocessableEntity {
	return &UpdateSamlConfigUnprocessableEntity{}
}

/*UpdateSamlConfigUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateSamlConfigUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateSamlConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /saml_config][%d] updateSamlConfigUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateSamlConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
