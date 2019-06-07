package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/1612180/chat_stranger/internal/dtos"
	"github.com/1612180/chat_stranger/internal/pkg/response"
	"github.com/1612180/chat_stranger/internal/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service *service.AdminService
}

func NewAdminHandler(s *service.AdminService) *AdminHandler {
	return &AdminHandler{
		service: s,
	}
}

func (h *AdminHandler) FetchAll(c *gin.Context) {
	adminRess, errs := h.service.FetchAll()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Response(402))
		return
	}

	res := response.Response(200)
	res["data"] = adminRess
	c.JSON(http.StatusOK, res)
}

func (h *AdminHandler) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Response(401))
		return
	}

	adminRes, errs := h.service.Find(id)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Response(403))
		return
	}

	res := response.Response(201)
	res["data"] = adminRes
	c.JSON(http.StatusOK, res)
}

func (h *AdminHandler) Create(c *gin.Context) {
	var adminReq dtos.AdminRequest
	if err := c.ShouldBindJSON(&adminReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Response(400))
		return
	}

	id, errs := h.service.Create(&adminReq)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Response(401))
		return
	}

	res := response.Response(205)
	res["id"] = id
	c.JSON(http.StatusOK, res)
}

func (h *AdminHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Response(401))
		return
	}

	var adminReq dtos.AdminRequest
	if err = c.ShouldBindJSON(&adminReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Response(400))
		return
	}

	if errs := h.service.UpdateInfo(id, &adminReq); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Response(403))
		return
	}

	c.JSON(http.StatusOK, response.Response(202))
}

func (h *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Response(401))
		return
	}

	if errs := h.service.Delete(id); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Response(403))
		return
	}

	c.JSON(http.StatusOK, response.Response(204))
}

func (h *AdminHandler) Authenticate(c *gin.Context) {
	var creReq dtos.CredentialRequest
	if err := c.ShouldBindJSON(&creReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Response(400))
		return
	}

	adminRes, errs := h.service.Authenticate(&creReq)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Response(405))
		return
	}

	s, err := service.CreateTokenString(service.JWTClaims{
		ID:             adminRes.ID,
		Role:           "admin",
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.Response(500))
		return
	}

	res := response.Response(206)
	res["token"] = s
	c.JSON(http.StatusOK, res)
}
