from gevent import monkey
monkey.patch_all()

from gevent.pywsgi import WSGIServer
from datetime import timedelta
import pandas as pd
import json
from minio import Minio
from flask import Flask, request
import tempfile
import os


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("Mapper")
with open(CONFIG_PATH, 'r') as f:
    config = json.load(f)
hour = timedelta(hours=1)


@app.route('/', methods=['POST'])
def main():
    args = request.get_json()

    client = Minio(
        config['minio']['url'],
        access_key=config['minio']['access_key'],
        secret_key=config['minio']['secret_key'],
        secure=config['minio']['secure']
    )

    # create a temparary directory for getting data file
    with tempfile.TemporaryDirectory() as tmpdir:
        filename = os.path.basename(args['file'])
        tmpfile = os.path.join(tmpdir, filename)
        client.fget_object(args['bucket'], args['file'], tmpfile)

        df = pd.read_parquet(tmpfile)
        time = (df['tpep_dropoff_datetime'] -
                df['tpep_pickup_datetime']) / hour
        distance = df['trip_distance']
        # filter out data with non-positive time and non-positive distance
        select = (time > 0) & (distance > 0)

        result = pd.DataFrame()
        result['speed'] = distance[select] / time[select]
        result['dayofweek'] = df['tpep_pickup_datetime'][select].dt.dayofweek.astype(
            int)
        # group by day of week and compute average speed
        result = result.groupby('dayofweek').mean()

    return {
        'result': result['speed'].tolist()
    }


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
