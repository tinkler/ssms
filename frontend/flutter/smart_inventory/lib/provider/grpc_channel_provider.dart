import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:grpc/grpc.dart';
import 'package:smart_inventory/const/global_config.dart';

final channelProvider = Provider.autoDispose((ref) {
  final uri = Uri.parse(grpcUri);
  final channel = ClientChannel(uri.host,
      port: uri.port,
      options:
          const ChannelOptions(credentials: ChannelCredentials.insecure()));
  ref.onDispose(() {
    channel.shutdown();
  });
  return channel;
});
