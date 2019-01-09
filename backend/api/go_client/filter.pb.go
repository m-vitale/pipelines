// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend/api/filter.proto

package go_client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Predicate_Op int32

const (
	Predicate_UNKNOWN             Predicate_Op = 0
	Predicate_EQUALS              Predicate_Op = 1
	Predicate_NOT_EQUALS          Predicate_Op = 2
	Predicate_GREATER_THAN        Predicate_Op = 3
	Predicate_GREATER_THAN_EQUALS Predicate_Op = 5
	Predicate_LESS_THAN           Predicate_Op = 6
	Predicate_LESS_THAN_EQUALS    Predicate_Op = 7
	Predicate_IN                  Predicate_Op = 8
	Predicate_IS_SUBSTRING        Predicate_Op = 9
)

var Predicate_Op_name = map[int32]string{
	0: "UNKNOWN",
	1: "EQUALS",
	2: "NOT_EQUALS",
	3: "GREATER_THAN",
	5: "GREATER_THAN_EQUALS",
	6: "LESS_THAN",
	7: "LESS_THAN_EQUALS",
	8: "IN",
	9: "IS_SUBSTRING",
}
var Predicate_Op_value = map[string]int32{
	"UNKNOWN":             0,
	"EQUALS":              1,
	"NOT_EQUALS":          2,
	"GREATER_THAN":        3,
	"GREATER_THAN_EQUALS": 5,
	"LESS_THAN":           6,
	"LESS_THAN_EQUALS":    7,
	"IN":                  8,
	"IS_SUBSTRING":        9,
}

func (x Predicate_Op) String() string {
	return proto.EnumName(Predicate_Op_name, int32(x))
}
func (Predicate_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_filter_836b9ddef08e4ec1, []int{0, 0}
}

type Predicate struct {
	Op  Predicate_Op `protobuf:"varint,1,opt,name=op,proto3,enum=api.Predicate_Op" json:"op,omitempty"`
	Key string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Predicate_IntValue
	//	*Predicate_LongValue
	//	*Predicate_StringValue
	//	*Predicate_TimestampValue
	//	*Predicate_IntValues
	//	*Predicate_LongValues
	//	*Predicate_StringValues
	Value                isPredicate_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Predicate) Reset()         { *m = Predicate{} }
