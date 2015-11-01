package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// LoginUserHandlerFunc turns a function with the right signature into a login user handler
type LoginUserHandlerFunc func(LoginUserParams) (*string, error)

func (fn LoginUserHandlerFunc) Handle(params LoginUserParams) (*string, error) {
	return fn(params)
}

// LoginUserHandler interface for that can handle valid login user params
type LoginUserHandler interface {
	Handle(LoginUserParams) (*string, error)
}

// NewLoginUser creates a new http.Handler for the login user operation
func NewLoginUser(ctx *middleware.Context, handler LoginUserHandler) *LoginUser {
	return &LoginUser{Context: ctx, Handler: handler}
}

/*
Logs user into the system
*/
type LoginUser struct {
	Context *middleware.Context
	Params  LoginUserParams
	Handler LoginUserHandler
}

func (o *LoginUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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
