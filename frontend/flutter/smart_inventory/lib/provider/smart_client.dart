import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:grpc/grpc.dart';
import 'package:smart_inventory/provider/grpc_channel_provider.dart';
import 'package:smart_inventory/src/generated/smart/v1/smart.pbgrpc.dart';

class SmartGsrvState extends StateNotifier<SmartGsrvClient> {
  final ClientChannel channel;
  SmartGsrvState(this.channel) : super(SmartGsrvClient(channel));

  reconnect() {
    state = SmartGsrvClient(channel);
  }
}

final smartClient =
    AutoDisposeStateNotifierProvider<SmartGsrvState, SmartGsrvClient>((ref) {
  final channel = ref.watch(channelProvider);
  return SmartGsrvState(channel);
});
