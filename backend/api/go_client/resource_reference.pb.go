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
// source: backend/api/resource_reference.proto

package go_client // import "github.com/kubeflow/pipelines/backend/api/go_client"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	return fileDescriptor_resource_reference_7070bd2b275e38fb, []int{0}
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
	return fileDescriptor_resource_reference_7070bd2b275e38fb, []int{1}
}

type ResourceKey struct {
	Type                 ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=api.ResourceType" json:"type,omitempty"`
	Id                   string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResourceKey) Reset()         { *m = ResourceKey{} }
func (m *ResourceKey) String() string { return proto.CompactTextString(m) }
func (*ResourceKey) ProtoMessage()    {}
func (*ResourceKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_reference_7070bd2b275e38fb, []int{0}
}
func (m *ResourceKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceKey.Unmarshal(m, b)
}
func (m *ResourceKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceKey.Marshal(b, m, deterministic)
}
func (dst *ResourceKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceKey.Merge(dst, src)
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
	Key                  *ResourceKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Relationship         Relationship `protobuf:"varint,2,opt,name=relationship,proto3,enum=api.Relationship" json:"relationship,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResourceReference) Reset()         { *m = ResourceReference{} }
func (m *ResourceReference) String() string { return proto.CompactTextString(m) }
func (*ResourceReference) ProtoMessage()    {}
func (*ResourceReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_reference_7070bd2b275e38fb, []int{1}
}
func (m *ResourceReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceReference.Unmarshal(m, b)
}
func (m *ResourceReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceReference.Marshal(b, m, deterministic)
}
func (dst *ResourceReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceReference.Merge(dst, src)
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
	proto.RegisterType((*ResourceKey)(nil), "api.ResourceKey")
	proto.RegisterType((*ResourceReference)(nil), "api.ResourceReference")
	proto.RegisterEnum("api.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("api.Relationship", Relationship_name, Relationship_value)
}

func init() {
	proto.RegisterFile("backend/api/resource_reference.proto", fileDescriptor_resource_reference_7070bd2b275e38fb)
}

var fileDescriptor_resource_reference_7070bd2b275e38fb = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xd1, 0x6b, 0xf2, 0x30,
	0x14, 0x47, 0x6d, 0xfd, 0xbe, 0x89, 0x57, 0x91, 0x18, 0x36, 0x70, 0x6f, 0x22, 0x1b, 0x88, 0x0f,
	0x2d, 0x28, 0xbe, 0x4f, 0x5d, 0x60, 0xce, 0xad, 0x95, 0x58, 0x71, 0xdb, 0x4b, 0x69, 0xeb, 0x55,
	0x83, 0x5d, 0x13, 0x62, 0x65, 0xf4, 0xbf, 0x1f, 0x96, 0x49, 0xf5, 0xf9, 0x1c, 0xce, 0x2f, 0xe4,
	0xc2, 0x43, 0x18, 0x44, 0x7b, 0x4c, 0xd6, 0x76, 0xa0, 0x84, 0xad, 0xf1, 0x20, 0x8f, 0x3a, 0x42,
	0x5f, 0xe3, 0x06, 0x35, 0x26, 0x11, 0x5a, 0x4a, 0xcb, 0x54, 0xd2, 0x72, 0xa0, 0x44, 0xe7, 0x19,
	0x6a, 0xfc, 0x4f, 0x98, 0x61, 0x46, 0x1f, 0xe1, 0x5f, 0x9a, 0x29, 0x6c, 0x19, 0x6d, 0xa3, 0xdb,
	0xe8, 0x37, 0xad, 0x40, 0x09, 0xeb, 0xcc, 0xbd, 0x4c, 0x21, 0xcf, 0x31, 0x6d, 0x80, 0x29, 0xd6,
	0x2d, 0xb3, 0x6d, 0x74, 0xab, 0xdc, 0x14, 0xeb, 0x4e, 0x02, 0xcd, 0xb3, 0xc5, 0xcf, 0x2b, 0xb4,
	0x03, 0xe5, 0x3d, 0x66, 0x79, 0xaa, 0xd6, 0x27, 0x57, 0xa9, 0x19, 0x66, 0xfc, 0x04, 0xe9, 0x10,
	0xea, 0x1a, 0xe3, 0x20, 0x15, 0x32, 0x39, 0xec, 0x84, 0xca, 0x93, 0xc5, 0x6e, 0x01, 0xf8, 0x95,
	0xd6, 0x1b, 0x43, 0xfd, 0xf2, 0x55, 0xf4, 0x1e, 0xee, 0x96, 0xce, 0xcc, 0x71, 0x57, 0x8e, 0xcf,
	0xd9, 0xc2, 0x5d, 0xf2, 0x09, 0xf3, 0xbd, 0xcf, 0x39, 0x23, 0x25, 0xda, 0x00, 0x60, 0x1f, 0x73,
	0xc6, 0xa7, 0xef, 0xcc, 0xf1, 0x88, 0x41, 0x2b, 0x50, 0x7e, 0x75, 0xc7, 0xc4, 0xec, 0x3d, 0x9d,
	0x1a, 0x45, 0x93, 0xb6, 0xe0, 0xb6, 0x68, 0xbc, 0x8d, 0xbc, 0xa9, 0xeb, 0x2c, 0x5e, 0xa6, 0x73,
	0x52, 0xa2, 0x55, 0xf8, 0xef, 0xae, 0x1c, 0xc6, 0x89, 0x41, 0x6b, 0x50, 0x99, 0x70, 0x36, 0xf2,
	0x5c, 0x4e, 0xcc, 0xf1, 0xf0, 0x6b, 0xb0, 0x15, 0xe9, 0xee, 0x18, 0x5a, 0x91, 0xfc, 0xb6, 0xf7,
	0xc7, 0x10, 0x37, 0xb1, 0xfc, 0xb1, 0x95, 0x50, 0x18, 0x8b, 0x04, 0x0f, 0xf6, 0xe5, 0x19, 0xb6,
	0xd2, 0x8f, 0x62, 0x81, 0x49, 0x1a, 0xde, 0xe4, 0xdf, 0x3f, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0x12, 0xb7, 0x99, 0xa6, 0x01, 0x00, 0x00,
}
