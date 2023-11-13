// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// Params represents a message for handling node parameters.
type Params struct {
	// Field 1: Deposit required for registering a node.
	// - (gogoproto.nullable) = false: Field is not nullable.
	Deposit types.Coin `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	// Field 2: Duration for which a node remains active.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.stdduration) = true: Use standard duration representation for Go.
	ActiveDuration time.Duration `protobuf:"bytes,2,opt,name=active_duration,json=activeDuration,proto3,stdduration" json:"active_duration"`
	// Field 3: Maximum prices in gigabytes for a node's subscription.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins":
	//   Type to cast to when repeating this field.
	MaxGigabytePrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=max_gigabyte_prices,json=maxGigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_gigabyte_prices"`
	// Field 4: Minimum prices in gigabytes for a node's subscription.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins":
	//   Type to cast to when repeating this field.
	MinGigabytePrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=min_gigabyte_prices,json=minGigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_gigabyte_prices"`
	// Field 5: Maximum prices in hours for a node's subscription.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins":
	//   Type to cast to when repeating this field.
	MaxHourlyPrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=max_hourly_prices,json=maxHourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_hourly_prices"`
	// Field 6: Minimum prices in hours for a node's subscription.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins":
	//   Type to cast to when repeating this field.
	MinHourlyPrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=min_hourly_prices,json=minHourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_hourly_prices"`
	// Field 7: Maximum gigabytes allowed for a subscription.
	MaxSubscriptionGigabytes int64 `protobuf:"varint,7,opt,name=max_subscription_gigabytes,json=maxSubscriptionGigabytes,proto3" json:"max_subscription_gigabytes,omitempty"`
	// Field 8: Minimum gigabytes required for a subscription.
	MinSubscriptionGigabytes int64 `protobuf:"varint,8,opt,name=min_subscription_gigabytes,json=minSubscriptionGigabytes,proto3" json:"min_subscription_gigabytes,omitempty"`
	// Field 9: Maximum hours allowed for a subscription.
	MaxSubscriptionHours int64 `protobuf:"varint,9,opt,name=max_subscription_hours,json=maxSubscriptionHours,proto3" json:"max_subscription_hours,omitempty"`
	// Field 10: Minimum hours required for a subscription.
	MinSubscriptionHours int64 `protobuf:"varint,10,opt,name=min_subscription_hours,json=minSubscriptionHours,proto3" json:"min_subscription_hours,omitempty"`
	// Field 11: Staking share required for a node.
	// - (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Dec":
	//   Custom type definition for the field.
	// - (gogoproto.nullable) = false: Field is not nullable.
	StakingShare github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=staking_share,json=stakingShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_share"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f1279255e9f358, []int{0}
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
	proto.RegisterType((*Params)(nil), "sentinel.node.v2.Params")
}

func init() { proto.RegisterFile("sentinel/node/v2/params.proto", fileDescriptor_02f1279255e9f358) }

