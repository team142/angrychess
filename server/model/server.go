package model

import (
	"github.com/team142/chessfor4/io/ws"
)

func CreateServer(address string, handler func(*Server, *ws.Client, []byte)) *Server {
	return &Server{Address: address, Handler: handler}
}

type Server struct {
	Address string
	Lobby   map[string]*Profile
	Games   []*Game
	Handler func(*Server, *ws.Client, []byte)
}

func (s *Server) HandleMessage(client *ws.Client, msg []byte) {
	s.Handler(s, client, msg)
}

func (s *Server) GetOrCreateProfile(clientID string) *Profile {
	p := s.Lobby[clientID]
	if p == nil {
		p = CreateProfile(clientID)
		s.Lobby[clientID] = p
	}
	return p
}
