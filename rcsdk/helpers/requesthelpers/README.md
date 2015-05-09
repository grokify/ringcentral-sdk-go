# Table of contents

1. [Overview](#overview)
2. [Fax](#fax)

***

# Overview

Request helpers are provided to generate HTTP request parameters for API calls including HTTP bodies and HTTP headers.

# Fax

The fax file helper takes a metadata byte array along with a file path and creates the requisite HTTP request body and headers.

```go
import(
	"github.com/grokify/ringcentral-sdk-go/rcsdk"
	"github.com/grokify/ringcentral-sdk-go/rcsdk/helpers/requesthelpers"
)

sdk := rcsdk.NewSdk("yourAppKey", "yourAppSecret", "https://platform.devtest.ringcentral.com")

platform := sdk.GetPlatform();

platform.Authorize('+16505551212','101','yourPassword')

fax := requesthelpers.NewReqHelperFaxFile([]byte(`{ 
	"to" : [{"phoneNumber": "16505551212"}],
	"faxResolution" : "High"
}`), "/path/to/myfile.pdf")

resp, err := platform.Post("/account/~/extension/~/fax", url.Values{}, fax.GetBody(), fax.GetHeaders())
```

## Examples

The following examples use the test files included in this repository, including `test_file.pdf`.

### Example 1: PDF

This example sends the included `test_file.pdf` as `Content-Type: application/octet-stream`. This content type is recommended for production use as it is smaller than the alternative base64 encoding.

#### Request

##### Go Request

```go
fax := requesthelpers.NewReqHelperFaxFile([]byte(`{ 
	"to" : [{"phoneNumber": "16505551212"}],
	"faxResolution" : "High"
}`), "/path/to/test_file.pdf")

resp, err := platform.Post("/account/~/extension/~/fax", url.Values{}, fax.GetBody(), fax.GetHeaders())
```

##### HTTP Request

```http
POST https://platform.ringcentral.com/restapi/v1.0/account/~/extension/~/fax HTTP/1.1
Authorization: Bearer U0pDMDFQMDFQQVMwMnxBQUFWZmY4ZXoxMlh
Accept: application/json
Content-Type: multipart/mixed; boundary=39aa294acba996517c87259618d1153df4ef812cfa2f2e627e7865844fcf

--39aa294acba996517c87259618d1153df4ef812cfa2f2e627e7865844fcf
Content-Type: application/json

{ 
		"to" : [{"phoneNumber": "16505551212"}],
		"faxResolution" : "High"
	}
--39aa294acba996517c87259618d1153df4ef812cfa2f2e627e7865844fcf
Content-Type: application/octet-stream
Content-Disposition: attachment; filename="test_file.pdf"

%????1.2
0000000016 00000 n                                              xref
0000000644 00000 n
0000000917 00000 n
0000001068 00000 n
0000001224 00000 n
0000001410 00000 n
0000001589 00000 n
0000001768 00000 n
0000002197 00000 n
0000002383 00000 n
0000002769 00000 n
0000003172 00000 n
0000003351 00000 n
0000000760 00000 n
0000000897 00000 n
stream36 /Filter /FlateDecode /Length 17 0 R >> a4fe8bd49a9dd0599b0b151>]
streamngth 355 /Filter /FlateDecode >> .71519 0.1192 0.1805 0.0722 0.9505 ] >> 
H?t??N?0D???=ڈn?q?$G?????7?!MR
?	")??gm?"UB???z<??s/??g@?7????,+?ژ
c#??ϣ????|              ʿ?\?K            ??b?M~'???f?6
                             ~5O?*??r.??5k?j?ܦ
                                              ?yA??EH?5Y??X,????j?J,?_gyc?\?{?~?????B9,$??7u?w?+P%Z??~????'??}5??????.?#?"Q"????"?F{h?^U?K?`
'?????`?=???{?֟i???"͢?q?MJL?u=)&%?	iyC???RK??ǔ??_?ŭ?F:~?7?G<sX3?xOq??y?????^0????
0000000000 65535 f0010829095457)ent)r Windows)0 0 222 833 556 556 0 0 0 0 0 
0000003429 00000 n
0000003658 00000 n
%%EOFxref46c5ba4fe8bd49a9dd0599b0b151><d70f46c5ba4fe8bd49a9dd0599b0b151>]
--39aa294acba996517c87259618d1153df4ef812cfa2f2e627e7865844fcf--
```

#### Response

```json
{
  "uri" : "https://platform.devtest.ringcentral.com/restapi/v1.0/account/111111111/extension/222222222/message-store/333333333",
  "id" : 4444444444,
  "to" : [ {
    "phoneNumber" : "+16505551212",
    "location" : "Redwood City, CA"
  } ],
  "type" : "Fax",
  "creationTime" : "2015-05-09T16:11:58.000Z",
  "readStatus" : "Unread",
  "priority" : "Normal",
  "attachments" : [ {
    "id" : 4444444444,
    "uri" : "https://platform.devtest.ringcentral.com/restapi/v1.0/account/111111111/extension/222222222/message-store/333333333/content/4444444444",
    "type" : "RenderedDocument",
    "contentType" : "image/tiff"
  } ],
  "direction" : "Outbound",
  "availability" : "Alive",
  "messageStatus" : "Queued",
  "faxResolution" : "High",
  "faxPageCount" : 0,
  "lastModifiedTime" : "2015-05-09T16:11:58.519Z"
}
```

#### Retrieval

The following code can be used to download and save a fax.

Note: this example uses `net/httputil` from `gotilla`.

```go
import (
	"net/http"
	"net/url"

	"github.com/grokify/gotilla/net/httputil"
)

fileUrl := "/account/111111111/extension/222222222/message-store/333333333/content/4444444444"

resp, _ := platform.Get(fileUrl, url.Values{}, []byte{}, http.Header{})
body, _ := httputil.ResponseBody(resp)

err := ioutil.WriteFile("/path/to/test_file_out.pdf", body, 0644)
```