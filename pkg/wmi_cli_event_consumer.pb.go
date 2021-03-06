// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/wmi_cli_event_consumer.proto

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

// WMI CommandLineEventConsumer, which can be used for persistence on Windows. See https://www.blackhat.com/docs/us-15/materials/us-15-Graeber-Abusing-Windows-Management-Instrumentation-WMI-To-Build-A-Persistent%20Asynchronous-And-Fileless-Backdoor-wp.pdf for more details.
type WmiCliEventConsumer struct {
	// Unique name of a consumer.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// Standard string template that specifies the process to be started. This property can be NULL
	CommandLineTemplate string `protobuf:"bytes,2,opt,name=command_line_template,json=commandLineTemplate,proto3" json:"command_line_template"`
	// Module to execute. The string can specify the full path and file name of the module to execute
	ExecutablePath string `protobuf:"bytes,3,opt,name=executable_path,json=executablePath,proto3" json:"executable_path"`
	// The name of the class.
	Class string `protobuf:"bytes,4,opt,name=class,proto3" json:"class"`
	// Relative path to the class or instance.
	RelativePath         string   `protobuf:"bytes,5,opt,name=relative_path,json=relativePath,proto3" json:"relative_path"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WmiCliEventConsumer) Reset()         { *m = WmiCliEventConsumer{} }
func (m *WmiCliEventConsumer) String() string { return proto.CompactTextString(m) }
func (*WmiCliEventConsumer) ProtoMessage()    {}
func (*WmiCliEventConsumer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db29cf87a15ac79, []int{0}
}

func (m *WmiCliEventConsumer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WmiCliEventConsumer.Unmarshal(m, b)
}
func (m *WmiCliEventConsumer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WmiCliEventConsumer.Marshal(b, m, deterministic)
}
func (m *WmiCliEventConsumer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WmiCliEventConsumer.Merge(m, src)
}
func (m *WmiCliEventConsumer) XXX_Size() int {
	return xxx_messageInfo_WmiCliEventConsumer.Size(m)
}
func (m *WmiCliEventConsumer) XXX_DiscardUnknown() {
	xxx_messageInfo_WmiCliEventConsumer.DiscardUnknown(m)
}

var xxx_messageInfo_WmiCliEventConsumer proto.InternalMessageInfo

func (m *WmiCliEventConsumer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WmiCliEventConsumer) GetCommandLineTemplate() string {
	if m != nil {
		return m.CommandLineTemplate
	}
	return ""
}

func (m *WmiCliEventConsumer) GetExecutablePath() string {
	if m != nil {
		return m.ExecutablePath
	}
	return ""
}

func (m *WmiCliEventConsumer) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *WmiCliEventConsumer) GetRelativePath() string {
	if m != nil {
		return m.RelativePath
	}
	return ""
}

func init() {
	proto.RegisterType((*WmiCliEventConsumer)(nil), "schemas.WmiCliEventConsumer")
}

func init() { proto.RegisterFile("pb/wmi_cli_event_consumer.proto", fileDescriptor_5db29cf87a15ac79) }

var fileDescriptor_5db29cf87a15ac79 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcd, 0x4a, 0x43, 0x31,
	0x10, 0x85, 0xa9, 0xb6, 0x8a, 0xc1, 0x1f, 0x48, 0x15, 0xee, 0x4e, 0xd1, 0x85, 0x6e, 0xec, 0x95,
	0x0a, 0xba, 0x70, 0x67, 0x71, 0xe7, 0x42, 0x44, 0x10, 0xdc, 0x84, 0x24, 0x8e, 0xbd, 0x43, 0x33,
	0x49, 0xbc, 0x99, 0x54, 0x7d, 0x3c, 0xdf, 0x4c, 0x1a, 0x23, 0xee, 0x66, 0xce, 0xf7, 0x9d, 0xc5,
	0x11, 0x87, 0xd1, 0xb4, 0x1f, 0x84, 0xca, 0x3a, 0x54, 0xb0, 0x04, 0xcf, 0xca, 0x06, 0x9f, 0x32,
	0x41, 0x3f, 0x89, 0x7d, 0xe0, 0x20, 0x37, 0x93, 0xed, 0x80, 0x74, 0x3a, 0xfe, 0x1e, 0x88, 0xf1,
	0x33, 0xe1, 0xcc, 0xe1, 0xdd, 0xca, 0x9b, 0x55, 0x4d, 0x4a, 0x31, 0xf4, 0x9a, 0xa0, 0x19, 0x1c,
	0x0d, 0xce, 0xb6, 0x1e, 0xcb, 0x2d, 0xa7, 0xe2, 0xc0, 0x06, 0x22, 0xed, 0x5f, 0x95, 0x43, 0x0f,
	0x8a, 0x81, 0xa2, 0xd3, 0x0c, 0xcd, 0x5a, 0x91, 0xc6, 0x15, 0xde, 0xa3, 0x87, 0xa7, 0x8a, 0xe4,
	0xa9, 0xd8, 0x83, 0x4f, 0xb0, 0x99, 0xb5, 0x71, 0xa0, 0xa2, 0xe6, 0xae, 0x59, 0x2f, 0xf6, 0xee,
	0x7f, 0xfc, 0xa0, 0xb9, 0x93, 0xfb, 0x62, 0x64, 0x9d, 0x4e, 0xa9, 0x19, 0x16, 0xfc, 0xfb, 0xc8,
	0x13, 0xb1, 0xd3, 0x83, 0xd3, 0x8c, 0xcb, 0x5a, 0x1e, 0x15, 0xba, 0xfd, 0x17, 0xae, 0xaa, 0xb7,
	0xd3, 0x97, 0x8b, 0x39, 0x72, 0x97, 0xcd, 0xc4, 0x06, 0x6a, 0x09, 0xed, 0x02, 0xe2, 0xf5, 0x55,
	0x1b, 0xd2, 0x7b, 0x86, 0xfe, 0xeb, 0xbc, 0x2c, 0x36, 0xf9, 0xad, 0x8d, 0x8b, 0xf9, 0x4d, 0xdd,
	0x6d, 0x36, 0x4a, 0x7a, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xb2, 0x9a, 0xf8, 0x2a, 0x01,
	0x00, 0x00,
}
