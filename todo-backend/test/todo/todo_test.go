package todo_test

import (
	"reflect"
	"testing"
	"todo-backend/models"
)

func TestTodo(t *testing.T) {
	// Test 5 todos
	t.Run("a todo", func(t *testing.T) {
		todo := models.Todo{
			Title:     "Hello World",
			Completed: false,
		}

		want := models.Todo{
			Title:     "Hello World",
			Completed: false,
		}

		if todo != want {
			AssertTodo(t, todo, want)
		}
	})

	t.Run("a list of todo", func(t *testing.T) {
		var (
			todo_one = models.Todo{
				Title:     "Hello World",
				Completed: false,
			}
			todo_two = models.Todo{
				Title:     "Hello World",
				Completed: false,
			}
			todo_three = models.Todo{
				Title:     "Hello World",
				Completed: false,
			}
		)

		todos := models.Todos{todo_one, todo_two, todo_three}

		AssertTodo(t, todos[0], models.Todo{
			Title:     "Hello World",
			Completed: false,
		})

		AssertTodos(t, todos, models.Todos{
			models.Todo{
				Title:     "Hello World",
				Completed: false,
			},
			models.Todo{
				Title:     "Hello World",
				Completed: false,
			},
			models.Todo{
				Title:     "Hello World",
				Completed: false,
			},
		})

	})

	// Add a Todo

	// Toggle a Todo

	// Delete a Todo

}

func AssertTodo(t *testing.T, got, want models.Todo) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTodos(t *testing.T, got, want models.Todos) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
