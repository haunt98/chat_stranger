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
	conn *websocket.Conn
	room *Room
	id   int
}

func NewClient(conn *websocket.Conn, room *Room, id int) {
	c := &Client{
		conn: conn,
		room: room,
		id:   id,
	}
	c.room.register <- c

	go c.Read()
}

func (c *Client) Read() {
	defer func() {
		c.room.unregister <- c
		if err := c.conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	for {
		var message Message
		if err := c.conn.ReadJSON(&message); err != nil {
			log.Println(err)
			return
		}
		c.room.broadcast <- message
	}
}
