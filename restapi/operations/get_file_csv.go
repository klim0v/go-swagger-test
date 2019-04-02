// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFileCsvHandlerFunc turns a function with the right signature into a get file csv handler
type GetFileCsvHandlerFunc func(GetFileCsvParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFileCsvHandlerFunc) Handle(params GetFileCsvParams) middleware.Responder {
	return fn(params)
}

// GetFileCsvHandler interface for that can handle valid get file csv params
type GetFileCsvHandler interface {
	Handle(GetFileCsvParams) middleware.Responder
}

// NewGetFileCsv creates a new http.Handler for the get file csv operation
func NewGetFileCsv(ctx *middleware.Context, handler GetFileCsvHandler) *GetFileCsv {
	return &GetFileCsv{Context: ctx, Handler: handler}
}

/*GetFileCsv swagger:route GET /file.csv getFileCsv

GetFileCsv get file csv API

*/
type GetFileCsv struct {
	Context *middleware.Context
	Handler GetFileCsvHandler
}

func (o *GetFileCsv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFileCsvParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
