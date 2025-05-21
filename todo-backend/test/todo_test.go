package todo_test

import (
	"reflect"
	"testing"

	"github.com/pss-coder/go-flutter-app-backend/todo/model"
)

func TestTodo(t *testing.T) {
	// Test 5 todos
	t.Run("a todo", func(t *testing.T) {
		todo := model.Todo{
			Id:        "ID ONE",
			Title:     "Hello World",
			Completed: false,
		}

		want := model.Todo{
			Id:        "ID ONE",
			Title:     "Hello World",
			Completed: false,
		}

		if todo != want {
			AssertTodo(t, todo, want)
		}
	})

	t.Run("a list of todo", func(t *testing.T) {
		var (
			todo_one = model.Todo{
				Id:        "ID ONE",
				Title:     "Hello World",
				Completed: false,
			}
			todo_two = model.Todo{
				Id:        "ID ONE",
				Title:     "Hello World",
				Completed: false,
			}
			todo_three = model.Todo{
				Id:        "ID ONE",
				Title:     "Hello World",
				Completed: false,
			}
		)

		todos := model.Todos{todo_one, todo_two, todo_three}

		AssertTodo(t, todos[0], model.Todo{
			Id:        "ID ONE",
			Title:     "Hello World",
			Completed: false,
		})

		AssertTodos(t, todos, model.Todos{
			model.Todo{
				Id:        "ID ONE",
				Title:     "Hello World",
				Completed: false,
			},
			model.Todo{
				Id:        "ID ONE",
				Title:     "Hello World",
				Completed: false,
			},
			model.Todo{
				Id:        "ID ONE",
				Title:     "Hello World",
				Completed: false,
			},
		})

	})

	// Add a Todo

	// Toggle a Todo

	// Delete a Todo

}

func AssertTodo(t *testing.T, got, want model.Todo) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTodos(t *testing.T, got, want model.Todos) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
