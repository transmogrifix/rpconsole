// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: redpanda/api/console/v1alpha1/debug_bundle.proto

package consolev1alpha1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetDebugBundleStatusResponse_Status int32

const (
	GetDebugBundleStatusResponse_STATUS_UNSPECIFIED GetDebugBundleStatusResponse_Status = 0
	GetDebugBundleStatusResponse_STATUS_RUNNING     GetDebugBundleStatusResponse_Status = 1
	GetDebugBundleStatusResponse_STATUS_SUCCESS     GetDebugBundleStatusResponse_Status = 2
	GetDebugBundleStatusResponse_STATUS_ERROR       GetDebugBundleStatusResponse_Status = 3
)

// Enum value maps for GetDebugBundleStatusResponse_Status.
var (
	GetDebugBundleStatusResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_RUNNING",
		2: "STATUS_SUCCESS",
		3: "STATUS_ERROR",
	}
	GetDebugBundleStatusResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_RUNNING":     1,
		"STATUS_SUCCESS":     2,
		"STATUS_ERROR":       3,
	}
)

func (x GetDebugBundleStatusResponse_Status) Enum() *GetDebugBundleStatusResponse_Status {
	p := new(GetDebugBundleStatusResponse_Status)
	*p = x
	return p
}

func (x GetDebugBundleStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetDebugBundleStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_enumTypes[0].Descriptor()
}

func (GetDebugBundleStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_redpanda_api_console_v1alpha1_debug_bundle_proto_enumTypes[0]
}

func (x GetDebugBundleStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetDebugBundleStatusResponse_Status.Descriptor instead.
func (GetDebugBundleStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{3, 0}
}

type CreateDebugBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerLogsSizeLimitBytes int32    `protobuf:"varint,1,opt,name=controller_logs_size_limit_bytes,json=controllerLogsSizeLimitBytes,proto3" json:"controller_logs_size_limit_bytes,omitempty"`
	CpuProfilerWaitSeconds       int32    `protobuf:"varint,2,opt,name=cpu_profiler_wait_seconds,json=cpuProfilerWaitSeconds,proto3" json:"cpu_profiler_wait_seconds,omitempty"`
	LogsSince                    string   `protobuf:"bytes,3,opt,name=logs_since,json=logsSince,proto3" json:"logs_since,omitempty"` // TODO add validation 'YYYY-MM-DD', 'yesterday', or 'today'
	LogsSizeLimitBytes           int32    `protobuf:"varint,4,opt,name=logs_size_limit_bytes,json=logsSizeLimitBytes,proto3" json:"logs_size_limit_bytes,omitempty"`
	LogsUntil                    string   `protobuf:"bytes,5,opt,name=logs_until,json=logsUntil,proto3" json:"logs_until,omitempty"` // TODO add validation 'YYYY-MM-DD', 'yesterday', or 'today'
	MetricsIntervalSeconds       int32    `protobuf:"varint,6,opt,name=metrics_interval_seconds,json=metricsIntervalSeconds,proto3" json:"metrics_interval_seconds,omitempty"`
	Partition                    []string `protobuf:"bytes,7,rep,name=partition,proto3" json:"partition,omitempty"`
}

func (x *CreateDebugBundleRequest) Reset() {
	*x = CreateDebugBundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDebugBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDebugBundleRequest) ProtoMessage() {}

func (x *CreateDebugBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDebugBundleRequest.ProtoReflect.Descriptor instead.
func (*CreateDebugBundleRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDebugBundleRequest) GetControllerLogsSizeLimitBytes() int32 {
	if x != nil {
		return x.ControllerLogsSizeLimitBytes
	}
	return 0
}

func (x *CreateDebugBundleRequest) GetCpuProfilerWaitSeconds() int32 {
	if x != nil {
		return x.CpuProfilerWaitSeconds
	}
	return 0
}

func (x *CreateDebugBundleRequest) GetLogsSince() string {
	if x != nil {
		return x.LogsSince
	}
	return ""
}

func (x *CreateDebugBundleRequest) GetLogsSizeLimitBytes() int32 {
	if x != nil {
		return x.LogsSizeLimitBytes
	}
	return 0
}

func (x *CreateDebugBundleRequest) GetLogsUntil() string {
	if x != nil {
		return x.LogsUntil
	}
	return ""
}

func (x *CreateDebugBundleRequest) GetMetricsIntervalSeconds() int32 {
	if x != nil {
		return x.MetricsIntervalSeconds
	}
	return 0
}

func (x *CreateDebugBundleRequest) GetPartition() []string {
	if x != nil {
		return x.Partition
	}
	return nil
}

type CreateDebugBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"` // UUID
}

