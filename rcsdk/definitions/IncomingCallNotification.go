package definitions

type IncomingCallNotification struct {
	Event          string `json:"event,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
	ExtensionId    string `json:"extensionId,omitempty"`
	ServerId       string `json:"serverId,omitempty"`
	From           string `json:"from,omitempty"`
	Action         string `json:"action,omitempty"`
	FromName       string `json:"fromName,omitempty"`
	SrvLvlExt      string `json:"srvLvlExt,omitempty"`
	RecUrl         string `json:"recUrl,omitempty"`
	ToName         string `json:"toName,omitempty"`
	Sid            string `json:"sid,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	SessionId      string `json:"sessionId,omitempty"`
	To             string `json:"to,omitempty"`
	ToUrl          string `json:"toUrl,omitempty"`
	SrvLvl         string `json:"srvLvl,omitempty"`
}
