from proto import data_pb2 as _data_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReadReviewsRequest(_message.Message):
    __slots__ = ["review_ids"]
    REVIEW_IDS_FIELD_NUMBER: _ClassVar[int]
    review_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, review_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class ReadReviewsRespond(_message.Message):
    __slots__ = ["reviews"]
    REVIEWS_FIELD_NUMBER: _ClassVar[int]
    reviews: _containers.RepeatedCompositeFieldContainer[_data_pb2.Review]
    def __init__(self, reviews: _Optional[_Iterable[_Union[_data_pb2.Review, _Mapping]]] = ...) -> None: ...

class StoreReviewRequest(_message.Message):
    __slots__ = ["movie_id", "rating", "text", "user_id"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    rating: int
    text: str
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., movie_id: _Optional[str] = ..., text: _Optional[str] = ..., rating: _Optional[int] = ...) -> None: ...

class StoreReviewRespond(_message.Message):
    __slots__ = ["review_id", "timestamp"]
    REVIEW_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    review_id: str
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, review_id: _Optional[str] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
