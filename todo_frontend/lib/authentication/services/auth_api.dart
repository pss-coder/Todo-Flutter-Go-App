
import 'package:todo_frontend/authentication/models/user.dart';

abstract class AuthApi {
  Future<User> login(String email, String password);
  Future<bool> signup(String email, String password, String name, String role);
  Future<bool> logout();
}