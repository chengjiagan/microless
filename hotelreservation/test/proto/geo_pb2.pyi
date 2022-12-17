from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AddHotelRequest(_message.Message):
    __slots__ = ["hotel_id", "lat", "lon"]
    HOTEL_ID_FIELD_NUMBER: _ClassVar[int]
    LAT_FIELD_NUMBER: _ClassVar[int]
    LON_FIELD_NUMBER: _ClassVar[int]
    hotel_id: str
    lat: float
    lon: float
    def __init__(self, hotel_id: _Optional[str] = ..., lat: _Optional[float] = ..., lon: _Optional[float] = ...) -> None: ...

class NearbyRequest(_message.Message):
    __slots__ = ["lat", "lon"]
    LAT_FIELD_NUMBER: _ClassVar[int]
    LON_FIELD_NUMBER: _ClassVar[int]
    lat: float
    lon: float
    def __init__(self, lat: _Optional[float] = ..., lon: _Optional[float] = ...) -> None: ...

class NearbyRespond(_message.Message):
    __slots__ = ["hotel_ids"]
    HOTEL_IDS_FIELD_NUMBER: _ClassVar[int]
    hotel_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, hotel_ids: _Optional[_Iterable[str]] = ...) -> None: ...
