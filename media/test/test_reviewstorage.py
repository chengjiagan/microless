import unittest

from bson import ObjectId
import utils
from proto.reviewstorage_pb2 import ReadReviewsRequest, ReadReviewsRespond, StoreReviewRequest
from proto.reviewstorage_pb2_grpc import ReviewStorageServiceStub


class TestReviewStorage(utils.TestMedia):
    stub: ReviewStorageServiceStub

    def setUp(self) -> None:
        super().setUp('reviewstorage', ReviewStorageServiceStub)

    def test_store_review(self) -> None:
        req = utils.get_proto(
            'data/test_store_review_proto.json', StoreReviewRequest)
        resp = self.stub.StoreReview(req)
        oid = ObjectId(resp.review_id)

        expect = utils.get_bson('data/test_store_review_bson.json', oid)
        actual = self.db['review'].find_one()
        self.assertEqual(expect, actual)

    def test_read_reviews(self) -> None:
        reviews = utils.get_bson('data/test_read_reviews_bson.json')
        self.db['review'].insert_many(reviews)

        review_ids = ['633942b584c48ef10d77f058', '63394592c051e3a5450fa286']
        req = ReadReviewsRequest(review_ids=review_ids)
        resp = self.stub.ReadReviews(req)

        expect = utils.get_proto(
            'data/test_read_reviews_proto.json', ReadReviewsRespond)
        actual = resp
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
