import 'dart:io';

class ApiService {
  // SIngleton instance
  static final ApiService _instance = ApiService._internal();
  factory ApiService() => _instance;
  ApiService._internal();

  final HttpClient client = HttpClient();
}