from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
DM: PostType
POST: PostType
REPLY: PostType
REPOST: PostType

class Creator(_message.Message):
    __slots__ = ["user_id", "username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    username: str
    def __init__(self, user_id: _Optional[str] = ..., username: _Optional[str] = ...) -> None: ...

class Media(_message.Message):
    __slots__ = ["media_id", "media_type"]
    MEDIA_ID_FIELD_NUMBER: _ClassVar[int]
    MEDIA_TYPE_FIELD_NUMBER: _ClassVar[int]
    media_id: int
    media_type: str
    def __init__(self, media_id: _Optional[int] = ..., media_type: _Optional[str] = ...) -> None: ...

class Post(_message.Message):
    __slots__ = ["creator", "media", "post_type", "text", "timestamp", "urls", "user_mentions"]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    MEDIA_FIELD_NUMBER: _ClassVar[int]
    POST_TYPE_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    URLS_FIELD_NUMBER: _ClassVar[int]
    USER_MENTIONS_FIELD_NUMBER: _ClassVar[int]
    creator: Creator
    media: _containers.RepeatedCompositeFieldContainer[Media]
    post_type: PostType
    text: str
    timestamp: _timestamp_pb2.Timestamp
    urls: _containers.RepeatedCompositeFieldContainer[Url]
    user_mentions: _containers.RepeatedCompositeFieldContainer[UserMention]
    def __init__(self, creator: _Optional[_Union[Creator, _Mapping]] = ..., text: _Optional[str] = ..., user_mentions: _Optional[_Iterable[_Union[UserMention, _Mapping]]] = ..., media: _Optional[_Iterable[_Union[Media, _Mapping]]] = ..., urls: _Optional[_Iterable[_Union[Url, _Mapping]]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., post_type: _Optional[_Union[PostType, str]] = ...) -> None: ...

class Url(_message.Message):
    __slots__ = ["expanded_url", "shortened_url"]
    EXPANDED_URL_FIELD_NUMBER: _ClassVar[int]
    SHORTENED_URL_FIELD_NUMBER: _ClassVar[int]
    expanded_url: str
    shortened_url: str
    def __init__(self, shortened_url: _Optional[str] = ..., expanded_url: _Optional[str] = ...) -> None: ...

class UserMention(_message.Message):
    __slots__ = ["user_id", "username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    username: str
    def __init__(self, user_id: _Optional[str] = ..., username: _Optional[str] = ...) -> None: ...

class PostType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
