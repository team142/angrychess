package model

import "github.com/team142/angrychess/util"

type MoveDescription struct {
	XDiff          int
	YDiff          int
	Down           bool //Y is decreasing
	Diagonal       bool
	BeingPlaced    bool
	BeingRemoved   bool
	MovingBoards   bool
	PawnOnSpawn    bool
	LastTwoRows    bool
	OtherBoard     bool
	LandingOnPiece *Piece
	PiecesBetween  []*Piece
}

func CalcMoveDescription(game *Game, player *Player, piece *Piece, move *MessageMove) *MoveDescription {
	result := &MoveDescription{}
	result.YDiff = util.Abs(piece.Y - move.ToY)
	result.XDiff = util.Abs(piece.X - move.ToX)
	result.Down = piece.Y > move.ToY
	result.BeingPlaced = piece.Cache && !move.Cache
	result.BeingRemoved = !piece.Cache && move.Cache
	result.MovingBoards = piece.Board != move.Board
	result.PawnOnSpawn = player.Team == 1 && piece.Y == 7 || player.Team == 2 && piece.Y == 2
	result.LastTwoRows = (1 == player.Team && 2 >= piece.Y) || (piece.Y >= 7 && 2 == player.Team)
	result.OtherBoard = player.Board != move.Board

	CalcPiecesBetween(game, player, piece, move, result)

Outer:
	for _, pl := range game.Players {
		for _, pi := range pl.Pieces {
			if pi.IsEqual(move) {
				result.LandingOnPiece = pi
				break Outer
			}
		}
	}
	return result

}

func CalcPiecesBetween(game *Game, player *Player, piece *Piece, move *MessageMove, result *MoveDescription) {
	//Don't worry about one tile
	if result.XDiff+result.YDiff <= 1 {
		return
	}
	//Knights don't worry about pieces between
	if piece.Identity == identityKnight {
		return
	}

	//Horizontal moves and is also greater than 1 diff
	if result.XDiff > 1 && result.YDiff == 0 {
		rx1, _, rx2, _ := util.OrderPoints(move.ToX, move.ToY, piece.X, piece.Y)
		for x := rx1 + 1; x < rx2; x++ {
			//TODO
			//Find piece add to result
		}
	}

}
