import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:todo_frontend/api/api_service.dart';
import 'package:todo_frontend/authentication/screens/auth_screen.dart';
import 'package:todo_frontend/authentication/screens/login_screen.dart';
import 'package:todo_frontend/authentication/screens/sign_up_screen.dart';
import 'package:todo_frontend/todo/screens/todo_list_screen.dart';

class MyTodoApp extends StatelessWidget {
  const MyTodoApp({super.key});

  static ApiService api = ApiService(); // initialise once

  Route<dynamic>? _customRoute(RouteSettings settings) {
    log(settings.name!);
    if (settings.name == '/todo-list') {
      return PageRouteBuilder(
        settings: settings,
        pageBuilder: (_, __, ___) => const TodoListScreen(),
        transitionsBuilder: (_, animation, __, child) {
          return FadeTransition(
            opacity: animation,
            child: child,
          );
        },
        transitionDuration: const Duration(milliseconds: 400),
      );
    } else if (settings.name == '/') {
      return PageRouteBuilder(
        settings: settings,
        pageBuilder: (_, __, ___) => const AuthScreen(),
        transitionsBuilder: (_, animation, __, child) {
          return FadeTransition(
            opacity: animation,
            child: child,
          );
        },
        transitionDuration: const Duration(milliseconds: 400),
      );
    }
    // Return null to use default MaterialPageRoute
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Todo App',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      initialRoute: '/',
     routes: {
        '/signup': (context) => const SignUpScreen(),
        '/login': (context) => const LoginScreen(),
        // '/todo-list' intentionally omitted here
      },
      // will only call if the route is not defined in the routes table
      onGenerateRoute: _customRoute,
      // home: const TodoListScreen(),
    );
  }
}