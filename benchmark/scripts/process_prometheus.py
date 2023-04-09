#! /usr/bin/env python

import pandas as pd

threads = [1, 2, 4, 8, 16, 32]
timestamp = '12012049'
cpu = pd.DataFrame()
for t in threads:
    df = pd.read_excel(f'../data/{timestamp}/prometheus_t{t}.xlsx', sheet_name='cpu', index_col='labels')
    cpu[t] = df.max(axis=1)

memory = pd.DataFrame()
for t in threads:
    df = pd.read_excel(f'../data/{timestamp}/prometheus_t{t}.xlsx', sheet_name='memory', index_col='labels')
    memory[t] = df.max(axis=1) / 1024**2 # in MiB

with pd.ExcelWriter('prometheus.xlsx') as ew:
    cpu.to_excel(ew, 'cpu', index_label='labels')
    memory.to_excel(ew, 'memory', index_label='labels')
