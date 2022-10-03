# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import plot_pb2 as proto_dot_plot__pb2


class PlotServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.WritePlot = channel.unary_unary(
                '/microless.media.plot.PlotService/WritePlot',
                request_serializer=proto_dot_plot__pb2.WritePlotRequest.SerializeToString,
                response_deserializer=proto_dot_plot__pb2.WritePlotRespond.FromString,
                )
        self.ReadPlot = channel.unary_unary(
                '/microless.media.plot.PlotService/ReadPlot',
                request_serializer=proto_dot_plot__pb2.ReadPlotRequest.SerializeToString,
                response_deserializer=proto_dot_plot__pb2.ReadPlotRespond.FromString,
                )


class PlotServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def WritePlot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReadPlot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PlotServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'WritePlot': grpc.unary_unary_rpc_method_handler(
                    servicer.WritePlot,
                    request_deserializer=proto_dot_plot__pb2.WritePlotRequest.FromString,
                    response_serializer=proto_dot_plot__pb2.WritePlotRespond.SerializeToString,
            ),
            'ReadPlot': grpc.unary_unary_rpc_method_handler(
                    servicer.ReadPlot,
                    request_deserializer=proto_dot_plot__pb2.ReadPlotRequest.FromString,
                    response_serializer=proto_dot_plot__pb2.ReadPlotRespond.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'microless.media.plot.PlotService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PlotService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def WritePlot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.media.plot.PlotService/WritePlot',
            proto_dot_plot__pb2.WritePlotRequest.SerializeToString,
            proto_dot_plot__pb2.WritePlotRespond.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ReadPlot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/microless.media.plot.PlotService/ReadPlot',
            proto_dot_plot__pb2.ReadPlotRequest.SerializeToString,
            proto_dot_plot__pb2.ReadPlotRespond.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
