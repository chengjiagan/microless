from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReadCastInfoRequest(_message.Message):
    __slots__ = ["cast_ids"]
    CAST_IDS_FIELD_NUMBER: _ClassVar[int]
    cast_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, cast_ids: _Optional[_Iterable[str]] = ...) -> None: ...

class ReadCastInfoRespond(_message.Message):
    __slots__ = ["cast_infos"]
    CAST_INFOS_FIELD_NUMBER: _ClassVar[int]
    cast_infos: _containers.RepeatedCompositeFieldContainer[_data_pb2.CastInfo]
    def __init__(self, cast_infos: _Optional[_Iterable[_Union[_data_pb2.CastInfo, _Mapping]]] = ...) -> None: ...

class WriteCastInfoRequest(_message.Message):
    __slots__ = ["gender", "intro", "name"]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    INTRO_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    gender: bool
    intro: str
    name: str
    def __init__(self, name: _Optional[str] = ..., gender: bool = ..., intro: _Optional[str] = ...) -> None: ...

class WriteCastInfoRespond(_message.Message):
    __slots__ = ["cast_info_id"]
    CAST_INFO_ID_FIELD_NUMBER: _ClassVar[int]
    cast_info_id: str
    def __init__(self, cast_info_id: _Optional[str] = ...) -> None: ...
