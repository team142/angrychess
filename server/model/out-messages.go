package model

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

//CreateMessageError is a general error message constructor
func CreateMessageError(t, m string) MessageError {
	return MessageError{Msg: "error", Title: t, Body: m}
}

//MessageError is a general error message
type MessageError struct {
	Msg   string `json:"msg"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateMessageListOfGames(games *ListOfGames) MessageListOfGames {
	return MessageListOfGames{Msg: "list-games", Games: games}
}

type MessageListOfGames struct {
	Msg   string      `json:"msg"`
	Games interface{} `json:"games"`
}
