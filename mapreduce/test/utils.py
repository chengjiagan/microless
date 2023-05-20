import json
import os
from typing import Any
import unittest
import minio


BUCKET = 'mapreduce-test'


class TestMapReduce(unittest.TestCase):
    config_file = os.environ.get('TEST_CONFIG') or '../config/dev.json'
    storage: minio.Minio
    # config
    addr: str

    def setUp(self, service: str) -> None:
        config = self.get_config()

        # get the address of service
        self.addr = config['service'][service]

        # connect to minio
        self.storage = minio.Minio(
            config['minio']['url'],
            access_key=config['minio']['access_key'],
            secret_key=config['minio']['secret_key'],
            secure=config['minio']['secure']
        )

        # do clean up
        self.clean()

    def clean(self):
        if not self.storage.bucket_exists(BUCKET):
            self.storage.make_bucket(BUCKET)
        for o in self.storage.list_objects(BUCKET, recursive=True):
            self.storage.remove_object(BUCKET, o.object_name)

    def get_config(self) -> dict[str, Any]:
        with open(self.config_file, 'r') as f:
            return json.load(f)

def get_json(file: str) -> dict[str, Any]:
    with open(file, 'r') as f:
        return json.load(f)
