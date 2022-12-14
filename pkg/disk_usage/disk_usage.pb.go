// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: disk_usage.proto

package disk_usage

import (
	"fmt"
	"math"

	proto "github.com/golang/protobuf/proto"
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

type DiskUsage int32

const (
	DiskUsage_Normal      DiskUsage = 0
	DiskUsage_AlmostFull  DiskUsage = 1
	DiskUsage_AlreadyFull DiskUsage = 2
)

var DiskUsage_name = map[int32]string{
	0: "Normal",
	1: "AlmostFull",
	2: "AlreadyFull",
}

var DiskUsage_value = map[string]int32{
	"Normal":      0,
	"AlmostFull":  1,
	"AlreadyFull": 2,
}

func (x DiskUsage) String() string {
	return proto.EnumName(DiskUsage_name, int32(x))
}

func (DiskUsage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_82acb53abfd87223, []int{0}
}

func init() {
	proto.RegisterEnum("disk_usage.DiskUsage", DiskUsage_name, DiskUsage_value)
}

func init() { proto.RegisterFile("disk_usage.proto", fileDescriptor_82acb53abfd87223) }

var fileDescriptor_82acb53abfd87223 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xc9, 0x2c, 0xce,
	0x8e, 0x2f, 0x2d, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0xf1, 0x17, 0x95, 0x16, 0x97, 0x80, 0x85, 0x21, 0x92, 0x5a, 0x16, 0x5c, 0x9c, 0x2e, 0x99,
	0xc5, 0xd9, 0xa1, 0x20, 0x59, 0x21, 0x2e, 0x2e, 0x36, 0xbf, 0xfc, 0xa2, 0xdc, 0xc4, 0x1c, 0x01,
	0x06, 0x21, 0x3e, 0x2e, 0x2e, 0xc7, 0x9c, 0xdc, 0xfc, 0xe2, 0x12, 0xb7, 0xd2, 0x9c, 0x1c, 0x01,
	0x46, 0x21, 0x7e, 0x2e, 0x6e, 0xc7, 0x9c, 0xa2, 0xd4, 0xc4, 0x94, 0x4a, 0xb0, 0x00, 0x93, 0x93,
	0xda, 0x8d, 0x15, 0x1c, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x5c, 0x02, 0xf9, 0x45, 0xe9, 0x7a, 0x25, 0x99, 0xd9, 0x65,
	0x7a, 0xd9, 0x65, 0x60, 0x1b, 0x92, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x36, 0x6c, 0xcd, 0x99, 0x00, 0x00, 0x00,
}
