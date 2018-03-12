// Code generated by protoc-gen-go. DO NOT EDIT.
// source: face.proto

/*
Package face_recog is a generated protocol buffer package.

It is generated from these files:
	face.proto

It has these top-level messages:
	IdentifyRequest
	IdentifyResponse
*/
package face_recog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IdentifyRequest struct {
	ImagePath string `protobuf:"bytes,1,opt,name=image_path,json=imagePath" json:"image_path,omitempty"`
}

func (m *IdentifyRequest) Reset()                    { *m = IdentifyRequest{} }
func (m *IdentifyRequest) String() string            { return proto.CompactTextString(m) }
func (*IdentifyRequest) ProtoMessage()               {}
func (*IdentifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdentifyRequest) GetImagePath() string {
	if m != nil {
		return m.ImagePath
	}
	return ""
}

type IdentifyResponse struct {
	PersonId int32 `protobuf:"varint,1,opt,name=person_id,json=personId" json:"person_id,omitempty"`
}

func (m *IdentifyResponse) Reset()                    { *m = IdentifyResponse{} }
func (m *IdentifyResponse) String() string            { return proto.CompactTextString(m) }
func (*IdentifyResponse) ProtoMessage()               {}
func (*IdentifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IdentifyResponse) GetPersonId() int32 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func init() {
	proto.RegisterType((*IdentifyRequest)(nil), "face_recog.IdentifyRequest")
	proto.RegisterType((*IdentifyResponse)(nil), "face_recog.IdentifyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Identify service

type IdentifyClient interface {
	Identify(ctx context.Context, in *IdentifyRequest, opts ...grpc.CallOption) (*IdentifyResponse, error)
}

type identifyClient struct {
	cc *grpc.ClientConn
}

func NewIdentifyClient(cc *grpc.ClientConn) IdentifyClient {
	return &identifyClient{cc}
}

func (c *identifyClient) Identify(ctx context.Context, in *IdentifyRequest, opts ...grpc.CallOption) (*IdentifyResponse, error) {
	out := new(IdentifyResponse)
	err := grpc.Invoke(ctx, "/face_recog.Identify/Identify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Identify service

type IdentifyServer interface {
	Identify(context.Context, *IdentifyRequest) (*IdentifyResponse, error)
}

func RegisterIdentifyServer(s *grpc.Server, srv IdentifyServer) {
	s.RegisterService(&_Identify_serviceDesc, srv)
}

func _Identify_Identify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentifyServer).Identify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/face_recog.Identify/Identify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentifyServer).Identify(ctx, req.(*IdentifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "face_recog.Identify",
	HandlerType: (*IdentifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Identify",
			Handler:    _Identify_Identify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "face.proto",
}

func init() { proto.RegisterFile("face.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4b, 0x4c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0xb3, 0xe3, 0x8b, 0x52, 0x93, 0xf3, 0xd3, 0x95,
	0x0c, 0xb8, 0xf8, 0x3d, 0x53, 0x52, 0xf3, 0x4a, 0x32, 0xd3, 0x2a, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x64, 0xb9, 0xb8, 0x32, 0x73, 0x13, 0xd3, 0x53, 0xe3, 0x0b, 0x12, 0x4b, 0x32,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38, 0xc1, 0x22, 0x01, 0x89, 0x25, 0x19, 0x4a, 0xfa,
	0x5c, 0x02, 0x08, 0x1d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0xd2, 0x5c, 0x9c, 0x05, 0xa9,
	0x45, 0xc5, 0xf9, 0x79, 0xf1, 0x99, 0x29, 0x60, 0x1d, 0xac, 0x41, 0x1c, 0x10, 0x01, 0xcf, 0x14,
	0xa3, 0x60, 0x2e, 0x0e, 0x98, 0x06, 0x21, 0x77, 0x24, 0xb6, 0xb4, 0x1e, 0xc2, 0x1d, 0x7a, 0x68,
	0x8e, 0x90, 0x92, 0xc1, 0x2e, 0x09, 0xb1, 0x4f, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x15, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0x52, 0x79, 0x30, 0xd8, 0x00, 0x00, 0x00,
}
