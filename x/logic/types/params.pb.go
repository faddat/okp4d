// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logic/v1beta2/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines all the configuration parameters of the "logic" module.
type Params struct {
	// Interpreter specifies the parameter for the logic interpreter.
	Interpreter Interpreter `protobuf:"bytes,1,opt,name=interpreter,proto3" json:"interpreter" yaml:"interpreter"`
	// Limits defines the limits of the logic module.
	// The limits are used to prevent the interpreter from running for too long.
	// If the interpreter runs for too long, the execution will be aborted.
	Limits Limits `protobuf:"bytes,2,opt,name=limits,proto3" json:"limits" yaml:"limits"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af0daa241de0fa3, []int{0}
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

func (m *Params) GetInterpreter() Interpreter {
	if m != nil {
		return m.Interpreter
	}
	return Interpreter{}
}

func (m *Params) GetLimits() Limits {
	if m != nil {
		return m.Limits
	}
	return Limits{}
}

// Limits defines the limits of the logic module.
type Limits struct {
	// max_gas specifies the maximum amount of computing power, measured in "gas," that is allowed to be consumed when
	// executing a request by the interpreter. The interpreter calculates the gas consumption based on the number and type
	// of operations that are executed, as well as, in some cases, the complexity of the processed data.
	// nil value remove max gas limitation.
	MaxGas *github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=max_gas,json=maxGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"max_gas,omitempty" yaml:"max_gas",omitempty`
	// max_size specifies the maximum size, in bytes, that is accepted for a program.
	// nil value remove size limitation.
	MaxSize *github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=max_size,json=maxSize,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"max_size,omitempty" yaml:"max_size"`
	// max_result_count specifies the maximum number of results that can be requested for a query.
	// nil value remove max result count limitation.
	MaxResultCount *github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=max_result_count,json=maxResultCount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"max_result_count,omitempty" yaml:"max_result_count"`
}

func (m *Limits) Reset()         { *m = Limits{} }
func (m *Limits) String() string { return proto.CompactTextString(m) }
func (*Limits) ProtoMessage()    {}
func (*Limits) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af0daa241de0fa3, []int{1}
}
func (m *Limits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Limits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Limits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Limits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Limits.Merge(m, src)
}
func (m *Limits) XXX_Size() int {
	return m.Size()
}
func (m *Limits) XXX_DiscardUnknown() {
	xxx_messageInfo_Limits.DiscardUnknown(m)
}

var xxx_messageInfo_Limits proto.InternalMessageInfo

// Interpreter defines the various parameters for the interpreter.
type Interpreter struct {
	// predicates_whitelist specifies a list of prolog predicates that are allowed and can be used by the interpreter.
	// The predicates are represented as `<predicate_name>/[<arity>]`, for example: `findall/3`, or `call`. If a predicate name without arity
	// is included in this list, then all predicates with that name will be considered regardless of arity. For example, if `call` is included
	// in the whitelist, then all predicates `call/1`, `call/2`, `call/3`... will be allowed.
	// If this field is not specified, the interpreter will use the default set of predicates.
	PredicatesWhitelist []string `protobuf:"bytes,1,rep,name=predicates_whitelist,json=predicatesWhitelist,proto3" json:"predicates_whitelist,omitempty" yaml:"predicates_whitelist"`
	// predicates_blacklist specifies a list of prolog predicates that are excluded from the set of registered predicates
	// and can never be executed by the interpreter.
	// The predicates are represented as `<predicate_name>/[<arity>]`, for example: `findall/3`, or `call`. If a predicate name without arity
	// is included in this list, then all predicates with that name will be considered regardless of arity. For example, if `call` is included
	// in the blacklist, then all predicates `call/1`, `call/2`, `call/3`... will be excluded.
	// If a predicate is included in both whitelist and blacklist, it will be excluded. This means that blacklisted predicates prevails
	// on whitelisted predicates.
	PredicatesBlacklist []string `protobuf:"bytes,2,rep,name=predicates_blacklist,json=predicatesBlacklist,proto3" json:"predicates_blacklist,omitempty" yaml:"predicates_blacklist"`
	// bootstrap specifies the initial program to run when booting the logic interpreter.
	// If not specified, the default boot sequence will be executed.
	Bootstrap string `protobuf:"bytes,3,opt,name=bootstrap,proto3" json:"bootstrap,omitempty" yaml:"bootstrap"`
}

func (m *Interpreter) Reset()         { *m = Interpreter{} }
func (m *Interpreter) String() string { return proto.CompactTextString(m) }
func (*Interpreter) ProtoMessage()    {}
func (*Interpreter) Descriptor() ([]byte, []int) {
	return fileDescriptor_3af0daa241de0fa3, []int{2}
}
func (m *Interpreter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Interpreter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Interpreter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Interpreter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interpreter.Merge(m, src)
}
func (m *Interpreter) XXX_Size() int {
	return m.Size()
}
func (m *Interpreter) XXX_DiscardUnknown() {
	xxx_messageInfo_Interpreter.DiscardUnknown(m)
}

