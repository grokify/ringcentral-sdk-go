package definitions

type IncomingCallNotification struct {
	Uuid           string `json:"uuid,omitempty"`
	Sid            string `json:"sid,omitempty"`
	ToUrl          string `json:"toUrl,omitempty"`
	FromName       string `json:"fromName,omitempty"`
	SrvLvl         string `json:"srvLvl,omitempty"`
	SrvLvlExt      string `json:"srvLvlExt,omitempty"`
	RecUrl         string `json:"recUrl,omitempty"`
	Event          string `json:"event,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	From           string `json:"from,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	ExtensionId    string `json:"extensionId,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	ToName         string `json:"toName,omitempty"`
	Action         string `json:"action,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	To             string `json:"to,omitempty"`
}
