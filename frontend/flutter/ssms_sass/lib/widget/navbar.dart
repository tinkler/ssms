import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:ssms_sass/provider/locale_provider.dart';
import 'package:ssms_sass/provider/theme_mode_provider.dart';

class Navbar extends StatelessWidget {
  const Navbar({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 50,
      width: double.infinity,
      alignment: Alignment.center,
      decoration: BoxDecoration(
          color: Theme.of(context).colorScheme.background,
          border: Border(
              bottom: BorderSide(
                  width: 0.5,
                  color:
                      Theme.of(context).colorScheme.primary.withOpacity(0.8)))),
      child: ConstrainedBox(
        constraints: const BoxConstraints(maxWidth: 900),
        child: Row(
          children: [
            GestureDetector(
              onTap: () => context.go('/'),
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(
                  AppLocalizations.of(context)!.appTitle,
                  style: Theme.of(context).appBarTheme.titleTextStyle,
                ),
              ),
            ),
            const Spacer(),
            Consumer(
              builder: (context, ref, _) {
                final locale = ref.read(localeProvider);
                return DropdownButton(
                    value: locale ?? const Locale('zh'),
                    onChanged: (v) {
                      ref.read(localeProvider.notifier).setLocale(v);
                    },
                    items: const [
                      DropdownMenuItem(value: Locale('zh'), child: Text('中文')),
                      DropdownMenuItem(
                          value: Locale('en'), child: Text('English')),
                    ]);
              },
            ),
            Padding(
              padding: const EdgeInsets.all(12.0),
              child: TextButton(
                onPressed: () {},
                child: Text(AppLocalizations.of(context)!.appBarSignIn),
              ),
            ),
            Consumer(builder: (context, ref, _) {
              final model = ref.watch(themeModeProvider);
              switchModel() =>
                  ref.read(themeModeProvider.notifier).switchMode();
              switch (model) {
                case ThemeMode.light:
                  return IconButton(
                      onPressed: switchModel,
                      icon: const Icon(Icons.light_mode));
                case ThemeMode.dark:
                  return IconButton(
                      onPressed: switchModel,
                      icon: const Icon(Icons.dark_mode));
                default:
                  return IconButton(
                      onPressed: switchModel,
                      icon: const Icon(Icons.light_mode));
              }
            })
          ],
        ),
      ),
    );
  }
}
