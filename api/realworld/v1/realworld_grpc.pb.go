// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// RealWorldClient is the client API for RealWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealWorldClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type realWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewRealWorldClient(cc grpc.ClientConnInterface) RealWorldClient {
	return &realWorldClient{cc}
}

func (c *realWorldClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.RealWorld/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealWorldServer is the server API for RealWorld service.
// All implementations must embed UnimplementedRealWorldServer
// for forward compatibility
type RealWorldServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedRealWorldServer()
}

// UnimplementedRealWorldServer must be embedded to have forward compatible implementations.
type UnimplementedRealWorldServer struct {
}

func (UnimplementedRealWorldServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedRealWorldServer) mustEmbedUnimplementedRealWorldServer() {}

// UnsafeRealWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealWorldServer will
// result in compilation errors.
type UnsafeRealWorldServer interface {
	mustEmbedUnimplementedRealWorldServer()
}

func RegisterRealWorldServer(s grpc.ServiceRegistrar, srv RealWorldServer) {
	s.RegisterService(&RealWorld_ServiceDesc, srv)
}

func _RealWorld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealWorldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.RealWorld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealWorldServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealWorld_ServiceDesc is the grpc.ServiceDesc for RealWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realworld.v1.RealWorld",
	HandlerType: (*RealWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _RealWorld_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/realworld/v1/realworld.proto",
}
