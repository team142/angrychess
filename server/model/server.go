package model

import (
	"fmt"
	"github.com/team142/angrychess/io/ws"
)

const (
	maxSupportedBoards = 2
)

//CreateServer starts a new server
func CreateServer(address string, handler func(*Server, *ws.Client, *[]byte), canStartBeforeFull bool) *Server {
	s := &Server{
		Address:            address,
		handler:            handler,
		Lobby:              make(map[*ws.Client]*Profile),
		Games:              make(map[string]*Game),
		Todo:               make(chan *item, 256),
		CanStartBeforeFull: canStartBeforeFull,
	}
	s.run()
	return s
}

//Server holds server state
type Server struct {
	Address            string
	Lobby              map[*ws.Client]*Profile
	Games              map[string]*Game
	handler            func(*Server, *ws.Client, *[]byte)
	Todo               chan *item
	CanStartBeforeFull bool
}

type item struct {
	client *ws.Client
	msg    *[]byte
}

func (s *Server) run() {
	go func() {
		for i := range s.Todo {
			s.handler(s, i.client, i.msg)
		}
	}()
}

//GameByClientOwner finds a game owned by client
func (s *Server) GameByClientOwner(client *ws.Client) (found bool, game *Game) {
	for _, game := range s.Games {
		if game.Owner.Profile.Client == client {
			return true, game
		}
	}
	return
}

//GameByClientPlaying find any player in a game
func (s *Server) GameByClientPlaying(client *ws.Client) (found bool, game *Game) {
	for _, game := range s.Games {
		for _, player := range game.Players {
			if player.Profile.Client == client {
				return true, game
			}
		}
	}
	return
}

//HandleMessage This message is called by other parts of the system - the interface to the server
func (s *Server) HandleMessage(client *ws.Client, msg *[]byte) {
	i := &item{
		client: client,
		msg:    msg,
	}
	s.Todo <- i

}

//GetOrCreateProfile creates profiles from a websocket client
func (s *Server) GetOrCreateProfile(client *ws.Client) *Profile {
	p := s.Lobby[client]
	if p == nil {
		p = createProfile(client)
		s.Lobby[client] = p
	}
	return p
}

//ListOfGames produces a light struct that describes the games hosted
func (s *Server) CreateListOfGames() *ListOfGames {
	result := ListOfGames{Games: []map[string]string{}}
	for _, game := range s.Games {
		row := make(map[string]string)
		row["id"] = game.ID
		row["title"] = game.Title
		row["players"] = fmt.Sprint(len(game.Players), "/", game.MaxPlayers())
		result.Games = append(result.Games, row)
	}
	return &result
}

//CreateMessageListOfGames creates a list of games
func (s *Server) CreateMessageListOfGames() *MessageListOfGames {
	list := s.CreateListOfGames()
	return CreateMessageListOfGames(list)

}
