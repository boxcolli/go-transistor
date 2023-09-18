//
//  Generated code. Do not modify.
//  source: app/v1/app.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use writeMessageRequestDescriptor instead')
const WriteMessageRequest$json = {
  '1': 'WriteMessageRequest',
  '2': [
    {'1': 'msg', '3': 1, '4': 1, '5': 11, '6': '.obj.v1.Msg', '10': 'msg'},
  ],
};

/// Descriptor for `WriteMessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List writeMessageRequestDescriptor = $convert.base64Decode(
    'ChNXcml0ZU1lc3NhZ2VSZXF1ZXN0Eh0KA21zZxgBIAEoCzILLm9iai52MS5Nc2dSA21zZw==');

@$core.Deprecated('Use writeMessageResponseDescriptor instead')
const WriteMessageResponse$json = {
  '1': 'WriteMessageResponse',
};

/// Descriptor for `WriteMessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List writeMessageResponseDescriptor = $convert.base64Decode(
    'ChRXcml0ZU1lc3NhZ2VSZXNwb25zZQ==');

@$core.Deprecated('Use readMessageWithPageRequestDescriptor instead')
const ReadMessageWithPageRequest$json = {
  '1': 'ReadMessageWithPageRequest',
  '2': [
    {'1': 'uid', '3': 1, '4': 1, '5': 12, '10': 'uid'},
    {'1': 'chat_id', '3': 2, '4': 1, '5': 12, '10': 'chatId'},
    {'1': 'offset', '3': 3, '4': 1, '5': 5, '10': 'offset'},
  ],
};

/// Descriptor for `ReadMessageWithPageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List readMessageWithPageRequestDescriptor = $convert.base64Decode(
    'ChpSZWFkTWVzc2FnZVdpdGhQYWdlUmVxdWVzdBIQCgN1aWQYASABKAxSA3VpZBIXCgdjaGF0X2'
    'lkGAIgASgMUgZjaGF0SWQSFgoGb2Zmc2V0GAMgASgFUgZvZmZzZXQ=');

@$core.Deprecated('Use readMessageWithPageResponseDescriptor instead')
const ReadMessageWithPageResponse$json = {
  '1': 'ReadMessageWithPageResponse',
};

/// Descriptor for `ReadMessageWithPageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List readMessageWithPageResponseDescriptor = $convert.base64Decode(
    'ChtSZWFkTWVzc2FnZVdpdGhQYWdlUmVzcG9uc2U=');

@$core.Deprecated('Use readMessageWithTimeRequestDescriptor instead')
const ReadMessageWithTimeRequest$json = {
  '1': 'ReadMessageWithTimeRequest',
};

/// Descriptor for `ReadMessageWithTimeRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List readMessageWithTimeRequestDescriptor = $convert.base64Decode(
    'ChpSZWFkTWVzc2FnZVdpdGhUaW1lUmVxdWVzdA==');

@$core.Deprecated('Use readMessageWithTimeResponseDescriptor instead')
const ReadMessageWithTimeResponse$json = {
  '1': 'ReadMessageWithTimeResponse',
};

/// Descriptor for `ReadMessageWithTimeResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List readMessageWithTimeResponseDescriptor = $convert.base64Decode(
    'ChtSZWFkTWVzc2FnZVdpdGhUaW1lUmVzcG9uc2U=');

@$core.Deprecated('Use subscribeMessageRequestDescriptor instead')
const SubscribeMessageRequest$json = {
  '1': 'SubscribeMessageRequest',
};

/// Descriptor for `SubscribeMessageRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List subscribeMessageRequestDescriptor = $convert.base64Decode(
    'ChdTdWJzY3JpYmVNZXNzYWdlUmVxdWVzdA==');

@$core.Deprecated('Use subscribeMessageResponseDescriptor instead')
const SubscribeMessageResponse$json = {
  '1': 'SubscribeMessageResponse',
};

/// Descriptor for `SubscribeMessageResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List subscribeMessageResponseDescriptor = $convert.base64Decode(
    'ChhTdWJzY3JpYmVNZXNzYWdlUmVzcG9uc2U=');

