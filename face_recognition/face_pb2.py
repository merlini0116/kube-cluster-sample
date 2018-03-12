# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: face.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='face.proto',
  package='facerecog',
  syntax='proto3',
  serialized_pb=_b('\n\nface.proto\x12\tfacerecog\"%\n\x0fIdentifyRequest\x12\x12\n\nimage_path\x18\x01 \x01(\t\"%\n\x10IdentifyResponse\x12\x11\n\tperson_id\x18\x01 \x01(\x05\x32Q\n\x08Identify\x12\x45\n\x08Identify\x12\x1a.facerecog.IdentifyRequest\x1a\x1b.facerecog.IdentifyResponse\"\x00\x62\x06proto3')
)




_IDENTIFYREQUEST = _descriptor.Descriptor(
  name='IdentifyRequest',
  full_name='facerecog.IdentifyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='image_path', full_name='facerecog.IdentifyRequest.image_path', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=25,
  serialized_end=62,
)


_IDENTIFYRESPONSE = _descriptor.Descriptor(
  name='IdentifyResponse',
  full_name='facerecog.IdentifyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='person_id', full_name='facerecog.IdentifyResponse.person_id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=64,
  serialized_end=101,
)

DESCRIPTOR.message_types_by_name['IdentifyRequest'] = _IDENTIFYREQUEST
DESCRIPTOR.message_types_by_name['IdentifyResponse'] = _IDENTIFYRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

IdentifyRequest = _reflection.GeneratedProtocolMessageType('IdentifyRequest', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFYREQUEST,
  __module__ = 'face_pb2'
  # @@protoc_insertion_point(class_scope:facerecog.IdentifyRequest)
  ))
_sym_db.RegisterMessage(IdentifyRequest)

IdentifyResponse = _reflection.GeneratedProtocolMessageType('IdentifyResponse', (_message.Message,), dict(
  DESCRIPTOR = _IDENTIFYRESPONSE,
  __module__ = 'face_pb2'
  # @@protoc_insertion_point(class_scope:facerecog.IdentifyResponse)
  ))
_sym_db.RegisterMessage(IdentifyResponse)



_IDENTIFY = _descriptor.ServiceDescriptor(
  name='Identify',
  full_name='facerecog.Identify',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=103,
  serialized_end=184,
  methods=[
  _descriptor.MethodDescriptor(
    name='Identify',
    full_name='facerecog.Identify.Identify',
    index=0,
    containing_service=None,
    input_type=_IDENTIFYREQUEST,
    output_type=_IDENTIFYRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_IDENTIFY)

DESCRIPTOR.services_by_name['Identify'] = _IDENTIFY

# @@protoc_insertion_point(module_scope)
