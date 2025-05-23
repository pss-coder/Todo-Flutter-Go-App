import 'dart:convert';
import 'dart:io';

import 'package:todo_frontend/api/api_service.dart';
import 'package:todo_frontend/todo/models/todo.dart';
import 'package:todo_frontend/todo/services/todo_api.dart';




class TodoGoService extends TodoApi {
  static String setupTodoBaseUrl()  {
  if (Platform.isAndroid) {
    // Android emulator
    return "http://10.0.2.2:8080/todos";
  } else if (Platform.isIOS) {
    // iOS simulator or device
    return "http://localhost:8080/todos";
  } else {
    // Other platforms, like web, Windows, etc.
    return "http://localhost:8080/todos";
  }
}


  static final String baseUrl = setupTodoBaseUrl();

  @override
  Future<Todo> addTodo(String title) async {
    final request = await ApiService().client.postUrl(Uri.parse(baseUrl));
    request.cookies.add(ApiService().cookies.first);
    request.headers.set('Content-Type', 'application/json');
    
    final body = jsonEncode({
      'title': title,
      'completed': false,
    });
    
    // use UTF8 encoding
    request.add(utf8.encode(body));

    // send request
    final response = await request.close();
    if (response.statusCode != 201) {
      throw Exception('Failed to add todo');
    }

    final data = await response.transform(utf8.decoder).join();
    return Todo.fromJson(jsonDecode(data));
    // throw UnimplementedError();
  }

  @override
  Future<Todo> deleteTodo(String id) async {
    print("Delete: $baseUrl/$id");
    final request = await ApiService().client.deleteUrl(Uri.parse("$baseUrl/$id"));
    request.cookies.add(ApiService().cookies.first);
    request.headers.set('Content-Type', 'application/json');

    // final body = jsonEncode({
    //   'id': id,
    // });
    // request.write(body);

    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to delete todo');
    }
    final data = await response.transform(utf8.decoder).join();
    return Todo.fromJson(jsonDecode(data));
  }

  @override
  Future<List<Todo>> getTodos() async {
    final requests = await ApiService().client.getUrl(Uri.parse(baseUrl));
    requests.cookies.add(ApiService().cookies.first);

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
    final request = await ApiService().client.putUrl(Uri.parse("$baseUrl/$id"));
    request.cookies.add(ApiService().cookies.first);
    request.headers.set('Content-Type', 'application/json');

    // final body = jsonEncode({
    //   'id': id,
    // });
    // request.write(body);

    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to toggle todo');
    }
    final data = await response.transform(utf8.decoder).join();
    return Todo.fromJson(jsonDecode(data));
  }

}