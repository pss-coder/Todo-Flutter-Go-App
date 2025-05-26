import 'package:flutter/material.dart';
import 'package:todo_frontend/api/api_service.dart';
import 'package:todo_frontend/authentication/services/auth_api.dart';
import 'package:todo_frontend/authentication/services/auth_go_service.dart';

class LogoutButton extends StatelessWidget {
  LogoutButton({super.key});

  AuthApi authService = AuthGoService();

  @override
  Widget build(BuildContext context) {
    return IconButton(onPressed: (){
            // Handle logout action
            print('Logout button pressed');
            
            authService.logout().then((isLogoutSuccess) {
              if (isLogoutSuccess) {
                ApiService().cookies = []; // clear cookies
                Navigator.pushReplacementNamed(context, '/');
                ScaffoldMessenger.of(context).showSnackBar( 
                                SnackBar(content: Text('Logout success')));
              } else {
                ScaffoldMessenger.of(context).showSnackBar( 
                                SnackBar(content: Text('Logout failed. something went wrong')));
              }


            });
          }, icon: Icon(Icons.logout_sharp));
  }
}