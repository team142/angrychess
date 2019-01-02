package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func handleInMessageChangeSeat(server *model.Server, client *ws.Client, msg []byte) {
	var message messages.MessageChangeSeat
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	server.ChangeSeat(client, message.Seat)

}
