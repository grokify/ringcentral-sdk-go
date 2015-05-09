package platform

type Auth struct {
	remember bool

	tokenType string

	accessToken string
	expiresIn   int
	expireTime  int

	refreshToken           string
	refreshTokenExpiresIn  int
	refresnTokenExpireTime int

	scope   string
	ownerId string
}

func NewAuth() Auth {
	a := Auth{}
	a.Reset()
	return a
}

func (a *Auth) SetData(authData AuthCallResponseData) *Auth {

	// MISC
	a.tokenType = authData.TokenType

	// ACCESS TOKEN
	a.accessToken = authData.AccessToken
	a.expiresIn = authData.ExpiresIn

	a.refreshToken = authData.RefreshToken
	a.refreshTokenExpiresIn = authData.RefreshTokenExpiresIn

	return a
}

func (a *Auth) Reset() *Auth {
	a.remember = false

	a.tokenType = ""

	a.accessToken = ""
	a.expiresIn = 0
	a.expireTime = 0

	a.refreshToken = ""
	a.refreshTokenExpiresIn = 0
	a.refresnTokenExpireTime = 0

	a.scope = ""
	a.ownerId = ""
	return a
}

func (a *Auth) GetAccessToken() string {
	return a.accessToken
}

func (a *Auth) GetTokenType() string {
	return a.tokenType
}
