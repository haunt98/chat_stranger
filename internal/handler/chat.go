package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/pkg/variable"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ChatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatServiceI service.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatServiceI}
}

type MessageSubmit struct {
	Body string `json:"body,omitempty"`
}

func (h *ChatHandler) FindRoom(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: find room failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: find room failed"))
		c.JSON(403, response.Create(999))
		return
	}

	status := c.Query(variable.StatusRoom)
	if status == variable.AnyRoom {
		room, err := h.chatService.FindAnyRoom(userID)
		if err != nil {
			logrus.Error(errors.Wrap(err, "account handler: find room failed"))
			c.JSON(500, response.Create(401))
			return
		}
		c.JSON(200, response.CreateWithData(400, room))
		return
	} else if status == variable.NextRoom {
		room, err := h.chatService.FindNextRoom(userID)
		if err != nil {
			logrus.Error(errors.Wrap(err, "account handler: find room failed"))
			c.JSON(500, response.Create(402))
			return
		}
		c.JSON(200, response.CreateWithData(400, room))
		return
	} else if status == variable.SameGenderRoom {
		room, err := h.chatService.FindSameGenderRoom(userID)
		if err != nil {
			logrus.Error(errors.Wrap(err, "account handler: find room failed"))
			c.JSON(500, response.Create(403))
			return
		}
		c.JSON(200, response.CreateWithData(400, room))
		return
	} else if status == variable.SameBirthYearRoom {
		room, err := h.chatService.FindSameBirthYearRoom(userID)
		if err != nil {
			logrus.Error(errors.Wrap(err, "account handler: find room failed"))
			c.JSON(500, response.Create(404))
			return
		}
		c.JSON(200, response.CreateWithData(400, room))
		return
	}
	c.JSON(400, response.Create(405))
}

func (h *ChatHandler) Join(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: join room failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: join room failed"))
		c.JSON(403, response.Create(999))
		return
	}

	roomID, err := strconv.Atoi(c.Query("roomID"))
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: join room failed"))
		c.JSON(400, response.Create(502))
		return
	}

	if err := h.chatService.Join(userID, roomID); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: join room failed"))
		c.JSON(500, response.Create(501))
		return
	}
	c.JSON(200, response.Create(500))
}

func (h *ChatHandler) Leave(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: leave room failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: leave room failed"))
		c.JSON(403, response.Create(999))
		return
	}

	if err := h.chatService.Leave(userID); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: leave room failed"))
		c.JSON(500, response.Create(601))
		return
	}
	c.JSON(200, response.Create(600))
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: send message failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: send message failed"))
		c.JSON(403, response.Create(999))
		return
	}

	var submit MessageSubmit
	if err := c.ShouldBindJSON(&submit); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: send message failed"))
		c.JSON(400, response.Create(702))
		return
	}

	if err := h.chatService.SendMessage(userID, submit.Body); err != nil {
		logrus.Error(errors.Wrap(err, "account handler: send message failed"))
		c.JSON(500, response.Create(701))
		return
	}
	c.JSON(200, response.Create(700))
}

func (h *ChatHandler) ReceiveMessage(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: receive message failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: receive message failed"))
		c.JSON(403, response.Create(999))
		return
	}

	from, err := time.Parse(time.RFC3339, c.Query(variable.FromTime))
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: receive message failed"))
		c.JSON(400, response.Create(802))
		return
	}

	msgs, err := h.chatService.ReceiveMessage(userID, from)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: receive message failed"))
		c.JSON(500, response.Create(801))
		return
	}
	c.JSON(200, response.CreateWithData(800, msgs))
}

func (h *ChatHandler) IsUserFree(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: is user free failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: is user free failed"))
		c.JSON(403, response.Create(999))
		return
	}

	yes, err := h.chatService.IsUserFree(userID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: is user free failed"))
		c.JSON(500, response.Create(902))
		return
	}
	if !yes {
		c.JSON(200, response.Create(901))
		return
	}
	c.JSON(200, response.Create(900))
}

func (h *ChatHandler) CountMember(c *gin.Context) {
	tempUserID, ok := c.Get("userID")
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("userID empty"), "account handler: count member failed"))
		c.JSON(403, response.Create(999))
		return
	}

	userID, ok := tempUserID.(int)
	if !ok {
		logrus.Error(errors.Wrap(fmt.Errorf("type assertion failed"), "account handler: count member failed"))
		c.JSON(403, response.Create(999))
		return
	}

	count, err := h.chatService.CountMembersInRoomOfUser(userID)
	if err != nil {
		logrus.Error(errors.Wrap(err, "account handler: count member failed"))
		c.JSON(500, response.Create(111))
	}
	c.JSON(200, response.CreateWithData(110, count))
}
