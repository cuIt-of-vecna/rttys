// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rtty.proto

/*
Package rtty is a generated protocol buffer package.

It is generated from these files:
	rtty.proto

It has these top-level messages:
	RttyMessage
*/
package rtty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RttyMessage_Type int32

const (
	RttyMessage_UNKNOWN  RttyMessage_Type = 0
	RttyMessage_LOGIN    RttyMessage_Type = 1
	RttyMessage_LOGINACK RttyMessage_Type = 2
	RttyMessage_LOGOUT   RttyMessage_Type = 3
	RttyMessage_TTY      RttyMessage_Type = 4
	RttyMessage_UPFILE   RttyMessage_Type = 5
	RttyMessage_DOWNFILE RttyMessage_Type = 6
	RttyMessage_COMMAND  RttyMessage_Type = 7
	RttyMessage_WINSIZE  RttyMessage_Type = 8
)

var RttyMessage_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOGIN",
	2: "LOGINACK",
	3: "LOGOUT",
	4: "TTY",
	5: "UPFILE",
	6: "DOWNFILE",
	7: "COMMAND",
	8: "WINSIZE",
}
var RttyMessage_Type_value = map[string]int32{
	"UNKNOWN":  0,
	"LOGIN":    1,
	"LOGINACK": 2,
	"LOGOUT":   3,
	"TTY":      4,
	"UPFILE":   5,
	"DOWNFILE": 6,
	"COMMAND":  7,
	"WINSIZE":  8,
}

func (x RttyMessage_Type) String() string {
	return proto.EnumName(RttyMessage_Type_name, int32(x))
}
func (RttyMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type RttyMessage_LoginCode int32

const (
	RttyMessage_OK      RttyMessage_LoginCode = 0
	RttyMessage_OFFLINE RttyMessage_LoginCode = 1
)

var RttyMessage_LoginCode_name = map[int32]string{
	0: "OK",
	1: "OFFLINE",
}
var RttyMessage_LoginCode_value = map[string]int32{
	"OK":      0,
	"OFFLINE": 1,
}

func (x RttyMessage_LoginCode) String() string {
	return proto.EnumName(RttyMessage_LoginCode_name, int32(x))
}
func (RttyMessage_LoginCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type RttyMessage_FileCode int32

const (
	RttyMessage_START     RttyMessage_FileCode = 0
	RttyMessage_RATELIMIT RttyMessage_FileCode = 1
	RttyMessage_FILEDATA  RttyMessage_FileCode = 2
	RttyMessage_CANCELED  RttyMessage_FileCode = 3
	RttyMessage_END       RttyMessage_FileCode = 4
)

var RttyMessage_FileCode_name = map[int32]string{
	0: "START",
	1: "RATELIMIT",
	2: "FILEDATA",
	3: "CANCELED",
	4: "END",
}
var RttyMessage_FileCode_value = map[string]int32{
	"START":     0,
	"RATELIMIT": 1,
	"FILEDATA":  2,
	"CANCELED":  3,
	"END":       4,
}

func (x RttyMessage_FileCode) String() string {
	return proto.EnumName(RttyMessage_FileCode_name, int32(x))
}
func (RttyMessage_FileCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

type RttyMessage_CommandErr int32

const (
	RttyMessage_NONE           RttyMessage_CommandErr = 0
	RttyMessage_TIMEOUT        RttyMessage_CommandErr = 1
	RttyMessage_NOTFOUND       RttyMessage_CommandErr = 2
	RttyMessage_READ           RttyMessage_CommandErr = 3
	RttyMessage_PERMISSION     RttyMessage_CommandErr = 4
	RttyMessage_SYSCALL        RttyMessage_CommandErr = 5
	RttyMessage_DEV_OFFLINE    RttyMessage_CommandErr = 6
	RttyMessage_CMD_REQUIRED   RttyMessage_CommandErr = 7
	RttyMessage_DEVID_REQUIRED RttyMessage_CommandErr = 8
)

var RttyMessage_CommandErr_name = map[int32]string{
	0: "NONE",
	1: "TIMEOUT",
	2: "NOTFOUND",
	3: "READ",
	4: "PERMISSION",
	5: "SYSCALL",
	6: "DEV_OFFLINE",
	7: "CMD_REQUIRED",
	8: "DEVID_REQUIRED",
}
var RttyMessage_CommandErr_value = map[string]int32{
	"NONE":           0,
	"TIMEOUT":        1,
	"NOTFOUND":       2,
	"READ":           3,
	"PERMISSION":     4,
	"SYSCALL":        5,
	"DEV_OFFLINE":    6,
	"CMD_REQUIRED":   7,
	"DEVID_REQUIRED": 8,
}

func (x RttyMessage_CommandErr) String() string {
	return proto.EnumName(RttyMessage_CommandErr_name, int32(x))
}
func (RttyMessage_CommandErr) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

type RttyMessage struct {
	Version  uint32              `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Type     RttyMessage_Type    `protobuf:"varint,2,opt,name=type,enum=RttyMessage_Type" json:"type,omitempty"`
	Sid      string              `protobuf:"bytes,3,opt,name=sid" json:"sid,omitempty"`
	Code     int32               `protobuf:"varint,4,opt,name=code" json:"code,omitempty"`
	Data     []byte              `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Name     string              `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Size     uint32              `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	Id       uint32              `protobuf:"varint,8,opt,name=id" json:"id,omitempty"`
	Err      int32               `protobuf:"varint,9,opt,name=err" json:"err,omitempty"`
	Username string              `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Password string              `protobuf:"bytes,11,opt,name=password" json:"password,omitempty"`
	StdOut   string              `protobuf:"bytes,12,opt,name=std_out,json=stdOut" json:"std_out,omitempty"`
	StdErr   string              `protobuf:"bytes,13,opt,name=std_err,json=stdErr" json:"std_err,omitempty"`
	Params   []string            `protobuf:"bytes,14,rep,name=params" json:"params,omitempty"`
	Env      map[string]string   `protobuf:"bytes,15,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Filelist []*RttyMessage_File `protobuf:"bytes,16,rep,name=filelist" json:"filelist,omitempty"`
	Cols     uint32              `protobuf:"varint,17,opt,name=cols" json:"cols,omitempty"`
	Rows     uint32              `protobuf:"varint,18,opt,name=rows" json:"rows,omitempty"`
}

func (m *RttyMessage) Reset()                    { *m = RttyMessage{} }
func (m *RttyMessage) String() string            { return proto.CompactTextString(m) }
func (*RttyMessage) ProtoMessage()               {}
func (*RttyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RttyMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RttyMessage) GetType() RttyMessage_Type {
	if m != nil {
		return m.Type
	}
	return RttyMessage_UNKNOWN
}

func (m *RttyMessage) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *RttyMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RttyMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RttyMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RttyMessage) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RttyMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RttyMessage) GetErr() int32 {
	if m != nil {
		return m.Err
	}
	return 0
}

func (m *RttyMessage) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RttyMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RttyMessage) GetStdOut() string {
	if m != nil {
		return m.StdOut
	}
	return ""
}

