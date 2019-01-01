package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func handleInMessageStartGame(server *model.Server, client *ws.Client, msg []byte) {

	found, game := server.FindGameByClient(client)
	game.SetupBoards()

	if !found {
		log.Println(fmt.Sprintf("Error finding game owned by, %v", client))
		log.Println(fmt.Sprintf("Error finding game owned by player with nick, %v", server.Lobby[client].Nick))
	}

	reply := messages.CreateMessageView(messages.ViewBoard)
	b, _ := json.Marshal(reply)
	game.Announce(b)
	game.ShareState()

}
