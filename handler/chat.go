package handler

import (
	"encoding/json"
	"fmt"
	"github.com/1612180/chat_stranger/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
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
func (client *Client) readPump() {
	defer func() {
		client.Room.Unregister <- client
		if err := client.Conn.Close(); err != nil {
			log.ServerLog(err)
		}
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	if err := client.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.ServerLog(err)
	}
	client.Conn.SetPongHandler(func(string) error {
		if err := client.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.ServerLog(err)
		}
		return nil
	})
	for {
		var message Message
		err := client.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.ServerLog(err)
			}
			break
		}
		client.Room.Broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := client.Conn.Close(); err != nil {
			log.ServerLog(err)
		}
	}()
	for {
		select {
		case message, ok := <-client.Send:
			if err := client.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.ServerLog(err)
			}
			if !ok {
				// The hub closed the channel.
				if err := client.Conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.ServerLog(err)
				}
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.ServerLog(err)
				return
			}
			if err := json.NewEncoder(w).Encode(message); err != nil {
				return
			}

			//if err := client.Conn.WriteJSON(message); err != nil {
			//	log.ServerLog(err)
			//	return
			//}

			// Add queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				if err := json.NewEncoder(w).Encode(<-client.Send); err != nil {
					log.ServerLog(err)
					return
				}
				//if err := client.Conn.WriteJSON(<-client.Send); err != nil {
				//	log.ServerLog(err)
				//	return
				//}
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := client.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.ServerLog(err)
			}
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
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

func (room *Room) Run() {
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
		case client := <-room.Unregister:
			if _, ok := room.Clients[client]; ok {
				delete(room.Clients, client)
				close(client.Send)
			}
		case message := <-room.Broadcast:
			for client := range room.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(room.Clients, client)
				}
			}
		}
	}
}

func (room *Room) IsFull() bool {
	room.mux.Lock()
	defer room.mux.Unlock()

	return len(room.Clients) >= 2
}

type Hub struct {
	Rooms map[int]*Room
	mux   sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[int]*Room),
	}
}

func (hub *Hub) NewRoom() int {
	hub.mux.Lock()
	defer hub.mux.Unlock()

	roomid := len(hub.Rooms) + 1
	hub.Rooms[roomid] = NewRoom()

	go hub.Rooms[roomid].Run()

	return roomid
}

func (hub *Hub) GetAvailableRoom() (int, error) {
	hub.mux.Lock()
	defer hub.mux.Unlock()

	for roomid := range hub.Rooms {
		if !hub.Rooms[roomid].IsFull() {
			return roomid, nil
		}
	}

	return -1, fmt.Errorf("no room available")
}

func (hub *Hub) IsAvailable(roomid int) bool {
	hub.mux.Lock()
	defer hub.mux.Unlock()

	room, ok := hub.Rooms[roomid]
	if !ok {
		return false
	}

	return !room.IsFull()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (hub *Hub) ChatHandler(c *gin.Context) {
	roomid, err := strconv.Atoi(c.Query("roomid"))
	if err != nil {
		log.ServerLog(err)
		return
	}

	if !hub.IsAvailable(roomid) {
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.ServerLog(err)
		return
	}

	client := &Client{
		Conn: conn,
		Room: hub.Rooms[roomid],
		Send: make(chan Message),
	}
	client.Room.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
