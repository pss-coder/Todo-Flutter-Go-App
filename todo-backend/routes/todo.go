package routes

import (
	"todo-backend/controllers"
	"todo-backend/store"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine, store store.TodoStore) {
	r.GET("/todos", controllers.GetTodos(store))
	r.POST("/todos", controllers.AddTodo(store))
	r.PUT("/todos/:id", controllers.ToggleTodo(store))
	r.DELETE("/todos/:id", controllers.DeleteTodo(store))
}
