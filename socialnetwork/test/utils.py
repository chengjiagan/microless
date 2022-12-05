import json
import pymongo
import unittest
import grpc
import redis
import os
from typing import Any, Dict, List, Type, TypeVar
from bson import json_util
from pymongo.collection import Collection
from google.protobuf.json_format import Parse, MessageToDict
from google.protobuf.message import Message
from pymemcache.client.base import Client


class TestSocialNetwork(unittest.TestCase):
    config_file = os.environ.get('TEST_CONFIG') or '../config/dev.json'
    # mongodb collection
    post_db: Collection
    usertimeline_db: Collection
    hometimeline_db: Collection
    user_db: Collection
    socialgraph_db: Collection
    url_db: Collection
    # config
    secret: str
    gateway: str

    def setUp(self, service: str, StubType: Type) -> None:
        config = self.get_config()

        # connect mongodb
        self.post_db = mongo_connect(config['mongodb'], 'post')
        self.usertimeline_db = mongo_connect(config['mongodb'], 'user-timeline')
        self.hometimeline_db = mongo_connect(config['mongodb'], 'home-timeline')
        self.user_db = mongo_connect(config['mongodb'], 'user')
        self.socialgraph_db = mongo_connect(config['mongodb'], 'social-graph')
        self.url_db = mongo_connect(config['mongodb'], 'url-shorten')

        # connect grpc service
        chan = grpc.insecure_channel(config['service'][service])
        self.stub = StubType(chan)

        self.secret = config['secret']
        self.gateway = config['gateway']

        self.clean()

    def clean(self):
        config = self.get_config()
        # clean mongodb
        self.post_db.delete_many({})
        self.usertimeline_db.delete_many({})
        self.hometimeline_db.delete_many({})
        self.user_db.delete_many({})
        self.socialgraph_db.delete_many({})
        self.url_db.delete_many({})
        # clean memcached
        for addr in config['memcached'].values():
            memcached_clean(addr)
        # clean redis
        for addr in config['redis'].values():
            redis_clean(addr)

    def get_config(self) -> Dict[str, Any]:
        with open(self.config_file, 'r') as f:
            return json.load(f)


def mongo_connect(config: Dict[str, str], collection: str) -> Collection:
    client = pymongo.MongoClient(config['url'])
    col = client[config['database']][collection]
    return col


def memcached_clean(addr: str) -> None:
    try:
        client = Client(addr)
        client.flush_all()
        client.close()
    except:
        pass


def redis_clean(url: str) -> None:
    try:
        client = redis.Redis.from_url(url)
        client.flushdb()
        client.close()
    except:
        pass


def get_bson(filename: str) -> Any:
    with open(filename, 'r') as f:
        return json_util.loads(f.read())


def get_json(filename: str) -> Any:
    with open(filename, 'r') as f:
        return json.load(f)


T = TypeVar('T')


def get_proto(filename: str, ProtoType: Type[T]) -> T:
    proto = ProtoType()
    with open(filename, 'r') as f:
        return Parse(f.read(), proto)


def get_text(filename: str) -> str:
    with open(filename, 'r') as f:
        return f.read()
