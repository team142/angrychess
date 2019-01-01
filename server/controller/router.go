package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

const (
	InMessageNick       = "nick"
	InMessageStartGame  = "start-game"
	InMessageJoinGame   = "join-game"
	InMessageMove       = "move"
	InMessagePlace      = "place"
	InMessageChangeSeat = "seat"
)

func HandleIncoming(server *model.Server, client *ws.Client, msg []byte) {
	var message messages.BaseMessage
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	if message.Msg == InMessageNick {
		handleInMessageNick(server, client, msg)

	} else if message.Msg == InMessageStartGame {
		handleInMessageCreateGame(server, client, msg)

	} else if message.Msg == InMessageJoinGame {
		//TODO: handle route

	} else if message.Msg == InMessageMove {
		//TODO: handle route

	} else if message.Msg == InMessagePlace {
		//TODO: handle route

	} else if message.Msg == InMessageChangeSeat {
		//TODO: handle route

	} else {
		log.Println("Unknown route: ", message.Msg)
	}

}
