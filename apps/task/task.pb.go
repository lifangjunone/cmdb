// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/task/pb/task.proto

package task

import (
	resource "github.com/lifangjunone/cmdb/apps/resource"
	page "github.com/lifangjunone/cmdb/common/pb/page"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 任务类型
type Type int32

const (
	// 资源同步任务
	Type_RESOURCE_SYNC Type = 0
	// 资源释放任务
	Type_RESOURCE_RELEASE Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "RESOURCE_SYNC",
		1: "RESOURCE_RELEASE",
	}
	Type_value = map[string]int32{
		"RESOURCE_SYNC":    0,
		"RESOURCE_RELEASE": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_task_pb_task_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_apps_task_pb_task_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{0}
}

// 任务运行阶段
type Stage int32

const (
	Stage_PENDING Stage = 0
	Stage_RUNNING Stage = 1
	Stage_SUCCESS Stage = 2
	Stage_FAILED  Stage = 3
	Stage_WARNING Stage = 4
)

// Enum value maps for Stage.
var (
	Stage_name = map[int32]string{
		0: "PENDING",
		1: "RUNNING",
		2: "SUCCESS",
		3: "FAILED",
		4: "WARNING",
	}
	Stage_value = map[string]int32{
		"PENDING": 0,
		"RUNNING": 1,
		"SUCCESS": 2,
		"FAILED":  3,
		"WARNING": 4,
	}
)

func (x Stage) Enum() *Stage {
	p := new(Stage)
	*p = x
	return p
}

func (x Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_task_pb_task_proto_enumTypes[1].Descriptor()
}

func (Stage) Type() protoreflect.EnumType {
	return &file_apps_task_pb_task_proto_enumTypes[1]
}

func (x Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Stage.Descriptor instead.
func (Stage) EnumDescriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{1}
}

type TaskSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// @gotags: json:"items"
	Items []*Task `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TaskSet) Reset() {
	*x = TaskSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSet) ProtoMessage() {}

func (x *TaskSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSet.ProtoReflect.Descriptor instead.
func (*TaskSet) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskSet) GetItems() []*Task {
	if x != nil {
		return x.Items
	}
	return nil
}

// Task 同个区域的同一种资源一次只能有1个task run
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// secret
	// @gotags: json:"secret_description"
	SecretDescription string `protobuf:"bytes,2,opt,name=secret_description,json=secretDescription,proto3" json:"secret_description,omitempty"`
	// 任务描述
	// @gotags: json:"data"
	Data *CreateTaskRequst `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// 任务状态
	// @gotags: json:"status"
	Status *Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetSecretDescription() string {
	if x != nil {
		return x.SecretDescription
	}
	return ""
}

func (x *Task) GetData() *CreateTaskRequst {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Task) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type CreateTaskRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务类型
	// @gotags: json:"type"
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=cmdb.task.Type" json:"type,omitempty"`
	// 测试运行
	// @gotags: json:"dry_run"
	DryRun bool `protobuf:"varint,2,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	// 任务使用的云商凭证Id
	// @gotags: json:"secret_id" validate:"required,lte=100"
	SecretId string `protobuf:"bytes,3,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	// 任务操作的资源类型
	// @gotags: json:"resource_type"
	ResourceType resource.Type `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=cmdb.resource.Type" json:"resource_type,omitempty"`
	// 通知资源的Region
	// @gotags: json:"region"
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	// 额外的一些参数
	// @gotags: json:"params"
	Params map[string]string `protobuf:"bytes,6,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 任务执行的超时时间, 单位时秒
	// @gotags: json:"timeout"
	Timeout int64 `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *CreateTaskRequst) Reset() {
	*x = CreateTaskRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequst) ProtoMessage() {}

func (x *CreateTaskRequst) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequst.ProtoReflect.Descriptor instead.
func (*CreateTaskRequst) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskRequst) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_RESOURCE_SYNC
}

func (x *CreateTaskRequst) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

func (x *CreateTaskRequst) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *CreateTaskRequst) GetResourceType() resource.Type {
	if x != nil {
		return x.ResourceType
	}
	return resource.Type(0)
}

func (x *CreateTaskRequst) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateTaskRequst) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *CreateTaskRequst) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 任务状态
	// @gotags: json:"stage"
	Stage Stage `protobuf:"varint,1,opt,name=stage,proto3,enum=cmdb.task.Stage" json:"stage,omitempty"`
	// 失败时的异常信息
	// @gotags: json:"message"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// 开始同步的时间
	// @gotags: json:"start_at"
	StartAt int64 `protobuf:"varint,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// 同步结束时间
	// @gotags: json:"end_at"
	EndAt int64 `protobuf:"varint,4,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	// 成功的条数
	// @gotags: json:"total_succeed"
	TotalSucceed int64 `protobuf:"varint,5,opt,name=total_succeed,json=totalSucceed,proto3" json:"total_succeed,omitempty"`
	// 失败的条数
	// @gotags: json:"total_failed"
	TotalFailed int64 `protobuf:"varint,6,opt,name=total_failed,json=totalFailed,proto3" json:"total_failed,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetStage() Stage {
	if x != nil {
		return x.Stage
	}
	return Stage_PENDING
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *Status) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *Status) GetTotalSucceed() int64 {
	if x != nil {
		return x.TotalSucceed
	}
	return 0
}

func (x *Status) GetTotalFailed() int64 {
	if x != nil {
		return x.TotalFailed
	}
	return 0
}

type QueryTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// 更加资源类型来过滤任务
	// @gotags: json:"resource_type"
	ResourceType resource.Type `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=cmdb.resource.Type" json:"resource_type,omitempty"`
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *QueryTaskRequest) Reset() {
	*x = QueryTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTaskRequest) ProtoMessage() {}

