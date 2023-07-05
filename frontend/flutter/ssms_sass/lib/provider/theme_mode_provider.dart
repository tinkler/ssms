import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class ThemeModeState extends StateNotifier<ThemeMode> {
  ThemeModeState() : super(ThemeMode.light);

  switchMode() {
    if (state == ThemeMode.light) {
      state = ThemeMode.dark;
    } else {
      state = ThemeMode.light;
    }
  }
}

final themeModeProvider =
    StateNotifierProvider<ThemeModeState, ThemeMode>((ref) => ThemeModeState());
