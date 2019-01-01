package model

func CreateGame(creator *Player) *Game {
	var players []*Player
	players = append(players, creator)
	return &Game{Players: players, Board: CreateBoard()}
}

type Game struct {
	Players []*Player
	Board   *Board
}
