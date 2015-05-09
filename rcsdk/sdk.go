package rcsdk

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/core"
	"github.com/grokify/ringcentral-sdk-go/rcsdk/platform"
)

type Sdk struct {
	platform platform.Platform
	context  core.Context
}

func NewSdk(appKey string, appSecret string, server string) Sdk {
	sdk := Sdk{}
	sdk.context = core.NewContext()
	sdk.platform = platform.NewPlatform(sdk.context, appKey, appSecret, server)
	return sdk
}

func (sdk *Sdk) GetPlatform() platform.Platform {
	return sdk.platform
}

func (sdk *Sdk) GetContext() core.Context {
	return sdk.context
}
