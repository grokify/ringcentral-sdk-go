package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request interface {
	Method() string
	Path() string
	Query() url.Values
	Url() string
	Body() io.Reader
	Headers() http.Header
	SetUrl(s string)
	Send() (*http.Response, error)
}

type BaseRequest struct {
	url string
}

func (req *BaseRequest) Method() string {
	return "get"
}

func (req *BaseRequest) Path() string {
	return ""
}

func (req *BaseRequest) Query() url.Values {
	return url.Values{}
}

func (req *BaseRequest) Url() string {
	return req.url
}

func (req *BaseRequest) SetUrl(url string) {
	fmt.Printf("GEN_SET_URL[%v]\n", url)
	req.url = url
}

func (req *BaseRequest) Body() io.Reader {
	return bytes.NewReader([]byte{})
}

func (req *BaseRequest) Headers() http.Header {
	return http.Header{}
}

func (req *BaseRequest) Send() (*http.Response, error) {
	// REQUEST
	corereq, _ := http.NewRequest(req.Method(), req.Url(), req.Body())

	corereq.Header = req.Headers()
	corereq.Header.Add("Accept", JSON_CONTENT_TYPE)

	// RESPONSE
	client := &http.Client{}
	return client.Do(corereq)
}

/*
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

	// RESPONSE
	client := &corehttp.Client{}
	return client.Do(corereq)
}
*/