func (x *CreateDebugBundleResponse) Reset() {
	*x = CreateDebugBundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDebugBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDebugBundleResponse) ProtoMessage() {}

func (x *CreateDebugBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDebugBundleResponse.ProtoReflect.Descriptor instead.
func (*CreateDebugBundleResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDebugBundleResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type GetDebugBundleStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDebugBundleStatusRequest) Reset() {
	*x = GetDebugBundleStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDebugBundleStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDebugBundleStatusRequest) ProtoMessage() {}

func (x *GetDebugBundleStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDebugBundleStatusRequest.ProtoReflect.Descriptor instead.
func (*GetDebugBundleStatusRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{2}
}

type GetDebugBundleStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId  string                              `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Status GetDebugBundleStatusResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=redpanda.api.console.v1alpha1.GetDebugBundleStatusResponse_Status" json:"status,omitempty"`
}

func (x *GetDebugBundleStatusResponse) Reset() {
	*x = GetDebugBundleStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDebugBundleStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDebugBundleStatusResponse) ProtoMessage() {}

func (x *GetDebugBundleStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDebugBundleStatusResponse.ProtoReflect.Descriptor instead.
func (*GetDebugBundleStatusResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{3}
}

func (x *GetDebugBundleStatusResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *GetDebugBundleStatusResponse) GetStatus() GetDebugBundleStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetDebugBundleStatusResponse_STATUS_UNSPECIFIED
}

type DeleteDebugBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *DeleteDebugBundleRequest) Reset() {
	*x = DeleteDebugBundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDebugBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDebugBundleRequest) ProtoMessage() {}

func (x *DeleteDebugBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDebugBundleRequest.ProtoReflect.Descriptor instead.
func (*DeleteDebugBundleRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDebugBundleRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type DeleteDebugBundleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDebugBundleResponse) Reset() {
	*x = DeleteDebugBundleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDebugBundleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDebugBundleResponse) ProtoMessage() {}

func (x *DeleteDebugBundleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDebugBundleResponse.ProtoReflect.Descriptor instead.
func (*DeleteDebugBundleResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{5}
}

type DeleteDebugBundleFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *DeleteDebugBundleFileRequest) Reset() {
	*x = DeleteDebugBundleFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDebugBundleFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDebugBundleFileRequest) ProtoMessage() {}

func (x *DeleteDebugBundleFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDebugBundleFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteDebugBundleFileRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDebugBundleFileRequest) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type DeleteDebugBundleFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDebugBundleFileResponse) Reset() {
	*x = DeleteDebugBundleFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDebugBundleFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDebugBundleFileResponse) ProtoMessage() {}

func (x *DeleteDebugBundleFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDebugBundleFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteDebugBundleFileResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP(), []int{7}
}

var File_redpanda_api_console_v1alpha1_debug_bundle_proto protoreflect.FileDescriptor

var file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDesc = []byte{
	0x0a, 0x30, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x22, 0xe6, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46,
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67,
	0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x70, 0x75, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x63, 0x70, 0x75, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x57, 0x61, 0x69, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65,
	0x12, 0x31, 0x0a, 0x15, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x6c, 0x6f, 0x67, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x75, 0x6e, 0x74, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x1d,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xed, 0x01,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x5a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x31, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a,
	0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xd5, 0x04, 0x0a, 0x12, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12,
	0x37, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x2e,
	0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x72, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x37,
	0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e,
	0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x94, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x3b, 0x2e,
	0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb1, 0x02, 0x0a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x10, 0x44, 0x65, 0x62, 0x75, 0x67, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x43, 0xaa,
	0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x29, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x52, 0x65,
	0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescOnce sync.Once
	file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescData = file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDesc
)

func file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescGZIP() []byte {
	file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescOnce.Do(func() {
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescData)
	})
	return file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDescData
}

