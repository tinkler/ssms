import 'package:flutter/material.dart';
import 'package:ssms_sass/widget/layout.dart';

class FaqPage extends StatelessWidget {
  const FaqPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: LayoutWidget(
          child: Column(
        children: [
          Text(
            'Frequently Asked Questions',
            style: TextStyle(fontSize: 40),
          )
        ],
      )),
    );
  }
}
