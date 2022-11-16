// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/castinfo.proto

package castinfo

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

// CastInfoServiceClient is the client API for CastInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CastInfoServiceClient interface {
	WriteCastInfo(ctx context.Context, in *WriteCastInfoRequest, opts ...grpc.CallOption) (*WriteCastInfoRespond, error)
	ReadCastInfo(ctx context.Context, in *ReadCastInfoRequest, opts ...grpc.CallOption) (*ReadCastInfoRespond, error)
}

type castInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCastInfoServiceClient(cc grpc.ClientConnInterface) CastInfoServiceClient {
	return &castInfoServiceClient{cc}
}

func (c *castInfoServiceClient) WriteCastInfo(ctx context.Context, in *WriteCastInfoRequest, opts ...grpc.CallOption) (*WriteCastInfoRespond, error) {
	out := new(WriteCastInfoRespond)
	err := c.cc.Invoke(ctx, "/microless.media.castinfo.CastInfoService/WriteCastInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *castInfoServiceClient) ReadCastInfo(ctx context.Context, in *ReadCastInfoRequest, opts ...grpc.CallOption) (*ReadCastInfoRespond, error) {
	out := new(ReadCastInfoRespond)
	err := c.cc.Invoke(ctx, "/microless.media.castinfo.CastInfoService/ReadCastInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CastInfoServiceServer is the server API for CastInfoService service.
// All implementations must embed UnimplementedCastInfoServiceServer
// for forward compatibility
type CastInfoServiceServer interface {
	WriteCastInfo(context.Context, *WriteCastInfoRequest) (*WriteCastInfoRespond, error)
	ReadCastInfo(context.Context, *ReadCastInfoRequest) (*ReadCastInfoRespond, error)
	mustEmbedUnimplementedCastInfoServiceServer()
}

// UnimplementedCastInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCastInfoServiceServer struct {
}

func (UnimplementedCastInfoServiceServer) WriteCastInfo(context.Context, *WriteCastInfoRequest) (*WriteCastInfoRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteCastInfo not implemented")
}
func (UnimplementedCastInfoServiceServer) ReadCastInfo(context.Context, *ReadCastInfoRequest) (*ReadCastInfoRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCastInfo not implemented")
}
func (UnimplementedCastInfoServiceServer) mustEmbedUnimplementedCastInfoServiceServer() {}

// UnsafeCastInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CastInfoServiceServer will
// result in compilation errors.
type UnsafeCastInfoServiceServer interface {
	mustEmbedUnimplementedCastInfoServiceServer()
}

func RegisterCastInfoServiceServer(s grpc.ServiceRegistrar, srv CastInfoServiceServer) {
	s.RegisterService(&CastInfoService_ServiceDesc, srv)
}

func _CastInfoService_WriteCastInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteCastInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CastInfoServiceServer).WriteCastInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.media.castinfo.CastInfoService/WriteCastInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CastInfoServiceServer).WriteCastInfo(ctx, req.(*WriteCastInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CastInfoService_ReadCastInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCastInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CastInfoServiceServer).ReadCastInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.media.castinfo.CastInfoService/ReadCastInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CastInfoServiceServer).ReadCastInfo(ctx, req.(*ReadCastInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CastInfoService_ServiceDesc is the grpc.ServiceDesc for CastInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CastInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microless.media.castinfo.CastInfoService",
	HandlerType: (*CastInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteCastInfo",
			Handler:    _CastInfoService_WriteCastInfo_Handler,
		},
		{
			MethodName: "ReadCastInfo",
			Handler:    _CastInfoService_ReadCastInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/castinfo.proto",
}