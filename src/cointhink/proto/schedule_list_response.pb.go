// Code generated by protoc-gen-go.
// source: proto/schedule_list_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleListResponse struct {
	Ok        bool        `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message   string      `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Schedules []*Schedule `protobuf:"bytes,3,rep,name=schedules" json:"schedules,omitempty"`
}

func (m *ScheduleListResponse) Reset()                    { *m = ScheduleListResponse{} }
func (m *ScheduleListResponse) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleListResponse) ProtoMessage()               {}
func (*ScheduleListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ScheduleListResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ScheduleListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ScheduleListResponse) GetSchedules() []*Schedule {
	if m != nil {
		return m.Schedules
	}
	return nil
}

func init() {
	proto1.RegisterType((*ScheduleListResponse)(nil), "proto.ScheduleListResponse")
}

func init() { proto1.RegisterFile("proto/schedule_list_response.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e, 0x89,
	0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x4b, 0x0a, 0xb1, 0x82, 0x29, 0x29,
	0x11, 0x54, 0xa5, 0x10, 0x49, 0xa5, 0x7c, 0x2e, 0x91, 0x60, 0xa8, 0x88, 0x4f, 0x66, 0x71, 0x49,
	0x10, 0x54, 0xab, 0x10, 0x1f, 0x17, 0x93, 0x7f, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10,
	0x93, 0x7f, 0xb6, 0x90, 0x04, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0xa4, 0xcb, 0xc5, 0x09, 0x33, 0xb3, 0x58, 0x82, 0x59,
	0x81, 0x59, 0x83, 0xdb, 0x88, 0x1f, 0x62, 0xb8, 0x1e, 0xcc, 0xe4, 0x20, 0x84, 0x8a, 0x24, 0x36,
	0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xfa, 0x8f, 0xab, 0xba, 0x00, 0x00, 0x00,
}
