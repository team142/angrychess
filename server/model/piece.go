package model

const (
	IdentityPawn   = 1
	IdentityKnight = 2
	IdentityBishop = 3
	IdentityRook   = 4
	IdentityQueen  = 5
	IdentityKing   = 6
)

type Piece struct {
	ID       string `json:"id"`
	Identity int    `json:"identity"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Color    bool   `json:"color"`
}

func (piece *Piece) TryMove(game *Game, color bool, fromX, fromY, toX, toY int) (err error) {
	//TODO: implement

	return
}
