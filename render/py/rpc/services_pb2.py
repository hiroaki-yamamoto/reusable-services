# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import models_pb2 as models__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='services.proto',
  package='render',
  syntax='proto3',
  serialized_options=b'Z)github.com/hiroaki-yamamoto/render/go/rpc',
  serialized_pb=b'\n\x0eservices.proto\x12\x06render\x1a\x0cmodels.proto2P\n\x0fTemplateService\x12=\n\x06render\x12\x18.render.RenderingRequest\x1a\x19.render.RenderingResponseB+Z)github.com/hiroaki-yamamoto/render/go/rpcb\x06proto3'
  ,
  dependencies=[models__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_TEMPLATESERVICE = _descriptor.ServiceDescriptor(
  name='TemplateService',
  full_name='render.TemplateService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=40,
  serialized_end=120,
  methods=[
  _descriptor.MethodDescriptor(
    name='render',
    full_name='render.TemplateService.render',
    index=0,
    containing_service=None,
    input_type=models__pb2._RENDERINGREQUEST,
    output_type=models__pb2._RENDERINGRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_TEMPLATESERVICE)

DESCRIPTOR.services_by_name['TemplateService'] = _TEMPLATESERVICE

# @@protoc_insertion_point(module_scope)