var file_redpanda_api_console_v1alpha1_debug_bundle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_redpanda_api_console_v1alpha1_debug_bundle_proto_goTypes = []interface{}{
	(GetDebugBundleStatusResponse_Status)(0), // 0: redpanda.api.console.v1alpha1.GetDebugBundleStatusResponse.Status
	(*CreateDebugBundleRequest)(nil),         // 1: redpanda.api.console.v1alpha1.CreateDebugBundleRequest
	(*CreateDebugBundleResponse)(nil),        // 2: redpanda.api.console.v1alpha1.CreateDebugBundleResponse
	(*GetDebugBundleStatusRequest)(nil),      // 3: redpanda.api.console.v1alpha1.GetDebugBundleStatusRequest
	(*GetDebugBundleStatusResponse)(nil),     // 4: redpanda.api.console.v1alpha1.GetDebugBundleStatusResponse
	(*DeleteDebugBundleRequest)(nil),         // 5: redpanda.api.console.v1alpha1.DeleteDebugBundleRequest
	(*DeleteDebugBundleResponse)(nil),        // 6: redpanda.api.console.v1alpha1.DeleteDebugBundleResponse
	(*DeleteDebugBundleFileRequest)(nil),     // 7: redpanda.api.console.v1alpha1.DeleteDebugBundleFileRequest
	(*DeleteDebugBundleFileResponse)(nil),    // 8: redpanda.api.console.v1alpha1.DeleteDebugBundleFileResponse
}
var file_redpanda_api_console_v1alpha1_debug_bundle_proto_depIdxs = []int32{
	0, // 0: redpanda.api.console.v1alpha1.GetDebugBundleStatusResponse.status:type_name -> redpanda.api.console.v1alpha1.GetDebugBundleStatusResponse.Status
	1, // 1: redpanda.api.console.v1alpha1.DebugBundleService.CreateDebugBundle:input_type -> redpanda.api.console.v1alpha1.CreateDebugBundleRequest
	3, // 2: redpanda.api.console.v1alpha1.DebugBundleService.GetDebugBundleStatus:input_type -> redpanda.api.console.v1alpha1.GetDebugBundleStatusRequest
	5, // 3: redpanda.api.console.v1alpha1.DebugBundleService.DeleteDebugBundle:input_type -> redpanda.api.console.v1alpha1.DeleteDebugBundleRequest
	7, // 4: redpanda.api.console.v1alpha1.DebugBundleService.DeleteDebugBundleFile:input_type -> redpanda.api.console.v1alpha1.DeleteDebugBundleFileRequest
	2, // 5: redpanda.api.console.v1alpha1.DebugBundleService.CreateDebugBundle:output_type -> redpanda.api.console.v1alpha1.CreateDebugBundleResponse
	4, // 6: redpanda.api.console.v1alpha1.DebugBundleService.GetDebugBundleStatus:output_type -> redpanda.api.console.v1alpha1.GetDebugBundleStatusResponse
	6, // 7: redpanda.api.console.v1alpha1.DebugBundleService.DeleteDebugBundle:output_type -> redpanda.api.console.v1alpha1.DeleteDebugBundleResponse
	8, // 8: redpanda.api.console.v1alpha1.DebugBundleService.DeleteDebugBundleFile:output_type -> redpanda.api.console.v1alpha1.DeleteDebugBundleFileResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_redpanda_api_console_v1alpha1_debug_bundle_proto_init() }
func file_redpanda_api_console_v1alpha1_debug_bundle_proto_init() {
	if File_redpanda_api_console_v1alpha1_debug_bundle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDebugBundleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDebugBundleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDebugBundleStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDebugBundleStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDebugBundleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDebugBundleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDebugBundleFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDebugBundleFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redpanda_api_console_v1alpha1_debug_bundle_proto_goTypes,
		DependencyIndexes: file_redpanda_api_console_v1alpha1_debug_bundle_proto_depIdxs,
		EnumInfos:         file_redpanda_api_console_v1alpha1_debug_bundle_proto_enumTypes,
		MessageInfos:      file_redpanda_api_console_v1alpha1_debug_bundle_proto_msgTypes,
	}.Build()
	File_redpanda_api_console_v1alpha1_debug_bundle_proto = out.File
	file_redpanda_api_console_v1alpha1_debug_bundle_proto_rawDesc = nil
	file_redpanda_api_console_v1alpha1_debug_bundle_proto_goTypes = nil
	file_redpanda_api_console_v1alpha1_debug_bundle_proto_depIdxs = nil
}
