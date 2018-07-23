// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

package api

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

type GetJobRequest struct {
	PipelineId           string   `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	JobId                string   `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobRequest) Reset()         { *m = GetJobRequest{} }
func (m *GetJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobRequest) ProtoMessage()    {}
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_30af660c247e8b64, []int{0}
}
func (m *GetJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobRequest.Unmarshal(m, b)
}
func (m *GetJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobRequest.Marshal(b, m, deterministic)
}
func (dst *GetJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobRequest.Merge(dst, src)
}
func (m *GetJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobRequest.Size(m)
}
func (m *GetJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobRequest proto.InternalMessageInfo

func (m *GetJobRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

func (m *GetJobRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type ListJobsRequest struct {
	PipelineId           string   `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortBy               string   `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsRequest) Reset()         { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()    {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_30af660c247e8b64, []int{1}
}
func (m *ListJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsRequest.Unmarshal(m, b)
}
func (m *ListJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsRequest.Marshal(b, m, deterministic)
}
func (dst *ListJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsRequest.Merge(dst, src)
}
func (m *ListJobsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobsRequest.Size(m)
}
func (m *ListJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsRequest proto.InternalMessageInfo

func (m *ListJobsRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

func (m *ListJobsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListJobsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListJobsRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}

type ListJobsResponse struct {
	Jobs                 []*Job   `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsResponse) Reset()         { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()    {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_30af660c247e8b64, []int{2}
}
func (m *ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsResponse.Unmarshal(m, b)
}
func (m *ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsResponse.Marshal(b, m, deterministic)
}
func (dst *ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsResponse.Merge(dst, src)
}
func (m *ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobsResponse.Size(m)
}
func (m *ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsResponse proto.InternalMessageInfo

func (m *ListJobsResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *ListJobsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type Job struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string               `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ScheduledAt          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
	Status               string               `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_30af660c247e8b64, []int{3}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (dst *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(dst, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Job) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Job) GetScheduledAt() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduledAt
	}
	return nil
}

func (m *Job) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type JobDetail struct {
	Job                  *Job     `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Workflow             string   `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobDetail) Reset()         { *m = JobDetail{} }
func (m *JobDetail) String() string { return proto.CompactTextString(m) }
func (*JobDetail) ProtoMessage()    {}
func (*JobDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_30af660c247e8b64, []int{4}
}
func (m *JobDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobDetail.Unmarshal(m, b)
}
func (m *JobDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobDetail.Marshal(b, m, deterministic)
}
func (dst *JobDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobDetail.Merge(dst, src)
}
func (m *JobDetail) XXX_Size() int {
	return xxx_messageInfo_JobDetail.Size(m)
}
func (m *JobDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_JobDetail.DiscardUnknown(m)
}

var xxx_messageInfo_JobDetail proto.InternalMessageInfo

func (m *JobDetail) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobDetail) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobRequest)(nil), "api.GetJobRequest")
	proto.RegisterType((*ListJobsRequest)(nil), "api.ListJobsRequest")
	proto.RegisterType((*ListJobsResponse)(nil), "api.ListJobsResponse")
	proto.RegisterType((*Job)(nil), "api.Job")
	proto.RegisterType((*JobDetail)(nil), "api.JobDetail")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*JobDetail, error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*JobDetail, error) {
	out := new(JobDetail)
	err := c.cc.Invoke(ctx, "/api.JobService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, "/api.JobService/ListJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	GetJob(context.Context, *GetJobRequest) (*JobDetail, error)
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.JobService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.JobService/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _JobService_GetJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _JobService_ListJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_job_30af660c247e8b64) }

var fileDescriptor_job_30af660c247e8b64 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xf3, 0x61, 0xe2, 0x09, 0x6d, 0xd1, 0x88, 0x42, 0x64, 0x82, 0x1a, 0xf9, 0x80, 0x22,
	0x21, 0x6c, 0x91, 0x22, 0x24, 0x24, 0x38, 0x14, 0x90, 0xaa, 0x46, 0x1c, 0x90, 0xdb, 0x03, 0x37,
	0x6b, 0x37, 0xde, 0xa6, 0xeb, 0x3a, 0xde, 0xc5, 0xbb, 0x69, 0x69, 0xab, 0x5e, 0xb8, 0xf0, 0x03,
	0xf8, 0x69, 0x9c, 0xb8, 0x23, 0x7e, 0x07, 0xda, 0xb5, 0x9d, 0x96, 0x72, 0x00, 0x4e, 0xf6, 0xbc,
	0x9d, 0x79, 0xf3, 0xf1, 0x1e, 0x78, 0x99, 0xa0, 0xa1, 0x2c, 0x85, 0x16, 0xd8, 0x26, 0x92, 0xfb,
	0xc3, 0xb9, 0x10, 0xf3, 0x9c, 0x45, 0x44, 0xf2, 0x88, 0x14, 0x85, 0xd0, 0x44, 0x73, 0x51, 0xa8,
	0x2a, 0xc5, 0xdf, 0xaa, 0x5f, 0x6d, 0x44, 0x97, 0x87, 0x91, 0xe6, 0x0b, 0xa6, 0x34, 0x59, 0xc8,
	0x2a, 0x21, 0xd8, 0x85, 0xb5, 0x5d, 0xa6, 0xa7, 0x82, 0xc6, 0xec, 0xe3, 0x92, 0x29, 0x8d, 0x5b,
	0xd0, 0x97, 0x5c, 0xb2, 0x9c, 0x17, 0x2c, 0xe1, 0xe9, 0xc0, 0x19, 0x39, 0x63, 0x2f, 0x86, 0x06,
	0xda, 0x4b, 0x71, 0x13, 0xdc, 0x4c, 0x50, 0xf3, 0xd6, 0xb2, 0x6f, 0xdd, 0x4c, 0xd0, 0xbd, 0x34,
	0xf8, 0xe2, 0xc0, 0xc6, 0x3b, 0xae, 0x0c, 0x95, 0xfa, 0x67, 0xae, 0x87, 0x00, 0x92, 0xcc, 0x59,
	0xa2, 0xc5, 0x31, 0x2b, 0x6a, 0x3e, 0xcf, 0x20, 0x07, 0x06, 0xc0, 0x07, 0x60, 0x83, 0x44, 0xf1,
	0x73, 0x36, 0x68, 0x8f, 0x9c, 0x71, 0x37, 0xee, 0x19, 0x60, 0x9f, 0x9f, 0x33, 0xbc, 0x0f, 0xb7,
	0x94, 0x28, 0x75, 0x42, 0xcf, 0x06, 0x1d, 0x5b, 0xe8, 0x9a, 0xf0, 0xf5, 0x59, 0xf0, 0x01, 0xee,
	0x5c, 0x0d, 0xa2, 0xa4, 0x28, 0x14, 0xc3, 0x21, 0x74, 0x32, 0x41, 0xd5, 0xc0, 0x19, 0xb5, 0xc7,
	0xfd, 0x49, 0x2f, 0x24, 0x92, 0x87, 0x66, 0x69, 0x8b, 0xe2, 0x23, 0xd8, 0x28, 0xd8, 0x27, 0x9d,
	0xfc, 0x31, 0xcb, 0x9a, 0x81, 0xdf, 0x37, 0xf3, 0x04, 0xdf, 0x1d, 0x68, 0x4f, 0x05, 0xc5, 0x75,
	0x68, 0xad, 0xd6, 0x69, 0xf1, 0x14, 0x11, 0x3a, 0x05, 0x59, 0xb0, 0xba, 0xc8, 0xfe, 0xe3, 0x10,
	0x3c, 0xf3, 0x55, 0x92, 0xcc, 0xaa, 0xd9, 0xbd, 0xf8, 0x0a, 0xc0, 0x17, 0x00, 0xb3, 0x92, 0x11,
	0xcd, 0xd2, 0x84, 0x68, 0x3b, 0x7f, 0x7f, 0xe2, 0x87, 0x95, 0x58, 0x61, 0x23, 0x56, 0x78, 0xd0,
	0x88, 0x15, 0x7b, 0x75, 0xf6, 0x8e, 0xc6, 0x57, 0x70, 0x5b, 0xcd, 0x8e, 0x58, 0xba, 0xcc, 0xab,
	0xe2, 0xee, 0x5f, 0x8b, 0xfb, 0xab, 0xfc, 0x1d, 0x8d, 0xf7, 0xc0, 0x55, 0x9a, 0xe8, 0xa5, 0x1a,
	0xb8, 0xf5, 0xd5, 0x6c, 0x14, 0xbc, 0x01, 0x6f, 0x2a, 0xe8, 0x5b, 0xa6, 0x09, 0xcf, 0xd1, 0x87,
	0x76, 0x26, 0xa8, 0xdd, 0xf0, 0xfa, 0xb5, 0x0c, 0x88, 0x3e, 0xf4, 0x4e, 0x45, 0x79, 0x7c, 0x98,
	0x8b, 0xd3, 0x7a, 0xe1, 0x55, 0x3c, 0xf9, 0xe9, 0x00, 0x4c, 0x05, 0xdd, 0x67, 0xe5, 0x09, 0x9f,
	0x31, 0xcc, 0xc0, 0xad, 0xcc, 0x85, 0x68, 0x39, 0x7e, 0x73, 0x9a, 0xbf, 0xde, 0xf0, 0x56, 0x4d,
	0x83, 0x97, 0x9f, 0xbf, 0xfd, 0xf8, 0xda, 0x7a, 0x8e, 0xcf, 0x8c, 0x97, 0x55, 0x74, 0xf2, 0x94,
	0xe4, 0xf2, 0x88, 0x4c, 0xa2, 0xc6, 0x2f, 0x2a, 0xba, 0xb8, 0xe6, 0xa6, 0xcb, 0xc8, 0x48, 0x17,
	0x5d, 0x54, 0x5e, 0xbc, 0xc4, 0x1c, 0x7a, 0x8d, 0xea, 0x78, 0xd7, 0x32, 0xdf, 0x70, 0xa3, 0xbf,
	0x79, 0x03, 0xad, 0xac, 0x11, 0x6c, 0xdb, 0xb6, 0x4f, 0xf0, 0xf1, 0x7f, 0xb4, 0xa5, 0xae, 0x3d,
	0xf3, 0xf6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x78, 0xb1, 0xed, 0x8e, 0x03, 0x00, 0x00,
}
