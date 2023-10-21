package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarasfara/pet-adoption-platoform/internal/entities"
)

type jsonResult struct {
	Data interface{} `json:"data"`
}

// @Summary		List todos
// @Description	Get list of all todos
// @Tags		Todos
// @Accept		json
// @Produce		json
// @Success		200  {object}  jsonResult{data=[]entities.Todo}
// @Failure		500  {object}  handler.errorResponse
// @Router		/todos [get]
func (h *Handler) getAllTodos(c *gin.Context) {
	todoList, err := h.services.Todo.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoList,
	})
}

// @Summary		Create a todo
// @Description	Create a todo
// @Tags		Todos
// @Accept		json
// @Produce		json
// @Success		201  {object}  jsonResult{data=string}
// @Failure		400  {object}  handler.errorResponse
// @Router		/todos [post]
func (h *Handler) createTodo(c *gin.Context) {
	var todo entities.Todo

	if err := c.BindJSON(&todo); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.services.Todo.CreateTodo(todo)

	c.JSON(http.StatusCreated, gin.H{
		"data": "todo has been created!",
	})
}

// @Summary		Show a todo
// @Description	Get todo by ID
// @Tags		Todos
// @Accept		json
// @Produce		json
// @Param		id   path      int  true  "Todo ID"
// @Success		200  {object}  jsonResult{data=entities.Todo}
// @Failure		404  {object}  handler.errorResponse
// @Failure		500  {object}  handler.errorResponse
// @Router		/todos/{id} [get]
func (h *Handler) getTodoById(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid todo id param")
		return
	}

	todo, err := h.services.Todo.GetTodoByID(todoId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, "todo not found")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

// @Summary		Update a todo
// @Description	Update a todo by id
// @Tags		Todos
// @Accept		json
// @Produce		json
// @Success		200  {object}  jsonResult{data=[]entities.Todo}
// @Failure		400  {object}  handler.errorResponse
// @Failure		500  {object}  handler.errorResponse
// @Router		/todos/{id} [put]
func (h *Handler) updateTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid todo id param")
		return
	}

	var input entities.UpdateTodoInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Todo.UpdateTodo(todoId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Todo was updated!",
	})
}

// @Summary		Delete a todo
// @Description	Delete a todo by id
// @Tags		Todos
// @Accept		json
// @Produce		json
// @Param		id   path      int  true  "Todo ID"
// @Success		200  {object}  jsonResult{data=string}
// @Failure		400  {object}  handler.errorResponse
// @Failure		500  {object}  handler.errorResponse
// @Router		/todos/{id} [delete]
func (h *Handler) deleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid todo id param.")
		return
	}

	if err := h.services.Todo.Delete(todoId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Todo was deleted!",
	})
}

// @Summary		Change status
// @Description	Change status a todo by id
// @Tags		Todos
// @Accept		json
// @Produce		json
// @Param		id   path      int  true  "Todo ID"
// @Success		200  {object}  jsonResult{data=string}
// @Failure		400  {object}  handler.errorResponse
// @Failure		500  {object}  handler.errorResponse
// @Router		/todos/{id} [patch]
func (h *Handler) changeStatus(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid todo id param.")
		return
	}

	if err := h.services.Todo.ChangeStatus(todoId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "Todo was updated.",
	})

}
