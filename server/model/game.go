package model

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	MaxSupportedBoards = 2
)

func CreateGame(creator *Player) *Game {
	var players []*Player
	players = append(players, creator)
	return &Game{Players: players, Boards: MaxSupportedBoards}
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

func (g *Game) ShareState() {
	b, err := json.Marshal(g)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	g.Players[0].Profile.Client.Send <- b
}
