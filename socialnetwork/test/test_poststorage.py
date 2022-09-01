#! /usr/bin/env python3
import unittest
import utils
from bson import ObjectId
from proto.data_pb2 import Post
from proto.poststorage_pb2_grpc import PostStorageServiceStub
from proto.poststorage_pb2 import ReadPostsRespond, StorePostRequest, ReadPostsRequest


class TestPostStorage(utils.TestSocialNetwork):
    stub: PostStorageServiceStub

    def setUp(self) -> None:
        super().setUp('poststorage', PostStorageServiceStub)

    def test_store_post(self) -> None:
        req = StorePostRequest(post=utils.get_proto(
            'json/test_store_post_proto.json', Post))
        resp = self.stub.StorePost(req)

        expect = utils.get_bson('json/test_store_post_bson.json')
        expect['_id'] = ObjectId(resp.post_id)

        actual = self.post_db.find_one()
        self.assertEqual(expect, actual)

    def test_read_posts(self) -> None:
        docs = utils.get_bson('json/test_read_posts_bson.json')
        self.post_db.insert_many(docs)

        req = ReadPostsRequest(
            post_ids=["630eebda0daff4bcd9a36c3e", "630eecc90daff4bcd9a36c40"])
        resp = self.stub.ReadPosts(req)

        expect = utils.get_proto(
            'json/test_read_posts_proto.json', ReadPostsRespond)
        self.assertEqual(expect, resp)


if __name__ == '__main__':
    unittest.main()
