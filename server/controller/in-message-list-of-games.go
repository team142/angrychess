package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func handleInMessageListOfGame(server *model.Server, client *ws.Client, msg []byte) {
	reply := server.CreateListOfGames()
	b, err := json.Marshal(&reply)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	log.Println(">> Sending list of games ")
	client.Send <- b

}
