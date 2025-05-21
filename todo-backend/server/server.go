package todobackend

import (
	"encoding/json"
	"net/http"

	model "github.com/pss-coder/go-flutter-app-backend/todo/model"
	store "github.com/pss-coder/go-flutter-app-backend/todo/store"
)

type TodoServer struct {
	store store.TodoStore
	hub   *WebSocketHub // our websocket data here
	http.Handler
}

const jsonContentType = "application/json"

func NewTodoServer(store store.TodoStore) *TodoServer {
	hub := NewWebSocketHub()   // create a new hub
	go hub.StartBroadcasting() // start sending updates

	server := &TodoServer{
		store: store,
		hub:   hub,
	}

	router := http.NewServeMux()
	router.HandleFunc("/todos", server.TodosHandler)
	router.HandleFunc("/ws", server.webSocket) // register WebSocket route

	// router.HandleFunc("/todos/toggle", server.toggleTodoHandler)
	// router.HandleFunc("/todos/delete", server.deleteTodoHandler)

	server.Handler = router
	return server
}

func (t *TodoServer) TodosHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", jsonContentType)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(t.store.GetTodos())
		return
	case http.MethodPost:
		w.Header().Set("Content-Type", jsonContentType)
		w.WriteHeader(http.StatusOK)
		var todo model.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo = t.store.AddTodo(todo.Title)
		t.hub.NotifyClients(t.store.GetTodos())
		json.NewEncoder(w).Encode(todo)
		return
	case http.MethodPut:
		w.Header().Set("Content-Type", jsonContentType)
		// w.WriteHeader(http.StatusOK)
		var todo model.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo = t.store.ToggleTodo(todo.Id)
		t.hub.NotifyClients(t.store.GetTodos())
		json.NewEncoder(w).Encode(todo)
		return
	case http.MethodDelete:
		w.Header().Set("Content-Type", jsonContentType)
		var todo model.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo = t.store.DeleteTodo(todo.Id)
		t.hub.NotifyClients(t.store.GetTodos())
		json.NewEncoder(w).Encode(todo)
		return
	}
}

func (t *TodoServer) webSocket(w http.ResponseWriter, r *http.Request) {
	t.hub.HandleWebSocket(w, r)
}
