import json
import pymongo
import unittest
import grpc
import bson
import redis
from typing import Any, Dict, List, Type, TypeVar
from bson import json_util
from pymongo.collection import Collection
from google.protobuf.json_format import Parse, MessageToDict
from google.protobuf.message import Message
from pymemcache.client.base import Client


class TestSocialNetwork(unittest.TestCase):
    # mongodb collection
    post_db: Collection
    timeline_db: Collection
    user_db: Collection
    socialgraph_db: Collection
    url_db: Collection
    # redis
    timeline_redis: redis.Redis
    # config
    secret: str
    rest: Dict[str, str]

    def setUp(self, service: str, StubType: Type) -> None:
        config = get_config()

        # connect mongodb
        self.post_db = mongo_connect_and_clean(config['mongodb'], 'post')
        self.timeline_db = mongo_connect_and_clean(
            config['mongodb'], 'user-timeline')
        self.user_db = mongo_connect_and_clean(config['mongodb'], 'user')
        self.socialgraph_db = mongo_connect_and_clean(
            config['mongodb'], 'social-graph')
        self.url_db = mongo_connect_and_clean(config['mongodb'], 'url-shorten')

        # clean memcached
        for addr in config['memcached'].values():
            memcached_clean(addr)

        # clean redis
        for addr in config['redis'].values():
            redis_clean(addr)
        # hometimelint uses redis as main database
        self.timeline_redis = redis.Redis.from_url(
            config['redis']['hometimeline'], decode_responses=True)

        # connect grpc service
        chan = grpc.insecure_channel(config['service'][service])
        self.stub = StubType(chan)

        self.secret = config['secret']
        self.rest = config['service-rest']


# path for config file
config_file = '../config/dev.json'


def get_config() -> Dict[str, Any]:
    with open(config_file, 'r') as f:
        return json.load(f)


def mongo_connect_and_clean(config: Dict[str, str], collection: str) -> Collection:
    client = pymongo.MongoClient(config['url'])
    col = client[config['database']][collection]
    col.delete_many({})
    return col


def memcached_clean(addr: str) -> None:
    client = Client(addr)
    client.flush_all()
    client.close()


def redis_clean(url: str) -> None:
    client = redis.Redis.from_url(url)
    client.flushdb()
    client.close()


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
