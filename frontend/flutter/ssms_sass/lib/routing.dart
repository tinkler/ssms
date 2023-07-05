import 'package:go_router/go_router.dart';
import 'package:ssms_sass/page/about/_view/about_page.dart';
import 'package:ssms_sass/page/faq/_view/faq_page.dart';
import 'package:ssms_sass/page/my_home/_view/my_home_page.dart';
import 'package:ssms_sass/page/pricing/_view/pricing_page.dart';

final routing = GoRouter(routes: [
  GoRoute(path: '/', builder: (context, state) => const MyHomePage()),
  GoRoute(
    path: '/about',
    builder: (context, state) => const AboutPage(),
  ),
  GoRoute(
    path: '/faq',
    builder: (context, state) => const FaqPage(),
  ),
  GoRoute(
    path: '/pricing',
    builder: (context, state) => const PricingPage(),
  )
]);
