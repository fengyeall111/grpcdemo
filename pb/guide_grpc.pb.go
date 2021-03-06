// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// GuideClient is the client API for Guide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuideClient interface {
	// 定义Confession方法
	S1R1(ctx context.Context, in *GuideRequest, opts ...grpc.CallOption) (*GuideResponse, error)
	S1Rn(ctx context.Context, in *GuideRequest, opts ...grpc.CallOption) (Guide_S1RnClient, error)
	SnR1(ctx context.Context, opts ...grpc.CallOption) (Guide_SnR1Client, error)
	SnRn(ctx context.Context, opts ...grpc.CallOption) (Guide_SnRnClient, error)
}

type guideClient struct {
	cc grpc.ClientConnInterface
}

func NewGuideClient(cc grpc.ClientConnInterface) GuideClient {
	return &guideClient{cc}
}

func (c *guideClient) S1R1(ctx context.Context, in *GuideRequest, opts ...grpc.CallOption) (*GuideResponse, error) {
	out := new(GuideResponse)
	err := c.cc.Invoke(ctx, "/pb.Guide/S1R1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideClient) S1Rn(ctx context.Context, in *GuideRequest, opts ...grpc.CallOption) (Guide_S1RnClient, error) {
	stream, err := c.cc.NewStream(ctx, &Guide_ServiceDesc.Streams[0], "/pb.Guide/S1Rn", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideS1RnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Guide_S1RnClient interface {
	Recv() (*GuideResponse, error)
	grpc.ClientStream
}

type guideS1RnClient struct {
	grpc.ClientStream
}

func (x *guideS1RnClient) Recv() (*GuideResponse, error) {
	m := new(GuideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guideClient) SnR1(ctx context.Context, opts ...grpc.CallOption) (Guide_SnR1Client, error) {
	stream, err := c.cc.NewStream(ctx, &Guide_ServiceDesc.Streams[1], "/pb.Guide/SnR1", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideSnR1Client{stream}
	return x, nil
}

type Guide_SnR1Client interface {
	Send(*GuideRequest) error
	CloseAndRecv() (*GuideResponse, error)
	grpc.ClientStream
}

type guideSnR1Client struct {
	grpc.ClientStream
}

func (x *guideSnR1Client) Send(m *GuideRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guideSnR1Client) CloseAndRecv() (*GuideResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GuideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guideClient) SnRn(ctx context.Context, opts ...grpc.CallOption) (Guide_SnRnClient, error) {
	stream, err := c.cc.NewStream(ctx, &Guide_ServiceDesc.Streams[2], "/pb.Guide/SnRn", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideSnRnClient{stream}
	return x, nil
}

type Guide_SnRnClient interface {
	Send(*GuideRequest) error
	Recv() (*GuideResponse, error)
	grpc.ClientStream
}

type guideSnRnClient struct {
	grpc.ClientStream
}

func (x *guideSnRnClient) Send(m *GuideRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *guideSnRnClient) Recv() (*GuideResponse, error) {
	m := new(GuideResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GuideServer is the server API for Guide service.
// All implementations must embed UnimplementedGuideServer
// for forward compatibility
type GuideServer interface {
	// 定义Confession方法
	S1R1(context.Context, *GuideRequest) (*GuideResponse, error)
	S1Rn(*GuideRequest, Guide_S1RnServer) error
	SnR1(Guide_SnR1Server) error
	SnRn(Guide_SnRnServer) error
	mustEmbedUnimplementedGuideServer()
}

// UnimplementedGuideServer must be embedded to have forward compatible implementations.
type UnimplementedGuideServer struct {
}

func (UnimplementedGuideServer) S1R1(context.Context, *GuideRequest) (*GuideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method S1R1 not implemented")
}
func (UnimplementedGuideServer) S1Rn(*GuideRequest, Guide_S1RnServer) error {
	return status.Errorf(codes.Unimplemented, "method S1Rn not implemented")
}
func (UnimplementedGuideServer) SnR1(Guide_SnR1Server) error {
	return status.Errorf(codes.Unimplemented, "method SnR1 not implemented")
}
func (UnimplementedGuideServer) SnRn(Guide_SnRnServer) error {
	return status.Errorf(codes.Unimplemented, "method SnRn not implemented")
}
func (UnimplementedGuideServer) mustEmbedUnimplementedGuideServer() {}

// UnsafeGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuideServer will
// result in compilation errors.
type UnsafeGuideServer interface {
	mustEmbedUnimplementedGuideServer()
}

func RegisterGuideServer(s grpc.ServiceRegistrar, srv GuideServer) {
	s.RegisterService(&Guide_ServiceDesc, srv)
}

func _Guide_S1R1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServer).S1R1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Guide/S1R1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServer).S1R1(ctx, req.(*GuideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guide_S1Rn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GuideRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GuideServer).S1Rn(m, &guideS1RnServer{stream})
}

type Guide_S1RnServer interface {
	Send(*GuideResponse) error
	grpc.ServerStream
}

type guideS1RnServer struct {
	grpc.ServerStream
}

func (x *guideS1RnServer) Send(m *GuideResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Guide_SnR1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuideServer).SnR1(&guideSnR1Server{stream})
}

type Guide_SnR1Server interface {
	SendAndClose(*GuideResponse) error
	Recv() (*GuideRequest, error)
	grpc.ServerStream
}

type guideSnR1Server struct {
	grpc.ServerStream
}

func (x *guideSnR1Server) SendAndClose(m *GuideResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guideSnR1Server) Recv() (*GuideRequest, error) {
	m := new(GuideRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Guide_SnRn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GuideServer).SnRn(&guideSnRnServer{stream})
}

type Guide_SnRnServer interface {
	Send(*GuideResponse) error
	Recv() (*GuideRequest, error)
	grpc.ServerStream
}

type guideSnRnServer struct {
	grpc.ServerStream
}

func (x *guideSnRnServer) Send(m *GuideResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *guideSnRnServer) Recv() (*GuideRequest, error) {
	m := new(GuideRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Guide_ServiceDesc is the grpc.ServiceDesc for Guide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Guide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Guide",
	HandlerType: (*GuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "S1R1",
			Handler:    _Guide_S1R1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "S1Rn",
			Handler:       _Guide_S1Rn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SnR1",
			Handler:       _Guide_SnR1_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SnRn",
			Handler:       _Guide_SnRn_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/guide.proto",
}
