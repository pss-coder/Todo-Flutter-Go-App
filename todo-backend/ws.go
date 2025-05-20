package todobackend

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketHub struct {
	clients   map[*websocket.Conn]bool // map of active ws connections
	broadcast chan Todos               // channel to send our todo list to all clients
	mu        sync.Mutex               // to protect shared access to clients
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Todos),
	}
}

func (hub *WebSocketHub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// convert incoming http request to websocket request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	hub.mu.Lock()
	hub.clients[conn] = true
	hub.mu.Unlock()

	go func() {
		defer func() {
			hub.mu.Lock()
			delete(hub.clients, conn)
			hub.mu.Unlock()
			conn.Close()
		}()
		for {
			if _, _, err := conn.NextReader(); err != nil {
				break
			}
		}
	}()
}

func (hub *WebSocketHub) StartBroadcasting() {
	for todos := range hub.broadcast {
		data, _ := json.Marshal(todos)
		hub.mu.Lock()
		for conn := range hub.clients {
			conn.WriteMessage(websocket.TextMessage, data)
		}
		hub.mu.Unlock()
	}
}

func (hub *WebSocketHub) NotifyClients(todos []Todo) {
	go func() {
		hub.broadcast <- todos
	}()
}
