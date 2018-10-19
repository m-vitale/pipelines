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

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type CreateRunRequest struct {
	Run                  *Run     `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRunRequest) Reset()         { *m = CreateRunRequest{} }
func (m *CreateRunRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRunRequest) ProtoMessage()    {}
func (*CreateRunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3419bc3417bf873, []int{0}
}

func (m *CreateRunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRunRequest.Unmarshal(m, b)
}
func (m *CreateRunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRunRequest.Marshal(b, m, deterministic)
}
func (m *CreateRunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRunRequest.Merge(m, src)
}
func (m *CreateRunRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRunRequest.Size(m)
}
func (m *CreateRunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRunRequest proto.InternalMessageInfo

func (m *CreateRunRequest) GetRun() *Run {
	if m != nil {
		return m.Run
	}
	return nil
}

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
	return fileDescriptor_e3419bc3417bf873, []int{1}
}

func (m *GetRunV2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunV2Request.Unmarshal(m, b)
}
func (m *GetRunV2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunV2Request.Marshal(b, m, deterministic)
}
func (m *GetRunV2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunV2Request.Merge(m, src)
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
	return fileDescriptor_e3419bc3417bf873, []int{2}
}

func (m *GetRunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRunRequest.Unmarshal(m, b)
}
func (m *GetRunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRunRequest.Marshal(b, m, deterministic)
}
func (m *GetRunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRunRequest.Merge(m, src)
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
	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Can be format of "field_name", "field_name asc" or "field_name des"
	// Ascending by default.
	SortBy string `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	// What resource reference to filter on. Expect {type=XX,id=XX}
	// E.g. Listing job for an experiment would be {type=EXPERIMENT,id=123}
	ResourceReference    string   `protobuf:"bytes,4,opt,name=resource_reference,json=resourceReference,proto3" json:"resource_reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRunsRequest) Reset()         { *m = ListRunsRequest{} }
func (m *ListRunsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRunsRequest) ProtoMessage()    {}
func (*ListRunsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3419bc3417bf873, []int{3}
}

func (m *ListRunsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsRequest.Unmarshal(m, b)
}
func (m *ListRunsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsRequest.Marshal(b, m, deterministic)
}
func (m *ListRunsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsRequest.Merge(m, src)
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

func (m *ListRunsRequest) GetResourceReference() string {
	if m != nil {
		return m.ResourceReference
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
	return fileDescriptor_e3419bc3417bf873, []int{4}
}

func (m *ListRunsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRunsResponse.Unmarshal(m, b)
}
func (m *ListRunsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRunsResponse.Marshal(b, m, deterministic)
}
func (m *ListRunsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRunsResponse.Merge(m, src)
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
	// Output. Unique run ID. Generated by API server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required input field. Name provided by user,
	// or auto generated if run is created by scheduled job. Not unique.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional input field. Describing the purpose of the run
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Required input field.
	// Describing what the pipeline manifest and parameters to use for the run.
	PipelineSpec *PipelineSpec `protobuf:"bytes,4,opt,name=pipeline_spec,json=pipelineSpec,proto3" json:"pipeline_spec,omitempty"`
	// Optional input field. Specify which resource this run belongs to.
	ResourceReferences []*ResourceReference `protobuf:"bytes,5,rep,name=resource_references,json=resourceReferences,proto3" json:"resource_references,omitempty"`
	// Output. The time that the run created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Output. When this run is scheduled to run. This could be different from
	// created_at. For example, if a run is from a backfilling job that was
	// supposed to run 2 month ago, the scheduled_at is 2 month ago,
	// v.s. created_at is the current time.
	ScheduledAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
	// Output. The status of the run.
	// One of [Pending, Running, Succeeded, Skipped, Failed, Error]
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	// In case any error happens retrieving a run field, only run ID
	// and the error message is returned. Client has the flexibility of choosing
	// how to handle error. This is especially useful during listing call.
	Error string `protobuf:"bytes,12,opt,name=error,proto3" json:"error,omitempty"`
	// TODO(yangpa): Following will be deprecated in v1beta1
	JobId                string   `protobuf:"bytes,13,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Namespace            string   `protobuf:"bytes,14,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3419bc3417bf873, []int{5}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
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

func (m *Run) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Run) GetPipelineSpec() *PipelineSpec {
	if m != nil {
		return m.PipelineSpec
	}
	return nil
}

