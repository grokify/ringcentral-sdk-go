package definitions

type IncomingCallNotification struct {
	RecUrl         string `json:"recUrl,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	ToName         string `json:"toName,omitempty"`
	Sid            string `json:"sid,omitempty"`
	SrvLvl         string `json:"srvLvl,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	FromName       string `json:"fromName,omitempty"`
	To             string `json:"to,omitempty"`
	Action         string `json:"action,omitempty"`
	From           string `json:"from,omitempty"`
	ToUrl          string `json:"toUrl,omitempty"`
	SrvLvlExt      string `json:"srvLvlExt,omitempty"`
	Event          string `json:"event,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
	ExtensionId    string `json:"extensionId,omitempty"`
}
