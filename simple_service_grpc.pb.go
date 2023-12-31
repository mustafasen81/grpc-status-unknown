// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: simple_service.proto

package grpcstatusunknown

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
	SimpleService_Subscribe_FullMethodName = "/grpcstatusunknown.SimpleService/Subscribe"
)

// SimpleServiceClient is the client API for SimpleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleServiceClient interface {
	Subscribe(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (SimpleService_SubscribeClient, error)
}

type simpleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleServiceClient(cc grpc.ClientConnInterface) SimpleServiceClient {
	return &simpleServiceClient{cc}
}

func (c *simpleServiceClient) Subscribe(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (SimpleService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimpleService_ServiceDesc.Streams[0], SimpleService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &simpleServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SimpleService_SubscribeClient interface {
	Recv() (*SimpleResponse, error)
	grpc.ClientStream
}

type simpleServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *simpleServiceSubscribeClient) Recv() (*SimpleResponse, error) {
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimpleServiceServer is the server API for SimpleService service.
// All implementations must embed UnimplementedSimpleServiceServer
// for forward compatibility
type SimpleServiceServer interface {
	Subscribe(*SimpleRequest, SimpleService_SubscribeServer) error
	mustEmbedUnimplementedSimpleServiceServer()
}

// UnimplementedSimpleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleServiceServer struct {
}

func (UnimplementedSimpleServiceServer) Subscribe(*SimpleRequest, SimpleService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSimpleServiceServer) mustEmbedUnimplementedSimpleServiceServer() {}

// UnsafeSimpleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServiceServer will
// result in compilation errors.
type UnsafeSimpleServiceServer interface {
	mustEmbedUnimplementedSimpleServiceServer()
}

func RegisterSimpleServiceServer(s grpc.ServiceRegistrar, srv SimpleServiceServer) {
	s.RegisterService(&SimpleService_ServiceDesc, srv)
}

func _SimpleService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SimpleServiceServer).Subscribe(m, &simpleServiceSubscribeServer{stream})
}

type SimpleService_SubscribeServer interface {
	Send(*SimpleResponse) error
	grpc.ServerStream
}

type simpleServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *simpleServiceSubscribeServer) Send(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SimpleService_ServiceDesc is the grpc.ServiceDesc for SimpleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcstatusunknown.SimpleService",
	HandlerType: (*SimpleServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _SimpleService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simple_service.proto",
}
