package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type ChatHandler struct {
	service *service.RoomService
}

func NewChatHandler(service *service.RoomService) *ChatHandler {
	return &ChatHandler{
		service: service,
	}
}

func (h *ChatHandler) WS(c *gin.Context) {
	// room id
	rid, err := strconv.Atoi(c.Query("rid"))
	if err != nil {
		log.Println(err)
		return
	}

	// user id
	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		log.Println(err)
		return
	}

	if err := h.service.CheckRoom(rid); err != nil {
		log.Println(err)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	service.NewClient(conn, h.service.GetRoom(rid), uid)
}

func (h *ChatHandler) FindRoom(c *gin.Context) {
	_, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, response.Response(501))
		return
	}

	q := c.Query("rid")
	if q == "" {
		res := response.Response(207)
		res["room"] = h.service.JoinRoom()
		c.JSON(http.StatusOK, res)
	} else {
		rid, err := strconv.Atoi(q)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, response.Response(401))
			return
		}

		res := response.Response(208)
		res["room"] = h.service.NextRoom(rid)
		c.JSON(http.StatusOK, res)
	}
}
