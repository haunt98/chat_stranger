package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// https://github.com/gorilla/websocket/tree/master/examples/chat
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Message struct {
	Fullname string `json:"fullname"`
	Body     string `json:"body"`
}

type Client struct {
	Room *Room
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan Message
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.Room.Unregister <- c
		if err := c.Conn.Close(); err != nil {
			log.Println(err)
		}
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
	}
	c.Conn.SetPongHandler(func(string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Println(err)
		}
		return nil
	})
	for {
		var message Message
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(err)
			}
			break
		}
		c.Room.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.Conn.Close(); err != nil {
			log.Println(err)
		}
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Println(err)
			}
			if !ok {
				// The hub closed the channel.
				if err := c.Conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println(err)
				}
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			if err := json.NewEncoder(w).Encode(message); err != nil {
				return
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				if err := json.NewEncoder(w).Encode(<-c.Send); err != nil {
					log.Println(err)
					return
				}
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Println(err)
			}
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
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
		case client := <-r.Register:
			r.Clients[client] = true
		case client := <-r.Unregister:
			if _, ok := r.Clients[client]; ok {
				delete(r.Clients, client)
				close(client.Send)
			}
		case message := <-r.Broadcast:
			for client := range r.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(r.Clients, client)
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *ChatHandler) WS(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Println(err)
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
		Send: make(chan Message),
	}
	client.Room.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
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
