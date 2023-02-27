// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetAPITaskIDSolveParams creates a new GetAPITaskIDSolveParams object
//
// There are no default values defined in the spec.
func NewGetAPITaskIDSolveParams() GetAPITaskIDSolveParams {

	return GetAPITaskIDSolveParams{}
}

// GetAPITaskIDSolveParams contains all the bound params for the get API task ID solve operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAPITaskIDSolve
type GetAPITaskIDSolveParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ID uint64
	/*
	  Required: true
	  In: query
	*/
	SolutionID uint64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAPITaskIDSolveParams() beforehand.
func (o *GetAPITaskIDSolveParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSolutionID, qhkSolutionID, _ := qs.GetOK("solutionID")
	if err := o.bindSolutionID(qSolutionID, qhkSolutionID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetAPITaskIDSolveParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "uint64", raw)
	}
	o.ID = value

	return nil
}

// bindSolutionID binds and validates parameter SolutionID from query.
func (o *GetAPITaskIDSolveParams) bindSolutionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("solutionID", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("solutionID", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertUint64(raw)
	if err != nil {
		return errors.InvalidType("solutionID", "query", "uint64", raw)
	}
	o.SolutionID = value

	return nil
}
