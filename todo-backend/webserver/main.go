package main

import (
	"log"
	"net/http"

	todobackend "github.com/pss-coder/go-flutter-app-backend" // Replace with your actual module path
)

func main() {
	// Initialize the TodoStore
	store := todobackend.NewTodoStore()

	// Create a new TodoServer with the store
	server := todobackend.NewTodoServer(store)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", server))
}
