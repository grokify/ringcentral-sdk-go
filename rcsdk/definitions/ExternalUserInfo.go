package definitions

type ExternalUserInfo struct {
	AccountId string `json:"accountId,omitempty"`
	UserId    string `json:"userId,omitempty"`
	UserToken string `json:"userToken,omitempty"`
	UserType  int    `json:"userType,omitempty"`
}
