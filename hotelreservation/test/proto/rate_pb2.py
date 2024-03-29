# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/rate.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10proto/rate.proto\x12\x1fmicroless.hotelreservation.rate\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/api/annotations.proto\"$\n\x0fGetRatesRequest\x12\x11\n\thotel_ids\x18\x01 \x03(\t\"L\n\x0fGetRatesRespond\x12\x39\n\x05rates\x18\x01 \x03(\x0b\x32*.microless.hotelreservation.rate.HotelRate\"\x9c\x01\n\x0e\x41\x64\x64RateRequest\x12\x10\n\x08hotel_id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12+\n\x07in_date\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08out_date\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0c\n\x04rate\x18\x05 \x01(\x05\"+\n\tHotelRate\x12\x10\n\x08hotel_id\x18\x01 \x01(\t\x12\x0c\n\x04rate\x18\x02 \x01(\x01\x32\xea\x01\n\x0bRateService\x12n\n\x08GetRates\x12\x30.microless.hotelreservation.rate.GetRatesRequest\x1a\x30.microless.hotelreservation.rate.GetRatesRespond\x12k\n\x07\x41\x64\x64Rate\x12/.microless.hotelreservation.rate.AddRateRequest\x1a\x16.google.protobuf.Empty\"\x17\x82\xd3\xe4\x93\x02\x11\"\x0c/api/v1/rate:\x01*B\'Z%microless/hotelreservation/proto/rateb\x06proto3')



_GETRATESREQUEST = DESCRIPTOR.message_types_by_name['GetRatesRequest']
_GETRATESRESPOND = DESCRIPTOR.message_types_by_name['GetRatesRespond']
_ADDRATEREQUEST = DESCRIPTOR.message_types_by_name['AddRateRequest']
_HOTELRATE = DESCRIPTOR.message_types_by_name['HotelRate']
GetRatesRequest = _reflection.GeneratedProtocolMessageType('GetRatesRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETRATESREQUEST,
  '__module__' : 'proto.rate_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.rate.GetRatesRequest)
  })
_sym_db.RegisterMessage(GetRatesRequest)

GetRatesRespond = _reflection.GeneratedProtocolMessageType('GetRatesRespond', (_message.Message,), {
  'DESCRIPTOR' : _GETRATESRESPOND,
  '__module__' : 'proto.rate_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.rate.GetRatesRespond)
  })
_sym_db.RegisterMessage(GetRatesRespond)

AddRateRequest = _reflection.GeneratedProtocolMessageType('AddRateRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDRATEREQUEST,
  '__module__' : 'proto.rate_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.rate.AddRateRequest)
  })
_sym_db.RegisterMessage(AddRateRequest)

HotelRate = _reflection.GeneratedProtocolMessageType('HotelRate', (_message.Message,), {
  'DESCRIPTOR' : _HOTELRATE,
  '__module__' : 'proto.rate_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.rate.HotelRate)
  })
_sym_db.RegisterMessage(HotelRate)

_RATESERVICE = DESCRIPTOR.services_by_name['RateService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z%microless/hotelreservation/proto/rate'
  _RATESERVICE.methods_by_name['AddRate']._options = None
  _RATESERVICE.methods_by_name['AddRate']._serialized_options = b'\202\323\344\223\002\021\"\014/api/v1/rate:\001*'
  _GETRATESREQUEST._serialized_start=145
  _GETRATESREQUEST._serialized_end=181
  _GETRATESRESPOND._serialized_start=183
  _GETRATESRESPOND._serialized_end=259
  _ADDRATEREQUEST._serialized_start=262
  _ADDRATEREQUEST._serialized_end=418
  _HOTELRATE._serialized_start=420
  _HOTELRATE._serialized_end=463
  _RATESERVICE._serialized_start=466
  _RATESERVICE._serialized_end=700
# @@protoc_insertion_point(module_scope)
