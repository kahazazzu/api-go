// Code generated by protoc-gen-go.
// source: proto/schedule_run.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleRun struct {
	Schedule *Schedule `protobuf:"bytes,1,opt,name=Schedule" json:"Schedule,omitempty"`
	Run      *Algorun  `protobuf:"bytes,2,opt,name=Run" json:"Run,omitempty"`
}

func (m *ScheduleRun) Reset()                    { *m = ScheduleRun{} }
func (m *ScheduleRun) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleRun) ProtoMessage()               {}
func (*ScheduleRun) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ScheduleRun) GetSchedule() *Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (m *ScheduleRun) GetRun() *Algorun {
	if m != nil {
		return m.Run
	}
	return nil
}

func init() {
	proto1.RegisterType((*ScheduleRun)(nil), "proto.ScheduleRun")
}

func init() { proto1.RegisterFile("proto/schedule_run.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0x2f, 0x2a, 0xcd, 0xd3, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x29, 0x29, 0x11, 0x54, 0x05, 0x10, 0x49, 0x29, 0x61, 0x88, 0x68, 0x62,
	0x4e, 0x7a, 0x3e, 0x5c, 0x87, 0x52, 0x0c, 0x17, 0x77, 0x30, 0x54, 0x59, 0x50, 0x69, 0x9e, 0x90,
	0x36, 0x17, 0x07, 0x8c, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0x0f, 0x51, 0xa8, 0x07,
	0x57, 0x05, 0x57, 0x20, 0xa4, 0xc0, 0xc5, 0x1c, 0x54, 0x9a, 0x27, 0xc1, 0x04, 0x56, 0xc7, 0x07,
	0x55, 0xe7, 0x08, 0x31, 0x3e, 0x08, 0x24, 0x95, 0xc4, 0x06, 0x16, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0xae, 0x04, 0xf3, 0xb2, 0x00, 0x00, 0x00,
}