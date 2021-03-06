// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/time_machine_destination.proto

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

// Locations backed up to using Time Machine.
type TimeMachineDestination struct {
	// Human readable name of drive
	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias"`
	// Time Machine destination ID
	DestinationId string `protobuf:"bytes,2,opt,name=destination_id,json=destinationId,proto3" json:"destination_id"`
	// Consistency scan date
	ConsistencyScanDate int32 `protobuf:"varint,3,opt,name=consistency_scan_date,json=consistencyScanDate,proto3" json:"consistency_scan_date"`
	// Root UUID of backup volume
	RootVolumeUuid string `protobuf:"bytes,4,opt,name=root_volume_uuid,json=rootVolumeUuid,proto3" json:"root_volume_uuid"`
	// Bytes available on volume
	BytesAvailable int32 `protobuf:"varint,5,opt,name=bytes_available,json=bytesAvailable,proto3" json:"bytes_available"`
	// Bytes used on volume
	BytesUsed int32 `protobuf:"varint,6,opt,name=bytes_used,json=bytesUsed,proto3" json:"bytes_used"`
	// Last known encrypted state
	Encryption           string   `protobuf:"bytes,7,opt,name=encryption,proto3" json:"encryption"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeMachineDestination) Reset()         { *m = TimeMachineDestination{} }
func (m *TimeMachineDestination) String() string { return proto.CompactTextString(m) }
func (*TimeMachineDestination) ProtoMessage()    {}
func (*TimeMachineDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_565c0b6b5f7018b4, []int{0}
}

func (m *TimeMachineDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeMachineDestination.Unmarshal(m, b)
}
func (m *TimeMachineDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeMachineDestination.Marshal(b, m, deterministic)
}
func (m *TimeMachineDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeMachineDestination.Merge(m, src)
}
func (m *TimeMachineDestination) XXX_Size() int {
	return xxx_messageInfo_TimeMachineDestination.Size(m)
}
func (m *TimeMachineDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeMachineDestination.DiscardUnknown(m)
}

var xxx_messageInfo_TimeMachineDestination proto.InternalMessageInfo

func (m *TimeMachineDestination) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *TimeMachineDestination) GetDestinationId() string {
	if m != nil {
		return m.DestinationId
	}
	return ""
}

func (m *TimeMachineDestination) GetConsistencyScanDate() int32 {
	if m != nil {
		return m.ConsistencyScanDate
	}
	return 0
}

func (m *TimeMachineDestination) GetRootVolumeUuid() string {
	if m != nil {
		return m.RootVolumeUuid
	}
	return ""
}

func (m *TimeMachineDestination) GetBytesAvailable() int32 {
	if m != nil {
		return m.BytesAvailable
	}
	return 0
}

func (m *TimeMachineDestination) GetBytesUsed() int32 {
	if m != nil {
		return m.BytesUsed
	}
	return 0
}

func (m *TimeMachineDestination) GetEncryption() string {
	if m != nil {
		return m.Encryption
	}
	return ""
}

func init() {
	proto.RegisterType((*TimeMachineDestination)(nil), "schemas.TimeMachineDestination")
}

func init() { proto.RegisterFile("pb/time_machine_destination.proto", fileDescriptor_565c0b6b5f7018b4) }

var fileDescriptor_565c0b6b5f7018b4 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x70, 0x5a, 0x6d, 0x4b, 0x03, 0x56, 0x89, 0x7f, 0xc8, 0x45, 0xa9, 0x82, 0xd8, 0x8b,
	0x5d, 0xa9, 0xa0, 0x07, 0x4f, 0x4a, 0x2f, 0x1e, 0xbc, 0x54, 0xeb, 0xc1, 0x4b, 0x98, 0x4d, 0xc6,
	0x76, 0xe8, 0x26, 0x59, 0x37, 0x49, 0x61, 0x9f, 0xc5, 0x97, 0x95, 0xa6, 0x8a, 0x7b, 0xcc, 0x6f,
	0x3e, 0xbe, 0x0c, 0xc3, 0xce, 0xcb, 0x3c, 0x0b, 0x64, 0x50, 0x1a, 0x50, 0x4b, 0xb2, 0x28, 0x35,
	0xfa, 0x40, 0x16, 0x02, 0x39, 0x3b, 0x2e, 0x2b, 0x17, 0x1c, 0xef, 0x79, 0xb5, 0x44, 0x03, 0xfe,
	0xe2, 0xbb, 0xcd, 0x4e, 0xde, 0xc8, 0xe0, 0xcb, 0x36, 0x3a, 0xfd, 0x4f, 0xf2, 0x23, 0xd6, 0x81,
	0x82, 0xc0, 0x8b, 0xd6, 0xb0, 0x35, 0xea, 0xcf, 0xb6, 0x0f, 0x7e, 0xc9, 0x06, 0x8d, 0x3a, 0x49,
	0x5a, 0xb4, 0xd3, 0x78, 0xaf, 0xa1, 0xcf, 0x9a, 0x4f, 0xd8, 0xb1, 0x72, 0xd6, 0x93, 0x0f, 0x68,
	0x55, 0x2d, 0xbd, 0x02, 0x2b, 0x35, 0x04, 0x14, 0x3b, 0xc3, 0xd6, 0xa8, 0x33, 0x3b, 0x6c, 0x0c,
	0x5f, 0x15, 0xd8, 0x29, 0x04, 0xe4, 0x23, 0x76, 0x50, 0x39, 0x17, 0xe4, 0xda, 0x15, 0xd1, 0xa0,
	0x8c, 0x91, 0xb4, 0xd8, 0x4d, 0xe5, 0x83, 0x8d, 0xbf, 0x27, 0x9e, 0x47, 0xd2, 0xfc, 0x8a, 0xed,
	0xe7, 0x75, 0x40, 0x2f, 0x61, 0x0d, 0x54, 0x40, 0x5e, 0xa0, 0xe8, 0xa4, 0xde, 0x41, 0xe2, 0xc7,
	0x3f, 0xe5, 0xa7, 0x8c, 0x6d, 0x83, 0xd1, 0xa3, 0x16, 0xdd, 0x94, 0xe9, 0x27, 0x99, 0x7b, 0xd4,
	0xfc, 0x8c, 0x31, 0xb4, 0xaa, 0xaa, 0xcb, 0xcd, 0xd6, 0xa2, 0x97, 0xfe, 0x6a, 0xc8, 0xd3, 0xe4,
	0xe3, 0x66, 0x41, 0x61, 0x19, 0xf3, 0xb1, 0x72, 0x26, 0x33, 0xa4, 0x56, 0x58, 0xde, 0xdf, 0x65,
	0xce, 0x7f, 0x45, 0xac, 0xea, 0xeb, 0x74, 0xcb, 0x3c, 0x7e, 0x66, 0xe5, 0x6a, 0xf1, 0xf0, 0x7b,
	0xd1, 0xbc, 0x9b, 0xf4, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x66, 0xf7, 0x44, 0x13, 0x86, 0x01,
	0x00, 0x00,
}
