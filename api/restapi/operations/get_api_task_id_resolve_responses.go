// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"quizzer/api/models"
)

// GetAPITaskIDResolveOKCode is the HTTP code returned for type GetAPITaskIDResolveOK
const GetAPITaskIDResolveOKCode int = 200

/*
GetAPITaskIDResolveOK Correct answers

swagger:response getApiTaskIdResolveOK
*/
type GetAPITaskIDResolveOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Answer `json:"body,omitempty"`
}

// NewGetAPITaskIDResolveOK creates GetAPITaskIDResolveOK with default headers values
func NewGetAPITaskIDResolveOK() *GetAPITaskIDResolveOK {

	return &GetAPITaskIDResolveOK{}
}

// WithPayload adds the payload to the get Api task Id resolve o k response
func (o *GetAPITaskIDResolveOK) WithPayload(payload []*models.Answer) *GetAPITaskIDResolveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api task Id resolve o k response
func (o *GetAPITaskIDResolveOK) SetPayload(payload []*models.Answer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPITaskIDResolveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Answer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
