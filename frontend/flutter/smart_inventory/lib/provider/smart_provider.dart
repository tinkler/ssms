import 'dart:async';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:grpc/grpc.dart';
import 'package:smart_inventory/provider/smart_client.dart';
import 'package:smart_inventory/src/generated/google/protobuf/any.pb.dart';
import 'package:smart_inventory/src/generated/mrz/v1/mrz.pb.dart';
import 'package:smart_inventory/src/generated/smart/v1/smart.pbgrpc.dart';

import 'error_refresh.dart';

class SmartEnquire extends StateNotifier<SmartHome> {
  final StreamController<SmartChat> sender = StreamController();
  final StreamController<SmartChat> _receiver = StreamController();
  Stream<SmartChat> get receiver => _receiver.stream;
  late ResponseStream<Any> _stub;

  SmartEnquire(SmartGsrvClient client,
      {SmartHome? smartHome, SmartChat? firstChat})
      : super(smartHome ?? SmartHome()) {
    _ingoing(client);
    _initModel();
  }

  _ingoing(SmartGsrvClient client) {
    try {
      _stub = client.smartHomeEnquire(_outgoing());
    } catch (e) {
      print('_ingoing error ${e.toString()}');
    }

    _stub.listen((event) {
      final res = event.unpackInto(Res());
      final data = res.data.unpackInto(SmartHome());
      final resp = res.resp.unpackInto(SmartChat());
      _receiver.sink.add(resp);
      state = data;
    }, onError: (e) {
      if (e is GrpcError) {
        if (e.code == StatusCode.cancelled) {
          return;
        }
      }
      _receiver.sink.addError(e);
    }, cancelOnError: true);
  }

  Stream<Any> _outgoing() async* {
    try {
      await for (var event in sender.stream.asBroadcastStream().map((event) {
        final req = Model()
          ..data = Any.pack(state)
          ..args = Any.pack(event);
        return Any.pack(req);
      })) {
        yield event;
      }
    } catch (e) {
      print('_outgoing ${e.toString()}');
    }
  }

  _initModel() {
    // First send
    sender.sink.add(SmartChat());
  }

  @override
  void dispose() {
    _stub.cancel();
    sender.close();
    _receiver.close();
    super.dispose();
  }
}

final smartHomeEnquireProvider =
    AutoDisposeStateNotifierProvider<SmartEnquire, SmartHome>((ref) {
  final client = ref.watch(smartClient);
  return SmartEnquire(client);
});

final smartHomeEnquireSmartChatProvider =
    AutoDisposeStreamProvider<SmartChat>((ref) async* {
  final sm = ref.watch(smartHomeEnquireProvider.notifier);
  try {
    await for (final sc in sm.receiver) {
      yield sc;
    }
  } catch (e) {
    ref.read(errorRefreshProvider.notifier).markAsNeedRefresh();
  }
});
