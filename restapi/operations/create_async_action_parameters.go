// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/mcastelino/testapi/models"
)

// NewCreateAsyncActionParams creates a new CreateAsyncActionParams object
// no default values defined in spec.
func NewCreateAsyncActionParams() CreateAsyncActionParams {

	return CreateAsyncActionParams{}
}

// CreateAsyncActionParams contains all the bound params for the create async action operation
// typically these are obtained from a http.Request
//
// swagger:parameters createAsyncAction
type CreateAsyncActionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Info *models.ActionInfo
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateAsyncActionParams() beforehand.
func (o *CreateAsyncActionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ActionInfo
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("info", "body"))
			} else {
				res = append(res, errors.NewParseError("info", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Info = &body
			}
		}
	} else {
		res = append(res, errors.Required("info", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}