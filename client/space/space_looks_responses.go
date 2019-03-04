// Code generated by go-swagger; DO NOT EDIT.

package space

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// SpaceLooksReader is a Reader for the SpaceLooks structure.
type SpaceLooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpaceLooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSpaceLooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSpaceLooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSpaceLooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSpaceLooksOK creates a SpaceLooksOK with default headers values
func NewSpaceLooksOK() *SpaceLooksOK {
	return &SpaceLooksOK{}
}

/*SpaceLooksOK handles this case with default header values.

Looks
*/
type SpaceLooksOK struct {
	Payload []*models.LookWithQuery
}

func (o *SpaceLooksOK) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/looks][%d] spaceLooksOK  %+v", 200, o.Payload)
}

func (o *SpaceLooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceLooksBadRequest creates a SpaceLooksBadRequest with default headers values
func NewSpaceLooksBadRequest() *SpaceLooksBadRequest {
	return &SpaceLooksBadRequest{}
}

/*SpaceLooksBadRequest handles this case with default header values.

Bad Request
*/
type SpaceLooksBadRequest struct {
	Payload *models.Error
}

func (o *SpaceLooksBadRequest) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/looks][%d] spaceLooksBadRequest  %+v", 400, o.Payload)
}

func (o *SpaceLooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpaceLooksNotFound creates a SpaceLooksNotFound with default headers values
func NewSpaceLooksNotFound() *SpaceLooksNotFound {
	return &SpaceLooksNotFound{}
}

/*SpaceLooksNotFound handles this case with default header values.

Not Found
*/
type SpaceLooksNotFound struct {
	Payload *models.Error
}

func (o *SpaceLooksNotFound) Error() string {
	return fmt.Sprintf("[GET /spaces/{space_id}/looks][%d] spaceLooksNotFound  %+v", 404, o.Payload)
}

func (o *SpaceLooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
