import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:go_router/go_router.dart';

class FooterWidget extends StatelessWidget {
  const FooterWidget({super.key});

  @override
  Widget build(BuildContext context) {
    final appLocalizations = AppLocalizations.of(context)!;
    return Container(
      width: double.infinity,
      decoration: BoxDecoration(
          border: Border(
              top: BorderSide(
                  width: 0.5,
                  color:
                      Theme.of(context).colorScheme.primary.withOpacity(0.8)))),
      child: Column(
        children: [
          Container(
            margin: const EdgeInsets.symmetric(vertical: 50),
            constraints: const BoxConstraints(
              maxWidth: 900,
            ),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisSize: MainAxisSize.max,
              children: [
                Container(
                  width: 300,
                  alignment: Alignment.centerLeft,
                  child: Column(
                    mainAxisSize: MainAxisSize.max,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Row(
                        children: [
                          Icon(Icons.logo_dev),
                          Text(
                            appLocalizations.companyName,
                            style: const TextStyle(fontSize: 25),
                          )
                        ],
                      ),
                      Text(appLocalizations.companyBrief),
                      const Text(
                        '© 2023 Company',
                        style: TextStyle(color: Colors.grey),
                      )
                    ],
                  ),
                ),
                const Spacer(),
                _MenuBox(
                  title: appLocalizations.menuListProduct,
                  children: [
                    TextButton(
                        onPressed: () => context.go('/pricing'),
                        child: Text(appLocalizations.pricing)),
                    TextButton(
                        onPressed: () => context.go('/faq'),
                        child: Text(appLocalizations.faq))
                  ],
                ),
                _MenuBox(
                  title: appLocalizations.menuListCompany,
                  children: [
                    TextButton(
                        onPressed: () => context.go('/about'),
                        child: Text(appLocalizations.about)),
                    TextButton(
                        onPressed: () {},
                        child: Text(appLocalizations.contact)),
                    TextButton(
                        onPressed: () {}, child: Text(appLocalizations.blog))
                  ],
                ),
                _MenuBox(title: appLocalizations.menuListSocial, children: [
                  TextButton(
                      onPressed: () {},
                      child: Row(
                        children: [
                          const Padding(
                            padding: EdgeInsets.only(right: 5),
                            child: Icon(Icons.wechat),
                          ),
                          Text(appLocalizations.wechat)
                        ],
                      ))
                ])
              ],
            ),
          )
        ],
      ),
    );
  }
}

class _MenuBox extends StatelessWidget {
  final String title;
  final List<Widget> children;
  const _MenuBox({required this.children, required this.title});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 50),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Padding(
            padding: const EdgeInsets.only(left: 10, bottom: 5),
            child: Text(
              title,
              style: const TextStyle(fontWeight: FontWeight.bold),
            ),
          ),
          ...children,
        ],
      ),
    );
  }
}
