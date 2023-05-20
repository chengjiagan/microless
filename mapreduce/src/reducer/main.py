from gevent import monkey
monkey.patch_all()

import json
from flask import Flask, request
import os
import numpy as np
from gevent.pywsgi import WSGIServer


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("Reducer")
with open(CONFIG_PATH, 'r') as f:
    config = json.load(f)


@app.route('/', methods=['POST'])
def main():
    args = request.get_json()

    data = np.asarray(args['data'])
    result = data.mean(axis=0).tolist()

    return {
        'result': result
    }


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
