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
	if boards == 2 && (spot == 1 || spot == 2) {
		p.Team = 1
	} else if boards == 2 && (spot == 3 || spot == 4) {
		p.Team = 2
	}

	if boards == 2 && (spot == 1 || spot == 4) {
		p.Color = false
	} else if boards == 2 && (spot == 2 || spot == 3) {
		p.Color = true
	}

}
