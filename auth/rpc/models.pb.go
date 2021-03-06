// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AuthInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	IsActive             bool     `protobuf:"varint,7,opt,name=isActive,proto3" json:"isActive,omitempty"`
	LockReason           string   `protobuf:"bytes,8,opt,name=lockReason,proto3" json:"lockReason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthInfo) Reset()         { *m = AuthInfo{} }
func (m *AuthInfo) String() string { return proto.CompactTextString(m) }
func (*AuthInfo) ProtoMessage()    {}
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}

func (m *AuthInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthInfo.Unmarshal(m, b)
}
func (m *AuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthInfo.Marshal(b, m, deterministic)
}
func (m *AuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfo.Merge(m, src)
}
func (m *AuthInfo) XXX_Size() int {
	return xxx_messageInfo_AuthInfo.Size(m)
}
func (m *AuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfo proto.InternalMessageInfo

func (m *AuthInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthInfo) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *AuthInfo) GetLockReason() string {
	if m != nil {
		return m.LockReason
	}
	return ""
}

type AccessToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessToken) Reset()         { *m = AccessToken{} }
func (m *AccessToken) String() string { return proto.CompactTextString(m) }
func (*AccessToken) ProtoMessage()    {}
func (*AccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{1}
}

func (m *AccessToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessToken.Unmarshal(m, b)
}
func (m *AccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessToken.Marshal(b, m, deterministic)
}
func (m *AccessToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessToken.Merge(m, src)
}
func (m *AccessToken) XXX_Size() int {
	return xxx_messageInfo_AccessToken.Size(m)
}
func (m *AccessToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessToken.DiscardUnknown(m)
}

var xxx_messageInfo_AccessToken proto.InternalMessageInfo

func (m *AccessToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthInfo)(nil), "auth.models.AuthInfo")
	proto.RegisterType((*AccessToken)(nil), "auth.models.AccessToken")
}

func init() { proto.RegisterFile("models.proto", fileDescriptor_0b5431a010549573) }

var fileDescriptor_0b5431a010549573 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x2c, 0x2d, 0xc9, 0xd0, 0x83,
	0x08, 0x29, 0x75, 0x30, 0x72, 0x71, 0x38, 0x96, 0x96, 0x64, 0x78, 0xe6, 0xa5, 0xe5, 0x0b, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x49,
	0x71, 0x71, 0x94, 0x16, 0xa7, 0x16, 0xf9, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45, 0xe1, 0x7c,
	0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc, 0x1c, 0x09, 0x66, 0xb0, 0x04, 0x84, 0x03, 0xd2,
	0x91, 0x59, 0xec, 0x98, 0x5c, 0x92, 0x59, 0x96, 0x2a, 0xc1, 0xae, 0xc0, 0xa8, 0xc1, 0x11, 0x04,
	0xe7, 0x0b, 0xc9, 0x71, 0x71, 0xe5, 0xe4, 0x27, 0x67, 0x07, 0xa5, 0x26, 0x16, 0xe7, 0xe7, 0x49,
	0x70, 0x80, 0xb5, 0x21, 0x89, 0x28, 0x29, 0x73, 0x71, 0x3b, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x87,
	0xe4, 0x67, 0xa7, 0xe6, 0x81, 0x2c, 0x28, 0x01, 0x31, 0xa0, 0xee, 0x81, 0x70, 0x9c, 0x58, 0xa3,
	0x98, 0x8b, 0x0a, 0x92, 0x93, 0xd8, 0xc0, 0x5e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0xf3, 0x73, 0x06, 0xda, 0x00, 0x00, 0x00,
}
