// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/proof.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/sentinel-official/hub/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Proof struct {
	Id        uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Duration  time.Duration   `protobuf:"bytes,2,opt,name=duration,proto3,stdduration" json:"duration"`
	Bandwidth types.Bandwidth `protobuf:"bytes,3,opt,name=bandwidth,proto3" json:"bandwidth"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c7e6824be930eaf, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Proof)(nil), "sentinel.session.v1.Proof")
}

func init() { proto.RegisterFile("sentinel/session/v1/proof.proto", fileDescriptor_5c7e6824be930eaf) }

var fileDescriptor_5c7e6824be930eaf = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x3d, 0x4e, 0x84, 0x40,
	0x14, 0xc7, 0x67, 0x70, 0x35, 0x2b, 0x26, 0x16, 0x68, 0x81, 0x1b, 0x33, 0xac, 0x56, 0xdb, 0xf8,
	0x46, 0xf4, 0x00, 0x1a, 0xe2, 0x01, 0x0c, 0xa5, 0x1d, 0x2c, 0x5f, 0x93, 0x20, 0x8f, 0xc0, 0x80,
	0x7a, 0x0b, 0x4b, 0xe3, 0x09, 0x3c, 0x0a, 0xe5, 0x96, 0x56, 0x7e, 0xc0, 0x45, 0x0c, 0xc3, 0x82,
	0xdd, 0x4c, 0xde, 0xef, 0xff, 0xff, 0xbd, 0x19, 0xdd, 0x2a, 0xc3, 0x4c, 0x8a, 0x2c, 0x4c, 0x79,
	0x19, 0x96, 0xa5, 0xc0, 0x8c, 0xd7, 0x36, 0xcf, 0x0b, 0xc4, 0x08, 0xf2, 0x02, 0x25, 0x1a, 0x47,
	0x23, 0x00, 0x5b, 0x00, 0x6a, 0x7b, 0x71, 0x1c, 0x63, 0x8c, 0x6a, 0xce, 0xfb, 0xd3, 0x80, 0x2e,
	0x58, 0x8c, 0x18, 0xa7, 0x21, 0x57, 0x37, 0xbf, 0x8a, 0x78, 0x50, 0x15, 0x9e, 0xec, 0x23, 0xc3,
	0xfc, 0x6c, 0x72, 0xc9, 0x97, 0x3c, 0x2c, 0x7b, 0x93, 0xef, 0x65, 0xc1, 0x93, 0x08, 0x64, 0x32,
	0x20, 0xe7, 0xef, 0x54, 0xdf, 0xbd, 0xef, 0xed, 0xc6, 0xa1, 0xae, 0x89, 0xc0, 0xa4, 0x4b, 0xba,
	0x9a, 0xb9, 0x9a, 0x08, 0x8c, 0x1b, 0x7d, 0x3e, 0xd6, 0x99, 0xda, 0x92, 0xae, 0x0e, 0xae, 0x4e,
	0x60, 0xf0, 0xc1, 0xe8, 0x83, 0xbb, 0x2d, 0xe0, 0xcc, 0x9b, 0x2f, 0x8b, 0xbc, 0x7d, 0x5b, 0xd4,
	0x9d, 0x42, 0xc6, 0xad, 0xbe, 0x3f, 0xd9, 0xcc, 0x1d, 0xd5, 0x70, 0x0a, 0xd3, 0xe3, 0xd4, 0x46,
	0x50, 0xdb, 0xe0, 0x8c, 0x8c, 0x33, 0xeb, 0x4b, 0xdc, 0xff, 0x90, 0xe3, 0x36, 0xbf, 0x8c, 0x7c,
	0xb4, 0x8c, 0x34, 0x2d, 0xa3, 0x9b, 0x96, 0xd1, 0x9f, 0x96, 0xd1, 0xd7, 0x8e, 0x91, 0x4d, 0xc7,
	0xc8, 0x67, 0xc7, 0xc8, 0xc3, 0x65, 0x2c, 0x64, 0x52, 0xf9, 0xb0, 0xc6, 0x47, 0x3e, 0x56, 0x5f,
	0x60, 0x14, 0x89, 0xb5, 0xf0, 0x52, 0x9e, 0x54, 0x3e, 0x7f, 0x9e, 0xfe, 0x59, 0x19, 0xfd, 0x3d,
	0xb5, 0xfc, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x7a, 0x9e, 0x80, 0x88, 0x01, 0x00,
	0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProof(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintProof(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintProof(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProof(uint64(m.Id))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovProof(uint64(l))
	l = m.Bandwidth.Size()
	n += 1 + l + sovProof(uint64(l))
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)
