package controller

import (
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
)

func handleInMessageDC(server *model.Server, client *ws.Client) {
	server.Disconnect(client)

}
