package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func createNewGame(s *model.Server, client *ws.Client) *model.Game {
	player := &model.Player{
		Profile: s.Lobby[client],
		Team:    1,
	}

	game := model.CreateGame(player)
	game.CanStartBeforeFull = s.CanStartBeforeFull
	s.Games[game.ID] = game

	game.DoWork(
		func(game *model.Game) {
			reply := model.CreateMessageView(model.ViewBoard)
			b, _ := json.Marshal(reply)
			game.Announce(b)
			game.ShareState()
		})

	log.Println(">> Created game ", game.Title)
	return game
}

//joinGame for easy access
func joinGame(s *model.Server, gameID string, p *model.Profile) *model.Game {
	player := &model.Player{
		Profile: s.Lobby[p.Client],
	}
	game := s.Games[gameID]
	ok := game.JoinGame(player)
	if !ok {
		reply := model.CreateMessageError("Could not join game", "Server is full")
		b, _ := json.Marshal(reply)
		p.Client.Send <- b
		return game
	}

	reply := model.CreateMessageView(model.ViewBoard)
	b, _ := json.Marshal(reply)

	game.DoWork(
		func(game *model.Game) {
			game.Announce(b)
			game.ShareState()
		})

	log.Println(">> ", player.Profile.Nick, " joined game ", game.Title)
	return game

}

//setNick sets profiles nickname
func setNick(s *model.Server, client *ws.Client, nick string) {

	nick = createUniqueNick(s, nick)

	profile := s.GetOrCreateProfile(client)
	profile.Nick = nick

	log.Println(">> Set profile nick: ", profile.Nick)

	reply := model.CreateMessageSecret(profile.Secret, profile.ID)
	b, _ := json.Marshal(reply)
	client.Send <- b

}

//StartGame starts a game if possible
func startGame(s *model.Server, client *ws.Client) {
	found, game := s.GameByClientOwner(client)
	if !found {
		log.Println(fmt.Sprintf("Error finding game owned by, %v with nick %v", client, s.Lobby[client].Nick))
		return
	}

	game.DoWork(
		func(game *model.Game) {
			game.StartGame()
		})

}

//move attempts to move a piece
func move(s *model.Server, message *model.MessageMove, client *ws.Client) {
	foundGame, game := s.GameByClientPlaying(client)
	if !foundGame {
		log.Println(fmt.Sprintf("Error finding game"))
		return
	}

	game.DoWork(
		func(game *model.Game) {
			//TODO: figure out where this logic should sit
			didMove := game.Move(client, message)
			if didMove {
				game.ChangeMoveFrom(client)
			}

		})

}

//changeSeat changes where a player sits
func changeSeat(s *model.Server, client *ws.Client, seat int) {
	_, game := s.GameByClientPlaying(client)

	game.DoWork(
		func(game *model.Game) {
			game.ChangeSeat(client, seat)
		})
}

func createUniqueNick(s *model.Server, nickIn string) string {
	nick := nickIn
	ok := false
	i := 1
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

//disconnect handles changes to server state when someone a websocket disconnects
func disconnect(s *model.Server, client *ws.Client) {
	log.Println(">> Going to handle disconnect")
	found, game := s.GameByClientPlaying(client)
	if found {
		game.RemoveClient(client)
		if len(game.Players) == 0 {
			log.Println(">> Game is empty. Removing game")
			game.Stop()
			delete(s.Games, game.ID)
		}
	} else {
		log.Println(">> Player disconnecting was not in game")
	}

	//Remove from server
	delete(s.Lobby, client)

}

//notifyLobby tells players without a game about a new game
func notifyLobby(s *model.Server) {
	reply := s.CreateListOfGames()
	b, _ := json.Marshal(&reply)

	for client := range s.Lobby {
		found, _ := s.GameByClientPlaying(client)
		if !found {
			client.Send <- b
		}
	}

}
