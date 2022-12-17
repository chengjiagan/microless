# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import search_pb2 as proto_dot_search__pb2


class SearchServiceStub(object):
    """Search service returns best hotel chocies for a user.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Search = channel.unary_unary(
                '/microless.hotelreservation.search.SearchService/Search',
                request_serializer=proto_dot_search__pb2.SearchRequest.SerializeToString,
                response_deserializer=proto_dot_search__pb2.SearchRespond.FromString,
                )


class SearchServiceServicer(object):
    """Search service returns best hotel chocies for a user.
    """

    def Search(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SearchServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Search': grpc.unary_unary_rpc_method_handler(
                    servicer.Search,
                    request_deserializer=proto_dot_search__pb2.SearchRequest.FromString,
                    response_serializer=proto_dot_search__pb2.SearchRespond.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'microless.hotelreservation.search.SearchService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SearchService(object):
    """Search service returns best hotel chocies for a user.
    """

    @staticmethod
    def Search(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.hotelreservation.search.SearchService/Search',
            proto_dot_search__pb2.SearchRequest.SerializeToString,
            proto_dot_search__pb2.SearchRespond.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
