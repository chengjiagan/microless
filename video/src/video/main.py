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

    # generate preview image
    ctx = contextvars.copy_context()
    t_preview = gevent.spawn(ctx.run, call_preview,
                             args['bucket'], args['input'], args['outdir'])

    if args['segment_time'] <= 0:
        video = simple_transcode(
            args['bucket'], args['input'], args['type'], args['outdir'])
    else:
        video = split_transcode(
            args['bucket'], args['input'], args['type'], args['segment_time'], args['outdir'])

    t_preview.join()
    return {
        'preview': t_preview.value,
        'transcoded': video
    }


def call_preview(bucket: str, file: str, outdir: str) -> str:
    _, fullname = os.path.split(file)
    filename, _ = os.path.splitext(fullname)
    output = os.path.join(outdir, f'{filename}_preview.gif')
    req = {
        'bucket': bucket,
        'input': file,
        'output': output
    }
    resp = requests.post(config['service']['preview'], json=req)
    return resp.json()['result']


def simple_transcode(bucket: str, file: str, target_type: str, outdir: str) -> str:
    base, _ = os.path.splitext(file)
    output = os.path.join(outdir, f'{base}.{target_type}')
    req = {
        'bucket': bucket,
        'input': file,
        'type': target_type,
        'output': output,
        'delete': False
    }
    resp = requests.post(config['service']['transcode'], json=req)
    return resp.json()['result']


def split_transcode(bucket: str, file: str, target_type: str, segment_time: float, outdir: str) -> str:
    # call split function to split the video
    req = {
        'bucket': bucket,
        'input': file,
        'segment_time': segment_time,
        'outdir': outdir
    }
    resp = requests.post(config['service']['split'], json=req)
    segments = resp.json()['result']

    # call transcode function to transcode segments
    tasks = []
    for seg in segments:
        ctx = contextvars.copy_context()
        t = gevent.spawn(ctx.run, call_transcode, bucket,
                         seg, target_type, outdir)
        tasks.append(t)
    gevent.wait(tasks)
    trans_segment = [t.value for t in tasks]

    # call merge function to merge transcoded segments
    base, _ = os.path.splitext(file)
    output = os.path.join(outdir, f'{base}.{target_type}')
    req = {
        'bucket': bucket,
        'segments': trans_segment,
        'output': output
    }
    resp = requests.post(config['service']['merge'], json=req)
    video = resp.json()['result']

    return video


def call_transcode(bucket: str, segment: str, target_type: str, outdir: str) -> str:
    base, _ = os.path.splitext(segment)
    output = f'{base}_transcoded.{target_type}'
    req = {
        'bucket': bucket,
        'input': segment,
        'type': target_type,
        'output': output,
        'delete': True
    }
    resp = requests.post(config['service']['transcode'], json=req)
    return resp.json()['result']


if __name__ == '__main__':
    server = WSGIServer(config['rest'], app)
    server.serve_forever()
