// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// MeReader is a Reader for the Me structure.
type MeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMeOK creates a MeOK with default headers values
func NewMeOK() *MeOK {
	return &MeOK{}
}

/*MeOK handles this case with default header values.

Current user.
*/
type MeOK struct {
	Payload *models.User
}

func (o *MeOK) Error() string {
	return fmt.Sprintf("[GET /user][%d] meOK  %+v", 200, o.Payload)
}

func (o *MeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMeNotFound creates a MeNotFound with default headers values
func NewMeNotFound() *MeNotFound {
	return &MeNotFound{}
}

/*MeNotFound handles this case with default header values.

Not Found
*/
type MeNotFound struct {
	Payload *models.Error
}

func (o *MeNotFound) Error() string {
	return fmt.Sprintf("[GET /user][%d] meNotFound  %+v", 404, o.Payload)
}

func (o *MeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
