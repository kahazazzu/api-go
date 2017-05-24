// Code generated by protoc-gen-go.
// source: proto/signup_form.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/signup_form.proto
	proto/schedule_create_response.proto
	proto/schedule_list.proto
	proto/schedule_list_response.proto
	proto/schedule_create.proto
	proto/schedule.proto
	proto/session_create_response.proto
	proto/account.proto
	proto/signup_form_response.proto
	proto/algorithm.proto
	proto/signin_email_response.proto
	proto/session_create.proto
	proto/signin_email.proto

It has these top-level messages:
	SignupForm
	ScheduleCreateResponse
	ScheduleList
	ScheduleListResponse
	ScheduleCreate
	Schedule
	SessionCreateResponse
	Account
	SignupFormResponse
	Algorithm
	SigninEmailResponse
	SessionCreate
	SigninEmail
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SignupForm struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Thing   string   `protobuf:"bytes,2,opt,name=thing" json:"thing,omitempty"`
}

func (m *SignupForm) Reset()                    { *m = SignupForm{} }
func (m *SignupForm) String() string            { return proto1.CompactTextString(m) }
func (*SignupForm) ProtoMessage()               {}
func (*SignupForm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignupForm) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *SignupForm) GetThing() string {
	if m != nil {
		return m.Thing
	}
	return ""
}

func init() {
	proto1.RegisterType((*SignupForm)(nil), "proto.SignupForm")
}

func init() { proto1.RegisterFile("proto/signup_form.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0x4c, 0xcf, 0x2b, 0x2d, 0x88, 0x4f, 0xcb, 0x2f, 0xca, 0xd5, 0x03, 0x8b,
	0x08, 0xb1, 0x82, 0x29, 0x29, 0x61, 0x88, 0x7c, 0x62, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x09, 0x44,
	0x4e, 0xc9, 0x87, 0x8b, 0x2b, 0x18, 0xac, 0xc1, 0x2d, 0xbf, 0x28, 0x57, 0x48, 0x83, 0x8b, 0x1d,
	0x2a, 0x2d, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0x07, 0x51, 0xa6, 0xe7, 0x08, 0x11, 0x0d,
	0x82, 0x49, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x64, 0xe6, 0xa5, 0x4b, 0x30, 0x29, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x38, 0x49, 0x6c, 0x60, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef,
	0xe2, 0x6a, 0x10, 0x8b, 0x00, 0x00, 0x00,
}
