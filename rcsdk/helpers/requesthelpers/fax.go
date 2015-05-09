package requesthelpers

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"
)

type ReqHelperFaxFile struct {
	metadata []byte
	filename string
	writer   *multipart.Writer
	headers  http.Header
	body     []byte
}

func NewReqHelperFaxFile(metadata []byte, filename string) ReqHelperFaxFile {
	fax := ReqHelperFaxFile{}
	fax.metadata = metadata
	fax.filename = filename
	fax.build()
	return fax
}

func (fax *ReqHelperFaxFile) build() {
	body := &bytes.Buffer{}
	fax.writer = multipart.NewWriter(body)
	fax.SetMetadata(fax.metadata)
	fax.SetFile(fax.filename)
	fax.writer.Close()
	fax.body = body.Bytes()
	fax.setHeaders()
}

func (fax *ReqHelperFaxFile) SetMetadata(metadata []byte) {
	// MIME HEADER
	metaHead := textproto.MIMEHeader{}
	metaHead.Add("Content-Type", "application/json")

	// MIME PART
	metaPart, _ := fax.writer.CreatePart(metaHead)
	metaPart.Write(metadata)
}

func (fax *ReqHelperFaxFile) SetFile(path string) error {
	// READ FILE
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// FILE HEAD
	_, filename := filepath.Split(path)

	fileHead := textproto.MIMEHeader{}
	fileHead.Add("Content-Type", "application/octet-stream")
	if len(filename) > 0 {
		fileHead.Add("Content-Disposition", strings.Join([]string{`attachment; filename="`, filename, `"`}, ""))
	} else {
		fileHead.Add("Content-Disposition", "attachment")
	}

	// FILE PART
	filePart, _ := fax.writer.CreatePart(fileHead)
	filePart.Write(bytes)
	return nil
}

func (fax *ReqHelperFaxFile) setHeaders() {
	fax.headers = http.Header{}
	fax.headers.Add("Content-Type", strings.Join([]string{"multipart/mixed; boundary=", fax.writer.Boundary()}, ""))
}

func (fax *ReqHelperFaxFile) GetHeaders() http.Header {
	return fax.headers
}

func (fax *ReqHelperFaxFile) GetBody() []byte {
	return fax.body
}
