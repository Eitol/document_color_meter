import logging
import os
from os.path import join

import grpc

# import the generated classes
from google.protobuf.pyext._message import RepeatedScalarContainer
from grpc._channel import _InactiveRpcError

import app.pdf_to_image_pb2_grpc
import app.pdf_to_image_pb2


def do_request():
    # open a gRPC channel
    options = [
        ('grpc.max_send_message_length', 512 * 1024 * 1024),
        ('grpc.max_receive_message_length', 512 * 1024 * 1024)
    ]
    channel = grpc.insecure_channel('localhost:50052', options=options)
    
    # create a stub (client)
    
    stub = app.pdf_to_image_pb2_grpc.PdfToImageServiceStub(channel)
    
    # create a valid request message
    with open(join("testdata", "test.pdf"), "rb") as f:
        request_body = app.pdf_to_image_pb2.PdfToImageRequest(file=f.read())
    try:
        response: app.pdf_to_image_pb2.PdfToImageResponse = stub.Convert(request_body)
    except _InactiveRpcError as e:
        logging.error("the server is down")
        return
    except (ValueError, Exception) as e:
        logging.error(e)
        return
    
    if response is None or response.pages is None or len(response.pages) == 0:
        logging.error("empty response")
        return
    out_path = os.path.join("testdata", "out")
    os.makedirs(os.path.join("testdata", "out"), exist_ok=True)
    count = 0
    for page in response.pages:
        out_page_img_path = os.path.join(out_path, f"page_{count}.jpg")
        with open(out_page_img_path, "wb") as f:
            f.write(page)
        count += 1


if __name__ == '__main__':
    do_request()
