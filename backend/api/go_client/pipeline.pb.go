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
// source: backend/api/pipeline.proto

package go_client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
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

type Url struct {
	PipelineUrl          string   `protobuf:"bytes,1,opt,name=pipeline_url,json=pipelineUrl,proto3" json:"pipeline_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Url) Reset()         { *m = Url{} }
func (m *Url) String() string { return proto.CompactTextString(m) }
func (*Url) ProtoMessage()    {}
func (*Url) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{0}
}
func (m *Url) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Url.Unmarshal(m, b)
}
func (m *Url) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Url.Marshal(b, m, deterministic)
}
func (dst *Url) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Url.Merge(dst, src)
}
func (m *Url) XXX_Size() int {
	return xxx_messageInfo_Url.Size(m)
}
func (m *Url) XXX_DiscardUnknown() {
	xxx_messageInfo_Url.DiscardUnknown(m)
}

var xxx_messageInfo_Url proto.InternalMessageInfo

func (m *Url) GetPipelineUrl() string {
	if m != nil {
		return m.PipelineUrl
	}
	return ""
}

type CreatePipelineRequest struct {
	Pipeline             *Pipeline `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreatePipelineRequest) Reset()         { *m = CreatePipelineRequest{} }
func (m *CreatePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineRequest) ProtoMessage()    {}
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{1}
}
func (m *CreatePipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePipelineRequest.Unmarshal(m, b)
}
func (m *CreatePipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePipelineRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePipelineRequest.Merge(dst, src)
}
func (m *CreatePipelineRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePipelineRequest.Size(m)
}
func (m *CreatePipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePipelineRequest proto.InternalMessageInfo

func (m *CreatePipelineRequest) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type GetPipelineRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPipelineRequest) Reset()         { *m = GetPipelineRequest{} }
func (m *GetPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRequest) ProtoMessage()    {}
func (*GetPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{2}
}
func (m *GetPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPipelineRequest.Unmarshal(m, b)
}
func (m *GetPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPipelineRequest.Marshal(b, m, deterministic)
}
func (dst *GetPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineRequest.Merge(dst, src)
}
func (m *GetPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_GetPipelineRequest.Size(m)
}
func (m *GetPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineRequest proto.InternalMessageInfo

func (m *GetPipelineRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListPipelinesRequest struct {
	PageToken            string   `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortBy               string   `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPipelinesRequest) Reset()         { *m = ListPipelinesRequest{} }
func (m *ListPipelinesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelinesRequest) ProtoMessage()    {}
func (*ListPipelinesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{3}
}
func (m *ListPipelinesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPipelinesRequest.Unmarshal(m, b)
}
func (m *ListPipelinesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPipelinesRequest.Marshal(b, m, deterministic)
}
func (dst *ListPipelinesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPipelinesRequest.Merge(dst, src)
}
func (m *ListPipelinesRequest) XXX_Size() int {
	return xxx_messageInfo_ListPipelinesRequest.Size(m)
}
func (m *ListPipelinesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPipelinesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPipelinesRequest proto.InternalMessageInfo

func (m *ListPipelinesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListPipelinesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPipelinesRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

func (m *ListPipelinesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type ListPipelinesResponse struct {
	Pipelines            []*Pipeline `protobuf:"bytes,1,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
	TotalSize            int32       `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	NextPageToken        string      `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListPipelinesResponse) Reset()         { *m = ListPipelinesResponse{} }
func (m *ListPipelinesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPipelinesResponse) ProtoMessage()    {}
func (*ListPipelinesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{4}
}
func (m *ListPipelinesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPipelinesResponse.Unmarshal(m, b)
}
func (m *ListPipelinesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPipelinesResponse.Marshal(b, m, deterministic)
}
func (dst *ListPipelinesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPipelinesResponse.Merge(dst, src)
}
func (m *ListPipelinesResponse) XXX_Size() int {
	return xxx_messageInfo_ListPipelinesResponse.Size(m)
}
func (m *ListPipelinesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPipelinesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPipelinesResponse proto.InternalMessageInfo

func (m *ListPipelinesResponse) GetPipelines() []*Pipeline {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

func (m *ListPipelinesResponse) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ListPipelinesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type DeletePipelineRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePipelineRequest) Reset()         { *m = DeletePipelineRequest{} }
func (m *DeletePipelineRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePipelineRequest) ProtoMessage()    {}
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{5}
}
func (m *DeletePipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePipelineRequest.Unmarshal(m, b)
}
func (m *DeletePipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePipelineRequest.Marshal(b, m, deterministic)
}
func (dst *DeletePipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePipelineRequest.Merge(dst, src)
}
func (m *DeletePipelineRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePipelineRequest.Size(m)
}
func (m *DeletePipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePipelineRequest proto.InternalMessageInfo

func (m *DeletePipelineRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTemplateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTemplateRequest) Reset()         { *m = GetTemplateRequest{} }
func (m *GetTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*GetTemplateRequest) ProtoMessage()    {}
func (*GetTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{6}
}
func (m *GetTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTemplateRequest.Unmarshal(m, b)
}
func (m *GetTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *GetTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTemplateRequest.Merge(dst, src)
}
func (m *GetTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_GetTemplateRequest.Size(m)
}
func (m *GetTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTemplateRequest proto.InternalMessageInfo

func (m *GetTemplateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTemplateResponse struct {
	Template             string   `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTemplateResponse) Reset()         { *m = GetTemplateResponse{} }
func (m *GetTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*GetTemplateResponse) ProtoMessage()    {}
func (*GetTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{7}
}
func (m *GetTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTemplateResponse.Unmarshal(m, b)
}
func (m *GetTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTemplateResponse.Marshal(b, m, deterministic)
}
func (dst *GetTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTemplateResponse.Merge(dst, src)
}
func (m *GetTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_GetTemplateResponse.Size(m)
}
func (m *GetTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTemplateResponse proto.InternalMessageInfo

func (m *GetTemplateResponse) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

type Pipeline struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Parameters           []*Parameter         `protobuf:"bytes,5,rep,name=parameters,proto3" json:"parameters,omitempty"`
	Url                  *Url                 `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Error                string               `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}
func (*Pipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_pipeline_38b26ddca62efc84, []int{8}
}
func (m *Pipeline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipeline.Unmarshal(m, b)
}
func (m *Pipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipeline.Marshal(b, m, deterministic)
}
func (dst *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(dst, src)
}
func (m *Pipeline) XXX_Size() int {
	return xxx_messageInfo_Pipeline.Size(m)
}
func (m *Pipeline) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipeline.DiscardUnknown(m)
}

var xxx_messageInfo_Pipeline proto.InternalMessageInfo

func (m *Pipeline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipeline) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Pipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pipeline) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Pipeline) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Pipeline) GetUrl() *Url {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Pipeline) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Url)(nil), "api.Url")
	proto.RegisterType((*CreatePipelineRequest)(nil), "api.CreatePipelineRequest")
	proto.RegisterType((*GetPipelineRequest)(nil), "api.GetPipelineRequest")
	proto.RegisterType((*ListPipelinesRequest)(nil), "api.ListPipelinesRequest")
	proto.RegisterType((*ListPipelinesResponse)(nil), "api.ListPipelinesResponse")
	proto.RegisterType((*DeletePipelineRequest)(nil), "api.DeletePipelineRequest")
	proto.RegisterType((*GetTemplateRequest)(nil), "api.GetTemplateRequest")
	proto.RegisterType((*GetTemplateResponse)(nil), "api.GetTemplateResponse")
	proto.RegisterType((*Pipeline)(nil), "api.Pipeline")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelineServiceClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error)
	DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error)
}

