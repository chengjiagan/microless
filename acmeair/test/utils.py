import json
import pymongo
import unittest
import grpc
import os
import typing
from typing import Any, Literal, Mapping, Optional, Tuple, Type, TypeVar
from bson import ObjectId, json_util
from pymongo.collection import Collection
from google.protobuf.json_format import Parse
from pymemcache.client.base import Client

DbNameT = Literal['bookings', 'customer', 'flights']
T = TypeVar('T')
DB_NAME: Tuple[str] = typing.get_args(DbNameT)


class TestAcmeair(unittest.TestCase):
    config_file = os.environ.get('TEST_CONFIG') or '../config/dev.json'
    # mongodb collection
    db: Mapping[DbNameT, Collection] = {}
    # config
    secret: str
    gateway: str

    def setUp(self, service: str, StubType: Type) -> None:
        config = self.get_config()

        # connect mongodb
        for name in DB_NAME:
            self.db[name] = mongo_connect(config['mongodb'], name)

        # connect grpc service
        chan = grpc.insecure_channel(config['service'][service])
        self.stub = StubType(chan)

        self.gateway = config['gateway']

        self.clean()

    def clean(self):
        config = self.get_config()
        # clean mongodb
        for name in DB_NAME:
            self.db[name].delete_many({})
        # clean memcached
        for addr in config['memcached'].values():
            memcached_clean(addr)

    def get_config(self) -> Mapping[str, Any]:
        with open(self.config_file, 'r') as f:
            return json.load(f)


def mongo_connect(config: Mapping[str, str], collection: str) -> Collection:
    client = pymongo.MongoClient(config['url'])
    col = client[config['database']][collection]
    return col


def memcached_clean(addr: str) -> None:
    client = Client(addr)
    client.flush_all()
    client.close()


def get_bson(filename: str, oid: Optional[ObjectId] = None) -> Any:
    with open(filename, 'r') as f:
        obj = json_util.loads(f.read())
    if oid is not None:
        obj['_id'] = oid
    return obj


def get_json(filename: str) -> Any:
    with open(filename, 'r') as f:
        return json.load(f)


def get_proto(filename: str, ProtoType: Type[T]) -> T:
    proto = ProtoType()
    with open(filename, 'r') as f:
        return Parse(f.read(), proto)


def get_text(filename: str) -> str:
    with open(filename, 'r') as f:
        return f.read()
