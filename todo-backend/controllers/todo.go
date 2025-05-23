package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-backend/models"
	"todo-backend/store"

	"github.com/gin-gonic/gin"
)

func GetTodos(store store.TodoStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos := store.GetTodos()
		c.JSON(http.StatusOK, todos)
	}
}

func AddTodo(store store.TodoStore) gin.HandlerFunc {
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
	}
}

func ToggleTodo(store store.TodoStore) gin.HandlerFunc {
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
	}
}

func DeleteTodo(store store.TodoStore) gin.HandlerFunc {
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
	}
}
