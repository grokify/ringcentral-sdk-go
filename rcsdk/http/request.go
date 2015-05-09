package http

import (
	"bytes"
	"io"
	corehttp "net/http"
	"net/url"
	"strings"
)

type Request struct {
	Headers
	method      string
	url         string
	queryParams url.Values
	body        []byte

	response *corehttp.Response
}

func NewRequest(method string, url string, queryParams url.Values, body []byte, headers corehttp.Header) Request {
	req := Request{method: method, url: url, queryParams: queryParams, body: body}
	req.headers = headers
	return req
}

func (req *Request) GetUrl() string {
	return req.url
}

func (req *Request) GetUrlWithQueryString() string {
	url := req.GetUrl()
	if len(req.queryParams) > 0 {
		queryString := req.queryParams.Encode()
		url = strings.Join([]string{url, queryString}, "?")
	}
	return url
}

func (req *Request) SetUrl(url string) {
	req.url = url
}

func (req *Request) GetBody() io.Reader {
	return bytes.NewReader(req.body)
}

func (req *Request) Send() (*corehttp.Response, error) {

	// REQUEST
	corereq, _ := corehttp.NewRequest(req.method, req.url, req.GetBody())

	corereq.Header = req.headers
	corereq.Header.Add("Accept", JSON_CONTENT_TYPE)
	corereq.Header.Add("Content-Type", JSON_CONTENT_TYPE)

	// RESPONSE
	client := &corehttp.Client{}
	return client.Do(corereq)
}
