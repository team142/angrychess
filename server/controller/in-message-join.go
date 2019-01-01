package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func handleInMessageJoinGame(server *model.Server, client *ws.Client, msg []byte) {
	var message messages.MessageJoinGame
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	game := server.JoinGame(message.ID, server.Lobby[client])

	reply := messages.CreateMessageView(messages.ViewBoard)
	b, err := json.Marshal(reply)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	fmt.Println(">> Created game ", game.Title)
	client.Send <- b

	game.ShareState()

}
