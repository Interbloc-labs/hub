// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v1/genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisSubscription struct {
	Subscription Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"_"`
	Quotas       []Quota      `protobuf:"bytes,2,rep,name=quotas,proto3" json:"quotas"`
}

func (m *GenesisSubscription) Reset()         { *m = GenesisSubscription{} }
func (m *GenesisSubscription) String() string { return proto.CompactTextString(m) }
func (*GenesisSubscription) ProtoMessage()    {}
func (*GenesisSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_81eef651065d2972, []int{0}
}
func (m *GenesisSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSubscription.Merge(m, src)
}
func (m *GenesisSubscription) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSubscription proto.InternalMessageInfo

type GenesisState struct {
	Subscriptions []GenesisSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"_,omitempty"`
	Params        Params                `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_81eef651065d2972, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisSubscription)(nil), "sentinel.subscription.v1.GenesisSubscription")
	proto.RegisterType((*GenesisState)(nil), "sentinel.subscription.v1.GenesisState")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v1/genesis.proto", fileDescriptor_81eef651065d2972)
}

var fileDescriptor_81eef651065d2972 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4a, 0xfb, 0x40,
	0x1c, 0xc7, 0x73, 0xfd, 0xff, 0x29, 0x78, 0xad, 0x4b, 0xea, 0x10, 0x3a, 0x5c, 0x4b, 0xd1, 0x52,
	0xd0, 0xe6, 0x68, 0x9d, 0x55, 0xc8, 0xe2, 0xaa, 0x15, 0x17, 0x97, 0x72, 0x09, 0xd7, 0xf4, 0xa0,
	0xc9, 0xc5, 0xde, 0xa5, 0xd8, 0xb7, 0xf0, 0x31, 0x04, 0x5f, 0xc0, 0x47, 0xe8, 0xd8, 0xd1, 0xa9,
	0x68, 0xb2, 0xf9, 0x14, 0x92, 0x4b, 0x02, 0x09, 0x78, 0x5b, 0x20, 0x9f, 0xef, 0x27, 0xdf, 0x5f,
	0xbe, 0x70, 0x28, 0x68, 0x28, 0x59, 0x48, 0x57, 0x58, 0xc4, 0xae, 0xf0, 0xd6, 0x2c, 0x92, 0x8c,
	0x87, 0x78, 0x33, 0xc1, 0x3e, 0x0d, 0xa9, 0x60, 0xc2, 0x8e, 0xd6, 0x5c, 0x72, 0xd3, 0x2a, 0x39,
	0xbb, 0xca, 0xd9, 0x9b, 0x49, 0xf7, 0xc4, 0xe7, 0x3e, 0x57, 0x10, 0xce, 0x9e, 0x72, 0xbe, 0x7b,
	0xa6, 0xf5, 0x46, 0x64, 0x4d, 0x82, 0x42, 0xdb, 0x3d, 0xd5, 0x62, 0xcf, 0x31, 0x97, 0xa4, 0xa0,
	0xce, 0xb5, 0x54, 0xad, 0x8c, 0x82, 0x07, 0xef, 0x00, 0x76, 0x6e, 0xf3, 0xee, 0x0f, 0x95, 0xb7,
	0xe6, 0x23, 0x6c, 0x57, 0x69, 0x0b, 0xf4, 0xc1, 0xa8, 0x35, 0x1d, 0xda, 0xba, 0xc3, 0xec, 0x6a,
	0xda, 0x39, 0xda, 0x1d, 0x7a, 0xc6, 0xcf, 0xa1, 0x07, 0xe6, 0xb3, 0x9a, 0xc6, 0xbc, 0x82, 0x4d,
	0x55, 0x55, 0x58, 0x8d, 0xfe, 0xbf, 0x51, 0x6b, 0xda, 0xd3, 0x0b, 0xef, 0x33, 0xce, 0xf9, 0x9f,
	0x99, 0x66, 0x45, 0x68, 0xf0, 0x01, 0x60, 0xbb, 0x6c, 0x2b, 0x89, 0xa4, 0xe6, 0x12, 0x1e, 0x57,
	0x73, 0xc2, 0x02, 0x4a, 0x3b, 0xd6, 0x6b, 0xff, 0x38, 0xd6, 0xe9, 0x14, 0x75, 0x5b, 0xf3, 0x0b,
	0x1e, 0x30, 0x49, 0x83, 0x48, 0x6e, 0x67, 0x75, 0xb1, 0x79, 0x0d, 0x9b, 0xf9, 0x16, 0x56, 0x43,
	0xfd, 0x8a, 0xbe, 0xfe, 0x13, 0x77, 0x8a, 0x2b, 0xab, 0xe7, 0x29, 0x87, 0xec, 0xbe, 0x91, 0xf1,
	0x96, 0x20, 0x63, 0x97, 0x20, 0xb0, 0x4f, 0x10, 0xf8, 0x4a, 0x10, 0x78, 0x4d, 0x91, 0xb1, 0x4f,
	0x91, 0xf1, 0x99, 0x22, 0xe3, 0xe9, 0xc6, 0x67, 0x72, 0x19, 0xbb, 0xb6, 0xc7, 0x03, 0x5c, 0xba,
	0xc7, 0x7c, 0xb1, 0x60, 0x1e, 0x23, 0x2b, 0xbc, 0x8c, 0x5d, 0xfc, 0x52, 0x5f, 0x74, 0x45, 0x7d,
	0xe2, 0x6d, 0xb3, 0x61, 0xe5, 0x36, 0xa2, 0xc2, 0x6d, 0xaa, 0x49, 0x2f, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x26, 0x77, 0xda, 0xb9, 0xa6, 0x02, 0x00, 0x00,
}

func (m *GenesisSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Quotas) > 0 {
		for iNdEx := len(m.Quotas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Quotas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Subscription.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Subscriptions) > 0 {
		for iNdEx := len(m.Subscriptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subscriptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Subscription.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Quotas) > 0 {
		for _, e := range m.Quotas {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for _, e := range m.Subscriptions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Subscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quotas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quotas = append(m.Quotas, Quota{})
			if err := m.Quotas[len(m.Quotas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, GenesisSubscription{})
			if err := m.Subscriptions[len(m.Subscriptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
