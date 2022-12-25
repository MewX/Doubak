// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/flag.proto

package proto

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

// Flag defines the names of all runtime flags.
type Flag int32

const (
	// The user name.
	Flag_user Flag = 0
	// Tasks to run.
	Flag_tasks Flag = 1
	// Categories to run on.
	Flag_categories Flag = 2
	// Output path/directory.
	Flag_output_dir Flag = 3
	// Incremental running or starting over with overriding existing files.
	Flag_incremental Flag = 4
	// Proxy used to send each request via.
	Flag_proxy Flag = 5
	// Max number of retries.
	Flag_max_retry Flag = 6
	// Min time between any two requests.
	Flag_req_delay Flag = 7
)

var Flag_name = map[int32]string{
	0: "user",
	1: "tasks",
	2: "categories",
	3: "output_dir",
	4: "incremental",
	5: "proxy",
	6: "max_retry",
	7: "req_delay",
}

var Flag_value = map[string]int32{
	"user":        0,
	"tasks":       1,
	"categories":  2,
	"output_dir":  3,
	"incremental": 4,
	"proxy":       5,
	"max_retry":   6,
	"req_delay":   7,
}

func (x Flag) String() string {
	return proto.EnumName(Flag_name, int32(x))
}

func (Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_840b304d78fa0728, []int{0}
}

func init() {
	proto.RegisterEnum("proto.Flag", Flag_name, Flag_value)
}

func init() { proto.RegisterFile("proto/flag.proto", fileDescriptor_840b304d78fa0728) }

var fileDescriptor_840b304d78fa0728 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x1c, 0x8e, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x46, 0xf9, 0x49, 0x0a, 0x1d, 0x04, 0x8c, 0x7c, 0x05, 0xc4, 0x06, 0xa9, 0xf5, 0x82, 0x1b,
	0xb0, 0xe0, 0x10, 0x6c, 0xa2, 0x49, 0x3c, 0x18, 0xab, 0x76, 0x1d, 0xc6, 0x63, 0xa9, 0xbe, 0x7d,
	0x15, 0xaf, 0xde, 0xf7, 0x16, 0x9f, 0xf4, 0x00, 0x57, 0xc9, 0x9a, 0xed, 0x6f, 0x24, 0x7f, 0xec,
	0xd3, 0x8c, 0x1d, 0x1f, 0x15, 0x86, 0xef, 0x48, 0xde, 0x3c, 0xc2, 0x50, 0x0b, 0x0b, 0xde, 0x98,
	0x3d, 0x8c, 0x4a, 0xe5, 0x54, 0xf0, 0xd6, 0xbc, 0x00, 0x2c, 0xa4, 0xec, 0xb3, 0x04, 0x2e, 0x78,
	0xb7, 0x79, 0xae, 0xba, 0x56, 0x9d, 0x5c, 0x10, 0xbc, 0x37, 0xaf, 0xf0, 0x14, 0xce, 0x8b, 0x70,
	0xe2, 0xb3, 0x52, 0xc4, 0x61, 0xfb, 0xae, 0x92, 0x2f, 0x0d, 0x47, 0xf3, 0x0c, 0xfb, 0x44, 0x97,
	0x49, 0x58, 0xa5, 0xe1, 0x6e, 0x53, 0xe1, 0xff, 0xc9, 0x71, 0xa4, 0x86, 0x0f, 0x5f, 0xef, 0x3f,
	0x6f, 0x3e, 0xe8, 0x5f, 0x9d, 0x8f, 0x4b, 0x4e, 0x36, 0x68, 0x39, 0xa4, 0x76, 0x70, 0xa4, 0x64,
	0x5d, 0xae, 0x33, 0x9d, 0x6c, 0xaf, 0x9b, 0x77, 0x1d, 0x9f, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x75, 0xa0, 0xf9, 0xfe, 0xbf, 0x00, 0x00, 0x00,
}
