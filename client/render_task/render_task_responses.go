// Code generated by go-swagger; DO NOT EDIT.

package render_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/billtrust/looker-go-sdk/models"
)

// RenderTaskReader is a Reader for the RenderTask structure.
type RenderTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRenderTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewRenderTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenderTaskOK creates a RenderTaskOK with default headers values
func NewRenderTaskOK() *RenderTaskOK {
	return &RenderTaskOK{}
}

/*RenderTaskOK handles this case with default header values.

Render Task
*/
type RenderTaskOK struct {
	Payload *models.RenderTask
}

func (o *RenderTaskOK) Error() string {
	return fmt.Sprintf("[GET /render_tasks/{render_task_id}][%d] renderTaskOK  %+v", 200, o.Payload)
}

func (o *RenderTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RenderTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderTaskNotFound creates a RenderTaskNotFound with default headers values
func NewRenderTaskNotFound() *RenderTaskNotFound {
	return &RenderTaskNotFound{}
}

/*RenderTaskNotFound handles this case with default header values.

Not Found
*/
type RenderTaskNotFound struct {
	Payload *models.Error
}

func (o *RenderTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /render_tasks/{render_task_id}][%d] renderTaskNotFound  %+v", 404, o.Payload)
}

func (o *RenderTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
