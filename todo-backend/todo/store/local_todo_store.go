package store

import model "github.com/pss-coder/go-flutter-app-backend/todo/model"

// Exported store type
type LocalTodoStore struct {
	todos model.Todos // If Todos is unexported (todos), use model.todos
}

// Constructor
func NewTodoStore() *LocalTodoStore {
	return &LocalTodoStore{
		todos: model.Todos{}, // If unexported: model.todos{}
	}
}

func (s *LocalTodoStore) AddTodo(title string) model.Todo {
	todo := model.NewTodo(title)
	s.todos = append(s.todos, *todo)
	return *todo
}

func (s *LocalTodoStore) ToggleTodo(id string) model.Todo {
	for i, todo := range s.todos {
		if todo.Id == id {
			s.todos[i].Completed = !todo.Completed
			return s.todos[i]
		}
	}
	return model.Todo{}
}

func (s *LocalTodoStore) DeleteTodo(id string) model.Todo {
	for i, todo := range s.todos {
		if todo.Id == id {
			todo := s.todos[i]
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return todo
		}
	}
	return model.Todo{}
}

func (s *LocalTodoStore) GetTodos() model.Todos {
	return s.todos
}
