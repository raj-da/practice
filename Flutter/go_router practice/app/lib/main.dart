import 'package:app/consts/route_name.dart';
import 'package:app/pages/dashboard_page.dart';
import 'package:app/pages/profile_page.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

void main() {
  runApp(myApp());
}

Widget myApp() {
  //* Set up the go route config
  final GoRouter router = GoRouter(
    initialLocation: '/dashboard',
    routes: [
      GoRoute(
        name: RouteName.dashboard,
        path: '/dashboard',
        builder: (context, state) => DashboardPage(),
      ),
      GoRoute(
        name: RouteName.profile,
        path: '/profile',
        builder: (context, state) {
          final String? name = state.uri.queryParameters['name'];
          return ProfilePage(name: name,);
        },
      ),
    ],
  );

  return MaterialApp.router(
    routerConfig: router,
    debugShowCheckedModeBanner: false,
  );
}
