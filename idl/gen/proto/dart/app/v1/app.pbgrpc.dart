//
//  Generated code. Do not modify.
//  source: app/v1/app.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'app.pb.dart' as $0;

export 'app.pb.dart';

@$pb.GrpcServiceName('app.v1.AppService')
class AppServiceClient extends $grpc.Client {
  static final _$writeMessage = $grpc.ClientMethod<$0.WriteMessageRequest, $0.WriteMessageResponse>(
      '/app.v1.AppService/WriteMessage',
      ($0.WriteMessageRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.WriteMessageResponse.fromBuffer(value));
  static final _$readMessageWithPage = $grpc.ClientMethod<$0.ReadMessageWithPageRequest, $0.ReadMessageWithPageResponse>(
      '/app.v1.AppService/ReadMessageWithPage',
      ($0.ReadMessageWithPageRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ReadMessageWithPageResponse.fromBuffer(value));
  static final _$readMessageWithTime = $grpc.ClientMethod<$0.ReadMessageWithTimeRequest, $0.ReadMessageWithTimeResponse>(
      '/app.v1.AppService/ReadMessageWithTime',
      ($0.ReadMessageWithTimeRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ReadMessageWithTimeResponse.fromBuffer(value));
  static final _$subscribeMessage = $grpc.ClientMethod<$0.SubscribeMessageRequest, $0.SubscribeMessageResponse>(
      '/app.v1.AppService/SubscribeMessage',
      ($0.SubscribeMessageRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.SubscribeMessageResponse.fromBuffer(value));

  AppServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.WriteMessageResponse> writeMessage($0.WriteMessageRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$writeMessage, request, options: options);
  }

  $grpc.ResponseFuture<$0.ReadMessageWithPageResponse> readMessageWithPage($0.ReadMessageWithPageRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$readMessageWithPage, request, options: options);
  }

  $grpc.ResponseFuture<$0.ReadMessageWithTimeResponse> readMessageWithTime($0.ReadMessageWithTimeRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$readMessageWithTime, request, options: options);
  }

  $grpc.ResponseStream<$0.SubscribeMessageResponse> subscribeMessage($0.SubscribeMessageRequest request, {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$subscribeMessage, $async.Stream.fromIterable([request]), options: options);
  }
}

@$pb.GrpcServiceName('app.v1.AppService')
abstract class AppServiceBase extends $grpc.Service {
  $core.String get $name => 'app.v1.AppService';

  AppServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.WriteMessageRequest, $0.WriteMessageResponse>(
        'WriteMessage',
        writeMessage_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.WriteMessageRequest.fromBuffer(value),
        ($0.WriteMessageResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ReadMessageWithPageRequest, $0.ReadMessageWithPageResponse>(
        'ReadMessageWithPage',
        readMessageWithPage_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ReadMessageWithPageRequest.fromBuffer(value),
        ($0.ReadMessageWithPageResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ReadMessageWithTimeRequest, $0.ReadMessageWithTimeResponse>(
        'ReadMessageWithTime',
        readMessageWithTime_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ReadMessageWithTimeRequest.fromBuffer(value),
        ($0.ReadMessageWithTimeResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SubscribeMessageRequest, $0.SubscribeMessageResponse>(
        'SubscribeMessage',
        subscribeMessage_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.SubscribeMessageRequest.fromBuffer(value),
        ($0.SubscribeMessageResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.WriteMessageResponse> writeMessage_Pre($grpc.ServiceCall call, $async.Future<$0.WriteMessageRequest> request) async {
    return writeMessage(call, await request);
  }

  $async.Future<$0.ReadMessageWithPageResponse> readMessageWithPage_Pre($grpc.ServiceCall call, $async.Future<$0.ReadMessageWithPageRequest> request) async {
    return readMessageWithPage(call, await request);
  }

  $async.Future<$0.ReadMessageWithTimeResponse> readMessageWithTime_Pre($grpc.ServiceCall call, $async.Future<$0.ReadMessageWithTimeRequest> request) async {
    return readMessageWithTime(call, await request);
  }

  $async.Stream<$0.SubscribeMessageResponse> subscribeMessage_Pre($grpc.ServiceCall call, $async.Future<$0.SubscribeMessageRequest> request) async* {
    yield* subscribeMessage(call, await request);
  }

  $async.Future<$0.WriteMessageResponse> writeMessage($grpc.ServiceCall call, $0.WriteMessageRequest request);
  $async.Future<$0.ReadMessageWithPageResponse> readMessageWithPage($grpc.ServiceCall call, $0.ReadMessageWithPageRequest request);
  $async.Future<$0.ReadMessageWithTimeResponse> readMessageWithTime($grpc.ServiceCall call, $0.ReadMessageWithTimeRequest request);
  $async.Stream<$0.SubscribeMessageResponse> subscribeMessage($grpc.ServiceCall call, $0.SubscribeMessageRequest request);
}
