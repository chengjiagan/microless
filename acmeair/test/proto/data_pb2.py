# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/data.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10proto/data.proto\x12\x11microless.acmeair\x1a\x1fgoogle/protobuf/timestamp.proto\"\x89\x02\n\x0b\x42ookingInfo\x12\x12\n\nbooking_id\x18\x01 \x01(\t\x12\x16\n\x0eone_way_flight\x18\x02 \x01(\x08\x12\x30\n\tto_flight\x18\x03 \x01(\x0b\x32\x1d.microless.acmeair.FlightInfo\x12\x34\n\rreturn_flight\x18\x04 \x01(\x0b\x32\x1d.microless.acmeair.FlightInfo\x12\x31\n\x08\x63ustomer\x18\x05 \x01(\x0b\x32\x1f.microless.acmeair.CustomerInfo\x12\x33\n\x0f\x64\x61te_of_booking\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x92\x02\n\x0c\x43ustomerInfo\x12\x13\n\x0b\x63ustomer_id\x18\x01 \x01(\t\x12\x10\n\x08username\x18\x02 \x01(\t\x12\x33\n\x06status\x18\x03 \x01(\x0e\x32#.microless.acmeair.MemberShipStatus\x12\x13\n\x0btotal_miles\x18\x04 \x01(\x05\x12\x11\n\tmiles_ytd\x18\x05 \x01(\x05\x12/\n\x07\x61\x64\x64ress\x18\x06 \x01(\x0b\x32\x1e.microless.acmeair.AddressInfo\x12\x14\n\x0cphone_number\x18\x07 \x01(\t\x12\x37\n\x11phone_number_type\x18\x08 \x01(\x0e\x32\x1c.microless.acmeair.PhoneType\"\x8b\x01\n\x0b\x41\x64\x64ressInfo\x12\x17\n\x0fstreet_address1\x18\x01 \x01(\t\x12\x17\n\x0fstreet_address2\x18\x02 \x01(\t\x12\x0c\n\x04\x63ity\x18\x03 \x01(\t\x12\x16\n\x0estate_province\x18\x04 \x01(\t\x12\x0f\n\x07\x63ountry\x18\x05 \x01(\t\x12\x13\n\x0bpostal_code\x18\x06 \x01(\t\"\xf1\x02\n\nFlightInfo\x12\x11\n\tflight_id\x18\x01 \x01(\t\x12<\n\x18scheduled_departure_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12:\n\x16scheduled_arrival_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1d\n\x15\x66irst_class_base_cost\x18\x05 \x01(\x03\x12\x1f\n\x17\x65\x63onomy_class_base_cost\x18\x06 \x01(\x03\x12\x1d\n\x15num_first_class_seats\x18\x07 \x01(\x05\x12\x1f\n\x17num_economy_class_seats\x18\x08 \x01(\x05\x12\x18\n\x10\x61irplane_type_id\x18\t \x01(\t\x12<\n\x0e\x66light_segment\x18\n \x01(\x0b\x32$.microless.acmeair.FlightSegmentInfo\"w\n\x11\x46lightSegmentInfo\x12\x13\n\x0b\x66light_name\x18\x01 \x01(\t\x12\x16\n\x0e\x66light_segment\x18\x02 \x01(\t\x12\x13\n\x0borigin_port\x18\x03 \x01(\t\x12\x11\n\tdest_port\x18\x04 \x01(\t\x12\r\n\x05miles\x18\x05 \x01(\x05*<\n\tPhoneType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04HOME\x10\x01\x12\x0c\n\x08\x42USINESS\x10\x02\x12\n\n\x06MOBILE\x10\x03*a\n\x10MemberShipStatus\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06SILVER\x10\x01\x12\x08\n\x04GOLD\x10\x02\x12\x0c\n\x08PLATINUM\x10\x03\x12\x11\n\rEXEC_PLATINUM\x10\x04\x12\x0c\n\x08GRAPHITE\x10\x05\x42\x19Z\x17microless/acmeair/protob\x06proto3')

