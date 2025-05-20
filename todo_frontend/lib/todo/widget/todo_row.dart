import 'package:flutter/material.dart';
import 'package:todo_frontend/todo/models/todo.dart';

class TodoRow extends StatefulWidget {
  const TodoRow({super.key, required this.todo});
  final Todo todo;

  @override
  State<TodoRow> createState() => _TodoRowState();
}

class _TodoRowState extends State<TodoRow> {
  bool isChecked = false;

  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: Checkbox(value: isChecked, onChanged: (bool? value) {
        setState(() {
          isChecked = value ?? false;
        });
      }),
      title: Text(widget.todo.title, 
      style: TextStyle(
        fontSize: 18,
        fontWeight: FontWeight.bold,
        decoration: isChecked ? TextDecoration.lineThrough : TextDecoration.none,
      )),
      // subtitle: const Text('This is a todo item description.'),
      trailing: IconButton(
        icon: const Icon(Icons.delete),
        onPressed: () {
          // Handle delete action
        },
      ),
    );
  }
}