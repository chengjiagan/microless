# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/movieinfo.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15proto/movieinfo.proto\x12\x19microless.media.movieinfo\x1a\x10proto/data.proto\x1a\x1bgoogle/protobuf/empty.proto\"\xc2\x01\n\x15WriteMovieInfoRequest\x12\r\n\x05title\x18\x01 \x01(\t\x12$\n\x05\x63\x61sts\x18\x02 \x03(\x0b\x32\x15.microless.media.Cast\x12\x0f\n\x07plot_id\x18\x03 \x01(\t\x12\x15\n\rthumbnail_ids\x18\x04 \x03(\t\x12\x11\n\tphoto_ids\x18\x05 \x03(\t\x12\x11\n\tvideo_ids\x18\x06 \x03(\t\x12\x12\n\navg_rating\x18\x07 \x01(\x01\x12\x12\n\nnum_rating\x18\x08 \x01(\x05\")\n\x15WriteMovieInfoRespond\x12\x10\n\x08movie_id\x18\x01 \x01(\t\"(\n\x14ReadMovieInfoRequest\x12\x10\n\x08movie_id\x18\x01 \x01(\t\"g\n\x13UpdateRatingRequest\x12\x10\n\x08movie_id\x18\x01 \x01(\t\x12\x1e\n\x16sum_uncommitted_rating\x18\x02 \x01(\x05\x12\x1e\n\x16num_uncommitted_rating\x18\x03 \x01(\x05\x32\xbe\x02\n\x10MovieInfoService\x12t\n\x0eWriteMovieInfo\x12\x30.microless.media.movieinfo.WriteMovieInfoRequest\x1a\x30.microless.media.movieinfo.WriteMovieInfoRespond\x12\\\n\rReadMovieInfo\x12/.microless.media.movieinfo.ReadMovieInfoRequest\x1a\x1a.microless.media.MovieInfo\x12V\n\x0cUpdateRating\x12..microless.media.movieinfo.UpdateRatingRequest\x1a\x16.google.protobuf.EmptyB!Z\x1fmicroless/media/proto/movieinfob\x06proto3')



_WRITEMOVIEINFOREQUEST = DESCRIPTOR.message_types_by_name['WriteMovieInfoRequest']
_WRITEMOVIEINFORESPOND = DESCRIPTOR.message_types_by_name['WriteMovieInfoRespond']
_READMOVIEINFOREQUEST = DESCRIPTOR.message_types_by_name['ReadMovieInfoRequest']
_UPDATERATINGREQUEST = DESCRIPTOR.message_types_by_name['UpdateRatingRequest']
WriteMovieInfoRequest = _reflection.GeneratedProtocolMessageType('WriteMovieInfoRequest', (_message.Message,), {
  'DESCRIPTOR' : _WRITEMOVIEINFOREQUEST,
  '__module__' : 'proto.movieinfo_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.movieinfo.WriteMovieInfoRequest)
  })
_sym_db.RegisterMessage(WriteMovieInfoRequest)

WriteMovieInfoRespond = _reflection.GeneratedProtocolMessageType('WriteMovieInfoRespond', (_message.Message,), {
  'DESCRIPTOR' : _WRITEMOVIEINFORESPOND,
  '__module__' : 'proto.movieinfo_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.movieinfo.WriteMovieInfoRespond)
  })
_sym_db.RegisterMessage(WriteMovieInfoRespond)

ReadMovieInfoRequest = _reflection.GeneratedProtocolMessageType('ReadMovieInfoRequest', (_message.Message,), {
  'DESCRIPTOR' : _READMOVIEINFOREQUEST,
  '__module__' : 'proto.movieinfo_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.movieinfo.ReadMovieInfoRequest)
  })
_sym_db.RegisterMessage(ReadMovieInfoRequest)

UpdateRatingRequest = _reflection.GeneratedProtocolMessageType('UpdateRatingRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATERATINGREQUEST,
  '__module__' : 'proto.movieinfo_pb2'
  # @@protoc_insertion_point(class_scope:microless.media.movieinfo.UpdateRatingRequest)
  })
_sym_db.RegisterMessage(UpdateRatingRequest)

_MOVIEINFOSERVICE = DESCRIPTOR.services_by_name['MovieInfoService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\037microless/media/proto/movieinfo'
  _WRITEMOVIEINFOREQUEST._serialized_start=100
  _WRITEMOVIEINFOREQUEST._serialized_end=294
  _WRITEMOVIEINFORESPOND._serialized_start=296
  _WRITEMOVIEINFORESPOND._serialized_end=337
  _READMOVIEINFOREQUEST._serialized_start=339
  _READMOVIEINFOREQUEST._serialized_end=379
  _UPDATERATINGREQUEST._serialized_start=381
  _UPDATERATINGREQUEST._serialized_end=484
  _MOVIEINFOSERVICE._serialized_start=487
  _MOVIEINFOSERVICE._serialized_end=805
# @@protoc_insertion_point(module_scope)