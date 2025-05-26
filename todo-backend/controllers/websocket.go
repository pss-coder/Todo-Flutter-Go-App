package controllers

import (
	"todo-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketHandler(hub *models.WebSockethub) gin.HandlerFunc {
	return func(c *gin.Context) {

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to upgrade connection"})
			return
		}
		defer conn.Close()

		hub.Mu.Lock()
		hub.Clients[conn] = true
		hub.Mu.Unlock()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				hub.Mu.Lock()
				delete(hub.Clients, conn)
				hub.Mu.Unlock()
				break
			}
		}
	}
}
