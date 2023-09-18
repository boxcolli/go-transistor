//
//  Generated code. Do not modify.
//  source: app/v1/app.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../obj/v1/obj.pb.dart' as $1;

class WriteMessageRequest extends $pb.GeneratedMessage {
  factory WriteMessageRequest({
    $1.Msg? msg,
  }) {
    final $result = create();
    if (msg != null) {
      $result.msg = msg;
    }
    return $result;
  }
  WriteMessageRequest._() : super();
  factory WriteMessageRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WriteMessageRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'WriteMessageRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..aOM<$1.Msg>(1, _omitFieldNames ? '' : 'msg', subBuilder: $1.Msg.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WriteMessageRequest clone() => WriteMessageRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WriteMessageRequest copyWith(void Function(WriteMessageRequest) updates) => super.copyWith((message) => updates(message as WriteMessageRequest)) as WriteMessageRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static WriteMessageRequest create() => WriteMessageRequest._();
  WriteMessageRequest createEmptyInstance() => create();
  static $pb.PbList<WriteMessageRequest> createRepeated() => $pb.PbList<WriteMessageRequest>();
  @$core.pragma('dart2js:noInline')
  static WriteMessageRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WriteMessageRequest>(create);
  static WriteMessageRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Msg get msg => $_getN(0);
  @$pb.TagNumber(1)
  set msg($1.Msg v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasMsg() => $_has(0);
  @$pb.TagNumber(1)
  void clearMsg() => clearField(1);
  @$pb.TagNumber(1)
  $1.Msg ensureMsg() => $_ensure(0);
}

class WriteMessageResponse extends $pb.GeneratedMessage {
  factory WriteMessageResponse() => create();
  WriteMessageResponse._() : super();
  factory WriteMessageResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory WriteMessageResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'WriteMessageResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  WriteMessageResponse clone() => WriteMessageResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  WriteMessageResponse copyWith(void Function(WriteMessageResponse) updates) => super.copyWith((message) => updates(message as WriteMessageResponse)) as WriteMessageResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static WriteMessageResponse create() => WriteMessageResponse._();
  WriteMessageResponse createEmptyInstance() => create();
  static $pb.PbList<WriteMessageResponse> createRepeated() => $pb.PbList<WriteMessageResponse>();
  @$core.pragma('dart2js:noInline')
  static WriteMessageResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<WriteMessageResponse>(create);
  static WriteMessageResponse? _defaultInstance;
}

class ReadMessageWithPageRequest extends $pb.GeneratedMessage {
  factory ReadMessageWithPageRequest({
    $core.List<$core.int>? uid,
    $core.List<$core.int>? chatId,
    $core.int? offset,
  }) {
    final $result = create();
    if (uid != null) {
      $result.uid = uid;
    }
    if (chatId != null) {
      $result.chatId = chatId;
    }
    if (offset != null) {
      $result.offset = offset;
    }
    return $result;
  }
  ReadMessageWithPageRequest._() : super();
  factory ReadMessageWithPageRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ReadMessageWithPageRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ReadMessageWithPageRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, _omitFieldNames ? '' : 'uid', $pb.PbFieldType.OY)
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'chatId', $pb.PbFieldType.OY)
    ..a<$core.int>(3, _omitFieldNames ? '' : 'offset', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ReadMessageWithPageRequest clone() => ReadMessageWithPageRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ReadMessageWithPageRequest copyWith(void Function(ReadMessageWithPageRequest) updates) => super.copyWith((message) => updates(message as ReadMessageWithPageRequest)) as ReadMessageWithPageRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ReadMessageWithPageRequest create() => ReadMessageWithPageRequest._();
  ReadMessageWithPageRequest createEmptyInstance() => create();
  static $pb.PbList<ReadMessageWithPageRequest> createRepeated() => $pb.PbList<ReadMessageWithPageRequest>();
  @$core.pragma('dart2js:noInline')
  static ReadMessageWithPageRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ReadMessageWithPageRequest>(create);
  static ReadMessageWithPageRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get uid => $_getN(0);
  @$pb.TagNumber(1)
  set uid($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUid() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get chatId => $_getN(1);
  @$pb.TagNumber(2)
  set chatId($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasChatId() => $_has(1);
  @$pb.TagNumber(2)
  void clearChatId() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get offset => $_getIZ(2);
  @$pb.TagNumber(3)
  set offset($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasOffset() => $_has(2);
  @$pb.TagNumber(3)
  void clearOffset() => clearField(3);
}

class ReadMessageWithPageResponse extends $pb.GeneratedMessage {
  factory ReadMessageWithPageResponse() => create();
  ReadMessageWithPageResponse._() : super();
  factory ReadMessageWithPageResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ReadMessageWithPageResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ReadMessageWithPageResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ReadMessageWithPageResponse clone() => ReadMessageWithPageResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ReadMessageWithPageResponse copyWith(void Function(ReadMessageWithPageResponse) updates) => super.copyWith((message) => updates(message as ReadMessageWithPageResponse)) as ReadMessageWithPageResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ReadMessageWithPageResponse create() => ReadMessageWithPageResponse._();
  ReadMessageWithPageResponse createEmptyInstance() => create();
  static $pb.PbList<ReadMessageWithPageResponse> createRepeated() => $pb.PbList<ReadMessageWithPageResponse>();
  @$core.pragma('dart2js:noInline')
  static ReadMessageWithPageResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ReadMessageWithPageResponse>(create);
  static ReadMessageWithPageResponse? _defaultInstance;
}

class ReadMessageWithTimeRequest extends $pb.GeneratedMessage {
  factory ReadMessageWithTimeRequest() => create();
  ReadMessageWithTimeRequest._() : super();
  factory ReadMessageWithTimeRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ReadMessageWithTimeRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ReadMessageWithTimeRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ReadMessageWithTimeRequest clone() => ReadMessageWithTimeRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ReadMessageWithTimeRequest copyWith(void Function(ReadMessageWithTimeRequest) updates) => super.copyWith((message) => updates(message as ReadMessageWithTimeRequest)) as ReadMessageWithTimeRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ReadMessageWithTimeRequest create() => ReadMessageWithTimeRequest._();
  ReadMessageWithTimeRequest createEmptyInstance() => create();
  static $pb.PbList<ReadMessageWithTimeRequest> createRepeated() => $pb.PbList<ReadMessageWithTimeRequest>();
  @$core.pragma('dart2js:noInline')
  static ReadMessageWithTimeRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ReadMessageWithTimeRequest>(create);
  static ReadMessageWithTimeRequest? _defaultInstance;
}

class ReadMessageWithTimeResponse extends $pb.GeneratedMessage {
  factory ReadMessageWithTimeResponse() => create();
  ReadMessageWithTimeResponse._() : super();
  factory ReadMessageWithTimeResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ReadMessageWithTimeResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ReadMessageWithTimeResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ReadMessageWithTimeResponse clone() => ReadMessageWithTimeResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ReadMessageWithTimeResponse copyWith(void Function(ReadMessageWithTimeResponse) updates) => super.copyWith((message) => updates(message as ReadMessageWithTimeResponse)) as ReadMessageWithTimeResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ReadMessageWithTimeResponse create() => ReadMessageWithTimeResponse._();
  ReadMessageWithTimeResponse createEmptyInstance() => create();
  static $pb.PbList<ReadMessageWithTimeResponse> createRepeated() => $pb.PbList<ReadMessageWithTimeResponse>();
  @$core.pragma('dart2js:noInline')
  static ReadMessageWithTimeResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ReadMessageWithTimeResponse>(create);
  static ReadMessageWithTimeResponse? _defaultInstance;
}

class SubscribeMessageRequest extends $pb.GeneratedMessage {
  factory SubscribeMessageRequest() => create();
  SubscribeMessageRequest._() : super();
  factory SubscribeMessageRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SubscribeMessageRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SubscribeMessageRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SubscribeMessageRequest clone() => SubscribeMessageRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SubscribeMessageRequest copyWith(void Function(SubscribeMessageRequest) updates) => super.copyWith((message) => updates(message as SubscribeMessageRequest)) as SubscribeMessageRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SubscribeMessageRequest create() => SubscribeMessageRequest._();
  SubscribeMessageRequest createEmptyInstance() => create();
  static $pb.PbList<SubscribeMessageRequest> createRepeated() => $pb.PbList<SubscribeMessageRequest>();
  @$core.pragma('dart2js:noInline')
  static SubscribeMessageRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SubscribeMessageRequest>(create);
  static SubscribeMessageRequest? _defaultInstance;
}

class SubscribeMessageResponse extends $pb.GeneratedMessage {
  factory SubscribeMessageResponse() => create();
  SubscribeMessageResponse._() : super();
  factory SubscribeMessageResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SubscribeMessageResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'SubscribeMessageResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'app.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  SubscribeMessageResponse clone() => SubscribeMessageResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  SubscribeMessageResponse copyWith(void Function(SubscribeMessageResponse) updates) => super.copyWith((message) => updates(message as SubscribeMessageResponse)) as SubscribeMessageResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static SubscribeMessageResponse create() => SubscribeMessageResponse._();
  SubscribeMessageResponse createEmptyInstance() => create();
  static $pb.PbList<SubscribeMessageResponse> createRepeated() => $pb.PbList<SubscribeMessageResponse>();
  @$core.pragma('dart2js:noInline')
  static SubscribeMessageResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SubscribeMessageResponse>(create);
  static SubscribeMessageResponse? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
