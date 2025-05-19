package todobackend

// ensure our todo stores 	implement the TodoStore interface
type TodoStore interface {
	addTodo(title string) Todo
	toggleTodo(id string) Todo
	deleteTodo(id string) Todo
	getTodos() Todos
}
