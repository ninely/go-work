// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: solve/v1/solve.proto

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

type SolveQuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 题目内容
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// 题目图片链接
	FileUrl string `protobuf:"bytes,2,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
}

func (x *SolveQuestionRequest) Reset() {
	*x = SolveQuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solve_v1_solve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveQuestionRequest) ProtoMessage() {}

func (x *SolveQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveQuestionRequest.ProtoReflect.Descriptor instead.
func (*SolveQuestionRequest) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{0}
}

func (x *SolveQuestionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SolveQuestionRequest) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type SolveQuestionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 输入题目
	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// 答案
	Solution string `protobuf:"bytes,2,opt,name=solution,proto3" json:"solution,omitempty"`
	// 步骤列表
	Steps []*SolveQuestionReply_Step `protobuf:"bytes,3,rep,name=steps,proto3" json:"steps,omitempty"`
	// 知识点
	Knowledge int32 `protobuf:"varint,4,opt,name=knowledge,proto3" json:"knowledge,omitempty"`
}

func (x *SolveQuestionReply) Reset() {
	*x = SolveQuestionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solve_v1_solve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveQuestionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveQuestionReply) ProtoMessage() {}

func (x *SolveQuestionReply) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveQuestionReply.ProtoReflect.Descriptor instead.
func (*SolveQuestionReply) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{1}
}

func (x *SolveQuestionReply) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *SolveQuestionReply) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

func (x *SolveQuestionReply) GetSteps() []*SolveQuestionReply_Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *SolveQuestionReply) GetKnowledge() int32 {
	if x != nil {
		return x.Knowledge
	}
	return 0
}

type UpdateSolveSolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 题目ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 题目结果
	DetailSolution string `protobuf:"bytes,2,opt,name=detail_solution,json=detailSolution,proto3" json:"detail_solution,omitempty"`
}

func (x *UpdateSolveSolutionRequest) Reset() {
	*x = UpdateSolveSolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solve_v1_solve_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSolveSolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSolveSolutionRequest) ProtoMessage() {}

func (x *UpdateSolveSolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSolveSolutionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSolveSolutionRequest) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSolveSolutionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSolveSolutionRequest) GetDetailSolution() string {
	if x != nil {
		return x.DetailSolution
	}
	return ""
}

type UpdateSolveSolutionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSolveSolutionReply) Reset() {
	*x = UpdateSolveSolutionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solve_v1_solve_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSolveSolutionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSolveSolutionReply) ProtoMessage() {}

func (x *UpdateSolveSolutionReply) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSolveSolutionReply.ProtoReflect.Descriptor instead.
func (*UpdateSolveSolutionReply) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{3}
}

// 步骤
type SolveQuestionReply_Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 步骤名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 步骤结果
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *SolveQuestionReply_Step) Reset() {
	*x = SolveQuestionReply_Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solve_v1_solve_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveQuestionReply_Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveQuestionReply_Step) ProtoMessage() {}

func (x *SolveQuestionReply_Step) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveQuestionReply_Step.ProtoReflect.Descriptor instead.
func (*SolveQuestionReply_Step) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SolveQuestionReply_Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SolveQuestionReply_Step) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_solve_v1_solve_proto protoreflect.FileDescriptor

var file_solve_v1_solve_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x4b, 0x0a, 0x14, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xd1, 0x01,
	0x0a, 0x12, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x1a, 0x32, 0x0a,
	0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x55, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0xbb, 0x01, 0x0a, 0x05, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x4f,
	0x0a, 0x0d, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x6c, 0x76, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_solve_v1_solve_proto_rawDescOnce sync.Once
	file_solve_v1_solve_proto_rawDescData = file_solve_v1_solve_proto_rawDesc
)

func file_solve_v1_solve_proto_rawDescGZIP() []byte {
	file_solve_v1_solve_proto_rawDescOnce.Do(func() {
		file_solve_v1_solve_proto_rawDescData = protoimpl.X.CompressGZIP(file_solve_v1_solve_proto_rawDescData)
	})
	return file_solve_v1_solve_proto_rawDescData
}

var file_solve_v1_solve_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_solve_v1_solve_proto_goTypes = []interface{}{
	(*SolveQuestionRequest)(nil),       // 0: parse.v1.SolveQuestionRequest
	(*SolveQuestionReply)(nil),         // 1: parse.v1.SolveQuestionReply
	(*UpdateSolveSolutionRequest)(nil), // 2: parse.v1.UpdateSolveSolutionRequest
	(*UpdateSolveSolutionReply)(nil),   // 3: parse.v1.UpdateSolveSolutionReply
	(*SolveQuestionReply_Step)(nil),    // 4: parse.v1.SolveQuestionReply.Step
}
var file_solve_v1_solve_proto_depIdxs = []int32{
	4, // 0: parse.v1.SolveQuestionReply.steps:type_name -> parse.v1.SolveQuestionReply.Step
	0, // 1: parse.v1.Solve.SolveQuestion:input_type -> parse.v1.SolveQuestionRequest
	2, // 2: parse.v1.Solve.UpdateSolveSolution:input_type -> parse.v1.UpdateSolveSolutionRequest
	1, // 3: parse.v1.Solve.SolveQuestion:output_type -> parse.v1.SolveQuestionReply
	3, // 4: parse.v1.Solve.UpdateSolveSolution:output_type -> parse.v1.UpdateSolveSolutionReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_solve_v1_solve_proto_init() }
func file_solve_v1_solve_proto_init() {
	if File_solve_v1_solve_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_solve_v1_solve_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveQuestionRequest); i {
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
		file_solve_v1_solve_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveQuestionReply); i {
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
		file_solve_v1_solve_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSolveSolutionRequest); i {
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
		file_solve_v1_solve_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSolveSolutionReply); i {
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
		file_solve_v1_solve_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveQuestionReply_Step); i {
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
			RawDescriptor: file_solve_v1_solve_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_solve_v1_solve_proto_goTypes,
		DependencyIndexes: file_solve_v1_solve_proto_depIdxs,
		MessageInfos:      file_solve_v1_solve_proto_msgTypes,
	}.Build()
	File_solve_v1_solve_proto = out.File
	file_solve_v1_solve_proto_rawDesc = nil
	file_solve_v1_solve_proto_goTypes = nil
	file_solve_v1_solve_proto_depIdxs = nil
}
