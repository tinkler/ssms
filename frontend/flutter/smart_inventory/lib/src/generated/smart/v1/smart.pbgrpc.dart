//
//  Generated code. Do not modify.
//  source: smart/v1/smart.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/any.pb.dart' as $0;

export 'smart.pb.dart';

@$pb.GrpcServiceName('smart.v1.SmartGsrv')
class SmartGsrvClient extends $grpc.Client {
  static final _$smartHomeEnquire = $grpc.ClientMethod<$0.Any, $0.Any>(
      '/smart.v1.SmartGsrv/SmartHomeEnquire',
      ($0.Any value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Any.fromBuffer(value));

  SmartGsrvClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseStream<$0.Any> smartHomeEnquire($async.Stream<$0.Any> request, {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$smartHomeEnquire, request, options: options);
  }
}

@$pb.GrpcServiceName('smart.v1.SmartGsrv')
abstract class SmartGsrvServiceBase extends $grpc.Service {
  $core.String get $name => 'smart.v1.SmartGsrv';

  SmartGsrvServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Any, $0.Any>(
        'SmartHomeEnquire',
        smartHomeEnquire,
        true,
        true,
        ($core.List<$core.int> value) => $0.Any.fromBuffer(value),
        ($0.Any value) => value.writeToBuffer()));
  }

  $async.Stream<$0.Any> smartHomeEnquire($grpc.ServiceCall call, $async.Stream<$0.Any> request);
}
