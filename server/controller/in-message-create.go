package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func handleInMessageCreateGame(server *model.Server, client *ws.Client, msg []byte) {

	player := &model.Player{
		Profile: server.Lobby[client],
		Team:    1,
	}
	server.CreateGame(player)
	reply := messages.CreateMessageView(messages.ViewBoard)
	b, err := json.Marshal(reply)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	client.Send <- b

}
