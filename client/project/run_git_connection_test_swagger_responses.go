// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// RunGitConnectionTestReader is a Reader for the RunGitConnectionTest structure.
type RunGitConnectionTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunGitConnectionTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunGitConnectionTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunGitConnectionTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunGitConnectionTestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRunGitConnectionTestUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRunGitConnectionTestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRunGitConnectionTestOK creates a RunGitConnectionTestOK with default headers values
func NewRunGitConnectionTestOK() *RunGitConnectionTestOK {
	return &RunGitConnectionTestOK{}
}

/* RunGitConnectionTestOK describes a response with status code 200, with default header values.

Git Connection Test Result
*/
type RunGitConnectionTestOK struct {
	Payload *models.GitConnectionTestResult
}

func (o *RunGitConnectionTestOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests/{test_id}][%d] runGitConnectionTestOK  %+v", 200, o.Payload)
}
func (o *RunGitConnectionTestOK) GetPayload() *models.GitConnectionTestResult {
	return o.Payload
}

func (o *RunGitConnectionTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitConnectionTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunGitConnectionTestBadRequest creates a RunGitConnectionTestBadRequest with default headers values
func NewRunGitConnectionTestBadRequest() *RunGitConnectionTestBadRequest {
	return &RunGitConnectionTestBadRequest{}
}

/* RunGitConnectionTestBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RunGitConnectionTestBadRequest struct {
	Payload *models.Error
}

func (o *RunGitConnectionTestBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests/{test_id}][%d] runGitConnectionTestBadRequest  %+v", 400, o.Payload)
}
func (o *RunGitConnectionTestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunGitConnectionTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunGitConnectionTestNotFound creates a RunGitConnectionTestNotFound with default headers values
func NewRunGitConnectionTestNotFound() *RunGitConnectionTestNotFound {
	return &RunGitConnectionTestNotFound{}
}

/* RunGitConnectionTestNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RunGitConnectionTestNotFound struct {
	Payload *models.Error
}

func (o *RunGitConnectionTestNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests/{test_id}][%d] runGitConnectionTestNotFound  %+v", 404, o.Payload)
}
func (o *RunGitConnectionTestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunGitConnectionTestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunGitConnectionTestUnprocessableEntity creates a RunGitConnectionTestUnprocessableEntity with default headers values
func NewRunGitConnectionTestUnprocessableEntity() *RunGitConnectionTestUnprocessableEntity {
	return &RunGitConnectionTestUnprocessableEntity{}
}

/* RunGitConnectionTestUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type RunGitConnectionTestUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *RunGitConnectionTestUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests/{test_id}][%d] runGitConnectionTestUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *RunGitConnectionTestUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RunGitConnectionTestUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunGitConnectionTestTooManyRequests creates a RunGitConnectionTestTooManyRequests with default headers values
func NewRunGitConnectionTestTooManyRequests() *RunGitConnectionTestTooManyRequests {
	return &RunGitConnectionTestTooManyRequests{}
}

/* RunGitConnectionTestTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RunGitConnectionTestTooManyRequests struct {
	Payload *models.Error
}

func (o *RunGitConnectionTestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_connection_tests/{test_id}][%d] runGitConnectionTestTooManyRequests  %+v", 429, o.Payload)
}
func (o *RunGitConnectionTestTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunGitConnectionTestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
