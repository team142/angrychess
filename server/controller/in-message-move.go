package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func handleInMessageMove(server *model.Server, client *ws.Client, msg []byte) {
	found, game := server.FindGameByClient(client)
	if !found {
		log.Println(fmt.Sprintf("Error finding game"))
		return
	}

	game.FindPlayerBySecret()

	var message messages.MessageMove
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}

	log.Println(">> Moving ")
	err := game.Move(message)
	if err != nil {
		log.Println(fmt.Sprintf("Error trying to move, %s", err))
		//TODO: send error message to players
		return
	}
	game.ShareState()

}
