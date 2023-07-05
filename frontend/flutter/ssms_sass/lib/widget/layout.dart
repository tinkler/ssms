import 'package:flutter/material.dart';
import 'package:ssms_sass/widget/footer.dart';
import 'package:ssms_sass/widget/navbar.dart';

class LayoutWidget extends StatelessWidget {
  final Widget child;
  const LayoutWidget({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: [
          const Navbar(),
          child,
          const FooterWidget(),
        ],
      ),
    );
  }
}
