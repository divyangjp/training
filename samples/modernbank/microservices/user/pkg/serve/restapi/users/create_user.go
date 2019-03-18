// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	model "github.com/tetrateio/training/samples/modernbank/microservices/user/pkg/model"
)

// CreateUserHandlerFunc turns a function with the right signature into a create user handler
type CreateUserHandlerFunc func(CreateUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateUserHandlerFunc) Handle(params CreateUserParams) middleware.Responder {
	return fn(params)
}

// CreateUserHandler interface for that can handle valid create user params
type CreateUserHandler interface {
	Handle(CreateUserParams) middleware.Responder
}

// NewCreateUser creates a new http.Handler for the create user operation
func NewCreateUser(ctx *middleware.Context, handler CreateUserHandler) *CreateUser {
	return &CreateUser{Context: ctx, Handler: handler}
}

/*CreateUser swagger:route POST /users users createUser

Create a new user identity for a customer

Creates a new user

*/
type CreateUser struct {
	Context *middleware.Context
	Handler CreateUserHandler
}

func (o *CreateUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateUserCreatedBody create user created body
// swagger:model CreateUserCreatedBody
type CreateUserCreatedBody struct {
	model.User

	model.Version
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CreateUserCreatedBody) UnmarshalJSON(raw []byte) error {
	// CreateUserCreatedBodyAO0
	var createUserCreatedBodyAO0 model.User
	if err := swag.ReadJSON(raw, &createUserCreatedBodyAO0); err != nil {
		return err
	}
	o.User = createUserCreatedBodyAO0

	// CreateUserCreatedBodyAO1
	var createUserCreatedBodyAO1 model.Version
	if err := swag.ReadJSON(raw, &createUserCreatedBodyAO1); err != nil {
		return err
	}
	o.Version = createUserCreatedBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CreateUserCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	createUserCreatedBodyAO0, err := swag.WriteJSON(o.User)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createUserCreatedBodyAO0)

	createUserCreatedBodyAO1, err := swag.WriteJSON(o.Version)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createUserCreatedBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create user created body
func (o *CreateUserCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with model.User
	if err := o.User.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with model.Version
	if err := o.Version.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *CreateUserCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateUserCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateUserCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
