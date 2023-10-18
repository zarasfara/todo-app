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
		// Register middlewares
		api.Use(jsonHeaderCheckMiddleware())
		api.Use(setJSONContentType())

		// Init routes
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("/", h.getAllUsers)
				users.POST("/", h.createUser)
			}

			todos := v1.Group("/todos")
			{
				todos.GET("/", h.getAllTodos)
				todos.POST("/", h.createTodo)
				todos.GET("/:id", h.getTodoById)
				todos.PUT("/:id", h.updateTodo)
				todos.PATCH("/change-status/:id", h.changeStatus)
				todos.DELETE("/:id", h.deleteTodo)
			}
		}

	}

	return router
}
