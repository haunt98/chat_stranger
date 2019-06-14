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

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) FetchAll(c *gin.Context) {
	userRess, errs := h.service.FetchAll()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(402))
		return
	}

	res := response.Make(200)
	res["data"] = userRess
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(401))
		return
	}

	userRes, errs := h.service.Find(id)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(403))
		return
	}

	res := response.Make(201)
	res["data"] = userRes
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Create(c *gin.Context) {
	var userReq dtos.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	if code := userReq.Check(); code != 0 {
		log.Println(response.Codes[code])
		c.JSON(http.StatusOK, response.Make(code))
		return
	}

	id, errs := h.service.Create(&userReq)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(404))
		return
	}

	res := response.Make(205)
	res["id"] = id
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(401))
		return
	}

	var userReq dtos.UserRequest
	if err = c.ShouldBindJSON(&userReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	if errs := h.service.UpdateInfo(id, &userReq); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(403))
		return
	}

	c.JSON(http.StatusOK, response.Make(202))
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(401))
		return
	}

	if errs := h.service.Delete(id); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(403))
		return
	}

	c.JSON(http.StatusOK, response.Make(204))
}

func (h *UserHandler) Authenticate(c *gin.Context) {
	var auth dtos.CredentialRequest
	if err := c.ShouldBindJSON(&auth); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	userRes, errs := h.service.Authenticate(&auth)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(405))
		return
	}

	s, err := service.CreateTokenString(service.JWTClaims{
		ID:             userRes.ID,
		Role:           "user",
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.Make(500))
		return
	}

	res := response.Make(206)
	res["token"] = s
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) VerifyFind(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	userRes, errs := h.service.Find(id.(int))
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(403))
		return
	}

	res := response.Make(201)
	res["data"] = userRes
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) VerifyDelete(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, response.Make(501))
		return
	}

	if errs := h.service.Delete(id.(int)); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(403))
		return
	}

	c.JSON(http.StatusOK, response.Make(204))
}

func (h *UserHandler) VerifyUpdateInfo(c *gin.Context) {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, response.Make(501))
		return
	}

	var userReq dtos.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.Make(400))
		return
	}

	if errs := h.service.UpdateInfo(id.(int), &userReq); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, response.Make(403))
		return
	}

	c.JSON(http.StatusOK, response.Make(202))
}
