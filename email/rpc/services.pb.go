// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SendRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	TxtBody              string   `protobuf:"bytes,2,opt,name=txtBody,proto3" json:"txtBody,omitempty"`
	HtmlBody             string   `protobuf:"bytes,3,opt,name=htmlBody,proto3" json:"htmlBody,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SendRequest) GetTxtBody() string {
	if m != nil {
		return m.TxtBody
	}
	return ""
}

func (m *SendRequest) GetHtmlBody() string {
	if m != nil {
		return m.HtmlBody
	}
	return ""
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "SendRequest")
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0x92,
	0x4a, 0x91, 0x5c, 0xdc, 0xc1, 0xa9, 0x79, 0x29, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x22, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10,
	0x8e, 0x90, 0x04, 0x17, 0x7b, 0x49, 0x45, 0x89, 0x53, 0x7e, 0x4a, 0xa5, 0x04, 0x13, 0x58, 0x1c,
	0xc6, 0x15, 0x92, 0xe2, 0xe2, 0xc8, 0x28, 0xc9, 0xcd, 0x01, 0x4b, 0x31, 0x83, 0xa5, 0xe0, 0x7c,
	0x23, 0x53, 0x2e, 0x56, 0x57, 0xb0, 0x76, 0x1d, 0x2e, 0x96, 0xe2, 0xd4, 0xbc, 0x14, 0x21, 0x1e,
	0x3d, 0x24, 0xab, 0xa4, 0xc4, 0xf4, 0x20, 0xee, 0xd2, 0x83, 0xb9, 0x4b, 0xcf, 0x15, 0xe4, 0xae,
	0x24, 0x36, 0x30, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x10, 0xba, 0x84, 0xb1, 0xc7, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type emailClient struct {
	cc *grpc.ClientConn
}

func NewEmailClient(cc *grpc.ClientConn) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Email/send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
type EmailServer interface {
	Send(context.Context, *SendRequest) (*empty.Empty, error)
}

// UnimplementedEmailServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServer struct {
}

func (*UnimplementedEmailServer) Send(ctx context.Context, req *SendRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterEmailServer(s *grpc.Server, srv EmailServer) {
	s.RegisterService(&_Email_serviceDesc, srv)
}

func _Email_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Email/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Email_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _Email_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
