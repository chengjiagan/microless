import unittest

from bson import ObjectId
import requests
import utils
from proto.composereview_pb2 import ComposeReviewRequest
from proto.composereview_pb2_grpc import ComposeReviewStub


class TestComposeReview(utils.TestMedia):
    stub: ComposeReviewStub

    def setUp(self) -> None:
        super().setUp('composereview', ComposeReviewStub)

    def test_compose_review(self) -> None:
        movie_id = '000000000000000000000000'
        user_id = '000000000000000000000000'
        docs = utils.get_bson('data/test_compose_review_bson.json')
        self.db['movie-review'].insert_one(docs['movie-review'])
        self.db['user-review'].insert_one(docs['user-review'])
        self.db['movie-info'].insert_one(docs['movie-info'])

        text = 'some random review text'
        rating = 3
        req = ComposeReviewRequest(
            movie_id=movie_id, user_id=user_id, text=text, rating=rating)
        self.stub.ComposeReview(req)

        # check review-storage
        expect = {
            'movie_id': ObjectId(movie_id),
            'user_id': ObjectId(user_id),
            'text': text,
            'rating': rating
        }
        actual = self.db['review'].find_one()
        review_id = actual['_id']
        del actual['_id']
        self.assertEqual(expect, actual)

        # check user-review
        expect = docs['user-review']
        expect['review_ids'] = [review_id]
        actual = self.db['user-review'].find_one()
        self.assertEqual(expect, actual)

        # check movie-review
        expect = docs['movie-review']
        expect['review_ids'] = [review_id]
        actual = self.db['movie-review'].find_one()
        self.assertEqual(expect, actual)

        # check rating
        expect = docs['movie-info']
        expect['avg_rating'] = 2.5
        expect['num_rating'] = 6
        actual = self.db['movie-info'].find_one()
        self.assertEqual(expect, actual)

    def test_compose_review_rest(self) -> None:
        movie_id = '000000000000000000000000'
        user_id = '000000000000000000000000'
        docs = utils.get_bson('data/test_compose_review_bson.json')
        self.db['movie-review'].insert_one(docs['movie-review'])
        self.db['user-review'].insert_one(docs['user-review'])
        self.db['movie-info'].insert_one(docs['movie-info'])

        text = 'some random review text'
        rating = 3
        req = {
            'movie_id': movie_id,
            'user_id': user_id,
            'text': text,
            'rating': rating
        }
        url = 'http://' + self.gateway + '/api/v1/composereview'
        requests.post(url, json=req)

        # check review-storage
        expect = {
            'movie_id': ObjectId(movie_id),
            'user_id': ObjectId(user_id),
            'text': text,
            'rating': rating
        }
        actual = self.db['review'].find_one()
        review_id = actual['_id']
        del actual['_id']
        self.assertEqual(expect, actual)

        # check user-review
        expect = docs['user-review']
        expect['review_ids'] = [review_id]
        actual = self.db['user-review'].find_one()
        self.assertEqual(expect, actual)

        # check movie-review
        expect = docs['movie-review']
        expect['review_ids'] = [review_id]
        actual = self.db['movie-review'].find_one()
        self.assertEqual(expect, actual)

        # check rating
        expect = docs['movie-info']
        expect['avg_rating'] = 2.5
        expect['num_rating'] = 6
        actual = self.db['movie-info'].find_one()
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
