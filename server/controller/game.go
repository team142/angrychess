package controller

import (
	"encoding/json"
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

//JoinGame gets a player into a game
func joinGame(game *model.Game, player *model.Player) bool {
	found, spot := game.FindSpot()
	if !found {
		return false
	}
	player.SetTeamColorAndBoard(spot, game.Boards)
	game.Players[spot] = player
	return true
}

//StartGame starts the game for all players
func startGame(game *model.Game) {
	ok, msg := game.IsReadyToStart()
	if !ok {
		reply := model.CreateMessageError("Failed to start game", msg)
		b, _ := json.Marshal(reply)
		game.Owner.Profile.Client.Send <- b
		return
	}

	game.SetupBoards()
	game.Started = true
	shareState(game)

}

//ShareState tells all players what is going on
func shareState(game *model.Game) {
	reply := model.CreateMessageShareState(game)
	b, err := json.Marshal(&reply)
	if err != nil {
		log.Println(err)
		return
	}
	for _, player := range game.Players {
		player.Profile.Client.Send <- b
	}

}

//Move moves a piece
func Move(game *model.Game, client *ws.Client, message *model.MessageMove) (didMove bool) {
	log.Println(">> Moving ")

	pieceFound, piece, piecePlayer := game.FindPiece(message.PieceID)

	//Can't move a piece that does not exists
	if !pieceFound {
		log.Println("Piece not found, " + message.PieceID)
		return
	}

	//Get the message sending player
	player, _, _ := game.PlayerByClient(client)

	//Can this player move?
	if !player.MyTurn {
		log.Println("Not my turn!, " + message.PieceID)
		return
	}

	//TODO: handle taking pieces off board here

	if piece.IsEqual(message) {
		log.Println("No move")
		return
	}

	//Does this player own this piece?
	if player != piecePlayer {
		log.Println("Player does not own piece, " + message.PieceID)
		return
	}

	//Check for bad state
	if message.Cache == false && message.Board == 0 {
		log.Println("Bad request.  Piece must be on board or in cache, not neither")
		return
	}

	description := model.CalcMoveDescription(game, player, piece, message)

	ok, taken, msg := model.IsMovePossible(player, piece, description)
	if !ok {
		log.Println(msg)
		return
	}

	if taken != nil {
		//TODO take piece
		//TakePiece(game, player, taken)
	}

	/*
		TODO: do other checks
	*/
	//if !player.OwnsPiece(move.PieceID) {
	//	err = fmt.Errorf("player doesnt not own piece: %s", move.PieceID)
	//}

	didMove = true
	piece.Move(message)
	return
}

//announce announces something to all players
func announce(game *model.Game, b []byte) {
	for _, player := range game.Players {
		player.Profile.Client.Send <- b
	}
}
