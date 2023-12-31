// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: transistor/v1/transistor.proto

package pb

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
	TransistorService_Command_FullMethodName   = "/transistor.v1.TransistorService/Command"
	TransistorService_Publish_FullMethodName   = "/transistor.v1.TransistorService/Publish"
	TransistorService_Subscribe_FullMethodName = "/transistor.v1.TransistorService/Subscribe"
)

// TransistorServiceClient is the client API for TransistorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransistorServiceClient interface {
	// Open a new command line interface
	Command(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (TransistorService_CommandClient, error)
	// Receive a stream from a non-cluster node
	Publish(ctx context.Context, opts ...grpc.CallOption) (TransistorService_PublishClient, error)
	// Receive a stream from both cluster/non-cluster nodes
	// Always the subscriber should approach to this server
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (TransistorService_SubscribeClient, error)
}

type transistorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransistorServiceClient(cc grpc.ClientConnInterface) TransistorServiceClient {
	return &transistorServiceClient{cc}
}

func (c *transistorServiceClient) Command(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (TransistorService_CommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransistorService_ServiceDesc.Streams[0], TransistorService_Command_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transistorServiceCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransistorService_CommandClient interface {
	Recv() (*CommandResponse, error)
	grpc.ClientStream
}

type transistorServiceCommandClient struct {
	grpc.ClientStream
}

func (x *transistorServiceCommandClient) Recv() (*CommandResponse, error) {
	m := new(CommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transistorServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (TransistorService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransistorService_ServiceDesc.Streams[1], TransistorService_Publish_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transistorServicePublishClient{stream}
	return x, nil
}

type TransistorService_PublishClient interface {
	Send(*PublishRequest) error
	CloseAndRecv() (*PublishResponse, error)
	grpc.ClientStream
}

type transistorServicePublishClient struct {
	grpc.ClientStream
}

func (x *transistorServicePublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transistorServicePublishClient) CloseAndRecv() (*PublishResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transistorServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (TransistorService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransistorService_ServiceDesc.Streams[2], TransistorService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transistorServiceSubscribeClient{stream}
	return x, nil
}

type TransistorService_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type transistorServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *transistorServiceSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transistorServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransistorServiceServer is the server API for TransistorService service.
// All implementations must embed UnimplementedTransistorServiceServer
// for forward compatibility
type TransistorServiceServer interface {
	// Open a new command line interface
	Command(*CommandRequest, TransistorService_CommandServer) error
	// Receive a stream from a non-cluster node
	Publish(TransistorService_PublishServer) error
	// Receive a stream from both cluster/non-cluster nodes
	// Always the subscriber should approach to this server
	Subscribe(TransistorService_SubscribeServer) error
	mustEmbedUnimplementedTransistorServiceServer()
}

// UnimplementedTransistorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransistorServiceServer struct {
}

func (UnimplementedTransistorServiceServer) Command(*CommandRequest, TransistorService_CommandServer) error {
	return status.Errorf(codes.Unimplemented, "method Command not implemented")
}
func (UnimplementedTransistorServiceServer) Publish(TransistorService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedTransistorServiceServer) Subscribe(TransistorService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedTransistorServiceServer) mustEmbedUnimplementedTransistorServiceServer() {}

// UnsafeTransistorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransistorServiceServer will
// result in compilation errors.
type UnsafeTransistorServiceServer interface {
	mustEmbedUnimplementedTransistorServiceServer()
}

func RegisterTransistorServiceServer(s grpc.ServiceRegistrar, srv TransistorServiceServer) {
	s.RegisterService(&TransistorService_ServiceDesc, srv)
}

func _TransistorService_Command_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransistorServiceServer).Command(m, &transistorServiceCommandServer{stream})
}

type TransistorService_CommandServer interface {
	Send(*CommandResponse) error
	grpc.ServerStream
}

type transistorServiceCommandServer struct {
	grpc.ServerStream
}

func (x *transistorServiceCommandServer) Send(m *CommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TransistorService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransistorServiceServer).Publish(&transistorServicePublishServer{stream})
}

type TransistorService_PublishServer interface {
	SendAndClose(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type transistorServicePublishServer struct {
	grpc.ServerStream
}

func (x *transistorServicePublishServer) SendAndClose(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transistorServicePublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TransistorService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransistorServiceServer).Subscribe(&transistorServiceSubscribeServer{stream})
}

type TransistorService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type transistorServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *transistorServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transistorServiceSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransistorService_ServiceDesc is the grpc.ServiceDesc for TransistorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransistorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transistor.v1.TransistorService",
	HandlerType: (*TransistorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Command",
			Handler:       _TransistorService_Command_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _TransistorService_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _TransistorService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "transistor/v1/transistor.proto",
}