_PHONETYPE = DESCRIPTOR.enum_types_by_name['PhoneType']
PhoneType = enum_type_wrapper.EnumTypeWrapper(_PHONETYPE)
_MEMBERSHIPSTATUS = DESCRIPTOR.enum_types_by_name['MemberShipStatus']
MemberShipStatus = enum_type_wrapper.EnumTypeWrapper(_MEMBERSHIPSTATUS)
UNKNOWN = 0
HOME = 1
BUSINESS = 2
MOBILE = 3
NONE = 0
SILVER = 1
GOLD = 2
PLATINUM = 3
EXEC_PLATINUM = 4
GRAPHITE = 5


_BOOKINGINFO = DESCRIPTOR.message_types_by_name['BookingInfo']
_CUSTOMERINFO = DESCRIPTOR.message_types_by_name['CustomerInfo']
_ADDRESSINFO = DESCRIPTOR.message_types_by_name['AddressInfo']
_FLIGHTINFO = DESCRIPTOR.message_types_by_name['FlightInfo']
_FLIGHTSEGMENTINFO = DESCRIPTOR.message_types_by_name['FlightSegmentInfo']
BookingInfo = _reflection.GeneratedProtocolMessageType('BookingInfo', (_message.Message,), {
  'DESCRIPTOR' : _BOOKINGINFO,
  '__module__' : 'proto.data_pb2'
  # @@protoc_insertion_point(class_scope:microless.acmeair.BookingInfo)
  })
_sym_db.RegisterMessage(BookingInfo)

CustomerInfo = _reflection.GeneratedProtocolMessageType('CustomerInfo', (_message.Message,), {
  'DESCRIPTOR' : _CUSTOMERINFO,
  '__module__' : 'proto.data_pb2'
  # @@protoc_insertion_point(class_scope:microless.acmeair.CustomerInfo)
  })
_sym_db.RegisterMessage(CustomerInfo)

AddressInfo = _reflection.GeneratedProtocolMessageType('AddressInfo', (_message.Message,), {
  'DESCRIPTOR' : _ADDRESSINFO,
  '__module__' : 'proto.data_pb2'
  # @@protoc_insertion_point(class_scope:microless.acmeair.AddressInfo)
  })
_sym_db.RegisterMessage(AddressInfo)

FlightInfo = _reflection.GeneratedProtocolMessageType('FlightInfo', (_message.Message,), {
  'DESCRIPTOR' : _FLIGHTINFO,
  '__module__' : 'proto.data_pb2'
  # @@protoc_insertion_point(class_scope:microless.acmeair.FlightInfo)
  })
_sym_db.RegisterMessage(FlightInfo)

FlightSegmentInfo = _reflection.GeneratedProtocolMessageType('FlightSegmentInfo', (_message.Message,), {
  'DESCRIPTOR' : _FLIGHTSEGMENTINFO,
  '__module__' : 'proto.data_pb2'
  # @@protoc_insertion_point(class_scope:microless.acmeair.FlightSegmentInfo)
  })
_sym_db.RegisterMessage(FlightSegmentInfo)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\027microless/acmeair/proto'
  _PHONETYPE._serialized_start=1252
  _PHONETYPE._serialized_end=1312
  _MEMBERSHIPSTATUS._serialized_start=1314
  _MEMBERSHIPSTATUS._serialized_end=1411
  _BOOKINGINFO._serialized_start=73
  _BOOKINGINFO._serialized_end=338
  _CUSTOMERINFO._serialized_start=341
  _CUSTOMERINFO._serialized_end=615
  _ADDRESSINFO._serialized_start=618
  _ADDRESSINFO._serialized_end=757
  _FLIGHTINFO._serialized_start=760
  _FLIGHTINFO._serialized_end=1129
  _FLIGHTSEGMENTINFO._serialized_start=1131
  _FLIGHTSEGMENTINFO._serialized_end=1250
# @@protoc_insertion_point(module_scope)