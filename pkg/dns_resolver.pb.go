// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/dns_resolver.proto

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

// Resolvers used by this host.
type DnsResolver struct {
	// Address type index or order
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// Address type: sortlist
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// Resolver IP/IPv6 address
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address"`
	// Address (sortlist) netmask length
	Netmask string `protobuf:"bytes,4,opt,name=netmask,proto3" json:"netmask"`
	// Resolver options
	Options              int64    `protobuf:"varint,5,opt,name=options,proto3" json:"options"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnsResolver) Reset()         { *m = DnsResolver{} }
func (m *DnsResolver) String() string { return proto.CompactTextString(m) }
func (*DnsResolver) ProtoMessage()    {}
func (*DnsResolver) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ef22f3277aeb493, []int{0}
}

func (m *DnsResolver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsResolver.Unmarshal(m, b)
}
func (m *DnsResolver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsResolver.Marshal(b, m, deterministic)
}
func (m *DnsResolver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsResolver.Merge(m, src)
}
func (m *DnsResolver) XXX_Size() int {
	return xxx_messageInfo_DnsResolver.Size(m)
}
func (m *DnsResolver) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsResolver.DiscardUnknown(m)
}

var xxx_messageInfo_DnsResolver proto.InternalMessageInfo

func (m *DnsResolver) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DnsResolver) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DnsResolver) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DnsResolver) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

func (m *DnsResolver) GetOptions() int64 {
	if m != nil {
		return m.Options
	}
	return 0
}

func init() {
	proto.RegisterType((*DnsResolver)(nil), "schemas.DnsResolver")
}

func init() { proto.RegisterFile("pb/dns_resolver.proto", fileDescriptor_4ef22f3277aeb493) }

var fileDescriptor_4ef22f3277aeb493 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xbb, 0x4a, 0xc0, 0x30,
	0x18, 0x46, 0x49, 0x2f, 0x16, 0x23, 0x38, 0x04, 0x84, 0x8c, 0xc5, 0xa9, 0x8b, 0x8d, 0x28, 0xe8,
	0xe0, 0x26, 0x3e, 0x41, 0x46, 0x17, 0x69, 0x9a, 0xdf, 0x36, 0xd4, 0x5c, 0xcc, 0x9f, 0x0a, 0x9d,
	0x7c, 0x75, 0x69, 0x6c, 0xb7, 0xef, 0x7c, 0x67, 0x39, 0xf4, 0x26, 0x28, 0xa1, 0x1d, 0x7e, 0x44,
	0x40, 0xff, 0xf5, 0x03, 0xb1, 0x0f, 0xd1, 0x27, 0xcf, 0x1a, 0x1c, 0x67, 0xb0, 0x03, 0xde, 0xfe,
	0xd2, 0xab, 0x37, 0x87, 0xf2, 0xb0, 0xec, 0x9a, 0x16, 0x46, 0x73, 0xd2, 0x92, 0xae, 0x96, 0x85,
	0xd1, 0x8c, 0xd1, 0x2a, 0x6d, 0x01, 0x78, 0xd1, 0x92, 0xee, 0x52, 0xe6, 0xcd, 0x38, 0x6d, 0x06,
	0xad, 0x23, 0x20, 0xf2, 0x32, 0xdf, 0x27, 0xee, 0xc6, 0x41, 0xb2, 0x03, 0x2e, 0xbc, 0xfa, 0x37,
	0x07, 0xee, 0xc6, 0x87, 0x64, 0xbc, 0x43, 0x5e, 0xb7, 0xa4, 0x2b, 0xe5, 0x89, 0xaf, 0x0f, 0xef,
	0xf7, 0x93, 0x49, 0xf3, 0xaa, 0xfa, 0xd1, 0x5b, 0x61, 0xcd, 0xb8, 0x40, 0x78, 0x7e, 0x12, 0x1e,
	0xbf, 0x57, 0x88, 0xdb, 0x5d, 0xce, 0x55, 0xeb, 0xa7, 0x08, 0xcb, 0xf4, 0x72, 0x44, 0xab, 0x8b,
	0xfc, 0x3e, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x4c, 0xef, 0x18, 0xdd, 0x00, 0x00, 0x00,
}
