// Code generated by go-swagger; DO NOT EDIT.

package color_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// ColorCollectionsStandardReader is a Reader for the ColorCollectionsStandard structure.
type ColorCollectionsStandardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColorCollectionsStandardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColorCollectionsStandardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewColorCollectionsStandardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewColorCollectionsStandardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewColorCollectionsStandardOK creates a ColorCollectionsStandardOK with default headers values
func NewColorCollectionsStandardOK() *ColorCollectionsStandardOK {
	return &ColorCollectionsStandardOK{}
}

/* ColorCollectionsStandardOK describes a response with status code 200, with default header values.

ColorCollections
*/
type ColorCollectionsStandardOK struct {
	Payload []*models.ColorCollection
}

func (o *ColorCollectionsStandardOK) Error() string {
	return fmt.Sprintf("[GET /color_collections/standard][%d] colorCollectionsStandardOK  %+v", 200, o.Payload)
}
func (o *ColorCollectionsStandardOK) GetPayload() []*models.ColorCollection {
	return o.Payload
}

func (o *ColorCollectionsStandardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColorCollectionsStandardBadRequest creates a ColorCollectionsStandardBadRequest with default headers values
func NewColorCollectionsStandardBadRequest() *ColorCollectionsStandardBadRequest {
	return &ColorCollectionsStandardBadRequest{}
}

/* ColorCollectionsStandardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ColorCollectionsStandardBadRequest struct {
	Payload *models.Error
}

func (o *ColorCollectionsStandardBadRequest) Error() string {
	return fmt.Sprintf("[GET /color_collections/standard][%d] colorCollectionsStandardBadRequest  %+v", 400, o.Payload)
}
func (o *ColorCollectionsStandardBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ColorCollectionsStandardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColorCollectionsStandardNotFound creates a ColorCollectionsStandardNotFound with default headers values
func NewColorCollectionsStandardNotFound() *ColorCollectionsStandardNotFound {
	return &ColorCollectionsStandardNotFound{}
}

/* ColorCollectionsStandardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ColorCollectionsStandardNotFound struct {
	Payload *models.Error
}

func (o *ColorCollectionsStandardNotFound) Error() string {
	return fmt.Sprintf("[GET /color_collections/standard][%d] colorCollectionsStandardNotFound  %+v", 404, o.Payload)
}
func (o *ColorCollectionsStandardNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ColorCollectionsStandardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
