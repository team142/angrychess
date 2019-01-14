package model

import (
	"log"
)

func IsMovePossible(game *Game, player *Player, piece *Piece, description *MoveDescription) (ok bool, taken *Piece, msg string) {
	if piece.Identity == identityPawn {
		return isMovePossiblePawn(game, player, piece, description)
	}
	return
}

func isMovePossiblePawn(game *Game, player *Player, piece *Piece, description *MoveDescription) (ok bool, taken *Piece, msg string) {

	if description.MovingBoards {
		log.Println("Can't move boards")
		ok = false
		return
	}

	if description.BeingPlaced {
		//TODO: other checks
		// - is empty tile
		// - not in check

		if description.LastTwoRows {
			log.Println("Can't place on last two enemy lines")
			ok = false
			return
		}
		ok = true
		return
	}

	if description.XDiff > 2 {
		log.Println("X diff must be only move 1 or 2 blocks")
		ok = false
		return
	}

	if description.YDiff > 1 {
		log.Println("Must only move 0 or 1 blocks")
		ok = false
		return
	}

	if player.shouldGoDown() != description.Down {
		log.Println("Expected goingDown ", description.Down, " equal to should go down ", player.shouldGoDown())
		ok = false
		return
	}

	if !description.PawnOnSpawn && description.XDiff == 2 {
		log.Println("Can't move two when not on start row")
		ok = false
		return
	}

	if description.OtherBoard {
		log.Println("Can't move onto another board")
		ok = false
		return
	}

}
