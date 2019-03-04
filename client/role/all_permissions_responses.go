// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// AllPermissionsReader is a Reader for the AllPermissions structure.
type AllPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAllPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAllPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAllPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAllPermissionsOK creates a AllPermissionsOK with default headers values
func NewAllPermissionsOK() *AllPermissionsOK {
	return &AllPermissionsOK{}
}

/*AllPermissionsOK handles this case with default header values.

Permission
*/
type AllPermissionsOK struct {
	Payload []*models.Permission
}

func (o *AllPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] allPermissionsOK  %+v", 200, o.Payload)
}

func (o *AllPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllPermissionsBadRequest creates a AllPermissionsBadRequest with default headers values
func NewAllPermissionsBadRequest() *AllPermissionsBadRequest {
	return &AllPermissionsBadRequest{}
}

/*AllPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type AllPermissionsBadRequest struct {
	Payload *models.Error
}

func (o *AllPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] allPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *AllPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllPermissionsNotFound creates a AllPermissionsNotFound with default headers values
func NewAllPermissionsNotFound() *AllPermissionsNotFound {
	return &AllPermissionsNotFound{}
}

/*AllPermissionsNotFound handles this case with default header values.

Not Found
*/
type AllPermissionsNotFound struct {
	Payload *models.Error
}

func (o *AllPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] allPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *AllPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
