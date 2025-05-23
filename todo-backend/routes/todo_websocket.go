package routes

import (
	"todo-backend/controllers"
	"todo-backend/middlewares"
	"todo-backend/models"

	"github.com/gin-gonic/gin"
)

func TodoWebSocketRoutes(r *gin.Engine, hub *models.WebSockethub) {
	r.Use(middlewares.IsAuthorised())
	{
		r.GET("/ws/todos", controllers.WebsocketHandler(hub))
	}

}