func (m *RttyMessage) GetStdErr() string {
	if m != nil {
		return m.StdErr
	}
	return ""
}

func (m *RttyMessage) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *RttyMessage) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *RttyMessage) GetFilelist() []*RttyMessage_File {
	if m != nil {
		return m.Filelist
	}
	return nil
}

func (m *RttyMessage) GetCols() uint32 {
	if m != nil {
		return m.Cols
	}
	return 0
}

func (m *RttyMessage) GetRows() uint32 {
	if m != nil {
		return m.Rows
	}
	return 0
}

type RttyMessage_File struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Dir   bool   `protobuf:"varint,2,opt,name=dir" json:"dir,omitempty"`
	Mtime uint64 `protobuf:"varint,3,opt,name=mtime" json:"mtime,omitempty"`
	Size  uint64 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *RttyMessage_File) Reset()                    { *m = RttyMessage_File{} }
func (m *RttyMessage_File) String() string            { return proto.CompactTextString(m) }
func (*RttyMessage_File) ProtoMessage()               {}
func (*RttyMessage_File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *RttyMessage_File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RttyMessage_File) GetDir() bool {
	if m != nil {
		return m.Dir
	}
	return false
}

func (m *RttyMessage_File) GetMtime() uint64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *RttyMessage_File) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*RttyMessage)(nil), "rtty_message")
	proto.RegisterType((*RttyMessage_File)(nil), "rtty_message.File")
	proto.RegisterEnum("RttyMessage_Type", RttyMessage_Type_name, RttyMessage_Type_value)
	proto.RegisterEnum("RttyMessage_LoginCode", RttyMessage_LoginCode_name, RttyMessage_LoginCode_value)
	proto.RegisterEnum("RttyMessage_FileCode", RttyMessage_FileCode_name, RttyMessage_FileCode_value)
	proto.RegisterEnum("RttyMessage_CommandErr", RttyMessage_CommandErr_name, RttyMessage_CommandErr_value)
}

