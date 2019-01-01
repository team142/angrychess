package messages

const (
	ViewMenu  = "view-menu"
	ViewBoard = "view-board"
)

func CreateMessageSecret(secret string) MessageSecret {
	return MessageSecret{Msg: "secret", Secret: secret}
}

type MessageSecret struct {
	Msg    string `json:"msg"`
	Secret string `json:"secret"`
}

func CreateMessageView(view string) MessageView {
	return MessageView{Msg: view, view: view}
}

type MessageView struct {
	Msg  string `json:"msg"`
	view string `json:"view"`
}
