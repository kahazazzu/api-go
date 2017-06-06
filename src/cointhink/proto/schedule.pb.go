// Code generated by protoc-gen-go.
// source: proto/schedule.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Schedule_States int32

const (
	Schedule_stopped Schedule_States = 0
	Schedule_running Schedule_States = 1
)

var Schedule_States_name = map[int32]string{
	0: "stopped",
	1: "running",
}
var Schedule_States_value = map[string]int32{
	"stopped": 0,
	"running": 1,
}

func (x Schedule_States) String() string {
	return proto1.EnumName(Schedule_States_name, int32(x))
}
func (Schedule_States) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0, 0} }

type Schedule struct {
	Id           string          `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId    string          `protobuf:"bytes,2,opt,name=AccountId" json:"AccountId,omitempty"`
	AlgorithmId  string          `protobuf:"bytes,3,opt,name=AlgorithmId" json:"AlgorithmId,omitempty"`
	Status       Schedule_States `protobuf:"varint,4,opt,name=Status,enum=proto.Schedule_States" json:"Status,omitempty"`
	InitialState string          `protobuf:"bytes,5,opt,name=InitialState" json:"InitialState,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto1.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Schedule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schedule) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Schedule) GetAlgorithmId() string {
	if m != nil {
		return m.AlgorithmId
	}
	return ""
}

func (m *Schedule) GetStatus() Schedule_States {
	if m != nil {
		return m.Status
	}
	return Schedule_stopped
}

func (m *Schedule) GetInitialState() string {
	if m != nil {
		return m.InitialState
	}
	return ""
}

func init() {
	proto1.RegisterType((*Schedule)(nil), "proto.Schedule")
	proto1.RegisterEnum("proto.Schedule_States", Schedule_States_name, Schedule_States_value)
}

func init() { proto1.RegisterFile("proto/schedule.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8d, 0xc1, 0x8a, 0x83, 0x30,
	0x14, 0x45, 0x27, 0xce, 0xe8, 0x8c, 0xcf, 0x41, 0x24, 0x0c, 0x43, 0x16, 0x5d, 0x48, 0x56, 0xae,
	0x52, 0x68, 0xbf, 0xc0, 0x65, 0xb6, 0xfa, 0x05, 0xd6, 0x04, 0x0d, 0xd8, 0x44, 0x4c, 0xf2, 0xa3,
	0xfd, 0xa2, 0x62, 0xb4, 0xb4, 0x5d, 0x3d, 0xee, 0x39, 0x8f, 0x7b, 0xe1, 0x6f, 0x5e, 0x8c, 0x33,
	0x47, 0xdb, 0x8f, 0x52, 0xf8, 0x49, 0xb2, 0x10, 0x71, 0x1c, 0x0e, 0xbd, 0x21, 0xf8, 0x69, 0x77,
	0x83, 0x73, 0x88, 0xb8, 0x20, 0xa8, 0x44, 0x55, 0xda, 0x44, 0x5c, 0xe0, 0x03, 0xa4, 0x75, 0xdf,
	0x1b, 0xaf, 0x1d, 0x17, 0x24, 0x0a, 0xf8, 0x09, 0x70, 0x09, 0x59, 0x3d, 0x0d, 0x66, 0x51, 0x6e,
	0xbc, 0x72, 0x41, 0x3e, 0x83, 0x7f, 0x45, 0x98, 0x41, 0xd2, 0xba, 0xce, 0x79, 0x4b, 0xbe, 0x4a,
	0x54, 0xe5, 0xa7, 0xff, 0x6d, 0x9b, 0x3d, 0x06, 0xd9, 0x6a, 0xa5, 0x6d, 0xf6, 0x2f, 0x4c, 0xe1,
	0x97, 0x6b, 0xe5, 0x54, 0x37, 0x05, 0x41, 0xe2, 0x50, 0xf9, 0xc6, 0x28, 0xdd, 0x3a, 0xa5, 0xc5,
	0x19, 0x7c, 0x5b, 0x67, 0xe6, 0x59, 0x8a, 0xe2, 0x63, 0x0d, 0x8b, 0xd7, 0x5a, 0xe9, 0xa1, 0x40,
	0x97, 0x24, 0xcc, 0x9c, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01, 0xe7, 0x75, 0xe2, 0xfa, 0x00,
	0x00, 0x00,
}
