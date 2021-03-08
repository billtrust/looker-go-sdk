// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"your-damain.com/swagger/looker-api-golang/models"
)

// TestLdapConfigUserInfoReader is a Reader for the TestLdapConfigUserInfo structure.
type TestLdapConfigUserInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestLdapConfigUserInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestLdapConfigUserInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestLdapConfigUserInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTestLdapConfigUserInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTestLdapConfigUserInfoUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTestLdapConfigUserInfoOK creates a TestLdapConfigUserInfoOK with default headers values
func NewTestLdapConfigUserInfoOK() *TestLdapConfigUserInfoOK {
	return &TestLdapConfigUserInfoOK{}
}

/* TestLdapConfigUserInfoOK describes a response with status code 200, with default header values.

Result info.
*/
type TestLdapConfigUserInfoOK struct {
	Payload *models.LDAPConfigTestResult
}

func (o *TestLdapConfigUserInfoOK) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_info][%d] testLdapConfigUserInfoOK  %+v", 200, o.Payload)
}
func (o *TestLdapConfigUserInfoOK) GetPayload() *models.LDAPConfigTestResult {
	return o.Payload
}

func (o *TestLdapConfigUserInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LDAPConfigTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigUserInfoBadRequest creates a TestLdapConfigUserInfoBadRequest with default headers values
func NewTestLdapConfigUserInfoBadRequest() *TestLdapConfigUserInfoBadRequest {
	return &TestLdapConfigUserInfoBadRequest{}
}

/* TestLdapConfigUserInfoBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TestLdapConfigUserInfoBadRequest struct {
	Payload *models.Error
}

func (o *TestLdapConfigUserInfoBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_info][%d] testLdapConfigUserInfoBadRequest  %+v", 400, o.Payload)
}
func (o *TestLdapConfigUserInfoBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestLdapConfigUserInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigUserInfoNotFound creates a TestLdapConfigUserInfoNotFound with default headers values
func NewTestLdapConfigUserInfoNotFound() *TestLdapConfigUserInfoNotFound {
	return &TestLdapConfigUserInfoNotFound{}
}

/* TestLdapConfigUserInfoNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TestLdapConfigUserInfoNotFound struct {
	Payload *models.Error
}

func (o *TestLdapConfigUserInfoNotFound) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_info][%d] testLdapConfigUserInfoNotFound  %+v", 404, o.Payload)
}
func (o *TestLdapConfigUserInfoNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestLdapConfigUserInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapConfigUserInfoUnprocessableEntity creates a TestLdapConfigUserInfoUnprocessableEntity with default headers values
func NewTestLdapConfigUserInfoUnprocessableEntity() *TestLdapConfigUserInfoUnprocessableEntity {
	return &TestLdapConfigUserInfoUnprocessableEntity{}
}

/* TestLdapConfigUserInfoUnprocessableEntity describes a response with status code 422, with default header values.

Validation Error
*/
type TestLdapConfigUserInfoUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *TestLdapConfigUserInfoUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /ldap_config/test_user_info][%d] testLdapConfigUserInfoUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *TestLdapConfigUserInfoUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *TestLdapConfigUserInfoUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
