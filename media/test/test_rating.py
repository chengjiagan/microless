import unittest

from bson import ObjectId
import requests
import utils
from proto.rating_pb2 import UploadRatingRequest
from proto.rating_pb2_grpc import RatingServiceStub


class TestRating(utils.TestMedia):
    stub: RatingServiceStub

    def setUp(self) -> None:
        super().setUp('rating', RatingServiceStub)

    def test_upload_rating(self) -> None:
        self.db['movie-info'].insert_one(utils.get_bson('data/test_upload_rating_original_bson.json'))

        movie_id = '000000000000000000000000'
        req = UploadRatingRequest(movie_id=movie_id, rating=3)
        self.stub.UploadRating(req)

        expect = utils.get_bson('data/test_upload_rating_expect_bson.json')
        actual = self.db['movie-info'].find_one()
        self.assertEqual(expect, actual)

    def test_upload_rating_rest(self) -> None:
        self.db['movie-info'].insert_one(utils.get_bson('data/test_upload_rating_original_bson.json'))

        movie_id = '000000000000000000000000'
        req = {
            'movie_id': movie_id,
            'rating': 3
        }
        url = 'http://' + self.gateway + '/api/v1/rating'
        requests.post(url, json=req)

        expect = utils.get_bson('data/test_upload_rating_expect_bson.json')
        actual = self.db['movie-info'].find_one()
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
