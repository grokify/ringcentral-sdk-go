package core

import (
	"io"
	"net/http"
	"net/url"

	rchttp "github.com/grokify/ringcentral-sdk-go/rcsdk/http"
)

type Context struct {
}

func NewContext() Context {
	c := Context{}
	return c
}

func (c *Context) GetRequest(method string, url string, queryParams url.Values, body io.Reader, headers http.Header) rchttp.Request {
	//rcreq := rchttp.NewRequest(method, url, queryParams, body, headers)
	rcreq := rchttp.BaseRequest{}
	rcreq.SetUrl(url)
	return &rcreq
}
