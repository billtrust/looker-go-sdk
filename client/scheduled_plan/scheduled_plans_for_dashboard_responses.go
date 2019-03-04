// Code generated by go-swagger; DO NOT EDIT.

package scheduled_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// ScheduledPlansForDashboardReader is a Reader for the ScheduledPlansForDashboard structure.
type ScheduledPlansForDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduledPlansForDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewScheduledPlansForDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewScheduledPlansForDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewScheduledPlansForDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewScheduledPlansForDashboardOK creates a ScheduledPlansForDashboardOK with default headers values
func NewScheduledPlansForDashboardOK() *ScheduledPlansForDashboardOK {
	return &ScheduledPlansForDashboardOK{}
}

/*ScheduledPlansForDashboardOK handles this case with default header values.

Scheduled Plan
*/
type ScheduledPlansForDashboardOK struct {
	Payload []*models.ScheduledPlan
}

func (o *ScheduledPlansForDashboardOK) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/dashboard/{dashboard_id}][%d] scheduledPlansForDashboardOK  %+v", 200, o.Payload)
}

func (o *ScheduledPlansForDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledPlansForDashboardBadRequest creates a ScheduledPlansForDashboardBadRequest with default headers values
func NewScheduledPlansForDashboardBadRequest() *ScheduledPlansForDashboardBadRequest {
	return &ScheduledPlansForDashboardBadRequest{}
}

/*ScheduledPlansForDashboardBadRequest handles this case with default header values.

Bad Request
*/
type ScheduledPlansForDashboardBadRequest struct {
	Payload *models.Error
}

func (o *ScheduledPlansForDashboardBadRequest) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/dashboard/{dashboard_id}][%d] scheduledPlansForDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *ScheduledPlansForDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduledPlansForDashboardNotFound creates a ScheduledPlansForDashboardNotFound with default headers values
func NewScheduledPlansForDashboardNotFound() *ScheduledPlansForDashboardNotFound {
	return &ScheduledPlansForDashboardNotFound{}
}

/*ScheduledPlansForDashboardNotFound handles this case with default header values.

Not Found
*/
type ScheduledPlansForDashboardNotFound struct {
	Payload *models.Error
}

func (o *ScheduledPlansForDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /scheduled_plans/dashboard/{dashboard_id}][%d] scheduledPlansForDashboardNotFound  %+v", 404, o.Payload)
}

func (o *ScheduledPlansForDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
