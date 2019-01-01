package model

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
)

const (
	MaxSupportedBoards = 2
)

func CreateGame(creator *Player) *Game {
	var players []*Player
	players = append(players, creator)
	return &Game{
		ID:      uuid.NewV4().String(),
		Players: players,
		Boards:  MaxSupportedBoards,
		Title:   fmt.Sprintf("%s's game", creator.Profile.Nick),
	}
}

type Game struct {
	ID      string
	Title   string
	Players []*Player
	Boards  int
}

func (game *Game) FindPlayerBySecret(secret string) (result *Player, found bool) {
	for _, p := range game.Players {
		if p.Profile.IsMe(secret) {
			result, found = p, true
			return
		}
	}
	return
}

func (game *Game) ShareState() {
	b, err := json.Marshal(game)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	game.Players[0].Profile.Client.Send <- b
}
