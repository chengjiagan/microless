// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/hometimeline.proto

package hometimeline

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HomeTimelineServiceClient is the client API for HomeTimelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeTimelineServiceClient interface {
	ReadHomeTimeline(ctx context.Context, in *ReadHomeTimelineRequest, opts ...grpc.CallOption) (*ReadHomeTimelineRespond, error)
	WriteHomeTimeline(ctx context.Context, in *WriteHomeTimelineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InsertUser(ctx context.Context, in *InsertUserResquest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type homeTimelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeTimelineServiceClient(cc grpc.ClientConnInterface) HomeTimelineServiceClient {
	return &homeTimelineServiceClient{cc}
}

func (c *homeTimelineServiceClient) ReadHomeTimeline(ctx context.Context, in *ReadHomeTimelineRequest, opts ...grpc.CallOption) (*ReadHomeTimelineRespond, error) {
	out := new(ReadHomeTimelineRespond)
	err := c.cc.Invoke(ctx, "/microless.socialnetwork.hometimeline.HomeTimelineService/ReadHomeTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeTimelineServiceClient) WriteHomeTimeline(ctx context.Context, in *WriteHomeTimelineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/microless.socialnetwork.hometimeline.HomeTimelineService/WriteHomeTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeTimelineServiceClient) InsertUser(ctx context.Context, in *InsertUserResquest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/microless.socialnetwork.hometimeline.HomeTimelineService/InsertUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeTimelineServiceServer is the server API for HomeTimelineService service.
// All implementations must embed UnimplementedHomeTimelineServiceServer
// for forward compatibility
type HomeTimelineServiceServer interface {
	ReadHomeTimeline(context.Context, *ReadHomeTimelineRequest) (*ReadHomeTimelineRespond, error)
	WriteHomeTimeline(context.Context, *WriteHomeTimelineRequest) (*emptypb.Empty, error)
	InsertUser(context.Context, *InsertUserResquest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHomeTimelineServiceServer()
}

// UnimplementedHomeTimelineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHomeTimelineServiceServer struct {
}

func (UnimplementedHomeTimelineServiceServer) ReadHomeTimeline(context.Context, *ReadHomeTimelineRequest) (*ReadHomeTimelineRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadHomeTimeline not implemented")
}
func (UnimplementedHomeTimelineServiceServer) WriteHomeTimeline(context.Context, *WriteHomeTimelineRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteHomeTimeline not implemented")
}
func (UnimplementedHomeTimelineServiceServer) InsertUser(context.Context, *InsertUserResquest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedHomeTimelineServiceServer) mustEmbedUnimplementedHomeTimelineServiceServer() {}

// UnsafeHomeTimelineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeTimelineServiceServer will
// result in compilation errors.
type UnsafeHomeTimelineServiceServer interface {
	mustEmbedUnimplementedHomeTimelineServiceServer()
}

func RegisterHomeTimelineServiceServer(s grpc.ServiceRegistrar, srv HomeTimelineServiceServer) {
	s.RegisterService(&HomeTimelineService_ServiceDesc, srv)
}

func _HomeTimelineService_ReadHomeTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHomeTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTimelineServiceServer).ReadHomeTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.socialnetwork.hometimeline.HomeTimelineService/ReadHomeTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTimelineServiceServer).ReadHomeTimeline(ctx, req.(*ReadHomeTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeTimelineService_WriteHomeTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteHomeTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTimelineServiceServer).WriteHomeTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.socialnetwork.hometimeline.HomeTimelineService/WriteHomeTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTimelineServiceServer).WriteHomeTimeline(ctx, req.(*WriteHomeTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeTimelineService_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserResquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTimelineServiceServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.socialnetwork.hometimeline.HomeTimelineService/InsertUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTimelineServiceServer).InsertUser(ctx, req.(*InsertUserResquest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeTimelineService_ServiceDesc is the grpc.ServiceDesc for HomeTimelineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeTimelineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microless.socialnetwork.hometimeline.HomeTimelineService",
	HandlerType: (*HomeTimelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadHomeTimeline",
			Handler:    _HomeTimelineService_ReadHomeTimeline_Handler,
		},
		{
			MethodName: "WriteHomeTimeline",
			Handler:    _HomeTimelineService_WriteHomeTimeline_Handler,
		},
		{
			MethodName: "InsertUser",
			Handler:    _HomeTimelineService_InsertUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hometimeline.proto",
}
