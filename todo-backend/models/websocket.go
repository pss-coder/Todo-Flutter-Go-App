package models

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSockethub struct {
	Clients   map[*websocket.Conn]bool
	broadcast chan Todos // Broadcast channel to send todos to all clients
	Mu        sync.Mutex
}

func NewWebSockethub() *WebSockethub {
	return &WebSockethub{
		Clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Todos),
	}
}

func (hub *WebSockethub) StartBroadcasting() {
	for todos := range hub.broadcast {
		data, _ := json.Marshal(todos)
		hub.Mu.Lock()
		for conn := range hub.Clients {
			conn.WriteMessage(websocket.TextMessage, data)
		}
		hub.Mu.Unlock()
	}
}

func (hub *WebSockethub) NotifyClients(todos Todos) {
	go func() {
		hub.broadcast <- todos
	}()
}
