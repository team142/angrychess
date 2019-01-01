package model

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
