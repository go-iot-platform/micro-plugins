// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/go-iot-platform/go-micro/errors/proto/errors.proto

package errors

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

type Error struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c88d4fdb61ee9f, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Error) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "errors.Error")
}

func init() {
	proto.RegisterFile("github.com/go-iot-platform/go-micro/errors/proto/errors.proto", fileDescriptor_65c88d4fdb61ee9f)
}

var fileDescriptor_65c88d4fdb61ee9f = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0xce, 0xcc, 0xd5, 0x4f, 0xcf, 0xd7, 0xcd,
	0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0xd6, 0x2f, 0x28, 0xca, 0x2f,
	0x81, 0x71, 0xf4, 0xc0, 0x1c, 0x21, 0x36, 0x08, 0x4f, 0x29, 0x9a, 0x8b, 0xd5, 0x15, 0xc4, 0x12,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3,
	0x85, 0xc4, 0xb8, 0xd8, 0x52, 0x52, 0x4b, 0x12, 0x33, 0x73, 0x24, 0x98, 0xc1, 0xea, 0xa0, 0x3c,
	0x90, 0x78, 0x71, 0x49, 0x62, 0x49, 0x69, 0xb1, 0x04, 0x0b, 0x44, 0x1c, 0xc2, 0x4b, 0x62, 0x03,
	0xdb, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xf1, 0xe3, 0xd1, 0xa1, 0x00, 0x00, 0x00,
}
