# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/hometimeline.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto import data_pb2 as proto_dot_data__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18proto/hometimeline.proto\x12$microless.socialnetwork.hometimeline\x1a\x10proto/data.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/api/annotations.proto\"G\n\x17ReadHomeTimelineRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\r\n\x05start\x18\x02 \x01(\x05\x12\x0c\n\x04stop\x18\x03 \x01(\x05\"G\n\x17ReadHomeTimelineRespond\x12,\n\x05posts\x18\x01 \x03(\x0b\x32\x1d.microless.socialnetwork.Post\"\x85\x01\n\x18WriteHomeTimelineRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0f\n\x07post_id\x18\x02 \x01(\t\x12-\n\ttimestamp\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x18\n\x10user_mentions_id\x18\x04 \x03(\t2\xbd\x02\n\x13HomeTimelineService\x12\xb8\x01\n\x10ReadHomeTimeline\x12=.microless.socialnetwork.hometimeline.ReadHomeTimelineRequest\x1a=.microless.socialnetwork.hometimeline.ReadHomeTimelineRespond\"&\x82\xd3\xe4\x93\x02 \x12\x1e/api/v1/hometimeline/{user_id}\x12k\n\x11WriteHomeTimeline\x12>.microless.socialnetwork.hometimeline.WriteHomeTimelineRequest\x1a\x16.google.protobuf.EmptyB,Z*microless/socialnetwork/proto/hometimelineb\x06proto3')



_READHOMETIMELINEREQUEST = DESCRIPTOR.message_types_by_name['ReadHomeTimelineRequest']
_READHOMETIMELINERESPOND = DESCRIPTOR.message_types_by_name['ReadHomeTimelineRespond']
_WRITEHOMETIMELINEREQUEST = DESCRIPTOR.message_types_by_name['WriteHomeTimelineRequest']
ReadHomeTimelineRequest = _reflection.GeneratedProtocolMessageType('ReadHomeTimelineRequest', (_message.Message,), {
  'DESCRIPTOR' : _READHOMETIMELINEREQUEST,
  '__module__' : 'proto.hometimeline_pb2'
  # @@protoc_insertion_point(class_scope:microless.socialnetwork.hometimeline.ReadHomeTimelineRequest)
  })
_sym_db.RegisterMessage(ReadHomeTimelineRequest)

ReadHomeTimelineRespond = _reflection.GeneratedProtocolMessageType('ReadHomeTimelineRespond', (_message.Message,), {
  'DESCRIPTOR' : _READHOMETIMELINERESPOND,
  '__module__' : 'proto.hometimeline_pb2'
  # @@protoc_insertion_point(class_scope:microless.socialnetwork.hometimeline.ReadHomeTimelineRespond)
  })
_sym_db.RegisterMessage(ReadHomeTimelineRespond)

WriteHomeTimelineRequest = _reflection.GeneratedProtocolMessageType('WriteHomeTimelineRequest', (_message.Message,), {
  'DESCRIPTOR' : _WRITEHOMETIMELINEREQUEST,
  '__module__' : 'proto.hometimeline_pb2'
  # @@protoc_insertion_point(class_scope:microless.socialnetwork.hometimeline.WriteHomeTimelineRequest)
  })
_sym_db.RegisterMessage(WriteHomeTimelineRequest)

_HOMETIMELINESERVICE = DESCRIPTOR.services_by_name['HomeTimelineService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z*microless/socialnetwork/proto/hometimeline'
  _HOMETIMELINESERVICE.methods_by_name['ReadHomeTimeline']._options = None
  _HOMETIMELINESERVICE.methods_by_name['ReadHomeTimeline']._serialized_options = b'\202\323\344\223\002 \022\036/api/v1/hometimeline/{user_id}'
  _READHOMETIMELINEREQUEST._serialized_start=176
  _READHOMETIMELINEREQUEST._serialized_end=247
  _READHOMETIMELINERESPOND._serialized_start=249
  _READHOMETIMELINERESPOND._serialized_end=320
  _WRITEHOMETIMELINEREQUEST._serialized_start=323
  _WRITEHOMETIMELINEREQUEST._serialized_end=456
  _HOMETIMELINESERVICE._serialized_start=459
  _HOMETIMELINESERVICE._serialized_end=776
# @@protoc_insertion_point(module_scope)
