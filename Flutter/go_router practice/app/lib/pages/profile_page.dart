import 'package:app/consts/route_name.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ProfilePage extends StatelessWidget {
  final String? name;
  const ProfilePage({this.name, super.key,});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: name == null ? Text('Hello, ') : Text('Hello $name'),
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
