// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateProfileHandlerFunc turns a function with the right signature into a update profile handler
type UpdateProfileHandlerFunc func(UpdateProfileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateProfileHandlerFunc) Handle(params UpdateProfileParams) middleware.Responder {
	return fn(params)
}

// UpdateProfileHandler interface for that can handle valid update profile params
type UpdateProfileHandler interface {
	Handle(UpdateProfileParams) middleware.Responder
}

// NewUpdateProfile creates a new http.Handler for the update profile operation
func NewUpdateProfile(ctx *middleware.Context, handler UpdateProfileHandler) *UpdateProfile {
	return &UpdateProfile{Context: ctx, Handler: handler}
}

/*
	UpdateProfile swagger:route PUT /profile user updateProfile

UpdateProfile update profile API
*/
type UpdateProfile struct {
	Context *middleware.Context
	Handler UpdateProfileHandler
}

func (o *UpdateProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateProfileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateProfileBody update profile body
//
// swagger:model UpdateProfileBody
type UpdateProfileBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// phone number
	PhoneNumber string `json:"phone_number,omitempty"`
}

// Validate validates this update profile body
func (o *UpdateProfileBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update profile body based on context it is used
func (o *UpdateProfileBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateProfileBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateProfileBody) UnmarshalBinary(b []byte) error {
	var res UpdateProfileBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateProfileOKBody update profile o k body
//
// swagger:model UpdateProfileOKBody
type UpdateProfileOKBody struct {

	// code
	Code float64 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update profile o k body
func (o *UpdateProfileOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update profile o k body based on context it is used
func (o *UpdateProfileOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateProfileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateProfileOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateProfileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
