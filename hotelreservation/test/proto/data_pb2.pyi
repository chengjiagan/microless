from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Address(_message.Message):
    __slots__ = ["city", "country", "lat", "lon", "postal_code", "state", "street_name", "street_number"]
    CITY_FIELD_NUMBER: _ClassVar[int]
    COUNTRY_FIELD_NUMBER: _ClassVar[int]
    LAT_FIELD_NUMBER: _ClassVar[int]
    LON_FIELD_NUMBER: _ClassVar[int]
    POSTAL_CODE_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    STREET_NAME_FIELD_NUMBER: _ClassVar[int]
    STREET_NUMBER_FIELD_NUMBER: _ClassVar[int]
    city: str
    country: str
    lat: float
    lon: float
    postal_code: str
    state: str
    street_name: str
    street_number: str
    def __init__(self, street_number: _Optional[str] = ..., street_name: _Optional[str] = ..., city: _Optional[str] = ..., state: _Optional[str] = ..., country: _Optional[str] = ..., postal_code: _Optional[str] = ..., lat: _Optional[float] = ..., lon: _Optional[float] = ...) -> None: ...

class Hotel(_message.Message):
    __slots__ = ["address", "description", "hotel_id", "images", "name", "phone_number", "room_number"]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    HOTEL_ID_FIELD_NUMBER: _ClassVar[int]
    IMAGES_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PHONE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    ROOM_NUMBER_FIELD_NUMBER: _ClassVar[int]
    address: Address
    description: str
    hotel_id: str
    images: _containers.RepeatedCompositeFieldContainer[Image]
    name: str
    phone_number: str
    room_number: int
    def __init__(self, hotel_id: _Optional[str] = ..., name: _Optional[str] = ..., phone_number: _Optional[str] = ..., description: _Optional[str] = ..., address: _Optional[_Union[Address, _Mapping]] = ..., images: _Optional[_Iterable[_Union[Image, _Mapping]]] = ..., room_number: _Optional[int] = ...) -> None: ...

class Image(_message.Message):
    __slots__ = ["default", "url"]
    DEFAULT_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    default: bool
    url: str
    def __init__(self, url: _Optional[str] = ..., default: bool = ...) -> None: ...