func init() { proto.RegisterFile("rtty.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x53, 0x4d, 0x6f, 0xdb, 0x38,
	0x10, 0x8d, 0x3e, 0x2c, 0xcb, 0x63, 0xc7, 0xe1, 0x12, 0x8b, 0x2c, 0x91, 0x93, 0xe0, 0xc3, 0xc2,
	0x27, 0x1f, 0xb2, 0xc0, 0xa2, 0xe8, 0x4d, 0x90, 0xe8, 0x40, 0x88, 0x4c, 0xa5, 0xb4, 0x9c, 0x20,
	0xbd, 0x18, 0x6a, 0xc5, 0x06, 0x42, 0x6d, 0xc9, 0x90, 0x64, 0x07, 0xee, 0x9f, 0xe8, 0x3f, 0xec,
	0x6f, 0x29, 0x86, 0xaa, 0x8d, 0xb4, 0xb7, 0xf7, 0xde, 0x0c, 0x1f, 0x87, 0x4f, 0x23, 0x80, 0xba,
	0x6d, 0x8f, 0xb3, 0x5d, 0x5d, 0xb5, 0xd5, 0xe4, 0x47, 0x1f, 0x46, 0x48, 0xd7, 0x5b, 0xd5, 0x34,
	0xd9, 0x8b, 0xa2, 0x0c, 0xfa, 0x07, 0x55, 0x37, 0x45, 0x55, 0x32, 0xc3, 0x33, 0xa6, 0x97, 0xf2,
	0x44, 0xe9, 0xbf, 0x60, 0xb7, 0xc7, 0x9d, 0x62, 0xa6, 0x67, 0x4c, 0xc7, 0xb7, 0x74, 0xf6, 0xf6,
	0xd8, 0x2c, 0x3d, 0xee, 0x94, 0xd4, 0x75, 0x4a, 0xc0, 0x6a, 0x8a, 0x9c, 0x59, 0x9e, 0x31, 0x1d,
	0x48, 0x84, 0x94, 0x82, 0xfd, 0xb9, 0xca, 0x15, 0xb3, 0x3d, 0x63, 0xda, 0x93, 0x1a, 0xa3, 0x96,
	0x67, 0x6d, 0xc6, 0x7a, 0x9e, 0x31, 0x1d, 0x49, 0x8d, 0x51, 0x2b, 0xb3, 0xad, 0x62, 0x8e, 0x3e,
	0xaa, 0x31, 0x6a, 0x4d, 0xf1, 0x4d, 0xb1, 0xbe, 0x1e, 0x46, 0x63, 0x3a, 0x06, 0xb3, 0xc8, 0x99,
	0xab, 0x15, 0xb3, 0xc8, 0xf1, 0x46, 0x55, 0xd7, 0x6c, 0xa0, 0xed, 0x11, 0xd2, 0x1b, 0x70, 0xf7,
	0x8d, 0xaa, 0xb5, 0x1b, 0x68, 0xb7, 0x33, 0xc7, 0xda, 0x2e, 0x6b, 0x9a, 0xd7, 0xaa, 0xce, 0xd9,
	0xb0, 0xab, 0x9d, 0x38, 0xfd, 0x07, 0xfa, 0x4d, 0x9b, 0xaf, 0xab, 0x7d, 0xcb, 0x46, 0xba, 0xe4,
	0x34, 0x6d, 0x9e, 0xec, 0xdb, 0x53, 0x01, 0xaf, 0xb9, 0x3c, 0x17, 0x78, 0x5d, 0xd3, 0x6b, 0x70,
	0x76, 0x59, 0x9d, 0x6d, 0x1b, 0x36, 0xf6, 0x2c, 0xd4, 0x3b, 0x46, 0xa7, 0x60, 0xa9, 0xf2, 0xc0,
	0xae, 0x3c, 0x6b, 0x3a, 0xbc, 0xbd, 0xfe, 0x3d, 0x2c, 0x5e, 0x1e, 0x78, 0xd9, 0xd6, 0x47, 0x89,
	0x2d, 0x74, 0x06, 0xee, 0x97, 0x62, 0xa3, 0x36, 0x45, 0xd3, 0x32, 0xa2, 0xdb, 0xff, 0xc8, 0x76,
	0x5e, 0x6c, 0x94, 0x3c, 0xf7, 0x74, 0x69, 0x6e, 0x1a, 0xf6, 0x57, 0x97, 0x08, 0x62, 0xd4, 0xea,
	0xea, 0xb5, 0x61, 0xb4, 0xd3, 0x10, 0xdf, 0x3c, 0x82, 0x8d, 0x27, 0xcf, 0xa9, 0x1a, 0x6f, 0x52,
	0x25, 0x60, 0xe5, 0x45, 0xad, 0x3f, 0xa5, 0x2b, 0x11, 0xd2, 0xbf, 0xa1, 0xb7, 0x6d, 0x8b, 0xad,
	0xd2, 0xdf, 0xcd, 0x96, 0x1d, 0x39, 0xa7, 0x6f, 0x6b, 0x51, 0xe3, 0x9b, 0xff, 0xc1, 0x3d, 0x3d,
	0x00, 0x7d, 0xbe, 0xaa, 0xe3, 0x2f, 0x6b, 0x84, 0xe8, 0x73, 0xc8, 0x36, 0xfb, 0x6e, 0x4d, 0x06,
	0xb2, 0x23, 0xef, 0xcd, 0x77, 0xc6, 0x64, 0x0f, 0x36, 0x6e, 0x09, 0x1d, 0x42, 0x7f, 0x25, 0xee,
	0x45, 0xf2, 0x24, 0xc8, 0x05, 0x1d, 0x40, 0x2f, 0x4e, 0xee, 0x22, 0x41, 0x0c, 0x3a, 0x02, 0x57,
	0x43, 0x3f, 0xb8, 0x27, 0x26, 0x05, 0x70, 0xe2, 0xe4, 0x2e, 0x59, 0xa5, 0xc4, 0xa2, 0x7d, 0xb0,
	0xd2, 0xf4, 0x99, 0xd8, 0x28, 0xae, 0x1e, 0xe6, 0x51, 0xcc, 0x49, 0x0f, 0xdb, 0xc3, 0xe4, 0x49,
	0x68, 0xe6, 0xa0, 0x69, 0x90, 0x2c, 0x16, 0xbe, 0x08, 0x49, 0x1f, 0xc9, 0x53, 0x24, 0x96, 0xd1,
	0x47, 0x4e, 0xdc, 0x89, 0x07, 0x83, 0xb8, 0x7a, 0x29, 0xca, 0x00, 0xb7, 0xce, 0x01, 0x33, 0xb9,
	0x27, 0x17, 0xd8, 0x91, 0xcc, 0xe7, 0x71, 0x24, 0x38, 0x31, 0x26, 0x11, 0xb8, 0x18, 0x94, 0x6e,
	0x18, 0x40, 0x6f, 0x99, 0xfa, 0x32, 0x25, 0x17, 0xf4, 0x12, 0x06, 0xd2, 0x4f, 0x79, 0x1c, 0x2d,
	0xa2, 0xb4, 0x1b, 0x0f, 0xef, 0x0a, 0xfd, 0xd4, 0x27, 0x26, 0xb2, 0xc0, 0x17, 0x01, 0x8f, 0x79,
	0xd8, 0x0d, 0xc8, 0x45, 0x48, 0xec, 0xc9, 0x77, 0x03, 0x20, 0xa8, 0xb6, 0xdb, 0xac, 0xd4, 0xcb,
	0xe1, 0x82, 0x2d, 0x12, 0xc1, 0xbb, 0x0b, 0xd3, 0x68, 0xc1, 0xf1, 0x3d, 0xda, 0x4a, 0x24, 0xe9,
	0x3c, 0x59, 0x89, 0x90, 0x98, 0xd8, 0x24, 0xb9, 0x8f, 0x36, 0x63, 0x80, 0x07, 0x2e, 0x17, 0xd1,
	0x72, 0x19, 0x25, 0x82, 0xd8, 0x78, 0x68, 0xf9, 0xbc, 0x0c, 0xfc, 0x38, 0x26, 0x3d, 0x7a, 0x05,
	0xc3, 0x90, 0x3f, 0xae, 0x4f, 0x63, 0x3b, 0x94, 0xc0, 0x28, 0x58, 0x84, 0x6b, 0xc9, 0x3f, 0xac,
	0x22, 0xc9, 0xf1, 0xdd, 0x14, 0xc6, 0x21, 0x7f, 0x8c, 0xde, 0x68, 0xee, 0x27, 0x47, 0xff, 0xe7,
	0xff, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x72, 0xf3, 0xbb, 0x98, 0xf5, 0x03, 0x00, 0x00,
}