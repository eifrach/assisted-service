// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ListBundlesOKCode is the HTTP code returned for type V2ListBundlesOK
const V2ListBundlesOKCode int = 200

/*
V2ListBundlesOK Success

swagger:response v2ListBundlesOK
*/
type V2ListBundlesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Bundle `json:"body,omitempty"`
}

// NewV2ListBundlesOK creates V2ListBundlesOK with default headers values
func NewV2ListBundlesOK() *V2ListBundlesOK {

	return &V2ListBundlesOK{}
}

// WithPayload adds the payload to the v2 list bundles o k response
func (o *V2ListBundlesOK) WithPayload(payload []*models.Bundle) *V2ListBundlesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list bundles o k response
func (o *V2ListBundlesOK) SetPayload(payload []*models.Bundle) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListBundlesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Bundle, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2ListBundlesInternalServerErrorCode is the HTTP code returned for type V2ListBundlesInternalServerError
const V2ListBundlesInternalServerErrorCode int = 500

/*
V2ListBundlesInternalServerError Internal server error

swagger:response v2ListBundlesInternalServerError
*/
type V2ListBundlesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListBundlesInternalServerError creates V2ListBundlesInternalServerError with default headers values
func NewV2ListBundlesInternalServerError() *V2ListBundlesInternalServerError {

	return &V2ListBundlesInternalServerError{}
}

// WithPayload adds the payload to the v2 list bundles internal server error response
func (o *V2ListBundlesInternalServerError) WithPayload(payload *models.Error) *V2ListBundlesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list bundles internal server error response
func (o *V2ListBundlesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListBundlesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
