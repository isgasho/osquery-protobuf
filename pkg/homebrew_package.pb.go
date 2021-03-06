// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/homebrew_package.proto

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

// The installed homebrew package database.
type HomebrewPackage struct {
	// Package name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Package install path
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path"`
	// Current 'linked' version
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HomebrewPackage) Reset()         { *m = HomebrewPackage{} }
func (m *HomebrewPackage) String() string { return proto.CompactTextString(m) }
func (*HomebrewPackage) ProtoMessage()    {}
func (*HomebrewPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6694af009c825f3, []int{0}
}

func (m *HomebrewPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HomebrewPackage.Unmarshal(m, b)
}
func (m *HomebrewPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HomebrewPackage.Marshal(b, m, deterministic)
}
func (m *HomebrewPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HomebrewPackage.Merge(m, src)
}
func (m *HomebrewPackage) XXX_Size() int {
	return xxx_messageInfo_HomebrewPackage.Size(m)
}
func (m *HomebrewPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_HomebrewPackage.DiscardUnknown(m)
}

var xxx_messageInfo_HomebrewPackage proto.InternalMessageInfo

func (m *HomebrewPackage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HomebrewPackage) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HomebrewPackage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*HomebrewPackage)(nil), "schemas.HomebrewPackage")
}

func init() { proto.RegisterFile("pb/homebrew_package.proto", fileDescriptor_f6694af009c825f3) }

var fileDescriptor_f6694af009c825f3 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x48, 0xd2, 0xcf,
	0xc8, 0xcf, 0x4d, 0x4d, 0x2a, 0x4a, 0x2d, 0x8f, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0x0a,
	0xe6, 0xe2, 0xf7, 0x80, 0x2a, 0x09, 0x80, 0xa8, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x41, 0x62, 0x05, 0x89, 0x25, 0x19, 0x12,
	0x4c, 0x10, 0x31, 0x10, 0x5b, 0x48, 0x82, 0x8b, 0xbd, 0x2c, 0xb5, 0xa8, 0x38, 0x33, 0x3f, 0x4f,
	0x82, 0x19, 0x2c, 0x0c, 0xe3, 0x3a, 0x19, 0x45, 0x19, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0xe7, 0x66, 0x26, 0x67, 0xa7, 0x16, 0x98, 0x9b, 0xe9, 0xe7, 0x17, 0x17,
	0x96, 0xa6, 0x16, 0x55, 0xea, 0x82, 0x9d, 0x90, 0x54, 0x9a, 0xa6, 0x5f, 0x90, 0x9d, 0x6e, 0x0d,
	0x75, 0x48, 0x12, 0x1b, 0x58, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x9d, 0x9b, 0x4d,
	0xb5, 0x00, 0x00, 0x00,
}
