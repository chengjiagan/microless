#! /usr/bin/env python

import sys
import pandas as pd
from datetime import datetime

file = sys.argv[1]
df = pd.read_csv(file)

start = df['start'].min()
start = datetime.fromtimestamp(start / 1000)
end = df['end'].max()
end = datetime.fromtimestamp(end / 1000)

print('start: ', start)
print('end: ', end)
