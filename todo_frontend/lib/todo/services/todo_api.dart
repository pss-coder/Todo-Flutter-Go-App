import 'package:todo_frontend/todo/models/todo.dart';

abstract class TodoApi {
  Future<List<Todo>> getTodos();
  Future<Todo> addTodo(Todo todo);
  Future<Todo> updateTodo(Todo todo);
  Future<Todo> deleteTodo(String id);
}

/// Error thrown when a [Todo] with a given id is not found.
class TodoNotFoundException implements Exception {}