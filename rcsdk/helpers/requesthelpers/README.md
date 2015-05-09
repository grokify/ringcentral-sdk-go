# Table of contents

1. [Overview](#overview)
4. [Fax](#fax)

***

# Overview

Request helpers are provided to generate HTTP request parameters for API calls including HTTP bodies and HTTP headers.

# Fax

The fax file helper takes a metadata byte array along with a file path and creates the requisite HTTP request body and headers.

```go
import(
	"github.com/grokify/ringcentral-sdk-go/rcsdk/helpers/requesthelpers"
)

fax := requesthelpers.NewReqHelperFaxFile([]byte(`{ 
	"to" : [{"phoneNumber": "16505551212"}],
	"faxResolution" : "High"
}`), "/path/to/myfile.pdf")

resp, _ := platform.Post("/account/~/extension/~/fax", url.Values{}, fax.GetBody(), fax.GetHeaders())
```