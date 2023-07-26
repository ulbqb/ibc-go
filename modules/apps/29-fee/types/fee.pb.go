// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/fee/v1/fee.proto

package types

import (
	fmt "fmt"
	github_com_Finschia_finschia_sdk_types "github.com/Finschia/finschia-sdk/types"
	types "github.com/Finschia/finschia-sdk/types"
	types1 "github.com/Finschia/ibc-go/v4/modules/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Fee defines the ICS29 receive, acknowledgement and timeout fees
type Fee struct {
	// the packet receive fee
	RecvFee github_com_Finschia_finschia_sdk_types.Coins `protobuf:"bytes,1,rep,name=recv_fee,json=recvFee,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coins" json:"recv_fee" yaml:"recv_fee"`
	// the packet acknowledgement fee
	AckFee github_com_Finschia_finschia_sdk_types.Coins `protobuf:"bytes,2,rep,name=ack_fee,json=ackFee,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coins" json:"ack_fee" yaml:"ack_fee"`
	// the packet timeout fee
	TimeoutFee github_com_Finschia_finschia_sdk_types.Coins `protobuf:"bytes,3,rep,name=timeout_fee,json=timeoutFee,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coins" json:"timeout_fee" yaml:"timeout_fee"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{0}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func (m *Fee) GetRecvFee() github_com_Finschia_finschia_sdk_types.Coins {
	if m != nil {
		return m.RecvFee
	}
	return nil
}

func (m *Fee) GetAckFee() github_com_Finschia_finschia_sdk_types.Coins {
	if m != nil {
		return m.AckFee
	}
	return nil
}

func (m *Fee) GetTimeoutFee() github_com_Finschia_finschia_sdk_types.Coins {
	if m != nil {
		return m.TimeoutFee
	}
	return nil
}

// PacketFee contains ICS29 relayer fees, refund address and optional list of permitted relayers
type PacketFee struct {
	// fee encapsulates the recv, ack and timeout fees associated with an IBC packet
	Fee Fee `protobuf:"bytes,1,opt,name=fee,proto3" json:"fee"`
	// the refund address for unspent fees
	RefundAddress string `protobuf:"bytes,2,opt,name=refund_address,json=refundAddress,proto3" json:"refund_address,omitempty" yaml:"refund_address"`
	// optional list of relayers permitted to receive fees
	Relayers []string `protobuf:"bytes,3,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *PacketFee) Reset()         { *m = PacketFee{} }
func (m *PacketFee) String() string { return proto.CompactTextString(m) }
func (*PacketFee) ProtoMessage()    {}
func (*PacketFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{1}
}
func (m *PacketFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketFee.Merge(m, src)
}
func (m *PacketFee) XXX_Size() int {
	return m.Size()
}
func (m *PacketFee) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketFee.DiscardUnknown(m)
}

var xxx_messageInfo_PacketFee proto.InternalMessageInfo

func (m *PacketFee) GetFee() Fee {
	if m != nil {
		return m.Fee
	}
	return Fee{}
}

func (m *PacketFee) GetRefundAddress() string {
	if m != nil {
		return m.RefundAddress
	}
	return ""
}

func (m *PacketFee) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

// PacketFees contains a list of type PacketFee
type PacketFees struct {
	// list of packet fees
	PacketFees []PacketFee `protobuf:"bytes,1,rep,name=packet_fees,json=packetFees,proto3" json:"packet_fees" yaml:"packet_fees"`
}

func (m *PacketFees) Reset()         { *m = PacketFees{} }
func (m *PacketFees) String() string { return proto.CompactTextString(m) }
func (*PacketFees) ProtoMessage()    {}
func (*PacketFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{2}
}
func (m *PacketFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketFees.Merge(m, src)
}
func (m *PacketFees) XXX_Size() int {
	return m.Size()
}
func (m *PacketFees) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketFees.DiscardUnknown(m)
}

var xxx_messageInfo_PacketFees proto.InternalMessageInfo

