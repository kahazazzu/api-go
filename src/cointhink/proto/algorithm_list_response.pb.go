// Code generated by protoc-gen-go.
// source: proto/algorithm_list_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AlgorithmListResponse struct {
	Ok         bool         `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message    string       `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Algorithms []*Algorithm `protobuf:"bytes,3,rep,name=Algorithms" json:"Algorithms,omitempty"`
}

func (m *AlgorithmListResponse) Reset()                    { *m = AlgorithmListResponse{} }
func (m *AlgorithmListResponse) String() string            { return proto1.CompactTextString(m) }
func (*AlgorithmListResponse) ProtoMessage()               {}
func (*AlgorithmListResponse) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

func (m *AlgorithmListResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *AlgorithmListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AlgorithmListResponse) GetAlgorithms() []*Algorithm {
	if m != nil {
		return m.Algorithms
	}
	return nil
}

func init() {
	proto1.RegisterType((*AlgorithmListResponse)(nil), "proto.AlgorithmListResponse")
}

func init() { proto1.RegisterFile("proto/algorithm_list_response.proto", fileDescriptor20) }

var fileDescriptor20 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcc, 0x49, 0xcf, 0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e,
	0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0xcb, 0x0a, 0xb1, 0x82, 0x29,
	0x29, 0x51, 0x34, 0xb5, 0x10, 0x59, 0xa5, 0x62, 0x2e, 0x51, 0x47, 0x98, 0x90, 0x4f, 0x66, 0x71,
	0x49, 0x10, 0x54, 0xb3, 0x10, 0x1f, 0x17, 0x93, 0x7f, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x93, 0x7f, 0xb6, 0x90, 0x04, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x64, 0xc0, 0xc5, 0x05, 0x37, 0xa2, 0x58, 0x82,
	0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x00, 0x62, 0xbc, 0x1e, 0x5c, 0x22, 0x08, 0x49, 0x4d, 0x12,
	0x1b, 0x58, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x87, 0x4a, 0xa8, 0x73, 0xc0, 0x00, 0x00,
	0x00,
}
