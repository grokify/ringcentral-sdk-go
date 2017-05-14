package definitions

type IncomingCallNotification struct {
	Uuid           string `json:"uuid,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	SrvLvl         string `json:"srvLvl,omitempty"`
	ExtensionId    string `json:"extensionId,omitempty"`
	Action         string `json:"action,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	RecUrl         string `json:"recUrl,omitempty"`
	Event          string `json:"event,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	From           string `json:"from,omitempty"`
	Sid            string `json:"sid,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	FromName       string `json:"fromName,omitempty"`
	To             string `json:"to,omitempty"`
	ToName         string `json:"toName,omitempty"`
	ToUrl          string `json:"toUrl,omitempty"`
	SrvLvlExt      string `json:"srvLvlExt,omitempty"`
}
