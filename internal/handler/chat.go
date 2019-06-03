package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type ChatHandler struct {
	Rooms map[int]*service.Room
	mux   sync.Mutex
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		Rooms: make(map[int]*service.Room),
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
	h.Rooms[id] = service.NewRoom()

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
	h.Rooms[id] = service.NewRoom()

	go h.Rooms[id].Run()

	return id
}

func (h *ChatHandler) WS(c *gin.Context) {
	rid, err := strconv.Atoi(c.Query("rid"))
	if err != nil {
		log.Println(err)
		return
	}

	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		log.Println(err)
		return
	}

	if _, ok := h.Rooms[rid]; !ok {
		log.Println(fmt.Errorf(ResponseCode[411]))
		return
	}
	if h.Rooms[rid].IsFull() {
		log.Println(fmt.Errorf(ResponseCode[410]))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &service.Client{
		Conn: conn,
		Room: h.Rooms[rid],
		ID:   uid,
	}
	client.Room.Register <- client

	go client.Read()
}

func (h *ChatHandler) FindRoom(c *gin.Context) {
	_, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusInternalServerError, Response(501))
		return
	}

	q := c.Query("rid")
	if q == "" {
		res := Response(207)
		res["room"] = h.JoinRoom()
		c.JSON(http.StatusOK, res)
	} else {
		rid, err := strconv.Atoi(q)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, Response(401))
			return
		}

		res := Response(208)
		res["room"] = h.NextRoom(rid)
		c.JSON(http.StatusOK, res)
	}
}
