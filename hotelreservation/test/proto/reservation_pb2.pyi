from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CheckAvailabilityRequest(_message.Message):
    __slots__ = ["hotel_ids", "in_date", "out_date", "room_number"]
    HOTEL_IDS_FIELD_NUMBER: _ClassVar[int]
    IN_DATE_FIELD_NUMBER: _ClassVar[int]
    OUT_DATE_FIELD_NUMBER: _ClassVar[int]
    ROOM_NUMBER_FIELD_NUMBER: _ClassVar[int]
    hotel_ids: _containers.RepeatedScalarFieldContainer[str]
    in_date: _timestamp_pb2.Timestamp
    out_date: _timestamp_pb2.Timestamp
    room_number: int
    def __init__(self, hotel_ids: _Optional[_Iterable[str]] = ..., in_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., out_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., room_number: _Optional[int] = ...) -> None: ...

class CheckAvailabilityRespond(_message.Message):
    __slots__ = ["hotel_ids"]
    HOTEL_IDS_FIELD_NUMBER: _ClassVar[int]
    hotel_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, hotel_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class MakeReservationRequest(_message.Message):
    __slots__ = ["hotel_id", "in_date", "out_date", "room_number", "user_id"]
    HOTEL_ID_FIELD_NUMBER: _ClassVar[int]
    IN_DATE_FIELD_NUMBER: _ClassVar[int]
    OUT_DATE_FIELD_NUMBER: _ClassVar[int]
    ROOM_NUMBER_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    hotel_id: str
    in_date: _timestamp_pb2.Timestamp
    out_date: _timestamp_pb2.Timestamp
    room_number: int
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., hotel_id: _Optional[str] = ..., in_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., out_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., room_number: _Optional[int] = ...) -> None: ...