func (m *Run) GetResourceReferences() []*ResourceReference {
	if m != nil {
		return m.ResourceReferences
	}
	return nil
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

func (m *Run) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Run) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Run) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type PipelineRuntime struct {
	// Output. The runtime manifest of the pipeline, including the status
	// of pipeline steps and fields need for UI visualization etc.
	PipelineManifest string `protobuf:"bytes,10,opt,name=pipeline_manifest,json=pipelineManifest,proto3" json:"pipeline_manifest,omitempty"`
	// Output. The runtime manifest of the argo workflow.
	// This is deprecated after pipeline_runtime_manifest is in use.
	WorkflowManifest     string   `protobuf:"bytes,11,opt,name=workflow_manifest,json=workflowManifest,proto3" json:"workflow_manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipelineRuntime) Reset()         { *m = PipelineRuntime{} }
func (m *PipelineRuntime) String() string { return proto.CompactTextString(m) }
func (*PipelineRuntime) ProtoMessage()    {}
func (*PipelineRuntime) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3419bc3417bf873, []int{6}
}

func (m *PipelineRuntime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineRuntime.Unmarshal(m, b)
}
func (m *PipelineRuntime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineRuntime.Marshal(b, m, deterministic)
}
func (m *PipelineRuntime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineRuntime.Merge(m, src)
}
func (m *PipelineRuntime) XXX_Size() int {
	return xxx_messageInfo_PipelineRuntime.Size(m)
}
func (m *PipelineRuntime) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineRuntime.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineRuntime proto.InternalMessageInfo

func (m *PipelineRuntime) GetPipelineManifest() string {
	if m != nil {
		return m.PipelineManifest
	}
	return ""
}

func (m *PipelineRuntime) GetWorkflowManifest() string {
	if m != nil {
		return m.WorkflowManifest
	}
	return ""
}

type RunDetail struct {
	Run             *Run             `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
	PipelineRuntime *PipelineRuntime `protobuf:"bytes,2,opt,name=pipeline_runtime,json=pipelineRuntime,proto3" json:"pipeline_runtime,omitempty"`
	// TODO(yangpa): Following will be deprecated in v1beta1
	Workflow             string   `protobuf:"bytes,3,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunDetail) Reset()         { *m = RunDetail{} }
func (m *RunDetail) String() string { return proto.CompactTextString(m) }
func (*RunDetail) ProtoMessage()    {}
func (*RunDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3419bc3417bf873, []int{7}
}

func (m *RunDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunDetail.Unmarshal(m, b)
}
func (m *RunDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunDetail.Marshal(b, m, deterministic)
}
func (m *RunDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunDetail.Merge(m, src)
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

func (m *RunDetail) GetPipelineRuntime() *PipelineRuntime {
	if m != nil {
		return m.PipelineRuntime
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
	proto.RegisterType((*CreateRunRequest)(nil), "api.CreateRunRequest")
	proto.RegisterType((*GetRunV2Request)(nil), "api.GetRunV2Request")
	proto.RegisterType((*GetRunRequest)(nil), "api.GetRunRequest")
	proto.RegisterType((*ListRunsRequest)(nil), "api.ListRunsRequest")
	proto.RegisterType((*ListRunsResponse)(nil), "api.ListRunsResponse")
	proto.RegisterType((*Run)(nil), "api.Run")
	proto.RegisterType((*PipelineRuntime)(nil), "api.PipelineRuntime")
	proto.RegisterType((*RunDetail)(nil), "api.RunDetail")
}

func init() { proto.RegisterFile("run.proto", fileDescriptor_e3419bc3417bf873) }

var fileDescriptor_e3419bc3417bf873 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xc6, 0x76, 0xe2, 0xd8, 0xe3, 0x38, 0x4e, 0x26, 0x4d, 0x39, 0x5d, 0x5d, 0x6a, 0x9d, 0x10,
	0xaa, 0xa0, 0xb1, 0x85, 0x91, 0x90, 0x40, 0xaa, 0x50, 0x02, 0x52, 0x55, 0x89, 0xa2, 0xe8, 0x12,
	0x55, 0x88, 0x2f, 0xc7, 0xfa, 0x6e, 0xec, 0x6c, 0xe2, 0xec, 0x1e, 0xfb, 0xd2, 0xd0, 0x54, 0xfd,
	0x82, 0xf8, 0x05, 0x20, 0x7e, 0x04, 0xbf, 0x87, 0xbf, 0xc0, 0x47, 0x7e, 0x04, 0xba, 0xbd, 0x97,
	0x5e, 0xec, 0x42, 0x3f, 0xd9, 0x3b, 0xf3, 0xcc, 0xb3, 0xcf, 0xce, 0x3d, 0x33, 0xd0, 0x55, 0x56,
	0x8c, 0x53, 0x25, 0x8d, 0xc4, 0x16, 0x4b, 0xb9, 0xdf, 0x23, 0xa5, 0xa4, 0xca, 0x23, 0xfe, 0x70,
	0x21, 0xe5, 0x62, 0x49, 0x13, 0x96, 0xf2, 0x09, 0x13, 0x42, 0x1a, 0x66, 0xb8, 0x14, 0xba, 0xc8,
	0x3e, 0x28, 0xb2, 0xee, 0x34, 0xb3, 0xf3, 0x89, 0xe1, 0x57, 0xa4, 0x0d, 0xbb, 0x4a, 0x0b, 0xc0,
	0x7e, 0xca, 0x53, 0x5a, 0x72, 0x41, 0x91, 0x4e, 0x29, 0x2e, 0x82, 0x9e, 0x22, 0x2d, 0xad, 0x8a,
	0x29, 0x52, 0x34, 0x27, 0x45, 0x22, 0xa6, 0x22, 0xf3, 0xc8, 0xfd, 0xc4, 0x87, 0x0b, 0x12, 0x87,
	0xfa, 0x9a, 0x2d, 0x16, 0xa4, 0x26, 0x32, 0x75, 0x37, 0xae, 0xdf, 0x1e, 0x8c, 0x61, 0xf7, 0x6b,
	0x45, 0xcc, 0x50, 0x68, 0x45, 0x48, 0x3f, 0x59, 0xd2, 0x06, 0x7d, 0x68, 0x29, 0x2b, 0xbc, 0xc6,
	0xa8, 0xf1, 0xb0, 0x37, 0xed, 0x8c, 0x59, 0xca, 0xc7, 0x59, 0x36, 0x0b, 0x06, 0x0f, 0x61, 0xf0,
	0x84, 0x4c, 0x68, 0xc5, 0xf3, 0x69, 0x09, 0x3f, 0x80, 0xb6, 0xb2, 0x22, 0xe2, 0x89, 0xab, 0xe8,
	0x86, 0x9b, 0xca, 0x8a, 0xa7, 0x49, 0xf0, 0x18, 0xfa, 0x39, 0xb2, 0x86, 0xbb, 0x90, 0xb3, 0x1a,
	0xee, 0x42, 0xce, 0x9e, 0x26, 0xb5, 0xf2, 0x66, 0xbd, 0xfc, 0x8f, 0x06, 0x0c, 0xbe, 0xe5, 0x3a,
	0x23, 0xd0, 0x25, 0xc3, 0x7d, 0x80, 0x94, 0x2d, 0x28, 0x32, 0xf2, 0x92, 0x44, 0xc1, 0xd2, 0xcd,
	0x22, 0x67, 0x59, 0x00, 0xef, 0x81, 0x3b, 0x44, 0x9a, 0xdf, 0x90, 0x23, 0xdb, 0x0c, 0x3b, 0x59,
	0xe0, 0x94, 0xdf, 0x10, 0xbe, 0x0f, 0x5b, 0x5a, 0x2a, 0x13, 0xcd, 0x5e, 0x7a, 0x2d, 0x57, 0xd8,
	0xce, 0x8e, 0xc7, 0x2f, 0xf1, 0x10, 0x70, 0xbd, 0x97, 0xde, 0x86, 0xc3, 0xec, 0x95, 0x99, 0xb0,
	0x4c, 0x04, 0xdf, 0xc3, 0xee, 0x1b, 0x59, 0x3a, 0x95, 0x42, 0x13, 0x0e, 0x61, 0x43, 0x59, 0xa1,
	0xbd, 0xc6, 0xa8, 0x75, 0xab, 0x63, 0x2e, 0x8a, 0x1f, 0xc1, 0x40, 0xd0, 0xcf, 0x26, 0xaa, 0x49,
	0xcf, 0x5f, 0xda, 0xcf, 0xc2, 0x27, 0xa5, 0xfc, 0xe0, 0xcf, 0x16, 0xb4, 0x42, 0x2b, 0x70, 0x07,
	0x9a, 0x55, 0x8f, 0x9a, 0x3c, 0x41, 0x84, 0x0d, 0xc1, 0xae, 0xa8, 0x28, 0x72, 0xff, 0x71, 0x04,
	0xbd, 0x84, 0x74, 0xac, 0xb8, 0xfb, 0xb0, 0xc5, 0x8b, 0xea, 0x21, 0xfc, 0x1c, 0xfa, 0xb7, 0x7c,
	0xe3, 0x5e, 0xd4, 0x9b, 0xee, 0x39, 0x71, 0x27, 0x45, 0xe6, 0x34, 0xa5, 0x38, 0xdc, 0x4e, 0x6b,
	0x27, 0x7c, 0x02, 0xfb, 0xeb, 0xed, 0xd0, 0xde, 0xa6, 0x7b, 0xda, 0xdd, 0xfc, 0x69, 0xab, 0x4d,
	0x09, 0x71, 0xad, 0x4f, 0x1a, 0xbf, 0x00, 0x88, 0x9d, 0xb3, 0x92, 0x88, 0x19, 0xaf, 0xed, 0x6e,
	0xf7, 0xc7, 0xb9, 0xd9, 0xc7, 0xa5, 0xd9, 0xc7, 0x67, 0xa5, 0xd9, 0xc3, 0x6e, 0x81, 0x3e, 0x32,
	0xf8, 0x18, 0xb6, 0x75, 0x7c, 0x4e, 0x89, 0x5d, 0xe6, 0xc5, 0x5b, 0xef, 0x2c, 0xee, 0x55, 0xf8,
	0x23, 0x83, 0x77, 0xa1, 0xad, 0x0d, 0x33, 0x56, 0x7b, 0x9d, 0xe2, 0x4b, 0xbb, 0x13, 0xde, 0x81,
	0x4d, 0x37, 0x96, 0xde, 0x76, 0x6e, 0x34, 0x77, 0xa8, 0xd9, 0xb2, 0x5f, 0xb7, 0xe5, 0x10, 0xba,
	0x59, 0xa7, 0x75, 0xca, 0x62, 0xf2, 0x76, 0x72, 0xab, 0x55, 0x81, 0xe0, 0x12, 0x06, 0x65, 0x0f,
	0x43, 0x2b, 0xb2, 0x89, 0xc5, 0x4f, 0x60, 0xaf, 0x6a, 0xf8, 0x15, 0x13, 0x7c, 0x4e, 0xda, 0x78,
	0xe0, 0x0a, 0x77, 0xcb, 0xc4, 0xb3, 0x22, 0x9e, 0x81, 0xaf, 0xa5, 0xba, 0x9c, 0x2f, 0xe5, 0xf5,
	0x1b, 0x70, 0x2f, 0x07, 0x97, 0x89, 0x12, 0x1c, 0xfc, 0xda, 0x80, 0x6e, 0x68, 0xc5, 0x37, 0x64,
	0x18, 0x5f, 0xfe, 0xdf, 0x74, 0xe2, 0x57, 0x50, 0x5d, 0x15, 0xa9, 0x5c, 0x97, 0xb3, 0x4d, 0x6f,
	0x7a, 0xe7, 0xd6, 0x77, 0x2f, 0x34, 0x87, 0x83, 0x74, 0xe5, 0x11, 0x3e, 0x74, 0xca, 0xeb, 0x0b,
	0x53, 0x55, 0xe7, 0xe9, 0x3f, 0x4d, 0x80, 0xd0, 0x8a, 0x53, 0x52, 0x2f, 0x78, 0x4c, 0xf8, 0x1d,
	0x74, 0xab, 0xcd, 0x81, 0x07, 0x8e, 0x7e, 0x75, 0x93, 0xf8, 0x95, 0xbc, 0x60, 0xf4, 0xcb, 0x5f,
	0x7f, 0xff, 0xde, 0xf4, 0x83, 0xfd, 0x6c, 0x0b, 0xea, 0xc9, 0x8b, 0x4f, 0xd9, 0x32, 0x3d, 0x67,
	0xd3, 0x49, 0x36, 0x21, 0x5f, 0x3a, 0xed, 0xcf, 0xa1, 0x53, 0x6e, 0x16, 0xcc, 0xd5, 0xae, 0x2c,
	0x1a, 0x7f, 0xa7, 0x64, 0xcb, 0x3b, 0x11, 0x7c, 0xe8, 0x38, 0x3f, 0xc0, 0xe1, 0x5b, 0x38, 0x27,
	0xaf, 0xf2, 0xa5, 0xf2, 0x1a, 0xcf, 0xa0, 0x53, 0x0e, 0x6c, 0xc1, 0xbb, 0xb2, 0x56, 0xfc, 0x83,
	0x95, 0x68, 0x3e, 0xd5, 0xc1, 0x3d, 0x47, 0x7f, 0x80, 0x6f, 0x93, 0x8c, 0x3f, 0x42, 0x3b, 0x97,
	0x87, 0x58, 0xd3, 0xfa, 0x5f, 0x4a, 0xa7, 0x8e, 0xea, 0x11, 0x7e, 0xbc, 0x42, 0x75, 0x21, 0x67,
	0x7a, 0xf2, 0x2a, 0xb7, 0xdf, 0xeb, 0xdb, 0xba, 0x8f, 0x4f, 0x7e, 0x3b, 0x7a, 0x16, 0x0e, 0x61,
	0x2b, 0xa1, 0x39, 0xb3, 0x4b, 0x83, 0x7b, 0x38, 0x80, 0xbe, 0xdf, 0x73, 0xcc, 0xa7, 0xce, 0xd0,
	0x3f, 0x3c, 0x80, 0xfb, 0xd0, 0x3e, 0x26, 0xa6, 0x48, 0xe1, 0x7e, 0xa7, 0xe9, 0xf7, 0x99, 0x35,
	0xe7, 0x52, 0xf1, 0x1b, 0xb7, 0xe2, 0x47, 0xcd, 0xd9, 0x36, 0x40, 0x05, 0x78, 0x6f, 0xd6, 0x76,
	0x83, 0xf3, 0xd9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75, 0xf1, 0xb7, 0x27, 0xad, 0x06, 0x00,
	0x00,
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
	CreateRun(ctx context.Context, in *CreateRunRequest, opts ...grpc.CallOption) (*Run, error)
	GetRunV2(ctx context.Context, in *GetRunV2Request, opts ...grpc.CallOption) (*RunDetail, error)
	ListRuns(ctx context.Context, in *ListRunsRequest, opts ...grpc.CallOption) (*ListRunsResponse, error)
	// TODO(yangpa): This will be deprecated in v1beta1
	GetRun(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*RunDetail, error)
}

type runServiceClient struct {
	cc *grpc.ClientConn
}

func NewRunServiceClient(cc *grpc.ClientConn) RunServiceClient {
	return &runServiceClient{cc}
}

func (c *runServiceClient) CreateRun(ctx context.Context, in *CreateRunRequest, opts ...grpc.CallOption) (*Run, error) {
	out := new(Run)
	err := c.cc.Invoke(ctx, "/api.RunService/CreateRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	CreateRun(context.Context, *CreateRunRequest) (*Run, error)
	GetRunV2(context.Context, *GetRunV2Request) (*RunDetail, error)
	ListRuns(context.Context, *ListRunsRequest) (*ListRunsResponse, error)
	// TODO(yangpa): This will be deprecated in v1beta1
	GetRun(context.Context, *GetRunRequest) (*RunDetail, error)
}

func RegisterRunServiceServer(s *grpc.Server, srv RunServiceServer) {
	s.RegisterService(&_RunService_serviceDesc, srv)
}

func _RunService_CreateRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).CreateRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RunService/CreateRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).CreateRun(ctx, req.(*CreateRunRequest))
	}
	return interceptor(ctx, in, info, handler)
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
			MethodName: "CreateRun",
			Handler:    _RunService_CreateRun_Handler,
		},
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
