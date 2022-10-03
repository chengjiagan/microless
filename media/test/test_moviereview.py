import unittest

from bson import ObjectId
import utils
from proto.moviereview_pb2 import CreateMovieRequest, ReadMovieReviewsRequest, ReadMovieReviewsRespond, UploadMovieReviewRequest
from proto.moviereview_pb2_grpc import MovieReviewServiceStub
from google.protobuf.timestamp_pb2 import Timestamp


class TestMovieReview(utils.TestMedia):
    stub: MovieReviewServiceStub

    def setUp(self) -> None:
        super().setUp('moviereview', MovieReviewServiceStub)

    def test_upload_movie_review(self) -> None:
        movie_id = '000000000000000000000000'
        self.db['movie-review'].insert_one({
            'movie_id': ObjectId(movie_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('633942b584c48ef10d77f058')]
        })

        dt = ObjectId('6339457ac051e3a5450fa284').generation_time
        timestamp = Timestamp()
        timestamp.FromDatetime(dt)
        req = UploadMovieReviewRequest(
            movie_id=movie_id, review_id='6339457ac051e3a5450fa284', timestamp=timestamp)
        self.stub.UploadMovieReview(req)

        expect = {
            'movie_id': ObjectId(movie_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('6339457ac051e3a5450fa284'), ObjectId('633942b584c48ef10d77f058')]
        }
        actual = self.db['movie-review'].find_one()
        del actual['_id']
        self.assertEqual(expect, actual)

    def test_read_movie_reviews(self) -> None:
        movie_id = '000000000000000000000000'
        self.db['movie-review'].insert_one({
            'movie_id': ObjectId(movie_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('6339457ac051e3a5450fa284'), ObjectId('633942b584c48ef10d77f058')]
        })
        self.db['review'].insert_many(utils.get_bson(
            'data/test_read_movie_reviews_bson.json'))

        req = ReadMovieReviewsRequest(movie_id=movie_id, start=1, stop=3)
        resp = self.stub.ReadMovieReviews(req)

        expect = utils.get_proto(
            'data/test_read_movie_reviews_proto.json', ReadMovieReviewsRespond)
        actual = resp
        self.assertEqual(expect, actual)

    def test_create_movie(self) -> None:
        movie_id = '000000000000000000000000'
        req = CreateMovieRequest(movie_id=movie_id)
        self.stub.CreateMovie(req)

        expect = {
            "movie_id": ObjectId(movie_id),
            "review_ids": []
        }
        actual = self.db['movie-review'].find_one()
        del actual['_id']
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
