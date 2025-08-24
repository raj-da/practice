import 'package:app/consts/route_name.dart';
import 'package:app/user/user.dart';
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
            final user = User(name: 'Rajaf', nickName: 'Rajo');

            context.goNamed(
              RouteName.profile,
              queryParameters: {'name': 'Rajaf'}, //? use the queryParameters field to pass query parameters
              extra: user, //? use the extra field to pass complex data type
            );
          },
          child: Text('Profile'),
        ),
      ),
    );
  }
}