func (m *PacketFees) GetPacketFees() []PacketFee {
	if m != nil {
		return m.PacketFees
	}
	return nil
}

// IdentifiedPacketFees contains a list of type PacketFee and associated PacketId
type IdentifiedPacketFees struct {
	// unique packet identifier comprised of the channel ID, port ID and sequence
	PacketId types1.PacketId `protobuf:"bytes,1,opt,name=packet_id,json=packetId,proto3" json:"packet_id" yaml:"packet_id"`
	// list of packet fees
	PacketFees []PacketFee `protobuf:"bytes,2,rep,name=packet_fees,json=packetFees,proto3" json:"packet_fees" yaml:"packet_fees"`
}

func (m *IdentifiedPacketFees) Reset()         { *m = IdentifiedPacketFees{} }
func (m *IdentifiedPacketFees) String() string { return proto.CompactTextString(m) }
func (*IdentifiedPacketFees) ProtoMessage()    {}
func (*IdentifiedPacketFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3319f1af2a53e5, []int{3}
}
func (m *IdentifiedPacketFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedPacketFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedPacketFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedPacketFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedPacketFees.Merge(m, src)
}
func (m *IdentifiedPacketFees) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedPacketFees) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedPacketFees.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedPacketFees proto.InternalMessageInfo

func (m *IdentifiedPacketFees) GetPacketId() types1.PacketId {
	if m != nil {
		return m.PacketId
	}
	return types1.PacketId{}
}

func (m *IdentifiedPacketFees) GetPacketFees() []PacketFee {
	if m != nil {
		return m.PacketFees
	}
	return nil
}

func init() {
	proto.RegisterType((*Fee)(nil), "ibc.applications.fee.v1.Fee")
	proto.RegisterType((*PacketFee)(nil), "ibc.applications.fee.v1.PacketFee")
	proto.RegisterType((*PacketFees)(nil), "ibc.applications.fee.v1.PacketFees")
	proto.RegisterType((*IdentifiedPacketFees)(nil), "ibc.applications.fee.v1.IdentifiedPacketFees")
}

func init() { proto.RegisterFile("ibc/applications/fee/v1/fee.proto", fileDescriptor_cb3319f1af2a53e5) }

