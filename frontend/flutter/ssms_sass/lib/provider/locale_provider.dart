import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class LocaleState extends StateNotifier<Locale?> {
  LocaleState() : super(null);

  setLocale(Locale? locale) {
    state = locale;
  }
}

final localeProvider =
    StateNotifierProvider<LocaleState, Locale?>((ref) => LocaleState());
