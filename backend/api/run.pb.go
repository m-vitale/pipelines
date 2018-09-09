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
// source: run.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type GetRunV2Request struct {
	RunId                string   `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRunV2Request) Reset()         { *m = GetRunV2Request{} }
func (m *GetRunV2Request) String() string { return proto.CompactTextString(m) }
func (*GetRunV2Request) ProtoMessage()    {}
func (*GetRunV2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_run_6044d52b97539d27, []int{0}
}
func (m *GetRunV2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunV2Request.Unmarshal(m, b)
}
func (m *GetRunV2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunV2Request.Marshal(b, m, deterministic)
}
func (dst *GetRunV2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunV2Request.Merge(dst, src)
}
func (m *GetRunV2Request) XXX_Size() int {
	return xxx_messageInfo_GetRunV2Request.Size(m)
}
func (m *GetRunV2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRunV2Request.DiscardUnknown(m)
}

var xxx_messageInfo_GetRunV2Request proto.InternalMessageInfo

func (m *GetRunV2Request) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type GetRunRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	RunId                string   `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRunRequest) Reset()         { *m = GetRunRequest{} }
func (m *GetRunRequest) String() string { return proto.CompactTextString(m) }
func (*GetRunRequest) ProtoMessage()    {}
func (*GetRunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_run_6044d52b97539d27, []int{1}
}
func (m *GetRunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunRequest.Unmarshal(m, b)
}
func (m *GetRunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunRequest.Marshal(b, m, deterministic)
}
func (dst *GetRunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunRequest.Merge(dst, src)
}
func (m *GetRunRequest) XXX_Size() int {
	return xxx_messageInfo_GetRunRequest.Size(m)
}
func (m *GetRunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRunRequest proto.InternalMessageInfo

func (m *GetRunRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *GetRunRequest) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type ListRunsRequest struct {
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Can be format of "field_name", "field_name asc" or "field_name des"
	// Ascending by default.
	SortBy               string   `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRunsRequest) Reset()         { *m = ListRunsRequest{} }
func (m *ListRunsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRunsRequest) ProtoMessage()    {}
func (*ListRunsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_run_6044d52b97539d27, []int{2}
}
func (m *ListRunsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsRequest.Unmarshal(m, b)
}
func (m *ListRunsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsRequest.Marshal(b, m, deterministic)
}
func (dst *ListRunsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsRequest.Merge(dst, src)
}
func (m *ListRunsRequest) XXX_Size() int {
	return xxx_messageInfo_ListRunsRequest.Size(m)
}
func (m *ListRunsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunsRequest proto.InternalMessageInfo

func (m *ListRunsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListRunsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRunsRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

type ListRunsResponse struct {
	Runs                 []*Run   `protobuf:"bytes,1,rep,name=runs,proto3" json:"runs,omitempty"`
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRunsResponse) Reset()         { *m = ListRunsResponse{} }
func (m *ListRunsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRunsResponse) ProtoMessage()    {}
func (*ListRunsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_run_6044d52b97539d27, []int{3}
}
func (m *ListRunsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsResponse.Unmarshal(m, b)
}
func (m *ListRunsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsResponse.Marshal(b, m, deterministic)
}
func (dst *ListRunsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsResponse.Merge(dst, src)
}
func (m *ListRunsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRunsResponse.Size(m)
}
func (m *ListRunsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRunsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRunsResponse proto.InternalMessageInfo

func (m *ListRunsResponse) GetRuns() []*Run {
	if m != nil {
		return m.Runs
	}
	return nil
}

func (m *ListRunsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type Run struct {
	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string               `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ScheduledAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
	Status      string               `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	JobId       string               `protobuf:"bytes,8,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// In case any error happens retrieving a run field, only run ID
	// and the error message is returned. Client has the flexibility of choosing
	// how to handle error. This is especially useful during listing call.
	Error                string   `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_run_6044d52b97539d27, []int{4}
}
func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (dst *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(dst, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Run) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Run) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Run) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Run) GetScheduledAt() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduledAt
	}
	return nil
}

func (m *Run) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Run) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Run) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RunDetail struct {
	Run                  *Run     `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
	Workflow             string   `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunDetail) Reset()         { *m = RunDetail{} }
func (m *RunDetail) String() string { return proto.CompactTextString(m) }
func (*RunDetail) ProtoMessage()    {}
func (*RunDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_run_6044d52b97539d27, []int{5}
}
func (m *RunDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunDetail.Unmarshal(m, b)
}
func (m *RunDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunDetail.Marshal(b, m, deterministic)
}
func (dst *RunDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunDetail.Merge(dst, src)
}
func (m *RunDetail) XXX_Size() int {
	return xxx_messageInfo_RunDetail.Size(m)
}
func (m *RunDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RunDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RunDetail proto.InternalMessageInfo

func (m *RunDetail) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

func (m *RunDetail) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRunV2Request)(nil), "api.GetRunV2Request")
	proto.RegisterType((*GetRunRequest)(nil), "api.GetRunRequest")
	proto.RegisterType((*ListRunsRequest)(nil), "api.ListRunsRequest")
	proto.RegisterType((*ListRunsResponse)(nil), "api.ListRunsResponse")
	proto.RegisterType((*Run)(nil), "api.Run")
	proto.RegisterType((*RunDetail)(nil), "api.RunDetail")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RunServiceClient is the client API for RunService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunServiceClient interface {
	GetRunV2(ctx context.Context, in *GetRunV2Request, opts ...grpc.CallOption) (*RunDetail, error)
	ListRuns(ctx context.Context, in *ListRunsRequest, opts ...grpc.CallOption) (*ListRunsResponse, error)
	GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*RunDetail, error)
}

type runServiceClient struct {
	cc *grpc.ClientConn
}

func NewRunServiceClient(cc *grpc.ClientConn) RunServiceClient {
	return &runServiceClient{cc}
}

func (c *runServiceClient) GetRunV2(ctx context.Context, in *GetRunV2Request, opts ...grpc.CallOption) (*RunDetail, error) {
	out := new(RunDetail)
	err := c.cc.Invoke(ctx, "/api.RunService/GetRunV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runServiceClient) ListRuns(ctx context.Context, in *ListRunsRequest, opts ...grpc.CallOption) (*ListRunsResponse, error) {
	out := new(ListRunsResponse)
	err := c.cc.Invoke(ctx, "/api.RunService/ListRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runServiceClient) GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*RunDetail, error) {
	out := new(RunDetail)
	err := c.cc.Invoke(ctx, "/api.RunService/GetRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunServiceServer is the server API for RunService service.
type RunServiceServer interface {
	GetRunV2(context.Context, *GetRunV2Request) (*RunDetail, error)
	ListRuns(context.Context, *ListRunsRequest) (*ListRunsResponse, error)
	GetRun(context.Context, *GetRunRequest) (*RunDetail, error)
}

func RegisterRunServiceServer(s *grpc.Server, srv RunServiceServer) {
	s.RegisterService(&_RunService_serviceDesc, srv)
}

func _RunService_GetRunV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).GetRunV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RunService/GetRunV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).GetRunV2(ctx, req.(*GetRunV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunService_ListRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).ListRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RunService/ListRuns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).ListRuns(ctx, req.(*ListRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunService_GetRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).GetRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RunService/GetRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).GetRun(ctx, req.(*GetRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RunService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RunService",
	HandlerType: (*RunServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRunV2",
			Handler:    _RunService_GetRunV2_Handler,
		},
		{
			MethodName: "ListRuns",
			Handler:    _RunService_ListRuns_Handler,
		},
		{
			MethodName: "GetRun",
			Handler:    _RunService_GetRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "run.proto",
}

func init() { proto.RegisterFile("run.proto", fileDescriptor_run_6044d52b97539d27) }

var fileDescriptor_run_6044d52b97539d27 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x92, 0x26, 0x8d, 0x27, 0xb4, 0x81, 0xa1, 0x85, 0xc8, 0x0d, 0xa5, 0x58, 0x08, 0x45,
	0xa8, 0x8d, 0x45, 0x38, 0x71, 0xe8, 0xa1, 0x05, 0x09, 0x55, 0xe2, 0x80, 0xdc, 0xaa, 0xe2, 0x16,
	0xd6, 0xf1, 0xc6, 0xdd, 0xd6, 0xdd, 0x35, 0xfb, 0xd1, 0xd2, 0x56, 0xbd, 0x70, 0xe0, 0x07, 0xc0,
	0x95, 0x7f, 0xc5, 0x5f, 0xe0, 0x87, 0x20, 0xaf, 0xed, 0xd4, 0xa4, 0x20, 0x4e, 0xce, 0xbc, 0x9d,
	0x79, 0xfb, 0x76, 0xde, 0x0b, 0x38, 0xd2, 0xf0, 0x61, 0x2a, 0x85, 0x16, 0xd8, 0x20, 0x29, 0x73,
	0x3b, 0x54, 0x4a, 0x21, 0x73, 0xc4, 0xed, 0xc7, 0x42, 0xc4, 0x09, 0xf5, 0x49, 0xca, 0x7c, 0xc2,
	0xb9, 0xd0, 0x44, 0x33, 0xc1, 0x55, 0x71, 0xfa, 0xb8, 0x38, 0xb5, 0x55, 0x68, 0xa6, 0xbe, 0x66,
	0xa7, 0x54, 0x69, 0x72, 0x9a, 0x16, 0x0d, 0x9b, 0xf6, 0x33, 0xd9, 0x8a, 0x29, 0xdf, 0x52, 0xe7,
	0x24, 0x8e, 0xa9, 0xf4, 0x45, 0x6a, 0x29, 0x6e, 0xd3, 0x79, 0x03, 0xe8, 0xbe, 0xa5, 0x3a, 0x30,
	0xfc, 0x70, 0x14, 0xd0, 0x4f, 0x86, 0x2a, 0x8d, 0xab, 0xd0, 0x92, 0x86, 0x8f, 0x59, 0xd4, 0xab,
	0x6d, 0xd4, 0x06, 0x4e, 0xd0, 0x94, 0x86, 0xef, 0x45, 0xde, 0x36, 0x2c, 0xe5, 0x9d, 0x95, 0xbe,
	0x63, 0x11, 0x56, 0xfa, 0x8e, 0x45, 0xb8, 0x17, 0x55, 0xc6, 0xeb, 0xd5, 0xf1, 0x29, 0x74, 0xdf,
	0x31, 0x95, 0xcd, 0xab, 0x92, 0xe0, 0x11, 0x40, 0x4a, 0x62, 0x3a, 0xd6, 0xe2, 0x84, 0xf2, 0xa2,
	0xdb, 0xc9, 0x90, 0x83, 0x0c, 0xc0, 0x35, 0xb0, 0xc5, 0x58, 0xb1, 0x4b, 0xda, 0x6b, 0x6c, 0xd4,
	0x06, 0xcd, 0xa0, 0x9d, 0x01, 0xfb, 0xec, 0x92, 0xe2, 0x43, 0x58, 0x54, 0x42, 0xea, 0x71, 0x78,
	0xd1, 0x5b, 0xb0, 0x83, 0xad, 0xac, 0xdc, 0xbd, 0xf0, 0x3e, 0xc0, 0xdd, 0x9b, 0x7b, 0x54, 0x2a,
	0xb8, 0xa2, 0xd8, 0x87, 0x05, 0x69, 0xb8, 0xea, 0xd5, 0x36, 0x1a, 0x83, 0xce, 0xa8, 0x3d, 0x24,
	0x29, 0x1b, 0x66, 0x0f, 0xb1, 0x28, 0x3e, 0x83, 0x2e, 0xa7, 0x9f, 0xf5, 0xf8, 0x96, 0x96, 0xa5,
	0x0c, 0x7e, 0x5f, 0xea, 0xf1, 0xbe, 0xd6, 0xa1, 0x11, 0x18, 0x8e, 0xcb, 0x50, 0x9f, 0xbd, 0xb9,
	0xce, 0x22, 0x44, 0x58, 0xe0, 0xe4, 0x94, 0x16, 0x43, 0xf6, 0x37, 0xf6, 0xc1, 0xc9, 0xbe, 0x2a,
	0x25, 0x93, 0x5c, 0xbb, 0x13, 0xdc, 0x00, 0xf8, 0x0a, 0x60, 0x22, 0x29, 0xd1, 0x34, 0x1a, 0x13,
	0x6d, 0xf5, 0x77, 0x46, 0xee, 0x30, 0x37, 0x76, 0x58, 0x1a, 0x3b, 0x3c, 0x28, 0x8d, 0x0d, 0x9c,
	0xa2, 0x7b, 0x47, 0xe3, 0x36, 0xdc, 0x51, 0x93, 0x23, 0x1a, 0x99, 0x24, 0x1f, 0x6e, 0xfe, 0x77,
	0xb8, 0x33, 0xeb, 0xdf, 0xd1, 0xf8, 0x00, 0x5a, 0x4a, 0x13, 0x6d, 0x54, 0xaf, 0x55, 0x6c, 0xcd,
	0x56, 0x15, 0x2f, 0xdb, 0x55, 0x2f, 0x57, 0xa0, 0x69, 0x93, 0xd9, 0x5b, 0xcc, 0x51, 0x5b, 0x78,
	0xaf, 0xc1, 0x09, 0x0c, 0x7f, 0x43, 0x35, 0x61, 0x09, 0xba, 0xd0, 0x90, 0x86, 0xdb, 0x75, 0x54,
	0x57, 0x9b, 0x81, 0xe8, 0x42, 0xfb, 0x5c, 0xc8, 0x93, 0x69, 0x22, 0xce, 0x8b, 0xed, 0xcc, 0xea,
	0xd1, 0x8f, 0x3a, 0x40, 0x60, 0xf8, 0x3e, 0x95, 0x67, 0x6c, 0x42, 0xf1, 0x10, 0xda, 0x65, 0x0e,
	0x71, 0xc5, 0xb2, 0xcc, 0xc5, 0xd2, 0x5d, 0x2e, 0xb9, 0xf3, 0x8b, 0xbd, 0xa7, 0x5f, 0x7e, 0xfe,
	0xfa, 0x5e, 0x5f, 0xc7, 0x7e, 0xf6, 0x47, 0x51, 0xfe, 0xd9, 0x0b, 0x92, 0xa4, 0x47, 0x64, 0xe4,
	0x67, 0x9e, 0xfa, 0x57, 0x79, 0x04, 0xaf, 0xf1, 0x00, 0xda, 0x65, 0x1c, 0x0a, 0xde, 0xb9, 0x14,
	0xba, 0xab, 0x73, 0x68, 0x9e, 0x19, 0x6f, 0xcd, 0xd2, 0xaf, 0xe2, 0xfd, 0xbf, 0xd0, 0xe3, 0x47,
	0x68, 0xe5, 0xf2, 0x10, 0x2b, 0x5a, 0xff, 0xa5, 0x74, 0x64, 0xa9, 0x36, 0xf1, 0xf9, 0x1c, 0xd5,
	0xb1, 0x08, 0x95, 0x7f, 0x95, 0xef, 0xfd, 0xfa, 0x4f, 0xdd, 0xbb, 0x4f, 0xbe, 0xed, 0xac, 0x07,
	0x7d, 0x58, 0x8c, 0xe8, 0x94, 0x98, 0x44, 0xe3, 0x3d, 0xec, 0xc2, 0x92, 0xdb, 0xb1, 0xcc, 0xfb,
	0xd6, 0xb3, 0xb0, 0x65, 0xcd, 0x7e, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x47, 0x70, 0x4d, 0xeb,
	0x4d, 0x04, 0x00, 0x00,
}
