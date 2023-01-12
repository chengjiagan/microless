from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.api import annotations_pb2 as _annotations_pb2
from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BrowseFlightsRequest(_message.Message):
    __slots__ = ["from_airport", "one_way_flight", "to_airport"]
    FROM_AIRPORT_FIELD_NUMBER: _ClassVar[int]
    ONE_WAY_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    TO_AIRPORT_FIELD_NUMBER: _ClassVar[int]
    from_airport: str
    one_way_flight: bool
    to_airport: str
    def __init__(self, from_airport: _Optional[str] = ..., to_airport: _Optional[str] = ..., one_way_flight: bool = ...) -> None: ...

class BrowseFlightsRespond(_message.Message):
    __slots__ = ["one_way_flight", "return_flights", "to_flights"]
    ONE_WAY_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    RETURN_FLIGHTS_FIELD_NUMBER: _ClassVar[int]
    TO_FLIGHTS_FIELD_NUMBER: _ClassVar[int]
    one_way_flight: bool
    return_flights: _containers.RepeatedCompositeFieldContainer[_data_pb2.FlightInfo]
    to_flights: _containers.RepeatedCompositeFieldContainer[_data_pb2.FlightInfo]
    def __init__(self, to_flights: _Optional[_Iterable[_Union[_data_pb2.FlightInfo, _Mapping]]] = ..., return_flights: _Optional[_Iterable[_Union[_data_pb2.FlightInfo, _Mapping]]] = ..., one_way_flight: bool = ...) -> None: ...

class GetFlightByIdRequest(_message.Message):
    __slots__ = ["flight_id"]
    FLIGHT_ID_FIELD_NUMBER: _ClassVar[int]
    flight_id: str
    def __init__(self, flight_id: _Optional[str] = ...) -> None: ...

class GetFlightByIdRespond(_message.Message):
    __slots__ = ["flight"]
    FLIGHT_FIELD_NUMBER: _ClassVar[int]
    flight: _data_pb2.FlightInfo
    def __init__(self, flight: _Optional[_Union[_data_pb2.FlightInfo, _Mapping]] = ...) -> None: ...

class GetTripFlightsRequest(_message.Message):
    __slots__ = ["from_airport", "from_date", "one_way_flight", "return_date", "to_airport"]
    FROM_AIRPORT_FIELD_NUMBER: _ClassVar[int]
    FROM_DATE_FIELD_NUMBER: _ClassVar[int]
    ONE_WAY_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    RETURN_DATE_FIELD_NUMBER: _ClassVar[int]
    TO_AIRPORT_FIELD_NUMBER: _ClassVar[int]
    from_airport: str
    from_date: _timestamp_pb2.Timestamp
    one_way_flight: bool
    return_date: _timestamp_pb2.Timestamp
    to_airport: str
    def __init__(self, from_airport: _Optional[str] = ..., to_airport: _Optional[str] = ..., from_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., return_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., one_way_flight: bool = ...) -> None: ...

class GetTripFlightsRespond(_message.Message):
    __slots__ = ["one_way_flight", "return_flights", "to_flights"]
    ONE_WAY_FLIGHT_FIELD_NUMBER: _ClassVar[int]
    RETURN_FLIGHTS_FIELD_NUMBER: _ClassVar[int]
    TO_FLIGHTS_FIELD_NUMBER: _ClassVar[int]
    one_way_flight: bool
    return_flights: _containers.RepeatedCompositeFieldContainer[_data_pb2.FlightInfo]
    to_flights: _containers.RepeatedCompositeFieldContainer[_data_pb2.FlightInfo]
    def __init__(self, to_flights: _Optional[_Iterable[_Union[_data_pb2.FlightInfo, _Mapping]]] = ..., return_flights: _Optional[_Iterable[_Union[_data_pb2.FlightInfo, _Mapping]]] = ..., one_way_flight: bool = ...) -> None: ...
