from __future__ import print_function

import logging

import grpc
import hello_world_pb2
import hello_world_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('127.0.0.1:8000') as channel:
        stub = hello_world_pb2_grpc.GreeterStub(channel)
        resp = stub.SayHello(hello_world_pb2.HelloRequest(name='jokereven'))
    print("Greeter client received: " + resp.reply)


if __name__ == '__main__':
    logging.basicConfig()
    run()
