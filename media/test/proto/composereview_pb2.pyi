from google.protobuf import empty_pb2 as _empty_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ComposeReviewRequest(_message.Message):
    __slots__ = ["movie_id", "rating", "text", "user_id"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    rating: int
    text: str
    user_id: str
    def __init__(self, movie_id: _Optional[str] = ..., user_id: _Optional[str] = ..., text: _Optional[str] = ..., rating: _Optional[int] = ...) -> None: ...
