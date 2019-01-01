package model

import (
	"github.com/team142/chessfor4/io/ws"
)

func CreateServer(address string, handler func(*Server, *ws.Client, []byte)) *Server {
	return &Server{Address: address, Handler: handler}
}

type Server struct {
	Address string
	Lobby   []*Profile
	Games   []*Game
	Handler func(*Server, *ws.Client, []byte)
}

func (s *Server) HandleMessage(client *ws.Client, msg []byte) {
	s.Handler(s, client, msg)
}