var fileDescriptor_cb3319f1af2a53e5 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0x3b, 0x5b, 0xd9, 0xdd, 0xa6, 0xb8, 0xca, 0xb0, 0x62, 0xb7, 0xe8, 0x74, 0x9d, 0x53,
	0x0f, 0x36, 0xa1, 0x75, 0x41, 0xf4, 0xa4, 0x23, 0x54, 0xd6, 0x93, 0x0e, 0x9e, 0xbc, 0x94, 0x4c,
	0xf2, 0xda, 0x86, 0x76, 0x26, 0xc3, 0x64, 0x5a, 0xa8, 0xe0, 0x55, 0xf0, 0xe6, 0x37, 0xf0, 0xee,
	0x27, 0xd9, 0x8b, 0xb0, 0x47, 0x4f, 0x55, 0xda, 0x6f, 0xb0, 0x9f, 0x40, 0x92, 0x49, 0xcb, 0xac,
	0xb8, 0x2c, 0x0b, 0x9e, 0xe6, 0xbd, 0xe4, 0xfd, 0xf3, 0xfb, 0x27, 0x79, 0x13, 0xf4, 0x48, 0x44,
	0x8c, 0xd0, 0x34, 0x9d, 0x0a, 0x46, 0x73, 0x21, 0x13, 0x45, 0x86, 0x00, 0x64, 0xde, 0xd5, 0x1f,
	0x9c, 0x66, 0x32, 0x97, 0xee, 0x7d, 0x11, 0x31, 0x5c, 0x2e, 0xc1, 0x7a, 0x6e, 0xde, 0x6d, 0x7a,
	0x4c, 0xaa, 0x58, 0x2a, 0x12, 0x51, 0xa5, 0x25, 0x11, 0xe4, 0xb4, 0x4b, 0x98, 0x14, 0x49, 0x21,
	0x6c, 0x1e, 0x8e, 0xe4, 0x48, 0x9a, 0x90, 0xe8, 0xc8, 0x8e, 0x1a, 0x22, 0x93, 0x19, 0x10, 0x36,
	0xa6, 0x49, 0x02, 0x53, 0x4d, 0xb3, 0x61, 0x51, 0xe2, 0x7f, 0xa9, 0xa2, 0x6a, 0x1f, 0xc0, 0xfd,
	0x84, 0xf6, 0x33, 0x60, 0xf3, 0xc1, 0x10, 0xa0, 0xe1, 0x1c, 0x57, 0xdb, 0xf5, 0xde, 0x11, 0x2e,
	0x98, 0x58, 0x33, 0xb1, 0x65, 0xe2, 0x57, 0x52, 0x24, 0xc1, 0xeb, 0xb3, 0x65, 0xab, 0x72, 0xb1,
	0x6c, 0xdd, 0x59, 0xd0, 0x78, 0xfa, 0xdc, 0xdf, 0x08, 0xfd, 0xef, 0xbf, 0x5a, 0x8f, 0x47, 0x22,
	0x1f, 0xcf, 0x22, 0xcc, 0x64, 0x4c, 0xfa, 0x22, 0x51, 0x6c, 0x2c, 0x28, 0x19, 0xda, 0xa0, 0xa3,
	0xf8, 0x84, 0xe4, 0x8b, 0x14, 0x94, 0x59, 0x47, 0x85, 0x7b, 0x5a, 0xaa, 0xf1, 0x1f, 0xd1, 0x1e,
	0x65, 0x13, 0x43, 0xdf, 0xb9, 0x8e, 0xde, 0xb7, 0xf4, 0x83, 0x82, 0x6e, 0x75, 0x37, 0x87, 0xef,
	0x52, 0x36, 0xd1, 0xec, 0xcf, 0x0e, 0xaa, 0xe7, 0x22, 0x06, 0x39, 0xcb, 0x8d, 0x81, 0xea, 0x75,
	0x06, 0xde, 0x58, 0x03, 0x6e, 0x61, 0xa0, 0xa4, 0xbd, 0xb9, 0x09, 0x64, 0xd5, 0x7d, 0x00, 0xff,
	0x9b, 0x83, 0x6a, 0x6f, 0x29, 0x9b, 0x80, 0xce, 0xdc, 0x13, 0x54, 0x2d, 0x2e, 0xc3, 0x69, 0xd7,
	0x7b, 0x0f, 0xf0, 0x15, 0x9d, 0x81, 0xfb, 0x00, 0xc1, 0x2d, 0x6d, 0x28, 0xd4, 0xe5, 0xee, 0x0b,
	0x74, 0x90, 0xc1, 0x70, 0x96, 0xf0, 0x01, 0xe5, 0x3c, 0x03, 0xa5, 0x1a, 0x3b, 0xc7, 0x4e, 0xbb,
	0x16, 0x1c, 0x5d, 0x2c, 0x5b, 0xf7, 0x36, 0xd7, 0x55, 0x9e, 0xf7, 0xc3, 0xdb, 0xc5, 0xc0, 0xcb,
	0x22, 0x77, 0x9b, 0xba, 0x13, 0xa6, 0x74, 0x01, 0x99, 0x32, 0x47, 0x51, 0x0b, 0xb7, 0xb9, 0x1f,
	0x23, 0xb4, 0x35, 0xa8, 0xdc, 0x01, 0xaa, 0xa7, 0x26, 0xd3, 0x5b, 0x57, 0xb6, 0x6d, 0xfc, 0x2b,
	0x9d, 0x6e, 0x95, 0x41, 0xf3, 0xf2, 0x01, 0x96, 0x16, 0xf1, 0x43, 0x94, 0x6e, 0x01, 0xfe, 0x0f,
	0x07, 0x1d, 0x9e, 0x72, 0x48, 0x72, 0x31, 0x14, 0xc0, 0x4b, 0xe4, 0xf7, 0xa8, 0x66, 0x45, 0x82,
	0xdb, 0x13, 0x7a, 0x68, 0xb8, 0xba, 0xd9, 0xf1, 0xa6, 0xc3, 0xb7, 0xcc, 0x53, 0x1e, 0x34, 0x2c,
	0xf2, 0xee, 0x25, 0xa4, 0xe0, 0x7e, 0xb8, 0x9f, 0xda, 0x9a, 0xbf, 0xf7, 0xb3, 0xf3, 0xbf, 0xf7,
	0x13, 0xbc, 0x3b, 0x5b, 0x79, 0xce, 0xf9, 0xca, 0x73, 0x7e, 0xaf, 0x3c, 0xe7, 0xeb, 0xda, 0xab,
	0x9c, 0xaf, 0xbd, 0xca, 0xcf, 0xb5, 0x57, 0xf9, 0xf0, 0xf4, 0x5f, 0x4d, 0x23, 0x22, 0xd6, 0x19,
	0x49, 0x32, 0x3f, 0x21, 0xb1, 0xe4, 0xb3, 0x29, 0x28, 0xfd, 0x7a, 0x28, 0xd2, 0x7b, 0xd6, 0xd1,
	0x0f, 0x87, 0xe9, 0xa2, 0x68, 0xd7, 0xfc, 0xc6, 0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0xbf, 0x56, 0x73, 0x5d, 0x04, 0x00, 0x00,
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TimeoutFee) > 0 {
		for iNdEx := len(m.TimeoutFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeoutFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AckFee) > 0 {
		for iNdEx := len(m.AckFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RecvFee) > 0 {
		for iNdEx := len(m.RecvFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecvFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PacketFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintFee(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RefundAddress) > 0 {
		i -= len(m.RefundAddress)
		copy(dAtA[i:], m.RefundAddress)
		i = encodeVarintFee(dAtA, i, uint64(len(m.RefundAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PacketFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PacketFees) > 0 {
		for iNdEx := len(m.PacketFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PacketFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedPacketFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedPacketFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedPacketFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PacketFees) > 0 {
		for iNdEx := len(m.PacketFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PacketFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.PacketId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RecvFee) > 0 {
		for _, e := range m.RecvFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.AckFee) > 0 {
		for _, e := range m.AckFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	if len(m.TimeoutFee) > 0 {
		for _, e := range m.TimeoutFee {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *PacketFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fee.Size()
	n += 1 + l + sovFee(uint64(l))
	l = len(m.RefundAddress)
	if l > 0 {
		n += 1 + l + sovFee(uint64(l))
	}
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *PacketFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PacketFees) > 0 {
		for _, e := range m.PacketFees {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func (m *IdentifiedPacketFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PacketId.Size()
	n += 1 + l + sovFee(uint64(l))
	if len(m.PacketFees) > 0 {
		for _, e := range m.PacketFees {
			l = e.Size()
			n += 1 + l + sovFee(uint64(l))
		}
	}
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecvFee = append(m.RecvFee, types.Coin{})
			if err := m.RecvFee[len(m.RecvFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckFee = append(m.AckFee, types.Coin{})
			if err := m.AckFee[len(m.AckFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutFee = append(m.TimeoutFee, types.Coin{})
			if err := m.TimeoutFee[len(m.TimeoutFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PacketFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PacketFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefundAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefundAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PacketFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PacketFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketFees = append(m.PacketFees, PacketFee{})
			if err := m.PacketFees[len(m.PacketFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IdentifiedPacketFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IdentifiedPacketFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedPacketFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketFees = append(m.PacketFees, PacketFee{})
			if err := m.PacketFees[len(m.PacketFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFee
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFee
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)
