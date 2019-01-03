package model

import "github.com/satori/go.uuid"

//Player describes a client in a game
type Player struct {
	Profile *Profile `json:"profile"`
	Color   bool     `json:"color"`
	Team    int      `json:"team"`
	Pieces  []*Piece `json:"pieces"`
	MyTurn  bool     `json:"myTurn"`
}

//GetPieceByID for easy access
func (p *Player) GetPieceByID(id string) (piece *Piece, found bool) {
	for _, p := range p.Pieces {
		if p.ID == id {
			piece, found = p, true
			return
		}
	}
	return
}

//SetTeamAndColor derives the color and team
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

//SetupBoard initializes the board and players
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
			Identity: identityPawn,
		}
		if p.Color {
			piece.Y = 7
		} else {
			piece.Y = 2
		}
		p.Pieces = append(p.Pieces, piece)
	}

	//Two free pawns :D
	for i := 1; i <= 2; i++ {
		piece := &Piece{
			ID:       uuid.NewV4().String(),
			Color:    p.Color,
			X:        0,
			Y:        0,
			Identity: identityPawn,
		}
		p.Pieces = append(p.Pieces, piece)
	}

}

//OwnsPiece determines if the piece is owned by a player
func (p *Player) OwnsPiece(ID string) bool {
	for _, p := range p.Pieces {
		if p.ID == ID {
			return true
		}
	}
	return false
}
