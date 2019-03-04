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

// PermissionSetReader is a Reader for the PermissionSet structure.
type PermissionSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermissionSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPermissionSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPermissionSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPermissionSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPermissionSetOK creates a PermissionSetOK with default headers values
func NewPermissionSetOK() *PermissionSetOK {
	return &PermissionSetOK{}
}

/*PermissionSetOK handles this case with default header values.

Permission Set
*/
type PermissionSetOK struct {
	Payload *models.PermissionSet
}

func (o *PermissionSetOK) Error() string {
	return fmt.Sprintf("[GET /permission_sets/{permission_set_id}][%d] permissionSetOK  %+v", 200, o.Payload)
}

func (o *PermissionSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermissionSetBadRequest creates a PermissionSetBadRequest with default headers values
func NewPermissionSetBadRequest() *PermissionSetBadRequest {
	return &PermissionSetBadRequest{}
}

/*PermissionSetBadRequest handles this case with default header values.

Bad Request
*/
type PermissionSetBadRequest struct {
	Payload *models.Error
}

func (o *PermissionSetBadRequest) Error() string {
	return fmt.Sprintf("[GET /permission_sets/{permission_set_id}][%d] permissionSetBadRequest  %+v", 400, o.Payload)
}

func (o *PermissionSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermissionSetNotFound creates a PermissionSetNotFound with default headers values
func NewPermissionSetNotFound() *PermissionSetNotFound {
	return &PermissionSetNotFound{}
}

/*PermissionSetNotFound handles this case with default header values.

Not Found
*/
type PermissionSetNotFound struct {
	Payload *models.Error
}

func (o *PermissionSetNotFound) Error() string {
	return fmt.Sprintf("[GET /permission_sets/{permission_set_id}][%d] permissionSetNotFound  %+v", 404, o.Payload)
}

func (o *PermissionSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
