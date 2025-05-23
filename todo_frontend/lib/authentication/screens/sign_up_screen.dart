import 'package:flutter/material.dart';
import 'package:todo_frontend/authentication/services/auth_api.dart';
import 'package:todo_frontend/authentication/services/auth_go_service.dart';
import 'package:todo_frontend/utils/email_validator.dart';


enum UserRole { user, admin }

class SignUpScreen extends StatefulWidget {
  const SignUpScreen({super.key});

  @override
  State<SignUpScreen> createState() => _SignUpScreenState();
}

class _SignUpScreenState extends State<SignUpScreen> {
  final formKey = GlobalKey<FormState>();

    final nameFieldController = TextEditingController();

    UserRole? userRoleRadio = UserRole.admin;

    final emailFieldController = TextEditingController();

    final passwordFieldController = TextEditingController();
    final confirmPasswordFieldController = TextEditingController();
    

    final AuthApi api = AuthGoService();


  @override
  Widget build(BuildContext context) {

    return Scaffold(
      appBar: AppBar(
        title: const Text('Sign Up'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text(
              'Welcome to the Sign Up Screen!',
            ),
            Form(
              autovalidateMode: AutovalidateMode.onUserInteraction,
              key: formKey,
              child: Padding(
                padding:  EdgeInsets.all(16.0),
                child: Column(
                  children: [
                    TextFormField(
                      controller: nameFieldController,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter your name';
                        }
                        return null;
                      },
                      decoration: const InputDecoration(
                        labelText: 'Name',
                      ),
                    ),
                    Row(
                      children: [
                        Expanded(
                          child: RadioListTile<UserRole>(
                            title: const Text("User"),
                            value: UserRole.user,
                            groupValue: userRoleRadio,
                            onChanged: (UserRole? value) {
                              setState(() {
                                userRoleRadio = value;
                              });
                            },
                          ),
                        ),
                        Expanded(
                          child: RadioListTile<UserRole>(
                            title: const Text("Admin"),
                            value: UserRole.admin,
                            groupValue: userRoleRadio,
                            onChanged: (UserRole? value) {
                              setState(() {
                                userRoleRadio = value;
                              });
                            },
                          ),
                        ),
                      ],
                    ),

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
                        } else if (confirmPasswordFieldController.text != value) {
                          return "Passwords do not match";
                        }
                        return null;
                      },
                      decoration: const InputDecoration(
                        labelText: 'Password',
                      ),
                      obscureText: true,
                    ),
                    TextFormField(
                      controller: confirmPasswordFieldController,
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'Please enter your password';
                        } else if (value.length < 8) {
                          return 'Please enter at least 8 characters';
                        } else if (passwordFieldController.text != value) {
                          return "Passwords do not match";
                        }
                        return null;
                      },
                      decoration: const InputDecoration(
                        labelText: 'Confirm Password',
                      ),
                      obscureText: true,
                    ),

                    ElevatedButton(
                      onPressed: () async {
                        if (formKey.currentState!.validate()) {
                           // Do API call to sign up and then login
                           api.signup(
                            emailFieldController.text,
                            passwordFieldController.text,
                            nameFieldController.text,
                            userRoleRadio!.name).then((isSignUpSuccess) {
                              if(isSignUpSuccess) {
                                ScaffoldMessenger.of(context).showSnackBar( 
                                SnackBar(content: Text('Sign up success, logging in now')));

                                api.login(emailFieldController.text, passwordFieldController.text)
                                .then((user){
                                  ScaffoldMessenger.of(context).showSnackBar( 
                                SnackBar(content: Text('Welcome ${user.name} with email ${user.email}, role ${user.role}')));
                                Navigator.pushNamedAndRemoveUntil(context, '/todo-list', (route) => false);
                                });
                              }
                            });
                        }
                      },
                      child: const Text('Sign up'),
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
      ),
    );
  }
}