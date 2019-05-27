package handler

import (
	"github.com/1612180/chat_stranger/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Client struct {
	Conn *websocket.Conn
	Hub  *Hub
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (client *Client) Read() {
	defer func() {
		if err := client.Conn.Close(); err != nil {
			log.ServerLog(err)
		}
		client.Hub.Unregister <- client
	}()

	for {
		messageType, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.ServerLog(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		client.Hub.Broadcast <- message
	}
}

type Hub struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (hub *Hub) Start() {
	for {
		select {
		case client := <-hub.Register:
			hub.Clients[client] = true
			for client := range hub.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."}); err != nil {
					log.ServerLog(err)
				}
			}
			break
		case client := <-hub.Unregister:
			delete(hub.Clients, client)
			for client := range hub.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."}); err != nil {
					log.ServerLog(err)
				}
			}
			break
		case message := <-hub.Broadcast:
			for client := range hub.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.ServerLog(err)
					return
				}
			}
		}
	}
}

func (hub *Hub) ChatHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.ServerLog(err)
		return
	}

	client := &Client{
		Conn: conn,
		Hub:  hub,
	}

	hub.Register <- client
	client.Read()
}
