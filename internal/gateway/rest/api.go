package rest

import (
	"net/http"

	"github.com/diyor200/uof/internal/domains"
	"github.com/diyor200/uof/internal/usecase/users"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
	uc     users.UserUseCase
}

func NewHandler(uc users.UserUseCase) *Handler {
	return &Handler{uc: uc}
}

type addUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Handler) AddUser(c *gin.Context) {
	var data addUserRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.AddUser(c, domains.User{
		Name:   data.Name,
		Email:  data.Email,
		Status: domains.UserStatusCreated,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}
