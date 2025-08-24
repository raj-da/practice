import 'package:app/consts/route_name.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('profile page'),
        centerTitle: true,
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
