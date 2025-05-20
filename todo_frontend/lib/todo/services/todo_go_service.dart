import 'dart:convert';
import 'dart:io';

import 'package:todo_frontend/todo/models/todo.dart';
import 'package:todo_frontend/todo/services/todo_api.dart';

class TodoGoService extends TodoApi {
  static const String baseUrl = "http://localhost:8080";
  final client = HttpClient();

  @override
  Future<Todo> addTodo(Todo todo) {
    // TODO: implement addTodo
    throw UnimplementedError();
  }

  @override
  Future<Todo> deleteTodo(String id) {
    // TODO: implement deleteTodo
    throw UnimplementedError();
  }

  @override
  Future<List<Todo>> getTodos() async {
    final requests = await client.getUrl(Uri.parse('$baseUrl/todos'));
    final response = await requests.close();

    if (response.statusCode != 200) {
      throw Exception('Failed to load todos');
    }

    

    final data = await response.transform(utf8.decoder).join();
    print(data);
    
    return jsonDecode(data) 
        .map<Todo>((json) => Todo.fromJson(json))
        .toList();
  }

  @override
  Future<Todo> updateTodo(Todo todo) {
    // TODO: implement updateTodo
    throw UnimplementedError();
  }

}