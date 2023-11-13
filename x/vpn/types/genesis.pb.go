// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/vpn/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/sentinel-official/hub/v1/x/deposit/types"
	types1 "github.com/sentinel-official/hub/v1/x/node/types"
	types2 "github.com/sentinel-official/hub/v1/x/plan/types"
	types3 "github.com/sentinel-official/hub/v1/x/provider/types"
	types4 "github.com/sentinel-official/hub/v1/x/session/types"
	types5 "github.com/sentinel-official/hub/v1/x/subscription/types"
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

// GenesisState represents the initial state of the Sentinel module at genesis.
type GenesisState struct {
	// Field 1: List of deposits associated with the Sentinel module.
	// This field is not nullable.
	Deposits []types.Deposit `protobuf:"bytes,1,rep,name=deposits,proto3" json:"deposits"`
	// Field 2: Genesis state for nodes in the Sentinel module.
	Nodes *types1.GenesisState `protobuf:"bytes,2,opt,name=nodes,proto3" json:"nodes,omitempty"`
	// Field 3: List of plans associated with the Sentinel module.
	// This field is not nullable.
	Plans []types2.GenesisPlan `protobuf:"bytes,3,rep,name=plans,proto3" json:"plans"`
	// Field 4: Genesis state for providers in the Sentinel module.
	Providers *types3.GenesisState `protobuf:"bytes,4,opt,name=providers,proto3" json:"providers,omitempty"`
	// Field 5: Genesis state for sessions in the Sentinel module.
	Sessions *types4.GenesisState `protobuf:"bytes,5,opt,name=sessions,proto3" json:"sessions,omitempty"`
	// Field 6: Genesis state for subscriptions in the Sentinel module.
	Subscriptions *types5.GenesisState `protobuf:"bytes,6,opt,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_255b384044b0bea0, []int{0}
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
	proto.RegisterType((*GenesisState)(nil), "sentinel.vpn.v1.GenesisState")
}

func init() { proto.RegisterFile("sentinel/vpn/v1/genesis.proto", fileDescriptor_255b384044b0bea0) }

var fileDescriptor_255b384044b0bea0 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0x4e, 0xfe, 0x5e, 0xd4, 0xdf, 0x05, 0x21, 0x45, 0x0c, 0x51, 0x45, 0x4d, 0xdb, 0xa1, 0xea,
	0x82, 0xad, 0x06, 0x16, 0x06, 0x10, 0xaa, 0x90, 0x58, 0x18, 0xa0, 0x6c, 0x6c, 0x49, 0xeb, 0xa6,
	0x96, 0x82, 0x1d, 0xc5, 0xae, 0x05, 0x2f, 0x81, 0x78, 0x0c, 0x1e, 0xa5, 0x63, 0x47, 0x26, 0x04,
	0xe9, 0x8b, 0xa0, 0xdc, 0x53, 0xa5, 0xdb, 0xd1, 0xf9, 0x6e, 0x39, 0x5f, 0x0c, 0xba, 0x82, 0x30,
	0x49, 0x19, 0xf1, 0xb0, 0xf2, 0x19, 0x56, 0x63, 0xec, 0x12, 0x46, 0x04, 0x15, 0xc8, 0x0f, 0xb8,
	0xe4, 0xc6, 0x51, 0x06, 0x23, 0xe5, 0x33, 0xa4, 0xc6, 0x9d, 0x63, 0x97, 0xbb, 0x3c, 0xc6, 0x70,
	0x34, 0x25, 0xb4, 0x4e, 0x3f, 0x77, 0x99, 0x13, 0x9f, 0x0b, 0x2a, 0x23, 0xa7, 0x74, 0x4c, 0x29,
	0x30, 0xa7, 0x30, 0x3e, 0x27, 0x58, 0x59, 0xbb, 0x49, 0x25, 0xdc, 0xf7, 0x6c, 0x56, 0xc5, 0x07,
	0x05, 0x1e, 0x70, 0x45, 0xe7, 0x24, 0xa8, 0x72, 0x8a, 0xcf, 0x10, 0x44, 0x08, 0xca, 0xf7, 0xd8,
	0x0c, 0x0b, 0xca, 0xca, 0x11, 0xb3, 0x80, 0xfa, 0x72, 0x1f, 0x6f, 0xf0, 0x5e, 0x03, 0x07, 0x77,
	0xc9, 0xe6, 0x49, 0xda, 0x92, 0x18, 0xd7, 0xa0, 0x95, 0x1e, 0x24, 0x4c, 0xbd, 0x57, 0x1b, 0xb5,
	0xad, 0x13, 0x94, 0x97, 0x93, 0x9d, 0xaa, 0xc6, 0xe8, 0x36, 0x19, 0x27, 0xf5, 0xf5, 0xf7, 0xa9,
	0x36, 0xcd, 0x35, 0xc6, 0x05, 0x68, 0x44, 0x87, 0x0b, 0xf3, 0x5f, 0x4f, 0x1f, 0xb5, 0x2d, 0x58,
	0x88, 0xa3, 0x35, 0x52, 0x16, 0x2a, 0xc7, 0x4d, 0x13, 0xb2, 0x71, 0x09, 0x1a, 0x51, 0x1d, 0xc2,
	0xac, 0xc5, 0x91, 0xdd, 0x42, 0x15, 0xad, 0x4b, 0xaa, 0x07, 0xcf, 0x66, 0x69, 0x66, 0xa2, 0x30,
	0x6e, 0xc0, 0xff, 0xac, 0x29, 0x61, 0xd6, 0xe3, 0xd0, 0x41, 0x49, 0x9e, 0x42, 0x95, 0xe0, 0x42,
	0x64, 0x5c, 0x81, 0x56, 0xda, 0xa3, 0x30, 0x1b, 0xb1, 0x41, 0xbf, 0x30, 0x48, 0x91, 0x8a, 0x3e,
	0x97, 0x18, 0xf7, 0xe0, 0xb0, 0xdc, 0xb1, 0x30, 0x9b, 0xb1, 0xc7, 0xb0, 0xe4, 0x51, 0x82, 0x2b,
	0x46, 0xbb, 0xe2, 0xc9, 0xe3, 0xfa, 0x17, 0x6a, 0x9f, 0x21, 0xd4, 0xd6, 0x21, 0xd4, 0x37, 0x21,
	0xd4, 0x7f, 0x42, 0xa8, 0x7f, 0x6c, 0xa1, 0xb6, 0xd9, 0x42, 0xed, 0x6b, 0x0b, 0xb5, 0x67, 0xec,
	0x52, 0xb9, 0x5c, 0x39, 0x68, 0xc6, 0x5f, 0x70, 0x16, 0x71, 0xc6, 0x17, 0x0b, 0x3a, 0xa3, 0xb6,
	0x87, 0x97, 0x2b, 0x27, 0x7a, 0x94, 0xaf, 0xf1, 0x3b, 0x97, 0x6f, 0x3e, 0x11, 0x4e, 0x33, 0xfe,
	0xd5, 0xe7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0x88, 0x98, 0x72, 0x04, 0x03, 0x00, 0x00,
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
	if m.Subscriptions != nil {
		{
			size, err := m.Subscriptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Sessions != nil {
		{
			size, err := m.Sessions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Providers != nil {
		{
			size, err := m.Providers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Plans) > 0 {
		for iNdEx := len(m.Plans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Plans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Nodes != nil {
		{
			size, err := m.Nodes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Nodes != nil {
		l = m.Nodes.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Plans) > 0 {
		for _, e := range m.Plans {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Providers != nil {
		l = m.Providers.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Sessions != nil {
		l = m.Sessions.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Subscriptions != nil {
		l = m.Subscriptions.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, types.Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
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
			if m.Nodes == nil {
				m.Nodes = &types1.GenesisState{}
			}
			if err := m.Nodes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plans", wireType)
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
			m.Plans = append(m.Plans, types2.GenesisPlan{})
			if err := m.Plans[len(m.Plans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Providers", wireType)
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
			if m.Providers == nil {
				m.Providers = &types3.GenesisState{}
			}
			if err := m.Providers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sessions", wireType)
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
			if m.Sessions == nil {
				m.Sessions = &types4.GenesisState{}
			}
			if err := m.Sessions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
			if m.Subscriptions == nil {
				m.Subscriptions = &types5.GenesisState{}
			}
			if err := m.Subscriptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
