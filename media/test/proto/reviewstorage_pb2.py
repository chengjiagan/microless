# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/reviewstorage.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto import data_pb2 as proto_dot_data__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19proto/reviewstorage.proto\x12\x1dmicroless.media.reviewstorage\x1a\x10proto/data.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"U\n\x12StoreReviewRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08movie_id\x18\x02 \x01(\t\x12\x0c\n\x04text\x18\x03 \x01(\t\x12\x0e\n\x06rating\x18\x04 \x01(\x05\"V\n\x12StoreReviewRespond\x12\x11\n\treview_id\x18\x01 \x01(\t\x12-\n\ttimestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"(\n\x12ReadReviewsRequest\x12\x12\n\nreview_ids\x18\x01 \x03(\t\">\n\x12ReadReviewsRespond\x12(\n\x07reviews\x18\x01 \x03(\x0b\x32\x17.microless.media.Review2\x80\x02\n\x14ReviewStorageService\x12s\n\x0bStoreReview\x12\x31.microless.media.reviewstorage.StoreReviewRequest\x1a\x31.microless.media.reviewstorage.StoreReviewRespond\x12s\n\x0bReadReviews\x12\x31.microless.media.reviewstorage.ReadReviewsRequest\x1a\x31.microless.media.reviewstorage.ReadReviewsRespondB%Z#microless/media/proto/reviewstorageb\x06proto3')



_STOREREVIEWREQUEST = DESCRIPTOR.message_types_by_name['StoreReviewRequest']
_STOREREVIEWRESPOND = DESCRIPTOR.message_types_by_name['StoreReviewRespond']
_READREVIEWSREQUEST = DESCRIPTOR.message_types_by_name['ReadReviewsRequest']
_READREVIEWSRESPOND = DESCRIPTOR.message_types_by_name['ReadReviewsRespond']
StoreReviewRequest = _reflection.GeneratedProtocolMessageType('StoreReviewRequest', (_message.Message,), {
  'DESCRIPTOR' : _STOREREVIEWREQUEST,
  '__module__' : 'proto.reviewstorage_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.reviewstorage.StoreReviewRequest)
  })
_sym_db.RegisterMessage(StoreReviewRequest)

StoreReviewRespond = _reflection.GeneratedProtocolMessageType('StoreReviewRespond', (_message.Message,), {
  'DESCRIPTOR' : _STOREREVIEWRESPOND,
  '__module__' : 'proto.reviewstorage_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.reviewstorage.StoreReviewRespond)
  })
_sym_db.RegisterMessage(StoreReviewRespond)

ReadReviewsRequest = _reflection.GeneratedProtocolMessageType('ReadReviewsRequest', (_message.Message,), {
  'DESCRIPTOR' : _READREVIEWSREQUEST,
  '__module__' : 'proto.reviewstorage_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.reviewstorage.ReadReviewsRequest)
  })
_sym_db.RegisterMessage(ReadReviewsRequest)

ReadReviewsRespond = _reflection.GeneratedProtocolMessageType('ReadReviewsRespond', (_message.Message,), {
  'DESCRIPTOR' : _READREVIEWSRESPOND,
  '__module__' : 'proto.reviewstorage_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.reviewstorage.ReadReviewsRespond)
  })
_sym_db.RegisterMessage(ReadReviewsRespond)

_REVIEWSTORAGESERVICE = DESCRIPTOR.services_by_name['ReviewStorageService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z#microless/media/proto/reviewstorage'
  _STOREREVIEWREQUEST._serialized_start=111
  _STOREREVIEWREQUEST._serialized_end=196
  _STOREREVIEWRESPOND._serialized_start=198
  _STOREREVIEWRESPOND._serialized_end=284
  _READREVIEWSREQUEST._serialized_start=286
  _READREVIEWSREQUEST._serialized_end=326
  _READREVIEWSRESPOND._serialized_start=328
  _READREVIEWSRESPOND._serialized_end=390
  _REVIEWSTORAGESERVICE._serialized_start=393
  _REVIEWSTORAGESERVICE._serialized_end=649
# @@protoc_insertion_point(module_scope)
