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

//Move moves piece
func (piece *Piece) Move(message MessageMove) {
	piece.X = message.ToX
	piece.Y = message.ToY

}

//Place places a piece on a board at a point
func (piece *Piece) Place(message MessagePlace) {
	piece.X = message.ToY
	piece.Y = message.ToY

}
