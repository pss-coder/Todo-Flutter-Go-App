import 'package:flutter/material.dart';
import 'package:todo_frontend/todo/models/todo.dart';

class TodoRow extends StatefulWidget {
  const TodoRow({super.key, required this.todo, 
  required this.onToggleComplete, required this.onDelete});
  
  final Todo todo;
  final Future<void> Function(String) onToggleComplete;
  final Future<void> Function(String) onDelete;
  

  @override
  State<TodoRow> createState() => _TodoRowState();
}

class _TodoRowState extends State<TodoRow> {

  @override
  Widget build(BuildContext context) {
    return ListTile(
      leading: Checkbox(value: widget.todo.completed, onChanged: (bool? value) async {
        await widget.onToggleComplete(widget.todo.id);
      }),
      title: Text(widget.todo.title, 
      style: TextStyle(
        fontSize: 18,
        fontWeight: FontWeight.bold,
        decoration: widget.todo.completed ? TextDecoration.lineThrough : TextDecoration.none,
      )),
      // subtitle: const Text('This is a todo item description.'),
      trailing: IconButton(
        icon: const Icon(Icons.delete),
        onPressed: () {
          widget.onDelete(widget.todo.id);
        },
      ),
    );
  }
}