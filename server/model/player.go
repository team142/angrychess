package model

type Player struct {
	Profile *Profile `json:"profile"`
	Color   bool     `json:"color"`
	Team    int      `json:"team"`
	Pieces  []*Piece `json:"pieces"`
}

func (p *Player) GetPieceByID(id string) (piece *Piece, found bool) {
	for _, p := range p.Pieces {
		if p.ID == id {
			piece, found = p, true
			return
		}
	}
	return
}