var xxx_messageInfo_Interpreter proto.InternalMessageInfo

func (m *Interpreter) GetPredicatesWhitelist() []string {
	if m != nil {
		return m.PredicatesWhitelist
	}
	return nil
}

func (m *Interpreter) GetPredicatesBlacklist() []string {
	if m != nil {
		return m.PredicatesBlacklist
	}
	return nil
}

func (m *Interpreter) GetBootstrap() string {
	if m != nil {
		return m.Bootstrap
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "logic.v1beta2.Params")
	proto.RegisterType((*Limits)(nil), "logic.v1beta2.Limits")
	proto.RegisterType((*Interpreter)(nil), "logic.v1beta2.Interpreter")
}

func init() { proto.RegisterFile("logic/v1beta2/params.proto", fileDescriptor_3af0daa241de0fa3) }

var fileDescriptor_3af0daa241de0fa3 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x33, 0x5b, 0x89, 0xee, 0x2c, 0xd5, 0x12, 0x5b, 0x0c, 0x2b, 0x24, 0x65, 0x7a, 0xb0,
	0x07, 0x4d, 0xb0, 0x8a, 0x87, 0x82, 0x97, 0x28, 0x88, 0x7f, 0x0e, 0x12, 0xf1, 0x0f, 0x82, 0x2c,
	0x93, 0xec, 0x90, 0x0e, 0x9b, 0xd9, 0x09, 0x99, 0x77, 0x75, 0xb7, 0x9f, 0xc2, 0xa3, 0x47, 0x3f,
	0x81, 0x9f, 0x63, 0xbd, 0xf5, 0x28, 0x1e, 0x82, 0xec, 0x7e, 0x83, 0xbd, 0x7a, 0x91, 0x4c, 0xd2,
	0x26, 0x2d, 0x42, 0xd9, 0x4b, 0x12, 0xde, 0xe7, 0x79, 0x7f, 0xcf, 0x9b, 0x77, 0x18, 0xdc, 0x4f,
	0x65, 0xc2, 0x63, 0xff, 0xf3, 0xfd, 0x88, 0x01, 0x3d, 0xf0, 0x33, 0x9a, 0x53, 0xa1, 0xbc, 0x2c,
	0x97, 0x20, 0xad, 0x4d, 0xad, 0x79, 0xb5, 0xd6, 0xdf, 0x4e, 0x64, 0x22, 0xb5, 0xe2, 0x97, 0x5f,
	0x95, 0x89, 0xfc, 0x40, 0xd8, 0x7c, 0xad, 0xbb, 0xac, 0x0f, 0xb8, 0xc7, 0xc7, 0xc0, 0xf2, 0x2c,
	0x67, 0xc0, 0x72, 0x1b, 0xed, 0xa2, 0xfd, 0xde, 0x41, 0xdf, 0x3b, 0x47, 0xf1, 0x9e, 0x37, 0x8e,
	0xa0, 0x3f, 0x2f, 0x5c, 0x63, 0x55, 0xb8, 0xd6, 0x8c, 0x8a, 0xf4, 0x90, 0xb4, 0x9a, 0x49, 0xd8,
	0x46, 0x59, 0x4f, 0xb1, 0x99, 0x72, 0xc1, 0x41, 0xd9, 0x1d, 0x0d, 0xdd, 0xb9, 0x00, 0x7d, 0xa5,
	0xc5, 0x60, 0xa7, 0xe6, 0x6d, 0x56, 0xbc, 0xaa, 0x85, 0x84, 0x75, 0xef, 0xe1, 0x95, 0x6f, 0xdf,
	0x5d, 0x83, 0xfc, 0xec, 0x60, 0xb3, 0xf2, 0x5b, 0x43, 0x7c, 0x55, 0xd0, 0xe9, 0x20, 0xa1, 0x4a,
	0x0f, 0xdb, 0x0d, 0x5e, 0xce, 0x0b, 0x17, 0xfd, 0x2e, 0xdc, 0x3b, 0x09, 0x87, 0xa3, 0x49, 0xe4,
	0xc5, 0x52, 0xf8, 0xb1, 0x54, 0x42, 0xaa, 0xfa, 0x75, 0x4f, 0x0d, 0x47, 0x3e, 0xcc, 0x32, 0xa6,
	0xbc, 0xb7, 0x7c, 0x0c, 0xab, 0xc2, 0xb5, 0xab, 0xac, 0x9a, 0x43, 0xee, 0x4a, 0xc1, 0x81, 0x89,
	0x0c, 0x66, 0xa1, 0x29, 0xe8, 0xf4, 0x19, 0x55, 0xd6, 0x27, 0x7c, 0xad, 0x54, 0x15, 0x3f, 0x66,
	0xf6, 0x86, 0x8e, 0x09, 0xd6, 0x8f, 0xb9, 0xd1, 0xc4, 0x94, 0x20, 0x12, 0x96, 0x93, 0xbf, 0xe1,
	0xc7, 0xcc, 0x02, 0xbc, 0x55, 0x56, 0x73, 0xa6, 0x26, 0x29, 0x0c, 0x62, 0x39, 0x19, 0x83, 0xde,
	0x52, 0x37, 0x78, 0xb1, 0x7e, 0xcc, 0xad, 0x26, 0xa6, 0x0d, 0x24, 0xe1, 0x75, 0x41, 0xa7, 0xa1,
	0xae, 0x3c, 0x29, 0x0b, 0x7a, 0x97, 0x88, 0xfc, 0x45, 0xb8, 0xd7, 0x3a, 0x50, 0xeb, 0x1d, 0xde,
	0xce, 0x72, 0x36, 0xe4, 0x31, 0x05, 0xa6, 0x06, 0x5f, 0x8e, 0x38, 0xb0, 0x94, 0x2b, 0xb0, 0xd1,
	0xee, 0xc6, 0x7e, 0x37, 0xd8, 0x2b, 0xe7, 0x59, 0x15, 0xee, 0xed, 0x2a, 0xe4, 0x7f, 0x4e, 0x12,
	0xde, 0x6c, 0xca, 0xef, 0x4f, 0xab, 0x17, 0xb8, 0x51, 0x4a, 0xe3, 0x91, 0xe6, 0x76, 0x2e, 0xe1,
	0x9e, 0x39, 0xcf, 0x71, 0x83, 0xd3, 0xaa, 0xf5, 0x08, 0x77, 0x23, 0x29, 0x41, 0x41, 0x4e, 0xb3,
	0xfa, 0x6c, 0xec, 0x1a, 0xb6, 0x55, 0xc1, 0xce, 0x64, 0x12, 0x36, 0xd6, 0xea, 0xef, 0x83, 0xc7,
	0xf3, 0x85, 0x83, 0x4e, 0x16, 0x0e, 0xfa, 0xb3, 0x70, 0xd0, 0xd7, 0xa5, 0x63, 0x9c, 0x2c, 0x1d,
	0xe3, 0xd7, 0xd2, 0x31, 0x3e, 0xee, 0xb5, 0x36, 0x2e, 0x47, 0xd9, 0x43, 0xfd, 0x18, 0xfa, 0x53,
	0xbf, 0xba, 0x6d, 0x7a, 0xe5, 0x91, 0xa9, 0x2f, 0xd0, 0x83, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x62, 0xe7, 0x7f, 0x83, 0x03, 0x00, 0x00,
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
		size, err := m.Limits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Interpreter.MarshalToSizedBuffer(dAtA[:i])
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

