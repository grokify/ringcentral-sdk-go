package platform

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/grokify/gotilla/net/httputil"
	"github.com/grokify/ringcentral-sdk-go/rcsdk/core"
	rchttp "github.com/grokify/ringcentral-sdk-go/rcsdk/http"
)

const (
	ACCESS_TOKEN_TTL           int    = 3600
	REFRESH_TOKEN_TTL          int    = 36000  // 10 hours
	REFRESH_TOKEN_TTL_REMEMBER int    = 604800 // 1 week
	ACCOUNT_PREFIX             string = "/account/"
	ACCOUNT_ID                 string = "~"
	TOKEN_ENDPOINT             string = "/restapi/oauth/token"
	REVOKE_ENDPOINT            string = "/restapi/oauth/revoke"
	API_VERSION                string = "v1.0"
	URL_PREFIX                 string = "/restapi"
)

type Platform struct {
	server    string
	appKey    string
	appSecret string
	auth      Auth
	context   core.Context
}

func NewPlatform(context core.Context, appKey string, appSecret string, server string) Platform {
	p := Platform{appKey: appKey, appSecret: appSecret, server: server, auth: NewAuth()}
	p.auth = NewAuth()
	p.context = context
	return p
}

func (p *Platform) IsAuthorized() *Platform {
	return p
}

func (p *Platform) ApiUrl(url string) string {
	absUrl := strings.Join([]string{p.server, URL_PREFIX, "/", API_VERSION, url}, "")
	return absUrl
}

func (p *Platform) Authorize(username string, extension string, password string, remember bool) (*http.Response, error) {
	res, err := p.authCall(username, extension, password)
	if err == nil {
		data, err2 := httputil.ResponseBody(res)
		if err2 != nil {
			return res, err2
		}
		authData, err3 := NewAuthCallResponseData(data)
		if err3 != nil {
			return res, err3
		}
		p.auth.SetData(authData)
	}
	return res, err
}

func (p *Platform) GetApiKey() string {
	data := strings.Join([]string{p.appKey, p.appSecret}, ":")
	apiKey := base64.StdEncoding.EncodeToString([]byte(data))
	return apiKey
}

func (p *Platform) getAuthHeader() string {
	return strings.Join([]string{p.auth.GetTokenType(), p.auth.GetAccessToken()}, " ")
}

func (p *Platform) apiCall(req rchttp.Request) (*http.Response, error) {
	p.IsAuthorized()

	req.SetHeader("Authorization", p.getAuthHeader())
	req.SetUrl(p.ApiUrl(req.GetUrl()))
	return req.Send()
}

func (p *Platform) authCall(username string, extension string, password string) (*http.Response, error) {
	client := &http.Client{}

	// URL
	tokenUrl := strings.Join([]string{p.server, TOKEN_ENDPOINT}, "")

	// BODY
	data := url.Values{}
	data.Add("grant_type", "password")
	data.Add("username", username)
	data.Add("extension", extension)
	data.Add("password", password)
	body := strings.NewReader(data.Encode())

	// REQUEST
	req, _ := http.NewRequest("POST", tokenUrl, body)

	authzHeaderVal := strings.Join([]string{"Basic", p.GetApiKey()}, " ")
	req.Header.Add("Authorization", authzHeaderVal)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	// RESPONSE
	return client.Do(req)
}

func (p *Platform) Post(url string, queryParameters url.Values, body []byte, headers http.Header) (*http.Response, error) {
	rcreq := p.context.GetRequest("POST", url, queryParameters, body, headers)
	return p.apiCall(rcreq)
}

type AuthCallResponseData struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	OwnerId               string `json:"owner_id"`
}

func NewAuthCallResponseData(body []byte) (AuthCallResponseData, error) {
	a := AuthCallResponseData{}
	err := json.Unmarshal(body, &a)
	return a, err
}
