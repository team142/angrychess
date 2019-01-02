package controller

import (
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func handleInMessageStartGame(server *model.Server, client *ws.Client) {

	found, game := server.FindGameByClient(client)
	if !found {
		log.Println(fmt.Sprintf("Error finding game owned by, %v", client))
		log.Println(fmt.Sprintf("Error finding game owned by player with nick, %v", server.Lobby[client].Nick))
		return
	}
	game.StartGame()

}
