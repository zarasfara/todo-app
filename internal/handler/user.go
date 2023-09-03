package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
