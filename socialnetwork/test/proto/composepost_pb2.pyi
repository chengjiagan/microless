from proto import data_pb2 as _data_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComposePostRequest(_message.Message):
    __slots__ = ["media_ids", "media_types", "post_type", "text", "user_id", "username"]
    MEDIA_IDS_FIELD_NUMBER: _ClassVar[int]
    MEDIA_TYPES_FIELD_NUMBER: _ClassVar[int]
    POST_TYPE_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    media_ids: _containers.RepeatedScalarFieldContainer[int]
    media_types: _containers.RepeatedScalarFieldContainer[str]
    post_type: _data_pb2.PostType
    text: str
    user_id: str
    username: str
    def __init__(self, username: _Optional[str] = ..., user_id: _Optional[str] = ..., text: _Optional[str] = ..., media_ids: _Optional[_Iterable[int]] = ..., media_types: _Optional[_Iterable[str]] = ..., post_type: _Optional[_Union[_data_pb2.PostType, str]] = ...) -> None: ...
