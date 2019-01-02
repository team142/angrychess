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
	inMessageNick        = "nick"
	inMessageStartGame   = "start-game"
	inMessageCreateGame  = "create-game"
	inMessageJoinGame    = "join-game"
	inMessageListOfGames = "list-games"
	inMessageMove        = "move"
	inMessagePlace       = "place"
	inMessageChangeSeat  = "seat"
)

//HandleIncoming handles messages arriving to the websocket server
func HandleIncoming(server *model.Server, client *ws.Client, msg []byte) {
	var message messages.BaseMessage
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	if message.Msg == inMessageNick {
		handleInMessageNick(server, client, msg)

	} else if message.Msg == inMessageCreateGame {
		handleInMessageCreateGame(server, client)

	} else if message.Msg == inMessageJoinGame {
		handleInMessageJoinGame(server, client, msg)

	} else if message.Msg == inMessageStartGame {
		handleInMessageStartGame(server, client)

	} else if message.Msg == inMessageListOfGames {
		handleInMessageListOfGame(server, client)

	} else if message.Msg == inMessageMove {
		handleInMessageMove(server, client, msg)

	} else if message.Msg == inMessagePlace {
		handleInMessagePlace(server, client, msg)

	} else if message.Msg == inMessageChangeSeat {
		handleInMessageChangeSeat(server, client, msg)

	} else {
		log.Println("Unknown route: ", message.Msg)
	}

}
