// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: link/v1/link.proto

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
	LinkService_GetLevelRange_FullMethodName = "/link.v1.LinkService/GetLevelRange"
	LinkService_GetTopology_FullMethodName   = "/link.v1.LinkService/GetTopology"
	LinkService_Join_FullMethodName          = "/link.v1.LinkService/Join"
	LinkService_Subscribe_FullMethodName     = "/link.v1.LinkService/Subscribe"
)

// LinkServiceClient is the client API for LinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkServiceClient interface {
	GetLevelRange(ctx context.Context, in *GetLevelRangeRequest, opts ...grpc.CallOption) (*GetLevelRangeResponse, error)
	GetTopology(ctx context.Context, in *GetTopologyRequest, opts ...grpc.CallOption) (*GetTopologyResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (LinkService_SubscribeClient, error)
}

type linkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkServiceClient(cc grpc.ClientConnInterface) LinkServiceClient {
	return &linkServiceClient{cc}
}

func (c *linkServiceClient) GetLevelRange(ctx context.Context, in *GetLevelRangeRequest, opts ...grpc.CallOption) (*GetLevelRangeResponse, error) {
	out := new(GetLevelRangeResponse)
	err := c.cc.Invoke(ctx, LinkService_GetLevelRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) GetTopology(ctx context.Context, in *GetTopologyRequest, opts ...grpc.CallOption) (*GetTopologyResponse, error) {
	out := new(GetTopologyResponse)
	err := c.cc.Invoke(ctx, LinkService_GetTopology_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, LinkService_Join_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (LinkService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &LinkService_ServiceDesc.Streams[0], LinkService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &linkServiceSubscribeClient{stream}
	return x, nil
}

type LinkService_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type linkServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *linkServiceSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *linkServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LinkServiceServer is the server API for LinkService service.
// All implementations must embed UnimplementedLinkServiceServer
// for forward compatibility
type LinkServiceServer interface {
	GetLevelRange(context.Context, *GetLevelRangeRequest) (*GetLevelRangeResponse, error)
	GetTopology(context.Context, *GetTopologyRequest) (*GetTopologyResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Subscribe(LinkService_SubscribeServer) error
	mustEmbedUnimplementedLinkServiceServer()
}

// UnimplementedLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLinkServiceServer struct {
}

func (UnimplementedLinkServiceServer) GetLevelRange(context.Context, *GetLevelRangeRequest) (*GetLevelRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLevelRange not implemented")
}
func (UnimplementedLinkServiceServer) GetTopology(context.Context, *GetTopologyRequest) (*GetTopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopology not implemented")
}
func (UnimplementedLinkServiceServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedLinkServiceServer) Subscribe(LinkService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedLinkServiceServer) mustEmbedUnimplementedLinkServiceServer() {}

// UnsafeLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServiceServer will
// result in compilation errors.
type UnsafeLinkServiceServer interface {
	mustEmbedUnimplementedLinkServiceServer()
}

func RegisterLinkServiceServer(s grpc.ServiceRegistrar, srv LinkServiceServer) {
	s.RegisterService(&LinkService_ServiceDesc, srv)
}

func _LinkService_GetLevelRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLevelRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).GetLevelRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_GetLevelRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).GetLevelRange(ctx, req.(*GetLevelRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_GetTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).GetTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_GetTopology_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).GetTopology(ctx, req.(*GetTopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LinkServiceServer).Subscribe(&linkServiceSubscribeServer{stream})
}

type LinkService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type linkServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *linkServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *linkServiceSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LinkService_ServiceDesc is the grpc.ServiceDesc for LinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "link.v1.LinkService",
	HandlerType: (*LinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLevelRange",
			Handler:    _LinkService_GetLevelRange_Handler,
		},
		{
			MethodName: "GetTopology",
			Handler:    _LinkService_GetTopology_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _LinkService_Join_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _LinkService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "link/v1/link.proto",
}