func (x *QueryTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTaskRequest.ProtoReflect.Descriptor instead.
func (*QueryTaskRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{4}
}

func (x *QueryTaskRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryTaskRequest) GetResourceType() resource.Type {
	if x != nil {
		return x.ResourceType
	}
	return resource.Type(0)
}

func (x *QueryTaskRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type DescribeTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Task id
	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeTaskRequest) Reset() {
	*x = DescribeTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_task_pb_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTaskRequest) ProtoMessage() {}

func (x *DescribeTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_task_pb_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTaskRequest.ProtoReflect.Descriptor instead.
func (*DescribeTaskRequest) Descriptor() ([]byte, []int) {
	return file_apps_task_pb_task_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeTaskRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_task_pb_task_proto protoreflect.FileDescriptor

var file_apps_task_pb_task_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd5, 0x02, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x3f,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xc4, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x26, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x10,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x59,
	0x4e, 0x43, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x01, 0x2a, 0x47, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x32, 0xc4, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e,
	0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6d, 0x64,
	0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3c, 0x0a, 0x09, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6d, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6d, 0x64, 0x62,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x66, 0x61, 0x6e, 0x67, 0x6a,
	0x75, 0x6e, 0x6f, 0x6e, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_task_pb_task_proto_rawDescOnce sync.Once
	file_apps_task_pb_task_proto_rawDescData = file_apps_task_pb_task_proto_rawDesc
)

func file_apps_task_pb_task_proto_rawDescGZIP() []byte {
	file_apps_task_pb_task_proto_rawDescOnce.Do(func() {
		file_apps_task_pb_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_task_pb_task_proto_rawDescData)
	})
	return file_apps_task_pb_task_proto_rawDescData
}

var file_apps_task_pb_task_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_task_pb_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_task_pb_task_proto_goTypes = []interface{}{
	(Type)(0),                   // 0: cmdb.task.Type
	(Stage)(0),                  // 1: cmdb.task.Stage
	(*TaskSet)(nil),             // 2: cmdb.task.TaskSet
	(*Task)(nil),                // 3: cmdb.task.Task
	(*CreateTaskRequst)(nil),    // 4: cmdb.task.CreateTaskRequst
	(*Status)(nil),              // 5: cmdb.task.Status
	(*QueryTaskRequest)(nil),    // 6: cmdb.task.QueryTaskRequest
	(*DescribeTaskRequest)(nil), // 7: cmdb.task.DescribeTaskRequest
	nil,                         // 8: cmdb.task.CreateTaskRequst.ParamsEntry
	(resource.Type)(0),          // 9: cmdb.resource.Type
	(*page.PageRequest)(nil),    // 10: cmdb.page.PageRequest
}
var file_apps_task_pb_task_proto_depIdxs = []int32{
	3,  // 0: cmdb.task.TaskSet.items:type_name -> cmdb.task.Task
	4,  // 1: cmdb.task.Task.data:type_name -> cmdb.task.CreateTaskRequst
	5,  // 2: cmdb.task.Task.status:type_name -> cmdb.task.Status
	0,  // 3: cmdb.task.CreateTaskRequst.type:type_name -> cmdb.task.Type
	9,  // 4: cmdb.task.CreateTaskRequst.resource_type:type_name -> cmdb.resource.Type
	8,  // 5: cmdb.task.CreateTaskRequst.params:type_name -> cmdb.task.CreateTaskRequst.ParamsEntry
	1,  // 6: cmdb.task.Status.stage:type_name -> cmdb.task.Stage
	10, // 7: cmdb.task.QueryTaskRequest.page:type_name -> cmdb.page.PageRequest
	9,  // 8: cmdb.task.QueryTaskRequest.resource_type:type_name -> cmdb.resource.Type
	4,  // 9: cmdb.task.Service.CreateTask:input_type -> cmdb.task.CreateTaskRequst
	6,  // 10: cmdb.task.Service.QueryTask:input_type -> cmdb.task.QueryTaskRequest
	7,  // 11: cmdb.task.Service.DescribeTask:input_type -> cmdb.task.DescribeTaskRequest
	3,  // 12: cmdb.task.Service.CreateTask:output_type -> cmdb.task.Task
	2,  // 13: cmdb.task.Service.QueryTask:output_type -> cmdb.task.TaskSet
	3,  // 14: cmdb.task.Service.DescribeTask:output_type -> cmdb.task.Task
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_apps_task_pb_task_proto_init() }
func file_apps_task_pb_task_proto_init() {
	if File_apps_task_pb_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_task_pb_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSet); i {
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
		file_apps_task_pb_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_apps_task_pb_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequst); i {
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
		file_apps_task_pb_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_apps_task_pb_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTaskRequest); i {
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
		file_apps_task_pb_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTaskRequest); i {
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
			RawDescriptor: file_apps_task_pb_task_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_task_pb_task_proto_goTypes,
		DependencyIndexes: file_apps_task_pb_task_proto_depIdxs,
		EnumInfos:         file_apps_task_pb_task_proto_enumTypes,
		MessageInfos:      file_apps_task_pb_task_proto_msgTypes,
	}.Build()
	File_apps_task_pb_task_proto = out.File
	file_apps_task_pb_task_proto_rawDesc = nil
	file_apps_task_pb_task_proto_goTypes = nil
	file_apps_task_pb_task_proto_depIdxs = nil
}
