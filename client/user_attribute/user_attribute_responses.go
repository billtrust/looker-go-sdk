// Code generated by go-swagger; DO NOT EDIT.

package user_attribute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// UserAttributeReader is a Reader for the UserAttribute structure.
type UserAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserAttributeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUserAttributeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUserAttributeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserAttributeOK creates a UserAttributeOK with default headers values
func NewUserAttributeOK() *UserAttributeOK {
	return &UserAttributeOK{}
}

/*UserAttributeOK handles this case with default header values.

User Attribute
*/
type UserAttributeOK struct {
	Payload *models.UserAttribute
}

func (o *UserAttributeOK) Error() string {
	return fmt.Sprintf("[GET /user_attributes/{user_attribute_id}][%d] userAttributeOK  %+v", 200, o.Payload)
}

func (o *UserAttributeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAttribute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAttributeBadRequest creates a UserAttributeBadRequest with default headers values
func NewUserAttributeBadRequest() *UserAttributeBadRequest {
	return &UserAttributeBadRequest{}
}

/*UserAttributeBadRequest handles this case with default header values.

Bad Request
*/
type UserAttributeBadRequest struct {
	Payload *models.Error
}

func (o *UserAttributeBadRequest) Error() string {
	return fmt.Sprintf("[GET /user_attributes/{user_attribute_id}][%d] userAttributeBadRequest  %+v", 400, o.Payload)
}

func (o *UserAttributeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserAttributeNotFound creates a UserAttributeNotFound with default headers values
func NewUserAttributeNotFound() *UserAttributeNotFound {
	return &UserAttributeNotFound{}
}

/*UserAttributeNotFound handles this case with default header values.

Not Found
*/
type UserAttributeNotFound struct {
	Payload *models.Error
}

func (o *UserAttributeNotFound) Error() string {
	return fmt.Sprintf("[GET /user_attributes/{user_attribute_id}][%d] userAttributeNotFound  %+v", 404, o.Payload)
}

func (o *UserAttributeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
