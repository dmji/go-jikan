package components

import (
	"net/http"
)

type HTTPMetadata struct {
	// Raw HTTP response; suitable for custom response parsing
	Response *http.Response `json:"-"`
	// Raw HTTP request; suitable for debugging
	Request *http.Request `json:"-"`
}

func (o *HTTPMetadata) GetResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.Response
}

func (o *HTTPMetadata) GetRequest() *http.Request {
	if o == nil {
		return nil
	}
	return o.Request
}
