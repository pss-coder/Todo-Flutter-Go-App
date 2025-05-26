package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-backend/models"
	"todo-backend/store"

	"github.com/gin-gonic/gin"
)

// GetTodos godoc
// @Summary      Get all todos
// @Description  get all todos (requires login)
// @Tags         todos
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}  models.Todo
// @Router       /todos [get]
func GetTodos(store store.TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos := store.GetTodos()
		c.JSON(http.StatusOK, todos)
	}
}

// AddTodo godoc
// @Summary      Add a new todo
// @Description  Add a new todo item (requires login)
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        todo  body      models.Todo  true  "Todo to add"
// @Success      201   {object}  models.Todo
// @Router       /todos [post]
func AddTodo(store store.TodoStore, ws *models.WebSockethub) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// print todo
		fmt.Println("Todo: ", todo)

		todo = store.AddTodo(todo)
		c.JSON(http.StatusCreated, todo)
		ws.NotifyClients(store.GetTodos())
	}
}

// ToggleTodo godoc
// @Summary      Toggle a todo's completion
// @Description  Toggle the completed state of a todo by ID (requires login)
// @Tags         todos
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id    path      int  true  "Todo ID"
// @Success      200   {object}  models.Todo
// @Router       /todos/{id} [put]
func ToggleTodo(store store.TodoStore, ws *models.WebSockethub) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// convert id to uint
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		todo := store.ToggleTodo(uint(idUint))
		c.JSON(http.StatusOK, todo)
		fmt.Println("Todo: ", todo)

		ws.NotifyClients(store.GetTodos())
	}
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Delete a todo by ID (requires login)
// @Tags         todos
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id    path      int  true  "Todo ID"
// @Success      200   {object}  models.Todo
// @Router       /todos/{id} [delete]
func DeleteTodo(store store.TodoStore, ws *models.WebSockethub) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// convert id to uint
		idUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		todo := store.DeleteTodo(uint(idUint))
		c.JSON(http.StatusOK, todo)

		ws.NotifyClients(store.GetTodos())
	}
}
