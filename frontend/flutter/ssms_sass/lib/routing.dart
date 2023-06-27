import 'package:go_router/go_router.dart';
import 'package:ssms_sass/page/my_home/_view/my_home_page.dart';

final routing = GoRouter(routes: [
  GoRoute(path: '/', builder: (context, state) => const MyHomePage())
]);
