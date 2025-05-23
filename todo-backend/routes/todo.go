package routes

import (
	"todo-backend/controllers"
	"todo-backend/models"
	"todo-backend/store"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine, store store.TodoStore, wsHub *models.WebSockethub) {
	r.GET("/todos", controllers.GetTodos(store))
	r.POST("/todos", controllers.AddTodo(store, wsHub))
	r.PUT("/todos/:id", controllers.ToggleTodo(store, wsHub))
	r.DELETE("/todos/:id", controllers.DeleteTodo(store, wsHub))
}
