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
import signal
import sys
import time

def handler(signum, frame):
    exit(0)

signal.signal(signal.SIGINT, handler)

async def run_stream(id_to_subscribe: str, client_name: str) -> None:
    async with grpc.aio.insecure_channel("localhost:9000") as channel:
        stub = grpcproject_pb2_grpc.GrpcProjectStub(channel)
        id_ = int(id_to_subscribe)

        # Direct read from the stub
        hello_stream = stub.Subscribe(
            grpcproject_pb2.SubscribeRequest(name=client_name, subscribtion_id=id_))
        hello_stream2 = stub.Subscribe(
            grpcproject_pb2.SubscribeRequest(name=client_name, subscribtion_id=2))
        while True:
            try:
                response = await hello_stream.read()
                if response != grpc.aio.EOF:
                    print(f"Event notification recived:\n \
                            {response.subscribtion_id}\n \
                            {response.message}\n \
                            {response.time}\n \
                            {response.times}\n \
                            {response.test_enum}\n")
                response = await hello_stream2.read()
                if response != grpc.aio.EOF:
                    print(f"Event notification recived:\n \
                            {response.subscribtion_id}\n \
                            {response.message}\n \
                            {response.time}\n \
                            {response.times}\n \
                            {response.test_enum}\n")
            except:
                print("waiting for server to start")
                hello_stream = stub.Subscribe(
                    grpcproject_pb2.SubscribeRequest(name=client_name, subscribtion_id=id_))
                hello_stream2 = stub.Subscribe(
                    grpcproject_pb2.SubscribeRequest(name=client_name, subscribtion_id=3))
                time.sleep(1)


if __name__ == '__main__':
    logging.basicConfig()
    while True:
        subID = int(input("subscription num: "))
        clientName = input("client name: ")
        asyncio.run(run_stream(subID, clientName))
