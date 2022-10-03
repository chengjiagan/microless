import unittest

from bson import ObjectId
import requests
import utils
from proto.userreview_pb2 import CreateUserRequest, ReadUserReviewsRequest, ReadUserReviewsRespond, UploadUserReviewRequest
from proto.userreview_pb2_grpc import UserReviewServiceStub
from google.protobuf.timestamp_pb2 import Timestamp


class TestUserReview(utils.TestMedia):
    stub: UserReviewServiceStub

    def setUp(self) -> None:
        super().setUp('userreview', UserReviewServiceStub)

    def test_upload_user_review(self) -> None:
        user_id = '000000000000000000000000'
        self.db['user-review'].insert_one({
            'user_id': ObjectId(user_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('633942b584c48ef10d77f058')]
        })

        dt = ObjectId('6339457ac051e3a5450fa284').generation_time
        timestamp = Timestamp()
        timestamp.FromDatetime(dt)
        req = UploadUserReviewRequest(
            user_id=user_id, review_id='6339457ac051e3a5450fa284', timestamp=timestamp)
        self.stub.UploadUserReview(req)

        expect = {
            'user_id': ObjectId(user_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('6339457ac051e3a5450fa284'), ObjectId('633942b584c48ef10d77f058')]
        }
        actual = self.db['user-review'].find_one()
        del actual['_id']
        self.assertEqual(expect, actual)

    def test_read_user_reviews(self) -> None:
        user_id = '000000000000000000000000'
        self.db['user-review'].insert_one({
            'user_id': ObjectId(user_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('6339457ac051e3a5450fa284'), ObjectId('633942b584c48ef10d77f058')]
        })
        self.db['review'].insert_many(utils.get_bson(
            'data/test_read_user_reviews_bson.json'))

        req = ReadUserReviewsRequest(user_id=user_id, start=1, stop=3)
        resp = self.stub.ReadUserReviews(req)

        expect = utils.get_proto(
            'data/test_read_user_reviews_proto.json', ReadUserReviewsRespond)
        actual = resp
        self.assertEqual(expect, actual)

    def test_read_user_reviews_rest(self) -> None:
        user_id = '000000000000000000000000'
        self.db['user-review'].insert_one({
            'user_id': ObjectId(user_id),
            'review_ids': [ObjectId('63394592c051e3a5450fa286'), ObjectId('6339457ac051e3a5450fa284'), ObjectId('633942b584c48ef10d77f058')]
        })
        self.db['review'].insert_many(utils.get_bson(
            'data/test_read_user_reviews_bson.json'))

        url = 'http://' + self.gateway + '/api/v1/userreview/' + user_id
        req = {
            'start': 1,
            'stop': 3
        }
        resp = requests.get(url, params=req)

        expect = utils.get_json('data/test_read_user_reviews_rest.json')
        actual = resp.json()
        self.assertEqual(expect, actual)

    def test_create_user(self) -> None:
        user_id = '000000000000000000000000'
        req = CreateUserRequest(user_id=user_id)
        self.stub.CreateUser(req)

        expect = {
            "user_id": ObjectId(user_id),
            "review_ids": []
        }
        actual = self.db['user-review'].find_one()
        del actual['_id']
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
