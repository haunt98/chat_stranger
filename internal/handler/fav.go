package handler

import (
	"log"
	"net/http"

	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	service *service.FavoriteService
}

func NewFavoriteHandler(service *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		service: service,
	}
}

func (h *FavoriteHandler) FetchAll(c *gin.Context) {
	favRess, errs := h.service.FetchAll()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(402))
		return
	}

	res := response.Make(200)
	res["data"] = favRess
	c.JSON(http.StatusOK, res)
}
