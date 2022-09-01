# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from proto import usertimeline_pb2 as proto_dot_usertimeline__pb2


class UserTimelineServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.WriteUserTimeline = channel.unary_unary(
                '/microless.socialnetwork.usertimeline.UserTimelineService/WriteUserTimeline',
                request_serializer=proto_dot_usertimeline__pb2.WriteUserTimelineRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.ReadUserTimeline = channel.unary_unary(
                '/microless.socialnetwork.usertimeline.UserTimelineService/ReadUserTimeline',
                request_serializer=proto_dot_usertimeline__pb2.ReadUserTimelineRequest.SerializeToString,
                response_deserializer=proto_dot_usertimeline__pb2.ReadUserTimelineRespond.FromString,
                )
        self.InsertUser = channel.unary_unary(
                '/microless.socialnetwork.usertimeline.UserTimelineService/InsertUser',
                request_serializer=proto_dot_usertimeline__pb2.InsertUserResquest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class UserTimelineServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def WriteUserTimeline(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReadUserTimeline(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def InsertUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserTimelineServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'WriteUserTimeline': grpc.unary_unary_rpc_method_handler(
                    servicer.WriteUserTimeline,
                    request_deserializer=proto_dot_usertimeline__pb2.WriteUserTimelineRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'ReadUserTimeline': grpc.unary_unary_rpc_method_handler(
                    servicer.ReadUserTimeline,
                    request_deserializer=proto_dot_usertimeline__pb2.ReadUserTimelineRequest.FromString,
                    response_serializer=proto_dot_usertimeline__pb2.ReadUserTimelineRespond.SerializeToString,
            ),
            'InsertUser': grpc.unary_unary_rpc_method_handler(
                    servicer.InsertUser,
                    request_deserializer=proto_dot_usertimeline__pb2.InsertUserResquest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'microless.socialnetwork.usertimeline.UserTimelineService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class UserTimelineService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def WriteUserTimeline(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.socialnetwork.usertimeline.UserTimelineService/WriteUserTimeline',
            proto_dot_usertimeline__pb2.WriteUserTimelineRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ReadUserTimeline(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.socialnetwork.usertimeline.UserTimelineService/ReadUserTimeline',
            proto_dot_usertimeline__pb2.ReadUserTimelineRequest.SerializeToString,
            proto_dot_usertimeline__pb2.ReadUserTimelineRespond.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def InsertUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.socialnetwork.usertimeline.UserTimelineService/InsertUser',
            proto_dot_usertimeline__pb2.InsertUserResquest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
