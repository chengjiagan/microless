#! /usr/bin/env python

import grpc
import json
import pandas as pd
import argparse
import model_pb2
from os import path
from datetime import datetime
from query_pb2_grpc import QueryServiceStub
from query_pb2 import FindTracesRequest, TraceQueryParameters
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.text_format import MessageToString

# get spans from jaeger
def get_spans(addr: str, start: datetime, end: datetime) -> list[model_pb2.Span]:
    with grpc.insecure_channel(addr) as chan:
        stub = QueryServiceStub(chan)

        # construct the query request
        endts = Timestamp()
        endts.FromDatetime(end)
        startts = Timestamp()
        startts.FromDatetime(start)
        query = TraceQueryParameters(
            start_time_min=startts,
            start_time_max=endts,
            search_depth=100000,
        )
        req = FindTracesRequest(query=query)
        # search traces from jaeger
        resp = stub.FindTraces(req)

        return [span for chunk in resp for span in chunk.spans]

# convert list of protos to list of dicts
def protos_to_strings(spans: list[model_pb2.Span]) -> list[str]:
    return [
        MessageToString(s, as_one_line=True)
        for s in spans
    ]

# get start and end timestamp from load record
def get_time(record: str) -> tuple[datetime, datetime]:
    df = pd.read_csv(record)
    # timestamp in record are in miliseconds, local in jaeger
    start = datetime.utcfromtimestamp(df['start'].min() / 1000)
    end = datetime.utcfromtimestamp(df['end'].max() / 1000)
    return (start, end)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-a', '--address', type=str, help='address to the jaeger server', required=True)
    parser.add_argument('-i', '--input', type=str, help='path to the load record file', required=True)
    parser.add_argument('-o', '--output', type=str, help='path to the output file')
    args = parser.parse_args()

    # generate an ouput path from 
    if args.output is None:
        fn, _ = path.splitext(path.basename(args.input))
        args.output = fn.replace('load', 'jaeger') + '.json'

    addr = args.address
    start, end = get_time(args.input)
    spans = protos_to_strings(get_spans(addr, start, end))
    with open(args.output, 'w') as f:
        for s in spans:
            f.write(s)
            f.write('\n')