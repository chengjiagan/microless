from proto import data_pb2 as _data_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComposeUrlsRequest(_message.Message):
    __slots__ = ["urls"]
    URLS_FIELD_NUMBER: _ClassVar[int]
    urls: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, urls: _Optional[_Iterable[str]] = ...) -> None: ...

class ComposeUrlsRespond(_message.Message):
    __slots__ = ["urls"]
    URLS_FIELD_NUMBER: _ClassVar[int]
    urls: _containers.RepeatedCompositeFieldContainer[_data_pb2.Url]
    def __init__(self, urls: _Optional[_Iterable[_Union[_data_pb2.Url, _Mapping]]] = ...) -> None: ...

class GetExtendedUrlsRequest(_message.Message):
    __slots__ = ["shortened_urls"]
    SHORTENED_URLS_FIELD_NUMBER: _ClassVar[int]
    shortened_urls: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, shortened_urls: _Optional[_Iterable[str]] = ...) -> None: ...

class GetExtendedUrlsRespond(_message.Message):
    __slots__ = ["urls"]
    URLS_FIELD_NUMBER: _ClassVar[int]
    urls: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, urls: _Optional[_Iterable[str]] = ...) -> None: ...
