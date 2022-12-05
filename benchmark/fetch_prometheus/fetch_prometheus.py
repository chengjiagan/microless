#! /usr/bin/env python

import argparse
import requests
import pandas as pd
from os import path
from datetime import datetime

query = {
    'cpu': 'sum by(pod) (irate(container_cpu_usage_seconds_total{{namespace="{namespace}"}}[1m]))',
    'memory': 'sum by(pod) (container_memory_usage_bytes{{namespace="{namespace}"}})'
}

# execute query in prometheus
def fetch(url: str, query: str, start: datetime, end: datetime) -> list[dict]:
    params = {
        'query': query,
        'start': int(start.timestamp()),
        'end': int(end.timestamp()),
        'step': '30s'
    }
    resp = requests.get(url, params=params)
    data = resp.json()
    return data['data']['result']

# get start and end timestamp from load record
def get_time(record: str) -> tuple[datetime, datetime]:
    df = pd.read_csv(record)
    # timestamp in record are in miliseconds, utc in prometheus
    start = datetime.fromtimestamp(df['start'].min() / 1000)
    end = datetime.fromtimestamp(df['end'].max() / 1000)
    return (start, end)

def to_dataframe(data: list[dict]) -> pd.DataFrame:
    df = pd.DataFrame()
    for t in data:
        labels = [f'{k}:{v}' for k, v in t['metric'].items()]
        name = ' '.join(labels)
        df[name] = pd.Series(dict(t['values']))
    df = df.astype(float)
    return df

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-a', '--address', type=str, help='address to the jaeger server', required=True)
    parser.add_argument('-i', '--input', type=str, help='path to the load record file', required=True)
    parser.add_argument('-o', '--output', type=str, help='path to the output file')
    parser.add_argument('-n', '--namespace', type=str, help='namespace to query', required=True)
    args = parser.parse_args()

    # generate an ouput path from 
    if args.output is None:
        fn, _ = path.splitext(path.basename(args.input))
        args.output = fn.replace('load', 'prometheus') + '.xlsx'

    url = f'http://{args.address}/api/v1/query_range'
    start, end = get_time(args.input)
    with pd.ExcelWriter(args.output) as ew:
        for n, q in query.items():
            data = fetch(url, q.format(namespace=args.namespace), start, end)
            df = to_dataframe(data)
            df.T.to_excel(ew, n, index_label='labels')