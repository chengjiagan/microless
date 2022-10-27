from gevent import monkey
monkey.patch_all()

import tempfile
from minio import Minio
from gevent.pywsgi import WSGIServer
import os
from flask import Flask, request
import json
import gevent
import contextvars
import glob


CONFIG_PATH = os.getenv('SERVICE_CONFIG') or ''

app = Flask("Split")
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

    _, fullname = os.path.split(args['input'])
    filename, ext = os.path.splitext(fullname)

    # a temparary directory for getting data
    with tempfile.TemporaryDirectory() as tmpdir:
        # get input video from minio
        input_tmp = os.path.join(tmpdir, f'input{ext}')
        client.fget_object(args['bucket'], args['input'], input_tmp)

        # split video into segments
        time = args['segment_time']
        segments_tmp = os.path.join(tmpdir, f'seg_%d{ext}')
        cmd = f'ffmpeg -i {input_tmp} -c copy -f segment -segment_time {time} -reset_timestamps 1 {segments_tmp}'
        os.system(cmd)

        # get all segments
        segments = glob.glob('seg_*', root_dir=tmpdir)

        # upload segments
        tasks = []
        segments_minio = []
        for seg in segments:
            seg_tmp = os.path.join(tmpdir, seg)
            seg_minio = os.path.join(args['outdir'], f'{filename}_{seg}')
            segments_minio.append(seg_minio)

            ctx = contextvars.copy_context()
            t = gevent.spawn(ctx.run, client.fput_object, args['bucket'], seg_minio, seg_tmp)
            tasks.append(t)
        gevent.wait(tasks)

    return {
        'result': segments_minio
    }


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