func (m *Predicate) String() string { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()    {}
func (*Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_836b9ddef08e4ec1, []int{0}
}
func (m *Predicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Predicate.Unmarshal(m, b)
}
func (m *Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Predicate.Marshal(b, m, deterministic)
}
func (dst *Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predicate.Merge(dst, src)
}
func (m *Predicate) XXX_Size() int {
	return xxx_messageInfo_Predicate.Size(m)
}
func (m *Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Predicate proto.InternalMessageInfo

func (m *Predicate) GetOp() Predicate_Op {
	if m != nil {
		return m.Op
	}
	return Predicate_UNKNOWN
}

func (m *Predicate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isPredicate_Value interface {
	isPredicate_Value()
}

type Predicate_IntValue struct {
	IntValue int32 `protobuf:"varint,3,opt,name=int_value,json=intValue,proto3,oneof"`
}

type Predicate_LongValue struct {
	LongValue int64 `protobuf:"varint,4,opt,name=long_value,json=longValue,proto3,oneof"`
}

type Predicate_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Predicate_TimestampValue struct {
	TimestampValue *timestamp.Timestamp `protobuf:"bytes,6,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Predicate_IntValues struct {
	IntValues *IntValues `protobuf:"bytes,7,opt,name=int_values,json=intValues,proto3,oneof"`
}

type Predicate_LongValues struct {
	LongValues *LongValues `protobuf:"bytes,8,opt,name=long_values,json=longValues,proto3,oneof"`
}

type Predicate_StringValues struct {
	StringValues *StringValues `protobuf:"bytes,9,opt,name=string_values,json=stringValues,proto3,oneof"`
}

func (*Predicate_IntValue) isPredicate_Value() {}

func (*Predicate_LongValue) isPredicate_Value() {}

func (*Predicate_StringValue) isPredicate_Value() {}

func (*Predicate_TimestampValue) isPredicate_Value() {}

func (*Predicate_IntValues) isPredicate_Value() {}

func (*Predicate_LongValues) isPredicate_Value() {}

func (*Predicate_StringValues) isPredicate_Value() {}

func (m *Predicate) GetValue() isPredicate_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Predicate) GetIntValue() int32 {
	if x, ok := m.GetValue().(*Predicate_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *Predicate) GetLongValue() int64 {
	if x, ok := m.GetValue().(*Predicate_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (m *Predicate) GetStringValue() string {
	if x, ok := m.GetValue().(*Predicate_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Predicate) GetTimestampValue() *timestamp.Timestamp {
	if x, ok := m.GetValue().(*Predicate_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

func (m *Predicate) GetIntValues() *IntValues {
	if x, ok := m.GetValue().(*Predicate_IntValues); ok {
		return x.IntValues
	}
	return nil
}

func (m *Predicate) GetLongValues() *LongValues {
	if x, ok := m.GetValue().(*Predicate_LongValues); ok {
		return x.LongValues
	}
	return nil
}

func (m *Predicate) GetStringValues() *StringValues {
	if x, ok := m.GetValue().(*Predicate_StringValues); ok {
		return x.StringValues
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Predicate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Predicate_OneofMarshaler, _Predicate_OneofUnmarshaler, _Predicate_OneofSizer, []interface{}{
		(*Predicate_IntValue)(nil),
		(*Predicate_LongValue)(nil),
		(*Predicate_StringValue)(nil),
		(*Predicate_TimestampValue)(nil),
		(*Predicate_IntValues)(nil),
		(*Predicate_LongValues)(nil),
		(*Predicate_StringValues)(nil),
	}
}

func _Predicate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Predicate)
	// value
	switch x := m.Value.(type) {
	case *Predicate_IntValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *Predicate_LongValue:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.LongValue))
	case *Predicate_StringValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Predicate_TimestampValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampValue); err != nil {
			return err
		}
	case *Predicate_IntValues:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IntValues); err != nil {
			return err
		}
	case *Predicate_LongValues:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LongValues); err != nil {
			return err
		}
	case *Predicate_StringValues:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringValues); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Predicate.Value has unexpected type %T", x)
	}
	return nil
}

func _Predicate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Predicate)
	switch tag {
	case 3: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Predicate_IntValue{int32(x)}
		return true, err
	case 4: // value.long_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Predicate_LongValue{int64(x)}
		return true, err
	case 5: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Predicate_StringValue{x}
		return true, err
	case 6: // value.timestamp_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.Value = &Predicate_TimestampValue{msg}
		return true, err
	case 7: // value.int_values
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IntValues)
		err := b.DecodeMessage(msg)
		m.Value = &Predicate_IntValues{msg}
		return true, err
	case 8: // value.long_values
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LongValues)
		err := b.DecodeMessage(msg)
		m.Value = &Predicate_LongValues{msg}
		return true, err
	case 9: // value.string_values
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringValues)
		err := b.DecodeMessage(msg)
		m.Value = &Predicate_StringValues{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Predicate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Predicate)
	// value
	switch x := m.Value.(type) {
	case *Predicate_IntValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntValue))
	case *Predicate_LongValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.LongValue))
	case *Predicate_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Predicate_TimestampValue:
		s := proto.Size(x.TimestampValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Predicate_IntValues:
		s := proto.Size(x.IntValues)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Predicate_LongValues:
		s := proto.Size(x.LongValues)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Predicate_StringValues:
		s := proto.Size(x.StringValues)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type IntValues struct {
	Values               []int32  `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntValues) Reset()         { *m = IntValues{} }
func (m *IntValues) String() string { return proto.CompactTextString(m) }
func (*IntValues) ProtoMessage()    {}
func (*IntValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_836b9ddef08e4ec1, []int{1}
}
func (m *IntValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntValues.Unmarshal(m, b)
}
func (m *IntValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntValues.Marshal(b, m, deterministic)
}
func (dst *IntValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntValues.Merge(dst, src)
}
func (m *IntValues) XXX_Size() int {
	return xxx_messageInfo_IntValues.Size(m)
}
func (m *IntValues) XXX_DiscardUnknown() {
	xxx_messageInfo_IntValues.DiscardUnknown(m)
}

var xxx_messageInfo_IntValues proto.InternalMessageInfo

func (m *IntValues) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type StringValues struct {
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValues) Reset()         { *m = StringValues{} }
func (m *StringValues) String() string { return proto.CompactTextString(m) }
func (*StringValues) ProtoMessage()    {}
func (*StringValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_836b9ddef08e4ec1, []int{2}
}
func (m *StringValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValues.Unmarshal(m, b)
}
func (m *StringValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValues.Marshal(b, m, deterministic)
}
func (dst *StringValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValues.Merge(dst, src)
}
func (m *StringValues) XXX_Size() int {
	return xxx_messageInfo_StringValues.Size(m)
}
func (m *StringValues) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValues.DiscardUnknown(m)
}

var xxx_messageInfo_StringValues proto.InternalMessageInfo

func (m *StringValues) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type LongValues struct {
	Values               []int64  `protobuf:"varint,3,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongValues) Reset()         { *m = LongValues{} }
func (m *LongValues) String() string { return proto.CompactTextString(m) }
func (*LongValues) ProtoMessage()    {}
func (*LongValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_836b9ddef08e4ec1, []int{3}
}
func (m *LongValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongValues.Unmarshal(m, b)
}
func (m *LongValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongValues.Marshal(b, m, deterministic)
}
func (dst *LongValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongValues.Merge(dst, src)
}
func (m *LongValues) XXX_Size() int {
	return xxx_messageInfo_LongValues.Size(m)
}
func (m *LongValues) XXX_DiscardUnknown() {
	xxx_messageInfo_LongValues.DiscardUnknown(m)
}

var xxx_messageInfo_LongValues proto.InternalMessageInfo

func (m *LongValues) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Filter struct {
	Predicates           []*Predicate `protobuf:"bytes,1,rep,name=predicates,proto3" json:"predicates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_filter_836b9ddef08e4ec1, []int{4}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (dst *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(dst, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPredicates() []*Predicate {
	if m != nil {
		return m.Predicates
	}
	return nil
}

func init() {
	proto.RegisterType((*Predicate)(nil), "api.Predicate")
	proto.RegisterType((*IntValues)(nil), "api.IntValues")
	proto.RegisterType((*StringValues)(nil), "api.StringValues")
	proto.RegisterType((*LongValues)(nil), "api.LongValues")
	proto.RegisterType((*Filter)(nil), "api.Filter")
	proto.RegisterEnum("api.Predicate_Op", Predicate_Op_name, Predicate_Op_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DummyFilterServiceClient is the client API for DummyFilterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DummyFilterServiceClient interface {
	GetFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Filter, error)
}

type dummyFilterServiceClient struct {
	cc *grpc.ClientConn
}

func NewDummyFilterServiceClient(cc *grpc.ClientConn) DummyFilterServiceClient {
	return &dummyFilterServiceClient{cc}
}

func (c *dummyFilterServiceClient) GetFilter(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Filter, error) {
	out := new(Filter)
	err := c.cc.Invoke(ctx, "/api.DummyFilterService/GetFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyFilterServiceServer is the server API for DummyFilterService service.
type DummyFilterServiceServer interface {
	GetFilter(context.Context, *Filter) (*Filter, error)
}

func RegisterDummyFilterServiceServer(s *grpc.Server, srv DummyFilterServiceServer) {
	s.RegisterService(&_DummyFilterService_serviceDesc, srv)
}

func _DummyFilterService_GetFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyFilterServiceServer).GetFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DummyFilterService/GetFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyFilterServiceServer).GetFilter(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

var _DummyFilterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DummyFilterService",
	HandlerType: (*DummyFilterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFilter",
			Handler:    _DummyFilterService_GetFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/filter.proto",
}

func init() { proto.RegisterFile("backend/api/filter.proto", fileDescriptor_filter_836b9ddef08e4ec1) }

var fileDescriptor_filter_836b9ddef08e4ec1 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xfd, 0x27, 0x18, 0x3c, 0x10, 0xe2, 0x6e, 0xab, 0xd6, 0x42, 0xad, 0xe2, 0x92, 0xaa,
	0xf5, 0xc9, 0x48, 0xf4, 0x92, 0x4b, 0x0f, 0x44, 0xa5, 0x80, 0x8a, 0x4c, 0xbb, 0x86, 0xf6, 0x88,
	0x0c, 0xd9, 0xa0, 0x55, 0xc0, 0x5e, 0xb1, 0x4b, 0xa4, 0x3c, 0x49, 0xdf, 0xa2, 0xcf, 0x18, 0x79,
	0xd7, 0x5e, 0x7c, 0x63, 0x66, 0x7e, 0xdf, 0xf0, 0x7d, 0xeb, 0x01, 0x7f, 0x93, 0x6e, 0x1f, 0x49,
	0x76, 0x3f, 0x48, 0x19, 0x1d, 0x3c, 0xd0, 0xbd, 0x20, 0xc7, 0x88, 0x1d, 0x73, 0x91, 0x23, 0x3b,
	0x65, 0xb4, 0xf7, 0x7e, 0x97, 0xe7, 0xbb, 0x3d, 0x91, 0xd3, 0x34, 0xcb, 0x72, 0x91, 0x0a, 0x9a,
	0x67, 0x5c, 0x21, 0xbd, 0xeb, 0x72, 0x2a, 0xab, 0xcd, 0xe9, 0x61, 0x20, 0xe8, 0x81, 0x70, 0x91,
	0x1e, 0x98, 0x02, 0xfa, 0xff, 0x2f, 0xc0, 0xfd, 0x75, 0x24, 0xf7, 0x74, 0x9b, 0x0a, 0x82, 0x3e,
	0x82, 0x95, 0x33, 0xdf, 0x0c, 0xcc, 0xb0, 0x3b, 0x7c, 0x15, 0xa5, 0x8c, 0x46, 0x7a, 0x16, 0x2d,
	0x18, 0xb6, 0x72, 0x86, 0x3c, 0xb0, 0x1f, 0xc9, 0xb3, 0x6f, 0x05, 0x66, 0xe8, 0xe2, 0xe2, 0x27,
	0xfa, 0x00, 0x2e, 0xcd, 0xc4, 0xfa, 0x29, 0xdd, 0x9f, 0x88, 0x6f, 0x07, 0x66, 0xd8, 0x98, 0x1a,
	0xb8, 0x45, 0x33, 0xf1, 0xa7, 0xe8, 0xa0, 0x6b, 0x80, 0x7d, 0x9e, 0xed, 0xca, 0xf9, 0x45, 0x60,
	0x86, 0xf6, 0xd4, 0xc0, 0x6e, 0xd1, 0x53, 0xc0, 0x0d, 0x74, 0xb8, 0x38, 0x52, 0x8d, 0x34, 0x8a,
	0xd5, 0x53, 0x03, 0xb7, 0x55, 0x57, 0x41, 0x63, 0xb8, 0xd2, 0xd6, 0x4b, 0xce, 0x09, 0xcc, 0xb0,
	0x3d, 0xec, 0x45, 0x2a, 0x62, 0x54, 0x45, 0x8c, 0x96, 0x15, 0x37, 0x35, 0x70, 0x57, 0x8b, 0xd4,
	0x9a, 0x01, 0x80, 0xf6, 0xca, 0xfd, 0xa6, 0xdc, 0xd0, 0x95, 0x41, 0x67, 0xa5, 0x5f, 0x5e, 0x98,
	0xab, 0xcc, 0x73, 0x34, 0x84, 0xf6, 0xd9, 0x3d, 0xf7, 0x5b, 0x52, 0x71, 0x25, 0x15, 0xf3, 0x2a,
	0x41, 0x21, 0x01, 0x9d, 0x87, 0xa3, 0x5b, 0xb8, 0xac, 0x07, 0xe2, 0xbe, 0x2b, 0x55, 0xea, 0x41,
	0x93, 0x73, 0xa8, 0x42, 0xd7, 0xa9, 0x85, 0xe4, 0xfd, 0x7f, 0x26, 0x58, 0x0b, 0x86, 0xda, 0xd0,
	0x5c, 0xc5, 0x3f, 0xe3, 0xc5, 0xdf, 0xd8, 0x33, 0x10, 0x80, 0x33, 0xfe, 0xbd, 0x1a, 0xcd, 0x13,
	0xcf, 0x44, 0x5d, 0x80, 0x78, 0xb1, 0x5c, 0x97, 0xb5, 0x85, 0x3c, 0xe8, 0x4c, 0xf0, 0x78, 0xb4,
	0x1c, 0xe3, 0xf5, 0x72, 0x3a, 0x8a, 0x3d, 0x1b, 0xbd, 0x83, 0xd7, 0xf5, 0x4e, 0x85, 0x36, 0xd0,
	0x25, 0xb8, 0xf3, 0x71, 0x92, 0x28, 0xce, 0x41, 0x6f, 0xc0, 0xd3, 0x65, 0x05, 0x35, 0x91, 0x03,
	0xd6, 0x2c, 0xf6, 0x5a, 0xc5, 0xde, 0x59, 0xb2, 0x4e, 0x56, 0x77, 0xc9, 0x12, 0xcf, 0xe2, 0x89,
	0xe7, 0xde, 0x35, 0xa1, 0x21, 0xc3, 0xf4, 0x6f, 0xc0, 0xd5, 0x4f, 0x85, 0xde, 0x82, 0x53, 0x46,
	0x34, 0x03, 0x3b, 0x6c, 0xe0, 0xb2, 0xea, 0x7f, 0x86, 0x4e, 0x3d, 0x67, 0x8d, 0xb3, 0x02, 0x3b,
	0x74, 0x35, 0xf7, 0x09, 0xe0, 0xfc, 0x8a, 0x35, 0xca, 0x0e, 0xec, 0xd0, 0xd6, 0xd4, 0x2d, 0x38,
	0x3f, 0xe4, 0xdd, 0xa3, 0x08, 0x80, 0x55, 0x07, 0xa9, 0xfe, 0xb3, 0xfa, 0x7c, 0xfa, 0x4e, 0x71,
	0x8d, 0x18, 0x7e, 0x03, 0xf4, 0xfd, 0x74, 0x38, 0x3c, 0x2b, 0x79, 0x42, 0x8e, 0x4f, 0x74, 0x4b,
	0xd0, 0x17, 0x70, 0x27, 0x44, 0x94, 0x2b, 0xdb, 0x52, 0xae, 0x8a, 0x5e, 0xbd, 0xe8, 0x1b, 0x1b,
	0x47, 0xde, 0xd4, 0xd7, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x08, 0x4b, 0x8b, 0x83, 0x03,
	0x00, 0x00,
}
