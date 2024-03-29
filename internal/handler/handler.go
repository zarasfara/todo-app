package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/zarasfara/pet-adoption-platoform/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
