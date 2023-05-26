package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platoform/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler {
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api", h.check)
	{
		api.GET("check")
	}

	return router
}
