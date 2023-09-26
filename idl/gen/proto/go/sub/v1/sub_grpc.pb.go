// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sub/v1/sub.proto

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
	SubService_SendMessage_FullMethodName    = "/sub.v1.SubService/SendMessage"
	SubService_SubscribeTopic_FullMethodName = "/sub.v1.SubService/SubscribeTopic"
)

// SubServiceClient is the client API for SubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubServiceClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (SubService_SubscribeTopicClient, error)
}

type subServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubServiceClient(cc grpc.ClientConnInterface) SubServiceClient {
	return &subServiceClient{cc}
}

func (c *subServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, SubService_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subServiceClient) SubscribeTopic(ctx context.Context, in *SubscribeTopicRequest, opts ...grpc.CallOption) (SubService_SubscribeTopicClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubService_ServiceDesc.Streams[0], SubService_SubscribeTopic_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &subServiceSubscribeTopicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubService_SubscribeTopicClient interface {
	Recv() (*SubscribeTopicResponse, error)
	grpc.ClientStream
}

type subServiceSubscribeTopicClient struct {
	grpc.ClientStream
}

func (x *subServiceSubscribeTopicClient) Recv() (*SubscribeTopicResponse, error) {
	m := new(SubscribeTopicResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubServiceServer is the server API for SubService service.
// All implementations must embed UnimplementedSubServiceServer
// for forward compatibility
type SubServiceServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	SubscribeTopic(*SubscribeTopicRequest, SubService_SubscribeTopicServer) error
	mustEmbedUnimplementedSubServiceServer()
}

// UnimplementedSubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubServiceServer struct {
}

func (UnimplementedSubServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedSubServiceServer) SubscribeTopic(*SubscribeTopicRequest, SubService_SubscribeTopicServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTopic not implemented")
}
func (UnimplementedSubServiceServer) mustEmbedUnimplementedSubServiceServer() {}

// UnsafeSubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubServiceServer will
// result in compilation errors.
type UnsafeSubServiceServer interface {
	mustEmbedUnimplementedSubServiceServer()
}

func RegisterSubServiceServer(s grpc.ServiceRegistrar, srv SubServiceServer) {
	s.RegisterService(&SubService_ServiceDesc, srv)
}

func _SubService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubService_SubscribeTopic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTopicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubServiceServer).SubscribeTopic(m, &subServiceSubscribeTopicServer{stream})
}

type SubService_SubscribeTopicServer interface {
	Send(*SubscribeTopicResponse) error
	grpc.ServerStream
}

type subServiceSubscribeTopicServer struct {
	grpc.ServerStream
}

func (x *subServiceSubscribeTopicServer) Send(m *SubscribeTopicResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SubService_ServiceDesc is the grpc.ServiceDesc for SubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sub.v1.SubService",
	HandlerType: (*SubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _SubService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeTopic",
			Handler:       _SubService_SubscribeTopic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sub/v1/sub.proto",
}
