package store

import (
	"todo-backend/models"

	"gorm.io/gorm"
)

type DBTodostore struct {
	// db *gorm.DB
	// Todos is the slice of Todos
	DB    *gorm.DB
	Todos models.Todos
}

func (store *DBTodostore) GetTodos() models.Todos {
	// Get the Todos from the database
	if err := store.DB.Find(&store.Todos).Error; err != nil {
		return nil
	}
	return store.Todos
}

func (store *DBTodostore) AddTodo(todo models.Todo) models.Todo {
	// Add a new todo to the database
	if err := store.DB.Create(&todo).Error; err != nil {
		return models.Todo{}
	}
	store.Todos = append(store.Todos, todo)

	return todo
}

func (store *DBTodostore) ToggleTodo(id uint) models.Todo {
	// Toggle the completed status of a todo
	var todo models.Todo
	if err := store.DB.First(&todo, id).Error; err != nil {
		return models.Todo{}
	}
	todo.Completed = !todo.Completed
	if err := store.DB.Save(&todo).Error; err != nil {
		return models.Todo{}
	}

	return todo
}

func (store *DBTodostore) DeleteTodo(id uint) models.Todo {
	// Delete a todo from the database
	var todo models.Todo
	if err := store.DB.First(&todo, id).Error; err != nil {
		return models.Todo{}
	}
	if err := store.DB.Delete(&todo).Error; err != nil {
		return models.Todo{}
	}
	// Remove the todo from the slice
	for i, t := range store.Todos {
		if t.ID == id {
			store.Todos = append(store.Todos[:i], store.Todos[i+1:]...)
			break
		}
	}

	return todo
}
