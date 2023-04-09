#! /usr/bin/env python

import pandas as pd
import sys
import glob
import re

path = sys.argv[1]

result = {}
files = glob.glob(f'{path}/load_*.csv')
files = sorted(files)
for f in files:
    t = re.search(r'load_(.*)\.csv', f).group(1)
    df = pd.read_csv(f)
    start = df['start'].min()
    end = df['end'].max()
    tp = len(df) / (end - start) * 1000 # query per second

    wdf = df[df['type'] == 'write']
    wtp = len(wdf) / (end - start) * 1000 # query per second
    wlat = wdf['end'] - wdf['start']
    wlat_avg = wlat.mean()
    wlat_med = wlat.quantile(0.5)
    wlat_95 = wlat.quantile(0.95)
    wlat_99 = wlat.quantile(0.99)
    wlat_var = wlat.std() / wlat.mean()
    
    rdf = df[df['type'] == 'read']
    rtp = len(rdf) / (end - start) * 1000 # query per second
    rlat = rdf['end'] - rdf['start']
    rlat_avg = rlat.mean()
    rlat_med = rlat.quantile(0.5)
    rlat_95 = rlat.quantile(0.95)
    rlat_99 = rlat.quantile(0.99)
    rlat_var = rlat.std() / rlat.mean()

    result[t] = [tp, wtp, wlat_avg, wlat_med, wlat_95, wlat_99, wlat_var, rtp, rlat_avg, rlat_med, rlat_95, rlat_99, rlat_var]

df = pd.DataFrame(result, index=['total throughput', 'write throughput', 'write average latency', 'write median latency', 'write p95 latency', 'write p99 latency', 'write latency stdvar', 'read throughput', 'read average latency', 'read median latency', 'read p95 latency', 'read p99 latency', 'read latency stdvar'])
df.to_csv(f'{path}/load.csv')
