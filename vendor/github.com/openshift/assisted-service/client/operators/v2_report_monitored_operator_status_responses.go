// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ReportMonitoredOperatorStatusReader is a Reader for the V2ReportMonitoredOperatorStatus structure.
type V2ReportMonitoredOperatorStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ReportMonitoredOperatorStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ReportMonitoredOperatorStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2ReportMonitoredOperatorStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2ReportMonitoredOperatorStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ReportMonitoredOperatorStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2ReportMonitoredOperatorStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2ReportMonitoredOperatorStatusMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2ReportMonitoredOperatorStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ReportMonitoredOperatorStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2ReportMonitoredOperatorStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ReportMonitoredOperatorStatusOK creates a V2ReportMonitoredOperatorStatusOK with default headers values
func NewV2ReportMonitoredOperatorStatusOK() *V2ReportMonitoredOperatorStatusOK {
	return &V2ReportMonitoredOperatorStatusOK{}
}

/*
V2ReportMonitoredOperatorStatusOK describes a response with status code 200, with default header values.

Success.
*/
type V2ReportMonitoredOperatorStatusOK struct {
}

// IsSuccess returns true when this v2 report monitored operator status o k response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 report monitored operator status o k response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status o k response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 report monitored operator status o k response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status o k response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2ReportMonitoredOperatorStatusOK) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusOK ", 200)
}

func (o *V2ReportMonitoredOperatorStatusOK) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusOK ", 200)
}

func (o *V2ReportMonitoredOperatorStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV2ReportMonitoredOperatorStatusBadRequest creates a V2ReportMonitoredOperatorStatusBadRequest with default headers values
func NewV2ReportMonitoredOperatorStatusBadRequest() *V2ReportMonitoredOperatorStatusBadRequest {
	return &V2ReportMonitoredOperatorStatusBadRequest{}
}

/*
V2ReportMonitoredOperatorStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V2ReportMonitoredOperatorStatusBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 report monitored operator status bad request response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status bad request response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status bad request response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 report monitored operator status bad request response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status bad request response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *V2ReportMonitoredOperatorStatusBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusBadRequest  %+v", 400, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusBadRequest) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusBadRequest  %+v", 400, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusUnauthorized creates a V2ReportMonitoredOperatorStatusUnauthorized with default headers values
func NewV2ReportMonitoredOperatorStatusUnauthorized() *V2ReportMonitoredOperatorStatusUnauthorized {
	return &V2ReportMonitoredOperatorStatusUnauthorized{}
}

/*
V2ReportMonitoredOperatorStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2ReportMonitoredOperatorStatusUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 report monitored operator status unauthorized response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status unauthorized response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status unauthorized response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 report monitored operator status unauthorized response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status unauthorized response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2ReportMonitoredOperatorStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusForbidden creates a V2ReportMonitoredOperatorStatusForbidden with default headers values
func NewV2ReportMonitoredOperatorStatusForbidden() *V2ReportMonitoredOperatorStatusForbidden {
	return &V2ReportMonitoredOperatorStatusForbidden{}
}

/*
V2ReportMonitoredOperatorStatusForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2ReportMonitoredOperatorStatusForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 report monitored operator status forbidden response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status forbidden response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status forbidden response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 report monitored operator status forbidden response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status forbidden response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2ReportMonitoredOperatorStatusForbidden) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusForbidden  %+v", 403, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusForbidden) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusForbidden  %+v", 403, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusNotFound creates a V2ReportMonitoredOperatorStatusNotFound with default headers values
func NewV2ReportMonitoredOperatorStatusNotFound() *V2ReportMonitoredOperatorStatusNotFound {
	return &V2ReportMonitoredOperatorStatusNotFound{}
}

/*
V2ReportMonitoredOperatorStatusNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2ReportMonitoredOperatorStatusNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 report monitored operator status not found response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status not found response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status not found response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 report monitored operator status not found response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status not found response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2ReportMonitoredOperatorStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusNotFound  %+v", 404, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusNotFound) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusNotFound  %+v", 404, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusMethodNotAllowed creates a V2ReportMonitoredOperatorStatusMethodNotAllowed with default headers values
func NewV2ReportMonitoredOperatorStatusMethodNotAllowed() *V2ReportMonitoredOperatorStatusMethodNotAllowed {
	return &V2ReportMonitoredOperatorStatusMethodNotAllowed{}
}

/*
V2ReportMonitoredOperatorStatusMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2ReportMonitoredOperatorStatusMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 report monitored operator status method not allowed response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status method not allowed response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status method not allowed response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 report monitored operator status method not allowed response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status method not allowed response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusConflict creates a V2ReportMonitoredOperatorStatusConflict with default headers values
func NewV2ReportMonitoredOperatorStatusConflict() *V2ReportMonitoredOperatorStatusConflict {
	return &V2ReportMonitoredOperatorStatusConflict{}
}

/*
V2ReportMonitoredOperatorStatusConflict describes a response with status code 409, with default header values.

Error.
*/
type V2ReportMonitoredOperatorStatusConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 report monitored operator status conflict response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status conflict response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status conflict response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 report monitored operator status conflict response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 report monitored operator status conflict response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusConflict) IsCode(code int) bool {
	return code == 409
}

func (o *V2ReportMonitoredOperatorStatusConflict) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusConflict  %+v", 409, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusConflict) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusConflict  %+v", 409, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusInternalServerError creates a V2ReportMonitoredOperatorStatusInternalServerError with default headers values
func NewV2ReportMonitoredOperatorStatusInternalServerError() *V2ReportMonitoredOperatorStatusInternalServerError {
	return &V2ReportMonitoredOperatorStatusInternalServerError{}
}

/*
V2ReportMonitoredOperatorStatusInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ReportMonitoredOperatorStatusInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 report monitored operator status internal server error response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status internal server error response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status internal server error response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 report monitored operator status internal server error response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 report monitored operator status internal server error response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2ReportMonitoredOperatorStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ReportMonitoredOperatorStatusServiceUnavailable creates a V2ReportMonitoredOperatorStatusServiceUnavailable with default headers values
func NewV2ReportMonitoredOperatorStatusServiceUnavailable() *V2ReportMonitoredOperatorStatusServiceUnavailable {
	return &V2ReportMonitoredOperatorStatusServiceUnavailable{}
}

/*
V2ReportMonitoredOperatorStatusServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type V2ReportMonitoredOperatorStatusServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 report monitored operator status service unavailable response has a 2xx status code
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 report monitored operator status service unavailable response has a 3xx status code
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 report monitored operator status service unavailable response has a 4xx status code
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 report monitored operator status service unavailable response has a 5xx status code
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 report monitored operator status service unavailable response a status code equal to that given
func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/monitored-operators][%d] v2ReportMonitoredOperatorStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ReportMonitoredOperatorStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}