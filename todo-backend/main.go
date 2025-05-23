package main

import (
	"fmt"
	"os"
	"todo-backend/models"
	"todo-backend/routes"
	"todo-backend/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// create gin instance
	r := gin.Default()

	// retrieve data from the database
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config := models.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	fmt.Println("DB Config: ", config)

	// initialize the database
	db, err := models.InitializeDatabase(config)
	if err != nil {
		panic("Failed to connect to the database")
	}

	DBTodoStore := store.DBTodostore{
		DB:    db,
		Todos: models.Todos{},
	}

	// load routes and store
	routes.TodoRoutes(r, &DBTodoStore)

	// run server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to run server")
	}
	fmt.Println("Server running on port 8080")
}
