// Code generated by protoc-gen-go.
// source: proto/trade_signal_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TradeSignalResponse struct {
	Ok      bool   `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	OrderId string `protobuf:"bytes,3,opt,name=OrderId" json:"OrderId,omitempty"`
}

func (m *TradeSignalResponse) Reset()                    { *m = TradeSignalResponse{} }
func (m *TradeSignalResponse) String() string            { return proto1.CompactTextString(m) }
func (*TradeSignalResponse) ProtoMessage()               {}
func (*TradeSignalResponse) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *TradeSignalResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TradeSignalResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TradeSignalResponse) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func init() {
	proto1.RegisterType((*TradeSignalResponse)(nil), "proto.TradeSignalResponse")
}

func init() { proto1.RegisterFile("proto/trade_signal_response.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x29, 0x4a, 0x4c, 0x49, 0x8d, 0x2f, 0xce, 0x4c, 0xcf, 0x4b, 0xcc, 0x89, 0x2f,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0xcb, 0x09, 0xb1, 0x82, 0x29, 0xa5, 0x48,
	0x2e, 0xe1, 0x10, 0x90, 0xaa, 0x60, 0xb0, 0xa2, 0x20, 0xa8, 0x1a, 0x21, 0x3e, 0x2e, 0x26, 0xff,
	0x6c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x26, 0xff, 0x6c, 0x21, 0x09, 0x2e, 0x76, 0xdf,
	0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x17, 0x24,
	0xe3, 0x5f, 0x94, 0x92, 0x5a, 0xe4, 0x99, 0x22, 0xc1, 0x0c, 0x91, 0x81, 0x72, 0x93, 0xd8, 0xc0,
	0x36, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0x52, 0x7f, 0x9e, 0x8d, 0x00, 0x00, 0x00,
}
