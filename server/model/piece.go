package model

const (
	identityPawn   = 1
	identityKnight = 2
	identityBishop = 3
	identityRook   = 4
	identityQueen  = 5
	identityKing   = 6
)

//Piece describes a piece on the board
type Piece struct {
	ID       string `json:"id"`
	Identity int    `json:"identity"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Color    bool   `json:"color"`
}

//TryMove attempts to move a piece.. probably should not be here
func (piece *Piece) TryMove(game *Game, color bool, fromX, fromY, toX, toY int) (ok bool, msg string) {
	//TODO: implement

	piece.X = toX
	piece.Y = toY
	ok = true

	return
}
