package faxrequest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/grokify/ringcentral-sdk-go/rcsdk/models"
)

const (
	JSON_CONTENT_TYPE = "application/json"
)

type Metadata struct {
	To            []models.CallerInfo `json:"to,omitempty"`
	FaxResolution string              `json:"faxResolution,omitempty"`
	SendTime      string              `json:"sendTime,omitempty"`
	CoverIndex    int8                `json:"coverIndex,omitempty"`
	CoverPageText string              `json:"coverPageText,omitempty"`
}

type FaxRequest struct {
	url        string
	metadata   Metadata
	writer     *multipart.Writer
	headers    http.Header
	bodyBuffer *bytes.Buffer
	body       []byte
}

func NewFaxRequest(urlpath string, metadata Metadata) (FaxRequest, error) {
	fax := FaxRequest{url: urlpath}
	fax.url = urlpath
	fax.initialize()

	err := fax.SetMetadata(metadata)
	if err != nil {
		return fax, err
	}
	return fax, nil
}

func (fax *FaxRequest) initialize() {
	fax.bodyBuffer = &bytes.Buffer{}
	fax.writer = multipart.NewWriter(fax.bodyBuffer) // io.Writer
}

func (fax *FaxRequest) Finalize() error {
	err := fax.writer.Close()
	if err != nil {
		return err
	}
	fax.body = fax.bodyBuffer.Bytes()
	fax.setHeaders()
	return nil
}

func (fax *FaxRequest) SetMetadata(metadata Metadata) error {
	fax.metadata = metadata

	// MIME HEADER
	headers := textproto.MIMEHeader{}
	headers.Add("Content-Type", "application/json")

	// MIME PART
	part, err := fax.writer.CreatePart(headers)
	if err != nil {
		return err
	}
	j, err := json.Marshal(metadata)
	if err != nil {
		return err
	}
	_, err = part.Write(j)
	return err
}

func (fax *FaxRequest) AddFile(path string) error {

	// READ FILE
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// FILE HEAD
	headers := textproto.MIMEHeader{}
	headers.Add("Content-Type", "application/octet-stream")

	_, filename := filepath.Split(path)
	if len(filename) > 0 {
		headers.Add("Content-Disposition", strings.Join([]string{`attachment; filename="`, filename, `"`}, ""))
	} else {
		headers.Add("Content-Disposition", "attachment")
	}

	// FILE PART
	part, err := fax.writer.CreatePart(headers)
	if err != nil {
		return err
	}
	_, err = part.Write(bytes)
	return err
}

func (fax *FaxRequest) AddText(bytes []byte, content_type string) error {

	if len(content_type) == 0 {
		content_type = "text/plain"
	}

	// FILE HEAD
	headers := textproto.MIMEHeader{}
	headers.Add("Content-Type", content_type)

	// FILE PART
	part, err := fax.writer.CreatePart(headers)
	if err != nil {
		return err
	}
	_, err = part.Write(bytes)
	return err
}

func (fax *FaxRequest) setHeaders() {
	fax.headers = http.Header{}
	fax.headers.Add("Content-Type", strings.Join([]string{"multipart/mixed; boundary=", fax.writer.Boundary()}, ""))
}

func (fax *FaxRequest) Method() string {
	return "POST"
}

func (fax *FaxRequest) Path() string {
	return ""
}

func (fax *FaxRequest) Url() string {
	return fax.url
}

func (fax *FaxRequest) SetUrl(url string) {
	fax.url = url
}

func (fax *FaxRequest) Query() url.Values {
	return url.Values{}
}

func (fax *FaxRequest) Headers() http.Header {
	return fax.headers
}

func (fax *FaxRequest) Body() io.Reader {
	return bytes.NewReader(fax.body)
}

func (req *FaxRequest) Send() (*http.Response, error) {
	// REQUEST
	r, _ := http.NewRequest(req.Method(), req.Url(), req.Body())

	r.Header = req.Headers()
	r.Header.Add("Accept", JSON_CONTENT_TYPE)

	// RESPONSE
	client := &http.Client{}
	return client.Do(r)
}
