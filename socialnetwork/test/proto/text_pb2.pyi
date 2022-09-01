from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComposeTextRequest(_message.Message):
    __slots__ = ["text"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class ComposeTextRespond(_message.Message):
    __slots__ = ["text", "urls", "user_mention"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    URLS_FIELD_NUMBER: _ClassVar[int]
    USER_MENTION_FIELD_NUMBER: _ClassVar[int]
    text: str
    urls: _containers.RepeatedCompositeFieldContainer[_data_pb2.Url]
    user_mention: _containers.RepeatedCompositeFieldContainer[_data_pb2.UserMention]
    def __init__(self, text: _Optional[str] = ..., user_mention: _Optional[_Iterable[_Union[_data_pb2.UserMention, _Mapping]]] = ..., urls: _Optional[_Iterable[_Union[_data_pb2.Url, _Mapping]]] = ...) -> None: ...
