# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='models.proto',
  package='render',
  syntax='proto3',
  serialized_options=b'Z)github.com/hiroaki-yamamoto/render/go/rpc',
  serialized_pb=b'\n\x0cmodels.proto\x12\x06render\"8\n\x10RenderingRequest\x12\x0f\n\x07tmpName\x18\x01 \x01(\t\x12\x13\n\x0b\x61rgumentMap\x18\x02 \x01(\x0c\"!\n\x11RenderingResponse\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\tB+Z)github.com/hiroaki-yamamoto/render/go/rpcb\x06proto3'
)




_RENDERINGREQUEST = _descriptor.Descriptor(
  name='RenderingRequest',
  full_name='render.RenderingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='tmpName', full_name='render.RenderingRequest.tmpName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='argumentMap', full_name='render.RenderingRequest.argumentMap', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=80,
)


_RENDERINGRESPONSE = _descriptor.Descriptor(
  name='RenderingResponse',
  full_name='render.RenderingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='render.RenderingResponse.data', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=82,
  serialized_end=115,
)

DESCRIPTOR.message_types_by_name['RenderingRequest'] = _RENDERINGREQUEST
DESCRIPTOR.message_types_by_name['RenderingResponse'] = _RENDERINGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RenderingRequest = _reflection.GeneratedProtocolMessageType('RenderingRequest', (_message.Message,), {
  'DESCRIPTOR' : _RENDERINGREQUEST,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:render.RenderingRequest)
  })
_sym_db.RegisterMessage(RenderingRequest)

RenderingResponse = _reflection.GeneratedProtocolMessageType('RenderingResponse', (_message.Message,), {
  'DESCRIPTOR' : _RENDERINGRESPONSE,
  '__module__' : 'models_pb2'
  # @@protoc_insertion_point(class_scope:render.RenderingResponse)
  })
_sym_db.RegisterMessage(RenderingResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)