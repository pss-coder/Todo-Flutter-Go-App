import 'dart:convert';
import 'dart:io';

import 'package:todo_frontend/todo/models/todo.dart';
import 'package:todo_frontend/todo/services/todo_api.dart';

class TodoGoService extends TodoApi {
  static const String baseUrl = "http://localhost:8080/todos";
  final client = HttpClient();

  @override
  Future<Todo> addTodo(String title) async {
    final request = await client.postUrl(Uri.parse(baseUrl));
    request.headers.set('Content-Type', 'application/json');
    
    final body = jsonEncode({
      'title': title,
      'completed': false,
    });
    request.write(body);

    // send request
    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to add todo');
    }

    final data = await response.transform(utf8.decoder).join();
    return Todo.fromJson(jsonDecode(data));
    // throw UnimplementedError();
  }

  @override
  Future<Todo> deleteTodo(String id) async {
    final request = await client.deleteUrl(Uri.parse(baseUrl));
    request.headers.set('Content-Type', 'application/json');

    final body = jsonEncode({
      'id': id,
    });
    request.write(body);

    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to delete todo');
    }
    final data = await response.transform(utf8.decoder).join();
    return Todo.fromJson(jsonDecode(data));
  }

  @override
  Future<List<Todo>> getTodos() async {
    final requests = await client.getUrl(Uri.parse(baseUrl));
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
  Future<Todo> toggleComplete(String id) async {
    final request = await client.putUrl(Uri.parse(baseUrl));
    request.headers.set('Content-Type', 'application/json');

    final body = jsonEncode({
      'id': id,
    });
    request.write(body);

    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to toggle todo');
    }
    final data = await response.transform(utf8.decoder).join();
    return Todo.fromJson(jsonDecode(data));
  }

}