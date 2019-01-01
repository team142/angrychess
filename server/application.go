package main

import (
	"flag"
	"github.com/team142/chessfor4/controller"
	"github.com/team142/chessfor4/io/ws"
	"log"
)

var addr = flag.String("addr", "127.0.0.1:8000", "http service address")

func main() {

	log.Println("Oh hai 🚀 Lets Go 🎠")

	//Blocking call
	ws.StartServer(addr, controller.HandleIncoming)

}
