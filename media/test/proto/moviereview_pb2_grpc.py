# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from proto import moviereview_pb2 as proto_dot_moviereview__pb2


class MovieReviewServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UploadMovieReview = channel.unary_unary(
                '/microless.media.moviereview.MovieReviewService/UploadMovieReview',
                request_serializer=proto_dot_moviereview__pb2.UploadMovieReviewRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ReadMovieReviews = channel.unary_unary(
                '/microless.media.moviereview.MovieReviewService/ReadMovieReviews',
                request_serializer=proto_dot_moviereview__pb2.ReadMovieReviewsRequest.SerializeToString,
                response_deserializer=proto_dot_moviereview__pb2.ReadMovieReviewsRespond.FromString,
                )
        self.CreateMovie = channel.unary_unary(
                '/microless.media.moviereview.MovieReviewService/CreateMovie',
                request_serializer=proto_dot_moviereview__pb2.CreateMovieRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class MovieReviewServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def UploadMovieReview(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReadMovieReviews(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateMovie(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MovieReviewServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UploadMovieReview': grpc.unary_unary_rpc_method_handler(
                    servicer.UploadMovieReview,
                    request_deserializer=proto_dot_moviereview__pb2.UploadMovieReviewRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ReadMovieReviews': grpc.unary_unary_rpc_method_handler(
                    servicer.ReadMovieReviews,
                    request_deserializer=proto_dot_moviereview__pb2.ReadMovieReviewsRequest.FromString,
                    response_serializer=proto_dot_moviereview__pb2.ReadMovieReviewsRespond.SerializeToString,
            ),
            'CreateMovie': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateMovie,
                    request_deserializer=proto_dot_moviereview__pb2.CreateMovieRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'microless.media.moviereview.MovieReviewService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MovieReviewService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def UploadMovieReview(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.media.moviereview.MovieReviewService/UploadMovieReview',
            proto_dot_moviereview__pb2.UploadMovieReviewRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ReadMovieReviews(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.media.moviereview.MovieReviewService/ReadMovieReviews',
            proto_dot_moviereview__pb2.ReadMovieReviewsRequest.SerializeToString,
            proto_dot_moviereview__pb2.ReadMovieReviewsRespond.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateMovie(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.media.moviereview.MovieReviewService/CreateMovie',
            proto_dot_moviereview__pb2.CreateMovieRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
