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

// ModelSetReader is a Reader for the ModelSet structure.
type ModelSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModelSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModelSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewModelSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModelSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModelSetOK creates a ModelSetOK with default headers values
func NewModelSetOK() *ModelSetOK {
	return &ModelSetOK{}
}

/* ModelSetOK describes a response with status code 200, with default header values.

Specified model set.
*/
type ModelSetOK struct {
	Payload *models.ModelSet
}

func (o *ModelSetOK) Error() string {
	return fmt.Sprintf("[GET /model_sets/{model_set_id}][%d] modelSetOK  %+v", 200, o.Payload)
}
func (o *ModelSetOK) GetPayload() *models.ModelSet {
	return o.Payload
}

func (o *ModelSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModelSetBadRequest creates a ModelSetBadRequest with default headers values
func NewModelSetBadRequest() *ModelSetBadRequest {
	return &ModelSetBadRequest{}
}

/* ModelSetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ModelSetBadRequest struct {
	Payload *models.Error
}

func (o *ModelSetBadRequest) Error() string {
	return fmt.Sprintf("[GET /model_sets/{model_set_id}][%d] modelSetBadRequest  %+v", 400, o.Payload)
}
func (o *ModelSetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModelSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModelSetNotFound creates a ModelSetNotFound with default headers values
func NewModelSetNotFound() *ModelSetNotFound {
	return &ModelSetNotFound{}
}

/* ModelSetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ModelSetNotFound struct {
	Payload *models.Error
}

func (o *ModelSetNotFound) Error() string {
	return fmt.Sprintf("[GET /model_sets/{model_set_id}][%d] modelSetNotFound  %+v", 404, o.Payload)
}
func (o *ModelSetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModelSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
