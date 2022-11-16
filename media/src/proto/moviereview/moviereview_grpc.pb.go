// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/moviereview.proto

package moviereview

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

// MovieReviewServiceClient is the client API for MovieReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieReviewServiceClient interface {
	UploadMovieReview(ctx context.Context, in *UploadMovieReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReadMovieReviews(ctx context.Context, in *ReadMovieReviewsRequest, opts ...grpc.CallOption) (*ReadMovieReviewsRespond, error)
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type movieReviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieReviewServiceClient(cc grpc.ClientConnInterface) MovieReviewServiceClient {
	return &movieReviewServiceClient{cc}
}

func (c *movieReviewServiceClient) UploadMovieReview(ctx context.Context, in *UploadMovieReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/microless.media.moviereview.MovieReviewService/UploadMovieReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieReviewServiceClient) ReadMovieReviews(ctx context.Context, in *ReadMovieReviewsRequest, opts ...grpc.CallOption) (*ReadMovieReviewsRespond, error) {
	out := new(ReadMovieReviewsRespond)
	err := c.cc.Invoke(ctx, "/microless.media.moviereview.MovieReviewService/ReadMovieReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieReviewServiceClient) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/microless.media.moviereview.MovieReviewService/CreateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieReviewServiceServer is the server API for MovieReviewService service.
// All implementations must embed UnimplementedMovieReviewServiceServer
// for forward compatibility
type MovieReviewServiceServer interface {
	UploadMovieReview(context.Context, *UploadMovieReviewRequest) (*emptypb.Empty, error)
	ReadMovieReviews(context.Context, *ReadMovieReviewsRequest) (*ReadMovieReviewsRespond, error)
	CreateMovie(context.Context, *CreateMovieRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMovieReviewServiceServer()
}

// UnimplementedMovieReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieReviewServiceServer struct {
}

func (UnimplementedMovieReviewServiceServer) UploadMovieReview(context.Context, *UploadMovieReviewRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadMovieReview not implemented")
}
func (UnimplementedMovieReviewServiceServer) ReadMovieReviews(context.Context, *ReadMovieReviewsRequest) (*ReadMovieReviewsRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMovieReviews not implemented")
}
func (UnimplementedMovieReviewServiceServer) CreateMovie(context.Context, *CreateMovieRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
}
func (UnimplementedMovieReviewServiceServer) mustEmbedUnimplementedMovieReviewServiceServer() {}

// UnsafeMovieReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieReviewServiceServer will
// result in compilation errors.
type UnsafeMovieReviewServiceServer interface {
	mustEmbedUnimplementedMovieReviewServiceServer()
}

func RegisterMovieReviewServiceServer(s grpc.ServiceRegistrar, srv MovieReviewServiceServer) {
	s.RegisterService(&MovieReviewService_ServiceDesc, srv)
}

func _MovieReviewService_UploadMovieReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMovieReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieReviewServiceServer).UploadMovieReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.media.moviereview.MovieReviewService/UploadMovieReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieReviewServiceServer).UploadMovieReview(ctx, req.(*UploadMovieReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieReviewService_ReadMovieReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMovieReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieReviewServiceServer).ReadMovieReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.media.moviereview.MovieReviewService/ReadMovieReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieReviewServiceServer).ReadMovieReviews(ctx, req.(*ReadMovieReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieReviewService_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieReviewServiceServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microless.media.moviereview.MovieReviewService/CreateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieReviewServiceServer).CreateMovie(ctx, req.(*CreateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieReviewService_ServiceDesc is the grpc.ServiceDesc for MovieReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microless.media.moviereview.MovieReviewService",
	HandlerType: (*MovieReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadMovieReview",
			Handler:    _MovieReviewService_UploadMovieReview_Handler,
		},
		{
			MethodName: "ReadMovieReviews",
			Handler:    _MovieReviewService_ReadMovieReviews_Handler,
		},
		{
			MethodName: "CreateMovie",
			Handler:    _MovieReviewService_CreateMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/moviereview.proto",
}