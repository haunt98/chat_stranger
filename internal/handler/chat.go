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

type ChatHandler struct {
	chatService *service.ChatService
}

func NewChatHandler(roomService *service.ChatService) *ChatHandler {
	return &ChatHandler{chatService: roomService}
}

func (h *ChatHandler) Find(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// query status
	status := c.Query("status")
	room, ok := h.chatService.Find(userID.(int), status)
	if !ok {
		c.JSON(200, response.Create(401))
		return
	}
	c.JSON(200, response.CreateWithData(400, room))
}

func (h *ChatHandler) Join(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// get room
	roomID, err := strconv.Atoi(c.Query("roomID"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "chat",
			"action": "join",
		}).Error(err)
		c.JSON(200, response.Create(502))
		return
	}

	if ok := h.chatService.Join(userID.(int), roomID); !ok {
		c.JSON(200, response.Create(501))
		return
	}
	c.JSON(200, response.Create(500))
}

func (h *ChatHandler) Leave(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	if ok := h.chatService.Leave(userID.(int)); !ok {
		c.JSON(200, response.Create(601))
		return
	}
	c.JSON(200, response.Create(600))
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// get message
	var message model.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "handler",
			"target": "chat",
			"action": "send message",
		}).Error(err)
		c.JSON(200, response.Create(702))
		return
	}

	message.UserID = userID.(int)
	if ok := h.chatService.SendMessage(&message); !ok {
		c.JSON(200, response.Create(701))
		return
	}
	c.JSON(200, response.Create(700))
}

func (h *ChatHandler) ReceiveMessage(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	// get time
	fromTime, err := time.Parse(time.RFC3339, c.Query("fromTime"))
	if err != nil {
		c.JSON(200, response.Create(802))
		return
	}

	messages, ok := h.chatService.ReceiveMessage(userID.(int), fromTime)
	if !ok {
		c.JSON(200, response.Create(801))
		return
	}
	c.JSON(200, response.CreateWithData(800, messages))
}

func (h *ChatHandler) IsUserFree(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	if ok := h.chatService.IsUserFree(userID.(int)); !ok {
		c.JSON(200, response.Create(901))
		return
	}
	c.JSON(200, response.Create(900))
}

func (h *ChatHandler) CountMember(c *gin.Context) {
	// get user
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(403, response.Create(999))
		return
	}

	count, ok := h.chatService.CountMember(userID.(int))
	if !ok {
		c.JSON(200, response.Create(111))
	}
	c.JSON(200, response.CreateWithData(110, count))
}