func (m *Limits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Limits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Limits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxSize != nil {
		{
			size := m.MaxSize.Size()
			i -= size
			if _, err := m.MaxSize.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxResultCount != nil {
		{
			size := m.MaxResultCount.Size()
			i -= size
			if _, err := m.MaxResultCount.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxGas != nil {
		{
			size := m.MaxGas.Size()
			i -= size
			if _, err := m.MaxGas.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Interpreter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Interpreter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Interpreter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bootstrap) > 0 {
		i -= len(m.Bootstrap)
		copy(dAtA[i:], m.Bootstrap)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Bootstrap)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PredicatesBlacklist) > 0 {
		for iNdEx := len(m.PredicatesBlacklist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PredicatesBlacklist[iNdEx])
			copy(dAtA[i:], m.PredicatesBlacklist[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.PredicatesBlacklist[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PredicatesWhitelist) > 0 {
		for iNdEx := len(m.PredicatesWhitelist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PredicatesWhitelist[iNdEx])
			copy(dAtA[i:], m.PredicatesWhitelist[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.PredicatesWhitelist[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
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
	l = m.Interpreter.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Limits.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Limits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxGas != nil {
		l = m.MaxGas.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxResultCount != nil {
		l = m.MaxResultCount.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxSize != nil {
		l = m.MaxSize.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Interpreter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PredicatesWhitelist) > 0 {
		for _, s := range m.PredicatesWhitelist {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.PredicatesBlacklist) > 0 {
		for _, s := range m.PredicatesBlacklist {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.Bootstrap)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Interpreter", wireType)
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
			if err := m.Interpreter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limits", wireType)
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
			if err := m.Limits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Limits) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Limits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Limits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGas", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Uint
			m.MaxGas = &v
			if err := m.MaxGas.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxResultCount", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Uint
			m.MaxResultCount = &v
			if err := m.MaxResultCount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSize", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Uint
			m.MaxSize = &v
			if err := m.MaxSize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Interpreter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Interpreter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Interpreter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredicatesWhitelist", wireType)
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
			m.PredicatesWhitelist = append(m.PredicatesWhitelist, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PredicatesBlacklist", wireType)
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
			m.PredicatesBlacklist = append(m.PredicatesBlacklist, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bootstrap", wireType)
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
			m.Bootstrap = string(dAtA[iNdEx:postIndex])
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
