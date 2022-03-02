// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// GetHostReader is a Reader for the GetHost structure.
type GetHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetHostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostOK creates a GetHostOK with default headers values
func NewGetHostOK() *GetHostOK {
	return &GetHostOK{}
}

/* GetHostOK describes a response with status code 200, with default header values.

Success.
*/
type GetHostOK struct {
	Payload *models.Host
}

func (o *GetHostOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/hosts/{host_id}][%d] getHostOK  %+v", 200, o.Payload)
}
func (o *GetHostOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *GetHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostUnauthorized creates a GetHostUnauthorized with default headers values
func NewGetHostUnauthorized() *GetHostUnauthorized {
	return &GetHostUnauthorized{}
}

/* GetHostUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetHostUnauthorized struct {
	Payload *models.InfraError
}

func (o *GetHostUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/hosts/{host_id}][%d] getHostUnauthorized  %+v", 401, o.Payload)
}
func (o *GetHostUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostForbidden creates a GetHostForbidden with default headers values
func NewGetHostForbidden() *GetHostForbidden {
	return &GetHostForbidden{}
}

/* GetHostForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type GetHostForbidden struct {
	Payload *models.InfraError
}

func (o *GetHostForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/hosts/{host_id}][%d] getHostForbidden  %+v", 403, o.Payload)
}
func (o *GetHostForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *GetHostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostNotFound creates a GetHostNotFound with default headers values
func NewGetHostNotFound() *GetHostNotFound {
	return &GetHostNotFound{}
}

/* GetHostNotFound describes a response with status code 404, with default header values.

Error.
*/
type GetHostNotFound struct {
	Payload *models.Error
}

func (o *GetHostNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/hosts/{host_id}][%d] getHostNotFound  %+v", 404, o.Payload)
}
func (o *GetHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostMethodNotAllowed creates a GetHostMethodNotAllowed with default headers values
func NewGetHostMethodNotAllowed() *GetHostMethodNotAllowed {
	return &GetHostMethodNotAllowed{}
}

/* GetHostMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type GetHostMethodNotAllowed struct {
	Payload *models.Error
}

func (o *GetHostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/hosts/{host_id}][%d] getHostMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *GetHostMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostInternalServerError creates a GetHostInternalServerError with default headers values
func NewGetHostInternalServerError() *GetHostInternalServerError {
	return &GetHostInternalServerError{}
}

/* GetHostInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type GetHostInternalServerError struct {
	Payload *models.Error
}

func (o *GetHostInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters/{cluster_id}/hosts/{host_id}][%d] getHostInternalServerError  %+v", 500, o.Payload)
}
func (o *GetHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}