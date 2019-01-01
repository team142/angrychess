package controller

import (
	"encoding/json"
	"fmt"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"github.com/team142/chessfor4/model/messages"
	"log"
)

func handleInMessageNick(server *model.Server, client *ws.Client, msg []byte) {
	var message messages.MessageNick
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	profile := server.GetOrCreateProfile(client)
	profile.Nick = message.Nick

	reply := messages.CreateMessageSecret(profile.Secret)
	b, err := json.Marshal(reply)
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling, %s", err))
	}
	client.Send <- b

}
