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
// source: backend/api/report.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

type ReportWorkflowRequest struct {
	Workflow             string   `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportWorkflowRequest) Reset()         { *m = ReportWorkflowRequest{} }
func (m *ReportWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*ReportWorkflowRequest) ProtoMessage()    {}
func (*ReportWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_a941d109d085a5b5, []int{0}
}
func (m *ReportWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportWorkflowRequest.Unmarshal(m, b)
}
func (m *ReportWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportWorkflowRequest.Marshal(b, m, deterministic)
}
func (dst *ReportWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportWorkflowRequest.Merge(dst, src)
}
func (m *ReportWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_ReportWorkflowRequest.Size(m)
}
func (m *ReportWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportWorkflowRequest proto.InternalMessageInfo

func (m *ReportWorkflowRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

type ReportScheduledWorkflowRequest struct {
	ScheduledWorkflow    string   `protobuf:"bytes,1,opt,name=scheduled_workflow,json=scheduledWorkflow,proto3" json:"scheduled_workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportScheduledWorkflowRequest) Reset()         { *m = ReportScheduledWorkflowRequest{} }
func (m *ReportScheduledWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*ReportScheduledWorkflowRequest) ProtoMessage()    {}
func (*ReportScheduledWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_report_a941d109d085a5b5, []int{1}
}
func (m *ReportScheduledWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportScheduledWorkflowRequest.Unmarshal(m, b)
}
func (m *ReportScheduledWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportScheduledWorkflowRequest.Marshal(b, m, deterministic)
}
func (dst *ReportScheduledWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportScheduledWorkflowRequest.Merge(dst, src)
}
func (m *ReportScheduledWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_ReportScheduledWorkflowRequest.Size(m)
}
func (m *ReportScheduledWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportScheduledWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportScheduledWorkflowRequest proto.InternalMessageInfo

func (m *ReportScheduledWorkflowRequest) GetScheduledWorkflow() string {
	if m != nil {
		return m.ScheduledWorkflow
	}
	return ""
}

func init() {
	proto.RegisterType((*ReportWorkflowRequest)(nil), "api.ReportWorkflowRequest")
	proto.RegisterType((*ReportScheduledWorkflowRequest)(nil), "api.ReportScheduledWorkflowRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportServiceClient interface {
	ReportWorkflow(ctx context.Context, in *ReportWorkflowRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReportScheduledWorkflow(ctx context.Context, in *ReportScheduledWorkflowRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type reportServiceClient struct {
	cc *grpc.ClientConn
}

func NewReportServiceClient(cc *grpc.ClientConn) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) ReportWorkflow(ctx context.Context, in *ReportWorkflowRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ReportService/ReportWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) ReportScheduledWorkflow(ctx context.Context, in *ReportScheduledWorkflowRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ReportService/ReportScheduledWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
type ReportServiceServer interface {
	ReportWorkflow(context.Context, *ReportWorkflowRequest) (*empty.Empty, error)
	ReportScheduledWorkflow(context.Context, *ReportScheduledWorkflowRequest) (*empty.Empty, error)
}

func RegisterReportServiceServer(s *grpc.Server, srv ReportServiceServer) {
	s.RegisterService(&_ReportService_serviceDesc, srv)
}

func _ReportService_ReportWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).ReportWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ReportService/ReportWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).ReportWorkflow(ctx, req.(*ReportWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_ReportScheduledWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportScheduledWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).ReportScheduledWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ReportService/ReportScheduledWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).ReportScheduledWorkflow(ctx, req.(*ReportScheduledWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportWorkflow",
			Handler:    _ReportService_ReportWorkflow_Handler,
		},
		{
			MethodName: "ReportScheduledWorkflow",
			Handler:    _ReportService_ReportScheduledWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/report.proto",
}

func init() { proto.RegisterFile("backend/api/report.proto", fileDescriptor_report_a941d109d085a5b5) }

var fileDescriptor_report_a941d109d085a5b5 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0xd1, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0x05, 0xcb, 0x26, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x94,
	0x48, 0x49, 0x43, 0x65, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88,
	0xa4, 0x92, 0x31, 0x97, 0x68, 0x10, 0xd8, 0xbc, 0xf0, 0xfc, 0xa2, 0xec, 0xb4, 0x9c, 0xfc, 0xf2,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x8e, 0x72, 0xa8, 0x90, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xaf, 0xe4, 0xcf, 0x25, 0x07, 0xd1, 0x14, 0x9c, 0x9c, 0x91,
	0x9a, 0x52, 0x9a, 0x93, 0x9a, 0x82, 0xae, 0x5b, 0x97, 0x4b, 0xa8, 0x18, 0x26, 0x17, 0x8f, 0x66,
	0x8e, 0x60, 0x31, 0xba, 0x2e, 0xa3, 0xe9, 0x4c, 0x5c, 0xbc, 0x50, 0x13, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0x85, 0xf2, 0xb9, 0xf8, 0x50, 0xdd, 0x25, 0x24, 0xa5, 0x97, 0x58, 0x90, 0xa9, 0x87,
	0xd5, 0xb1, 0x52, 0x62, 0x7a, 0x10, 0x3f, 0xea, 0xc1, 0xfc, 0xa8, 0xe7, 0x0a, 0xf2, 0xa3, 0x92,
	0x66, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x94, 0x95, 0xc4, 0x41, 0x41, 0x53, 0xac, 0x5f, 0x66, 0x98,
	0x94, 0x5a, 0x92, 0x68, 0xa8, 0x0f, 0x73, 0x50, 0xb1, 0x15, 0xdc, 0x4f, 0x42, 0x53, 0x18, 0xb9,
	0xc4, 0x71, 0x78, 0x4a, 0x48, 0x19, 0xc9, 0x6a, 0x5c, 0x5e, 0xc6, 0xe9, 0x06, 0x1b, 0xb0, 0x1b,
	0xcc, 0x94, 0x14, 0x50, 0xdd, 0x00, 0x0f, 0x04, 0x84, 0x63, 0xb0, 0x04, 0x59, 0x12, 0x1b, 0xd8,
	0x34, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xa9, 0x38, 0xc4, 0x02, 0x02, 0x00, 0x00,
}
