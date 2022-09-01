from proto import data_pb2 as _data_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComposeCreatorWithUserIdRequest(_message.Message):
    __slots__ = ["user_id", "username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    username: str
    def __init__(self, user_id: _Optional[str] = ..., username: _Optional[str] = ...) -> None: ...

class ComposeCreatorWithUserIdRespond(_message.Message):
    __slots__ = ["creator"]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    creator: _data_pb2.Creator
    def __init__(self, creator: _Optional[_Union[_data_pb2.Creator, _Mapping]] = ...) -> None: ...

class ComposeCreatorWithUsernameRequest(_message.Message):
    __slots__ = ["username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class ComposeCreatorWithUsernameRespond(_message.Message):
    __slots__ = ["creator"]
    CREATOR_FIELD_NUMBER: _ClassVar[int]
    creator: _data_pb2.Creator
    def __init__(self, creator: _Optional[_Union[_data_pb2.Creator, _Mapping]] = ...) -> None: ...

class GetUserIdRequest(_message.Message):
    __slots__ = ["username"]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class GetUserIdRespond(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...

class LoginRequest(_message.Message):
    __slots__ = ["password", "username"]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    password: str
    username: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class LoginRespond(_message.Message):
    __slots__ = ["token"]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    token: str
    def __init__(self, token: _Optional[str] = ...) -> None: ...

class RegisterUserRequest(_message.Message):
    __slots__ = ["first_name", "last_name", "password", "username"]
    FIRST_NAME_FIELD_NUMBER: _ClassVar[int]
    LAST_NAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    first_name: str
    last_name: str
    password: str
    username: str
    def __init__(self, first_name: _Optional[str] = ..., last_name: _Optional[str] = ..., username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...

class RegisterUserRespond(_message.Message):
    __slots__ = ["user_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    def __init__(self, user_id: _Optional[str] = ...) -> None: ...
