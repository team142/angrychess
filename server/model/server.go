package model

import (
	"github.com/team142/chessfor4/io/ws"
)

func CreateServer(address string, handler func(*Server, *ws.Client, []byte)) *Server {
	return &Server{Address: address, Handler: handler}
}

type Server struct {
	Address string
	Lobby   map[*ws.Client]*Profile
	Games   []*Game
	Handler func(*Server, *ws.Client, []byte)
}

func (s *Server) HandleMessage(client *ws.Client, msg []byte) {
	s.Handler(s, client, msg)
}

func (s *Server) GetOrCreateProfile(client *ws.Client) *Profile {
	p := s.Lobby[client]
	if p == nil {
		p = CreateProfile(client)
		s.Lobby[client] = p
	}
	return p
}
