"""
Inspiracja: https://grpc.io/docs/languages/python/quickstart/
"""

from __future__ import print_function

import asyncio
import logging

import grpc
import grpcproject_pb2
import grpcproject_pb2_grpc

import random

async def run() -> None:


    # while True: … 
    # id_ -> co ile sekund wysyłamy powiadomienie o akcji
    # akcja -> czas działania serwera i operacja modulo
    print("Will try to greet world ...")

    async with grpc.aio.insecure_channel("localhost:9000") as channel:
        stub = grpcproject_pb2_grpc.GrpcProjectStub(channel)
        id_ = random.randint(1,5)

        # Read from an async generator
        async for response in stub.FetchResponse(
            grpcproject_pb2.Request(id=id_)):
            print("Greeter client received from async generator: " +
                  response.result)

        # Direct read from the stub
        hello_stream = stub.FetchResponse(
            grpcproject_pb2.Request(id=id_))
        while True:
            response = await hello_stream.read()
            if response == grpc.aio.EOF:
                break
            print("Greeter client received from direct read: " +
                  response.result)

if __name__ == '__main__':
    logging.basicConfig()
    asyncio.run(run())
