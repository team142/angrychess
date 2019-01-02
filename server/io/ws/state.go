package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Hub keeps track of all connected clients
type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//Client describes a websocket client
type Client struct {
	Hub     *Hub
	conn    *websocket.Conn
	Send    chan []byte
	handler func(*Client, []byte)
}

func (c *Client) handleMessage(msg []byte) {
	go c.handler(c, msg) //TODO: is this right?
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Println("New client registered")
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)

				data := struct {
					Msg string `json:"msg"`
				}{
					"disconnect",
				}
				b, _ := json.Marshal(data)
				client.handler(client, b)
				log.Println("Client unregistered")
			}
		}
	}
}
