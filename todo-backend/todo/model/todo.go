package model

import "github.com/google/uuid"

// we create a type called Todos which is a slice of Todo
type Todos []Todo

type Todo struct {
	Id string `json:"id"` // our json body will have a field called id
	// and it will be mapped to the Id field of the Todo struct
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTodo(title string) *Todo {
	return &Todo{Id: generateUniqueId(), Title: title, Completed: false}
}

func generateUniqueId() string {
	return uuid.New().String() // Use a UUID library to generate unique IDs
}
