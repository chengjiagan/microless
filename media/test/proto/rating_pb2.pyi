from google.protobuf import empty_pb2 as _empty_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class UploadRatingRequest(_message.Message):
    __slots__ = ["movie_id", "rating"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    rating: int
    def __init__(self, movie_id: _Optional[str] = ..., rating: _Optional[int] = ...) -> None: ...
