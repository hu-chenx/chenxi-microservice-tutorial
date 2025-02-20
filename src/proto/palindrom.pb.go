// Code generated by protoc-gen-go. DO NOT EDIT.
// source: palindrom.proto

package palindrom

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

type PalindromeRequest struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PalindromeRequest) Reset()         { *m = PalindromeRequest{} }
func (m *PalindromeRequest) String() string { return proto.CompactTextString(m) }
func (*PalindromeRequest) ProtoMessage()    {}
func (*PalindromeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11459a9a22e6cc0b, []int{0}
}

func (m *PalindromeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PalindromeRequest.Unmarshal(m, b)
}
func (m *PalindromeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PalindromeRequest.Marshal(b, m, deterministic)
}
func (m *PalindromeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PalindromeRequest.Merge(m, src)
}
func (m *PalindromeRequest) XXX_Size() int {
	return xxx_messageInfo_PalindromeRequest.Size(m)
}
func (m *PalindromeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PalindromeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PalindromeRequest proto.InternalMessageInfo

func (m *PalindromeRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type PalindromeResponse struct {
	Palindrome           string   `protobuf:"bytes,2,opt,name=palindrome,proto3" json:"palindrome,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PalindromeResponse) Reset()         { *m = PalindromeResponse{} }
func (m *PalindromeResponse) String() string { return proto.CompactTextString(m) }
func (*PalindromeResponse) ProtoMessage()    {}
func (*PalindromeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11459a9a22e6cc0b, []int{1}
}

func (m *PalindromeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PalindromeResponse.Unmarshal(m, b)
}
func (m *PalindromeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PalindromeResponse.Marshal(b, m, deterministic)
}
func (m *PalindromeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PalindromeResponse.Merge(m, src)
}
func (m *PalindromeResponse) XXX_Size() int {
	return xxx_messageInfo_PalindromeResponse.Size(m)
}
func (m *PalindromeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PalindromeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PalindromeResponse proto.InternalMessageInfo

func (m *PalindromeResponse) GetPalindrome() string {
	if m != nil {
		return m.Palindrome
	}
	return ""
}

func init() {
	proto.RegisterType((*PalindromeRequest)(nil), "PalindromeRequest")
	proto.RegisterType((*PalindromeResponse)(nil), "PalindromeResponse")
}

func init() { proto.RegisterFile("palindrom.proto", fileDescriptor_11459a9a22e6cc0b) }

var fileDescriptor_11459a9a22e6cc0b = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0xcc, 0xc9,
	0xcc, 0x4b, 0x29, 0xca, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe4, 0x12, 0x0c,
	0x80, 0x09, 0xa5, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe6,
	0x15, 0x94, 0x96, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x26, 0x5c, 0x42,
	0xc8, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xe4, 0xb8, 0xb8, 0xe0, 0x66, 0xa6, 0x4a,
	0x30, 0x81, 0x35, 0x20, 0x89, 0x18, 0x79, 0x72, 0x71, 0x21, 0x74, 0x09, 0x59, 0x73, 0xf1, 0xa5,
	0x65, 0xe6, 0xa5, 0x20, 0x89, 0x08, 0xe9, 0x61, 0xd8, 0x2f, 0x25, 0xac, 0x87, 0x69, 0x91, 0x12,
	0x43, 0x12, 0x1b, 0xd8, 0xc9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x21, 0x70, 0x80,
	0xc5, 0x00, 0x00, 0x00,
}
