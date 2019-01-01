package controller

import (
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func handleInMessageCreateGame(server *model.Server, client *ws.Client, msg []byte) {

	player := &model.Player{
		Profile: server.Lobby[client],
		Team:    1,
	}
	g := server.CreateGame(player)
	log.Println(">> Created game ", g.Title)

	g.ShareState()

}
