package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/1612180/chat_stranger/internal/models"
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
	users, errs := h.service.FetchAll()
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(402))
		return
	}

	res := Response(200)
	res["users"] = users
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Find(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	user, errs := h.service.Find(uint(id))
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	res := Response(201)
	res["user"] = user
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Create(c *gin.Context) {
	var upload models.UserUpload
	if err := c.ShouldBindJSON(&upload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	id, errs := h.service.Create(&upload)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(404))
		return
	}

	res := Response(205)
	res["userid"] = id
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) UpdateInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	var upload models.UserUpload
	if err = c.ShouldBindJSON(&upload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := h.service.UpdateInfo(uint(id), &upload); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(202))
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(401))
	}

	var auth models.Authentication
	if err = c.ShouldBindJSON(&auth); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := h.service.UpdatePassword(uint(id), &auth); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(203))
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(401))
		return
	}

	if errs := h.service.Delete(uint(id)); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(204))
}

func (h *UserHandler) Authenticate(c *gin.Context) {
	var auth models.Authentication
	if err := c.ShouldBindJSON(&auth); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	user, errs := h.service.Authenticate(&auth)
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(405))
		return
	}

	s, err := service.CreateTokenString(models.JWTClaims{
		ID:             user.ID,
		Role:           "User",
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Response(500))
		return
	}

	res := Response(206)
	res["token"] = s
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) VerifyFind(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(501))
		return
	}

	user, errs := h.service.Find(id.(uint))
	if len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	res := Response(201)
	res["user"] = user
	c.JSON(http.StatusOK, res)
}

func (h *UserHandler) VerifyDelete(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusBadRequest, Response(501))
		return
	}

	if errs := h.service.Delete(id.(uint)); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(204))
}

func (h *UserHandler) VerifyUpdateInfo(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusInternalServerError, Response(501))
		return
	}

	var upload models.UserUpload
	if err := c.ShouldBindJSON(&upload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := h.service.UpdateInfo(id.(uint), &upload); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(202))
}

func (h *UserHandler) VerifyUpdatePassword(c *gin.Context) {
	id, ok := c.Get("ID")
	if !ok {
		c.JSON(http.StatusInternalServerError, Response(501))
		return
	}

	var auth models.Authentication
	if err := c.ShouldBindJSON(&auth); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response(400))
		return
	}

	if errs := h.service.UpdatePassword(id.(uint), &auth); len(errs) != 0 {
		for _, err := range errs {
			log.Println(err)
		}
		c.JSON(http.StatusOK, Response(403))
		return
	}

	c.JSON(http.StatusOK, Response(203))
}
