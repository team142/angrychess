package controller

import (
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
)

func handleInMessageCreateGame(server *model.Server, client *ws.Client) {
	createGameByClient(server, client)
	notifyLobby(server)
}
