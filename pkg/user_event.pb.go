// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user_event.proto

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

// Track user events from the audit framework.
type UserEvent struct {
	// User ID
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// Audit User ID
	Auid int64 `protobuf:"varint,2,opt,name=auid,proto3" json:"auid"`
	// Process (or thread) ID
	Pid int64 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid"`
	// Message from the event
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message"`
	// The file description for the process socket
	Type int32 `protobuf:"varint,5,opt,name=type,proto3" json:"type"`
	// Supplied path from event
	Path string `protobuf:"bytes,6,opt,name=path,proto3" json:"path"`
	// The Internet protocol address or family ID
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address"`
	// The network protocol ID
	Terminal string `protobuf:"bytes,8,opt,name=terminal,proto3" json:"terminal"`
	// Time of execution in UNIX time
	Time int64 `protobuf:"varint,9,opt,name=time,proto3" json:"time"`
	// Time of execution in system uptime
	Uptime int64 `protobuf:"varint,10,opt,name=uptime,proto3" json:"uptime"`
	// Event ID
	Eid                  string   `protobuf:"bytes,11,opt,name=eid,proto3" json:"eid"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEvent) Reset()         { *m = UserEvent{} }
func (m *UserEvent) String() string { return proto.CompactTextString(m) }
func (*UserEvent) ProtoMessage()    {}
func (*UserEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_751c3bd39377d6b2, []int{0}
}

func (m *UserEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEvent.Unmarshal(m, b)
}
func (m *UserEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEvent.Marshal(b, m, deterministic)
}
func (m *UserEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEvent.Merge(m, src)
}
func (m *UserEvent) XXX_Size() int {
	return xxx_messageInfo_UserEvent.Size(m)
}
func (m *UserEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEvent.DiscardUnknown(m)
}

var xxx_messageInfo_UserEvent proto.InternalMessageInfo

func (m *UserEvent) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserEvent) GetAuid() int64 {
	if m != nil {
		return m.Auid
	}
	return 0
}

func (m *UserEvent) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *UserEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserEvent) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *UserEvent) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UserEvent) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UserEvent) GetTerminal() string {
	if m != nil {
		return m.Terminal
	}
	return ""
}

func (m *UserEvent) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *UserEvent) GetUptime() int64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *UserEvent) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func init() {
	proto.RegisterType((*UserEvent)(nil), "schemas.UserEvent")
}

func init() { proto.RegisterFile("pb/user_event.proto", fileDescriptor_751c3bd39377d6b2) }

var fileDescriptor_751c3bd39377d6b2 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x3f, 0x4b, 0xf4, 0x40,
	0x10, 0xc6, 0xc9, 0x9b, 0xbb, 0xe4, 0xb2, 0x6f, 0x23, 0x2b, 0xc8, 0x60, 0x15, 0xac, 0xd2, 0x78,
	0x11, 0x05, 0x2d, 0xec, 0x04, 0xbf, 0x40, 0xc0, 0xc6, 0x46, 0x36, 0xd9, 0x31, 0x59, 0xce, 0xbd,
	0xac, 0xfb, 0x47, 0xb8, 0x8f, 0x6e, 0x27, 0x33, 0x89, 0x76, 0xbf, 0xe7, 0xb7, 0xb3, 0x0f, 0xc3,
	0x88, 0x73, 0xd7, 0xb7, 0x29, 0xa0, 0x7f, 0xc3, 0x2f, 0x3c, 0xc6, 0xbd, 0xf3, 0x73, 0x9c, 0x65,
	0x19, 0x86, 0x09, 0xad, 0x0a, 0x57, 0xdf, 0x99, 0xa8, 0x5e, 0x02, 0xfa, 0x67, 0x7a, 0x94, 0x67,
	0x22, 0x4f, 0x46, 0x43, 0x56, 0x67, 0x4d, 0xde, 0x11, 0x4a, 0x29, 0x36, 0x8a, 0xd4, 0x3f, 0x56,
	0xcc, 0x34, 0xe5, 0x8c, 0x86, 0x7c, 0x99, 0x72, 0x46, 0x4b, 0x10, 0xa5, 0xc5, 0x10, 0xd4, 0x88,
	0xb0, 0xa9, 0xb3, 0xa6, 0xea, 0x7e, 0x23, 0xfd, 0x8f, 0x27, 0x87, 0xb0, 0xad, 0xb3, 0x66, 0xdb,
	0x31, 0x93, 0x73, 0x2a, 0x4e, 0x50, 0xf0, 0x28, 0x33, 0x35, 0x28, 0xad, 0x3d, 0x86, 0x00, 0xe5,
	0xd2, 0xb0, 0x46, 0x79, 0x29, 0x76, 0x11, 0xbd, 0x35, 0x47, 0xf5, 0x01, 0x3b, 0x7e, 0xfa, 0xcb,
	0xdc, 0x6e, 0x2c, 0x42, 0xb5, 0x6c, 0x47, 0x2c, 0x2f, 0x44, 0x91, 0x1c, 0x5b, 0xc1, 0x76, 0x4d,
	0xb4, 0x35, 0x1a, 0x0d, 0xff, 0xb9, 0x82, 0xf0, 0xe9, 0xf6, 0xf5, 0x66, 0x34, 0x71, 0x4a, 0xfd,
	0x7e, 0x98, 0x6d, 0x6b, 0xcd, 0x70, 0x40, 0xf7, 0x70, 0xdf, 0xce, 0xe1, 0x33, 0xa1, 0x3f, 0x5d,
	0xf3, 0xa5, 0xfa, 0xf4, 0xde, 0xba, 0xc3, 0xf8, 0xb8, 0xde, 0xab, 0x2f, 0xd8, 0xde, 0xfd, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x33, 0x34, 0x30, 0xf9, 0x56, 0x01, 0x00, 0x00,
}
