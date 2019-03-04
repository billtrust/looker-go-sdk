// Code generated by go-swagger; DO NOT EDIT.

package content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bmccarthy/looker-go-sdk/models"
)

// DeleteContentFavoriteReader is a Reader for the DeleteContentFavorite structure.
type DeleteContentFavoriteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentFavoriteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteContentFavoriteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteContentFavoriteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteContentFavoriteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteContentFavoriteNoContent creates a DeleteContentFavoriteNoContent with default headers values
func NewDeleteContentFavoriteNoContent() *DeleteContentFavoriteNoContent {
	return &DeleteContentFavoriteNoContent{}
}

/*DeleteContentFavoriteNoContent handles this case with default header values.

Successfully deleted.
*/
type DeleteContentFavoriteNoContent struct {
	Payload string
}

func (o *DeleteContentFavoriteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /content_favorite/{content_favorite_id}][%d] deleteContentFavoriteNoContent  %+v", 204, o.Payload)
}

func (o *DeleteContentFavoriteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentFavoriteBadRequest creates a DeleteContentFavoriteBadRequest with default headers values
func NewDeleteContentFavoriteBadRequest() *DeleteContentFavoriteBadRequest {
	return &DeleteContentFavoriteBadRequest{}
}

/*DeleteContentFavoriteBadRequest handles this case with default header values.

Bad Request
*/
type DeleteContentFavoriteBadRequest struct {
	Payload *models.Error
}

func (o *DeleteContentFavoriteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /content_favorite/{content_favorite_id}][%d] deleteContentFavoriteBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContentFavoriteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentFavoriteNotFound creates a DeleteContentFavoriteNotFound with default headers values
func NewDeleteContentFavoriteNotFound() *DeleteContentFavoriteNotFound {
	return &DeleteContentFavoriteNotFound{}
}

/*DeleteContentFavoriteNotFound handles this case with default header values.

Not Found
*/
type DeleteContentFavoriteNotFound struct {
	Payload *models.Error
}

func (o *DeleteContentFavoriteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /content_favorite/{content_favorite_id}][%d] deleteContentFavoriteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteContentFavoriteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
