from proto import data_pb2 as _data_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReadPostsRequest(_message.Message):
    __slots__ = ["post_ids"]
    POST_IDS_FIELD_NUMBER: _ClassVar[int]
    post_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, post_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class ReadPostsRespond(_message.Message):
    __slots__ = ["posts"]
    POSTS_FIELD_NUMBER: _ClassVar[int]
    posts: _containers.RepeatedCompositeFieldContainer[_data_pb2.Post]
    def __init__(self, posts: _Optional[_Iterable[_Union[_data_pb2.Post, _Mapping]]] = ...) -> None: ...

class StorePostRequest(_message.Message):
    __slots__ = ["post"]
    POST_FIELD_NUMBER: _ClassVar[int]
    post: _data_pb2.Post
    def __init__(self, post: _Optional[_Union[_data_pb2.Post, _Mapping]] = ...) -> None: ...

class StorePostRespond(_message.Message):
    __slots__ = ["post_id", "timestamp"]
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    post_id: str
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, post_id: _Optional[str] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
