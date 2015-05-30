package faxrequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	obj "github.com/grokify/ringcentral-sdk-go/rcsdk/helpers/objects"
)

type Metadata struct {
	To            []obj.CallerInfo `json:"to,omitempty"`
	FaxResolution string           `json:"faxResolution,omitempty"`
	SendTime      string           `json:"sendTime,omitempty"`
	CoverIndex    int8             `json:"coverIndex,omitempty"`
	CoverPageText string           `json:"coverPageText,omitempty"`
}

type RequestHelper struct {
	metadata   Metadata
	writer     *multipart.Writer
	headers    http.Header
	bodyBuffer *bytes.Buffer
	body       []byte
}

func NewRequestHelper(metadata Metadata) (RequestHelper, error) {
	fax := RequestHelper{}
	fax.initialize()
	err := fax.SetMetadata(metadata)
	if err != nil {
		return fax, err
	}
	return fax, nil
}

func (fax *RequestHelper) initialize() {
	fax.bodyBuffer = &bytes.Buffer{}
	fax.writer = multipart.NewWriter(fax.bodyBuffer) // io.Writer
}

func (fax *RequestHelper) Finalize() error {
	err := fax.writer.Close()
	if err != nil {
		return err
	}
	fax.body = fax.bodyBuffer.Bytes()
	fax.setHeaders()
	return nil
}

func (fax *RequestHelper) SetMetadata(metadata Metadata) error {
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

func (fax *RequestHelper) AddFile(path string) error {

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

func (fax *RequestHelper) AddText(bytes []byte, content_type string) error {

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

func (fax *RequestHelper) setHeaders() {
	fax.headers = http.Header{}
	fax.headers.Add("Content-Type", strings.Join([]string{"multipart/mixed; boundary=", fax.writer.Boundary()}, ""))
}

func (fax *RequestHelper) GetHeaders() http.Header {
	return fax.headers
}

func (fax *RequestHelper) GetBody() []byte {
	return fax.body
}
