// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/kernel_extension.proto

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

// OS X's kernel extensions, both loaded and within the load search path.
type KernelExtension struct {
	// Extension load tag or index
	Idx int32 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx"`
	// Reference count
	Refs int32 `protobuf:"varint,2,opt,name=refs,proto3" json:"refs"`
	// Bytes of wired memory used by extension
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	// Extension label
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// Extension version
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version"`
	// Optional path to extension bundle
	Path                 string   `protobuf:"bytes,6,opt,name=path,proto3" json:"path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelExtension) Reset()         { *m = KernelExtension{} }
func (m *KernelExtension) String() string { return proto.CompactTextString(m) }
func (*KernelExtension) ProtoMessage()    {}
func (*KernelExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f5d7e8371d6829, []int{0}
}

func (m *KernelExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelExtension.Unmarshal(m, b)
}
func (m *KernelExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelExtension.Marshal(b, m, deterministic)
}
func (m *KernelExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelExtension.Merge(m, src)
}
func (m *KernelExtension) XXX_Size() int {
	return xxx_messageInfo_KernelExtension.Size(m)
}
func (m *KernelExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelExtension.DiscardUnknown(m)
}

var xxx_messageInfo_KernelExtension proto.InternalMessageInfo

func (m *KernelExtension) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *KernelExtension) GetRefs() int32 {
	if m != nil {
		return m.Refs
	}
	return 0
}

func (m *KernelExtension) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *KernelExtension) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelExtension) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *KernelExtension) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*KernelExtension)(nil), "schemas.KernelExtension")
}

func init() { proto.RegisterFile("pb/kernel_extension.proto", fileDescriptor_50f5d7e8371d6829) }

var fileDescriptor_50f5d7e8371d6829 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xbb, 0x4e, 0xc4, 0x30,
	0x10, 0x00, 0x65, 0xf2, 0x12, 0x6e, 0x40, 0xae, 0x4c, 0x17, 0x51, 0xa5, 0x21, 0x46, 0x20, 0x41,
	0x41, 0x87, 0x44, 0x45, 0x97, 0x92, 0x06, 0xc5, 0x61, 0x93, 0x58, 0xc1, 0x0f, 0x6c, 0x07, 0x05,
	0xfe, 0xe1, 0xfe, 0xf9, 0xe4, 0xbd, 0x4b, 0x37, 0x3b, 0x1a, 0xad, 0x76, 0xe9, 0x8d, 0x93, 0x62,
	0x01, 0x6f, 0xe0, 0xfb, 0x13, 0xb6, 0x08, 0x26, 0x28, 0x6b, 0x5a, 0xe7, 0x6d, 0xb4, 0xac, 0x0a,
	0xc3, 0x0c, 0xba, 0x0f, 0xb7, 0x07, 0x42, 0xaf, 0xde, 0xb1, 0x79, 0xdb, 0x13, 0x76, 0x4d, 0x33,
	0xf5, 0xb5, 0x71, 0x52, 0x93, 0xa6, 0xe8, 0x12, 0x32, 0x46, 0x73, 0x0f, 0x63, 0xe0, 0x17, 0xa8,
	0x90, 0x93, 0x0b, 0xea, 0x1f, 0x78, 0x56, 0x93, 0x26, 0xeb, 0x90, 0x93, 0x33, 0xbd, 0x06, 0x9e,
	0xd7, 0xa4, 0xb9, 0xec, 0x90, 0x19, 0xa7, 0xd5, 0x2f, 0xf8, 0xb4, 0x98, 0x17, 0xa8, 0xf7, 0x31,
	0xd5, 0xae, 0x8f, 0x33, 0x2f, 0x4f, 0x75, 0xe2, 0xd7, 0x87, 0x8f, 0xfb, 0x49, 0xc5, 0x79, 0x95,
	0xed, 0x60, 0xb5, 0xd0, 0x6a, 0x58, 0xc0, 0x3d, 0x3f, 0x09, 0x1b, 0x7e, 0x56, 0xf0, 0x7f, 0x77,
	0x78, 0xbd, 0x5c, 0x47, 0xe1, 0x96, 0xe9, 0xe5, 0xfc, 0x83, 0x2c, 0xd1, 0x3e, 0x1e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xff, 0x47, 0x9c, 0x56, 0xf0, 0x00, 0x00, 0x00,
}