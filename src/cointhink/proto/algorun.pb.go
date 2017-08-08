// Code generated by protoc-gen-go.
// source: proto/algorun.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Algorun_States int32

const (
	Algorun_unknown    Algorun_States = 0
	Algorun_building   Algorun_States = 1
	Algorun_starting   Algorun_States = 2
	Algorun_running    Algorun_States = 3
	Algorun_stopped    Algorun_States = 4
	Algorun_destroying Algorun_States = 5
	Algorun_deleted    Algorun_States = 6
)

var Algorun_States_name = map[int32]string{
	0: "unknown",
	1: "building",
	2: "starting",
	3: "running",
	4: "stopped",
	5: "destroying",
	6: "deleted",
}
var Algorun_States_value = map[string]int32{
	"unknown":    0,
	"building":   1,
	"starting":   2,
	"running":    3,
	"stopped":    4,
	"destroying": 5,
	"deleted":    6,
}

func (x Algorun_States) String() string {
	return proto1.EnumName(Algorun_States_name, int32(x))
}
func (Algorun_States) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0, 0} }

type Algorun struct {
	Id          string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AlgorithmId string `protobuf:"bytes,2,opt,name=AlgorithmId" json:"AlgorithmId,omitempty"`
	AccountId   string `protobuf:"bytes,3,opt,name=AccountId" json:"AccountId,omitempty"`
	ScheduleId  string `protobuf:"bytes,4,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=Status" json:"Status,omitempty"`
	Code        string `protobuf:"bytes,6,opt,name=Code" json:"Code,omitempty"`
}

func (m *Algorun) Reset()                    { *m = Algorun{} }
func (m *Algorun) String() string            { return proto1.CompactTextString(m) }
func (*Algorun) ProtoMessage()               {}
func (*Algorun) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Algorun) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Algorun) GetAlgorithmId() string {
	if m != nil {
		return m.AlgorithmId
	}
	return ""
}

func (m *Algorun) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Algorun) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *Algorun) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Algorun) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto1.RegisterType((*Algorun)(nil), "proto.Algorun")
	proto1.RegisterEnum("proto.Algorun_States", Algorun_States_name, Algorun_States_value)
}

func init() { proto1.RegisterFile("proto/algorun.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x9a, 0xa4, 0x76, 0x2a, 0x25, 0x8c, 0x20, 0x7b, 0x10, 0x29, 0x3d, 0x79, 0xd2,
	0x83, 0x4f, 0x50, 0x3c, 0xed, 0xd5, 0x3e, 0x41, 0x9a, 0x19, 0x9a, 0x60, 0xdc, 0x0d, 0xbb, 0xb3,
	0x88, 0x4f, 0xe1, 0x2b, 0xcb, 0x4e, 0x85, 0xf6, 0xb4, 0xfb, 0x7d, 0xdf, 0x7f, 0x1a, 0xb8, 0x9f,
	0x83, 0x17, 0xff, 0xda, 0x4d, 0x27, 0x1f, 0x92, 0x7b, 0x51, 0xc2, 0x5a, 0x9f, 0xdd, 0x6f, 0x09,
	0xcb, 0xfd, 0x39, 0xe0, 0x06, 0x4a, 0x4b, 0xa6, 0xd8, 0x16, 0xcf, 0xab, 0x8f, 0xd2, 0x12, 0x6e,
	0x61, 0xad, 0x69, 0x94, 0xe1, 0xcb, 0x92, 0x29, 0x35, 0x5c, 0x2b, 0x7c, 0x84, 0xd5, 0xbe, 0xef,
	0x7d, 0x72, 0x62, 0xc9, 0x2c, 0xb4, 0x5f, 0x04, 0x3e, 0x01, 0x1c, 0xfa, 0x81, 0x29, 0x4d, 0x6c,
	0xc9, 0x54, 0x9a, 0xaf, 0x0c, 0x3e, 0x40, 0x73, 0x90, 0x4e, 0x52, 0x34, 0xb5, 0xb6, 0x7f, 0x42,
	0x84, 0xea, 0xdd, 0x13, 0x9b, 0x46, 0xad, 0xfe, 0x77, 0xc3, 0x79, 0xcb, 0x11, 0xd7, 0xb0, 0x4c,
	0xee, 0xd3, 0xf9, 0x6f, 0xd7, 0xde, 0xe0, 0x1d, 0xdc, 0x1e, 0xd3, 0x38, 0xd1, 0xe8, 0x4e, 0x6d,
	0x91, 0x29, 0x4a, 0x17, 0x24, 0x53, 0x99, 0x87, 0x21, 0x39, 0x97, 0x61, 0x91, 0x21, 0x8a, 0x9f,
	0x67, 0xa6, 0xb6, 0xc2, 0x0d, 0x00, 0x71, 0x94, 0xe0, 0x7f, 0x72, 0xac, 0x73, 0x24, 0x9e, 0x58,
	0x98, 0xda, 0xe6, 0xd8, 0xe8, 0x61, 0xde, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xa9, 0x81,
	0xcd, 0x36, 0x01, 0x00, 0x00,
}
