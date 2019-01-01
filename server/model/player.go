package model

import "github.com/satori/go.uuid"

type Player struct {
	Profile *Profile `json:"profile"`
	Color   bool     `json:"color"`
	Team    int      `json:"team"`
	Pieces  []*Piece `json:"pieces"`
	MyTurn  bool     `json:"myTurn"`
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

func (p *Player) SetupBoard() {
	if !p.Color {
		p.MyTurn = true
	}

	//Pawns
	for i := 1; i <= 8; i++ {
		piece := &Piece{
			ID:       uuid.NewV4().String(),
			Color:    p.Color,
			X:        i,
			Identity: IdentityPawn,
		}
		if p.Color {
			piece.Y = 7
		} else {
			piece.Y = 2
		}
		p.Pieces = append(p.Pieces, piece)
	}

}

func (p *Player) OwnsPiece(ID string) bool {
	for _, p := range p.Pieces {
		if p.ID == ID {
			return true
		}
	}
	return false
}
