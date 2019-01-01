package messages

func CreateMessageSecret(secret string) MessageSecret {
	return MessageSecret{Msg: "secret", Secret: secret}
}

type MessageSecret struct {
	Msg    string `json:"msg"`
	Secret string `json:"secret"`
}
