import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:ssms_sass/widget/footer.dart';
import 'package:ssms_sass/widget/layout.dart';
import 'package:ssms_sass/widget/navbar.dart';
import 'package:ssms_sass/widget/started.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: LayoutWidget(
            child: Column(
      mainAxisSize: MainAxisSize.max,
      children: <Widget>[
        Container(
          height: 300,
          width: double.infinity,
          alignment: Alignment.center,
          decoration: const BoxDecoration(
              image: DecorationImage(
                  colorFilter: ColorFilter.srgbToLinearGamma(),
                  image: AssetImage('assets/images/banner.png'),
                  fit: BoxFit.fitWidth,
                  alignment: Alignment.bottomLeft)),
          child: ConstrainedBox(
            constraints: const BoxConstraints(
              maxWidth: 600,
            ),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  AppLocalizations.of(context)!.bannerTitle,
                  style: const TextStyle(fontSize: 20, color: Colors.white),
                ),
                Text(
                  AppLocalizations.of(context)!.bannerDescription(
                      AppLocalizations.of(context)!.appTitle),
                  style: const TextStyle(color: Colors.white),
                ),
                const SizedBox(
                  height: 30,
                ),
                const StartedButton(),
              ],
            ),
          ),
        ),
        const SizedBox(
          height: 60,
        ),
        const _FeatureWidget(),
        const SizedBox(
          height: 30,
        ),
        const StartedWidget()
      ],
    )));
  }
}

class _FeatureWidget extends StatelessWidget {
  const _FeatureWidget();

  @override
  Widget build(BuildContext context) {
    final borderSide = BorderSide(
        color:
            Theme.of(context).colorScheme.secondaryContainer.withOpacity(0.8));
    return Container(
      constraints: const BoxConstraints(maxWidth: 900),
      decoration: BoxDecoration(
          color: Theme.of(context).colorScheme.background,
          borderRadius: const BorderRadius.all(Radius.circular(5)),
          boxShadow: const [
            BoxShadow(offset: Offset(0.5, 0.5), blurRadius: 1)
          ]),
      child: GridView(
        gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
          childAspectRatio: 1.5,
          crossAxisCount: 2,
        ),
        physics: const NeverScrollableScrollPhysics(),
        shrinkWrap: true,
        children: [
          _FeatureBox(
            title: AppLocalizations.of(context)!.featureTitle1,
            content: AppLocalizations.of(context)!
                .featureDescription1(AppLocalizations.of(context)!.appTitle),
            borderSide: borderSide,
            assetName: 'assets/images/feature1.png',
          ),
          _FeatureBox(
            title: AppLocalizations.of(context)!.featureTitle2,
            content: AppLocalizations.of(context)!.featureDescription2,
            borderSide: borderSide,
          ),
          _FeatureBox(
            title: AppLocalizations.of(context)!.featureTitle3,
            content: AppLocalizations.of(context)!
                .featureDescription3(AppLocalizations.of(context)!.appTitle),
            borderSide: borderSide,
          ),
          _FeatureBox(
            title: AppLocalizations.of(context)!.featureTitle4,
            content: AppLocalizations.of(context)!.featureDescription4,
            borderSide: borderSide,
            assetName: 'assets/images/feature4.png',
          ),
        ],
      ),
    );
  }
}

class _FeatureBox extends StatelessWidget {
  final BorderSide borderSide;
  final String title;
  final String content;
  final String? assetName;
  const _FeatureBox(
      {required this.borderSide,
      required this.title,
      required this.content,
      this.assetName});

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        final maxLines = constraints.maxWidth > 500 ? 3 : 2;
        return Container(
          width: constraints.maxWidth,
          height: constraints.maxHeight,
          decoration: BoxDecoration(
              border: Border(bottom: borderSide, right: borderSide)),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            mainAxisSize: MainAxisSize.max,
            children: [
              Container(
                height: constraints.maxWidth / 3,
                width: constraints.maxWidth / 3,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(10),
                  image: assetName != null
                      ? DecorationImage(
                          image: AssetImage(assetName!), fit: BoxFit.cover)
                      : null,
                ),
              ),
              Text(
                title,
                style: const TextStyle(fontSize: 20),
              ),
              Padding(
                padding:
                    const EdgeInsets.symmetric(vertical: 5, horizontal: 30),
                child: Text(
                  content,
                  overflow: TextOverflow.ellipsis,
                  maxLines: maxLines,
                ),
              )
            ],
          ),
        );
      },
    );
  }
}

class MiddleCircularCutClipper extends CustomClipper<Path> {
  @override
  Path getClip(Size size) {
    final path = Path();

    // Define the dimensions of the rectangle
    final rectWidth = size.width;
    final rectHeight = size.height;

    // Define the radius of the circular cut
    final circularRadius = 40.0;

    // Calculate the position and size of the circular cut
    final cutRect = Rect.fromCircle(
        center: Offset(rectWidth / 2, rectHeight / 2), radius: circularRadius);

    // Add the rectangle to the path
    path.addRect(Rect.fromLTRB(0, 0, rectWidth, rectHeight));

    // Add the circle to be cut out
    path.addOval(cutRect);
    path.close();

    return path;
  }

  @override
  bool shouldReclip(CustomClipper<Path> oldClipper) {
    return false;
  }
}
