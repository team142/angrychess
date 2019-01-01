package model

import (
	"github.com/team142/chessfor4/controller"
	"github.com/team142/chessfor4/io/ws"
)

func CreateServer(address string) *Server {
	return &Server{Address: address}
}

type Server struct {
	Address string
	Lobby   []*Profile
	Games   []*Game
}

func (s *Server) HandleMessage(client *ws.Client, msg []byte) {
	controller.HandleIncoming(s, client, msg)
}
