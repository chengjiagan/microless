from proto import data_pb2 as _data_pb2
from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ReadMovieInfoRequest(_message.Message):
    __slots__ = ["movie_id"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    def __init__(self, movie_id: _Optional[str] = ...) -> None: ...

class UpdateRatingRequest(_message.Message):
    __slots__ = ["movie_id", "num_uncommitted_rating", "sum_uncommitted_rating"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    NUM_UNCOMMITTED_RATING_FIELD_NUMBER: _ClassVar[int]
    SUM_UNCOMMITTED_RATING_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    num_uncommitted_rating: int
    sum_uncommitted_rating: int
    def __init__(self, movie_id: _Optional[str] = ..., sum_uncommitted_rating: _Optional[int] = ..., num_uncommitted_rating: _Optional[int] = ...) -> None: ...

class WriteMovieInfoRequest(_message.Message):
    __slots__ = ["avg_rating", "casts", "num_rating", "photo_ids", "plot_id", "thumbnail_ids", "title", "video_ids"]
    AVG_RATING_FIELD_NUMBER: _ClassVar[int]
    CASTS_FIELD_NUMBER: _ClassVar[int]
    NUM_RATING_FIELD_NUMBER: _ClassVar[int]
    PHOTO_IDS_FIELD_NUMBER: _ClassVar[int]
    PLOT_ID_FIELD_NUMBER: _ClassVar[int]
    THUMBNAIL_IDS_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    VIDEO_IDS_FIELD_NUMBER: _ClassVar[int]
    avg_rating: float
    casts: _containers.RepeatedCompositeFieldContainer[_data_pb2.Cast]
    num_rating: int
    photo_ids: _containers.RepeatedScalarFieldContainer[str]
    plot_id: str
    thumbnail_ids: _containers.RepeatedScalarFieldContainer[str]
    title: str
    video_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, title: _Optional[str] = ..., casts: _Optional[_Iterable[_Union[_data_pb2.Cast, _Mapping]]] = ..., plot_id: _Optional[str] = ..., thumbnail_ids: _Optional[_Iterable[str]] = ..., photo_ids: _Optional[_Iterable[str]] = ..., video_ids: _Optional[_Iterable[str]] = ..., avg_rating: _Optional[float] = ..., num_rating: _Optional[int] = ...) -> None: ...

class WriteMovieInfoRespond(_message.Message):
    __slots__ = ["movie_id"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    def __init__(self, movie_id: _Optional[str] = ...) -> None: ...
