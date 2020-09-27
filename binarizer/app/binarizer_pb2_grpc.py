# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import binarizer_pb2 as binarizer__pb2


class BinarizerStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Binarize = channel.unary_unary(
                '/binarizer.Binarizer/Binarize',
                request_serializer=binarizer__pb2.BinarizeRequest.SerializeToString,
                response_deserializer=binarizer__pb2.BinarizeResponse.FromString,
                )


class BinarizerServicer(object):
    """Missing associated documentation comment in .proto file"""

    def Binarize(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BinarizerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Binarize': grpc.unary_unary_rpc_method_handler(
                    servicer.Binarize,
                    request_deserializer=binarizer__pb2.BinarizeRequest.FromString,
                    response_serializer=binarizer__pb2.BinarizeResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'binarizer.Binarizer', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Binarizer(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def Binarize(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/binarizer.Binarizer/Binarize',
            binarizer__pb2.BinarizeRequest.SerializeToString,
            binarizer__pb2.BinarizeResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)