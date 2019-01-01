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
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Players []*Player `json:"players"`
	Boards  int       `json:"boards"`
}

func (game *Game) JoinGame(player *Player) bool {
	if len(game.Players) >= game.GetMaxPlayers() {
		return false
	}
	found, team, color := game.findSpot()
	if !found {
		return false
	}
	player.Team = team
	player.Color = color

	game.Players = append(game.Players, player)

	return true
}

func (game *Game) findSpot() (found bool, team int, color bool) {
	var slots [][]bool
	for x := 1; x <= 2; x++ {
		for y := 1; y <= 2; y++ {
			slots[x][y] = false
		}
	}

	for _, player := range game.Players {
		team := player.Team - 1
		color := 0
		if player.Color {
			color = 1
		}
		slots[team][color] = true
	}

	for x := 1; x <= 2; x++ {
		for y := 1; y <= 2; y++ {
			if slots[x][y] == false {
				return true, x, y == 0
			}
		}
	}
	found = false
	return

}

func (game *Game) GetMaxPlayers() int {
	return game.Boards * 2
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
	b, err := json.Marshal(&game)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	game.Players[0].Profile.Client.Send <- b
}
