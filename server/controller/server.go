package controller

import (
	"fmt"
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

func createGameByClient(s *model.Server, client *ws.Client) *model.Game {
	player := &model.Player{
		Profile: s.Lobby[client],
		Team:    1,
	}

	game := model.CreateGameAndRun(player)
	game.CanStartBeforeFull = s.CanStartBeforeFull
	s.Games[game.ID] = game

	game.DoWork(
		func(game *model.Game) {
			reply := model.CreateMessageView(model.ViewBoard)
			announce(game, reply)
			shareState(game)
		})

	log.Println(">> Created game ", game.Title)
	return game
}

//joinGame for easy access
func joinGameByClient(s *model.Server, gameID string, p *model.Profile) *model.Game {
	player := &model.Player{
		Profile: s.Lobby[p.Client],
	}
	game := s.Games[gameID]
	ok := joinGame(game, player)
	if !ok {
		reply := model.CreateMessageError("Could not join game", "Server is full")
		p.Client.SendObject(reply)
		return game
	}

	reply := model.CreateMessageView(model.ViewBoard)

	game.DoWork(
		func(game *model.Game) {
			announce(game, reply)
			shareState(game)
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
	client.SendObject(reply)

}

//StartGame starts a game if possible
func startGameByClient(s *model.Server, client *ws.Client) {
	found, game := s.GameByClientOwner(client)
	if !found {
		log.Println(fmt.Sprintf("Error finding game owned by, %v with nick %v", client, s.Lobby[client].Nick))
		return
	}

	game.DoWork(
		func(game *model.Game) {
			startGame(game)
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
			didMove := Move(game, client, message)
			if didMove {
				game.ChangeMoveFrom(client)
			}
			shareState(game)

		})

}

//changeSeat changes where a player sits
func changeSeatByClient(s *model.Server, client *ws.Client, seat int) {
	_, game := s.GameByClientPlaying(client)

	game.DoWork(
		func(game *model.Game) {
			game.ChangeSeat(client, seat)
			shareState(game)

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
	notify := false
	if found {
		game.RemoveClient(client)
		if len(game.Players) == 0 {
			log.Println(">> Game is empty. Removing game")
			game.Stop()
			delete(s.Games, game.ID)
			notify = true
		}
		shareState(game)
	} else {
		log.Println(">> Player disconnecting was not in game")
	}

	//Remove from server
	delete(s.Lobby, client)
	if notify {
		notifyLobby(s)
	}

}

//notifyLobby tells players without a game about a new game
func notifyLobby(s *model.Server) {
	reply := model.CreateMessageListOfGames(s.CreateListOfGames())
	for client := range s.Lobby {
		found, _ := s.GameByClientPlaying(client)
		if !found {
			client.SendObject(reply)
		}
	}

}
