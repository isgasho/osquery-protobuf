// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/logon_session.proto

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

// Windows Logon Session.
type LogonSession struct {
	// A locally unique identifier (LUID) that identifies a logon session.
	LogonId int32 `protobuf:"varint,1,opt,name=logon_id,json=logonId,proto3" json:"logon_id"`
	// The account name of the security principal that owns the logon session.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user"`
	// The name of the domain used to authenticate the owner of the logon session.
	LogonDomain string `protobuf:"bytes,3,opt,name=logon_domain,json=logonDomain,proto3" json:"logon_domain"`
	// The authentication package used to authenticate the owner of the logon session.
	AuthenticationPackage string `protobuf:"bytes,4,opt,name=authentication_package,json=authenticationPackage,proto3" json:"authentication_package"`
	// The logon method.
	LogonType string `protobuf:"bytes,5,opt,name=logon_type,json=logonType,proto3" json:"logon_type"`
	// The Terminal Services session identifier.
	SessionId int32 `protobuf:"varint,6,opt,name=session_id,json=sessionId,proto3" json:"session_id"`
	// The user's security identifier (SID).
	LogonSid string `protobuf:"bytes,7,opt,name=logon_sid,json=logonSid,proto3" json:"logon_sid"`
	// The time the session owner logged on.
	LogonTime int64 `protobuf:"varint,8,opt,name=logon_time,json=logonTime,proto3" json:"logon_time"`
	// The name of the server used to authenticate the owner of the logon session.
	LogonServer string `protobuf:"bytes,9,opt,name=logon_server,json=logonServer,proto3" json:"logon_server"`
	// The DNS name for the owner of the logon session.
	DnsDomainName string `protobuf:"bytes,10,opt,name=dns_domain_name,json=dnsDomainName,proto3" json:"dns_domain_name"`
	// The user principal name (UPN) for the owner of the logon session.
	Upn string `protobuf:"bytes,11,opt,name=upn,proto3" json:"upn"`
	// The script used for logging on.
	LogonScript string `protobuf:"bytes,12,opt,name=logon_script,json=logonScript,proto3" json:"logon_script"`
	// The home directory for the logon session.
	ProfilePath string `protobuf:"bytes,13,opt,name=profile_path,json=profilePath,proto3" json:"profile_path"`
	// The home directory for the logon session.
	HomeDirectory string `protobuf:"bytes,14,opt,name=home_directory,json=homeDirectory,proto3" json:"home_directory"`
	// The drive location of the home directory of the logon session.
	HomeDirectoryDrive   string   `protobuf:"bytes,15,opt,name=home_directory_drive,json=homeDirectoryDrive,proto3" json:"home_directory_drive"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogonSession) Reset()         { *m = LogonSession{} }
func (m *LogonSession) String() string { return proto.CompactTextString(m) }
func (*LogonSession) ProtoMessage()    {}
func (*LogonSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_b13f3070512d118c, []int{0}
}

func (m *LogonSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogonSession.Unmarshal(m, b)
}
func (m *LogonSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogonSession.Marshal(b, m, deterministic)
}
func (m *LogonSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogonSession.Merge(m, src)
}
func (m *LogonSession) XXX_Size() int {
	return xxx_messageInfo_LogonSession.Size(m)
}
func (m *LogonSession) XXX_DiscardUnknown() {
	xxx_messageInfo_LogonSession.DiscardUnknown(m)
}

var xxx_messageInfo_LogonSession proto.InternalMessageInfo

func (m *LogonSession) GetLogonId() int32 {
	if m != nil {
		return m.LogonId
	}
	return 0
}

func (m *LogonSession) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *LogonSession) GetLogonDomain() string {
	if m != nil {
		return m.LogonDomain
	}
	return ""
}

func (m *LogonSession) GetAuthenticationPackage() string {
	if m != nil {
		return m.AuthenticationPackage
	}
	return ""
}

func (m *LogonSession) GetLogonType() string {
	if m != nil {
		return m.LogonType
	}
	return ""
}

func (m *LogonSession) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *LogonSession) GetLogonSid() string {
	if m != nil {
		return m.LogonSid
	}
	return ""
}

func (m *LogonSession) GetLogonTime() int64 {
	if m != nil {
		return m.LogonTime
	}
	return 0
}

func (m *LogonSession) GetLogonServer() string {
	if m != nil {
		return m.LogonServer
	}
	return ""
}

func (m *LogonSession) GetDnsDomainName() string {
	if m != nil {
		return m.DnsDomainName
	}
	return ""
}

func (m *LogonSession) GetUpn() string {
	if m != nil {
		return m.Upn
	}
	return ""
}

func (m *LogonSession) GetLogonScript() string {
	if m != nil {
		return m.LogonScript
	}
	return ""
}

func (m *LogonSession) GetProfilePath() string {
	if m != nil {
		return m.ProfilePath
	}
	return ""
}

func (m *LogonSession) GetHomeDirectory() string {
	if m != nil {
		return m.HomeDirectory
	}
	return ""
}

func (m *LogonSession) GetHomeDirectoryDrive() string {
	if m != nil {
		return m.HomeDirectoryDrive
	}
	return ""
}

func init() {
	proto.RegisterType((*LogonSession)(nil), "schemas.LogonSession")
}

func init() { proto.RegisterFile("pb/logon_session.proto", fileDescriptor_b13f3070512d118c) }

var fileDescriptor_b13f3070512d118c = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x89, 0xed, 0x6e, 0x37, 0xb3, 0xdd, 0x5d, 0x19, 0x74, 0x19, 0x11, 0xa1, 0x0a, 0x4a,
	0x2f, 0x36, 0x45, 0x51, 0x0f, 0xde, 0xa4, 0x97, 0x82, 0x48, 0x69, 0x3d, 0x79, 0x09, 0x93, 0xcc,
	0x6b, 0x32, 0xb4, 0xf3, 0xc3, 0x99, 0x49, 0x21, 0x7f, 0x84, 0xff, 0xb3, 0xe4, 0x4d, 0xca, 0xa6,
	0xb7, 0xc9, 0xe7, 0xfb, 0x99, 0x97, 0xf9, 0xc2, 0x23, 0x8f, 0xb6, 0xc8, 0x8e, 0xa6, 0x32, 0x3a,
	0xf7, 0xe0, 0xbd, 0x34, 0x7a, 0x61, 0x9d, 0x09, 0x86, 0x4e, 0x7c, 0x59, 0x83, 0xe2, 0xfe, 0xdd,
	0xbf, 0x31, 0x99, 0xfe, 0xec, 0x84, 0x5d, 0xcc, 0xe9, 0x2b, 0x72, 0x13, 0x2f, 0x48, 0xc1, 0x92,
	0x59, 0x32, 0xbf, 0xda, 0x4e, 0xf0, 0x7b, 0x2d, 0x28, 0x25, 0xe3, 0xc6, 0x83, 0x63, 0xcf, 0x66,
	0xc9, 0x3c, 0xdd, 0xe2, 0x99, 0xbe, 0x25, 0xd3, 0xa8, 0x0b, 0xa3, 0xb8, 0xd4, 0x6c, 0x84, 0xd9,
	0x2d, 0xb2, 0x15, 0x22, 0xfa, 0x85, 0x3c, 0xf2, 0x26, 0xd4, 0xa0, 0x83, 0x2c, 0x79, 0x90, 0x46,
	0xe7, 0x96, 0x97, 0x07, 0x5e, 0x01, 0x1b, 0xa3, 0xfc, 0xf2, 0x32, 0xdd, 0xc4, 0x90, 0xbe, 0x21,
	0x24, 0x4e, 0x0e, 0xad, 0x05, 0x76, 0x85, 0x6a, 0x8a, 0xe4, 0x77, 0x6b, 0x31, 0xee, 0x2b, 0x75,
	0x2f, 0xbd, 0xc6, 0x97, 0xa6, 0x3d, 0x59, 0x0b, 0xfa, 0x9a, 0xa4, 0x7d, 0x6f, 0x29, 0xd8, 0x04,
	0x2f, 0xc7, 0x5e, 0x3b, 0x29, 0x06, 0xa3, 0xa5, 0x02, 0x76, 0x33, 0x4b, 0xe6, 0xa3, 0xf3, 0x68,
	0xa9, 0xe0, 0xa9, 0x93, 0x07, 0x77, 0x02, 0xc7, 0xd2, 0x41, 0xa7, 0x1d, 0x22, 0xfa, 0x81, 0x3c,
	0x08, 0xed, 0xfb, 0xd2, 0xb9, 0xe6, 0x0a, 0x18, 0x41, 0xeb, 0x4e, 0x68, 0x1f, 0x7b, 0xff, 0xe2,
	0x0a, 0xe8, 0x73, 0x32, 0x6a, 0xac, 0x66, 0xb7, 0x98, 0x75, 0xc7, 0xc1, 0xf0, 0xd2, 0x49, 0x1b,
	0xd8, 0x74, 0x38, 0x1c, 0x51, 0xa7, 0x58, 0x67, 0xf6, 0xf2, 0x08, 0xb9, 0xe5, 0xa1, 0x66, 0x77,
	0x51, 0xe9, 0xd9, 0x86, 0x87, 0x9a, 0xbe, 0x27, 0xf7, 0xb5, 0x51, 0x90, 0x0b, 0xe9, 0xa0, 0x0c,
	0xc6, 0xb5, 0xec, 0x3e, 0xfe, 0xbe, 0xa3, 0xab, 0x33, 0xa4, 0x4b, 0xf2, 0xe2, 0x52, 0xcb, 0x85,
	0x93, 0x27, 0x60, 0x0f, 0x28, 0xd3, 0x0b, 0x79, 0xd5, 0x25, 0x3f, 0x3e, 0xfd, 0x59, 0x56, 0x32,
	0xd4, 0x4d, 0xb1, 0x28, 0x8d, 0xca, 0x94, 0x2c, 0x0f, 0x60, 0xbf, 0x7d, 0xcd, 0x8c, 0xff, 0xdb,
	0x80, 0x6b, 0x3f, 0xe2, 0xf6, 0x14, 0xcd, 0x3e, 0xb3, 0x87, 0xea, 0x7b, 0xbf, 0x43, 0xc5, 0x35,
	0xd2, 0xcf, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x87, 0x47, 0xa1, 0xa4, 0x6d, 0x02, 0x00, 0x00,
}
