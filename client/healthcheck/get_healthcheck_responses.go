// Code generated by go-swagger; DO NOT EDIT.

package healthcheck

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/myles-mcdonnell/sendmailserviceproxy/models"
)

// GetHealthcheckReader is a Reader for the GetHealthcheck structure.
type GetHealthcheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthcheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHealthcheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetHealthcheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHealthcheckOK creates a GetHealthcheckOK with default headers values
func NewGetHealthcheckOK() *GetHealthcheckOK {
	return &GetHealthcheckOK{}
}

/*GetHealthcheckOK handles this case with default header values.

confirm that the service is healthy
*/
type GetHealthcheckOK struct {
}

func (o *GetHealthcheckOK) Error() string {
	return fmt.Sprintf("[GET /healthcheck][%d] getHealthcheckOK ", 200)
}

func (o *GetHealthcheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthcheckDefault creates a GetHealthcheckDefault with default headers values
func NewGetHealthcheckDefault(code int) *GetHealthcheckDefault {
	return &GetHealthcheckDefault{
		_statusCode: code,
	}
}

/*GetHealthcheckDefault handles this case with default header values.

generic error response
*/
type GetHealthcheckDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// Code gets the status code for the get healthcheck default response
func (o *GetHealthcheckDefault) Code() int {
	return o._statusCode
}

func (o *GetHealthcheckDefault) Error() string {
	return fmt.Sprintf("[GET /healthcheck][%d] GetHealthcheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetHealthcheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
