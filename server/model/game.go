package model

func CreateGame(creator *Player) *Game {
	var players []*Player
	players = append(players, creator)
	return &Game{Players: players, Board: CreateBoard()}
}

type Game struct {
	ID      string
	Players []*Player
	Board   *Board
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
