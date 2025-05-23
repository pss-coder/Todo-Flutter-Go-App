import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:todo_frontend/api/api_service.dart';
import 'package:todo_frontend/authentication/widget/logout_button.dart';
import 'package:todo_frontend/todo/models/todo.dart';
import 'package:todo_frontend/todo/services/todo_api.dart';
import 'package:todo_frontend/todo/services/todo_go_service.dart';
import 'package:todo_frontend/todo/widget/todo_row.dart';

class TodoListScreen extends StatefulWidget {
  const TodoListScreen({super.key});

  @override
  State<TodoListScreen> createState() => _TodoListScreenState();
}

class _TodoListScreenState extends State<TodoListScreen> {
  final TextEditingController _controller = TextEditingController();
  FocusNode myFocusNode = FocusNode();

  final TodoApi api = TodoGoService();
  
  // late Future<List<Todo>> _todos;
  List<Todo> _todos = [];

  WebSocket? socket;

  bool isAuthenticated = true;

  static String setupWSTodoBaseUrl()  {
  if (Platform.isAndroid) {
    // Android emulator
    return "ws://10.0.2.2:8080/ws/todos";
  } else if (Platform.isIOS) {
    // iOS simulator or device
    return "ws://localhost:8080/ws/todos";
  } else {
    // Other platforms, like web, Windows, etc.
    return "ws://localhost:8080/ws/todos";
  }
  }

  void connectToWebSocket() async {
    try {
      // Convert cookies to header string
      String cookieHeader = ApiService().cookies.map((c) => '${c.name}=${c.value}').join('; ');
      print(cookieHeader);


      socket = await WebSocket.connect(setupWSTodoBaseUrl(),
      headers: {
        'Cookie': cookieHeader
      }
      );
      print("connect to websocket");

      // we get our initial data from API once
      await api.getTodos().then((todos) {
        setState(() {
          _todos = todos;
        });
      });

      socket!.listen((data) {
        print("Received: $data");

      final todos = jsonDecode(data) 
        .map<Todo>((json) => Todo.fromJson(json))
        .toList();

        setState(() {
          _todos = todos;
        });

      });
    } catch (e) {
      setState(() {
        isAuthenticated = false;
      });
      print('Websocket error: $e');
    }
  }
  

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    connectToWebSocket();

    // _todos = api.getTodos();
  }

  @override
  void dispose() {
    if (socket != null) {
      socket!.close();
    }
    _controller.dispose();
    myFocusNode.dispose();
    super.dispose();
  }

  

  Future<void> deleteTodo(String id) async {
    // Call the API to delete the todo item
    await api.deleteTodo(id).then((_) {
      // Refresh the todo list
      // setState(() {
      //   _todos = api.getTodos();
      // });
    }).catchError((error) {
      // Handle error
      print('Error deleting todo: $error');
    });
  }

  Future<void> toggleComplete(String id) async {
    // Call the API to toggle the todo item
    await api.toggleComplete(id).then((_) {
      // Refresh the todo list
      // setState(() {
      //   _todos = api.getTodos();
      // });
    }).catchError((error) {
      // Handle error
      print('Error toggling todo: $error');
    });
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Todo List'),
        actions: [
          LogoutButton(),
        ],
      ),
      body: 
      !isAuthenticated ? Center(child: Text("Fail to connect to websocket, might be unathenticated"),) :
      _todos.isEmpty
      ? 
      Column(
              children: [
                const Center(child: Text('No todos found. Add some!')),
                Padding(
              padding: const EdgeInsets.all(32.0),
              child: TextFormField(
                controller: _controller,
                focusNode: myFocusNode,
                onEditingComplete: () {},
                textInputAction: TextInputAction.go,
                onFieldSubmitted: (value) async {
                  // Handle the submission of the new todo item
                  print('New todo item: $value');

                  // Here you can call your API to add the new todo item
                  await api.addTodo(value);

                  // setState(() {
                  //   // Refresh the todo list
                  //   _todos = api.getTodos();
                  // });
        
                  // clear the text field
                  _controller.clear();
                  myFocusNode.requestFocus();
                },
                decoration: InputDecoration(
                  labelText: 'Add a new todo',
                  border: OutlineInputBorder(),
                ),
              ),
            ),

              ],
            )
      : 

      Column(
          children: [
            Expanded(
              child: ListView.builder(
                itemCount: _todos.length, // Replace with your todo items count
                itemBuilder: (context, index) {
                  return TodoRow(
                    todo: _todos[index],
                    onToggleComplete: toggleComplete,
                    onDelete: deleteTodo,
                  );
                },
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(32.0),
              child: TextFormField(
                controller: _controller,
                textInputAction: TextInputAction.go,
                focusNode: myFocusNode,
                onTapUpOutside: (event) {
                  myFocusNode.unfocus();
                },
                // onEditingComplete: () {},
                onFieldSubmitted: (value) async {
                  // Handle the submission of the new todo item
                  print('New todo item: $value');

                  // Here you can call your API to add the new todo item
                  await api.addTodo(value);

                  // setState(() {
                  //   // Refresh the todo list
                  //   _todos = api.getTodos();
                  // });
        
                  // clear the text field
                  _controller.clear();
                  myFocusNode.requestFocus();
                },
                decoration: InputDecoration(
                  labelText: 'Add a new todo',
                  border: OutlineInputBorder(),
                ),
              ),
            ),
          ],
        )
    );
  }
}