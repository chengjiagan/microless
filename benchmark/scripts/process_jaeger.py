#! /usr/bin/env python

import json
from ..fetch_jaeger import model_pb2
import pandas as pd
from datetime import datetime, timedelta
from google.protobuf.text_format import Parse

intervals = {}
children = {}
names = {}

MS = timedelta(milliseconds=1)

def strings_to_protos(spans: list[str]) -> list[model_pb2.Span]:
    return [
        Parse(s, model_pb2.Span)
        for s in spans
    ]

def compose_key(trace_id: bytes, span_id: bytes) -> str:
    return trace_id.hex() + '-' + span_id.hex()

def span_children(spans: list[model_pb2.Span]) -> dict[str, list[str]]:
    children = {}

    for span in spans:
        key = compose_key(span.trace_id, span.span_id)

        for ref in span.references:
            # skip references that aren't child relation
            if ref.ref_type != model_pb2.CHILD_OF:
                continue

            parent_key = compose_key(ref.trace_id, ref.span_id)
            if parent_key in children:
                children[parent_key].append(key)
            else:
                children[parent_key] = [key]

    return children

def span_interval(spans: list[model_pb2.Span]) -> dict[str, tuple[datetime, datetime]]:
    intervals = {}

    for span in spans:
        key = compose_key(span.trace_id, span.span_id)

        start = span.start_time.ToDatetime()
        duration = span.duration.ToTimedelta()
        end = start + duration
        intervals[key] = (start, end)

    return intervals

def filter_span(spans: list[model_pb2.Span]) -> dict[str, str]:
    names = {}

    for span in spans:
        key = compose_key(span.trace_id, span.span_id)

        service = span.process.service_name
        operation = span.operation_name
        name = service + '-' + operation

        for kv in span.tags:
            if kv.key == 'span.kind' and kv.v_str == 'server':
                names[key] = name
                break

    return names

def calc_ratio(children: dict[str, list[str]], intervals: dict[str, tuple[datetime, datetime]], key: str) -> float:
    # have no children
    if key not in children:
        return 0.0

    # parent's duration
    s, e = intervals[key]
    duration = (e - s) / MS

    children_interval = [intervals[c] for c in children[key]]
    children_interval.sort(key=lambda x: x[0]) # sort by start time

    # calculate total duration of children spans, overlap is counted once
    end = datetime.min # end time of previous interval
    subduration = 0.0
    for s, e in children_interval:
        if s < end:
            # have overlap with previous interval
            if e > end:
                # not covered by previous interval
                subduration += (e - end) / MS
                end = e
        else:
            # no overlap
            subduration += (e - s) / MS
            end = e

    return subduration / duration

filename = 'spans.json'
with open(filename, 'r') as f:
    spans = strings_to_protos(json.load(f))
children = span_children(spans)
intervals = span_interval(spans)
names = filter_span(spans)

latency = {}
ratio = {}

for key, name in names.items():
    s, e = intervals[key]
    duration = (e - s) / MS
    if name in latency:
        latency[name].append(duration)
    else:
        latency[name] = [duration]

    r = calc_ratio(children, intervals, key)
    if name in ratio:
        ratio[name].append(r)
    else:
        ratio[name] = [r]

for k, v in latency.items():
    latency[k] = sum(v) / len(v)

for k, v in ratio.items():
    ratio[k] = sum(v) / len(v)

latency = pd.Series(latency, name='latency')
ratio = pd.Series(ratio, name='ratio')
df = pd.concat([latency, ratio], axis=1)
df.reset_index(inplace=True)
df = df.rename(columns = {'index':'function'})
df.to_csv('result.csv', index=False)
