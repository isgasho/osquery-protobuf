// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/browser_plugin.proto

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

// All C/NPAPI browser plugin details for all users.
type BrowserPlugin struct {
	// Plugin display name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Plugin identifier
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier"`
	// Plugin short version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version"`
	// Build SDK used to compile plugin
	Sdk string `protobuf:"bytes,4,opt,name=sdk,proto3" json:"sdk"`
	// Plugin description text
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description"`
	// Plugin language-localization
	DevelopmentRegion string `protobuf:"bytes,6,opt,name=development_region,json=developmentRegion,proto3" json:"development_region"`
	// Plugin requires native execution
	Native int32 `protobuf:"varint,7,opt,name=native,proto3" json:"native"`
	// Path to plugin bundle
	Path string `protobuf:"bytes,8,opt,name=path,proto3" json:"path"`
	// Is the plugin disabled. 1 = Disabled
	Disabled             int32    `protobuf:"varint,9,opt,name=disabled,proto3" json:"disabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrowserPlugin) Reset()         { *m = BrowserPlugin{} }
func (m *BrowserPlugin) String() string { return proto.CompactTextString(m) }
func (*BrowserPlugin) ProtoMessage()    {}
func (*BrowserPlugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_aecfbd426550a47b, []int{0}
}

func (m *BrowserPlugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrowserPlugin.Unmarshal(m, b)
}
func (m *BrowserPlugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrowserPlugin.Marshal(b, m, deterministic)
}
func (m *BrowserPlugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrowserPlugin.Merge(m, src)
}
func (m *BrowserPlugin) XXX_Size() int {
	return xxx_messageInfo_BrowserPlugin.Size(m)
}
func (m *BrowserPlugin) XXX_DiscardUnknown() {
	xxx_messageInfo_BrowserPlugin.DiscardUnknown(m)
}

var xxx_messageInfo_BrowserPlugin proto.InternalMessageInfo

func (m *BrowserPlugin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BrowserPlugin) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *BrowserPlugin) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *BrowserPlugin) GetSdk() string {
	if m != nil {
		return m.Sdk
	}
	return ""
}

func (m *BrowserPlugin) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BrowserPlugin) GetDevelopmentRegion() string {
	if m != nil {
		return m.DevelopmentRegion
	}
	return ""
}

func (m *BrowserPlugin) GetNative() int32 {
	if m != nil {
		return m.Native
	}
	return 0
}

func (m *BrowserPlugin) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *BrowserPlugin) GetDisabled() int32 {
	if m != nil {
		return m.Disabled
	}
	return 0
}

func init() {
	proto.RegisterType((*BrowserPlugin)(nil), "schemas.BrowserPlugin")
}

func init() { proto.RegisterFile("pb/browser_plugin.proto", fileDescriptor_aecfbd426550a47b) }

var fileDescriptor_aecfbd426550a47b = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0xfe, 0x24, 0xad, 0x11, 0x12, 0xdc, 0x01, 0x2c, 0x06, 0x14, 0x31, 0x75, 0x69,
	0x83, 0x40, 0x82, 0x81, 0xad, 0x4f, 0x80, 0x32, 0xb2, 0x54, 0x71, 0x7c, 0x9b, 0x58, 0x49, 0x6c,
	0x63, 0x3b, 0x41, 0xbc, 0x01, 0x8f, 0x8d, 0x72, 0x5b, 0x50, 0xb6, 0x73, 0xce, 0xf7, 0x2d, 0xf7,
	0xb2, 0x5b, 0x2b, 0x32, 0xe1, 0xcc, 0x97, 0x47, 0x77, 0xb0, 0x6d, 0x5f, 0x29, 0xbd, 0xb3, 0xce,
	0x04, 0x03, 0x89, 0x2f, 0x6b, 0xec, 0x0a, 0xff, 0xf0, 0x33, 0x63, 0x97, 0xfb, 0x93, 0xf1, 0x4e,
	0x02, 0x00, 0x5b, 0xe8, 0xa2, 0x43, 0x1e, 0xa5, 0xd1, 0x66, 0x9d, 0x53, 0x86, 0x7b, 0xc6, 0x94,
	0x44, 0x1d, 0xd4, 0x51, 0xa1, 0xe3, 0x33, 0x22, 0x93, 0x05, 0x38, 0x4b, 0x06, 0x74, 0x5e, 0x19,
	0xcd, 0xe7, 0x04, 0xff, 0x2a, 0x5c, 0xb1, 0xb9, 0x97, 0x0d, 0x5f, 0xd0, 0x3a, 0x46, 0x48, 0xd9,
	0x85, 0x44, 0x5f, 0x3a, 0x65, 0xc3, 0xe8, 0x2f, 0x89, 0x4c, 0x27, 0xd8, 0x32, 0x90, 0x38, 0x60,
	0x6b, 0x6c, 0x87, 0x3a, 0x1c, 0x1c, 0x56, 0xa3, 0x18, 0x93, 0x78, 0x3d, 0x21, 0x39, 0x01, 0xb8,
	0x61, 0xb1, 0x2e, 0x82, 0x1a, 0x90, 0x27, 0x69, 0xb4, 0x59, 0xe6, 0xe7, 0x36, 0x1e, 0x62, 0x8b,
	0x50, 0xf3, 0xd5, 0xe9, 0x90, 0x31, 0xc3, 0x1d, 0x5b, 0x49, 0xe5, 0x0b, 0xd1, 0xa2, 0xe4, 0x6b,
	0xb2, 0xff, 0xfb, 0xfe, 0xe9, 0xe3, 0xb1, 0x52, 0xa1, 0xee, 0xc5, 0xae, 0x34, 0x5d, 0xd6, 0xa9,
	0xb2, 0x41, 0xfb, 0xfa, 0x92, 0x19, 0xff, 0xd9, 0xa3, 0xfb, 0xde, 0xd2, 0xe3, 0x44, 0x7f, 0xcc,
	0x6c, 0x53, 0xbd, 0x9d, 0xdf, 0x27, 0x62, 0x5a, 0x9f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0e,
	0x50, 0x80, 0xfe, 0x69, 0x01, 0x00, 0x00,
}
