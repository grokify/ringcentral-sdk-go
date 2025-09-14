package platform

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/grokify/mogo/net/http/httpsimple"
	"github.com/grokify/mogo/net/http/httputilmore"

	"github.com/grokify/ringcentral-sdk-go/rcsdk/core"
	rchttp "github.com/grokify/ringcentral-sdk-go/rcsdk/http"
)

const (
	AccessTokenTTL          int    = 3600
	RefreshTokenTTL         int    = 36000  // 10 hours
	RefreshTokenTTLRemember int    = 604800 // 1 week
	AccountPrefix           string = "/account/"
	AccountID               string = "~"
	TokenEndpoint           string = "/restapi/oauth/token" // #nosec G101
	RevokeEndpoint          string = "/restapi/oauth/revoke"
	APIVersion              string = "v1.0"
	URLPrefix               string = "/restapi"
)

type Platform struct {
	server    string
	appKey    string
	appSecret string
	Auth      Auth
	context   core.Context
}

func NewPlatform(context core.Context, appKey string, appSecret string, server string) Platform {
	p := Platform{appKey: appKey, appSecret: appSecret, server: server, Auth: NewAuth()}
	p.Auth = NewAuth()
	p.context = context
	return p
}

func (p *Platform) IsAuthorized() *Platform {
	return p
}

func (p *Platform) APIURL(url string) string {
	absUrl := strings.Join([]string{p.server, URLPrefix, "/", APIVersion, url}, "")
	return absUrl
}

func (p *Platform) Authorize(ctx context.Context, username string, extension string, password string, remember bool) (*http.Response, error) {
	res, err := p.authCall(ctx, username, extension, password)
	if err != nil {
		return res, err
	}

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		err = errors.New("HTTP_RESPONSE_ERROR " + strconv.Itoa(res.StatusCode))
		return res, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return res, err
	}
	authData, err := NewAuthCallResponseData(data)
	if err != nil {
		return res, err
	}
	p.Auth.SetData(authData)

	return res, err
}

func (p *Platform) GetApiKey() string {
	data := strings.Join([]string{p.appKey, p.appSecret}, ":")
	apiKey := base64.StdEncoding.EncodeToString([]byte(data))
	return apiKey
}

func (p *Platform) GetAuthHeader() string {
	return strings.Join([]string{p.Auth.GetTokenType(), p.Auth.GetAccessToken()}, " ")
}

func (p *Platform) apiCall(req rchttp.Request) (*http.Response, error) {
	p.IsAuthorized()

	head := req.Headers()
	head.Add(httputilmore.HeaderAuthorization, p.GetAuthHeader())

	log.Printf("[INFO] SDK_Request_URL [%s]", req.Url())

	req.SetUrl(p.APIURL(req.Url()))
	return req.Send()
}

func (p *Platform) APICall(ctx context.Context, req httpsimple.Request) (*http.Response, error) {
	p.IsAuthorized()

	req.Headers.Add(httputilmore.HeaderAuthorization, p.GetAuthHeader())

	return req.Do(ctx, nil)
}

/*
func (p *Platform) APICall2(req rchttp.Request2) (*http.Response, error) {
	p.IsAuthorized()

	req.Headers.Add(httputilmore.HeaderAuthorization, p.GetAuthHeader())

	log.Printf("[INFO] SDK_Request_URL [%s]", req.URL)

	absURL := p.ApiUrl(req.URL)
	queryString := req.Query.Encode()
	if len(queryString) > 0 {
		absURL = fmt.Sprintf("%s?%s", absURL, queryString)
	}
	req.URL = absURL

	return req.Send()
}
*/

func (p *Platform) authCall(ctx context.Context, username string, extension string, password string) (*http.Response, error) {
	// BODY
	data := url.Values{}
	data.Add("grant_type", "password")
	data.Add("username", username)
	data.Add("extension", extension)
	data.Add("password", password)

	req := httpsimple.Request{
		Method: http.MethodPost,
		URL:    strings.Join([]string{p.server, TokenEndpoint}, ""),
		Headers: http.Header{
			httputilmore.HeaderAuthorization: []string{strings.Join([]string{"Basic", p.GetApiKey()}, " ")},
			httputilmore.HeaderContentType:   []string{httputilmore.ContentTypeAppFormURLEncodedUtf8},
		},
		Body: data.Encode,
	}

	return req.Do(ctx, nil)
	/*
		if 1 == 0 {
			// URL
			tokenURL := strings.Join([]string{p.server, TOKEN_ENDPOINT}, "")

			// BODY
			data := url.Values{}
			data.Add("grant_type", "password")
			data.Add("username", username)
			data.Add("extension", extension)
			data.Add("password", password)
			body := strings.NewReader(data.Encode())

			// REQUEST
			req, _ := http.NewRequest(http.MethodPost, tokenURL, body)

			authzHeaderVal := strings.Join([]string{"Basic", p.GetApiKey()}, " ")
			req.Header.Add(httputilmore.HeaderAuthorization, authzHeaderVal)
			req.Header.Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppFormURLEncodedUtf8)

			// RESPONSE
			client := &http.Client{}
			return client.Do(req)
		}
	*/
}

func (p *Platform) Send(req rchttp.Request) (*http.Response, error) {
	return p.apiCall(req)
	// rcreq := p.context.GetRequest(strings.ToUpper(req.Method()), req.Url(), req.Query(), req.Body(), req.Headers())
	// return p.apiCall(rcreq)
}

func (p *Platform) Get(url string, queryParameters url.Values, body io.Reader, headers http.Header) (*http.Response, error) {
	rcreq := p.context.GetRequest(http.MethodGet, url, queryParameters, body, headers)
	return p.apiCall(rcreq)
}

func (p *Platform) Post(url string, queryParameters url.Values, body io.Reader, headers http.Header) (*http.Response, error) {
	rcreq := p.context.GetRequest(http.MethodPost, url, queryParameters, body, headers)
	return p.apiCall(rcreq)
}

func (p *Platform) Put(url string, queryParameters url.Values, body io.Reader, headers http.Header) (*http.Response, error) {
	rcreq := p.context.GetRequest(http.MethodPut, url, queryParameters, body, headers)
	return p.apiCall(rcreq)
}

func (p *Platform) Delete(url string, queryParameters url.Values, body io.Reader, headers http.Header) (*http.Response, error) {
	rcreq := p.context.GetRequest(http.MethodDelete, url, queryParameters, body, headers)
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