var fileDescriptor_02f1279255e9f358 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xd3, 0x30,
	0x00, 0xc6, 0x13, 0x36, 0xda, 0xcd, 0x03, 0x06, 0x65, 0x42, 0xa1, 0x12, 0x6e, 0xc5, 0x01, 0xf5,
	0x32, 0x7b, 0x2d, 0x5c, 0x90, 0x38, 0x95, 0x49, 0x70, 0xe0, 0x30, 0xa5, 0x37, 0x2e, 0x95, 0x93,
	0xba, 0xa9, 0xb5, 0xc6, 0x8e, 0x62, 0xa7, 0xa4, 0xe2, 0x25, 0x38, 0xf2, 0x08, 0x88, 0x47, 0xe0,
	0x09, 0x7a, 0xdc, 0x11, 0x71, 0xd8, 0xa0, 0x7d, 0x11, 0xe4, 0x3f, 0x41, 0x5d, 0x07, 0x88, 0x03,
	0x3b, 0xc5, 0x89, 0xfd, 0xfb, 0x7e, 0xce, 0x17, 0xc5, 0xe0, 0x91, 0xa4, 0x5c, 0x31, 0x4e, 0xa7,
	0x98, 0x8b, 0x11, 0xc5, 0xb3, 0x1e, 0xce, 0x48, 0x4e, 0x52, 0x89, 0xb2, 0x5c, 0x28, 0xd1, 0xb8,
	0x5b, 0x4d, 0x23, 0x3d, 0x8d, 0x66, 0xbd, 0x26, 0x8c, 0x85, 0x4c, 0x85, 0xc4, 0x11, 0x91, 0x14,
	0xcf, 0xba, 0x11, 0x55, 0xa4, 0x8b, 0x63, 0xc1, 0xb8, 0x25, 0x9a, 0x07, 0x89, 0x48, 0x84, 0x19,
	0x62, 0x3d, 0x72, 0x4f, 0x61, 0x22, 0x44, 0x32, 0xa5, 0xd8, 0xdc, 0x45, 0xc5, 0x18, 0x8f, 0x8a,
	0x9c, 0x28, 0x26, 0x1c, 0xf5, 0xf8, 0x4b, 0x1d, 0xd4, 0x4e, 0x8c, 0xb8, 0xf1, 0x1c, 0xd4, 0x47,
	0x34, 0x13, 0x92, 0xa9, 0xc0, 0x6f, 0xfb, 0x9d, 0xbd, 0xde, 0x43, 0x64, 0x95, 0x48, 0x2b, 0x91,
	0x53, 0xa2, 0x97, 0x82, 0xf1, 0xfe, 0xf6, 0xe2, 0xbc, 0xe5, 0x85, 0xd5, 0xfa, 0xc6, 0x1b, 0xb0,
	0x4f, 0x62, 0xc5, 0x66, 0x74, 0x58, 0xc5, 0x07, 0x37, 0x5c, 0x84, 0xf5, 0xa3, 0xca, 0x8f, 0x8e,
	0xdd, 0x82, 0xfe, 0x8e, 0x8e, 0xf8, 0x78, 0xd1, 0xf2, 0xc3, 0x3b, 0x96, 0xad, 0x66, 0x1a, 0xef,
	0xc1, 0xfd, 0x94, 0x94, 0xc3, 0x84, 0x25, 0x24, 0x9a, 0x2b, 0x3a, 0xcc, 0x72, 0x16, 0x53, 0x19,
	0x6c, 0xb5, 0xb7, 0xfe, 0xbe, 0xa9, 0x23, 0x9d, 0xf8, 0xf9, 0xa2, 0xd5, 0x49, 0x98, 0x9a, 0x14,
	0x11, 0x8a, 0x45, 0x8a, 0x5d, 0x69, 0xf6, 0x72, 0x28, 0x47, 0xa7, 0x58, 0xcd, 0x33, 0x2a, 0x0d,
	0x20, 0xc3, 0x7b, 0x29, 0x29, 0x5f, 0x39, 0xcd, 0x89, 0xb1, 0x18, 0x39, 0xe3, 0x57, 0xe4, 0xdb,
	0xd7, 0x21, 0x67, 0x7c, 0x43, 0xfe, 0x0e, 0xe8, 0x1d, 0x0d, 0x27, 0xa2, 0xc8, 0xa7, 0xf3, 0x4a,
	0x7d, 0xf3, 0xff, 0xab, 0xf7, 0x53, 0x52, 0xbe, 0x36, 0x92, 0x35, 0x31, 0xe3, 0x1b, 0xe2, 0xda,
	0x75, 0x88, 0x19, 0xbf, 0x24, 0x7e, 0x01, 0x9a, 0xfa, 0x8d, 0x65, 0x11, 0xc9, 0x38, 0x67, 0x99,
	0xfe, 0xfe, 0xbf, 0xba, 0x97, 0x41, 0xbd, 0xed, 0x77, 0xb6, 0xc2, 0x20, 0x25, 0xe5, 0x60, 0x6d,
	0x41, 0x55, 0x9a, 0xa5, 0x19, 0xff, 0x13, 0xbd, 0xe3, 0x68, 0xc6, 0x7f, 0x4f, 0x3f, 0x03, 0x0f,
	0xae, 0xb8, 0x75, 0x03, 0x32, 0xd8, 0x35, 0xe4, 0xc1, 0x86, 0x57, 0x6f, 0xdc, 0x52, 0x9b, 0x4e,
	0x4b, 0x01, 0x47, 0x5d, 0xf6, 0x59, 0x6a, 0x00, 0x6e, 0x4b, 0x45, 0x4e, 0x19, 0x4f, 0x86, 0x72,
	0x42, 0x72, 0x1a, 0xec, 0xb5, 0xfd, 0xce, 0x6e, 0x1f, 0xe9, 0x06, 0xbf, 0x9d, 0xb7, 0x9e, 0xfc,
	0x43, 0x83, 0xc7, 0x34, 0x0e, 0x6f, 0xb9, 0x90, 0x81, 0xce, 0xe8, 0x87, 0x8b, 0x1f, 0xd0, 0xfb,
	0xb4, 0x84, 0xde, 0x62, 0x09, 0xfd, 0xb3, 0x25, 0xf4, 0xbf, 0x2f, 0xa1, 0xff, 0x61, 0x05, 0xbd,
	0xb3, 0x15, 0xf4, 0xbe, 0xae, 0xa0, 0xf7, 0xf6, 0x68, 0x2d, 0xb7, 0x3a, 0x51, 0x0e, 0xc5, 0x78,
	0xcc, 0x62, 0x46, 0xa6, 0x78, 0x52, 0x44, 0x78, 0xd6, 0xc5, 0xa5, 0x3d, 0x82, 0x8c, 0x25, 0xaa,
	0x99, 0x3f, 0xf5, 0xe9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xbd, 0xec, 0x1a, 0xa0, 0x04,
	0x00, 0x00,
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
	{
		size := m.StakingShare.Size()
		i -= size
		if _, err := m.StakingShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if m.MinSubscriptionHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSubscriptionHours))
		i--
		dAtA[i] = 0x50
	}
	if m.MaxSubscriptionHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSubscriptionHours))
		i--
		dAtA[i] = 0x48
	}
	if m.MinSubscriptionGigabytes != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSubscriptionGigabytes))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxSubscriptionGigabytes != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSubscriptionGigabytes))
		i--
		dAtA[i] = 0x38
	}
	if len(m.MinHourlyPrices) > 0 {
		for iNdEx := len(m.MinHourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinHourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.MaxHourlyPrices) > 0 {
		for iNdEx := len(m.MaxHourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxHourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MinGigabytePrices) > 0 {
		for iNdEx := len(m.MinGigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinGigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MaxGigabytePrices) > 0 {
		for iNdEx := len(m.MaxGigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxGigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ActiveDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
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
	l = m.Deposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	if len(m.MaxGigabytePrices) > 0 {
		for _, e := range m.MaxGigabytePrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinGigabytePrices) > 0 {
		for _, e := range m.MinGigabytePrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MaxHourlyPrices) > 0 {
		for _, e := range m.MaxHourlyPrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinHourlyPrices) > 0 {
		for _, e := range m.MinHourlyPrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxSubscriptionGigabytes != 0 {
		n += 1 + sovParams(uint64(m.MaxSubscriptionGigabytes))
	}
	if m.MinSubscriptionGigabytes != 0 {
		n += 1 + sovParams(uint64(m.MinSubscriptionGigabytes))
	}
	if m.MaxSubscriptionHours != 0 {
		n += 1 + sovParams(uint64(m.MaxSubscriptionHours))
	}
	if m.MinSubscriptionHours != 0 {
		n += 1 + sovParams(uint64(m.MinSubscriptionHours))
	}
	l = m.StakingShare.Size()
	n += 1 + l + sovParams(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveDuration", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGigabytePrices", wireType)
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
			m.MaxGigabytePrices = append(m.MaxGigabytePrices, types.Coin{})
			if err := m.MaxGigabytePrices[len(m.MaxGigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGigabytePrices", wireType)
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
			m.MinGigabytePrices = append(m.MinGigabytePrices, types.Coin{})
			if err := m.MinGigabytePrices[len(m.MinGigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxHourlyPrices", wireType)
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
			m.MaxHourlyPrices = append(m.MaxHourlyPrices, types.Coin{})
			if err := m.MaxHourlyPrices[len(m.MaxHourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinHourlyPrices", wireType)
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
			m.MinHourlyPrices = append(m.MinHourlyPrices, types.Coin{})
			if err := m.MinHourlyPrices[len(m.MinHourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubscriptionGigabytes", wireType)
			}
			m.MaxSubscriptionGigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubscriptionGigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSubscriptionGigabytes", wireType)
			}
			m.MinSubscriptionGigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSubscriptionGigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubscriptionHours", wireType)
			}
			m.MaxSubscriptionHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubscriptionHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSubscriptionHours", wireType)
			}
			m.MinSubscriptionHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSubscriptionHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
