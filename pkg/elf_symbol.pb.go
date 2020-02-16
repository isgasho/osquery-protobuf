// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/elf_symbol.proto

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

// ELF symbol list.
type ElfSymbol struct {
	// Symbol name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Symbol address (value)
	Addr int32 `protobuf:"varint,2,opt,name=addr,proto3" json:"addr"`
	// Size of object
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size"`
	// Symbol type
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	// Binding type
	Binding string `protobuf:"bytes,5,opt,name=binding,proto3" json:"binding"`
	// Section table index
	Offset int32 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset"`
	// Table name containing symbol
	Table string `protobuf:"bytes,7,opt,name=table,proto3" json:"table"`
	// Path to ELF file
	Path                 string   `protobuf:"bytes,8,opt,name=path,proto3" json:"path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElfSymbol) Reset()         { *m = ElfSymbol{} }
func (m *ElfSymbol) String() string { return proto.CompactTextString(m) }
func (*ElfSymbol) ProtoMessage()    {}
func (*ElfSymbol) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed9afaa5dc02982b, []int{0}
}

func (m *ElfSymbol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElfSymbol.Unmarshal(m, b)
}
func (m *ElfSymbol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElfSymbol.Marshal(b, m, deterministic)
}
func (m *ElfSymbol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElfSymbol.Merge(m, src)
}
func (m *ElfSymbol) XXX_Size() int {
	return xxx_messageInfo_ElfSymbol.Size(m)
}
func (m *ElfSymbol) XXX_DiscardUnknown() {
	xxx_messageInfo_ElfSymbol.DiscardUnknown(m)
}

var xxx_messageInfo_ElfSymbol proto.InternalMessageInfo

func (m *ElfSymbol) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ElfSymbol) GetAddr() int32 {
	if m != nil {
		return m.Addr
	}
	return 0
}

func (m *ElfSymbol) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ElfSymbol) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ElfSymbol) GetBinding() string {
	if m != nil {
		return m.Binding
	}
	return ""
}

func (m *ElfSymbol) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ElfSymbol) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ElfSymbol) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*ElfSymbol)(nil), "schemas.ElfSymbol")
}

func init() { proto.RegisterFile("pb/elf_symbol.proto", fileDescriptor_ed9afaa5dc02982b) }

var fileDescriptor_ed9afaa5dc02982b = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x15, 0x68, 0x12, 0xea, 0xd1, 0x20, 0xf4, 0xc6, 0x8a, 0xa9, 0x0b, 0x35, 0x02, 0x09,
	0x06, 0x36, 0x24, 0x7e, 0xa0, 0x6c, 0x2c, 0xc8, 0x4e, 0x9e, 0x13, 0xab, 0x76, 0x6c, 0x62, 0x67,
	0x08, 0x3f, 0xc6, 0xef, 0x21, 0x3f, 0x97, 0xed, 0xdc, 0x63, 0xbf, 0x2b, 0x5d, 0x76, 0x1d, 0x94,
	0x40, 0xab, 0xbf, 0xe2, 0xea, 0x94, 0xb7, 0x87, 0x30, 0xfb, 0xe4, 0x79, 0x1b, 0xbb, 0x11, 0x9d,
	0x8c, 0x77, 0xbf, 0x15, 0xdb, 0xbe, 0x5b, 0xfd, 0x41, 0x8f, 0x9c, 0xb3, 0xcd, 0x24, 0x1d, 0x42,
	0xb5, 0xab, 0xf6, 0xdb, 0x23, 0x71, 0x76, 0xb2, 0xef, 0x67, 0xb8, 0xd8, 0x55, 0xfb, 0xfa, 0x48,
	0x9c, 0x5d, 0x34, 0x3f, 0x08, 0x97, 0xc5, 0x65, 0xce, 0x2e, 0xad, 0x01, 0x61, 0x53, 0x6e, 0x33,
	0x73, 0x60, 0xad, 0x32, 0x53, 0x6f, 0xa6, 0x01, 0x6a, 0xd2, 0xff, 0x91, 0xdf, 0xb2, 0xc6, 0x6b,
	0x1d, 0x31, 0x41, 0x43, 0x1d, 0xe7, 0xc4, 0x6f, 0x58, 0x9d, 0xa4, 0xb2, 0x08, 0x2d, 0xfd, 0x2f,
	0x21, 0x77, 0x07, 0x99, 0x46, 0xb8, 0x2a, 0xdd, 0x99, 0xdf, 0x1e, 0x3f, 0x1f, 0x06, 0x93, 0xc6,
	0x45, 0x1d, 0x3a, 0xef, 0x84, 0x33, 0xdd, 0x09, 0xc3, 0xcb, 0xb3, 0xf0, 0xf1, 0x7b, 0xc1, 0x79,
	0xbd, 0xa7, 0x9d, 0x6a, 0xd1, 0x22, 0x9c, 0x86, 0xd7, 0xf3, 0x5a, 0xd5, 0x90, 0x7d, 0xfa, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x10, 0xce, 0xa3, 0x8c, 0x14, 0x01, 0x00, 0x00,
}