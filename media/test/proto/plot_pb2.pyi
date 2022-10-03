from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ReadPlotRequest(_message.Message):
    __slots__ = ["plot_id"]
    PLOT_ID_FIELD_NUMBER: _ClassVar[int]
    plot_id: str
    def __init__(self, plot_id: _Optional[str] = ...) -> None: ...

class ReadPlotRespond(_message.Message):
    __slots__ = ["plot"]
    PLOT_FIELD_NUMBER: _ClassVar[int]
    plot: str
    def __init__(self, plot: _Optional[str] = ...) -> None: ...

class WritePlotRequest(_message.Message):
    __slots__ = ["plot"]
    PLOT_FIELD_NUMBER: _ClassVar[int]
    plot: str
    def __init__(self, plot: _Optional[str] = ...) -> None: ...

class WritePlotRespond(_message.Message):
    __slots__ = ["plot_id"]
    PLOT_ID_FIELD_NUMBER: _ClassVar[int]
    plot_id: str
    def __init__(self, plot_id: _Optional[str] = ...) -> None: ...
