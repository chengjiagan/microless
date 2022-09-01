from proto import data_pb2 as _data_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReadHomeTimelineRequest(_message.Message):
    __slots__ = ["start", "stop", "user_id"]
    START_FIELD_NUMBER: _ClassVar[int]
    STOP_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    start: int
    stop: int
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., start: _Optional[int] = ..., stop: _Optional[int] = ...) -> None: ...

class ReadHomeTimelineRespond(_message.Message):
    __slots__ = ["posts"]
    POSTS_FIELD_NUMBER: _ClassVar[int]
    posts: _containers.RepeatedCompositeFieldContainer[_data_pb2.Post]
    def __init__(self, posts: _Optional[_Iterable[_Union[_data_pb2.Post, _Mapping]]] = ...) -> None: ...

class WriteHomeTimelineRequest(_message.Message):
    __slots__ = ["post_id", "timestamp", "user_id", "user_mentions_id"]
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    USER_MENTIONS_ID_FIELD_NUMBER: _ClassVar[int]
    post_id: str
    timestamp: _timestamp_pb2.Timestamp
    user_id: str
    user_mentions_id: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, user_id: _Optional[str] = ..., post_id: _Optional[str] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., user_mentions_id: _Optional[_Iterable[str]] = ...) -> None: ...
