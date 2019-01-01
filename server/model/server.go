package model

import (
	"fmt"
	"github.com/team142/chessfor4/io/ws"
)

//CreateServer starts a new server
func CreateServer(address string, handler func(*Server, *ws.Client, []byte)) *Server {
	return &Server{
		Address: address,
		Handler: handler,
		Lobby:   make(map[*ws.Client]*Profile),
		Games:   make(map[string]*Game),
	}
}

//Server holds server state
type Server struct {
	Address string
	Lobby   map[*ws.Client]*Profile
	Games   map[string]*Game
	Handler func(*Server, *ws.Client, []byte)
}

//FindGameByClient for easy access
func (s *Server) FindGameByClient(client *ws.Client) (found bool, game *Game) {
	for _, game := range s.Games {
		if game.Owner.Profile.Client == client {
			return true, game
		}
	}
	return
}

//HandleMessage A handler for messages being given to this server
func (s *Server) HandleMessage(client *ws.Client, msg []byte) {
	s.Handler(s, client, msg)
}

//GetOrCreateProfile creates profiles from a websocket client
func (s *Server) GetOrCreateProfile(client *ws.Client) *Profile {
	p := s.Lobby[client]
	if p == nil {
		p = CreateProfile(client)
		s.Lobby[client] = p
	}
	return p
}

//CreateGame for easy access
func (s *Server) CreateGame(p *Player) *Game {
	g := CreateGame(p)
	s.Games[g.ID] = g
	return g
}

//JoinGame for easy access
func (s *Server) JoinGame(gameID string, p *Profile) *Game {
	player := &Player{
		Profile: s.Lobby[p.Client],
	}
	game := s.Games[gameID]
	game.JoinGame(player)
	return game

}

//CreateListOfGames produces a light struct that describes the games hosted
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
