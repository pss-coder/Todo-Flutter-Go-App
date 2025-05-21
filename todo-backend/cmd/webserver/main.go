package main

import (
	"log"
	"net/http"

	server "github.com/pss-coder/go-flutter-app-backend/server"
	store "github.com/pss-coder/go-flutter-app-backend/todo/store"
)

func main() {
	// Initialize the TodoStore
	store := store.NewTodoStore()

	// Create a new TodoServer with the store
	server := server.NewTodoServer(store)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", server))
}
