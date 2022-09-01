# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/composepost.proto
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
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17proto/composepost.proto\x12#microless.socialnetwork.composepost\x1a\x10proto/data.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/api/annotations.proto\"\xa3\x01\n\x12\x43omposePostRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x0c\n\x04text\x18\x03 \x01(\t\x12\x11\n\tmedia_ids\x18\x04 \x03(\x03\x12\x13\n\x0bmedia_types\x18\x05 \x03(\t\x12\x34\n\tpost_type\x18\x06 \x01(\x0e\x32!.microless.socialnetwork.PostType2\x94\x01\n\x12\x43omposePostService\x12~\n\x0b\x43omposePost\x12\x37.microless.socialnetwork.composepost.ComposePostRequest\x1a\x16.google.protobuf.Empty\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/api/v1/composepost:\x01*B+Z)microless/socialnetwork/proto/composepostb\x06proto3')



_COMPOSEPOSTREQUEST = DESCRIPTOR.message_types_by_name['ComposePostRequest']
ComposePostRequest = _reflection.GeneratedProtocolMessageType('ComposePostRequest', (_message.Message,), {
  'DESCRIPTOR' : _COMPOSEPOSTREQUEST,
  '__module__' : 'proto.composepost_pb2'
  # @@protoc_insertion_point(class_scope:microless.socialnetwork.composepost.ComposePostRequest)
  })
_sym_db.RegisterMessage(ComposePostRequest)

_COMPOSEPOSTSERVICE = DESCRIPTOR.services_by_name['ComposePostService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z)microless/socialnetwork/proto/composepost'
  _COMPOSEPOSTSERVICE.methods_by_name['ComposePost']._options = None
  _COMPOSEPOSTSERVICE.methods_by_name['ComposePost']._serialized_options = b'\202\323\344\223\002\030\"\023/api/v1/composepost:\001*'
  _COMPOSEPOSTREQUEST._serialized_start=142
  _COMPOSEPOSTREQUEST._serialized_end=305
  _COMPOSEPOSTSERVICE._serialized_start=308
  _COMPOSEPOSTSERVICE._serialized_end=456
# @@protoc_insertion_point(module_scope)
