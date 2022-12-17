from proto import data_pb2 as _data_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddProfileRequest(_message.Message):
    __slots__ = ["hotel"]
    HOTEL_FIELD_NUMBER: _ClassVar[int]
    hotel: _data_pb2.Hotel
    def __init__(self, hotel: _Optional[_Union[_data_pb2.Hotel, _Mapping]] = ...) -> None: ...

class AddProfileRespond(_message.Message):
    __slots__ = ["hotel_id"]
    HOTEL_ID_FIELD_NUMBER: _ClassVar[int]
    hotel_id: str
    def __init__(self, hotel_id: _Optional[str] = ...) -> None: ...

class GetProfilesRequest(_message.Message):
    __slots__ = ["hotel_ids"]
    HOTEL_IDS_FIELD_NUMBER: _ClassVar[int]
    hotel_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, hotel_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class GetProfilesRespond(_message.Message):
    __slots__ = ["hotels"]
    HOTELS_FIELD_NUMBER: _ClassVar[int]
    hotels: _containers.RepeatedCompositeFieldContainer[_data_pb2.Hotel]
    def __init__(self, hotels: _Optional[_Iterable[_Union[_data_pb2.Hotel, _Mapping]]] = ...) -> None: ...

class GetRoomNumberRequest(_message.Message):
    __slots__ = ["hotel_id"]
    HOTEL_ID_FIELD_NUMBER: _ClassVar[int]
    hotel_id: str
    def __init__(self, hotel_id: _Optional[str] = ...) -> None: ...

class GetRoomNumberRespond(_message.Message):
    __slots__ = ["room_number"]
    ROOM_NUMBER_FIELD_NUMBER: _ClassVar[int]
    room_number: int
    def __init__(self, room_number: _Optional[int] = ...) -> None: ...
