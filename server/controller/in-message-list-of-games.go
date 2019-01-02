package controller

import (
	"encoding/json"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func handleInMessageListOfGame(server *model.Server, client *ws.Client) {
	list := server.ListOfGames()
	reply := model.CreateMessageListOfGames(list)
	b, _ := json.Marshal(&reply)
	log.Println(">> Sending list of games ")
	client.Send <- b

}
