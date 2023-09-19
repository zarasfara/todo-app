package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	usersList, err := h.services.User.GetAll()
	if err != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"users": usersList,
	})
}

func (h *Handler) createUser(c *gin.Context) {
	var user entities.User

	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	h.services.User.CreateUser(user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "user has been created!",
	})
}
