// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRegisterUserParams creates a new RegisterUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterUserParams() *RegisterUserParams {
	return &RegisterUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterUserParamsWithTimeout creates a new RegisterUserParams object
// with the ability to set a timeout on a request.
func NewRegisterUserParamsWithTimeout(timeout time.Duration) *RegisterUserParams {
	return &RegisterUserParams{
		timeout: timeout,
	}
}

// NewRegisterUserParamsWithContext creates a new RegisterUserParams object
// with the ability to set a context for a request.
func NewRegisterUserParamsWithContext(ctx context.Context) *RegisterUserParams {
	return &RegisterUserParams{
		Context: ctx,
	}
}

// NewRegisterUserParamsWithHTTPClient creates a new RegisterUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterUserParamsWithHTTPClient(client *http.Client) *RegisterUserParams {
	return &RegisterUserParams{
		HTTPClient: client,
	}
}

/*
RegisterUserParams contains all the parameters to send to the API endpoint

	for the register user operation.

	Typically these are written to a http.Request.
*/
type RegisterUserParams struct {

	// Body.
	Body RegisterUserBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterUserParams) WithDefaults() *RegisterUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register user params
func (o *RegisterUserParams) WithTimeout(timeout time.Duration) *RegisterUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register user params
func (o *RegisterUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register user params
func (o *RegisterUserParams) WithContext(ctx context.Context) *RegisterUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register user params
func (o *RegisterUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register user params
func (o *RegisterUserParams) WithHTTPClient(client *http.Client) *RegisterUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register user params
func (o *RegisterUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register user params
func (o *RegisterUserParams) WithBody(body RegisterUserBody) *RegisterUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register user params
func (o *RegisterUserParams) SetBody(body RegisterUserBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}