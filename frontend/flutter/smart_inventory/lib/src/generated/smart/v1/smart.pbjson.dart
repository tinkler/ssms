//
//  Generated code. Do not modify.
//  source: smart/v1/smart.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use chatMessageDescriptor instead')
const ChatMessage$json = {
  '1': 'ChatMessage',
  '2': [
    {'1': 'content', '3': 1, '4': 1, '5': 9, '10': 'content'},
    {'1': 'sender', '3': 2, '4': 1, '5': 9, '10': 'sender'},
  ],
};

/// Descriptor for `ChatMessage`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List chatMessageDescriptor = $convert.base64Decode(
    'CgtDaGF0TWVzc2FnZRIYCgdjb250ZW50GAEgASgJUgdjb250ZW50EhYKBnNlbmRlchgCIAEoCV'
    'IGc2VuZGVy');

@$core.Deprecated('Use smartChatDescriptor instead')
const SmartChat$json = {
  '1': 'SmartChat',
  '2': [
    {'1': 'messages', '3': 1, '4': 3, '5': 11, '6': '.smart.v1.ChatMessage', '10': 'messages'},
  ],
};

/// Descriptor for `SmartChat`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List smartChatDescriptor = $convert.base64Decode(
    'CglTbWFydENoYXQSMQoIbWVzc2FnZXMYASADKAsyFS5zbWFydC52MS5DaGF0TWVzc2FnZVIIbW'
    'Vzc2FnZXM=');

@$core.Deprecated('Use smartHomeDescriptor instead')
const SmartHome$json = {
  '1': 'SmartHome',
  '2': [
    {'1': 'current_model', '3': 1, '4': 1, '5': 9, '10': 'currentModel'},
  ],
};

/// Descriptor for `SmartHome`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List smartHomeDescriptor = $convert.base64Decode(
    'CglTbWFydEhvbWUSIwoNY3VycmVudF9tb2RlbBgBIAEoCVIMY3VycmVudE1vZGVs');

