from gevent import monkey
monkey.patch_all()

from gevent.pywsgi import WSGIServer
import os
from flask import Flask, request
import requests
import json
import gevent
import contextvars


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("MapReduce")
with open(CONFIG_PATH, 'r') as f:
    config = json.load(f)


@app.route('/', methods=['POST'])
def main():
    args = request.get_json()

    # get file names from file-service
    req = {
        'bucket': args['bucket'],
        'prefix': args['prefix']
    }
    resp = requests.post(config['service']['file'], json=req)
    filenames = resp.json()['filenames']

    # get average speed from mapper-service
    tasks = []
    for fn in filenames:
        ctx = contextvars.copy_context()
        tasks.append(gevent.spawn(ctx.run, call_mapper, args['bucket'], fn))
    gevent.joinall(tasks)
    speed = [t.value for t in tasks]

    # get reduced result from reducer-service
    req = {
        'data': speed
    }
    resp = requests.post(config['service']['reducer'], json=req)
    result = resp.json()['result']

    return {
        'result': result
    }


def call_mapper(bucket: str, file: str) -> list[float]:
    req = {
        'bucket': bucket,
        'file': file
    }
    resp = requests.post(config['service']['mapper'], json=req)
    return resp.json()['result']


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
