// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: homework/v1/homework.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FindDetailSolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 题目内容
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// 题目图片链接
	FileUrl string `protobuf:"bytes,2,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	// 用户ID
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FindDetailSolutionRequest) Reset() {
	*x = FindDetailSolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homework_v1_homework_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDetailSolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDetailSolutionRequest) ProtoMessage() {}

func (x *FindDetailSolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_homework_v1_homework_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDetailSolutionRequest.ProtoReflect.Descriptor instead.
func (*FindDetailSolutionRequest) Descriptor() ([]byte, []int) {
	return file_homework_v1_homework_proto_rawDescGZIP(), []int{0}
}

func (x *FindDetailSolutionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FindDetailSolutionRequest) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *FindDetailSolutionRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindDetailSolutionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 输入题目
	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// 答案
	Solution string `protobuf:"bytes,2,opt,name=solution,proto3" json:"solution,omitempty"`
	// 步骤列表
	Steps []*FindDetailSolutionReply_Step `protobuf:"bytes,3,rep,name=steps,proto3" json:"steps,omitempty"`
	// 知识点
	Knowledge int32 `protobuf:"varint,4,opt,name=knowledge,proto3" json:"knowledge,omitempty"`
}

func (x *FindDetailSolutionReply) Reset() {
	*x = FindDetailSolutionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homework_v1_homework_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDetailSolutionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDetailSolutionReply) ProtoMessage() {}

func (x *FindDetailSolutionReply) ProtoReflect() protoreflect.Message {
	mi := &file_homework_v1_homework_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDetailSolutionReply.ProtoReflect.Descriptor instead.
func (*FindDetailSolutionReply) Descriptor() ([]byte, []int) {
	return file_homework_v1_homework_proto_rawDescGZIP(), []int{1}
}

func (x *FindDetailSolutionReply) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *FindDetailSolutionReply) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

func (x *FindDetailSolutionReply) GetSteps() []*FindDetailSolutionReply_Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *FindDetailSolutionReply) GetKnowledge() int32 {
	if x != nil {
		return x.Knowledge
	}
	return 0
}

type GetRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetRecordsRequest) Reset() {
	*x = GetRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homework_v1_homework_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsRequest) ProtoMessage() {}

func (x *GetRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_homework_v1_homework_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsRequest.ProtoReflect.Descriptor instead.
func (*GetRecordsRequest) Descriptor() ([]byte, []int) {
	return file_homework_v1_homework_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecordsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetRecordsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*GetRecordsReply_Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GetRecordsReply) Reset() {
	*x = GetRecordsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homework_v1_homework_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsReply) ProtoMessage() {}

func (x *GetRecordsReply) ProtoReflect() protoreflect.Message {
	mi := &file_homework_v1_homework_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsReply.ProtoReflect.Descriptor instead.
func (*GetRecordsReply) Descriptor() ([]byte, []int) {
	return file_homework_v1_homework_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecordsReply) GetRecords() []*GetRecordsReply_Record {
	if x != nil {
		return x.Records
	}
	return nil
}

// 步骤
type FindDetailSolutionReply_Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 步骤名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 步骤结果
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *FindDetailSolutionReply_Step) Reset() {
	*x = FindDetailSolutionReply_Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homework_v1_homework_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDetailSolutionReply_Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDetailSolutionReply_Step) ProtoMessage() {}

func (x *FindDetailSolutionReply_Step) ProtoReflect() protoreflect.Message {
	mi := &file_homework_v1_homework_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDetailSolutionReply_Step.ProtoReflect.Descriptor instead.
func (*FindDetailSolutionReply_Step) Descriptor() ([]byte, []int) {
	return file_homework_v1_homework_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindDetailSolutionReply_Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindDetailSolutionReply_Step) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type GetRecordsReply_Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 题目内容
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// 答案
	DetailSolution string `protobuf:"bytes,2,opt,name=detail_solution,json=detailSolution,proto3" json:"detail_solution,omitempty"`
}

func (x *GetRecordsReply_Record) Reset() {
	*x = GetRecordsReply_Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homework_v1_homework_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsReply_Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsReply_Record) ProtoMessage() {}

func (x *GetRecordsReply_Record) ProtoReflect() protoreflect.Message {
	mi := &file_homework_v1_homework_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordsReply_Record.ProtoReflect.Descriptor instead.
func (*GetRecordsReply_Record) Descriptor() ([]byte, []int) {
	return file_homework_v1_homework_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetRecordsReply_Record) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetRecordsReply_Record) GetDetailSolution() string {
	if x != nil {
		return x.DetailSolution
	}
	return ""
}

var File_homework_v1_homework_proto protoreflect.FileDescriptor

var file_homework_v1_homework_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xde, 0x01, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x1a,
	0x32, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x1a, 0x4b, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0xef, 0x01, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x7b,
	0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x66, 0x69, 0x6e,
	0x64, 0x2d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x42, 0x1a, 0x5a, 0x18, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_homework_v1_homework_proto_rawDescOnce sync.Once
	file_homework_v1_homework_proto_rawDescData = file_homework_v1_homework_proto_rawDesc
)

func file_homework_v1_homework_proto_rawDescGZIP() []byte {
	file_homework_v1_homework_proto_rawDescOnce.Do(func() {
		file_homework_v1_homework_proto_rawDescData = protoimpl.X.CompressGZIP(file_homework_v1_homework_proto_rawDescData)
	})
	return file_homework_v1_homework_proto_rawDescData
}

var file_homework_v1_homework_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_homework_v1_homework_proto_goTypes = []interface{}{
	(*FindDetailSolutionRequest)(nil),    // 0: homework.v1.FindDetailSolutionRequest
	(*FindDetailSolutionReply)(nil),      // 1: homework.v1.FindDetailSolutionReply
	(*GetRecordsRequest)(nil),            // 2: homework.v1.GetRecordsRequest
	(*GetRecordsReply)(nil),              // 3: homework.v1.GetRecordsReply
	(*FindDetailSolutionReply_Step)(nil), // 4: homework.v1.FindDetailSolutionReply.Step
	(*GetRecordsReply_Record)(nil),       // 5: homework.v1.GetRecordsReply.Record
}
var file_homework_v1_homework_proto_depIdxs = []int32{
	4, // 0: homework.v1.FindDetailSolutionReply.steps:type_name -> homework.v1.FindDetailSolutionReply.Step
	5, // 1: homework.v1.GetRecordsReply.records:type_name -> homework.v1.GetRecordsReply.Record
	0, // 2: homework.v1.Homework.FindDetailSolution:input_type -> homework.v1.FindDetailSolutionRequest
	2, // 3: homework.v1.Homework.GetRecords:input_type -> homework.v1.GetRecordsRequest
	1, // 4: homework.v1.Homework.FindDetailSolution:output_type -> homework.v1.FindDetailSolutionReply
	3, // 5: homework.v1.Homework.GetRecords:output_type -> homework.v1.GetRecordsReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_homework_v1_homework_proto_init() }
func file_homework_v1_homework_proto_init() {
	if File_homework_v1_homework_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_homework_v1_homework_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDetailSolutionRequest); i {
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
		file_homework_v1_homework_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDetailSolutionReply); i {
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
		file_homework_v1_homework_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordsRequest); i {
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
		file_homework_v1_homework_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordsReply); i {
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
		file_homework_v1_homework_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDetailSolutionReply_Step); i {
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
		file_homework_v1_homework_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordsReply_Record); i {
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
			RawDescriptor: file_homework_v1_homework_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_homework_v1_homework_proto_goTypes,
		DependencyIndexes: file_homework_v1_homework_proto_depIdxs,
		MessageInfos:      file_homework_v1_homework_proto_msgTypes,
	}.Build()
	File_homework_v1_homework_proto = out.File
	file_homework_v1_homework_proto_rawDesc = nil
	file_homework_v1_homework_proto_goTypes = nil
	file_homework_v1_homework_proto_depIdxs = nil
}