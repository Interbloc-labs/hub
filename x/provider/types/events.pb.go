// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/provider/v1/events.proto

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

type EventRegisterProvider struct {
	From        string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Identity    string `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Website     string `protobuf:"bytes,5,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *EventRegisterProvider) Reset()         { *m = EventRegisterProvider{} }
func (m *EventRegisterProvider) String() string { return proto.CompactTextString(m) }
func (*EventRegisterProvider) ProtoMessage()    {}
func (*EventRegisterProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec27b3fe604bfcc, []int{0}
}
func (m *EventRegisterProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegisterProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegisterProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegisterProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegisterProvider.Merge(m, src)
}
func (m *EventRegisterProvider) XXX_Size() int {
	return m.Size()
}
func (m *EventRegisterProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegisterProvider.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegisterProvider proto.InternalMessageInfo

type EventUpdateProvider struct {
	From        string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Identity    string `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Website     string `protobuf:"bytes,5,opt,name=website,proto3" json:"website,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *EventUpdateProvider) Reset()         { *m = EventUpdateProvider{} }
func (m *EventUpdateProvider) String() string { return proto.CompactTextString(m) }
func (*EventUpdateProvider) ProtoMessage()    {}
func (*EventUpdateProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ec27b3fe604bfcc, []int{1}
}
func (m *EventUpdateProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateProvider.Merge(m, src)
}
func (m *EventUpdateProvider) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateProvider.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateProvider proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRegisterProvider)(nil), "sentinel.provider.v1.EventRegisterProvider")
	proto.RegisterType((*EventUpdateProvider)(nil), "sentinel.provider.v1.EventUpdateProvider")
}

func init() { proto.RegisterFile("sentinel/provider/v1/events.proto", fileDescriptor_3ec27b3fe604bfcc) }

var fileDescriptor_3ec27b3fe604bfcc = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x91, 0x31, 0x4e, 0xf3, 0x30,
	0x1c, 0xc5, 0xe3, 0xef, 0x2b, 0x05, 0xcc, 0x66, 0x8a, 0x64, 0x75, 0xb0, 0x4a, 0x27, 0x16, 0x62,
	0x45, 0xdc, 0x00, 0x89, 0x1d, 0x15, 0xb1, 0xb0, 0x25, 0xf1, 0x3f, 0xa9, 0xa5, 0x26, 0xb6, 0x6c,
	0x37, 0xd0, 0x5b, 0x70, 0x0c, 0x36, 0xc4, 0x2d, 0x3a, 0x76, 0x64, 0x84, 0xe4, 0x22, 0xc8, 0x4e,
	0x03, 0x1c, 0x81, 0xed, 0xbd, 0xe7, 0x9f, 0xff, 0x7a, 0xd2, 0xc3, 0xe7, 0x16, 0x6a, 0x27, 0x6b,
	0x58, 0x71, 0x6d, 0x54, 0x23, 0x05, 0x18, 0xde, 0x24, 0x1c, 0x1a, 0xa8, 0x9d, 0x8d, 0xb5, 0x51,
	0x4e, 0x91, 0xc9, 0x80, 0xc4, 0x03, 0x12, 0x37, 0xc9, 0x74, 0x52, 0xaa, 0x52, 0x05, 0x80, 0x7b,
	0xd5, 0xb3, 0xf3, 0x37, 0x84, 0xcf, 0x6e, 0xfc, 0xe7, 0x05, 0x94, 0xd2, 0x3a, 0x30, 0xb7, 0xfb,
	0x2f, 0x84, 0xe0, 0x51, 0x61, 0x54, 0x45, 0xd1, 0x0c, 0x5d, 0x1c, 0x2f, 0x82, 0x26, 0x14, 0x1f,
	0xa6, 0x42, 0x18, 0xb0, 0x96, 0xfe, 0x0b, 0xf1, 0x60, 0x3d, 0x5d, 0xa7, 0x15, 0xd0, 0xff, 0x3d,
	0xed, 0x35, 0x99, 0xe2, 0x23, 0x29, 0x7c, 0x15, 0xb7, 0xa1, 0xa3, 0x90, 0x7f, 0x7b, 0x7f, 0xe9,
	0x11, 0x32, 0x2b, 0x1d, 0xd0, 0x83, 0xfe, 0xd2, 0xde, 0x92, 0x19, 0x3e, 0x11, 0x60, 0x73, 0x23,
	0xb5, 0x93, 0xaa, 0xa6, 0xe3, 0xf0, 0xfa, 0x3b, 0x9a, 0xbf, 0x22, 0x7c, 0x1a, 0x3a, 0xdf, 0x6b,
	0x91, 0x3a, 0xf8, 0xfb, 0x8d, 0xaf, 0xef, 0xb6, 0x9f, 0x2c, 0x7a, 0x69, 0x59, 0xb4, 0x6d, 0x19,
	0xda, 0xb5, 0x0c, 0x7d, 0xb4, 0x0c, 0x3d, 0x77, 0x2c, 0xda, 0x75, 0x2c, 0x7a, 0xef, 0x58, 0xf4,
	0x90, 0x94, 0xd2, 0x2d, 0xd7, 0x59, 0x9c, 0xab, 0x8a, 0x0f, 0xf3, 0x5d, 0xaa, 0xa2, 0x90, 0xb9,
	0x4c, 0x57, 0x7c, 0xb9, 0xce, 0xf8, 0xd3, 0xcf, 0xe0, 0x6e, 0xa3, 0xc1, 0x66, 0xe3, 0xb0, 0xe0,
	0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xbd, 0xdd, 0xaf, 0x12, 0x02, 0x00, 0x00,
}

func (m *EventRegisterProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegisterProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegisterProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
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
func (m *EventRegisterProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Description)
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
func (m *EventRegisterProvider) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventRegisterProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegisterProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
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
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
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
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
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
func (m *EventUpdateProvider) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventUpdateProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
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
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
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
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
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
