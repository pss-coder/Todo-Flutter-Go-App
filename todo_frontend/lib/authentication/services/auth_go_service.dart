import 'dart:convert';
import 'dart:ffi';
import 'dart:io';

import 'package:todo_frontend/api/api_service.dart';
import 'package:todo_frontend/authentication/models/user.dart';
import 'package:todo_frontend/authentication/services/auth_api.dart';

class AuthGoService extends AuthApi {

  static String setupAuthBaseURL()  {
  if (Platform.isAndroid) {
    // Android emulator
    return "http://10.0.2.2:8080";
  } else if (Platform.isIOS) {
    // iOS simulator or device
    return "http://localhost:8080";
  } else {
    // Other platforms, like web, Windows, etc.
    return "http://localhost:8080";
  }
}

  static final String baseUrl = setupAuthBaseURL();

  @override
  Future<User> login(String email, String password) async {
    final request = await ApiService().client.postUrl(Uri.parse("$baseUrl/login"));
    request.headers.set('Content-Type', 'application/json');
    
    final body = jsonEncode({
      'email': email,
      'password': password,
    });

    // use UTF8 encoding
    request.add(utf8.encode(body));

    // send request
    final response = await request.close();

    if (response.statusCode != 200) {
      throw Exception('Failed to login');
    }

    final data = await response.transform(utf8.decoder).join();
    return User.fromJson(jsonDecode(data));
  }

  @override
  Future<bool> logout() async {
    final request = await ApiService().client.getUrl(Uri.parse("$baseUrl/logout"));
    request.headers.set('Content-Type', 'application/json');

    // send request
    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to logout');
    }
    final data = await response.transform(utf8.decoder).join();
    print(data);

    return response.statusCode == 200;
  }

  @override
  Future<bool> signup(String email, String password, String name, String role) async {
    final request = await ApiService().client.postUrl(Uri.parse("$baseUrl/signup"));
    request.headers.set('Content-Type', 'application/json');

    final body = jsonEncode({
      'email': email,
      'password': password,
      'name': name,
      'role': role,
    });

    // use UTF8 encoding
    request.add(utf8.encode(body));

    // send request
    final response = await request.close();
    if (response.statusCode != 200) {
      throw Exception('Failed to logout');
    }
    final data = await response.transform(utf8.decoder).join();
    print(data);

    return response.statusCode == 200;


  }

}