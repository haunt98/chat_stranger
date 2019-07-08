package handler

import (
	"strconv"
	"time"

	"github.com/1612180/chat_stranger/internal/model"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoomHandler struct {
	roomService *service.RoomService
}

func NewRoomHandler(roomService *service.RoomService) *RoomHandler {
	return &RoomHandler{roomService: roomService}
}

func (h *RoomHandler) Find(c *gin.Context) { // get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	status := c.Query("status")
	room, ok := h.roomService.Find(userID.(int), status)
	if !ok {
		c.JSON(200, response.Create(40))
		return
	}
	c.JSON(200, response.CreateWithData(4, room))
}

func (h *RoomHandler) Join(c *gin.Context) {
	// get room
	roomID, err := strconv.Atoi(c.Query("roomID"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "room",
			"action": "join",
		}).Error(err)
		c.JSON(200, response.Create(51))
		return
	}

	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	if ok := h.roomService.Join(userID.(int), roomID); !ok {
		c.JSON(200, response.Create(50))
		return
	}
	c.JSON(200, response.Create(5))
}

func (h *RoomHandler) Leave(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	if ok := h.roomService.Leave(userID.(int)); !ok {
		c.JSON(200, response.Create(60))
		return
	}
	c.JSON(200, response.Create(6))
}

func (h *RoomHandler) SendMessage(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	var message model.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "room",
			"action": "send message",
		}).Error(err)
		c.JSON(200, response.Create(71))
		return
	}

	message.UserID = userID.(int)
	if ok := h.roomService.SendMessage(&message); !ok {
		c.JSON(200, response.Create(70))
		return
	}
	c.JSON(200, response.Create(7))
}

func (h *RoomHandler) ReceiveMessage(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// get time
	fromTime, err := time.Parse(time.RFC3339, c.Query("fromTime"))
	if err != nil {
		c.JSON(200, response.Create(81))
		return
	}

	messages, ok := h.roomService.ReceiveMessage(userID.(int), fromTime)
	if !ok {
		c.JSON(200, response.Create(80))
		return
	}
	c.JSON(200, response.CreateWithData(8, messages))
}

func (h *RoomHandler) IsUserFree(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	if ok := h.roomService.IsUserFree(userID.(int)); !ok {
		c.JSON(200, response.Create(90))
		return
	}
	c.JSON(200, response.Create(9))
}

func (h *RoomHandler) CountMember(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	count, ok := h.roomService.CountMember(userID.(int))
	if !ok {
		c.JSON(200, response.Create(9010))
	}
	c.JSON(200, response.CreateWithData(901, count))
}
