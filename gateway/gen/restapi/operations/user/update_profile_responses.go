// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateProfileOKCode is the HTTP code returned for type UpdateProfileOK
const UpdateProfileOKCode int = 200

/*
UpdateProfileOK create user

swagger:response updateProfileOK
*/
type UpdateProfileOK struct {

	/*
	  In: Body
	*/
	Payload *UpdateProfileOKBody `json:"body,omitempty"`
}

// NewUpdateProfileOK creates UpdateProfileOK with default headers values
func NewUpdateProfileOK() *UpdateProfileOK {

	return &UpdateProfileOK{}
}

// WithPayload adds the payload to the update profile o k response
func (o *UpdateProfileOK) WithPayload(payload *UpdateProfileOKBody) *UpdateProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update profile o k response
func (o *UpdateProfileOK) SetPayload(payload *UpdateProfileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
