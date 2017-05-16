package definitions

type IncomingCallNotification struct {
	Action         string `json:"action,omitempty"`
	Event          string `json:"event,omitempty"`
	ExtensionId    string `json:"extensionId,omitempty"`
	From           string `json:"from,omitempty"`
	FromName       string `json:"fromName,omitempty"`
	RecUrl         string `json:"recUrl,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	Sid            string `json:"sid,omitempty"`
	SrvLvl         string `json:"srvLvl,omitempty"`
	SrvLvlExt      string `json:"srvLvlExt,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	To             string `json:"to,omitempty"`
	ToName         string `json:"toName,omitempty"`
	ToUrl          string `json:"toUrl,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
}
