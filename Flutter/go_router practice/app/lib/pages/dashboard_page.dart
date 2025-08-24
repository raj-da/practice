import 'package:app/consts/route_name.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class DashboardPage extends StatelessWidget {
  const DashboardPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('DashBoard page'),
        centerTitle: true,
        backgroundColor: Colors.blue,
      ),

      body: Center(
        child: ElevatedButton(
          onPressed: () {
            context.goNamed(RouteName.profile, pathParameters: {'name': 'Rajaf'});
          },
          child: Text('Profile'),
        ),
      ),
    );
  }
}
