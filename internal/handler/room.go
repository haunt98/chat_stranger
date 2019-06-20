package handler

import (
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
	roomid, errs := h.service.FindEmpty()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(415))
		return
	}

	res := response.Make(209)
	res["data"] = roomid
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) NextEmpty(c *gin.Context) {
	userid, ok := c.Get("id")
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

	roomid, errs := h.service.NextEmpty(userid.(int), roomReq.ID)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(423))
		return
	}

	res := response.Make(214)
	res["data"] = roomid
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) Join(c *gin.Context) {
	userid, ok := c.Get("id")
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

	if errs := h.service.Join(userid.(int), roomReq.ID); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(416))
		return
	}

	c.JSON(http.StatusOK, response.Make(210))
}

func (h *RoomHandler) Leave(c *gin.Context) {
	userid, ok := c.Get("id")
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

	if errs := h.service.Leave(userid.(int), roomReq.ID); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(417))
		return
	}

	c.JSON(http.StatusOK, response.Make(211))
}

func (h *RoomHandler) SendLatestMsg(c *gin.Context) {
	userid, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	var latestReq dtos.LatestRequest
	if err := c.ShouldBindJSON(&latestReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	msgRes, newLatest, err := h.service.SendLatestMsg(userid.(int), latestReq.RoomID, latestReq.Latest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, response.Make(420))
		return
	}

	// already have latest msg or any user join or leave
	if msgRes == nil {
		res := response.Make(213)
		res["latest"] = newLatest
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.Make(213)
	res["data"] = msgRes
	res["latest"] = newLatest
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) ReceiveMsg(c *gin.Context) {
	userid, ok := c.Get("id")
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

	if errs := h.service.ReceiveMsg(userid.(int), &msgReq); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(419))
		return
	}

	c.JSON(http.StatusOK, response.Make(212))
}
