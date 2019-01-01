package model

type Player struct {
	Profile Profile
	Color   bool
	Team    int
	Board   *Board
	Pieces  []*Piece
}
