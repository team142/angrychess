package model

type Player struct {
	Profile *Profile
	Color   bool
	Team    int
	Board   *Board
	Pieces  []*Piece
}

func (p *Player) GetPieceByID(id string) (piece Piece, found bool) {
	for _, p := range p.Pieces {
		if p.ID == id {
			piece, found = p, true
			return
		}
	}
	return
}
