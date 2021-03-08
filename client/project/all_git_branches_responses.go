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

// AllGitBranchesReader is a Reader for the AllGitBranches structure.
type AllGitBranchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllGitBranchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllGitBranchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAllGitBranchesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAllGitBranchesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAllGitBranchesOK creates a AllGitBranchesOK with default headers values
func NewAllGitBranchesOK() *AllGitBranchesOK {
	return &AllGitBranchesOK{}
}

/* AllGitBranchesOK describes a response with status code 200, with default header values.

Git Branch
*/
type AllGitBranchesOK struct {
	Payload []*models.GitBranch
}

func (o *AllGitBranchesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_branches][%d] allGitBranchesOK  %+v", 200, o.Payload)
}
func (o *AllGitBranchesOK) GetPayload() []*models.GitBranch {
	return o.Payload
}

func (o *AllGitBranchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGitBranchesBadRequest creates a AllGitBranchesBadRequest with default headers values
func NewAllGitBranchesBadRequest() *AllGitBranchesBadRequest {
	return &AllGitBranchesBadRequest{}
}

/* AllGitBranchesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AllGitBranchesBadRequest struct {
	Payload *models.Error
}

func (o *AllGitBranchesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_branches][%d] allGitBranchesBadRequest  %+v", 400, o.Payload)
}
func (o *AllGitBranchesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllGitBranchesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllGitBranchesNotFound creates a AllGitBranchesNotFound with default headers values
func NewAllGitBranchesNotFound() *AllGitBranchesNotFound {
	return &AllGitBranchesNotFound{}
}

/* AllGitBranchesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AllGitBranchesNotFound struct {
	Payload *models.Error
}

func (o *AllGitBranchesNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/git_branches][%d] allGitBranchesNotFound  %+v", 404, o.Payload)
}
func (o *AllGitBranchesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AllGitBranchesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
