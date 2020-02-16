// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/md_device.proto

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

// Software RAID array settings.
type MdDevice struct {
	// md device name
	DeviceName string `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name"`
	// Current state of the array
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status"`
	// Current raid level of the array
	RaidLevel int32 `protobuf:"varint,3,opt,name=raid_level,json=raidLevel,proto3" json:"raid_level"`
	// size of the array in blocks
	Size int64 `protobuf:"varint,4,opt,name=size,proto3" json:"size"`
	// chunk size in bytes
	ChunkSize int64 `protobuf:"varint,5,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size"`
	// Number of configured RAID disks in array
	RaidDisks int32 `protobuf:"varint,6,opt,name=raid_disks,json=raidDisks,proto3" json:"raid_disks"`
	// Number of working disks in array
	WorkingDisks int32 `protobuf:"varint,7,opt,name=working_disks,json=workingDisks,proto3" json:"working_disks"`
	// Number of active disks in array
	ActiveDisks int32 `protobuf:"varint,8,opt,name=active_disks,json=activeDisks,proto3" json:"active_disks"`
	// Number of active disks in array
	FailedDisks int32 `protobuf:"varint,9,opt,name=failed_disks,json=failedDisks,proto3" json:"failed_disks"`
	// Number of active disks in array
	SpareDisks int32 `protobuf:"varint,10,opt,name=spare_disks,json=spareDisks,proto3" json:"spare_disks"`
	// State of the superblock
	SuperblockState string `protobuf:"bytes,11,opt,name=superblock_state,json=superblockState,proto3" json:"superblock_state"`
	// Version of the superblock
	SuperblockVersion string `protobuf:"bytes,12,opt,name=superblock_version,json=superblockVersion,proto3" json:"superblock_version"`
	// Unix timestamp of last update
	SuperblockUpdateTime int64 `protobuf:"varint,13,opt,name=superblock_update_time,json=superblockUpdateTime,proto3" json:"superblock_update_time"`
	// Bitmap chunk size
	BitmapChunkSize string `protobuf:"bytes,14,opt,name=bitmap_chunk_size,json=bitmapChunkSize,proto3" json:"bitmap_chunk_size"`
	// External referenced bitmap file
	BitmapExternalFile string `protobuf:"bytes,15,opt,name=bitmap_external_file,json=bitmapExternalFile,proto3" json:"bitmap_external_file"`
	// Progress of the recovery activity
	RecoveryProgress string `protobuf:"bytes,16,opt,name=recovery_progress,json=recoveryProgress,proto3" json:"recovery_progress"`
	// Estimated duration of recovery activity
	RecoveryFinish string `protobuf:"bytes,17,opt,name=recovery_finish,json=recoveryFinish,proto3" json:"recovery_finish"`
	// Speed of recovery activity
	RecoverySpeed string `protobuf:"bytes,18,opt,name=recovery_speed,json=recoverySpeed,proto3" json:"recovery_speed"`
	// Progress of the resync activity
	ResyncProgress string `protobuf:"bytes,19,opt,name=resync_progress,json=resyncProgress,proto3" json:"resync_progress"`
	// Estimated duration of resync activity
	ResyncFinish string `protobuf:"bytes,20,opt,name=resync_finish,json=resyncFinish,proto3" json:"resync_finish"`
	// Speed of resync activity
	ResyncSpeed string `protobuf:"bytes,21,opt,name=resync_speed,json=resyncSpeed,proto3" json:"resync_speed"`
	// Progress of the reshape activity
	ReshapeProgress string `protobuf:"bytes,22,opt,name=reshape_progress,json=reshapeProgress,proto3" json:"reshape_progress"`
	// Estimated duration of reshape activity
	ReshapeFinish string `protobuf:"bytes,23,opt,name=reshape_finish,json=reshapeFinish,proto3" json:"reshape_finish"`
	// Speed of reshape activity
	ReshapeSpeed string `protobuf:"bytes,24,opt,name=reshape_speed,json=reshapeSpeed,proto3" json:"reshape_speed"`
	// Progress of the resync activity
	CheckArrayProgress string `protobuf:"bytes,25,opt,name=check_array_progress,json=checkArrayProgress,proto3" json:"check_array_progress"`
	// Estimated duration of resync activity
	CheckArrayFinish string `protobuf:"bytes,26,opt,name=check_array_finish,json=checkArrayFinish,proto3" json:"check_array_finish"`
	// Speed of resync activity
	CheckArraySpeed string `protobuf:"bytes,27,opt,name=check_array_speed,json=checkArraySpeed,proto3" json:"check_array_speed"`
	// Unused devices
	UnusedDevices        string   `protobuf:"bytes,28,opt,name=unused_devices,json=unusedDevices,proto3" json:"unused_devices"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MdDevice) Reset()         { *m = MdDevice{} }
func (m *MdDevice) String() string { return proto.CompactTextString(m) }
func (*MdDevice) ProtoMessage()    {}
func (*MdDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bd23044ee5b844a, []int{0}
}

func (m *MdDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdDevice.Unmarshal(m, b)
}
func (m *MdDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdDevice.Marshal(b, m, deterministic)
}
func (m *MdDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdDevice.Merge(m, src)
}
func (m *MdDevice) XXX_Size() int {
	return xxx_messageInfo_MdDevice.Size(m)
}
func (m *MdDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_MdDevice.DiscardUnknown(m)
}

var xxx_messageInfo_MdDevice proto.InternalMessageInfo

func (m *MdDevice) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *MdDevice) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MdDevice) GetRaidLevel() int32 {
	if m != nil {
		return m.RaidLevel
	}
	return 0
}

func (m *MdDevice) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MdDevice) GetChunkSize() int64 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *MdDevice) GetRaidDisks() int32 {
	if m != nil {
		return m.RaidDisks
	}
	return 0
}

func (m *MdDevice) GetWorkingDisks() int32 {
	if m != nil {
		return m.WorkingDisks
	}
	return 0
}

func (m *MdDevice) GetActiveDisks() int32 {
	if m != nil {
		return m.ActiveDisks
	}
	return 0
}

func (m *MdDevice) GetFailedDisks() int32 {
	if m != nil {
		return m.FailedDisks
	}
	return 0
}

func (m *MdDevice) GetSpareDisks() int32 {
	if m != nil {
		return m.SpareDisks
	}
	return 0
}

func (m *MdDevice) GetSuperblockState() string {
	if m != nil {
		return m.SuperblockState
	}
	return ""
}

func (m *MdDevice) GetSuperblockVersion() string {
	if m != nil {
		return m.SuperblockVersion
	}
	return ""
}

func (m *MdDevice) GetSuperblockUpdateTime() int64 {
	if m != nil {
		return m.SuperblockUpdateTime
	}
	return 0
}

func (m *MdDevice) GetBitmapChunkSize() string {
	if m != nil {
		return m.BitmapChunkSize
	}
	return ""
}

func (m *MdDevice) GetBitmapExternalFile() string {
	if m != nil {
		return m.BitmapExternalFile
	}
	return ""
}

func (m *MdDevice) GetRecoveryProgress() string {
	if m != nil {
		return m.RecoveryProgress
	}
	return ""
}

func (m *MdDevice) GetRecoveryFinish() string {
	if m != nil {
		return m.RecoveryFinish
	}
	return ""
}

func (m *MdDevice) GetRecoverySpeed() string {
	if m != nil {
		return m.RecoverySpeed
	}
	return ""
}

func (m *MdDevice) GetResyncProgress() string {
	if m != nil {
		return m.ResyncProgress
	}
	return ""
}

func (m *MdDevice) GetResyncFinish() string {
	if m != nil {
		return m.ResyncFinish
	}
	return ""
}

func (m *MdDevice) GetResyncSpeed() string {
	if m != nil {
		return m.ResyncSpeed
	}
	return ""
}

func (m *MdDevice) GetReshapeProgress() string {
	if m != nil {
		return m.ReshapeProgress
	}
	return ""
}

func (m *MdDevice) GetReshapeFinish() string {
	if m != nil {
		return m.ReshapeFinish
	}
	return ""
}

func (m *MdDevice) GetReshapeSpeed() string {
	if m != nil {
		return m.ReshapeSpeed
	}
	return ""
}

func (m *MdDevice) GetCheckArrayProgress() string {
	if m != nil {
		return m.CheckArrayProgress
	}
	return ""
}

func (m *MdDevice) GetCheckArrayFinish() string {
	if m != nil {
		return m.CheckArrayFinish
	}
	return ""
}

func (m *MdDevice) GetCheckArraySpeed() string {
	if m != nil {
		return m.CheckArraySpeed
	}
	return ""
}

func (m *MdDevice) GetUnusedDevices() string {
	if m != nil {
		return m.UnusedDevices
	}
	return ""
}

func init() {
	proto.RegisterType((*MdDevice)(nil), "schemas.MdDevice")
}

func init() { proto.RegisterFile("pb/md_device.proto", fileDescriptor_4bd23044ee5b844a) }

var fileDescriptor_4bd23044ee5b844a = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x94, 0x4b, 0x6f, 0x13, 0x3d,
	0x14, 0x86, 0x95, 0xaf, 0xd7, 0x9c, 0x24, 0x6d, 0xe2, 0x2f, 0x14, 0x73, 0x53, 0x5b, 0xaa, 0x8a,
	0xb6, 0xd0, 0xa6, 0x02, 0x04, 0x0b, 0x56, 0x40, 0xe9, 0x0a, 0x10, 0x6a, 0x81, 0x05, 0x9b, 0x91,
	0x33, 0x73, 0x92, 0x58, 0x73, 0xc5, 0x9e, 0x09, 0x84, 0x1f, 0xc8, 0xef, 0x42, 0x73, 0x8e, 0xe7,
	0xb2, 0x4b, 0x9e, 0xf7, 0x89, 0xdf, 0xe3, 0x13, 0xc9, 0x20, 0xb2, 0xe9, 0x24, 0x0e, 0xbc, 0x00,
	0x97, 0xda, 0xc7, 0x8b, 0xcc, 0xa4, 0x79, 0x2a, 0xb6, 0xac, 0xbf, 0xc0, 0x58, 0xd9, 0xc7, 0x7f,
	0xb7, 0x61, 0xfb, 0x53, 0x70, 0x45, 0x99, 0xd8, 0x87, 0x1e, 0x5b, 0x5e, 0xa2, 0x62, 0x94, 0x9d,
	0x83, 0xce, 0x49, 0xf7, 0x06, 0x18, 0x7d, 0x56, 0x31, 0x8a, 0x3d, 0xd8, 0xb4, 0xb9, 0xca, 0x0b,
	0x2b, 0xff, 0xa3, 0xcc, 0x7d, 0x13, 0x8f, 0x00, 0x8c, 0xd2, 0x81, 0x17, 0xe1, 0x12, 0x23, 0xb9,
	0x76, 0xd0, 0x39, 0xd9, 0xb8, 0xe9, 0x96, 0xe4, 0x63, 0x09, 0x84, 0x80, 0x75, 0xab, 0xff, 0xa0,
	0x5c, 0x3f, 0xe8, 0x9c, 0xac, 0xdd, 0xd0, 0xe7, 0xf2, 0x27, 0xfe, 0xa2, 0x48, 0x42, 0x8f, 0x92,
	0x0d, 0x4a, 0xba, 0x44, 0x6e, 0x5d, 0x4c, 0x27, 0x06, 0xda, 0x86, 0x56, 0x6e, 0x36, 0x27, 0x5e,
	0x95, 0x40, 0x1c, 0xc1, 0xe0, 0x57, 0x6a, 0x42, 0x9d, 0xcc, 0x9d, 0xb1, 0x45, 0x46, 0xdf, 0x41,
	0x96, 0x0e, 0xa1, 0xaf, 0xfc, 0x5c, 0x2f, 0xd1, 0x39, 0xdb, 0xe4, 0xf4, 0x98, 0xd5, 0xca, 0x4c,
	0xe9, 0x08, 0xab, 0xa2, 0x2e, 0x2b, 0xcc, 0x58, 0xd9, 0x87, 0x9e, 0xcd, 0x94, 0xa9, 0x0e, 0x01,
	0x32, 0x80, 0x10, 0x0b, 0xa7, 0x30, 0xb4, 0x45, 0x86, 0x66, 0x1a, 0xa5, 0x7e, 0xe8, 0x95, 0x1b,
	0x41, 0xd9, 0xa3, 0xf5, 0xec, 0x36, 0xfc, 0xb6, 0xc4, 0xe2, 0x1c, 0x44, 0x4b, 0x5d, 0xa2, 0xb1,
	0x3a, 0x4d, 0x64, 0x9f, 0xe4, 0x51, 0x93, 0x7c, 0xe7, 0x40, 0xbc, 0x84, 0xbd, 0x96, 0x5e, 0x64,
	0x81, 0xca, 0xd1, 0xcb, 0x75, 0x8c, 0x72, 0x40, 0xfb, 0x1a, 0x37, 0xe9, 0x37, 0x0a, 0xbf, 0xea,
	0x18, 0xc5, 0x19, 0x8c, 0xa6, 0x3a, 0x8f, 0x55, 0xe6, 0xb5, 0x16, 0xbc, 0xc3, 0x03, 0x71, 0xf0,
	0xbe, 0x5e, 0xf3, 0x25, 0x8c, 0x9d, 0x8b, 0xbf, 0x73, 0x34, 0x89, 0x8a, 0xbc, 0x99, 0x8e, 0x50,
	0xee, 0x92, 0x2e, 0x38, 0xfb, 0xe0, 0xa2, 0x6b, 0x1d, 0xa1, 0x78, 0x0a, 0x23, 0x83, 0x7e, 0xba,
	0x44, 0xb3, 0xf2, 0x32, 0x93, 0xce, 0x0d, 0x5a, 0x2b, 0x87, 0xa4, 0x0f, 0xab, 0xe0, 0x8b, 0xe3,
	0xe2, 0x09, 0xec, 0xd6, 0xf2, 0x4c, 0x27, 0xda, 0x2e, 0xe4, 0x88, 0xd4, 0x9d, 0x0a, 0x5f, 0x13,
	0x15, 0xc7, 0x50, 0x13, 0xcf, 0x66, 0x88, 0x81, 0x14, 0xe4, 0x0d, 0x2a, 0x7a, 0x5b, 0x42, 0x3e,
	0xcf, 0xae, 0x12, 0xbf, 0xa9, 0xfe, 0xbf, 0x3a, 0xaf, 0xc4, 0x75, 0xf1, 0x11, 0x0c, 0x9c, 0xe8,
	0x6a, 0xc7, 0xa4, 0xf5, 0x19, 0xba, 0xd2, 0x43, 0x70, 0xdf, 0x5d, 0xe5, 0x1d, 0x72, 0x7a, 0xcc,
	0xb8, 0xf0, 0x14, 0x86, 0x06, 0xed, 0x42, 0x65, 0xd8, 0x34, 0xee, 0xf1, 0x2a, 0x1d, 0xaf, 0x2b,
	0xe9, 0x0a, 0xac, 0xba, 0xce, 0xbb, 0xd5, 0x15, 0x88, 0xba, 0x52, 0x9e, 0x8c, 0x34, 0x6e, 0x95,
	0xf5, 0x64, 0x25, 0xe4, 0xda, 0x4b, 0x18, 0xfb, 0x0b, 0xf4, 0x43, 0x4f, 0x19, 0xa3, 0x5a, 0x7b,
	0xbe, 0xc7, 0x7f, 0x0b, 0x65, 0x6f, 0xcb, 0xa8, 0x6e, 0x7f, 0x06, 0xa2, 0xfd, 0x0b, 0x37, 0xc1,
	0x7d, 0xfe, 0x5f, 0x1a, 0xdf, 0x0d, 0x71, 0x06, 0xa3, 0xb6, 0xcd, 0x83, 0x3c, 0xe0, 0x7b, 0x35,
	0x32, 0xcf, 0x72, 0x0c, 0x3b, 0x45, 0x52, 0x58, 0xac, 0x5e, 0x10, 0x2b, 0x1f, 0xf2, 0xbd, 0x98,
	0xf2, 0xd3, 0x61, 0xdf, 0x3d, 0xff, 0x71, 0x39, 0xd7, 0xf9, 0xa2, 0x98, 0x5e, 0xf8, 0x69, 0x3c,
	0x89, 0xb5, 0x1f, 0x62, 0xf6, 0xfa, 0xd5, 0x24, 0xb5, 0x3f, 0x0b, 0x34, 0xab, 0x73, 0x7a, 0x76,
	0xa6, 0xc5, 0x6c, 0x92, 0x85, 0xf3, 0x37, 0xee, 0xf1, 0x99, 0x6e, 0x12, 0x7d, 0xf1, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x3b, 0x3e, 0xcd, 0xb1, 0xa2, 0x04, 0x00, 0x00,
}