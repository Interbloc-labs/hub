// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Params struct {
	InactiveDuration         time.Duration `protobuf:"bytes,1,opt,name=inactive_duration,json=inactiveDuration,proto3,stdduration" json:"inactive_duration"`
	ProofVerificationEnabled bool          `protobuf:"varint,2,opt,name=proof_verification_enabled,json=proofVerificationEnabled,proto3" json:"proof_verification_enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b86e8f1be7738cd, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "sentinel.session.v1.Params")
}

func init() { proto.RegisterFile("sentinel/session/v1/params.proto", fileDescriptor_6b86e8f1be7738cd) }

var fileDescriptor_6b86e8f1be7738cd = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xed, 0x3b, 0x54, 0x55, 0xee, 0x02, 0x85, 0x21, 0x64, 0x70, 0x23, 0xa6, 0x2e, 0xd8,
	0x2a, 0x6c, 0x88, 0xa9, 0x82, 0xbd, 0xea, 0xc0, 0x80, 0x90, 0x22, 0x27, 0x75, 0x5c, 0x4b, 0x69,
	0xbe, 0x28, 0x76, 0x22, 0xfa, 0x16, 0x8c, 0x7d, 0x04, 0x1e, 0x25, 0x63, 0x47, 0x26, 0xfe, 0x24,
	0x2f, 0x82, 0xe2, 0xd4, 0x88, 0xcd, 0x3e, 0xe7, 0x77, 0x8e, 0xf4, 0x1d, 0x2f, 0xd4, 0x22, 0x37,
	0x2a, 0x17, 0x19, 0xd3, 0x42, 0x6b, 0x05, 0x39, 0xab, 0xe7, 0xac, 0xe0, 0x25, 0xdf, 0x6a, 0x5a,
	0x94, 0x60, 0x60, 0x72, 0xe6, 0x08, 0x7a, 0x24, 0x68, 0x3d, 0x0f, 0xce, 0x25, 0x48, 0xb0, 0x3e,
	0xeb, 0x5f, 0x03, 0x1a, 0x10, 0x09, 0x20, 0x33, 0xc1, 0xec, 0x2f, 0xae, 0x52, 0xb6, 0xae, 0x4a,
	0x6e, 0xfa, 0x88, 0x55, 0x2e, 0xf7, 0xd8, 0x1b, 0x2d, 0x6d, 0xf7, 0x64, 0xe9, 0x9d, 0xaa, 0x9c,
	0x27, 0x46, 0xd5, 0x22, 0x72, 0x94, 0x8f, 0x43, 0x3c, 0xfb, 0x7f, 0x7d, 0x41, 0x87, 0x1a, 0xea,
	0x6a, 0xe8, 0xfd, 0x11, 0x58, 0x8c, 0x9b, 0x8f, 0x29, 0xda, 0x7f, 0x4e, 0xf1, 0xea, 0xc4, 0xa5,
	0x9d, 0x37, 0xb9, 0xf3, 0x82, 0xa2, 0x04, 0x48, 0xa3, 0x5a, 0x94, 0x2a, 0x55, 0x89, 0x55, 0x23,
	0x91, 0xf3, 0x38, 0x13, 0x6b, 0xff, 0x5f, 0x88, 0x67, 0xe3, 0x95, 0x6f, 0x89, 0xc7, 0x3f, 0xc0,
	0xc3, 0xe0, 0x2f, 0x9e, 0x9b, 0x6f, 0x82, 0xde, 0x5a, 0x82, 0x9a, 0x96, 0xe0, 0x43, 0x4b, 0xf0,
	0x57, 0x4b, 0xf0, 0x6b, 0x47, 0xd0, 0xa1, 0x23, 0xe8, 0xbd, 0x23, 0xe8, 0xe9, 0x56, 0x2a, 0xb3,
	0xa9, 0x62, 0x9a, 0xc0, 0x96, 0xb9, 0x49, 0xae, 0x20, 0x4d, 0x55, 0xa2, 0x78, 0xc6, 0x36, 0x55,
	0xcc, 0x5e, 0x7e, 0x37, 0xcc, 0x84, 0xe4, 0xc9, 0xae, 0x9f, 0xd2, 0xec, 0x0a, 0xa1, 0xe3, 0x91,
	0x3d, 0xe5, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x80, 0xd0, 0xdf, 0xd2, 0x6e, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProofVerificationEnabled {
		i--
		if m.ProofVerificationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.InactiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.InactiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.InactiveDuration)
	n += 1 + l + sovParams(uint64(l))
	if m.ProofVerificationEnabled {
		n += 2
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.InactiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofVerificationEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProofVerificationEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
