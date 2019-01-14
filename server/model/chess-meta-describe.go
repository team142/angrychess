package model

import "github.com/team142/chessfor4/util"

type MoveDescription struct {
	XDiff        int
	YDiff        int
	Down         bool //Y is decreasing
	Diagonal     bool
	BeingPlaced  bool
	BeingRemoved bool
	MovingBoards bool
	PawnOnSpawn  bool
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

	return result
}