type pipelineServiceClient struct {
	cc *grpc.ClientConn
}

func NewPipelineServiceClient(cc *grpc.ClientConn) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, "/api.PipelineService/CreatePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, "/api.PipelineService/GetPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error) {
	out := new(ListPipelinesResponse)
	err := c.cc.Invoke(ctx, "/api.PipelineService/ListPipelines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) DeletePipeline(ctx context.Context, in *DeletePipelineRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.PipelineService/DeletePipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...grpc.CallOption) (*GetTemplateResponse, error) {
	out := new(GetTemplateResponse)
	err := c.cc.Invoke(ctx, "/api.PipelineService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServiceServer is the server API for PipelineService service.
type PipelineServiceServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*Pipeline, error)
	GetPipeline(context.Context, *GetPipelineRequest) (*Pipeline, error)
	ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error)
	DeletePipeline(context.Context, *DeletePipelineRequest) (*empty.Empty, error)
	GetTemplate(context.Context, *GetTemplateRequest) (*GetTemplateResponse, error)
}

func RegisterPipelineServiceServer(s *grpc.Server, srv PipelineServiceServer) {
	s.RegisterService(&_PipelineService_serviceDesc, srv)
}

func _PipelineService_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PipelineService/CreatePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PipelineService/GetPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipeline(ctx, req.(*GetPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PipelineService/ListPipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).ListPipelines(ctx, req.(*ListPipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PipelineService/DeletePipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).DeletePipeline(ctx, req.(*DeletePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PipelineService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetTemplate(ctx, req.(*GetTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PipelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineService_CreatePipeline_Handler,
		},
		{
			MethodName: "GetPipeline",
			Handler:    _PipelineService_GetPipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _PipelineService_ListPipelines_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineService_DeletePipeline_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _PipelineService_GetTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/api/pipeline.proto",
}

func init() {
	proto.RegisterFile("backend/api/pipeline.proto", fileDescriptor_pipeline_38b26ddca62efc84)
}

var fileDescriptor_pipeline_38b26ddca62efc84 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4f, 0x53, 0x13, 0x4f,
	0x10, 0xfd, 0x25, 0x81, 0x90, 0x74, 0x48, 0xa8, 0xdf, 0xf0, 0x27, 0xcb, 0x02, 0x12, 0x57, 0x0a,
	0x83, 0xca, 0xa6, 0xc0, 0x93, 0xde, 0x88, 0x5a, 0x5e, 0xb4, 0x8a, 0x0a, 0x70, 0xd1, 0x43, 0x6a,
	0x92, 0x74, 0xe2, 0xc8, 0x66, 0x77, 0x9d, 0x99, 0xa0, 0x60, 0x79, 0xd1, 0xa3, 0x37, 0x3d, 0xfb,
	0xa9, 0xfc, 0x0a, 0x5e, 0xfd, 0x0e, 0xd6, 0xf6, 0xee, 0x84, 0xfc, 0x83, 0x53, 0x32, 0xaf, 0xdf,
	0x76, 0xf7, 0xeb, 0x79, 0x3d, 0x60, 0xb7, 0x78, 0xfb, 0x1c, 0xfd, 0x4e, 0x8d, 0x87, 0xa2, 0x16,
	0x8a, 0x10, 0x3d, 0xe1, 0xa3, 0x1b, 0xca, 0x40, 0x07, 0x2c, 0xc3, 0x43, 0x61, 0x6f, 0xf6, 0x82,
	0xa0, 0xe7, 0x21, 0xc5, 0xb9, 0xef, 0x07, 0x9a, 0x6b, 0x11, 0xf8, 0x2a, 0xa6, 0xd8, 0xdb, 0x49,
	0x94, 0x4e, 0xad, 0x41, 0xb7, 0xa6, 0x45, 0x1f, 0x95, 0xe6, 0xfd, 0x30, 0x21, 0x6c, 0x4c, 0x12,
	0xb0, 0x1f, 0xea, 0x4b, 0x13, 0x1c, 0x2b, 0xce, 0x25, 0xef, 0xa3, 0x46, 0x99, 0x04, 0xcb, 0xa3,
	0x41, 0x94, 0x32, 0x30, 0x81, 0x47, 0xf4, 0xd3, 0xde, 0xef, 0xa1, 0xbf, 0xaf, 0x3e, 0xf2, 0x5e,
	0x0f, 0x65, 0x2d, 0x08, 0xa9, 0xab, 0xe9, 0x0e, 0x9d, 0x2a, 0x64, 0xce, 0xa4, 0xc7, 0xee, 0xc2,
	0xa2, 0x51, 0xd7, 0x1c, 0x48, 0xcf, 0x4a, 0x55, 0x52, 0xd5, 0x7c, 0xa3, 0x60, 0xb0, 0x33, 0xe9,
	0x39, 0x75, 0x58, 0x7d, 0x26, 0x91, 0x6b, 0x3c, 0x4e, 0xc0, 0x06, 0x7e, 0x18, 0xa0, 0xd2, 0x6c,
	0x0f, 0x72, 0x86, 0x47, 0xdf, 0x15, 0x0e, 0x8b, 0x2e, 0x0f, 0x85, 0x3b, 0xe4, 0x0d, 0xc3, 0xce,
	0x0e, 0xb0, 0x97, 0xa8, 0x27, 0x13, 0x94, 0x20, 0x2d, 0x3a, 0x49, 0xc9, 0xb4, 0xe8, 0x38, 0xdf,
	0x52, 0xb0, 0xf2, 0x4a, 0xa8, 0x21, 0x4f, 0x19, 0xe2, 0x16, 0x40, 0xc8, 0x7b, 0xd8, 0xd4, 0xc1,
	0x39, 0xfa, 0xc9, 0x07, 0xf9, 0x08, 0x39, 0x8d, 0x00, 0xb6, 0x01, 0x74, 0x68, 0x2a, 0x71, 0x85,
	0x56, 0xba, 0x92, 0xaa, 0xce, 0x37, 0x72, 0x11, 0x70, 0x22, 0xae, 0x90, 0x95, 0x61, 0x41, 0x05,
	0x52, 0x37, 0x5b, 0x97, 0x56, 0x86, 0x3e, 0xcc, 0x46, 0xc7, 0xfa, 0x25, 0x5b, 0x83, 0x6c, 0x57,
	0x78, 0x1a, 0xa5, 0x35, 0x17, 0xe3, 0xf1, 0xc9, 0xf9, 0x9e, 0x82, 0xd5, 0x89, 0x2e, 0x54, 0x18,
	0xf8, 0x0a, 0xd9, 0x43, 0xc8, 0x1b, 0x45, 0xca, 0x4a, 0x55, 0x32, 0xd3, 0x8a, 0xaf, 0xe3, 0x51,
	0xcf, 0x3a, 0xd0, 0xdc, 0x8b, 0xbb, 0xca, 0x50, 0x57, 0x79, 0x42, 0xa8, 0xad, 0x5d, 0x58, 0xf2,
	0xf1, 0x93, 0x6e, 0x8e, 0xe8, 0x4a, 0x53, 0x1b, 0xc5, 0x08, 0x3e, 0x36, 0xda, 0x9c, 0xfb, 0xb0,
	0xfa, 0x1c, 0x3d, 0x9c, 0x9e, 0xfe, 0xe4, 0xf0, 0xe2, 0x11, 0x9f, 0x62, 0x3f, 0xf4, 0xb8, 0xbe,
	0x91, 0x75, 0x00, 0xcb, 0x63, 0xac, 0x44, 0x99, 0x0d, 0x39, 0x9d, 0x60, 0x09, 0x79, 0x78, 0x76,
	0xfe, 0xa6, 0x20, 0x67, 0x8a, 0x4f, 0xe6, 0x63, 0x4f, 0x00, 0xda, 0x64, 0x8e, 0x4e, 0x93, 0x6b,
	0x52, 0x50, 0x38, 0xb4, 0xdd, 0xd8, 0xdc, 0xae, 0x31, 0xb7, 0x7b, 0x6a, 0xdc, 0xdf, 0xc8, 0x27,
	0xec, 0x23, 0xcd, 0x18, 0xcc, 0xf9, 0xbc, 0x8f, 0xc9, 0xad, 0xd0, 0x7f, 0x56, 0x81, 0x42, 0x07,
	0x55, 0x5b, 0x0a, 0xf2, 0x6d, 0x72, 0x31, 0xa3, 0x10, 0x73, 0x23, 0x2b, 0x24, 0x1b, 0xa1, 0xac,
	0x79, 0xba, 0x84, 0x52, 0x7c, 0x09, 0x06, 0x6e, 0x8c, 0x30, 0x98, 0x0d, 0x99, 0xc8, 0xd7, 0x0b,
	0xd4, 0x59, 0x8e, 0x88, 0x67, 0xd2, 0x6b, 0x44, 0x20, 0x5b, 0x81, 0x79, 0x5a, 0x20, 0x2b, 0x4b,
	0x75, 0xe2, 0xc3, 0xe1, 0xaf, 0x39, 0x58, 0x32, 0x7a, 0x4f, 0x50, 0x5e, 0x88, 0x36, 0xb2, 0x2e,
	0x94, 0xc6, 0x77, 0x80, 0xd9, 0x94, 0x6a, 0xe6, 0x62, 0xd8, 0xe3, 0xa6, 0x70, 0xf6, 0xbe, 0xfe,
	0xfe, 0xf3, 0x33, 0x7d, 0xcf, 0x29, 0x47, 0x2b, 0xab, 0x6a, 0x17, 0x07, 0x2d, 0xd4, 0xfc, 0x60,
	0xf8, 0xaa, 0xa8, 0xa7, 0xc3, 0x3d, 0x61, 0x6f, 0xa1, 0x30, 0xb2, 0x27, 0xac, 0x4c, 0x89, 0xa6,
	0x37, 0x67, 0xb2, 0xc2, 0x0e, 0x55, 0xb8, 0xc3, 0x36, 0x6f, 0xa8, 0x50, 0xfb, 0x2c, 0x3a, 0x5f,
	0x58, 0x0f, 0x8a, 0x63, 0xbe, 0x66, 0xeb, 0x94, 0x65, 0xd6, 0xc6, 0xd9, 0xf6, 0xac, 0x50, 0x6c,
	0x16, 0x67, 0x9b, 0xaa, 0xad, 0xb3, 0x9b, 0xf4, 0xb0, 0xf7, 0x50, 0x1a, 0xf7, 0x6c, 0x32, 0xad,
	0x99, 0x46, 0xb6, 0xd7, 0xa6, 0xec, 0xf2, 0x22, 0x7a, 0x0b, 0x8d, 0xa8, 0x07, 0xb7, 0x8b, 0x0a,
	0x69, 0x62, 0xc6, 0xd0, 0xd7, 0x13, 0x9b, 0x58, 0x04, 0xdb, 0x9a, 0x0e, 0x24, 0x72, 0x5c, 0xaa,
	0x53, 0x65, 0xbb, 0xb7, 0xd5, 0xa9, 0x99, 0x75, 0x50, 0xf5, 0xe3, 0x1f, 0x47, 0xaf, 0x1b, 0x9b,
	0xb0, 0xd0, 0xc1, 0x2e, 0x1f, 0x78, 0x9a, 0xfd, 0xcf, 0x96, 0xa0, 0x68, 0x17, 0x28, 0xff, 0x89,
	0xe6, 0x7a, 0xa0, 0xde, 0x6c, 0xc3, 0x16, 0x64, 0xeb, 0xc8, 0x25, 0x4a, 0xb6, 0x9c, 0x4b, 0xdb,
	0x45, 0x3e, 0xd0, 0xef, 0x02, 0x29, 0xae, 0xe8, 0x09, 0xae, 0xa4, 0x5b, 0x8b, 0x00, 0x43, 0xc2,
	0x7f, 0xad, 0x2c, 0x29, 0x7f, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xec, 0xd8, 0xa0, 0x22, 0x75,
	0x06, 0x00, 0x00,
}
