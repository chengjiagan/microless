from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComposeUserMentionsRequest(_message.Message):
    __slots__ = ["usernames"]
    USERNAMES_FIELD_NUMBER: _ClassVar[int]
    usernames: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, usernames: _Optional[_Iterable[str]] = ...) -> None: ...

class ComposeUserMentionsRespond(_message.Message):
    __slots__ = ["user_mentions"]
    USER_MENTIONS_FIELD_NUMBER: _ClassVar[int]
    user_mentions: _containers.RepeatedCompositeFieldContainer[_data_pb2.UserMention]
    def __init__(self, user_mentions: _Optional[_Iterable[_Union[_data_pb2.UserMention, _Mapping]]] = ...) -> None: ...
