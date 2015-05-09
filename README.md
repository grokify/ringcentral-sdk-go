# Table of contents

1. [Overview](#overview)
2. [Installation](#installation)
3. [Core Module](#core-module)
4. [SMS](#sms)

***

# Overview

This is a alpha SDK for the RingCentral for Developers platform. It attempts to mirror the structure of the official RingCentral SDKs. As an alpha, many parts of the SDK are not yet implemented yet.

# Installation

```bash
$ go get github.com/grokify/ringcentral-sdk-go
```

# Core Module

## Instantiate the RCSDK object

The SDK is represented by the global RCSDK constructor. Your application must create an instance of this object.

```go
import(
	"github.com/grokify/ringcentral-sdk-go/rcsdk"
)
// For Production use: "https://platform.ringcentral.com"
// For Sandbox use:  "https://platform.devtest.ringcentral.com"
sdk := rcsdk.NewSdk("yourAppKey", "yourAppSecret", "https://platform.devtest.ringcentral.com")
```

## Get the Platform Singleton

```js
platform := sdk.GetPlatform();
```

## Login

Login is accomplished by calling the `platform.Authorize()` method of the Platform singleton with username, extension
(optional), and password as parameters. A `Promise` instance is returned, resolved with an AJAX `Response` object.

The `username` should be a phone number in E.164 format with or without the leading `+`.

```go
platform.Authorize('+16505551212','101','yourPassword')
```

# SMS

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

