// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/node.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/sentinel-official/hub/types"
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

type Node struct {
	Address        string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	GigabytePrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=gigabyte_prices,json=gigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"gigabyte_prices"`
	HourlyPrices   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=hourly_prices,json=hourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"hourly_prices"`
	RemoteURL      string                                   `protobuf:"bytes,4,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
	Status         types1.Status                            `protobuf:"varint,5,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt       int64                                    `protobuf:"varint,6,opt,name=status_at,json=statusAt,proto3" json:"status_at,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_f03f9fca82881dc6, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Node)(nil), "sentinel.node.v2.Node")
}

func init() { proto.RegisterFile("sentinel/node/v2/node.proto", fileDescriptor_f03f9fca82881dc6) }

var fileDescriptor_f03f9fca82881dc6 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0xa1, 0xb8, 0x6e, 0xd9, 0xd6, 0x6e, 0x85, 0x7a, 0xc0, 0xb6, 0xb4, 0x46, 0x3d, 0x71,
	0xa8, 0x77, 0x0b, 0x7d, 0x82, 0xba, 0xd7, 0xaa, 0xad, 0xa8, 0x7c, 0xc9, 0xc5, 0xe2, 0xcf, 0x1a,
	0xaf, 0x82, 0x59, 0xc4, 0x2e, 0x28, 0x7e, 0x81, 0x9c, 0xf3, 0x18, 0x51, 0x9e, 0xc4, 0x47, 0x1f,
	0x73, 0x72, 0x12, 0xfc, 0x22, 0x91, 0x77, 0xc1, 0xca, 0x03, 0xe4, 0x34, 0xc3, 0x37, 0x33, 0xfa,
	0x7d, 0xcc, 0x2c, 0x98, 0x70, 0x92, 0x0b, 0x9a, 0x93, 0x0c, 0xe7, 0x2c, 0x21, 0xb8, 0xf6, 0x65,
	0x44, 0x45, 0xc9, 0x04, 0xb3, 0x3e, 0x77, 0x45, 0x24, 0xc5, 0xda, 0x1f, 0xc3, 0x98, 0xf1, 0x0d,
	0xe3, 0x38, 0x0a, 0x39, 0xc1, 0xb5, 0x17, 0x11, 0x11, 0x7a, 0x38, 0x66, 0x34, 0x57, 0x13, 0xe3,
	0x2f, 0x29, 0x4b, 0x99, 0x4c, 0xf1, 0x29, 0x6b, 0x55, 0x78, 0x86, 0x88, 0x6d, 0x41, 0x38, 0xae,
	0x3d, 0xcc, 0x45, 0x28, 0x2a, 0xae, 0xea, 0x5f, 0xaf, 0x0d, 0xd0, 0xfb, 0xc3, 0x12, 0x62, 0xd9,
	0xe0, 0x5d, 0x98, 0x24, 0x25, 0xe1, 0xdc, 0xd6, 0x1d, 0xdd, 0x35, 0x83, 0xee, 0xd3, 0x12, 0xe0,
	0x53, 0x4a, 0xd3, 0x30, 0xda, 0x0a, 0xb2, 0x2c, 0x4a, 0x1a, 0x13, 0x6e, 0xbf, 0x71, 0x0c, 0xf7,
	0x83, 0x3f, 0x42, 0xca, 0x12, 0x3a, 0x59, 0x42, 0xad, 0x25, 0xf4, 0x8b, 0xd1, 0x7c, 0xfe, 0x7d,
	0x77, 0x98, 0x6a, 0x77, 0x0f, 0x53, 0x37, 0xa5, 0x62, 0x5d, 0x45, 0x28, 0x66, 0x1b, 0xdc, 0xfa,
	0x57, 0x61, 0xc6, 0x93, 0x4b, 0x65, 0x49, 0x0e, 0xf0, 0x60, 0xd8, 0x31, 0xfe, 0x49, 0x84, 0x55,
	0x80, 0xc1, 0x9a, 0x55, 0x65, 0xb6, 0xed, 0x98, 0xc6, 0xeb, 0x33, 0x3f, 0x2a, 0x42, 0x4b, 0xfc,
	0x06, 0x40, 0x49, 0x36, 0x4c, 0x90, 0x65, 0x55, 0x66, 0x76, 0xef, 0xb4, 0x84, 0xf9, 0xa0, 0x39,
	0x4c, 0xcd, 0x40, 0xaa, 0x8b, 0xe0, 0x77, 0x60, 0xaa, 0x86, 0x45, 0x99, 0x59, 0x1e, 0xe8, 0xab,
	0x45, 0xda, 0x6f, 0x1d, 0xdd, 0x1d, 0xfa, 0x23, 0x74, 0xbe, 0x98, 0x42, 0xd4, 0x1e, 0xfa, 0x2f,
	0x1b, 0x82, 0xb6, 0xd1, 0x9a, 0x00, 0x53, 0x65, 0xcb, 0x50, 0xd8, 0x7d, 0x47, 0x77, 0x8d, 0xe0,
	0xbd, 0x12, 0x7e, 0x8a, 0xf9, 0xdf, 0xdd, 0x13, 0xd4, 0x6e, 0x1b, 0xa8, 0xed, 0x1a, 0xa8, 0xef,
	0x1b, 0xa8, 0x3f, 0x36, 0x50, 0xbf, 0x39, 0x42, 0x6d, 0x7f, 0x84, 0xda, 0xfd, 0x11, 0x6a, 0x17,
	0xb3, 0x17, 0xff, 0xd5, 0xb1, 0x66, 0x6c, 0xb5, 0xa2, 0x31, 0x0d, 0x33, 0xbc, 0xae, 0x22, 0x7c,
	0xa5, 0x5e, 0x92, 0xe4, 0x47, 0x7d, 0x79, 0xe0, 0x1f, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xce,
	0xf7, 0xa0, 0x30, 0x67, 0x02, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StatusAt != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.StatusAt))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RemoteURL) > 0 {
		i -= len(m.RemoteURL)
		copy(dAtA[i:], m.RemoteURL)
		i = encodeVarintNode(dAtA, i, uint64(len(m.RemoteURL)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HourlyPrices) > 0 {
		for iNdEx := len(m.HourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.GigabytePrices) > 0 {
		for iNdEx := len(m.GigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if len(m.GigabytePrices) > 0 {
		for _, e := range m.GigabytePrices {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	if len(m.HourlyPrices) > 0 {
		for _, e := range m.HourlyPrices {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	l = len(m.RemoteURL)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovNode(uint64(m.Status))
	}
	if m.StatusAt != 0 {
		n += 1 + sovNode(uint64(m.StatusAt))
	}
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GigabytePrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GigabytePrices = append(m.GigabytePrices, types.Coin{})
			if err := m.GigabytePrices[len(m.GigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HourlyPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HourlyPrices = append(m.HourlyPrices, types.Coin{})
			if err := m.HourlyPrices[len(m.HourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			m.StatusAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)
