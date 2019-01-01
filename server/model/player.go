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

func (p *Player) SetTeamAndColor(spot int, boards int) {
	c := spot + boards
	b := c%2 == 0
	p.Color = b
	if spot <= boards {
		p.Team = 1
	} else {
		p.Team = 2
	}

}
