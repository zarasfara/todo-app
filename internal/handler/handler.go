package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platoform/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.Use(jsonHeaderCheckMiddleware())
		
		v1 := api.Group("/v1")
		{
			v1.GET("/check",h.check)

			users := v1.Group("/users")
			{
				users.GET("/", h.getAllUsers)
			}
		}

	}

	return router
}
