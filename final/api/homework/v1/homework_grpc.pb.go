// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: homework/v1/homework.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HomeworkClient is the client API for Homework service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeworkClient interface {
	// 获取答案
	FindDetailSolution(ctx context.Context, in *FindDetailSolutionRequest, opts ...grpc.CallOption) (*FindDetailSolutionReply, error)
	// 搜索记录列表
	GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsReply, error)
}

type homeworkClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeworkClient(cc grpc.ClientConnInterface) HomeworkClient {
	return &homeworkClient{cc}
}

func (c *homeworkClient) FindDetailSolution(ctx context.Context, in *FindDetailSolutionRequest, opts ...grpc.CallOption) (*FindDetailSolutionReply, error) {
	out := new(FindDetailSolutionReply)
	err := c.cc.Invoke(ctx, "/homework.v1.Homework/FindDetailSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkClient) GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsReply, error) {
	out := new(GetRecordsReply)
	err := c.cc.Invoke(ctx, "/homework.v1.Homework/GetRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeworkServer is the server API for Homework service.
// All implementations must embed UnimplementedHomeworkServer
// for forward compatibility
type HomeworkServer interface {
	// 获取答案
	FindDetailSolution(context.Context, *FindDetailSolutionRequest) (*FindDetailSolutionReply, error)
	// 搜索记录列表
	GetRecords(context.Context, *GetRecordsRequest) (*GetRecordsReply, error)
	mustEmbedUnimplementedHomeworkServer()
}

// UnimplementedHomeworkServer must be embedded to have forward compatible implementations.
type UnimplementedHomeworkServer struct {
}

func (UnimplementedHomeworkServer) FindDetailSolution(context.Context, *FindDetailSolutionRequest) (*FindDetailSolutionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDetailSolution not implemented")
}
func (UnimplementedHomeworkServer) GetRecords(context.Context, *GetRecordsRequest) (*GetRecordsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecords not implemented")
}
func (UnimplementedHomeworkServer) mustEmbedUnimplementedHomeworkServer() {}

// UnsafeHomeworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeworkServer will
// result in compilation errors.
type UnsafeHomeworkServer interface {
	mustEmbedUnimplementedHomeworkServer()
}

func RegisterHomeworkServer(s grpc.ServiceRegistrar, srv HomeworkServer) {
	s.RegisterService(&Homework_ServiceDesc, srv)
}

func _Homework_FindDetailSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDetailSolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeworkServer).FindDetailSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homework.v1.Homework/FindDetailSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeworkServer).FindDetailSolution(ctx, req.(*FindDetailSolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Homework_GetRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeworkServer).GetRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/homework.v1.Homework/GetRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeworkServer).GetRecords(ctx, req.(*GetRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Homework_ServiceDesc is the grpc.ServiceDesc for Homework service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Homework_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "homework.v1.Homework",
	HandlerType: (*HomeworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindDetailSolution",
			Handler:    _Homework_FindDetailSolution_Handler,
		},
		{
			MethodName: "GetRecords",
			Handler:    _Homework_GetRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "homework/v1/homework.proto",
}