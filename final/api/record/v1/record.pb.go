// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: record/v1/record.proto

package v1

import (
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
		mi := &file_record_v1_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsRequest) ProtoMessage() {}

func (x *GetRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[0]
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
	return file_record_v1_record_proto_rawDescGZIP(), []int{0}
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
		mi := &file_record_v1_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsReply) ProtoMessage() {}

func (x *GetRecordsReply) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[1]
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
	return file_record_v1_record_proto_rawDescGZIP(), []int{1}
}

func (x *GetRecordsReply) GetRecords() []*GetRecordsReply_Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type CreateRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 题目内容
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateRecordRequest) Reset() {
	*x = CreateRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_v1_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecordRequest) ProtoMessage() {}

func (x *CreateRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateRecordRequest) Descriptor() ([]byte, []int) {
	return file_record_v1_record_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRecordRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRecordRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateRecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId int64 `protobuf:"varint,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
}

func (x *CreateRecordReply) Reset() {
	*x = CreateRecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_v1_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRecordReply) ProtoMessage() {}

func (x *CreateRecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRecordReply.ProtoReflect.Descriptor instead.
func (*CreateRecordReply) Descriptor() ([]byte, []int) {
	return file_record_v1_record_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRecordReply) GetRecordId() int64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

type UpdateRecordDetailSolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 记录ID
	RecordId int64 `protobuf:"varint,1,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	// 答案来源
	SolutionSource int32 `protobuf:"varint,2,opt,name=solution_source,json=solutionSource,proto3" json:"solution_source,omitempty"`
	// 题目结果
	DetailSolution string `protobuf:"bytes,3,opt,name=detail_solution,json=detailSolution,proto3" json:"detail_solution,omitempty"`
}

func (x *UpdateRecordDetailSolutionRequest) Reset() {
	*x = UpdateRecordDetailSolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_v1_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRecordDetailSolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordDetailSolutionRequest) ProtoMessage() {}

func (x *UpdateRecordDetailSolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordDetailSolutionRequest.ProtoReflect.Descriptor instead.
func (*UpdateRecordDetailSolutionRequest) Descriptor() ([]byte, []int) {
	return file_record_v1_record_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRecordDetailSolutionRequest) GetRecordId() int64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *UpdateRecordDetailSolutionRequest) GetSolutionSource() int32 {
	if x != nil {
		return x.SolutionSource
	}
	return 0
}

func (x *UpdateRecordDetailSolutionRequest) GetDetailSolution() string {
	if x != nil {
		return x.DetailSolution
	}
	return ""
}

type UpdateRecordDetailSolutionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRecordDetailSolutionReply) Reset() {
	*x = UpdateRecordDetailSolutionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_record_v1_record_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRecordDetailSolutionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordDetailSolutionReply) ProtoMessage() {}

func (x *UpdateRecordDetailSolutionReply) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordDetailSolutionReply.ProtoReflect.Descriptor instead.
func (*UpdateRecordDetailSolutionReply) Descriptor() ([]byte, []int) {
	return file_record_v1_record_proto_rawDescGZIP(), []int{5}
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
		mi := &file_record_v1_record_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordsReply_Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordsReply_Record) ProtoMessage() {}

func (x *GetRecordsReply_Record) ProtoReflect() protoreflect.Message {
	mi := &file_record_v1_record_proto_msgTypes[6]
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
	return file_record_v1_record_proto_rawDescGZIP(), []int{1, 0}
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

var File_record_v1_record_proto protoreflect.FileDescriptor

var file_record_v1_record_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x9a, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x1a, 0x4b, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x21, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21,
	0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x32, 0x96, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x46, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x76, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_record_v1_record_proto_rawDescOnce sync.Once
	file_record_v1_record_proto_rawDescData = file_record_v1_record_proto_rawDesc
)

func file_record_v1_record_proto_rawDescGZIP() []byte {
	file_record_v1_record_proto_rawDescOnce.Do(func() {
		file_record_v1_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_record_v1_record_proto_rawDescData)
	})
	return file_record_v1_record_proto_rawDescData
}

var file_record_v1_record_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_record_v1_record_proto_goTypes = []interface{}{
	(*GetRecordsRequest)(nil),                 // 0: parse.v1.GetRecordsRequest
	(*GetRecordsReply)(nil),                   // 1: parse.v1.GetRecordsReply
	(*CreateRecordRequest)(nil),               // 2: parse.v1.CreateRecordRequest
	(*CreateRecordReply)(nil),                 // 3: parse.v1.CreateRecordReply
	(*UpdateRecordDetailSolutionRequest)(nil), // 4: parse.v1.UpdateRecordDetailSolutionRequest
	(*UpdateRecordDetailSolutionReply)(nil),   // 5: parse.v1.UpdateRecordDetailSolutionReply
	(*GetRecordsReply_Record)(nil),            // 6: parse.v1.GetRecordsReply.Record
}
var file_record_v1_record_proto_depIdxs = []int32{
	6, // 0: parse.v1.GetRecordsReply.records:type_name -> parse.v1.GetRecordsReply.Record
	0, // 1: parse.v1.Record.GetRecords:input_type -> parse.v1.GetRecordsRequest
	2, // 2: parse.v1.Record.CreateRecord:input_type -> parse.v1.CreateRecordRequest
	4, // 3: parse.v1.Record.UpdateRecordDetailSolution:input_type -> parse.v1.UpdateRecordDetailSolutionRequest
	1, // 4: parse.v1.Record.GetRecords:output_type -> parse.v1.GetRecordsReply
	3, // 5: parse.v1.Record.CreateRecord:output_type -> parse.v1.CreateRecordReply
	5, // 6: parse.v1.Record.UpdateRecordDetailSolution:output_type -> parse.v1.UpdateRecordDetailSolutionReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_record_v1_record_proto_init() }
func file_record_v1_record_proto_init() {
	if File_record_v1_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_record_v1_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_record_v1_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_record_v1_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRecordRequest); i {
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
		file_record_v1_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRecordReply); i {
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
		file_record_v1_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRecordDetailSolutionRequest); i {
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
		file_record_v1_record_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRecordDetailSolutionReply); i {
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
		file_record_v1_record_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_record_v1_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_record_v1_record_proto_goTypes,
		DependencyIndexes: file_record_v1_record_proto_depIdxs,
		MessageInfos:      file_record_v1_record_proto_msgTypes,
	}.Build()
	File_record_v1_record_proto = out.File
	file_record_v1_record_proto_rawDesc = nil
	file_record_v1_record_proto_goTypes = nil
	file_record_v1_record_proto_depIdxs = nil
}
