package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pss-coder/go-flutter-app-backend/server"
	"github.com/pss-coder/go-flutter-app-backend/todo/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hello := os.Getenv("HELLO")
	fmt.Printf("%v\n", hello)

	// // Initialize the TodoStore
	store, err := store.NewDBTodoStore("dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	//store.NewTodoStore()

	// // Create a new TodoServer with the store
	server := server.NewTodoServer(store)

	// // Start the server
	log.Fatal(http.ListenAndServe(":8080", server))
}
