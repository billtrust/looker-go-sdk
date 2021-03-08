// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// CreateModelSetReader is a Reader for the CreateModelSet structure.
type CreateModelSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateModelSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateModelSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateModelSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateModelSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateModelSetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateModelSetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateModelSetOK creates a CreateModelSetOK with default headers values
func NewCreateModelSetOK() *CreateModelSetOK {
	return &CreateModelSetOK{}
}

/* CreateModelSetOK describes a response with status code 200, with default header values.

Newly created ModelSet
*/
type CreateModelSetOK struct {
	Payload *models.ModelSet
}

func (o *CreateModelSetOK) Error() string {
	return fmt.Sprintf("[POST /model_sets][%d] createModelSetOK  %+v", 200, o.Payload)
}
func (o *CreateModelSetOK) GetPayload() *models.ModelSet {
	return o.Payload
}

func (o *CreateModelSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelSetBadRequest creates a CreateModelSetBadRequest with default headers values
func NewCreateModelSetBadRequest() *CreateModelSetBadRequest {
	return &CreateModelSetBadRequest{}
}

/* CreateModelSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateModelSetBadRequest struct {
	Payload *models.Error
}

func (o *CreateModelSetBadRequest) Error() string {
	return fmt.Sprintf("[POST /model_sets][%d] createModelSetBadRequest  %+v", 400, o.Payload)
}
func (o *CreateModelSetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateModelSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelSetNotFound creates a CreateModelSetNotFound with default headers values
func NewCreateModelSetNotFound() *CreateModelSetNotFound {
	return &CreateModelSetNotFound{}
}

/* CreateModelSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateModelSetNotFound struct {
	Payload *models.Error
}

func (o *CreateModelSetNotFound) Error() string {
	return fmt.Sprintf("[POST /model_sets][%d] createModelSetNotFound  %+v", 404, o.Payload)
}
func (o *CreateModelSetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateModelSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelSetConflict creates a CreateModelSetConflict with default headers values
func NewCreateModelSetConflict() *CreateModelSetConflict {
	return &CreateModelSetConflict{}
}

/* CreateModelSetConflict describes a response with status code 409, with default header values.

Resource Already Exists
*/
type CreateModelSetConflict struct {
	Payload *models.Error
}

func (o *CreateModelSetConflict) Error() string {
	return fmt.Sprintf("[POST /model_sets][%d] createModelSetConflict  %+v", 409, o.Payload)
}
func (o *CreateModelSetConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateModelSetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateModelSetUnprocessableEntity creates a CreateModelSetUnprocessableEntity with default headers values
func NewCreateModelSetUnprocessableEntity() *CreateModelSetUnprocessableEntity {
	return &CreateModelSetUnprocessableEntity{}
}

/* CreateModelSetUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type CreateModelSetUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateModelSetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /model_sets][%d] createModelSetUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateModelSetUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateModelSetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
