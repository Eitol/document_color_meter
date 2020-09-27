import os
import traceback
from concurrent import futures
import logging

import grpc
from google.protobuf import any_pb2
from grpc_status import rpc_status

import app.binarizer_pb2_grpc
import app.binarizer_pb2
from google.rpc import code_pb2, status_pb2, error_details_pb2

from app.binarizer import Binarizer, Method

DEFAULT_PORT = 50053


class Server(app.binarizer_pb2_grpc.BinarizerServicer):
    def Binarize(self, request, context) -> None:
        if len(request.images) == 0:
            context.abort_with_status(status_pb2.Status(
                code=code_pb2.INVALID_ARGUMENT,
                message='empty "images" argument',
            ))
            return
        try:
            response, out_path = Binarizer.binarize(request.images, Method(request.binarization_method))
        except (ValueError, Exception) as e:
            self._abort(context, 'error while decoding the file', code_pb2.INVALID_ARGUMENT, str(e))
            return
        return app.binarizer_pb2.BinarizeResponse(
            images=response,
            out_path=out_path,
        )
    
    @staticmethod
    def _abort(context, err_msg: str, code: int, detail_msg: str = ""):
        detail = any_pb2.Any()
        detail.Pack(
            error_details_pb2.DebugInfo(
                stack_entries=traceback.format_stack(),
                detail=detail_msg,
            )
        )
        rich_status = status_pb2.Status(
            code=code,
            message=err_msg,
            details=[detail]
        )
        context.abort_with_status(rpc_status.to_status(rich_status))


def get_port() -> int:
    port = os.getenv("PORT")
    return int(port) if port is not None else DEFAULT_PORT


def serve():
    options = [
        ('grpc.max_send_message_length', 512 * 1024 * 1024),
        ('grpc.max_receive_message_length', 512 * 1024 * 1024)
    ]
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10), options=options)
    app.binarizer_pb2_grpc.add_BinarizerServicer_to_server(Server(), server)
    server.add_insecure_port(f"[::]:{get_port()}")
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
