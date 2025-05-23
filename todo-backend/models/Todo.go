package models

import (
	"gorm.io/gorm"
)

// we create a type called Todos which is a slice of Todo
type Todos []Todo

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTodo(title string) *Todo {
	return &Todo{Title: title, Completed: false}
}
