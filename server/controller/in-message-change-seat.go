package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

func handleInMessageChangeSeat(server *model.Server, client *ws.Client, msg []byte) {
	var message model.MessageChangeSeat
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
		return
	}
	changeSeatByClient(server, client, message.Seat)

}
