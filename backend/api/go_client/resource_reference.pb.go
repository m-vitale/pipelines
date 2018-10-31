// Copyright 2018 Google LLC
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
// source: resource_reference.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ResourceType int32

const (
	ResourceType_UNKNOWN_RESOURCE_TYPE ResourceType = 0
	ResourceType_EXPERIMENT            ResourceType = 1
	ResourceType_JOB                   ResourceType = 2
)

var ResourceType_name = map[int32]string{
	0: "UNKNOWN_RESOURCE_TYPE",
	1: "EXPERIMENT",
	2: "JOB",
}

var ResourceType_value = map[string]int32{
	"UNKNOWN_RESOURCE_TYPE": 0,
	"EXPERIMENT":            1,
	"JOB":                   2,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52a3fd5cc3dcce29, []int{0}
}

type Relationship int32

const (
	Relationship_UNKNOWN_RELATIONSHIP Relationship = 0
	Relationship_OWNER                Relationship = 1
	Relationship_CREATOR              Relationship = 2
)

var Relationship_name = map[int32]string{
	0: "UNKNOWN_RELATIONSHIP",
	1: "OWNER",
	2: "CREATOR",
}

var Relationship_value = map[string]int32{
	"UNKNOWN_RELATIONSHIP": 0,
	"OWNER":                1,
	"CREATOR":              2,
}

func (x Relationship) String() string {
	return proto.EnumName(Relationship_name, int32(x))
}

func (Relationship) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52a3fd5cc3dcce29, []int{1}
}

type ResourceKey struct {
	// The type of the resource that referred to.
	Type ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=api.ResourceType" json:"type,omitempty"`
	// The ID of the resource that referred to.
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceKey) Reset()         { *m = ResourceKey{} }
func (m *ResourceKey) String() string { return proto.CompactTextString(m) }
func (*ResourceKey) ProtoMessage()    {}
func (*ResourceKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a3fd5cc3dcce29, []int{0}
}

func (m *ResourceKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceKey.Unmarshal(m, b)
}
func (m *ResourceKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceKey.Marshal(b, m, deterministic)
}
func (m *ResourceKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceKey.Merge(m, src)
}
func (m *ResourceKey) XXX_Size() int {
	return xxx_messageInfo_ResourceKey.Size(m)
}
func (m *ResourceKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceKey.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceKey proto.InternalMessageInfo

func (m *ResourceKey) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_UNKNOWN_RESOURCE_TYPE
}

func (m *ResourceKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ResourceReference struct {
	Key *ResourceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Required field. The relationship from referred resource to the object.
	Relationship         Relationship `protobuf:"varint,2,opt,name=relationship,proto3,enum=api.Relationship" json:"relationship,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResourceReference) Reset()         { *m = ResourceReference{} }
func (m *ResourceReference) String() string { return proto.CompactTextString(m) }
func (*ResourceReference) ProtoMessage()    {}
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a3fd5cc3dcce29, []int{1}
}

func (m *ResourceReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceReference.Unmarshal(m, b)
}
func (m *ResourceReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceReference.Marshal(b, m, deterministic)
}
func (m *ResourceReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceReference.Merge(m, src)
}
func (m *ResourceReference) XXX_Size() int {
	return xxx_messageInfo_ResourceReference.Size(m)
}
func (m *ResourceReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceReference.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceReference proto.InternalMessageInfo

func (m *ResourceReference) GetKey() *ResourceKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ResourceReference) GetRelationship() Relationship {
	if m != nil {
		return m.Relationship
	}
	return Relationship_UNKNOWN_RELATIONSHIP
}

func init() {
	proto.RegisterEnum("api.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("api.Relationship", Relationship_name, Relationship_value)
	proto.RegisterType((*ResourceKey)(nil), "api.ResourceKey")
	proto.RegisterType((*ResourceReference)(nil), "api.ResourceReference")
}

func init() { proto.RegisterFile("resource_reference.proto", fileDescriptor_52a3fd5cc3dcce29) }

var fileDescriptor_52a3fd5cc3dcce29 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x46, 0x9b, 0x54, 0x2d, 0x9d, 0x96, 0xb2, 0x5d, 0x14, 0xe2, 0xad, 0x04, 0x84, 0xd2, 0x43,
	0x0e, 0x15, 0xef, 0xb6, 0x75, 0xc1, 0x18, 0xdd, 0x0d, 0xd3, 0x94, 0xea, 0x29, 0xc4, 0x76, 0xc4,
	0x45, 0x49, 0x96, 0x6d, 0x3c, 0xe4, 0xdf, 0x4b, 0x83, 0x21, 0xf6, 0xfc, 0x1e, 0xef, 0x1b, 0x06,
	0x3c, 0x4b, 0x87, 0xe2, 0xc7, 0xee, 0x28, 0xb5, 0xf4, 0x41, 0x96, 0xf2, 0x1d, 0x05, 0xc6, 0x16,
	0x65, 0xc1, 0xbb, 0x99, 0xd1, 0xfe, 0x03, 0x0c, 0xf0, 0x4f, 0x88, 0xa8, 0xe2, 0x37, 0x70, 0x56,
	0x56, 0x86, 0x3c, 0x67, 0xe2, 0x4c, 0x47, 0xf3, 0x71, 0x90, 0x19, 0x1d, 0x34, 0x3c, 0xa9, 0x0c,
	0x61, 0x8d, 0xf9, 0x08, 0x5c, 0xbd, 0xf7, 0xdc, 0x89, 0x33, 0xed, 0xa3, 0xab, 0xf7, 0x7e, 0x0e,
	0xe3, 0xc6, 0xc2, 0x66, 0x85, 0xfb, 0xd0, 0xfd, 0xa2, 0xaa, 0x4e, 0x0d, 0xe6, 0xec, 0x24, 0x15,
	0x51, 0x85, 0x47, 0xc8, 0xef, 0x60, 0x68, 0xe9, 0x3b, 0x2b, 0x75, 0x91, 0x1f, 0x3e, 0xb5, 0xa9,
	0x93, 0xed, 0x6e, 0x0b, 0xf0, 0x44, 0x9b, 0x2d, 0x61, 0xf8, 0xff, 0x2a, 0x7e, 0x0d, 0x57, 0x1b,
	0x19, 0x49, 0xb5, 0x95, 0x29, 0x8a, 0xb5, 0xda, 0xe0, 0x4a, 0xa4, 0xc9, 0x5b, 0x2c, 0x58, 0x87,
	0x8f, 0x00, 0xc4, 0x6b, 0x2c, 0x30, 0x7c, 0x11, 0x32, 0x61, 0x0e, 0xef, 0x41, 0xf7, 0x49, 0x2d,
	0x99, 0x3b, 0xbb, 0x3f, 0x36, 0xda, 0x26, 0xf7, 0xe0, 0xb2, 0x6d, 0x3c, 0x2f, 0x92, 0x50, 0xc9,
	0xf5, 0x63, 0x18, 0xb3, 0x0e, 0xef, 0xc3, 0xb9, 0xda, 0x4a, 0x81, 0xcc, 0xe1, 0x03, 0xe8, 0xad,
	0x50, 0x2c, 0x12, 0x85, 0xcc, 0x7d, 0xbf, 0xa8, 0xff, 0x78, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x35, 0xb6, 0x7b, 0xd3, 0x63, 0x01, 0x00, 0x00,
}
