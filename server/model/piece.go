package model

import (
	"github.com/team142/chessfor4/util"
	"log"
)

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
	Cache    bool   `json:"cache"`
	Board    int    `json:"board"`
}

//Move moves piece
func (piece *Piece) Move(message *MessageMove) {
	piece.Cache = message.Cache
	piece.Board = message.Board
	if !piece.Cache {
		piece.X = message.ToX
		piece.Y = message.ToY
	}

}

//CanMoveLikeThat checks that the piece can make those sorts of moves
func (piece *Piece) CanMoveLikeThat(player *Player, move *MessageMove) (ok bool) {
	ok = true
	if piece.Identity == identityPawn {
		shouldGoDown := player.Team == 1
		movingOne := util.Abs(piece.Y-move.ToY) == 1
		movingTwo := util.Abs(piece.Y-move.ToY) == 2
		goingDown := piece.Y > move.ToY
		placeOntoBoard := piece.Cache && !move.Cache
		movingBoards := piece.Board != move.Board
		isOnStartRow := player.Team == 1 && piece.Y == 7 || player.Team == 2 && piece.Y == 2

		if movingBoards {
			log.Println("Can't move boards")
			ok = false
			return
		}

		if placeOntoBoard {
			//TODO: other checks
			if isLastTwo(player, move.ToY) {
				log.Println("Can't place on last two enemy lines")
				ok = false
				return
			}
			ok = true
			return
		}

		if !movingTwo && !movingOne {
			log.Println("Must only move 1 or 2 blocks")
			ok = false
			return
		}

		if shouldGoDown != goingDown {
			log.Println("Expected goingDown ", goingDown, " equal to should go down ", shouldGoDown)
			ok = false
			return
		}

		if !isOnStartRow && movingTwo {
			log.Println("Can't move two when not on start row")
			ok = false
			return
		}

		if player.Board != move.Board {
			log.Println("Can't move onto another board")
			ok = false
			return
		}

	}

	return
}

func isLastTwo(player *Player, y int) bool {
	return (1 == player.Team && 2 >= y) || (y >= 7 && 2 == player.Team)
}

func (piece *Piece) isEqual(move *MessageMove) bool {
	if piece.Board != move.Board {
		return false
	}
	if piece.Cache != move.Cache {
		return false
	}
	if piece.X != move.ToX {
		return false
	}
	if piece.Y != move.ToY {
		return false
	}
	return true
}

func createRook(id string, board int, color bool) *Piece {
	return &Piece{
		ID:       id,
		Board:    board,
		Color:    color,
		Identity: identityRook,
	}
}
