from proto import data_pb2 as _data_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateMovieRequest(_message.Message):
    __slots__ = ["movie_id"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    def __init__(self, movie_id: _Optional[str] = ...) -> None: ...

class ReadMovieReviewsRequest(_message.Message):
    __slots__ = ["movie_id", "start", "stop"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    START_FIELD_NUMBER: _ClassVar[int]
    STOP_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    start: int
    stop: int
    def __init__(self, movie_id: _Optional[str] = ..., start: _Optional[int] = ..., stop: _Optional[int] = ...) -> None: ...

class ReadMovieReviewsRespond(_message.Message):
    __slots__ = ["reviews"]
    REVIEWS_FIELD_NUMBER: _ClassVar[int]
    reviews: _containers.RepeatedCompositeFieldContainer[_data_pb2.Review]
    def __init__(self, reviews: _Optional[_Iterable[_Union[_data_pb2.Review, _Mapping]]] = ...) -> None: ...

class UploadMovieReviewRequest(_message.Message):
    __slots__ = ["movie_id", "review_id", "timestamp"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    REVIEW_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    review_id: str
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, movie_id: _Optional[str] = ..., review_id: _Optional[str] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
