import 'package:flutter/material.dart';
import 'package:todo_frontend/authentication/services/auth_api.dart';
import 'package:todo_frontend/authentication/services/auth_go_service.dart';
import 'package:todo_frontend/utils/email_validator.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final formKey = GlobalKey<FormState>();
    final emailFieldController = TextEditingController();
    final passwordFieldController = TextEditingController();

    final AuthApi api = AuthGoService();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Login'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text(
              'Welcome to the Login Screen!',
            ),
            Form(
              autovalidateMode: AutovalidateMode.onUserInteraction,
              key: formKey,
              child: Padding(
                padding:  EdgeInsets.all(16.0),
                child: Column(
                  children: [
                    TextFormField(
                      controller: emailFieldController,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter your email';
                        } else if (!EmailValidator.validate(value)) {
                          return 'Please enter a valid email';
                        }
                        return null;
                      },
                      decoration: const InputDecoration(
                        labelText: 'Email',
                      ),
                    ),
                    TextFormField(
                      controller: passwordFieldController,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter your password';
                        } else if (value.length < 8) {
                          return 'Please enter at least 8 characters';
                        }
                        return null;
                      },
                      decoration: const InputDecoration(
                        labelText: 'Password',
                      ),
                      obscureText: true,
                    ),

                    ElevatedButton(
                      onPressed: () async {
                        if (formKey.currentState!.validate()) {
                           ScaffoldMessenger.of(context).showSnackBar( SnackBar(content: Text('Processing email ${emailFieldController.text} and password ${passwordFieldController.text}')),);

                           // Do API call to login
                           try {
                            await api.login(emailFieldController.text, passwordFieldController.text).then((user) => {
                              ScaffoldMessenger.of(context).showSnackBar( 
                                SnackBar(content: Text('Welcome ${user.name} with email ${user.email}, role ${user.role}')))
                            });
                            
                            Navigator.pushNamedAndRemoveUntil(context, '/todo-list', (route) => false);
                           }
                           catch(e) {
                            print(e);
                             ScaffoldMessenger.of(context).showSnackBar( 
                                SnackBar(content: Text('Failed to Login. Please check again')));
                           }
                           
                          
                           
                        }
                      },
                      child: const Text('Login'),
                    ),
                  ],
                ),
              ),
            ),
            // ElevatedButton(
            //   onPressed: () {
            //     Navigator.pushNamedAndRemoveUntil(context, '/todo-list', (route) => false);
            //   },
            //   child: const Text('Go to Todo List'),
            // ),
          ],
        ),
    )
    );
  }
}