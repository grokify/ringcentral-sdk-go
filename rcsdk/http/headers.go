package http

import (
	corehttp "net/http"
)

const (
	CONTENT_TYPE             string = "Content-Type"
	AUTHORIZATION            string = "Authorization"
	URL_ENCODED_CONTENT_TYPE string = "application/x-www-form-urlencoded"
	JSON_CONTENT_TYPE        string = "application/json"
	MULTIPART_CONTENT_TYPE   string = "multipart/mixed"
)

type Headers struct {
	headers corehttp.Header
}

func (h *Headers) SetHeader(name string, value string) {
	h.headers.Set(name, value)
}

func (h *Headers) GetHeaders() corehttp.Header {
	return h.headers
}
