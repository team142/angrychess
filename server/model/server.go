package model

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"log"
)

//CreateServer starts a new server
func CreateServer(address string, handler func(*Server, *ws.Client, []byte)) *Server {
	s := &Server{
		Address: address,
		handler: handler,
		Lobby:   make(map[*ws.Client]*Profile),
		Games:   make(map[string]*Game),
		todo:    make(chan *item, 256),
	}
	s.run()
	return s
}

//Server holds server state
type Server struct {
	Address string
	Lobby   map[*ws.Client]*Profile
	Games   map[string]*Game
	handler func(*Server, *ws.Client, []byte)
	todo    chan *item
}

type item struct {
	client *ws.Client
	msg    []byte
}

func (s *Server) run() {
	go func() {
		for i := range s.todo {
			s.handler(s, i.client, i.msg)
		}
	}()
}

//GameByClient for easy access
func (s *Server) GameByClient(client *ws.Client) (found bool, game *Game) {
	for _, game := range s.Games {
		if game.Owner.Profile.Client == client {
			return true, game
		}
	}
	return
}

//HandleMessage This message is called by other parts of the system - the interface to the server
func (s *Server) HandleMessage(client *ws.Client, msg []byte) {
	i := &item{
		client: client,
		msg:    msg,
	}
	s.todo <- i

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
func (s *Server) CreateGame(client *ws.Client) *Game {
	player := &Player{
		Profile: s.Lobby[client],
		Team:    1,
	}

	g := CreateGame(player)
	s.Games[g.ID] = g

	reply := CreateMessageView(ViewBoard)
	b, _ := json.Marshal(reply)
	g.Announce(b)

	g.ShareState()
	log.Println(">> Created game ", g.Title)
	return g
}

//JoinGame for easy access
func (s *Server) JoinGame(gameID string, p *Profile) *Game {
	player := &Player{
		Profile: s.Lobby[p.Client],
	}
	game := s.Games[gameID]
	game.JoinGame(player)

	reply := CreateMessageView(ViewBoard)
	b, _ := json.Marshal(reply)
	game.Announce(b)

	log.Println(">> ", player.Profile.Nick, " joined game ", game.Title)
	game.ShareState()
	return game

}

//ListOfGames produces a light struct that describes the games hosted
func (s *Server) ListOfGames() *ListOfGames {
	result := ListOfGames{Games: []map[string]string{}}
	for _, game := range s.Games {
		item := make(map[string]string)
		item["id"] = game.ID
		item["title"] = game.Title
		item["players"] = fmt.Sprint(len(game.Players), "/", game.MaxPlayers())
		result.Games = append(result.Games, item)
	}
	return &result
}

//SetNick sets profiles nickname
func (s *Server) SetNick(client *ws.Client, nick string) {

	nick = s.createUniqueNick(nick)

	profile := s.GetOrCreateProfile(client)
	profile.Nick = nick

	log.Println(">> Set profile nick: ", profile.Nick)

	reply := CreateMessageSecret(profile.Secret, profile.ID)
	b, _ := json.Marshal(reply)
	client.Send <- b

}

//StartGame starts a game if possible
func (s *Server) StartGame(client *ws.Client) {
	found, game := s.GameByClient(client)
	if !found {
		log.Println(fmt.Sprintf("Error finding game owned by, %v with nick %v", client, s.Lobby[client].Nick))
		return
	}
	game.StartGame()

}

//Move attempts to move a piece
func (s *Server) Move(message MessageMove, client *ws.Client) {
	foundGame, game := s.GameByClient(client)
	if !foundGame {
		log.Println(fmt.Sprintf("Error finding game"))
		return
	}
	game.Move(client, message)

}

//Place attempts to place a piece
func (s *Server) Place(message MessagePlace, client *ws.Client) {
	foundGame, game := s.GameByClient(client)
	if !foundGame {
		log.Println(fmt.Sprintf("Error finding game"))
		return
	}
	game.Place(client, message)

}

func (s *Server) ChangeSeat(client *ws.Client, seat int) {
	_, game := s.GameByClient(client)
	game.ChangeSeat(client, seat)
}

func (s *Server) createUniqueNick(nickIn string) string {
	nick := nickIn
	ok := false
	i := 0
	for !ok {
		ok = true
		for _, b := range s.Lobby {
			if b.Nick == nick {
				ok = false
				break
			}
		}
		if ok {
			break
		}
		i++
		nick = fmt.Sprintf("%s%v", nickIn, i)
	}
	return nick

}

func (s *Server) Disconnect(client *ws.Client) {
	log.Println(">> Going to handle disconnect")
	found, game := s.GameByClient(client)
	if found {
		game.RemoveClient(client)
		if len(game.Players) == 0 {
			log.Println(">> Game is empty. Removing game")
			delete(s.Games, game.ID)
		}
	} else {
		log.Println(">> Player disconnecting was not in game")
	}

	//Remove from server
	delete(s.Lobby, client)

}
