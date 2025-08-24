import 'package:app/consts/route_name.dart';
import 'package:app/pages/dashboard_page.dart';
import 'package:app/pages/profile_page.dart';
import 'package:app/user/user.dart';
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
        pageBuilder: (context, state) => CustomTransitionPage(
          key: state.pageKey,
          child: DashboardPage(),
          transitionsBuilder: (context, animation, secondaryAnimation, child) =>
              FadeTransition(opacity: animation, child: child),
        ),
        // builder: (context, state) => DashboardPage(),
      ),
      GoRoute(
        name: RouteName.profile,
        path: '/profile',
        pageBuilder: (context, state) {
          final String? name = state.uri.queryParameters['name'];
          final user = state.extra as User; //* cast to User class

          return CustomTransitionPage(
            key: state.pageKey, //? unique key to navigate the stack
            child: ProfilePage(user: user, name: name),
            transitionsBuilder:
                (context, animation, secondaryAnimation, child) {
                  //? define the animation here
                  return FadeTransition(opacity: animation, child: child);
                },
          );
        },
      ),
    ],
  );

  return MaterialApp.router(
    routerConfig: router,
    debugShowCheckedModeBanner: false,
  );
}
