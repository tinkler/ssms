import 'package:flutter/material.dart';
import 'package:ssms_sass/widget/layout.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class PricingPage extends StatelessWidget {
  const PricingPage({super.key});

  @override
  Widget build(BuildContext context) {
    final appLocalizations = AppLocalizations.of(context)!;
    final features = [
      appLocalizations.featureTitle1,
      appLocalizations.featureTitle2,
      appLocalizations.featureTitle3,
      appLocalizations.featureTitle4
    ];
    return Scaffold(
        body: LayoutWidget(
            child: ConstrainedBox(
      constraints: const BoxConstraints(
        maxWidth: 900,
      ),
      child: Column(
        children: [
          const SizedBox(
            height: 50,
          ),
          Text(
            appLocalizations.pricing,
            style: const TextStyle(fontSize: 40),
          ),
          const SizedBox(
            height: 15,
          ),
          GridView(
            physics: const NeverScrollableScrollPhysics(),
            shrinkWrap: true,
            gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2, crossAxisSpacing: 30, childAspectRatio: 0.8),
            children: [
              _PriceTag(
                price: appLocalizations.localeName == 'zh'
                    ? _Price.rmb(price: '')
                    : _Price.dollar(price: ''),
                name: appLocalizations.personal,
                books: appLocalizations.localeName == 'zh'
                    ? ['V1版账套数 10个', '']
                    : [],
                features: features,
              ),
              _PriceTag(
                price: appLocalizations.localeName == 'zh'
                    ? _Price.rmb(price: '15')
                    : _Price.dollar(price: '1.99'),
                name: appLocalizations.priceTagPro,
                books: appLocalizations.localeName == 'zh'
                    ? ['V1版账套数 10', 'V2版账套数 5个']
                    : [],
                features: features,
              )
            ],
          ),
          const SizedBox(
            height: 85,
          )
        ],
      ),
    )));
  }
}

class _Price {
  final String timeUnit;
  final String price;
  final String priceUnit;
  final bool prefix;

  _Price.rmb({required this.price})
      : timeUnit = '月',
        priceUnit = '元',
        prefix = false;
  _Price.dollar({required this.price})
      : timeUnit = 'mo',
        priceUnit = "\$",
        prefix = true;
}

class _PriceTag extends StatelessWidget {
  final String name;
  final _Price price;
  final List<String> books;
  final List<String> features;
  const _PriceTag(
      {required this.name,
      required this.price,
      required this.books,
      required this.features});

  @override
  Widget build(BuildContext context) {
    String priceStr = price.price;
    late final Widget priceWidget;
    if (price.price != '') {
      if (price.prefix) {
        priceStr = price.priceUnit + priceStr;
      } else {
        priceStr = priceStr + price.priceUnit;
      }
      priceWidget = Padding(
        padding: const EdgeInsets.all(5.0),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            Text(
              priceStr,
              style: const TextStyle(fontSize: 35, fontWeight: FontWeight.bold),
            ),
            Padding(
              padding: const EdgeInsets.only(bottom: 5),
              child: Text(
                '/${price.timeUnit}',
                style: const TextStyle(fontSize: 20, color: Colors.grey),
              ),
            )
          ],
        ),
      );
    } else {
      priceStr = AppLocalizations.of(context)!.priceTagFree;
      priceWidget = Padding(
        padding: const EdgeInsets.all(10.0),
        child: Text(
          priceStr,
          style: TextStyle(fontSize: 25, color: Colors.red.withOpacity(0.8)),
        ),
      );
    }

    return Container(
      padding: const EdgeInsets.symmetric(vertical: 25, horizontal: 20),
      decoration: BoxDecoration(
          color: Theme.of(context).colorScheme.background,
          boxShadow: const [
            BoxShadow(
                offset: Offset(-0.05, 1), blurRadius: 0.5, color: Colors.grey)
          ],
          borderRadius: const BorderRadius.all(Radius.circular(5))),
      child: Column(crossAxisAlignment: CrossAxisAlignment.start, children: [
        Text(
          name,
          style: const TextStyle(fontSize: 20),
        ),
        priceWidget,
        ListView.builder(
          shrinkWrap: true,
          // list features and books with a separete between them.
          itemCount: features.length + books.length + 1,
          itemBuilder: (context, index) {
            if (index == books.length) {
              return const Divider();
            }
            final String title = index > books.length
                ? features[index - books.length - 1]
                : books[index];
            if (title == '') {
              return const SizedBox(
                height: 36,
              );
            }
            return Padding(
              padding: const EdgeInsets.symmetric(vertical: 3),
              child: Row(
                children: [
                  const Icon(
                    Icons.check,
                    color: Colors.green,
                  ),
                  const SizedBox(
                    width: 8,
                  ),
                  Text(title)
                ],
              ),
            );
          },
        ),
        const Spacer(),
        ElevatedButton(
            onPressed: () {},
            style: ElevatedButton.styleFrom(shape: const LinearBorder()),
            child: Container(
                width: double.infinity,
                alignment: Alignment.center,
                child: Text(AppLocalizations.of(context)!.priceTagChoose)))
      ]),
    );
  }
}
