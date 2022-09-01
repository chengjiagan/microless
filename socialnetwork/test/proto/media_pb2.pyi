from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComposeMediaRequest(_message.Message):
    __slots__ = ["media_ids", "media_types"]
    MEDIA_IDS_FIELD_NUMBER: _ClassVar[int]
    MEDIA_TYPES_FIELD_NUMBER: _ClassVar[int]
    media_ids: _containers.RepeatedScalarFieldContainer[int]
    media_types: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, media_types: _Optional[_Iterable[str]] = ..., media_ids: _Optional[_Iterable[int]] = ...) -> None: ...

class ComposeMediaRespond(_message.Message):
    __slots__ = ["media"]
    MEDIA_FIELD_NUMBER: _ClassVar[int]
    media: _containers.RepeatedCompositeFieldContainer[_data_pb2.Media]
    def __init__(self, media: _Optional[_Iterable[_Union[_data_pb2.Media, _Mapping]]] = ...) -> None: ...
