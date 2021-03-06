// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/apt_source.proto

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

// Current list of APT repositories or software channels.
type AptSource struct {
	// Repository name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Source file
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
	// Repository base URI
	BaseUri string `protobuf:"bytes,3,opt,name=base_uri,json=baseUri,proto3" json:"base_uri"`
	// Release name
	Release string `protobuf:"bytes,4,opt,name=release,proto3" json:"release"`
	// Repository source version
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version"`
	// Repository maintainer
	Maintainer string `protobuf:"bytes,6,opt,name=maintainer,proto3" json:"maintainer"`
	// Repository components
	Components string `protobuf:"bytes,7,opt,name=components,proto3" json:"components"`
	// Repository architectures
	Architectures        string   `protobuf:"bytes,8,opt,name=architectures,proto3" json:"architectures"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AptSource) Reset()         { *m = AptSource{} }
func (m *AptSource) String() string { return proto.CompactTextString(m) }
func (*AptSource) ProtoMessage()    {}
func (*AptSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b730ca557630dd, []int{0}
}

func (m *AptSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AptSource.Unmarshal(m, b)
}
func (m *AptSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AptSource.Marshal(b, m, deterministic)
}
func (m *AptSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AptSource.Merge(m, src)
}
func (m *AptSource) XXX_Size() int {
	return xxx_messageInfo_AptSource.Size(m)
}
func (m *AptSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AptSource.DiscardUnknown(m)
}

var xxx_messageInfo_AptSource proto.InternalMessageInfo

func (m *AptSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AptSource) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *AptSource) GetBaseUri() string {
	if m != nil {
		return m.BaseUri
	}
	return ""
}

func (m *AptSource) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *AptSource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AptSource) GetMaintainer() string {
	if m != nil {
		return m.Maintainer
	}
	return ""
}

func (m *AptSource) GetComponents() string {
	if m != nil {
		return m.Components
	}
	return ""
}

func (m *AptSource) GetArchitectures() string {
	if m != nil {
		return m.Architectures
	}
	return ""
}

func init() {
	proto.RegisterType((*AptSource)(nil), "schemas.AptSource")
}

func init() { proto.RegisterFile("pb/apt_source.proto", fileDescriptor_60b730ca557630dd) }

var fileDescriptor_60b730ca557630dd = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3f, 0x4b, 0x34, 0x31,
	0x10, 0x87, 0xd9, 0xf7, 0x3d, 0x77, 0xef, 0x06, 0x6c, 0x22, 0xc8, 0xd8, 0x88, 0x88, 0x85, 0x8d,
	0xb7, 0xa2, 0xa0, 0x85, 0x95, 0x7e, 0x04, 0xc5, 0xc6, 0xe6, 0x48, 0xc2, 0x78, 0x1b, 0xce, 0xfc,
	0x71, 0x26, 0x11, 0xfc, 0xce, 0x7e, 0x08, 0xd9, 0xec, 0x82, 0xda, 0xe5, 0x79, 0x9e, 0x4c, 0xf3,
	0x83, 0x83, 0x64, 0x7a, 0x9d, 0xf2, 0x46, 0x62, 0x61, 0x4b, 0xeb, 0xc4, 0x31, 0x47, 0xd5, 0x89,
	0x1d, 0xc8, 0x6b, 0x39, 0xfd, 0x6a, 0x60, 0x75, 0x9f, 0xf2, 0x53, 0x8d, 0x4a, 0xc1, 0x22, 0x68,
	0x4f, 0xd8, 0x9c, 0x34, 0xe7, 0xab, 0xc7, 0xfa, 0x56, 0x87, 0xd0, 0x4e, 0xa7, 0xf8, 0xaf, 0xda,
	0x99, 0xd4, 0x11, 0x2c, 0x8d, 0x16, 0xda, 0x14, 0x76, 0xf8, 0xbf, 0x96, 0x6e, 0xe4, 0x67, 0x76,
	0x0a, 0xa1, 0x63, 0x7a, 0x23, 0x2d, 0x84, 0x8b, 0xa9, 0xcc, 0x38, 0x96, 0x0f, 0x62, 0x71, 0x31,
	0xe0, 0xde, 0x54, 0x66, 0x54, 0xc7, 0x00, 0x5e, 0xbb, 0x90, 0xb5, 0x0b, 0xc4, 0xd8, 0xd6, 0xf8,
	0xcb, 0x8c, 0xdd, 0x46, 0x9f, 0x62, 0xa0, 0x90, 0x05, 0xbb, 0xa9, 0xff, 0x18, 0x75, 0x06, 0xfb,
	0x9a, 0xed, 0xe0, 0x32, 0xd9, 0x5c, 0x98, 0x04, 0x97, 0xf5, 0xcb, 0x5f, 0xf9, 0x70, 0xf5, 0x72,
	0xb9, 0x75, 0x79, 0x28, 0x66, 0x6d, 0xa3, 0xef, 0xbd, 0xb3, 0x3b, 0x4a, 0xb7, 0x37, 0x7d, 0x94,
	0xf7, 0x42, 0xfc, 0x79, 0x51, 0xc7, 0x31, 0xe5, 0xb5, 0x4f, 0xbb, 0xed, 0xdd, 0x3c, 0x91, 0x69,
	0xab, 0xbd, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x60, 0x44, 0x35, 0x6a, 0x49, 0x01, 0x00, 0x00,
}
