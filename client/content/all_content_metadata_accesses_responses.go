// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// AllContentMetadataAccessesReader is a Reader for the AllContentMetadataAccesses structure.
type AllContentMetadataAccessesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllContentMetadataAccessesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllContentMetadataAccessesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllContentMetadataAccessesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllContentMetadataAccessesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllContentMetadataAccessesOK creates a AllContentMetadataAccessesOK with default headers values
func NewAllContentMetadataAccessesOK() *AllContentMetadataAccessesOK {
	return &AllContentMetadataAccessesOK{}
}

/* AllContentMetadataAccessesOK describes a response with status code 200, with default header values.

Content Metadata Access
*/
type AllContentMetadataAccessesOK struct {
	Payload []*models.ContentMetaGroupUser
}

func (o *AllContentMetadataAccessesOK) Error() string {
	return fmt.Sprintf("[GET /content_metadata_access][%d] allContentMetadataAccessesOK  %+v", 200, o.Payload)
}
func (o *AllContentMetadataAccessesOK) GetPayload() []*models.ContentMetaGroupUser {
	return o.Payload
}

func (o *AllContentMetadataAccessesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllContentMetadataAccessesBadRequest creates a AllContentMetadataAccessesBadRequest with default headers values
func NewAllContentMetadataAccessesBadRequest() *AllContentMetadataAccessesBadRequest {
	return &AllContentMetadataAccessesBadRequest{}
}

/* AllContentMetadataAccessesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllContentMetadataAccessesBadRequest struct {
	Payload *models.Error
}

func (o *AllContentMetadataAccessesBadRequest) Error() string {
	return fmt.Sprintf("[GET /content_metadata_access][%d] allContentMetadataAccessesBadRequest  %+v", 400, o.Payload)
}
func (o *AllContentMetadataAccessesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllContentMetadataAccessesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllContentMetadataAccessesNotFound creates a AllContentMetadataAccessesNotFound with default headers values
func NewAllContentMetadataAccessesNotFound() *AllContentMetadataAccessesNotFound {
	return &AllContentMetadataAccessesNotFound{}
}

/* AllContentMetadataAccessesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllContentMetadataAccessesNotFound struct {
	Payload *models.Error
}

func (o *AllContentMetadataAccessesNotFound) Error() string {
	return fmt.Sprintf("[GET /content_metadata_access][%d] allContentMetadataAccessesNotFound  %+v", 404, o.Payload)
}
func (o *AllContentMetadataAccessesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllContentMetadataAccessesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}