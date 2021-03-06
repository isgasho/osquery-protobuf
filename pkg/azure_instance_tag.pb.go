// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/azure_instance_tag.proto

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

// Azure instance tags.
type AzureInstanceTag struct {
	// Unique identifier for the VM
	VmId string `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id"`
	// The tag key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key"`
	// The tag value
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureInstanceTag) Reset()         { *m = AzureInstanceTag{} }
func (m *AzureInstanceTag) String() string { return proto.CompactTextString(m) }
func (*AzureInstanceTag) ProtoMessage()    {}
func (*AzureInstanceTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b7a6f14001a594b, []int{0}
}

func (m *AzureInstanceTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureInstanceTag.Unmarshal(m, b)
}
func (m *AzureInstanceTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureInstanceTag.Marshal(b, m, deterministic)
}
func (m *AzureInstanceTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureInstanceTag.Merge(m, src)
}
func (m *AzureInstanceTag) XXX_Size() int {
	return xxx_messageInfo_AzureInstanceTag.Size(m)
}
func (m *AzureInstanceTag) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureInstanceTag.DiscardUnknown(m)
}

var xxx_messageInfo_AzureInstanceTag proto.InternalMessageInfo

func (m *AzureInstanceTag) GetVmId() string {
	if m != nil {
		return m.VmId
	}
	return ""
}

func (m *AzureInstanceTag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AzureInstanceTag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*AzureInstanceTag)(nil), "schemas.AzureInstanceTag")
}

func init() { proto.RegisterFile("pb/azure_instance_tag.proto", fileDescriptor_8b7a6f14001a594b) }

var fileDescriptor_8b7a6f14001a594b = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x48, 0xd2, 0x4f,
	0xac, 0x2a, 0x2d, 0x4a, 0x8d, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x8d, 0x2f, 0x49,
	0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c,
	0x56, 0xf2, 0xe7, 0x12, 0x70, 0x04, 0x29, 0xf2, 0x84, 0xaa, 0x09, 0x49, 0x4c, 0x17, 0x12, 0xe6,
	0x62, 0x2d, 0xcb, 0x8d, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x29, 0xcb,
	0xf5, 0x4c, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x02, 0x0b, 0x81, 0x98, 0x42,
	0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0xcc, 0x60, 0x31, 0x08, 0xc7, 0xc9, 0x28,
	0xca, 0x20, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x37, 0x33, 0x39,
	0x3b, 0xb5, 0xc0, 0xdc, 0x4c, 0x3f, 0xbf, 0xb8, 0xb0, 0x34, 0xb5, 0xa8, 0x52, 0x17, 0x6c, 0x7d,
	0x52, 0x69, 0x9a, 0x7e, 0x41, 0x76, 0xba, 0x35, 0xd4, 0x11, 0x49, 0x6c, 0x60, 0x51, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x02, 0xee, 0x63, 0xb3, 0x00, 0x00, 0x00,
}
