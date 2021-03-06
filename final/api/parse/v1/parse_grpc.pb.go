// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: parse/v1/parse.proto

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

// ParseClient is the client API for Parse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParseClient interface {
	ParseQuestion(ctx context.Context, in *ParseQuestionRequest, opts ...grpc.CallOption) (*ParseQuestionReply, error)
}

type parseClient struct {
	cc grpc.ClientConnInterface
}

func NewParseClient(cc grpc.ClientConnInterface) ParseClient {
	return &parseClient{cc}
}

func (c *parseClient) ParseQuestion(ctx context.Context, in *ParseQuestionRequest, opts ...grpc.CallOption) (*ParseQuestionReply, error) {
	out := new(ParseQuestionReply)
	err := c.cc.Invoke(ctx, "/parse.v1.Parse/ParseQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParseServer is the server API for Parse service.
// All implementations must embed UnimplementedParseServer
// for forward compatibility
type ParseServer interface {
	ParseQuestion(context.Context, *ParseQuestionRequest) (*ParseQuestionReply, error)
	mustEmbedUnimplementedParseServer()
}

// UnimplementedParseServer must be embedded to have forward compatible implementations.
type UnimplementedParseServer struct {
}

func (UnimplementedParseServer) ParseQuestion(context.Context, *ParseQuestionRequest) (*ParseQuestionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseQuestion not implemented")
}
func (UnimplementedParseServer) mustEmbedUnimplementedParseServer() {}

// UnsafeParseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParseServer will
// result in compilation errors.
type UnsafeParseServer interface {
	mustEmbedUnimplementedParseServer()
}

func RegisterParseServer(s grpc.ServiceRegistrar, srv ParseServer) {
	s.RegisterService(&Parse_ServiceDesc, srv)
}

func _Parse_ParseQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParseServer).ParseQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parse.v1.Parse/ParseQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParseServer).ParseQuestion(ctx, req.(*ParseQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Parse_ServiceDesc is the grpc.ServiceDesc for Parse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parse.v1.Parse",
	HandlerType: (*ParseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseQuestion",
			Handler:    _Parse_ParseQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parse/v1/parse.proto",
}
