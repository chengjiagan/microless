import unittest

from bson import ObjectId
import requests
import utils
from proto.page_pb2 import ReadPageRequest, ReadPageRespond
from proto.page_pb2_grpc import PageServiceStub


class TestPage(utils.TestMedia):
    stub: PageServiceStub

    def setUp(self) -> None:
        super().setUp('page', PageServiceStub)

    def test_read_page(self) -> None:
        docs = utils.get_bson('data/test_read_page_bson.json')
        self.db['review'].insert_many(docs['review'])
        self.db['movie-review'].insert_one(docs['movie-review'])
        self.db['movie-info'].insert_one(docs['movie-info'])
        self.db['cast-info'].insert_many(docs['cast-info'])
        self.db['plot'].insert_one(docs['plot'])

        movie_id = '000000000000000000000000'
        req = ReadPageRequest(movie_id=movie_id, review_start=1, review_stop=3)
        resp = self.stub.ReadPage(req)

        expect = utils.get_proto('data/test_read_page_proto.json', ReadPageRespond)
        actual = resp
        self.assertEqual(expect, actual)

    def test_read_page_rest(self) -> None:
        docs = utils.get_bson('data/test_read_page_bson.json')
        self.db['review'].insert_many(docs['review'])
        self.db['movie-review'].insert_one(docs['movie-review'])
        self.db['movie-info'].insert_one(docs['movie-info'])
        self.db['cast-info'].insert_many(docs['cast-info'])
        self.db['plot'].insert_one(docs['plot'])

        movie_id = '000000000000000000000000'
        url = 'http://' + self.gateway + '/api/v1/page/' + movie_id
        req = {
            'review_start': 1,
            'review_stop': 3
        }
        resp = requests.get(url, params=req)

        expect = utils.get_json('data/test_read_page_json.json')
        actual = resp.json()
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
