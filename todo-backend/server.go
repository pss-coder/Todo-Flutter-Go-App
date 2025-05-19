package todobackend

import (
	"encoding/json"
	"net/http"
)

type TodoServer struct {
	store TodoStore
	http.Handler
}

const jsonContentType = "application/json"

func NewTodoServer(store TodoStore) *TodoServer {
	server := &TodoServer{
		store: store,
	}

	router := http.NewServeMux()
	router.HandleFunc("/todos", server.TodosHandler)
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
		json.NewEncoder(w).Encode(t.store.getTodos())
		return
	case http.MethodPost:
		w.Header().Set("Content-Type", jsonContentType)
		w.WriteHeader(http.StatusOK)
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo = t.store.addTodo(todo.Title)
		json.NewEncoder(w).Encode(todo)
		return
	case http.MethodPut:
		w.Header().Set("Content-Type", jsonContentType)
		// w.WriteHeader(http.StatusOK)
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo = t.store.toggleTodo(todo.Id)
		json.NewEncoder(w).Encode(todo)
		return
	case http.MethodDelete:
		w.Header().Set("Content-Type", jsonContentType)
		var todo Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todo = t.store.deleteTodo(todo.Id)
		json.NewEncoder(w).Encode(todo)
		return
	}

}

// func (t *TodoServer) addTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", jsonContentType)
// 	w.WriteHeader(http.StatusOK)
// 	var todo Todo
// 	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	t.store.addTodo(todo.Title)
// 	json.NewEncoder(w).Encode(todo)
// }
