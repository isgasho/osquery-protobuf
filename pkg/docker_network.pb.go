// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/docker_network.proto

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

// Docker networks information.
type DockerNetwork struct {
	// Network ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// Network name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// Network driver
	Driver string `protobuf:"bytes,3,opt,name=driver,proto3" json:"driver"`
	// Time of creation as UNIX time
	Created int64 `protobuf:"varint,4,opt,name=created,proto3" json:"created"`
	// 1 if IPv6 is enabled on this network. 0 otherwise
	EnableIpv6 int32 `protobuf:"varint,5,opt,name=enable_ipv6,json=enableIpv6,proto3" json:"enable_ipv6"`
	// Network subnet
	Subnet string `protobuf:"bytes,6,opt,name=subnet,proto3" json:"subnet"`
	// Network gateway
	Gateway              string   `protobuf:"bytes,7,opt,name=gateway,proto3" json:"gateway"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerNetwork) Reset()         { *m = DockerNetwork{} }
func (m *DockerNetwork) String() string { return proto.CompactTextString(m) }
func (*DockerNetwork) ProtoMessage()    {}
func (*DockerNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1f4548f986cabd, []int{0}
}

func (m *DockerNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerNetwork.Unmarshal(m, b)
}
func (m *DockerNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerNetwork.Marshal(b, m, deterministic)
}
func (m *DockerNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerNetwork.Merge(m, src)
}
func (m *DockerNetwork) XXX_Size() int {
	return xxx_messageInfo_DockerNetwork.Size(m)
}
func (m *DockerNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_DockerNetwork proto.InternalMessageInfo

func (m *DockerNetwork) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DockerNetwork) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerNetwork) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *DockerNetwork) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *DockerNetwork) GetEnableIpv6() int32 {
	if m != nil {
		return m.EnableIpv6
	}
	return 0
}

func (m *DockerNetwork) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *DockerNetwork) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func init() {
	proto.RegisterType((*DockerNetwork)(nil), "schemas.DockerNetwork")
}

func init() { proto.RegisterFile("pb/docker_network.proto", fileDescriptor_1e1f4548f986cabd) }

var fileDescriptor_1e1f4548f986cabd = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0xb4, 0x4d, 0xc4, 0x21, 0x18, 0x3c, 0xc0, 0x6d, 0x44, 0x4c, 0x59, 0x68, 0x10,
	0x48, 0x61, 0x60, 0x43, 0x2c, 0x2c, 0x0c, 0x19, 0x59, 0x2a, 0xff, 0x39, 0x52, 0x2b, 0x24, 0x36,
	0x8e, 0x93, 0xaa, 0xdf, 0x8c, 0x8f, 0x87, 0xea, 0x34, 0xdb, 0xfd, 0x7e, 0xf7, 0xf4, 0xa4, 0x07,
	0xb7, 0x4e, 0x96, 0xda, 0xaa, 0x96, 0xfc, 0xae, 0xa7, 0x70, 0xb0, 0xbe, 0xdd, 0x3a, 0x6f, 0x83,
	0xe5, 0xd9, 0xa0, 0xf6, 0xd4, 0x89, 0xe1, 0xfe, 0x8f, 0xc1, 0xd5, 0x7b, 0x4c, 0x7c, 0xce, 0x01,
	0x7e, 0x0d, 0x89, 0xd1, 0xc8, 0x72, 0x56, 0x5c, 0xd4, 0x89, 0xd1, 0x9c, 0xc3, 0xba, 0x17, 0x1d,
	0x61, 0x12, 0x4d, 0xbc, 0xf9, 0x0d, 0xa4, 0xda, 0x9b, 0x89, 0x3c, 0xae, 0xa2, 0x3d, 0x13, 0x47,
	0xc8, 0x94, 0x27, 0x11, 0x48, 0xe3, 0x3a, 0x67, 0xc5, 0xaa, 0x5e, 0x90, 0xdf, 0xc1, 0x25, 0xf5,
	0x42, 0xfe, 0xd0, 0xce, 0xb8, 0xa9, 0xc2, 0x4d, 0xce, 0x8a, 0x4d, 0x0d, 0xb3, 0xfa, 0x70, 0x53,
	0x75, 0xaa, 0x1c, 0x46, 0xd9, 0x53, 0xc0, 0x74, 0xae, 0x9c, 0xe9, 0x54, 0xd9, 0x88, 0x40, 0x07,
	0x71, 0xc4, 0x2c, 0x3e, 0x16, 0x7c, 0x7b, 0xfa, 0x7a, 0x6c, 0x4c, 0xd8, 0x8f, 0x72, 0xab, 0x6c,
	0x57, 0x76, 0x46, 0xb5, 0xe4, 0x5e, 0xaa, 0xd2, 0x0e, 0xbf, 0x23, 0xf9, 0xe3, 0x43, 0x1c, 0x2a,
	0xc7, 0xef, 0xd2, 0xb5, 0xcd, 0xeb, 0x79, 0xae, 0x4c, 0xa3, 0x7d, 0xfe, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x22, 0x6c, 0x76, 0xa3, 0x19, 0x01, 0x00, 0x00,
}
