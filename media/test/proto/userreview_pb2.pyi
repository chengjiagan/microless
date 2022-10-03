from proto import data_pb2 as _data_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateUserRequest(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...

class ReadUserReviewsRequest(_message.Message):
    __slots__ = ["start", "stop", "user_id"]
    START_FIELD_NUMBER: _ClassVar[int]
    STOP_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    start: int
    stop: int
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., start: _Optional[int] = ..., stop: _Optional[int] = ...) -> None: ...

class ReadUserReviewsRespond(_message.Message):
    __slots__ = ["reviews"]
    REVIEWS_FIELD_NUMBER: _ClassVar[int]
    reviews: _containers.RepeatedCompositeFieldContainer[_data_pb2.Review]
    def __init__(self, reviews: _Optional[_Iterable[_Union[_data_pb2.Review, _Mapping]]] = ...) -> None: ...

class UploadUserReviewRequest(_message.Message):
    __slots__ = ["review_id", "timestamp", "user_id"]
    REVIEW_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    review_id: str
    timestamp: _timestamp_pb2.Timestamp
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., review_id: _Optional[str] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
