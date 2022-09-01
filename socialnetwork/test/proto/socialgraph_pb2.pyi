from google.protobuf import empty_pb2 as _empty_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class FollowRequest(_message.Message):
    __slots__ = ["followee_id", "user_id"]
    FOLLOWEE_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    followee_id: str
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., followee_id: _Optional[str] = ...) -> None: ...

class FollowWithUsernameRequest(_message.Message):
    __slots__ = ["followee_username", "user_username"]
    FOLLOWEE_USERNAME_FIELD_NUMBER: _ClassVar[int]
    USER_USERNAME_FIELD_NUMBER: _ClassVar[int]
    followee_username: str
    user_username: str
    def __init__(self, user_username: _Optional[str] = ..., followee_username: _Optional[str] = ...) -> None: ...

class GetFolloweesRequest(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...

class GetFolloweesRespond(_message.Message):
    __slots__ = ["followees_id"]
    FOLLOWEES_ID_FIELD_NUMBER: _ClassVar[int]
    followees_id: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, followees_id: _Optional[_Iterable[str]] = ...) -> None: ...

class GetFollowersRequest(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...

class GetFollowersRespond(_message.Message):
    __slots__ = ["followers_id"]
    FOLLOWERS_ID_FIELD_NUMBER: _ClassVar[int]
    followers_id: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, followers_id: _Optional[_Iterable[str]] = ...) -> None: ...

class InsertUserRequest(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...

class UnfollowRequest(_message.Message):
    __slots__ = ["followee_id", "user_id"]
    FOLLOWEE_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    followee_id: str
    user_id: str
    def __init__(self, user_id: _Optional[str] = ..., followee_id: _Optional[str] = ...) -> None: ...

class UnfollowWithUsernameRequest(_message.Message):
    __slots__ = ["followee_username", "user_username"]
    FOLLOWEE_USERNAME_FIELD_NUMBER: _ClassVar[int]
    USER_USERNAME_FIELD_NUMBER: _ClassVar[int]
    followee_username: str
    user_username: str
    def __init__(self, user_username: _Optional[str] = ..., followee_username: _Optional[str] = ...) -> None: ...
