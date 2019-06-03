package service

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	Fullname string `json:"fullname"`
	Body     string `json:"body"`
}

type Client struct {
	Room *Room
	Conn *websocket.Conn
	ID   int
}

func (c *Client) Read() {
	defer func() {
		c.Room.Unregister <- c
		if err := c.Conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		var message Message
		if err := c.Conn.ReadJSON(&message); err != nil {
			log.Println(err)
			return
		}
		c.Room.Broadcast <- message
	}
}
