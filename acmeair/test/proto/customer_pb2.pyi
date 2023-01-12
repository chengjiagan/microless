from google.protobuf import empty_pb2 as _empty_pb2
from google.api import annotations_pb2 as _annotations_pb2
from proto import data_pb2 as _data_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCustomerRequest(_message.Message):
    __slots__ = ["customer_id"]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    customer_id: str
    def __init__(self, customer_id: _Optional[str] = ...) -> None: ...

class GetCustomerRespond(_message.Message):
    __slots__ = ["customer"]
    CUSTOMER_FIELD_NUMBER: _ClassVar[int]
    customer: _data_pb2.CustomerInfo
    def __init__(self, customer: _Optional[_Union[_data_pb2.CustomerInfo, _Mapping]] = ...) -> None: ...

class PutCustomerRequest(_message.Message):
    __slots__ = ["customer", "customer_id"]
    CUSTOMER_FIELD_NUMBER: _ClassVar[int]
    CUSTOMER_ID_FIELD_NUMBER: _ClassVar[int]
    customer: _data_pb2.CustomerInfo
    customer_id: str
    def __init__(self, customer_id: _Optional[str] = ..., customer: _Optional[_Union[_data_pb2.CustomerInfo, _Mapping]] = ...) -> None: ...
