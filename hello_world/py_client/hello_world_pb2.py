# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello_world.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11hello_world.proto\x12\x05proto\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1e\n\rHelloResponse\x12\r\n\x05reply\x18\x01 \x01(\t2B\n\x07Greeter\x12\x37\n\x08SayHello\x12\x13.proto.HelloRequest\x1a\x14.proto.HelloResponse\"\x00\x42\x11Z\x0fpy_client/protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'hello_world_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\017py_client/proto'
  _HELLOREQUEST._serialized_start=28
  _HELLOREQUEST._serialized_end=56
  _HELLORESPONSE._serialized_start=58
  _HELLORESPONSE._serialized_end=88
  _GREETER._serialized_start=90
  _GREETER._serialized_end=156
# @@protoc_insertion_point(module_scope)