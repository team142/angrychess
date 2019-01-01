package main

import (
	"flag"
	"github.com/team142/chessfor4/io/ws"
	"github.com/team142/chessfor4/model"
	"log"
)

var addr = flag.String("addr", "127.0.0.1:8000", "http service address")
var publicAddress = flag.String("public", "local", "Server public domain name")

func main() {

	log.Println("Oh hai 🚀 Lets Go 🎠")

	server := model.CreateServer(*publicAddress)

	//Blocking call
	ws.StartWSServer(addr, server.HandleMessage)

}
