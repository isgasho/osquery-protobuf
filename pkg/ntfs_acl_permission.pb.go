// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/ntfs_acl_permission.proto

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

// Retrieve NTFS ACL permission information for files and directories.
type NtfsAclPermission struct {
	// Path to the file or directory.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path"`
	// Type of access mode for the access control entry.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	// User or group to which the ACE applies.
	Principal string `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal"`
	// Specific permissions that indicate the rights described by the ACE.
	Access string `protobuf:"bytes,4,opt,name=access,proto3" json:"access"`
	// The inheritance policy of the ACE.
	InheritedFrom        string   `protobuf:"bytes,5,opt,name=inherited_from,json=inheritedFrom,proto3" json:"inherited_from"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NtfsAclPermission) Reset()         { *m = NtfsAclPermission{} }
func (m *NtfsAclPermission) String() string { return proto.CompactTextString(m) }
func (*NtfsAclPermission) ProtoMessage()    {}
func (*NtfsAclPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_59de2852e4a6b0c2, []int{0}
}

func (m *NtfsAclPermission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NtfsAclPermission.Unmarshal(m, b)
}
func (m *NtfsAclPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NtfsAclPermission.Marshal(b, m, deterministic)
}
func (m *NtfsAclPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NtfsAclPermission.Merge(m, src)
}
func (m *NtfsAclPermission) XXX_Size() int {
	return xxx_messageInfo_NtfsAclPermission.Size(m)
}
func (m *NtfsAclPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_NtfsAclPermission.DiscardUnknown(m)
}

var xxx_messageInfo_NtfsAclPermission proto.InternalMessageInfo

func (m *NtfsAclPermission) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *NtfsAclPermission) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NtfsAclPermission) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *NtfsAclPermission) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *NtfsAclPermission) GetInheritedFrom() string {
	if m != nil {
		return m.InheritedFrom
	}
	return ""
}

func init() {
	proto.RegisterType((*NtfsAclPermission)(nil), "schemas.NtfsAclPermission")
}

func init() { proto.RegisterFile("pb/ntfs_acl_permission.proto", fileDescriptor_59de2852e4a6b0c2) }

var fileDescriptor_59de2852e4a6b0c2 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x05, 0x60, 0x56, 0xcf, 0x93, 0x0b, 0x28, 0x98, 0x42, 0x52, 0x5c, 0x21, 0x82, 0x60, 0xe3,
	0xad, 0x28, 0x68, 0x61, 0xa5, 0x85, 0xa5, 0x88, 0xa5, 0xcd, 0x92, 0x1d, 0x27, 0x97, 0xe1, 0x36,
	0xc9, 0x98, 0xc9, 0x16, 0xf7, 0x4f, 0xfc, 0xb9, 0x72, 0xf1, 0x58, 0xbb, 0x37, 0xdf, 0x9b, 0xe6,
	0xa9, 0x25, 0xf7, 0x6d, 0x2c, 0x4e, 0x3a, 0x0b, 0x43, 0xc7, 0x98, 0x03, 0x89, 0x50, 0x8a, 0x2b,
	0xce, 0xa9, 0x24, 0x7d, 0x2c, 0xe0, 0x31, 0x58, 0xb9, 0xfc, 0x69, 0xd4, 0xd9, 0x5b, 0x71, 0xf2,
	0x0c, 0xc3, 0xfb, 0xf4, 0xa4, 0xb5, 0x9a, 0xb1, 0x2d, 0xde, 0x34, 0x17, 0xcd, 0xf5, 0xe2, 0xa3,
	0xe6, 0x9d, 0x95, 0x2d, 0xa3, 0x39, 0xf8, 0xb3, 0x5d, 0xd6, 0x4b, 0xb5, 0xe0, 0x4c, 0x11, 0x88,
	0xed, 0x60, 0x0e, 0x6b, 0xf1, 0x0f, 0xfa, 0x5c, 0xcd, 0x2d, 0x00, 0x8a, 0x98, 0x59, 0xad, 0xf6,
	0x97, 0xbe, 0x52, 0xa7, 0x14, 0x3d, 0x66, 0x2a, 0xf8, 0xd5, 0xb9, 0x9c, 0x82, 0x39, 0xaa, 0xfd,
	0xc9, 0xa4, 0xaf, 0x39, 0x85, 0x97, 0xbb, 0xcf, 0xdb, 0x35, 0x15, 0x3f, 0xf6, 0x2b, 0x48, 0xa1,
	0x0d, 0x04, 0x1b, 0xe4, 0xc7, 0x87, 0x36, 0xc9, 0xf7, 0x88, 0x79, 0x7b, 0x53, 0x87, 0xf4, 0xa3,
	0x6b, 0x79, 0xb3, 0x7e, 0xda, 0xcf, 0xe9, 0xe7, 0x55, 0xef, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0xe7, 0xed, 0xb1, 0xfe, 0x00, 0x00, 0x00,
}
