package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/angrychess/io/ws"
	"github.com/team142/angrychess/model"
	"log"
)

func handleInMessageStartGame(server *model.Server, client *ws.Client) {
	startGameByClient(server, client)
}

func handleInMessageNick(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.MessageNick
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}
	setNick(server, client, message.Nick)
}

func handleInMessageMove(server *model.Server, client *ws.Client, msg *[]byte) {
	message := &model.MessageMove{}
	if err := json.Unmarshal(*msg, message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	move(server, message, client)
}

func handleInMessageListOfGame(server *model.Server, client *ws.Client) {
	go func() {
		reply := server.CreateMessageListOfGames()
		log.Println(">> Sending list of games ")
		client.SendObject(reply)
	}()
}

func handleInMessageJoinGame(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.MessageJoinGame
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}
	joinGameByClient(server, message.ID, server.Lobby[client])
	notifyLobby(server)
}

func handleInMessageDC(server *model.Server, client *ws.Client) {
	disconnect(server, client)
}

func handleInMessageCreateGame(server *model.Server, client *ws.Client) {
	createGameByClient(server, client)
	notifyLobby(server)
}

func handleInMessageChangeSeat(server *model.Server, client *ws.Client, msg *[]byte) {
	var message model.MessageChangeSeat
	if err := json.Unmarshal(*msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	changeSeatByClient(server, client, message.Seat)
}
