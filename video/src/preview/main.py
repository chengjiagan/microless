from gevent import monkey
monkey.patch_all()

import json
from flask import Flask, request
import os
from gevent.pywsgi import WSGIServer
from minio import Minio
import tempfile


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("Preview")
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

    # should have input = path/file.ext
    _, ext = os.path.splitext(args['input'])

    # a temparary directory for getting data
    with tempfile.TemporaryDirectory() as tmpdir:
        # get input video from minio
        input_tmp = os.path.join(tmpdir, f'input{ext}')
        client.fget_object(args['bucket'], args['input'], input_tmp)

        # generate preview
        preview = 'preview.gif'
        preview_tmp = os.path.join(tmpdir, preview)
        cmd = f'ffmpeg -t 3 -ss 00:00:02 -i {input_tmp} {preview_tmp}'
        os.system(cmd)

        # upload preview
        client.fput_object(args['bucket'], args['output'], preview_tmp)

    return {
        'result': args['output']
    }


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
