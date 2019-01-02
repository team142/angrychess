package messages

const (
	//ViewMenu for displaying the game menu
	ViewMenu = "view-menu"
	//ViewBoard for displaying the game board
	ViewBoard = "view-board"
)

//CreateMessageSecret for easy access
func CreateMessageSecret(secret string) MessageSecret {
	return MessageSecret{Msg: "secret", Secret: secret}
}

//MessageSecret for easy access
type MessageSecret struct {
	Msg    string `json:"msg"`
	Secret string `json:"secret"`
}

//CreateMessageView for easy access
func CreateMessageView(view string) MessageView {
	return MessageView{Msg: "view", View: view}
}

//MessageView for easy access
type MessageView struct {
	Msg  string `json:"msg"`
	View string `json:"view"`
}

func CreateMessageError(t, m string) MessageError {
	return MessageError{Title: t, Msg: m}
}

type MessageError struct {
	Title string `json:"title"`
	Msg   string `json:"msg"`
}
