// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// HomepageReader is a Reader for the Homepage structure.
type HomepageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HomepageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHomepageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHomepageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHomepageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHomepageOK creates a HomepageOK with default headers values
func NewHomepageOK() *HomepageOK {
	return &HomepageOK{}
}

/* HomepageOK describes a response with status code 200, with default header values.

Homepage
*/
type HomepageOK struct {
	Payload *models.Homepage
}

func (o *HomepageOK) Error() string {
	return fmt.Sprintf("[GET /homepages/{homepage_id}][%d] homepageOK  %+v", 200, o.Payload)
}
func (o *HomepageOK) GetPayload() *models.Homepage {
	return o.Payload
}

func (o *HomepageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Homepage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHomepageBadRequest creates a HomepageBadRequest with default headers values
func NewHomepageBadRequest() *HomepageBadRequest {
	return &HomepageBadRequest{}
}

/* HomepageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type HomepageBadRequest struct {
	Payload *models.Error
}

func (o *HomepageBadRequest) Error() string {
	return fmt.Sprintf("[GET /homepages/{homepage_id}][%d] homepageBadRequest  %+v", 400, o.Payload)
}
func (o *HomepageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *HomepageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHomepageNotFound creates a HomepageNotFound with default headers values
func NewHomepageNotFound() *HomepageNotFound {
	return &HomepageNotFound{}
}

/* HomepageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type HomepageNotFound struct {
	Payload *models.Error
}

func (o *HomepageNotFound) Error() string {
	return fmt.Sprintf("[GET /homepages/{homepage_id}][%d] homepageNotFound  %+v", 404, o.Payload)
}
func (o *HomepageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *HomepageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
