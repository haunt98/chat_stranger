package handler

import (
	"log"
	"net/http"

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
	if len(errs) != 0{
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(415))
	}

	res := response.Make(209)
	res["data"] = rid
	c.JSON(http.StatusOK, res)
}

func (h *RoomHandler) Join(c *gin.Context) {

}