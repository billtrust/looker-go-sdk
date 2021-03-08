// Code generated by go-swagger; DO NOT EDIT.

package folder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// DeleteFolderReader is a Reader for the DeleteFolder structure.
type DeleteFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFolderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFolderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFolderNoContent creates a DeleteFolderNoContent with default headers values
func NewDeleteFolderNoContent() *DeleteFolderNoContent {
	return &DeleteFolderNoContent{}
}

/* DeleteFolderNoContent describes a response with status code 204, with default header values.

Successfully deleted.
*/
type DeleteFolderNoContent struct {
	Payload string
}

func (o *DeleteFolderNoContent) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_id}][%d] deleteFolderNoContent  %+v", 204, o.Payload)
}
func (o *DeleteFolderNoContent) GetPayload() string {
	return o.Payload
}

func (o *DeleteFolderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderBadRequest creates a DeleteFolderBadRequest with default headers values
func NewDeleteFolderBadRequest() *DeleteFolderBadRequest {
	return &DeleteFolderBadRequest{}
}

/* DeleteFolderBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteFolderBadRequest struct {
	Payload *models.Error
}

func (o *DeleteFolderBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_id}][%d] deleteFolderBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteFolderBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderNotFound creates a DeleteFolderNotFound with default headers values
func NewDeleteFolderNotFound() *DeleteFolderNotFound {
	return &DeleteFolderNotFound{}
}

/* DeleteFolderNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteFolderNotFound struct {
	Payload *models.Error
}

func (o *DeleteFolderNotFound) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_id}][%d] deleteFolderNotFound  %+v", 404, o.Payload)
}
func (o *DeleteFolderNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderTooManyRequests creates a DeleteFolderTooManyRequests with default headers values
func NewDeleteFolderTooManyRequests() *DeleteFolderTooManyRequests {
	return &DeleteFolderTooManyRequests{}
}

/* DeleteFolderTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteFolderTooManyRequests struct {
	Payload *models.Error
}

func (o *DeleteFolderTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_id}][%d] deleteFolderTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteFolderTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteFolderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
