from gevent import monkey
monkey.patch_all()

import tempfile
from minio import Minio
from gevent.pywsgi import WSGIServer
import os
from flask import Flask, request
import json


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("Transcode")
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

    _, origin_ext = os.path.splitext(args['input'])
    new_ext = '.' + args['type']

    # a temparary directory for getting data
    with tempfile.TemporaryDirectory() as tmpdir:
        # get input video from minio
        input_tmp = os.path.join(tmpdir, f'input{origin_ext}')
        client.fget_object(args['bucket'], args['input'], input_tmp)

        # transcode video
        transcoded_tmp = os.path.join(tmpdir, f'transcoded{new_ext}')
        cmd = f'ffmpeg -y -i {input_tmp} -preset superfast {transcoded_tmp}'
        os.system(cmd)

        # upload transcoded video
        client.fput_object(args['bucket'], args['output'], transcoded_tmp)

    # remove input video from minio
    if args['delete']:
        client.remove_object(args['bucket'], args['input'])

    return {
        'result': args['output']
    }


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
