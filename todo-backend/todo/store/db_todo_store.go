package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	model "github.com/pss-coder/go-flutter-app-backend/todo/model"
)

type DBTodoStore struct {
	db    *sqlx.DB
	todos model.Todos
}

func NewDBTodoStore(dataSourceName string) (*DBTodoStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DBTodoStore{db: db}, nil
}

func (dbStore *DBTodoStore) GetTodos() model.Todos {
	err := dbStore.db.Select(&dbStore.todos, "SELECT id, title, completed from todos")
	if err != nil {
		fmt.Println("Failed to get todos:", err)
	}

	return dbStore.todos
}

func (dbStore *DBTodoStore) AddTodo(title string) model.Todo {
	var todo model.Todo

	query := `
	INSERT INTO todos (title)
	VALUES ($1)
	RETURNING id, title, completed
	`

	err := dbStore.db.Get(&todo, query, title)
	if err != nil {
		fmt.Println("Failed to insert todo:", err)
		return model.Todo{}
	}

	return todo
}

func (dbStore *DBTodoStore) ToggleTodo(id string) model.Todo {
	var todo model.Todo

	query := `
	UPDATE todos
	SET completed = NOT completed
	WHERE id = ($1)
	RETURNING id, title, completed
	`

	err := dbStore.db.Get(&todo, query, id)
	if err != nil {
		fmt.Printf("Failed to toggle todo with id %v: %v\n", id, err)
		return model.Todo{}
	}

	return todo
}

func (dbStore *DBTodoStore) DeleteTodo(id string) model.Todo {
	var todo model.Todo

	query := `
	DELETE FROM todos
	WHERE id = ($1)
	RETURNING id, title, completed
	`

	err := dbStore.db.Get(&todo, query, id)
	if err != nil {
		fmt.Printf("Failed to toggle todo with id %v: %v \n", id, err)
		return model.Todo{}
	}

	return todo
}
