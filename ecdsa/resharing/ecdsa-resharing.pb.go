// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob/ecdsa-resharing.proto

package resharing

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

//
// The Round 1 data is broadcast to peers of the New Committee in this message.
type DGRound1Message struct {
	EcdsaPubX            []byte   `protobuf:"bytes,1,opt,name=ecdsa_pub_x,json=ecdsaPubX,proto3" json:"ecdsa_pub_x,omitempty"`
	EcdsaPubY            []byte   `protobuf:"bytes,2,opt,name=ecdsa_pub_y,json=ecdsaPubY,proto3" json:"ecdsa_pub_y,omitempty"`
	VCommitment          []byte   `protobuf:"bytes,3,opt,name=v_commitment,json=vCommitment,proto3" json:"v_commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound1Message) Reset()         { *m = DGRound1Message{} }
func (m *DGRound1Message) String() string { return proto.CompactTextString(m) }
func (*DGRound1Message) ProtoMessage()    {}
func (*DGRound1Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7d3ae1dc68dc295, []int{0}
}

func (m *DGRound1Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound1Message.Unmarshal(m, b)
}
func (m *DGRound1Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound1Message.Marshal(b, m, deterministic)
}
func (m *DGRound1Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound1Message.Merge(m, src)
}
func (m *DGRound1Message) XXX_Size() int {
	return xxx_messageInfo_DGRound1Message.Size(m)
}
func (m *DGRound1Message) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound1Message.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound1Message proto.InternalMessageInfo

func (m *DGRound1Message) GetEcdsaPubX() []byte {
	if m != nil {
		return m.EcdsaPubX
	}
	return nil
}

func (m *DGRound1Message) GetEcdsaPubY() []byte {
	if m != nil {
		return m.EcdsaPubY
	}
	return nil
}

func (m *DGRound1Message) GetVCommitment() []byte {
	if m != nil {
		return m.VCommitment
	}
	return nil
}

//
// The Round 2 data is broadcast to other peers of the New Committee in this message.
type DGRound2Message1 struct {
	PaillierN            []byte   `protobuf:"bytes,1,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	PaillierProof        [][]byte `protobuf:"bytes,2,rep,name=paillier_proof,json=paillierProof,proto3" json:"paillier_proof,omitempty"`
	NTilde               []byte   `protobuf:"bytes,3,opt,name=n_tilde,json=nTilde,proto3" json:"n_tilde,omitempty"`
	H1                   []byte   `protobuf:"bytes,4,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   []byte   `protobuf:"bytes,5,opt,name=h2,proto3" json:"h2,omitempty"`
	Dlnproof_1           [][]byte `protobuf:"bytes,6,rep,name=dlnproof_1,json=dlnproof1,proto3" json:"dlnproof_1,omitempty"`
	Dlnproof_2           [][]byte `protobuf:"bytes,7,rep,name=dlnproof_2,json=dlnproof2,proto3" json:"dlnproof_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound2Message1) Reset()         { *m = DGRound2Message1{} }
func (m *DGRound2Message1) String() string { return proto.CompactTextString(m) }
func (*DGRound2Message1) ProtoMessage()    {}
func (*DGRound2Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7d3ae1dc68dc295, []int{1}
}

func (m *DGRound2Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound2Message1.Unmarshal(m, b)
}
func (m *DGRound2Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound2Message1.Marshal(b, m, deterministic)
}
func (m *DGRound2Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound2Message1.Merge(m, src)
}
func (m *DGRound2Message1) XXX_Size() int {
	return xxx_messageInfo_DGRound2Message1.Size(m)
}
func (m *DGRound2Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound2Message1.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound2Message1 proto.InternalMessageInfo

func (m *DGRound2Message1) GetPaillierN() []byte {
	if m != nil {
		return m.PaillierN
	}
	return nil
}

func (m *DGRound2Message1) GetPaillierProof() [][]byte {
	if m != nil {
		return m.PaillierProof
	}
	return nil
}

func (m *DGRound2Message1) GetNTilde() []byte {
	if m != nil {
		return m.NTilde
	}
	return nil
}

func (m *DGRound2Message1) GetH1() []byte {
	if m != nil {
		return m.H1
	}
	return nil
}

func (m *DGRound2Message1) GetH2() []byte {
	if m != nil {
		return m.H2
	}
	return nil
}

func (m *DGRound2Message1) GetDlnproof_1() [][]byte {
	if m != nil {
		return m.Dlnproof_1
	}
	return nil
}

func (m *DGRound2Message1) GetDlnproof_2() [][]byte {
	if m != nil {
		return m.Dlnproof_2
	}
	return nil
}

//
// The Round 2 "ACK" is broadcast to peers of the Old Committee in this message.
type DGRound2Message2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound2Message2) Reset()         { *m = DGRound2Message2{} }
func (m *DGRound2Message2) String() string { return proto.CompactTextString(m) }
func (*DGRound2Message2) ProtoMessage()    {}
func (*DGRound2Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7d3ae1dc68dc295, []int{2}
}

func (m *DGRound2Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound2Message2.Unmarshal(m, b)
}
func (m *DGRound2Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound2Message2.Marshal(b, m, deterministic)
}
func (m *DGRound2Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound2Message2.Merge(m, src)
}
func (m *DGRound2Message2) XXX_Size() int {
	return xxx_messageInfo_DGRound2Message2.Size(m)
}
func (m *DGRound2Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound2Message2.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound2Message2 proto.InternalMessageInfo

//
// The Round 3 data is sent to peers of the New Committee in this message.
type DGRound3Message1 struct {
	Share                []byte   `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound3Message1) Reset()         { *m = DGRound3Message1{} }
func (m *DGRound3Message1) String() string { return proto.CompactTextString(m) }
func (*DGRound3Message1) ProtoMessage()    {}
func (*DGRound3Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7d3ae1dc68dc295, []int{3}
}

func (m *DGRound3Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound3Message1.Unmarshal(m, b)
}
func (m *DGRound3Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound3Message1.Marshal(b, m, deterministic)
}
func (m *DGRound3Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound3Message1.Merge(m, src)
}
func (m *DGRound3Message1) XXX_Size() int {
	return xxx_messageInfo_DGRound3Message1.Size(m)
}
func (m *DGRound3Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound3Message1.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound3Message1 proto.InternalMessageInfo

func (m *DGRound3Message1) GetShare() []byte {
	if m != nil {
		return m.Share
	}
	return nil
}

//
// The Round 3 data is broadcast to peers of the New Committee in this message.
type DGRound3Message2 struct {
	VDecommitment        [][]byte `protobuf:"bytes,1,rep,name=v_decommitment,json=vDecommitment,proto3" json:"v_decommitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound3Message2) Reset()         { *m = DGRound3Message2{} }
func (m *DGRound3Message2) String() string { return proto.CompactTextString(m) }
func (*DGRound3Message2) ProtoMessage()    {}
func (*DGRound3Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7d3ae1dc68dc295, []int{4}
}

func (m *DGRound3Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound3Message2.Unmarshal(m, b)
}
func (m *DGRound3Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound3Message2.Marshal(b, m, deterministic)
}
func (m *DGRound3Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound3Message2.Merge(m, src)
}
func (m *DGRound3Message2) XXX_Size() int {
	return xxx_messageInfo_DGRound3Message2.Size(m)
}
func (m *DGRound3Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound3Message2.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound3Message2 proto.InternalMessageInfo

func (m *DGRound3Message2) GetVDecommitment() [][]byte {
	if m != nil {
		return m.VDecommitment
	}
	return nil
}

//
// The Round 4 "ACK" is broadcast to peers of the Old and New Committees from the New Committee in this message.
type DGRound4Message struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound4Message) Reset()         { *m = DGRound4Message{} }
func (m *DGRound4Message) String() string { return proto.CompactTextString(m) }
func (*DGRound4Message) ProtoMessage()    {}
func (*DGRound4Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7d3ae1dc68dc295, []int{5}
}

func (m *DGRound4Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound4Message.Unmarshal(m, b)
}
func (m *DGRound4Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound4Message.Marshal(b, m, deterministic)
}
func (m *DGRound4Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound4Message.Merge(m, src)
}
func (m *DGRound4Message) XXX_Size() int {
	return xxx_messageInfo_DGRound4Message.Size(m)
}
func (m *DGRound4Message) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound4Message.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound4Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DGRound1Message)(nil), "DGRound1Message")
	proto.RegisterType((*DGRound2Message1)(nil), "DGRound2Message1")
	proto.RegisterType((*DGRound2Message2)(nil), "DGRound2Message2")
	proto.RegisterType((*DGRound3Message1)(nil), "DGRound3Message1")
	proto.RegisterType((*DGRound3Message2)(nil), "DGRound3Message2")
	proto.RegisterType((*DGRound4Message)(nil), "DGRound4Message")
}

