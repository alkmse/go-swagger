package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/fixtures/goparsing/petstore/models"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetOrderByIDHandlerFunc turns a function with the right signature into a get order by id handler
type GetOrderByIDHandlerFunc func(GetOrderByIDParams) (*models.Order, error)

func (fn GetOrderByIDHandlerFunc) Handle(params GetOrderByIDParams) (*models.Order, error) {
	return fn(params)
}

// GetOrderByIDHandler interface for that can handle valid get order by id params
type GetOrderByIDHandler interface {
	Handle(GetOrderByIDParams) (*models.Order, error)
}

// NewGetOrderByID creates a new http.Handler for the get order by id operation
func NewGetOrderByID(ctx *middleware.Context, handler GetOrderByIDHandler) *GetOrderByID {
	return &GetOrderByID{Context: ctx, Handler: handler}
}

/*
Find purchase order by ID

For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
*/
type GetOrderByID struct {
	Context *middleware.Context
	Params  GetOrderByIDParams
	Handler GetOrderByIDHandler
}

func (o *GetOrderByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res, err := o.Handler.Handle(o.Params) // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	o.Context.Respond(rw, r, route.Produces, route, res)

}
