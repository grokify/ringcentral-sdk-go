package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/grokify/mogo/net/http/httputilmore"
)

/*
type RequestInterface interface{}

type Request2 struct {
	Method  string
	Path    string
	Query   url.Values
	URL     string
	Body    io.Reader
	Headers http.Header
}

func NewRequest2() Request2 {
	return Request2{
		Method:  http.MethodGet,
		Query:   url.Values{},
		Body:    bytes.NewReader([]byte("")),
		Headers: http.Header{}}
}

func (req *Request2) Send() (*http.Response, error) {
	corereq, err := http.NewRequest(strings.ToUpper(req.Method), req.URL, req.Body)
	if err != nil {
		return nil, err
	}

	corereq.Header = req.Headers
	corereq.Header.Add(httputilmore.HeaderAccept, httputilmore.ContentTypeAppJSONUtf8)

	client := &http.Client{}
	return client.Do(corereq)
}
*/

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
	corereq.Header.Add(httputilmore.HeaderAccept, httputilmore.ContentTypeAppJSONUtf8)

	// RESPONSE
	client := &http.Client{}
	return client.Do(corereq)
}
