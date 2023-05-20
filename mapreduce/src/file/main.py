from gevent import monkey
monkey.patch_all()

import os
from gevent.pywsgi import WSGIServer
from flask import Flask, request
from minio import Minio
import json


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("File")
with open(CONFIG_PATH, 'r') as f:
    config = json.load(f)


@app.route('/', methods=['POST'])
def main():
    args = request.get_json()

    client = Minio(
        config['minio']['url'],
        access_key=config['minio']['access_key'],
        secret_key=config['minio']['secret_key'],
        secure=config['minio']['secure']
    )

    objs = client.list_objects(args['bucket'], prefix=args['prefix'])
    filenames = [o.object_name for o in objs]

    return {
        'filenames': filenames
    }


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
