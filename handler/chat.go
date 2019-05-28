package handler

import (
	"github.com/1612180/chat_stranger/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
)

type Message struct {
	Fullname string `json:"fullname"`
	Body     string `json:"body"`
}

type Client struct {
	Conn *websocket.Conn
	Room *Room
}

func (client *Client) Read() {
	defer func() {
		if err := client.Conn.Close(); err != nil {
			log.ServerLog(err)
		}
		client.Room.Unregister <- client
	}()

	for {
		var message Message
		err := client.Conn.ReadJSON(&message)
		if err != nil {
			log.ServerLog(err)
			return
		}
		client.Room.Broadcast <- message
	}
}

type Room struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewRoom() *Room {
	return &Room{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (room *Room) Start() {
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
			for client := range room.Clients {
				if err := client.Conn.WriteJSON(Message{
					Fullname: "Server",
					Body:     "New User Joined...",
				}); err != nil {
					log.ServerLog(err)
				}
			}
			break
		case client := <-room.Unregister:
			delete(room.Clients, client)
			for client := range room.Clients {
				if err := client.Conn.WriteJSON(Message{
					Fullname: "Server",
					Body:     "User Disconnected...",
				}); err != nil {
					log.ServerLog(err)
				}
			}
			break
		case message := <-room.Broadcast:
			for client := range room.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.ServerLog(err)
					return
				}
			}
		}
	}
}

type Hub struct {
	Rooms map[int]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[int]*Room),
	}
}

func (hub *Hub) NewRoom() int {
	roomid := len(hub.Rooms) + 1
	hub.Rooms[roomid] = NewRoom()
	go hub.Rooms[roomid].Start()
	return roomid
}

var upgrader = websocket.Upgrader{}

func (hub *Hub) ChatHandler(c *gin.Context) {
	queryRoomID := c.Query("roomid")
	RoomID, _ := strconv.Atoi(queryRoomID)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.ServerLog(err)
		return
	}

	client := &Client{
		Conn: conn,
		Room: hub.Rooms[RoomID],
	}

	hub.Rooms[RoomID].Register <- client
	client.Read()
}
