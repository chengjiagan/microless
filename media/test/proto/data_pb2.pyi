from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Cast(_message.Message):
    __slots__ = ["cast_id", "cast_info_id", "character"]
    CAST_ID_FIELD_NUMBER: _ClassVar[int]
    CAST_INFO_ID_FIELD_NUMBER: _ClassVar[int]
    CHARACTER_FIELD_NUMBER: _ClassVar[int]
    cast_id: int
    cast_info_id: str
    character: str
    def __init__(self, cast_id: _Optional[int] = ..., character: _Optional[str] = ..., cast_info_id: _Optional[str] = ...) -> None: ...

class CastInfo(_message.Message):
    __slots__ = ["cast_info_id", "gender", "intro", "name"]
    CAST_INFO_ID_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    INTRO_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    cast_info_id: str
    gender: bool
    intro: str
    name: str
    def __init__(self, cast_info_id: _Optional[str] = ..., name: _Optional[str] = ..., gender: bool = ..., intro: _Optional[str] = ...) -> None: ...

class MovieInfo(_message.Message):
    __slots__ = ["avg_rating", "casts", "movie_id", "num_rating", "photo_ids", "plot_id", "thumbnail_ids", "title", "video_ids"]
    AVG_RATING_FIELD_NUMBER: _ClassVar[int]
    CASTS_FIELD_NUMBER: _ClassVar[int]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    NUM_RATING_FIELD_NUMBER: _ClassVar[int]
    PHOTO_IDS_FIELD_NUMBER: _ClassVar[int]
    PLOT_ID_FIELD_NUMBER: _ClassVar[int]
    THUMBNAIL_IDS_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    VIDEO_IDS_FIELD_NUMBER: _ClassVar[int]
    avg_rating: float
    casts: _containers.RepeatedCompositeFieldContainer[Cast]
    movie_id: str
    num_rating: int
    photo_ids: _containers.RepeatedScalarFieldContainer[str]
    plot_id: str
    thumbnail_ids: _containers.RepeatedScalarFieldContainer[str]
    title: str
    video_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, movie_id: _Optional[str] = ..., title: _Optional[str] = ..., casts: _Optional[_Iterable[_Union[Cast, _Mapping]]] = ..., plot_id: _Optional[str] = ..., thumbnail_ids: _Optional[_Iterable[str]] = ..., photo_ids: _Optional[_Iterable[str]] = ..., video_ids: _Optional[_Iterable[str]] = ..., avg_rating: _Optional[float] = ..., num_rating: _Optional[int] = ...) -> None: ...

class Review(_message.Message):
    __slots__ = ["movie_id", "rating", "review_id", "text", "timestamp", "user_id"]
    MOVIE_ID_FIELD_NUMBER: _ClassVar[int]
    RATING_FIELD_NUMBER: _ClassVar[int]
    REVIEW_ID_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    movie_id: str
    rating: int
    review_id: str
    text: str
    timestamp: _timestamp_pb2.Timestamp
    user_id: str
    def __init__(self, review_id: _Optional[str] = ..., user_id: _Optional[str] = ..., text: _Optional[str] = ..., movie_id: _Optional[str] = ..., rating: _Optional[int] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
