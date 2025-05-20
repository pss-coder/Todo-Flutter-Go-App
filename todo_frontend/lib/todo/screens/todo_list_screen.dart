import 'package:flutter/material.dart';
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
  // TextField Controller
  final TextEditingController _controller = TextEditingController();

  final TodoApi api = TodoGoService();

  late Future<List<Todo>> _todos;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    _todos = api.getTodos();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Todo List'),
      ),
      body: FutureBuilder(
        future: _todos,
        builder: (context, AsyncSnapshot<List<Todo>> snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(child: Text('Error: ${snapshot.error}'));
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(child: Text('No todos found.'));
          }

          final todos = snapshot.data!;
          return Column(
          children: [
            Expanded(
              child: ListView.builder(
                itemCount: todos.length, // Replace with your todo items count
                itemBuilder: (context, index) {
                  return TodoRow(
                    todo: todos[index],
                  );
                },
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(32.0),
              child: TextFormField(
                controller: _controller,
                textInputAction: TextInputAction.go,
                onFieldSubmitted: (value) async {
                  // Handle the submission of the new todo item
                  print('New todo item: $value');

                  // Here you can call your API to add the new todo item
                  await api.addTodo(value);

                  setState(() {
                    // Refresh the todo list
                    _todos = api.getTodos();
                  });
        
                  // clear the text field
                  _controller.clear();
                },
                decoration: InputDecoration(
                  labelText: 'Add a new todo',
                  border: OutlineInputBorder(),
                ),
              ),
            ),
          ],
        );

        },
      ),
    );
  }
}