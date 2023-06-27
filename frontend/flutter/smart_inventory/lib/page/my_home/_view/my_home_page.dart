import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:smart_inventory/page/my_home/_widget/slide_in_input.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:smart_inventory/page/my_home/_widget/smart_screen.dart';
import 'package:smart_inventory/provider/error_refresh.dart';
import 'package:smart_inventory/provider/smart_client.dart';
import 'package:smart_inventory/provider/smart_provider.dart';

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});
  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final GlobalKey<SlideInInputWidgetState> _inputWidget = GlobalKey();

  void _incrementCounter() {
    _inputWidget.currentState?.startAnimation();
  }

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    return Scaffold(
      appBar: AppBar(
        title: Text(AppLocalizations.of(context)?.appName ?? 'SmartInventory'),
      ),
      body: SizedBox(
        width: double.infinity,
        height: double.infinity,
        child: LayoutBuilder(
          builder: (context, constraints) {
            return Stack(
              // Center is a layout widget. It takes a single child and positions it
              // in the middle of the parent.
              children: [
                GestureDetector(
                  onTap: () {
                    _inputWidget.currentState?.foldAnimation();
                  },
                  child: const SmartScreenWidget(),
                ),
                Positioned(
                    bottom: 10,
                    child: SizedBox(
                      width: constraints.maxWidth,
                      child: SlideInInputWidget(
                        key: _inputWidget,
                      ),
                    )),
                Consumer(
                  builder: (context, ref, _) {
                    final showRefresh = ref.watch(errorRefreshProvider);
                    if (showRefresh) {
                      return Positioned.fill(
                          child: ConstrainedBox(
                        constraints: const BoxConstraints(
                          minWidth: 100,
                          minHeight: 80,
                        ),
                        child: AlertDialog(
                          title: const Text('发生网络错误'),
                          actions: [
                            TextButton(
                                onPressed: () {
                                  ref.read(smartClient.notifier).reconnect();
                                  ref
                                      .read(errorRefreshProvider.notifier)
                                      .reset();
                                },
                                child: const Text('刷新'))
                          ],
                        ),
                      ));
                    }
                    return Container();
                  },
                )
              ],
            );
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
