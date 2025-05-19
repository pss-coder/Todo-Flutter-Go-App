package todobackend

func NewTodoStore() *Todos {
	return &Todos{}
}

// / addTodo adds a new todo to the list of todos
func (t *Todos) addTodo(title string) Todo {
	todo := newTodo(title)
	*t = append(*t, *todo)
	return *todo
}

// Toggle the completed status of a todo
func (t *Todos) toggleTodo(id string) Todo {
	// Find the todo with the given ID and toggle its completed status
	for i, todo := range *t {
		if todo.Id == id {
			(*t)[i].Completed = !todo.Completed
			return (*t)[i]
		}
	}
	return Todo{}
}

func (t *Todos) deleteTodo(id string) Todo {
	for i, todo := range *t {
		if todo.Id == id {
			todo := (*t)[i]
			*t = append((*t)[:i], (*t)[i+1:]...)
			return todo
		}
	}
	return Todo{}
}

func (t *Todos) getTodos() Todos {
	return *t
}
