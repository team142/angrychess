package model

import (
	"fmt"
	"github.com/team142/chessfor4/io/ws"
)

func CreateServer(address string, handler func(*Server, *ws.Client, []byte)) *Server {
	return &Server{
		Address: address,
		Handler: handler,
		Lobby:   make(map[*ws.Client]*Profile),
		Games:   make(map[string]*Game),
	}
}

type Server struct {
	Address string
	Lobby   map[*ws.Client]*Profile
	Games   map[string]*Game
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

func (s *Server) CreateGame(p *Player) *Game {
	g := CreateGame(p)
	s.Games[g.ID] = g
	return g
}

func (s *Server) JoinGame(gameID string, p *Profile) *Game {
	player := &Player{
		Profile: s.Lobby[p.Client],
	}
	game := s.Games[gameID]
	game.JoinGame(player)
	return game

}

func (s *Server) CreateListOfGames() *ListOfGames {
	result := ListOfGames{Games: []map[string]string{}}
	for _, game := range s.Games {
		item := make(map[string]string)
		item["id"] = game.ID
		item["title"] = game.Title
		item["players"] = fmt.Sprint(len(game.Players), "/", game.GetMaxPlayers())
		result.Games = append(result.Games, item)
	}
	return &result
}
