package store

import "todo-backend/models"

type TodoStore interface {
	AddTodo(todo models.Todo) models.Todo
	ToggleTodo(id uint) models.Todo
	DeleteTodo(id uint) models.Todo
	GetTodos() models.Todos
}
