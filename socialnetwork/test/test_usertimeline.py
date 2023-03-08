#! /usr/bin/env python3
import unittest
import utils
import requests
from bson import ObjectId
from proto.usertimeline_pb2_grpc import UserTimelineServiceStub
from proto.usertimeline_pb2 import InsertUserResquest, ReadUserTimelineRequest, ReadUserTimelineRespond, WriteUserTimelineRequest


class TestUserTimeline(utils.TestSocialNetwork):
    stub: UserTimelineServiceStub

    def setUp(self) -> None:
        super().setUp('usertimeline', UserTimelineServiceStub)

    def test_insert_user(self) -> None:
        user_id = '000000000000000000000001'
        req = InsertUserResquest(user_id=user_id)
        self.stub.InsertUser(req)

        expect = {
            'user_id': ObjectId(user_id),
            'post_ids': [],
        }

        actual = self.usertimeline_db.find_one()
        del actual['_id']

        self.assertEqual(expect, actual)

    def test_write_user_timeline(self) -> None:
        user_id = '000000000000000000000001'
        self.usertimeline_db.insert_one(
            {'user_id': ObjectId(user_id), 'post_ids': []})

        # unordered write requests
        post_id_0 = '000000000000000000000001'
        self.stub.WriteUserTimeline(WriteUserTimelineRequest(
            user_id=user_id, post_id=post_id_0))
        post_id_1 = '000000000000000000000010'
        self.stub.WriteUserTimeline(WriteUserTimelineRequest(
            user_id=user_id, post_id=post_id_1))
        post_id_2 = '000000000000000000000005'
        self.stub.WriteUserTimeline(WriteUserTimelineRequest(
            user_id=user_id, post_id=post_id_2))

        expect = {
            'user_id': ObjectId(user_id),
            'post_ids': [
                ObjectId(post_id_2),
                ObjectId(post_id_1),
                ObjectId(post_id_0),
            ],
        }

        actual = self.usertimeline_db.find_one()
        del actual['_id']

        self.assertEqual(expect, actual)

    def test_read_user_timeline(self) -> None:
        user_id = '000000000000000000000001'
        self.post_db.insert_many(utils.get_bson(
            'json/test_read_user_timeline_posts.json'))
        self.usertimeline_db.insert_one({
            'user_id': ObjectId(user_id),
            'post_ids': [
                ObjectId('630f084e6b6cedf0046302ef'),
                ObjectId('630eecc90daff4bcd9a36c40'),
                ObjectId('630eebda0daff4bcd9a36c3e'),
            ]
        })

        req = ReadUserTimelineRequest(user_id=user_id, start=0, stop=2)
        resp = self.stub.ReadUserTimeline(req)

        expect = utils.get_proto(
            'json/test_read_user_timeline_proto.json', ReadUserTimelineRespond)

        self.assertEqual(expect, resp)

    def test_read_user_timeline_rest(self) -> None:
        self.read_user_timeline_rest(self.gateway)

    def read_user_timeline_rest(self, addr: str) -> None:
        user_id = '000000000000000000000001'
        self.post_db.insert_many(utils.get_bson(
            'json/test_read_user_timeline_posts.json'))
        self.usertimeline_db.insert_one({
            'user_id': ObjectId(user_id),
            'post_ids': [
                ObjectId('630f084e6b6cedf0046302ef'),
                ObjectId('630eecc90daff4bcd9a36c40'),
                ObjectId('630eebda0daff4bcd9a36c3e'),
            ]
        })

        url = 'http://' + addr + '/api/v1/usertimeline/' + user_id
        req = {
            'start': 0,
            'stop': 2
        }
        resp = requests.get(url, params=req)

        actual = resp.json()
        expect = utils.get_json('json/test_read_user_timeline_rest.json')
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
