import 'package:flutter/material.dart';
import 'package:ssms_sass/widget/layout.dart';

class AboutPage extends StatelessWidget {
  const AboutPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: LayoutWidget(child: Text('about')),
    );
  }
}
