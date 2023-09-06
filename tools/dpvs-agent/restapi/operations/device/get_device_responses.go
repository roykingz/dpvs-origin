// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetDeviceOKCode is the HTTP code returned for type GetDeviceOK
const GetDeviceOKCode int = 200

/*
GetDeviceOK Success

swagger:response getDeviceOK
*/
type GetDeviceOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetDeviceOK creates GetDeviceOK with default headers values
func NewGetDeviceOK() *GetDeviceOK {

	return &GetDeviceOK{}
}

// WithPayload adds the payload to the get device o k response
func (o *GetDeviceOK) WithPayload(payload string) *GetDeviceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get device o k response
func (o *GetDeviceOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeviceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}