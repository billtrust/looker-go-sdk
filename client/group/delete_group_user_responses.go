// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// DeleteGroupUserReader is a Reader for the DeleteGroupUser structure.
type DeleteGroupUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGroupUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteGroupUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGroupUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGroupUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGroupUserNoContent creates a DeleteGroupUserNoContent with default headers values
func NewDeleteGroupUserNoContent() *DeleteGroupUserNoContent {
	return &DeleteGroupUserNoContent{}
}

/* DeleteGroupUserNoContent describes a response with status code 204, with default header values.

User successfully removed from group
*/
type DeleteGroupUserNoContent struct {
}

func (o *DeleteGroupUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}/users/{user_id}][%d] deleteGroupUserNoContent ", 204)
}

func (o *DeleteGroupUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupUserBadRequest creates a DeleteGroupUserBadRequest with default headers values
func NewDeleteGroupUserBadRequest() *DeleteGroupUserBadRequest {
	return &DeleteGroupUserBadRequest{}
}

/* DeleteGroupUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteGroupUserBadRequest struct {
	Payload *models.Error
}

func (o *DeleteGroupUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}/users/{user_id}][%d] deleteGroupUserBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteGroupUserBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupUserForbidden creates a DeleteGroupUserForbidden with default headers values
func NewDeleteGroupUserForbidden() *DeleteGroupUserForbidden {
	return &DeleteGroupUserForbidden{}
}

/* DeleteGroupUserForbidden describes a response with status code 403, with default header values.

Permission Denied
*/
type DeleteGroupUserForbidden struct {
	Payload *models.Error
}

func (o *DeleteGroupUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}/users/{user_id}][%d] deleteGroupUserForbidden  %+v", 403, o.Payload)
}
func (o *DeleteGroupUserForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupUserNotFound creates a DeleteGroupUserNotFound with default headers values
func NewDeleteGroupUserNotFound() *DeleteGroupUserNotFound {
	return &DeleteGroupUserNotFound{}
}

/* DeleteGroupUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteGroupUserNotFound struct {
	Payload *models.Error
}

func (o *DeleteGroupUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /groups/{group_id}/users/{user_id}][%d] deleteGroupUserNotFound  %+v", 404, o.Payload)
}
func (o *DeleteGroupUserNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
