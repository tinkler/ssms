//
//  Generated code. Do not modify.
//  source: mrz/v1/mrz.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/any.pb.dart' as $0;

class Model extends $pb.GeneratedMessage {
  factory Model() => create();
  Model._() : super();
  factory Model.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory Model.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Model',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'mrz.v1'),
      createEmptyInstance: create)
    ..aOM<$0.Any>(1, _omitFieldNames ? '' : 'data', subBuilder: $0.Any.create)
    ..aOM<$0.Any>(2, _omitFieldNames ? '' : 'Args',
        protoName: 'Args', subBuilder: $0.Any.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Model clone() => Model()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Model copyWith(void Function(Model) updates) =>
      super.copyWith((message) => updates(message as Model)) as Model;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Model create() => Model._();
  Model createEmptyInstance() => create();
  static $pb.PbList<Model> createRepeated() => $pb.PbList<Model>();
  @$core.pragma('dart2js:noInline')
  static Model getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Model>(create);
  static Model? _defaultInstance;

  @$pb.TagNumber(1)
  $0.Any get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($0.Any v) {
    setField(1, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $0.Any ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.Any get args => $_getN(1);
  @$pb.TagNumber(2)
  set args($0.Any v) {
    setField(2, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasArgs() => $_has(1);
  @$pb.TagNumber(2)
  void clearArgs() => clearField(2);
  @$pb.TagNumber(2)
  $0.Any ensureArgs() => $_ensure(1);
}

class Res extends $pb.GeneratedMessage {
  factory Res() => create();
  Res._() : super();
  factory Res.fromBuffer($core.List<$core.int> i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(i, r);
  factory Res.fromJson($core.String i,
          [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Res',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'mrz.v1'),
      createEmptyInstance: create)
    ..aOM<$0.Any>(1, _omitFieldNames ? '' : 'data', subBuilder: $0.Any.create)
    ..aOM<$0.Any>(2, _omitFieldNames ? '' : 'resp', subBuilder: $0.Any.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
      'Will be removed in next major version')
  Res clone() => Res()..mergeFromMessage(this);
  @$core.Deprecated('Using this can add significant overhead to your binary. '
      'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
      'Will be removed in next major version')
  Res copyWith(void Function(Res) updates) =>
      super.copyWith((message) => updates(message as Res)) as Res;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Res create() => Res._();
  Res createEmptyInstance() => create();
  static $pb.PbList<Res> createRepeated() => $pb.PbList<Res>();
  @$core.pragma('dart2js:noInline')
  static Res getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Res>(create);
  static Res? _defaultInstance;

  @$pb.TagNumber(1)
  $0.Any get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($0.Any v) {
    setField(1, v);
  }

  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $0.Any ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $0.Any get resp => $_getN(1);
  @$pb.TagNumber(2)
  set resp($0.Any v) {
    setField(2, v);
  }

  @$pb.TagNumber(2)
  $core.bool hasResp() => $_has(1);
  @$pb.TagNumber(2)
  void clearResp() => clearField(2);
  @$pb.TagNumber(2)
  $0.Any ensureResp() => $_ensure(1);
}

const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
