// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/selinux_event.proto

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

// Track SELinux events.
type SelinuxEvent struct {
	// Event type
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	// Message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	// Time of execution in UNIX time
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time"`
	// Time of execution in system uptime
	Uptime int64 `protobuf:"varint,4,opt,name=uptime,proto3" json:"uptime"`
	// Event ID
	Eid                  string   `protobuf:"bytes,5,opt,name=eid,proto3" json:"eid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelinuxEvent) Reset()         { *m = SelinuxEvent{} }
func (m *SelinuxEvent) String() string { return proto.CompactTextString(m) }
func (*SelinuxEvent) ProtoMessage()    {}
func (*SelinuxEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8aa3badd287a2091, []int{0}
}

func (m *SelinuxEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelinuxEvent.Unmarshal(m, b)
}
func (m *SelinuxEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelinuxEvent.Marshal(b, m, deterministic)
}
func (m *SelinuxEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelinuxEvent.Merge(m, src)
}
func (m *SelinuxEvent) XXX_Size() int {
	return xxx_messageInfo_SelinuxEvent.Size(m)
}
func (m *SelinuxEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SelinuxEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SelinuxEvent proto.InternalMessageInfo

func (m *SelinuxEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SelinuxEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SelinuxEvent) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SelinuxEvent) GetUptime() int64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *SelinuxEvent) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func init() {
	proto.RegisterType((*SelinuxEvent)(nil), "schemas.SelinuxEvent")
}

func init() { proto.RegisterFile("pb/selinux_event.proto", fileDescriptor_8aa3badd287a2091) }

var fileDescriptor_8aa3badd287a2091 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x2f,
	0x4e, 0xcd, 0xc9, 0xcc, 0x2b, 0xad, 0x88, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0x56, 0xaa, 0xe2, 0xe2, 0x09,
	0x86, 0xc8, 0xbb, 0x82, 0xa4, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4,
	0x54, 0x09, 0x26, 0xb0, 0x30, 0x8c, 0x0b, 0x56, 0x9d, 0x99, 0x9b, 0x2a, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x1c, 0x04, 0x66, 0x0b, 0x89, 0x71, 0xb1, 0x95, 0x16, 0x80, 0x45, 0x59, 0xc0, 0xa2, 0x50,
	0x9e, 0x90, 0x00, 0x17, 0x73, 0x6a, 0x66, 0x8a, 0x04, 0x2b, 0xd8, 0x04, 0x10, 0xd3, 0xc9, 0x28,
	0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x37, 0x33, 0x39,
	0x3b, 0xb5, 0xc0, 0xdc, 0x4c, 0x3f, 0xbf, 0xb8, 0xb0, 0x34, 0xb5, 0xa8, 0x52, 0x17, 0xec, 0xd2,
	0xa4, 0xd2, 0x34, 0xfd, 0x82, 0xec, 0x74, 0x6b, 0xa8, 0x7b, 0x93, 0xd8, 0xc0, 0xa2, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x3f, 0x94, 0x60, 0xd9, 0x00, 0x00, 0x00,
}