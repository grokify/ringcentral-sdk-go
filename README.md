RingCentral SDK in Go
=====================

[![Build Status](https://img.shields.io/travis/grokify/ringcentral-sdk-go/master.svg)](https://travis-ci.org/grokify/ringcentral-sdk-go)

## Table of contents

1. [Overview](#overview)
  1. [Included](#included)
  2. [To Do](#to-do)
2. [Installation](#installation)
2. [Usage](#usage)
  1. [Instantiation](#instantiation)
  1. [Authorization](#authorization)
  1. [API Requests](#api-requests)
    1. [SMS Example](#sms-example)
    2. [Fax Example](#fax-example)
6. [Links](#links)
7. [Contributions](#contributions)
8. [License](#license)

***

## Overview

This is an in-development SDK for the RingCentral for Developers platform. At present, it is used for testing purposes only. Many parts of this SDK are not yet implemented and are subject to breaking changes. Use at your own risk.

It attempts to mirror the structure of the official RingCentral SDKs. 

### Included

* OAuth2 authorization
* Generic API requests
* Fax request helper to create multipart/mixed messages

### To Do

The following items are still needed for this SDK. Contributions are most welcome.

* OAuth2 token refresh
* Subscriptions
* Tests

## Installation

```bash
$ go get github.com/grokify/ringcentral-sdk-go
```

## Usage

### Instantation

The SDK is represented by the global RCSDK constructor. Your application must create an instance of this object.

```go
import(
	"github.com/grokify/ringcentral-sdk-go/rcsdk"
)

// For Production use: rcsdk.RC_SERVER_PRODUCTION or "https://platform.ringcentral.com"
// For Sandbox use: rcsdk.RC_SERVER_SANDBOX or "https://platform.devtest.ringcentral.com"

sdk := rcsdk.NewSdk("yourAppKey", "yourAppSecret", rcsdk.RC_SERVER_SANDBOX)

// Get the Platform Singleton

platform := sdk.GetPlatform();
```

### Authorization

Login is accomplished by calling the `platform.Authorize()` method of the Platform singleton with username, extension
(optional), and password as parameters. A `Promise` instance is returned, resolved with an AJAX `Response` object.

The `username` should be a phone number in E.164 format without the leading `+`.

```go
platform.Authorize("16505551212", "101", "yourPassword")
```

### API Requests

General API requests can be made via the `platform` object's `Get`, `Post`, `Put`, and `Delete` methods.
Individual API calls are documented on the [API Developer and Reference Guide](https://developers.ringcentral.com/api-docs/latest/index.html).

The below SMS example shows how this can be used.

#### SMS Example

In order to send an SMS using the API, make a POST request to `/account/~/extension/~/sms`:

```go
import(
	"net/http"
	"net/url"
)

resp, err := platform.Post("/account/~/extension/~/sms", url.Values{}, []byte(`{ 
	"to"   : [{"phoneNumber": "14155551212"}],
	"from" :  {"phoneNumber": "16505551212"}, 
	"text" : "Test from Go"
}`), http.Header{})

```

#### Fax Example

Request helpers are provided for more complicated requests. The below shows usage of the
fax request helper which is used to help create the `multipart/mixed` HTTP request.

More information on usage is available in `./rcsdk/helpers/faxrequest/README.md`.

```go
import(
	"github.com/grokify/ringcentral-sdk-go/rcsdk/helpers/faxrequest"
	obj "github.com/grokify/ringcentral-sdk-go/rcsdk/helpers/objects"
)

fax, err := faxrequest.NewRequestHelper(faxrequest.Metadata{
	To: []obj.CallerInfo{
		obj.CallerInfo{PhoneNumber: "+16505626570"}},
	CoverPageText: "RingCentral fax example in Go!"})
err = fax.AddText([]byte("Hello World!"), "text/plain")
err = fax.AddFile("/path/to/myfile1.pdf")
err = fax.AddFile("/path/to/myfile2.tif")
err = fax.Finalize()

resp, err := platform.Post("/account/~/extension/~/fax", url.Values{}, fax.GetBody(), fax.GetHeaders())
```

## Links

Project Repo

* https://github.com/grokify/ringcentral-sdk-go

RingCentral API Docs

* https://developers.ringcentral.com/library.html

RingCentral API Explorer

* http://ringcentral.github.io/api-explorer

RingCentral Official SDKs

* https://github.com/ringcentral

## Contributions

Any reports of problems, comments or suggestions are most welcome.

Please report these on [Github](https://github.com/grokify/ringcentral-sdk-go)

## License

RingCentral SDK is available under an MIT-style license. See {file:LICENSE.txt} for details.

RingCentral SDK &copy; 2015 by John Wang
