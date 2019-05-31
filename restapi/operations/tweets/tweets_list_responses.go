// Code generated by go-swagger; DO NOT EDIT.

package tweets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/suda/go-swagger-example/models"
)

// TweetsListOKCode is the HTTP code returned for type TweetsListOK
const TweetsListOKCode int = 200

/*TweetsListOK OK

swagger:response tweetsListOK
*/
type TweetsListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Tweet `json:"body,omitempty"`
}

// NewTweetsListOK creates TweetsListOK with default headers values
func NewTweetsListOK() *TweetsListOK {

	return &TweetsListOK{}
}

// WithPayload adds the payload to the tweets list o k response
func (o *TweetsListOK) WithPayload(payload []*models.Tweet) *TweetsListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tweets list o k response
func (o *TweetsListOK) SetPayload(payload []*models.Tweet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TweetsListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Tweet, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}