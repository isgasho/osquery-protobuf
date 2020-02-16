// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/extended_attribute.proto

package schemas

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

// Returns the extended attributes for files (similar to Windows ADS).
type ExtendedAttribute struct {
	// Absolute file path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
	// Directory of file(s)
	Directory string `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory"`
	// Name of the value generated from the extended attribute
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key"`
	// The parsed information from the attribute
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value"`
	// 1 if the value is base64 encoded else 0
	Base64               int32    `protobuf:"varint,5,opt,name=base64,proto3" json:"base64"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtendedAttribute) Reset()         { *m = ExtendedAttribute{} }
func (m *ExtendedAttribute) String() string { return proto.CompactTextString(m) }
func (*ExtendedAttribute) ProtoMessage()    {}
func (*ExtendedAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d83a7339e298673, []int{0}
}

func (m *ExtendedAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendedAttribute.Unmarshal(m, b)
}
func (m *ExtendedAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendedAttribute.Marshal(b, m, deterministic)
}
func (m *ExtendedAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedAttribute.Merge(m, src)
}
func (m *ExtendedAttribute) XXX_Size() int {
	return xxx_messageInfo_ExtendedAttribute.Size(m)
}
func (m *ExtendedAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedAttribute proto.InternalMessageInfo

func (m *ExtendedAttribute) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ExtendedAttribute) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *ExtendedAttribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ExtendedAttribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ExtendedAttribute) GetBase64() int32 {
	if m != nil {
		return m.Base64
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtendedAttribute)(nil), "schemas.ExtendedAttribute")
}

func init() { proto.RegisterFile("pb/extended_attribute.proto", fileDescriptor_9d83a7339e298673) }

var fileDescriptor_9d83a7339e298673 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x48, 0xd2, 0x4f,
	0xad, 0x28, 0x49, 0xcd, 0x4b, 0x49, 0x4d, 0x89, 0x4f, 0x2c, 0x29, 0x29, 0xca, 0x4c, 0x2a, 0x2d,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c,
	0x56, 0x6a, 0x65, 0xe4, 0x12, 0x74, 0x85, 0xaa, 0x72, 0x84, 0x29, 0x12, 0x12, 0xe2, 0x62, 0x29,
	0x48, 0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x64, 0xb8, 0x38,
	0x53, 0x32, 0x8b, 0x52, 0x93, 0x4b, 0xf2, 0x8b, 0x2a, 0x25, 0x98, 0xc0, 0x12, 0x08, 0x01, 0x21,
	0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x66, 0xb0, 0x38, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a,
	0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x02, 0x16, 0x83, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x92, 0x12,
	0x8b, 0x53, 0xcd, 0x4c, 0x24, 0x58, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x27, 0xa3, 0x28,
	0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xdc, 0xcc, 0xe4, 0xec,
	0xd4, 0x02, 0x73, 0x33, 0xfd, 0xfc, 0xe2, 0xc2, 0xd2, 0xd4, 0xa2, 0x4a, 0x5d, 0xb0, 0xab, 0x93,
	0x4a, 0xd3, 0xf4, 0x0b, 0xb2, 0xd3, 0xad, 0xa1, 0x6e, 0x4f, 0x62, 0x03, 0x8b, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x10, 0x69, 0x91, 0x8f, 0xea, 0x00, 0x00, 0x00,
}