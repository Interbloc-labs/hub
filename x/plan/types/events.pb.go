// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/plan/v2/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sentinel-official/hub/types"
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

type EventCreate struct {
	ID              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	ProviderAddress string `protobuf:"bytes,2,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty" yaml:"provider_address"`
}

func (m *EventCreate) Reset()         { *m = EventCreate{} }
func (m *EventCreate) String() string { return proto.CompactTextString(m) }
func (*EventCreate) ProtoMessage()    {}
func (*EventCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{0}
}
func (m *EventCreate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreate.Merge(m, src)
}
func (m *EventCreate) XXX_Size() int {
	return m.Size()
}
func (m *EventCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreate.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreate proto.InternalMessageInfo

type EventSetStatus struct {
	ID              uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	ProviderAddress string       `protobuf:"bytes,2,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty" yaml:"provider_address"`
	Status          types.Status `protobuf:"varint,3,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
}

func (m *EventSetStatus) Reset()         { *m = EventSetStatus{} }
func (m *EventSetStatus) String() string { return proto.CompactTextString(m) }
func (*EventSetStatus) ProtoMessage()    {}
func (*EventSetStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{1}
}
func (m *EventSetStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSetStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSetStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSetStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSetStatus.Merge(m, src)
}
func (m *EventSetStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventSetStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSetStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventSetStatus proto.InternalMessageInfo

type EventLinkNode struct {
	ID              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	NodeAddress     string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ProviderAddress string `protobuf:"bytes,3,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty" yaml:"provider_address"`
}

func (m *EventLinkNode) Reset()         { *m = EventLinkNode{} }
func (m *EventLinkNode) String() string { return proto.CompactTextString(m) }
func (*EventLinkNode) ProtoMessage()    {}
func (*EventLinkNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{2}
}
func (m *EventLinkNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventLinkNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventLinkNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventLinkNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLinkNode.Merge(m, src)
}
func (m *EventLinkNode) XXX_Size() int {
	return m.Size()
}
func (m *EventLinkNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLinkNode.DiscardUnknown(m)
}

var xxx_messageInfo_EventLinkNode proto.InternalMessageInfo

type EventUnlinkNode struct {
	ID              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	NodeAddress     string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ProviderAddress string `protobuf:"bytes,3,opt,name=provider_address,json=providerAddress,proto3" json:"provider_address,omitempty" yaml:"provider_address"`
}

func (m *EventUnlinkNode) Reset()         { *m = EventUnlinkNode{} }
func (m *EventUnlinkNode) String() string { return proto.CompactTextString(m) }
func (*EventUnlinkNode) ProtoMessage()    {}
func (*EventUnlinkNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4222ab50472303c2, []int{3}
}
func (m *EventUnlinkNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUnlinkNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUnlinkNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUnlinkNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUnlinkNode.Merge(m, src)
}
func (m *EventUnlinkNode) XXX_Size() int {
	return m.Size()
}
func (m *EventUnlinkNode) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUnlinkNode.DiscardUnknown(m)
}

var xxx_messageInfo_EventUnlinkNode proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventCreate)(nil), "sentinel.plan.v2.EventCreate")
	proto.RegisterType((*EventSetStatus)(nil), "sentinel.plan.v2.EventSetStatus")
	proto.RegisterType((*EventLinkNode)(nil), "sentinel.plan.v2.EventLinkNode")
	proto.RegisterType((*EventUnlinkNode)(nil), "sentinel.plan.v2.EventUnlinkNode")
}

func init() { proto.RegisterFile("sentinel/plan/v2/events.proto", fileDescriptor_4222ab50472303c2) }

var fileDescriptor_4222ab50472303c2 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4b, 0xcd, 0xd1, 0x2f, 0xc8, 0x49, 0xcc, 0xd3, 0x2f, 0x33, 0xd2, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x49, 0xeb, 0x81, 0xa4,
	0xf5, 0xca, 0x8c, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d,
	0x94, 0x1c, 0xdc, 0x98, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0x32, 0x43, 0xfd, 0xe2, 0x92, 0xc4,
	0x92, 0x52, 0xa8, 0x39, 0x4a, 0x55, 0x5c, 0xdc, 0xae, 0x20, 0x73, 0x9d, 0x8b, 0x52, 0x13, 0x4b,
	0x52, 0x85, 0x94, 0xb9, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x9c, 0x84, 0x1f,
	0xdd, 0x93, 0x67, 0xf2, 0x74, 0xf9, 0x74, 0x4f, 0x9e, 0xb3, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x29,
	0x33, 0x45, 0x29, 0x88, 0x29, 0x33, 0x45, 0xc8, 0x8d, 0x4b, 0xa0, 0xa0, 0x28, 0xbf, 0x2c, 0x33,
	0x25, 0xb5, 0x28, 0x3e, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83,
	0xd3, 0x49, 0xfa, 0xd3, 0x3d, 0x79, 0x71, 0x88, 0x62, 0x74, 0x15, 0x4a, 0x41, 0xfc, 0x30, 0x21,
	0x47, 0xa8, 0xc8, 0x61, 0x46, 0x2e, 0x3e, 0xb0, 0xe5, 0xc1, 0xa9, 0x25, 0xc1, 0x60, 0x47, 0xd1,
	0xd5, 0x7e, 0x21, 0x17, 0x2e, 0x36, 0x48, 0x58, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x19, 0x49,
	0xea, 0xc1, 0x03, 0x15, 0x1c, 0x58, 0x7a, 0x65, 0x86, 0x7a, 0x10, 0x77, 0x39, 0x09, 0x7e, 0xba,
	0x27, 0xcf, 0x0b, 0x31, 0x18, 0xa2, 0x45, 0x29, 0x08, 0xaa, 0x57, 0x69, 0x07, 0x23, 0x17, 0x2f,
	0xd8, 0x17, 0x3e, 0x99, 0x79, 0xd9, 0x7e, 0xf9, 0x29, 0x44, 0x06, 0xa2, 0x15, 0x17, 0x4f, 0x5e,
	0x7e, 0x4a, 0x2a, 0x9a, 0x07, 0xc4, 0x3f, 0xdd, 0x93, 0x17, 0x86, 0x28, 0x44, 0x96, 0x55, 0x0a,
	0xe2, 0x06, 0x71, 0x61, 0x0e, 0xc7, 0x16, 0x00, 0xcc, 0x64, 0x44, 0xc0, 0x2e, 0x46, 0x2e, 0x7e,
	0xb0, 0xd3, 0x43, 0xf3, 0x72, 0x86, 0x9a, 0xe3, 0x9d, 0xfc, 0x4f, 0x3c, 0x94, 0x63, 0x58, 0xf1,
	0x48, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x61, 0x31, 0xab, 0x9b, 0x9f, 0x96, 0x96,
	0x99, 0x9c, 0x99, 0x98, 0xa3, 0x9f, 0x51, 0x9a, 0xa4, 0x5f, 0x01, 0xc9, 0x5c, 0xe0, 0xd8, 0x4e,
	0x62, 0x03, 0xe7, 0x08, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x7e, 0x49, 0xbc, 0x7a,
	0x03, 0x00, 0x00,
}

func (m *EventCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProviderAddress) > 0 {
		i -= len(m.ProviderAddress)
		copy(dAtA[i:], m.ProviderAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ProviderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventSetStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSetStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSetStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ProviderAddress) > 0 {
		i -= len(m.ProviderAddress)
		copy(dAtA[i:], m.ProviderAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ProviderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventLinkNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventLinkNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventLinkNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProviderAddress) > 0 {
		i -= len(m.ProviderAddress)
		copy(dAtA[i:], m.ProviderAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ProviderAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventUnlinkNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUnlinkNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUnlinkNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProviderAddress) > 0 {
		i -= len(m.ProviderAddress)
		copy(dAtA[i:], m.ProviderAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ProviderAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	l = len(m.ProviderAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSetStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	l = len(m.ProviderAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventLinkNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ProviderAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUnlinkNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ProviderAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSetStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventLinkNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventLinkNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventLinkNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUnlinkNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUnlinkNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUnlinkNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
