from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
test1: TestEnum
test2: TestEnum
test3: TestEnum

class Notification(_message.Message):
    __slots__ = ["message", "subscribtion_id", "test_enum", "time", "times"]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    SUBSCRIBTION_ID_FIELD_NUMBER: _ClassVar[int]
    TEST_ENUM_FIELD_NUMBER: _ClassVar[int]
    TIMES_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    message: str
    subscribtion_id: int
    test_enum: TestEnum
    time: int
    times: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, subscribtion_id: _Optional[int] = ..., message: _Optional[str] = ..., time: _Optional[int] = ..., times: _Optional[_Iterable[int]] = ..., test_enum: _Optional[_Union[TestEnum, str]] = ...) -> None: ...

class SubscribeRequest(_message.Message):
    __slots__ = ["name", "subscribtion_id"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    SUBSCRIBTION_ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    subscribtion_id: int
    def __init__(self, subscribtion_id: _Optional[int] = ..., name: _Optional[str] = ...) -> None: ...

class TestEnum(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
