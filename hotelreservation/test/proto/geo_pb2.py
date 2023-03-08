# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/geo.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0fproto/geo.proto\x12\x1emicroless.hotelreservation.geo\x1a\x1bgoogle/protobuf/empty.proto\")\n\rNearbyRequest\x12\x0b\n\x03lat\x18\x01 \x01(\x01\x12\x0b\n\x03lon\x18\x02 \x01(\x01\"\"\n\rNearbyRespond\x12\x11\n\thotel_ids\x18\x01 \x03(\t\"=\n\x0f\x41\x64\x64HotelRequest\x12\x10\n\x08hotel_id\x18\x01 \x01(\t\x12\x0b\n\x03lat\x18\x02 \x01(\x01\x12\x0b\n\x03lon\x18\x03 \x01(\x01\x32\xc9\x01\n\nGeoService\x12\x66\n\x06Nearby\x12-.microless.hotelreservation.geo.NearbyRequest\x1a-.microless.hotelreservation.geo.NearbyRespond\x12S\n\x08\x41\x64\x64Hotel\x12/.microless.hotelreservation.geo.AddHotelRequest\x1a\x16.google.protobuf.EmptyB&Z$microless/hotelreservation/proto/geob\x06proto3')



_NEARBYREQUEST = DESCRIPTOR.message_types_by_name['NearbyRequest']
_NEARBYRESPOND = DESCRIPTOR.message_types_by_name['NearbyRespond']
_ADDHOTELREQUEST = DESCRIPTOR.message_types_by_name['AddHotelRequest']
NearbyRequest = _reflection.GeneratedProtocolMessageType('NearbyRequest', (_message.Message,), {
  'DESCRIPTOR' : _NEARBYREQUEST,
  '__module__' : 'proto.geo_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.geo.NearbyRequest)
  })
_sym_db.RegisterMessage(NearbyRequest)

NearbyRespond = _reflection.GeneratedProtocolMessageType('NearbyRespond', (_message.Message,), {
  'DESCRIPTOR' : _NEARBYRESPOND,
  '__module__' : 'proto.geo_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.geo.NearbyRespond)
  })
_sym_db.RegisterMessage(NearbyRespond)

AddHotelRequest = _reflection.GeneratedProtocolMessageType('AddHotelRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDHOTELREQUEST,
  '__module__' : 'proto.geo_pb2'
  # @@protoc_insertion_point(class_scope:microless.hotelreservation.geo.AddHotelRequest)
  })
_sym_db.RegisterMessage(AddHotelRequest)

_GEOSERVICE = DESCRIPTOR.services_by_name['GeoService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z$microless/hotelreservation/proto/geo'
  _NEARBYREQUEST._serialized_start=80
  _NEARBYREQUEST._serialized_end=121
  _NEARBYRESPOND._serialized_start=123
  _NEARBYRESPOND._serialized_end=157
  _ADDHOTELREQUEST._serialized_start=159
  _ADDHOTELREQUEST._serialized_end=220
  _GEOSERVICE._serialized_start=223
  _GEOSERVICE._serialized_end=424
# @@protoc_insertion_point(module_scope)