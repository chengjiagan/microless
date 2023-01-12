from google.protobuf import empty_pb2 as _empty_pb2
from google.api import annotations_pb2 as _annotations_pb2
from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BookFlightsRequest(_message.Message):
    __slots__ = ["customer_id", "one_way_flight", "ret_flight_id", "to_flight_id"]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    ONE_WAY_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    RET_FLIGHT_ID_FIELD_NUMBER: _ClassVar[int]
    TO_FLIGHT_ID_FIELD_NUMBER: _ClassVar[int]
    customer_id: str
    one_way_flight: bool
    ret_flight_id: str
    to_flight_id: str
    def __init__(self, customer_id: _Optional[str] = ..., to_flight_id: _Optional[str] = ..., ret_flight_id: _Optional[str] = ..., one_way_flight: bool = ...) -> None: ...

class BookFlightsRespond(_message.Message):
    __slots__ = ["booking"]
    BOOKING_FIELD_NUMBER: _ClassVar[int]
    booking: _data_pb2.BookingInfo
    def __init__(self, booking: _Optional[_Union[_data_pb2.BookingInfo, _Mapping]] = ...) -> None: ...

class CancelBookingByIdRequest(_message.Message):
    __slots__ = ["booking_id", "customer_id"]
    BOOKING_ID_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    booking_id: str
    customer_id: str
    def __init__(self, booking_id: _Optional[str] = ..., customer_id: _Optional[str] = ...) -> None: ...

class GetBookingByIdRequest(_message.Message):
    __slots__ = ["booking_id"]
    BOOKING_ID_FIELD_NUMBER: _ClassVar[int]
    booking_id: str
    def __init__(self, booking_id: _Optional[str] = ...) -> None: ...

class GetBookingByIdRespond(_message.Message):
    __slots__ = ["booking"]
    BOOKING_FIELD_NUMBER: _ClassVar[int]
    booking: _data_pb2.BookingInfo
    def __init__(self, booking: _Optional[_Union[_data_pb2.BookingInfo, _Mapping]] = ...) -> None: ...

class GetBookingByUserRequest(_message.Message):
    __slots__ = ["customer_id"]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    customer_id: str
    def __init__(self, customer_id: _Optional[str] = ...) -> None: ...

class GetBookingByUserRespond(_message.Message):
    __slots__ = ["bookings"]
    BOOKINGS_FIELD_NUMBER: _ClassVar[int]
    bookings: _containers.RepeatedCompositeFieldContainer[_data_pb2.BookingInfo]
    def __init__(self, bookings: _Optional[_Iterable[_Union[_data_pb2.BookingInfo, _Mapping]]] = ...) -> None: ...
