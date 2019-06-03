package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Fullname string `json:"fullname"`
	Body     string `json:"body"`
}

type Client struct {
	Room *Room
	Conn *websocket.Conn
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

type Room struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	mux        sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (r *Room) Run() {
	for {
		select {
		case c := <-r.Register:
			for oldC := range r.Clients {
				if err := oldC.Conn.WriteJSON(Message{
					Fullname: "Server",
					Body:     "User has joined",
				}); err != nil {
					log.Println(err)
					delete(r.Clients, oldC)
					if err := oldC.Conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
			r.Clients[c] = true
		case c := <-r.Unregister:
			delete(r.Clients, c)
			if err := c.Conn.Close(); err != nil {
				log.Println(err)
			}
			for oldC := range r.Clients {
				if err := oldC.Conn.WriteJSON(Message{
					Fullname: "Server",
					Body:     "User has leaved",
				}); err != nil {
					log.Println(err)
					delete(r.Clients, oldC)
					if err := oldC.Conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
		case message := <-r.Broadcast:
			for c := range r.Clients {
				if err := c.Conn.WriteJSON(message); err != nil {
					log.Println(err)
					delete(r.Clients, c)
					if err := c.Conn.Close(); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func (r *Room) IsFull() bool {
	r.mux.Lock()
	defer r.mux.Unlock()

	return len(r.Clients) >= 2
}

type ChatHandler struct {
	Rooms map[int]*Room
	mux   sync.Mutex
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		Rooms: make(map[int]*Room),
	}
}

func (h *ChatHandler) NextRoom(cur int) int {
	h.mux.Lock()
	defer h.mux.Unlock()

	for id := range h.Rooms {
		if !h.Rooms[id].IsFull() && id != cur {
			return id
		}
	}

	id := len(h.Rooms) + 1
	h.Rooms[id] = NewRoom()

	go h.Rooms[id].Run()

	return id
}

func (h *ChatHandler) JoinRoom() int {
	h.mux.Lock()
	defer h.mux.Unlock()

	for id := range h.Rooms {
		if !h.Rooms[id].IsFull() {
			return id
		}
	}

	id := len(h.Rooms) + 1
	h.Rooms[id] = NewRoom()

	go h.Rooms[id].Run()

	return id
}

var upgrader = websocket.Upgrader{}

func (h *ChatHandler) WS(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err)
		return
	}

	if _, ok := h.Rooms[id]; !ok {
		log.Println(fmt.Errorf(ResponseCode[411]))
		return
	}
	if h.Rooms[id].IsFull() {
		log.Println(fmt.Errorf(ResponseCode[410]))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		Conn: conn,
		Room: h.Rooms[id],
	}
	client.Room.Register <- client

	go client.Read()
}

func (h *ChatHandler) FindRoom(c *gin.Context) {
	q := c.Query("id")
	if q == "" {
		res := Response(207)
		res["room"] = h.JoinRoom()
		c.JSON(http.StatusOK, res)
	} else {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, Response(401))
			return
		}

		res := Response(208)
		res["room"] = h.NextRoom(id)
		c.JSON(http.StatusOK, res)
	}
}
