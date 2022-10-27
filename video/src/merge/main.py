import contextvars
from gevent import monkey
import gevent
monkey.patch_all()

import tempfile
from minio import Minio
from gevent.pywsgi import WSGIServer
import os
from flask import Flask, request
import json


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("Merge")
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

    _, ext = os.path.splitext(args['segments'][0])

    with tempfile.TemporaryDirectory() as tmpdir:
        # download segments from minio
        segments_tmp = download_segment(client, args['bucket'], args['segments'], tmpdir, ext)

        # merge segments
        merge_tmp = os.path.join(tmpdir, f'merge{ext}')
        merge_segment(tmpdir, segments_tmp, merge_tmp)

        # upload merge video to minio
        client.fput_object(args['bucket'], args['output'], merge_tmp)

    # remove segments from minio
    for seg in args['segments']:
        client.remove_object(args['bucket'], seg)

    return {
        'result': args['output']
    }


def download_segment(client: Minio, bucket: str, segments: list[str], tmpdir: str, ext: str) -> list[str]:
    tasks = []
    segments_tmp = []
    for i, seg_minio in enumerate(segments):
        seg_tmp = os.path.join(tmpdir, f'seg_{i}{ext}')
        segments_tmp.append(seg_tmp)

        ctx = contextvars.copy_context()
        t = gevent.spawn(ctx.run, client.fget_object, bucket, seg_minio, seg_tmp)
        tasks.append(t)
    gevent.wait(tasks)

    return segments_tmp

def merge_segment(tmpdir: str, segments: list[str], output: str) -> None:
    # generate list file
    list_path = os.path.join(tmpdir, 'list.txt')
    with open(list_path, 'w') as f:
        for seg in segments:
            f.write(f"file '{seg}'\n")
    
    # call ffmpeg to merge segments
    cmd = f'ffmpeg -f concat -safe 0 -i {list_path} -c copy -fflags +genpts {output}'
    os.system(cmd)


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
