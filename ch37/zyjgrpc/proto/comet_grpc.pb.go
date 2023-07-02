// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/comet.proto

package zyjgrpc

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

// CometClient is the client API for Comet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CometClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Comet_StreamClient, error)
	// PushMsg push by key or mid
	PushMsg(ctx context.Context, opts ...grpc.CallOption) (Comet_PushMsgClient, error)
	// Broadcast send to every enrity
	Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (Comet_BroadcastClient, error)
	// BroadcastGroup broadcast to one group
	BroadcastGroup(ctx context.Context, in *BroadcastGroupReq, opts ...grpc.CallOption) (Comet_BroadcastGroupClient, error)
	// Groups get all groups
	Groups(ctx context.Context, in *GroupsReq, opts ...grpc.CallOption) (Comet_GroupsClient, error)
}

type cometClient struct {
	cc grpc.ClientConnInterface
}

func NewCometClient(cc grpc.ClientConnInterface) CometClient {
	return &cometClient{cc}
}

func (c *cometClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Comet_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Comet_ServiceDesc.Streams[0], "/v1.Comet/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &cometStreamClient{stream}
	return x, nil
}

type Comet_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type cometStreamClient struct {
	grpc.ClientStream
}

func (x *cometStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cometStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cometClient) PushMsg(ctx context.Context, opts ...grpc.CallOption) (Comet_PushMsgClient, error) {
	stream, err := c.cc.NewStream(ctx, &Comet_ServiceDesc.Streams[1], "/v1.Comet/PushMsg", opts...)
	if err != nil {
		return nil, err
	}
	x := &cometPushMsgClient{stream}
	return x, nil
}

type Comet_PushMsgClient interface {
	Send(*PushMsgReq) error
	CloseAndRecv() (*PushMsgReply, error)
	grpc.ClientStream
}

type cometPushMsgClient struct {
	grpc.ClientStream
}

func (x *cometPushMsgClient) Send(m *PushMsgReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cometPushMsgClient) CloseAndRecv() (*PushMsgReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PushMsgReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cometClient) Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (Comet_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &Comet_ServiceDesc.Streams[2], "/v1.Comet/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &cometBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Comet_BroadcastClient interface {
	Recv() (*BroadcastReply, error)
	grpc.ClientStream
}

type cometBroadcastClient struct {
	grpc.ClientStream
}

func (x *cometBroadcastClient) Recv() (*BroadcastReply, error) {
	m := new(BroadcastReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cometClient) BroadcastGroup(ctx context.Context, in *BroadcastGroupReq, opts ...grpc.CallOption) (Comet_BroadcastGroupClient, error) {
	stream, err := c.cc.NewStream(ctx, &Comet_ServiceDesc.Streams[3], "/v1.Comet/BroadcastGroup", opts...)
	if err != nil {
		return nil, err
	}
	x := &cometBroadcastGroupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Comet_BroadcastGroupClient interface {
	Recv() (*BroadcastGroupReply, error)
	grpc.ClientStream
}

type cometBroadcastGroupClient struct {
	grpc.ClientStream
}

func (x *cometBroadcastGroupClient) Recv() (*BroadcastGroupReply, error) {
	m := new(BroadcastGroupReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cometClient) Groups(ctx context.Context, in *GroupsReq, opts ...grpc.CallOption) (Comet_GroupsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Comet_ServiceDesc.Streams[4], "/v1.Comet/Groups", opts...)
	if err != nil {
		return nil, err
	}
	x := &cometGroupsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Comet_GroupsClient interface {
	Recv() (*GroupsReply, error)
	grpc.ClientStream
}

type cometGroupsClient struct {
	grpc.ClientStream
}

func (x *cometGroupsClient) Recv() (*GroupsReply, error) {
	m := new(GroupsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CometServer is the server API for Comet service.
// All implementations must embed UnimplementedCometServer
// for forward compatibility
type CometServer interface {
	Stream(Comet_StreamServer) error
	// PushMsg push by key or mid
	PushMsg(Comet_PushMsgServer) error
	// Broadcast send to every enrity
	Broadcast(*BroadcastReq, Comet_BroadcastServer) error
	// BroadcastGroup broadcast to one group
	BroadcastGroup(*BroadcastGroupReq, Comet_BroadcastGroupServer) error
	// Groups get all groups
	Groups(*GroupsReq, Comet_GroupsServer) error
	mustEmbedUnimplementedCometServer()
}

// UnimplementedCometServer must be embedded to have forward compatible implementations.
type UnimplementedCometServer struct {
}

func (UnimplementedCometServer) Stream(Comet_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedCometServer) PushMsg(Comet_PushMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}
func (UnimplementedCometServer) Broadcast(*BroadcastReq, Comet_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedCometServer) BroadcastGroup(*BroadcastGroupReq, Comet_BroadcastGroupServer) error {
	return status.Errorf(codes.Unimplemented, "method BroadcastGroup not implemented")
}
func (UnimplementedCometServer) Groups(*GroupsReq, Comet_GroupsServer) error {
	return status.Errorf(codes.Unimplemented, "method Groups not implemented")
}
func (UnimplementedCometServer) mustEmbedUnimplementedCometServer() {}

// UnsafeCometServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CometServer will
// result in compilation errors.
type UnsafeCometServer interface {
	mustEmbedUnimplementedCometServer()
}

func RegisterCometServer(s grpc.ServiceRegistrar, srv CometServer) {
	s.RegisterService(&Comet_ServiceDesc, srv)
}

func _Comet_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CometServer).Stream(&cometStreamServer{stream})
}

type Comet_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type cometStreamServer struct {
	grpc.ServerStream
}

func (x *cometStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cometStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Comet_PushMsg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CometServer).PushMsg(&cometPushMsgServer{stream})
}

type Comet_PushMsgServer interface {
	SendAndClose(*PushMsgReply) error
	Recv() (*PushMsgReq, error)
	grpc.ServerStream
}

type cometPushMsgServer struct {
	grpc.ServerStream
}

func (x *cometPushMsgServer) SendAndClose(m *PushMsgReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cometPushMsgServer) Recv() (*PushMsgReq, error) {
	m := new(PushMsgReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Comet_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CometServer).Broadcast(m, &cometBroadcastServer{stream})
}

type Comet_BroadcastServer interface {
	Send(*BroadcastReply) error
	grpc.ServerStream
}

type cometBroadcastServer struct {
	grpc.ServerStream
}

func (x *cometBroadcastServer) Send(m *BroadcastReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Comet_BroadcastGroup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastGroupReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CometServer).BroadcastGroup(m, &cometBroadcastGroupServer{stream})
}

type Comet_BroadcastGroupServer interface {
	Send(*BroadcastGroupReply) error
	grpc.ServerStream
}

type cometBroadcastGroupServer struct {
	grpc.ServerStream
}

func (x *cometBroadcastGroupServer) Send(m *BroadcastGroupReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Comet_Groups_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GroupsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CometServer).Groups(m, &cometGroupsServer{stream})
}

type Comet_GroupsServer interface {
	Send(*GroupsReply) error
	grpc.ServerStream
}

type cometGroupsServer struct {
	grpc.ServerStream
}

func (x *cometGroupsServer) Send(m *GroupsReply) error {
	return x.ServerStream.SendMsg(m)
}

// Comet_ServiceDesc is the grpc.ServiceDesc for Comet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Comet",
	HandlerType: (*CometServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Comet_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PushMsg",
			Handler:       _Comet_PushMsg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _Comet_Broadcast_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BroadcastGroup",
			Handler:       _Comet_BroadcastGroup_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Groups",
			Handler:       _Comet_Groups_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/comet.proto",
}