func init() { proto.RegisterFile("protob/ecdsa-resharing.proto", fileDescriptor_f7d3ae1dc68dc295) }

var fileDescriptor_f7d3ae1dc68dc295 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x69, 0xe7, 0x3a, 0x96, 0xd5, 0xcd, 0x05, 0xc1, 0x1c, 0x54, 0x66, 0x40, 0xe8, 0x45,
	0x47, 0x32, 0x2f, 0x5e, 0x75, 0xe0, 0x49, 0x19, 0xc5, 0x83, 0x7a, 0x09, 0xed, 0x1a, 0xd7, 0x42,
	0x9b, 0x94, 0xfe, 0x42, 0xff, 0x4a, 0xff, 0x25, 0x69, 0x96, 0x86, 0x95, 0x1d, 0xdf, 0xe7, 0x7d,
	0x9b, 0xc7, 0xeb, 0x03, 0x97, 0x79, 0x21, 0x2b, 0x19, 0x2e, 0xf9, 0x36, 0x2a, 0x83, 0xbb, 0x82,
	0x97, 0x71, 0x50, 0x24, 0x62, 0x77, 0xaf, 0x30, 0xae, 0xc0, 0x6c, 0xfd, 0xe2, 0xcb, 0x5a, 0x44,
	0xe4, 0x95, 0x97, 0x65, 0xb0, 0xe3, 0xf0, 0x1a, 0x4c, 0xd4, 0x2d, 0xcb, 0xeb, 0x90, 0xfd, 0x20,
	0x6b, 0x61, 0x79, 0xae, 0x3f, 0x56, 0x68, 0x53, 0x87, 0x1f, 0x7d, 0xff, 0x17, 0xd9, 0x7d, 0xff,
	0x13, 0xde, 0x00, 0xb7, 0x61, 0x5b, 0x99, 0x65, 0x49, 0x95, 0x71, 0x51, 0xa1, 0x81, 0x3a, 0x98,
	0x34, 0xcf, 0x06, 0xe1, 0x3f, 0x0b, 0x9c, 0xe9, 0x58, 0xaa, 0x63, 0x09, 0xbc, 0x02, 0x20, 0x0f,
	0x92, 0x34, 0x4d, 0x78, 0xc1, 0x44, 0x17, 0xdb, 0x91, 0x37, 0x78, 0x0b, 0xa6, 0xc6, 0xce, 0x0b,
	0x29, 0xbf, 0x91, 0xbd, 0x18, 0x78, 0xae, 0x7f, 0xda, 0xd1, 0x4d, 0x0b, 0xe1, 0x05, 0x18, 0x09,
	0x56, 0x25, 0x69, 0xc4, 0x75, 0xb0, 0x23, 0xde, 0x5b, 0x05, 0xa7, 0xc0, 0x8e, 0x09, 0x3a, 0x51,
	0xcc, 0x8e, 0x89, 0xd2, 0x14, 0x0d, 0xb5, 0xa6, 0x6d, 0x7c, 0x94, 0x0a, 0xf5, 0x32, 0x23, 0xc8,
	0x51, 0x6f, 0x8f, 0x3b, 0x42, 0x7a, 0x36, 0x45, 0xa3, 0xbe, 0x4d, 0x31, 0x3c, 0x2a, 0x44, 0xb1,
	0x67, 0xd8, 0xca, 0x94, 0x3c, 0x07, 0xc3, 0x76, 0x00, 0xae, 0xfb, 0xed, 0x05, 0x7e, 0x3c, 0xba,
	0xa4, 0x6d, 0xdf, 0x86, 0x45, 0xfc, 0xe0, 0x47, 0x5a, 0xfb, 0xbe, 0xcd, 0xfa, 0x00, 0xe2, 0xb9,
	0x19, 0xf0, 0x41, 0x7f, 0xfa, 0x34, 0xff, 0x9a, 0xa9, 0x35, 0x96, 0x66, 0xec, 0xd0, 0x51, 0x6b,
	0xaf, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x14, 0x8e, 0x99, 0x0d, 0x02, 0x00, 0x00,
}
