// Code generated by protoc-gen-go. DO NOT EDIT.
// source: natsrpc.proto

package natsrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type MethodType int32

const (
	MethodType_A MethodType = 0
	MethodType_S MethodType = 1
	MethodType_P MethodType = 2
)

var MethodType_name = map[int32]string{
	0: "A",
	1: "S",
	2: "P",
}

var MethodType_value = map[string]int32{
	"A": 0,
	"S": 1,
	"P": 2,
}

func (x MethodType) String() string {
	return proto.EnumName(MethodType_name, int32(x))
}

func (MethodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_210d289958f39804, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_210d289958f39804, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Int32 struct {
	Val                  int32    `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_210d289958f39804, []int{1}
}

func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (m *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(m, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type String struct {
	Val                  string   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_210d289958f39804, []int{2}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

var E_ServceAsync = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         43230,
	Name:          "natsrpc.servceAsync",
	Tag:           "varint,43230,opt,name=servceAsync",
	Filename:      "natsrpc.proto",
}

var E_ClientAsync = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         43231,
	Name:          "natsrpc.clientAsync",
	Tag:           "varint,43231,opt,name=clientAsync",
	Filename:      "natsrpc.proto",
}

var E_MethodType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*MethodType)(nil),
	Field:         2360,
	Name:          "natsrpc.methodType",
	Tag:           "varint,2360,opt,name=methodType,enum=natsrpc.MethodType",
	Filename:      "natsrpc.proto",
}

func init() {
	proto.RegisterEnum("natsrpc.MethodType", MethodType_name, MethodType_value)
	proto.RegisterType((*Empty)(nil), "natsrpc.Empty")
	proto.RegisterType((*Int32)(nil), "natsrpc.Int32")
	proto.RegisterType((*String)(nil), "natsrpc.String")
	proto.RegisterExtension(E_ServceAsync)
	proto.RegisterExtension(E_ClientAsync)
	proto.RegisterExtension(E_MethodType)
}

func init() {
	proto.RegisterFile("natsrpc.proto", fileDescriptor_210d289958f39804)
}

var fileDescriptor_210d289958f39804 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4b, 0x2c, 0x29,
	0x2e, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x14, 0xd2,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xc2, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9,
	0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0xa5, 0x4a, 0xec, 0x5c, 0xac, 0xae, 0xb9, 0x05, 0x25,
	0x95, 0x4a, 0x92, 0x5c, 0xac, 0x9e, 0x79, 0x25, 0xc6, 0x46, 0x42, 0x02, 0x5c, 0xcc, 0x65, 0x89,
	0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x20, 0xa6, 0x92, 0x14, 0x17, 0x5b, 0x70, 0x49,
	0x51, 0x66, 0x5e, 0x3a, 0xb2, 0x1c, 0x27, 0x58, 0x4e, 0x4b, 0x91, 0x8b, 0xcb, 0x37, 0xb5, 0x24,
	0x23, 0x3f, 0x25, 0xa4, 0xb2, 0x20, 0x55, 0x88, 0x95, 0x8b, 0xd1, 0x51, 0x80, 0x01, 0x44, 0x05,
	0x0b, 0x30, 0x82, 0xa8, 0x00, 0x01, 0x26, 0x2b, 0x67, 0x2e, 0xee, 0xe2, 0xd4, 0xa2, 0xb2, 0xe4,
	0x54, 0xc7, 0xe2, 0xca, 0xbc, 0x64, 0x21, 0x79, 0x3d, 0x88, 0xa3, 0xf4, 0x60, 0x8e, 0xd2, 0x0b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xf5, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0xb8, 0x77,
	0x91, 0x49, 0x81, 0x51, 0x83, 0x23, 0x08, 0x59, 0x17, 0xc8, 0x90, 0xe4, 0x9c, 0xcc, 0xd4, 0xbc,
	0x12, 0x22, 0x0d, 0xb9, 0x0f, 0x33, 0x04, 0x49, 0x97, 0x55, 0x08, 0x17, 0x57, 0x2e, 0xc2, 0xb1,
	0x72, 0x18, 0x66, 0x40, 0x7c, 0x02, 0x33, 0x62, 0x87, 0x90, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0xb0,
	0x1e, 0x2c, 0x70, 0x11, 0x1e, 0x0d, 0x42, 0x32, 0xc7, 0x89, 0x33, 0x0a, 0x16, 0xde, 0x49, 0x6c,
	0x60, 0xa3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x9d, 0x09, 0x39, 0x90, 0x01, 0x00,
	0x00,
}
