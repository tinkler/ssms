import 'package:flutter_riverpod/flutter_riverpod.dart';

class ErrorRefresh extends StateNotifier<bool> {
  ErrorRefresh() : super(false);

  markAsNeedRefresh() {
    state = true;
  }

  reset() {
    state = false;
  }
}

final errorRefreshProvider =
    StateNotifierProvider<ErrorRefresh, bool>((ref) => ErrorRefresh());
