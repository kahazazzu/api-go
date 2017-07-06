// Code generated by protoc-gen-go.
// source: proto/rpc.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Rpc struct {
	Id     string               `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Method string               `protobuf:"bytes,2,opt,name=Method" json:"Method,omitempty"`
	Token  string               `protobuf:"bytes,3,opt,name=Token" json:"Token,omitempty"`
	Object *google_protobuf.Any `protobuf:"bytes,4,opt,name=Object" json:"Object,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto1.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *Rpc) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rpc) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Rpc) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Rpc) GetObject() *google_protobuf.Any {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto1.RegisterType((*Rpc)(nil), "proto.Rpc")
}

func init() { proto1.RegisterFile("proto/rpc.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0xb3, 0x84, 0x58, 0xc1, 0x94, 0x94, 0x64, 0x7a, 0x7e,
	0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x98, 0x57, 0x09, 0x51, 0xa1,
	0x54, 0xc8, 0xc5, 0x1c, 0x54, 0x90, 0x2c, 0xc4, 0xc7, 0xc5, 0xe4, 0x99, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0xc4, 0xe4, 0x99, 0x22, 0x24, 0xc6, 0xc5, 0xe6, 0x9b, 0x5a, 0x92, 0x91, 0x9f,
	0x22, 0xc1, 0x04, 0x16, 0x83, 0xf2, 0x84, 0x44, 0xb8, 0x58, 0x43, 0xf2, 0xb3, 0x53, 0xf3, 0x24,
	0x98, 0xc1, 0xc2, 0x10, 0x8e, 0x90, 0x0e, 0x17, 0x9b, 0x7f, 0x52, 0x56, 0x6a, 0x72, 0x89, 0x04,
	0x8b, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xc4, 0x42, 0x3d, 0x98, 0x85, 0x7a, 0x8e, 0x79,
	0x95, 0x41, 0x50, 0x35, 0x49, 0x6c, 0x60, 0x51, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x0b, 0xae, 0xe3, 0xae, 0x00, 0x00, 0x00,
}
