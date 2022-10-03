from proto import data_pb2 as _data_pb2
from google.api import annotations_pb2 as _annotations_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReadPageRequest(_message.Message):
    __slots__ = ["movie_id", "review_start", "review_stop"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    REVIEW_START_FIELD_NUMBER: _ClassVar[int]
    REVIEW_STOP_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    review_start: int
    review_stop: int
    def __init__(self, movie_id: _Optional[str] = ..., review_start: _Optional[int] = ..., review_stop: _Optional[int] = ...) -> None: ...

class ReadPageRespond(_message.Message):
    __slots__ = ["cast_infos", "movie_info", "plot", "reviews"]
    CAST_INFOS_FIELD_NUMBER: _ClassVar[int]
    MOVIE_INFO_FIELD_NUMBER: _ClassVar[int]
    PLOT_FIELD_NUMBER: _ClassVar[int]
    REVIEWS_FIELD_NUMBER: _ClassVar[int]
    cast_infos: _containers.RepeatedCompositeFieldContainer[_data_pb2.CastInfo]
    movie_info: _data_pb2.MovieInfo
    plot: str
    reviews: _containers.RepeatedCompositeFieldContainer[_data_pb2.Review]
    def __init__(self, movie_info: _Optional[_Union[_data_pb2.MovieInfo, _Mapping]] = ..., cast_infos: _Optional[_Iterable[_Union[_data_pb2.CastInfo, _Mapping]]] = ..., plot: _Optional[str] = ..., reviews: _Optional[_Iterable[_Union[_data_pb2.Review, _Mapping]]] = ...) -> None: ...
