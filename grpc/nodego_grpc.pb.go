// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc2
// source: nodego.proto

package grpc

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

// NodegoClient is the client API for Nodego service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodegoClient interface {
	// Execute next flow node.
	ExecuteNext(ctx context.Context, in *ExecuteNextRequest, opts ...grpc.CallOption) (*ExecuteNextResponse, error)
}

type nodegoClient struct {
	cc grpc.ClientConnInterface
}

func NewNodegoClient(cc grpc.ClientConnInterface) NodegoClient {
	return &nodegoClient{cc}
}

func (c *nodegoClient) ExecuteNext(ctx context.Context, in *ExecuteNextRequest, opts ...grpc.CallOption) (*ExecuteNextResponse, error) {
	out := new(ExecuteNextResponse)
	err := c.cc.Invoke(ctx, "/nodego.Nodego/ExecuteNext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodegoServer is the server API for Nodego service.
// All implementations must embed UnimplementedNodegoServer
// for forward compatibility
type NodegoServer interface {
	// Execute next flow node.
	ExecuteNext(context.Context, *ExecuteNextRequest) (*ExecuteNextResponse, error)
	mustEmbedUnimplementedNodegoServer()
}

// UnimplementedNodegoServer must be embedded to have forward compatible implementations.
type UnimplementedNodegoServer struct {
}

func (UnimplementedNodegoServer) ExecuteNext(context.Context, *ExecuteNextRequest) (*ExecuteNextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteNext not implemented")
}
func (UnimplementedNodegoServer) mustEmbedUnimplementedNodegoServer() {}

// UnsafeNodegoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodegoServer will
// result in compilation errors.
type UnsafeNodegoServer interface {
	mustEmbedUnimplementedNodegoServer()
}

func RegisterNodegoServer(s grpc.ServiceRegistrar, srv NodegoServer) {
	s.RegisterService(&Nodego_ServiceDesc, srv)
}

func _Nodego_ExecuteNext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteNextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodegoServer).ExecuteNext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodego.Nodego/ExecuteNext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodegoServer).ExecuteNext(ctx, req.(*ExecuteNextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Nodego_ServiceDesc is the grpc.ServiceDesc for Nodego service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nodego_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodego.Nodego",
	HandlerType: (*NodegoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteNext",
			Handler:    _Nodego_ExecuteNext_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodego.proto",
}
