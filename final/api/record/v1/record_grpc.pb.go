// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: record/v1/record.proto

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

// RecordClient is the client API for Record service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordClient interface {
	GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsReply, error)
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordReply, error)
	UpdateRecordSolution(ctx context.Context, in *UpdateRecordSolutionRequest, opts ...grpc.CallOption) (*UpdateRecordSolutionReply, error)
}

type recordClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordClient(cc grpc.ClientConnInterface) RecordClient {
	return &recordClient{cc}
}

func (c *recordClient) GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsReply, error) {
	out := new(GetRecordsReply)
	err := c.cc.Invoke(ctx, "/parse.v1.Record/GetRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordReply, error) {
	out := new(CreateRecordReply)
	err := c.cc.Invoke(ctx, "/parse.v1.Record/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordClient) UpdateRecordSolution(ctx context.Context, in *UpdateRecordSolutionRequest, opts ...grpc.CallOption) (*UpdateRecordSolutionReply, error) {
	out := new(UpdateRecordSolutionReply)
	err := c.cc.Invoke(ctx, "/parse.v1.Record/UpdateRecordSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServer is the server API for Record service.
// All implementations must embed UnimplementedRecordServer
// for forward compatibility
type RecordServer interface {
	GetRecords(context.Context, *GetRecordsRequest) (*GetRecordsReply, error)
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordReply, error)
	UpdateRecordSolution(context.Context, *UpdateRecordSolutionRequest) (*UpdateRecordSolutionReply, error)
	mustEmbedUnimplementedRecordServer()
}

// UnimplementedRecordServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServer struct {
}

func (UnimplementedRecordServer) GetRecords(context.Context, *GetRecordsRequest) (*GetRecordsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecords not implemented")
}
func (UnimplementedRecordServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedRecordServer) UpdateRecordSolution(context.Context, *UpdateRecordSolutionRequest) (*UpdateRecordSolutionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecordSolution not implemented")
}
func (UnimplementedRecordServer) mustEmbedUnimplementedRecordServer() {}

// UnsafeRecordServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServer will
// result in compilation errors.
type UnsafeRecordServer interface {
	mustEmbedUnimplementedRecordServer()
}

func RegisterRecordServer(s grpc.ServiceRegistrar, srv RecordServer) {
	s.RegisterService(&Record_ServiceDesc, srv)
}

func _Record_GetRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).GetRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parse.v1.Record/GetRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).GetRecords(ctx, req.(*GetRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parse.v1.Record/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Record_UpdateRecordSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordSolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServer).UpdateRecordSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parse.v1.Record/UpdateRecordSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServer).UpdateRecordSolution(ctx, req.(*UpdateRecordSolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Record_ServiceDesc is the grpc.ServiceDesc for Record service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Record_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parse.v1.Record",
	HandlerType: (*RecordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecords",
			Handler:    _Record_GetRecords_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _Record_CreateRecord_Handler,
		},
		{
			MethodName: "UpdateRecordSolution",
			Handler:    _Record_UpdateRecordSolution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record/v1/record.proto",
}
