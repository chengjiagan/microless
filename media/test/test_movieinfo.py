import unittest

from bson import ObjectId
import utils
from proto.movieinfo_pb2 import ReadMovieInfoRequest, UpdateRatingRequest, WriteMovieInfoRequest
from proto.movieinfo_pb2_grpc import MovieInfoServiceStub
from proto.data_pb2 import MovieInfo


class TestMovieInfo(utils.TestMedia):
    stub: MovieInfoServiceStub

    def setUp(self) -> None:
        super().setUp('movieinfo', MovieInfoServiceStub)

    def test_write_movie_info(self) -> None:
        req = utils.get_proto(
            'data/test_write_movie_info_proto.json', WriteMovieInfoRequest)
        resp = self.stub.WriteMovieInfo(req)
        oid = ObjectId(resp.movie_id)

        expect = utils.get_bson('data/test_write_movie_info_bson.json', oid)
        actual = self.db['movie-info'].find_one()
        self.assertEqual(expect, actual)

        # WriteMovieInfo will alse create a new record in movie-review database
        expect = {
            'movie_id': oid,
            'review_ids': []
        }
        actual = self.db['movie-review'].find_one()
        del actual['_id']
        self.assertEqual(expect, actual)

    def test_read_movie_info(self) -> None:
        info = utils.get_bson('data/test_read_movie_info_bson.json')
        self.db['movie-info'].insert_one(info)

        req = ReadMovieInfoRequest(movie_id='000000000000000000000000')
        resp = self.stub.ReadMovieInfo(req)

        expect = utils.get_proto(
            'data/test_read_movie_info_proto.json', MovieInfo)
        actual = resp
        self.assertEqual(expect, actual)

    def test_update_rating(self) -> None:
        info = utils.get_bson('data/test_update_rating_original_bson.json')
        self.db['movie-info'].insert_one(info)

        req = UpdateRatingRequest(movie_id='000000000000000000000000',
                                  sum_uncommitted_rating=20, num_uncommitted_rating=5)
        self.stub.UpdateRating(req)

        expect = utils.get_bson('data/test_update_rating_expect_bson.json')
        actual = self.db['movie-info'].find_one()
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
