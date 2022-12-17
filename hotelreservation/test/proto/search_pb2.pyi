from proto import data_pb2 as _data_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SearchRequest(_message.Message):
    __slots__ = ["in_date", "lat", "lon", "out_date", "room_number"]
    IN_DATE_FIELD_NUMBER: _ClassVar[int]
    LAT_FIELD_NUMBER: _ClassVar[int]
    LON_FIELD_NUMBER: _ClassVar[int]
    OUT_DATE_FIELD_NUMBER: _ClassVar[int]
    ROOM_NUMBER_FIELD_NUMBER: _ClassVar[int]
    in_date: _timestamp_pb2.Timestamp
    lat: float
    lon: float
    out_date: _timestamp_pb2.Timestamp
    room_number: int
    def __init__(self, lat: _Optional[float] = ..., lon: _Optional[float] = ..., in_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., out_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., room_number: _Optional[int] = ...) -> None: ...

class SearchRespond(_message.Message):
    __slots__ = ["hotels"]
    HOTELS_FIELD_NUMBER: _ClassVar[int]
    hotels: _containers.RepeatedCompositeFieldContainer[_data_pb2.Hotel]
    def __init__(self, hotels: _Optional[_Iterable[_Union[_data_pb2.Hotel, _Mapping]]] = ...) -> None: ...
