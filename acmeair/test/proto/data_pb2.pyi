from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

BUSINESS: PhoneType
DESCRIPTOR: _descriptor.FileDescriptor
EXEC_PLATINUM: MemberShipStatus
GOLD: MemberShipStatus
GRAPHITE: MemberShipStatus
HOME: PhoneType
MOBILE: PhoneType
NONE: MemberShipStatus
PLATINUM: MemberShipStatus
SILVER: MemberShipStatus
UNKNOWN: PhoneType

class AddressInfo(_message.Message):
    __slots__ = ["city", "country", "postal_code", "state_province", "street_address1", "street_address2"]
    CITY_FIELD_NUMBER: _ClassVar[int]
    COUNTRY_FIELD_NUMBER: _ClassVar[int]
    POSTAL_CODE_FIELD_NUMBER: _ClassVar[int]
    STATE_PROVINCE_FIELD_NUMBER: _ClassVar[int]
    STREET_ADDRESS1_FIELD_NUMBER: _ClassVar[int]
    STREET_ADDRESS2_FIELD_NUMBER: _ClassVar[int]
    city: str
    country: str
    postal_code: str
    state_province: str
    street_address1: str
    street_address2: str
    def __init__(self, street_address1: _Optional[str] = ..., street_address2: _Optional[str] = ..., city: _Optional[str] = ..., state_province: _Optional[str] = ..., country: _Optional[str] = ..., postal_code: _Optional[str] = ...) -> None: ...

class BookingInfo(_message.Message):
    __slots__ = ["booking_id", "customer", "date_of_booking", "one_way_flight", "return_flight", "to_flight"]
    BOOKING_ID_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_FIELD_NUMBER: _ClassVar[int]
    DATE_OF_BOOKING_FIELD_NUMBER: _ClassVar[int]
    ONE_WAY_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    RETURN_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    TO_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    booking_id: str
    customer: CustomerInfo
    date_of_booking: _timestamp_pb2.Timestamp
    one_way_flight: bool
    return_flight: FlightInfo
    to_flight: FlightInfo
    def __init__(self, booking_id: _Optional[str] = ..., one_way_flight: bool = ..., to_flight: _Optional[_Union[FlightInfo, _Mapping]] = ..., return_flight: _Optional[_Union[FlightInfo, _Mapping]] = ..., customer: _Optional[_Union[CustomerInfo, _Mapping]] = ..., date_of_booking: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CustomerInfo(_message.Message):
    __slots__ = ["address", "customer_id", "miles_ytd", "phone_number", "phone_number_type", "status", "total_miles", "username"]
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    MILES_YTD_FIELD_NUMBER: _ClassVar[int]
    PHONE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PHONE_NUMBER_TYPE_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TOTAL_MILES_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    address: AddressInfo
    customer_id: str
    miles_ytd: int
    phone_number: str
    phone_number_type: PhoneType
    status: MemberShipStatus
    total_miles: int
    username: str
    def __init__(self, customer_id: _Optional[str] = ..., username: _Optional[str] = ..., status: _Optional[_Union[MemberShipStatus, str]] = ..., total_miles: _Optional[int] = ..., miles_ytd: _Optional[int] = ..., address: _Optional[_Union[AddressInfo, _Mapping]] = ..., phone_number: _Optional[str] = ..., phone_number_type: _Optional[_Union[PhoneType, str]] = ...) -> None: ...

class FlightInfo(_message.Message):
    __slots__ = ["airplane_type_id", "economy_class_base_cost", "first_class_base_cost", "flight_id", "flight_segment", "num_economy_class_seats", "num_first_class_seats", "scheduled_arrival_time", "scheduled_departure_time"]
    AIRPLANE_TYPE_ID_FIELD_NUMBER: _ClassVar[int]
    ECONOMY_CLASS_BASE_COST_FIELD_NUMBER: _ClassVar[int]
    FIRST_CLASS_BASE_COST_FIELD_NUMBER: _ClassVar[int]
    FLIGHT_ID_FIELD_NUMBER: _ClassVar[int]
    FLIGHT_SEGMENT_FIELD_NUMBER: _ClassVar[int]
    NUM_ECONOMY_CLASS_SEATS_FIELD_NUMBER: _ClassVar[int]
    NUM_FIRST_CLASS_SEATS_FIELD_NUMBER: _ClassVar[int]
    SCHEDULED_ARRIVAL_TIME_FIELD_NUMBER: _ClassVar[int]
    SCHEDULED_DEPARTURE_TIME_FIELD_NUMBER: _ClassVar[int]
    airplane_type_id: str
    economy_class_base_cost: int
    first_class_base_cost: int
    flight_id: str
    flight_segment: FlightSegmentInfo
    num_economy_class_seats: int
    num_first_class_seats: int
    scheduled_arrival_time: _timestamp_pb2.Timestamp
    scheduled_departure_time: _timestamp_pb2.Timestamp
    def __init__(self, flight_id: _Optional[str] = ..., scheduled_departure_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., scheduled_arrival_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., first_class_base_cost: _Optional[int] = ..., economy_class_base_cost: _Optional[int] = ..., num_first_class_seats: _Optional[int] = ..., num_economy_class_seats: _Optional[int] = ..., airplane_type_id: _Optional[str] = ..., flight_segment: _Optional[_Union[FlightSegmentInfo, _Mapping]] = ...) -> None: ...

class FlightSegmentInfo(_message.Message):
    __slots__ = ["dest_port", "flight_name", "flight_segment", "miles", "origin_port"]
    DEST_PORT_FIELD_NUMBER: _ClassVar[int]
    FLIGHT_NAME_FIELD_NUMBER: _ClassVar[int]
    FLIGHT_SEGMENT_FIELD_NUMBER: _ClassVar[int]
    MILES_FIELD_NUMBER: _ClassVar[int]
    ORIGIN_PORT_FIELD_NUMBER: _ClassVar[int]
    dest_port: str
    flight_name: str
    flight_segment: str
    miles: int
    origin_port: str
    def __init__(self, flight_name: _Optional[str] = ..., flight_segment: _Optional[str] = ..., origin_port: _Optional[str] = ..., dest_port: _Optional[str] = ..., miles: _Optional[int] = ...) -> None: ...

class PhoneType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []

class MemberShipStatus(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
