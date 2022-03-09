// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetHostHandlerFunc turns a function with the right signature into a get host handler
type GetHostHandlerFunc func(GetHostParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHostHandlerFunc) Handle(params GetHostParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetHostHandler interface for that can handle valid get host params
type GetHostHandler interface {
	Handle(GetHostParams, interface{}) middleware.Responder
}

// NewGetHost creates a new http.Handler for the get host operation
func NewGetHost(ctx *middleware.Context, handler GetHostHandler) *GetHost {
	return &GetHost{Context: ctx, Handler: handler}
}

/* GetHost swagger:route GET /v1/clusters/{cluster_id}/hosts/{host_id} installer getHost

Retrieves the details of the OpenShift host.

*/
type GetHost struct {
	Context *middleware.Context
	Handler GetHostHandler
}

func (o *GetHost) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHostParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}