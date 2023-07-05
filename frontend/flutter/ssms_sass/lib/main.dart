import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:ssms_sass/provider/locale_provider.dart';
import 'package:ssms_sass/provider/theme_mode_provider.dart';
import 'package:ssms_sass/routing.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

void main() {
  runApp(const ProviderScope(child: MyApp()));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return Consumer(
      builder: (context, ref, _) {
        final themeMode = ref.watch(themeModeProvider);
        final locale = ref.watch(localeProvider);
        print(locale.toString());
        return MaterialApp.router(
          debugShowCheckedModeBanner: false,
          onGenerateTitle: (BuildContext context) =>
              AppLocalizations.of(context)!.appTitle,
          darkTheme: ThemeData(
              colorScheme: ColorScheme.fromSeed(
                  brightness: Brightness.dark,
                  seedColor: const Color.fromARGB(255, 33, 93, 243)),
              useMaterial3: true,
              appBarTheme: AppBarTheme.of(context).copyWith(
                  titleTextStyle: TextStyle(
                      fontSize: 20,
                      color: Theme.of(context).colorScheme.primary))),
          theme: ThemeData(
              // This is the theme of your application.
              //
              // TRY THIS: Try running your application with "flutter run". You'll see
              // the application has a blue toolbar. Then, without quitting the app,
              // try changing the seedColor in the colorScheme below to Colors.green
              // and then invoke "hot reload" (save your changes or press the "hot
              // reload" button in a Flutter-supported IDE, or press "r" if you used
              // the command line to start the app).
              //
              // Notice that the counter didn't reset back to zero; the application
              // state is not lost during the reload. To reset the state, use hot
              // restart instead.
              //
              // This works for code too, not just values: Most code changes can be
              // tested with just a hot reload.
              colorScheme: ColorScheme.fromSeed(
                  seedColor: const Color.fromARGB(255, 33, 93, 243)),
              useMaterial3: true,
              appBarTheme: AppBarTheme.of(context).copyWith(
                  titleTextStyle: TextStyle(
                      fontSize: 20,
                      color: Theme.of(context).colorScheme.primary))),
          routerConfig: routing,
          localizationsDelegates: const [
            AppLocalizations.delegate,
            GlobalMaterialLocalizations.delegate,
            GlobalWidgetsLocalizations.delegate,
            GlobalCupertinoLocalizations.delegate,
          ],
          supportedLocales: const [Locale('zh'), Locale('en')],
          themeMode: themeMode,
          locale: locale,
        );
      },
    );
  }
}
