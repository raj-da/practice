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
      GoRoute(path: '/dashboard', builder: (context, state) => DashboardPage()),
      GoRoute(path: '/profile', builder: (context, state) => ProfilePage()),
    ],
  );

  return MaterialApp.router(
    routerConfig: router,
    debugShowCheckedModeBanner: false,
  );
}
