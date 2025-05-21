package store

import model "github.com/pss-coder/go-flutter-app-backend/todo/model"

// ensure our todo stores 	implement the TodoStore interface
type TodoStore interface {
	AddTodo(title string) model.Todo
	ToggleTodo(id string) model.Todo
	DeleteTodo(id string) model.Todo
	GetTodos() model.Todos
}
