package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	service *service.RoomService
}

func NewRoomHandler(service *service.RoomService) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}

func (h *RoomHandler) FetchAll(c *gin.Context) {
	roomRess, errs := h.service.FetchAll()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(402))
		return
	}

	res := response.Make(200)
	res["data"] = roomRess
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) FindEmpty(c *gin.Context) {
	rid, errs := h.service.FindEmpty()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(415))
		return
	}

	res := response.Make(209)
	res["data"] = rid
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) Join(c *gin.Context) {
	uid, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	var roomReq dtos.RoomRequest
	if err := c.ShouldBindJSON(&roomReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	// TODO user leave all room

	if errs := h.service.Join(uid.(int), roomReq.ID); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(416))
		return
	}

	c.JSON(http.StatusOK, response.Make(210))
}

func (h *RoomHandler) Leave(c *gin.Context) {
	uid, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	var roomReq dtos.RoomRequest
	if err := c.ShouldBindJSON(&roomReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	// to leave room, user must join that room
	if errs := h.service.Check(uid.(int), roomReq.ID); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(418))
		return
	}

	if errs := h.service.Leave(uid.(int), roomReq.ID); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(417))
		return
	}

	c.JSON(http.StatusOK, response.Make(211))
}

func (h *RoomHandler) SendMsg(c *gin.Context) {
	uid, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	var roomReq dtos.RoomRequest
	if err := c.ShouldBindJSON(&roomReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	if errs := h.service.Check(uid.(int), roomReq.ID); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(418))
		return
	}

	msgRes, err := h.service.SendMsg(roomReq.ID, 30e9)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.Make(420))
		return
	}

	res := response.Make(213)
	res["data"] = msgRes
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) ReceiveMsg(c *gin.Context) {
	uid, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	var msgReq dtos.MessageRequest
	if err := c.ShouldBindJSON(&msgReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	if errs := h.service.Check(uid.(int), msgReq.Rid); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(418))
		return
	}

	fmt.Println(uid.(int), msgReq)

	if errs := h.service.ReceiveMsg(uid.(int), msgReq); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(419))
		return
	}

	fmt.Println("Hello")

	c.JSON(http.StatusOK, response.Make(212))
}
