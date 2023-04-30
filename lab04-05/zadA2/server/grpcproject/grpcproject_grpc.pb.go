// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: grpcproject.proto

package grpcproject

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

const (
	GrpcProject_SayHello_FullMethodName = "/GrpcProject/SayHello"
)

// GrpcProjectClient is the client API for GrpcProject service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcProjectClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type grpcProjectClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcProjectClient(cc grpc.ClientConnInterface) GrpcProjectClient {
	return &grpcProjectClient{cc}
}

func (c *grpcProjectClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, GrpcProject_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcProjectServer is the server API for GrpcProject service.
// All implementations must embed UnimplementedGrpcProjectServer
// for forward compatibility
type GrpcProjectServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGrpcProjectServer()
}

// UnimplementedGrpcProjectServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcProjectServer struct {
}

func (UnimplementedGrpcProjectServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGrpcProjectServer) mustEmbedUnimplementedGrpcProjectServer() {}

// UnsafeGrpcProjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcProjectServer will
// result in compilation errors.
type UnsafeGrpcProjectServer interface {
	mustEmbedUnimplementedGrpcProjectServer()
}

func RegisterGrpcProjectServer(s grpc.ServiceRegistrar, srv GrpcProjectServer) {
	s.RegisterService(&GrpcProject_ServiceDesc, srv)
}

func _GrpcProject_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcProjectServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcProject_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcProjectServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GrpcProject_ServiceDesc is the grpc.ServiceDesc for GrpcProject service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcProject_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GrpcProject",
	HandlerType: (*GrpcProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GrpcProject_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcproject.proto",
}
