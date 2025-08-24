import 'package:app/consts/route_name.dart';
import 'package:app/user/user.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ProfilePage extends StatelessWidget {
  final String? name;
  final User user;
  const ProfilePage({required this.user, this.name, super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: name == null ? Text('Hello, ') : Text('Hello $name (${user.nickName})'),
        backgroundColor: Colors.blue,
      ),

      body: Center(
        child: ElevatedButton(
          onPressed: () {
            context.goNamed(RouteName.dashboard);
          },
          child: Text('dashboard'),
        ),
      ),
    );
  }
}
