import unittest
import utils
import requests
from proto.hometimeline_pb2_grpc import HomeTimelineServiceStub
from proto.hometimeline_pb2 import ReadHomeTimelineRequest, ReadHomeTimelineRespond, WriteHomeTimelineRequest, InsertUserResquest
from google.protobuf.timestamp_pb2 import Timestamp
from bson import ObjectId


class TestHomeTimeline(utils.TestSocialNetwork):
    stub: HomeTimelineServiceStub

    def setUp(self) -> None:
        super().setUp('hometimeline', HomeTimelineServiceStub)

    def test_insert_user(self) -> None:
        user_id = '000000000000000000000001'
        req = InsertUserResquest(user_id=user_id)
        self.stub.InsertUser(req)

        expect = {
            'user_id': ObjectId(user_id),
            'post_ids': [],
        }

        actual = self.hometimeline_db.find_one()
        del actual['_id']

        self.assertEqual(expect, actual)

    def test_write_home_timeline(self) -> None:
        user_id = '000000000000000000000000'
        post_id = '000000000000000000000000'
        follower_ids = [
            '000000000000000000000001',
            '000000000000000000000002'
        ]
        mention_ids = [
            '000000000000000000000002',
            '000000000000000000000003'
        ]
        timestamp = Timestamp()
        timestamp.GetCurrentTime()
        self.socialgraph_db.insert_one({
            'user_id': ObjectId(user_id),
            'followers': [ObjectId(i) for i in follower_ids],
            'followees': [],
        })
        for user in set(follower_ids) & set(mention_ids):
            self.hometimeline_db.insert_one(
                {'user_id': ObjectId(user), 'post_ids': []}
            )

        req = WriteHomeTimelineRequest(
            user_id=user_id, post_id=post_id, timestamp=timestamp, user_mentions_id=mention_ids)
        self.stub.WriteHomeTimeline(req)

        for user in set(follower_ids) & set(mention_ids):
            expect = {
                'user_id': ObjectId(user),
                'post_ids': [ObjectId(post_id)],
            }
            actual = self.hometimeline_db.find_one({'user_id': ObjectId(user)})
            del actual['_id']
            self.assertEqual(expect, actual)

    def test_read_home_timeline(self) -> None:
        user_id = '000000000000000000000001'
        posts = utils.get_bson('json/test_read_home_timeline_posts.json')
        self.post_db.insert_many(posts)
        self.hometimeline_db.insert_one({
            'user_id': ObjectId(user_id),
            'post_ids': [
                ObjectId('630f084e6b6cedf0046302ef'),
                ObjectId('630eecc90daff4bcd9a36c40'),
                ObjectId('630eebda0daff4bcd9a36c3e'),
            ]
        })

        req = ReadHomeTimelineRequest(user_id=user_id, start=0, stop=2)
        actual = self.stub.ReadHomeTimeline(req)

        expect = utils.get_proto(
            'json/test_read_home_timeline_proto.json', ReadHomeTimelineRespond)
        self.assertEqual(expect, actual)

    def test_read_home_timeline_rest(self) -> None:
        self.read_home_timeline_rest(self.gateway)

    def read_home_timeline_rest(self, addr: str) -> None:
        user_id = '000000000000000000000001'
        posts = utils.get_bson('json/test_read_home_timeline_posts.json')
        self.post_db.insert_many(posts)
        self.hometimeline_db.insert_one({
            'user_id': ObjectId(user_id),
            'post_ids': [
                ObjectId('630f084e6b6cedf0046302ef'),
                ObjectId('630eecc90daff4bcd9a36c40'),
                ObjectId('630eebda0daff4bcd9a36c3e'),
            ]
        })

        url = 'http://' + addr + '/api/v1/hometimeline/' + user_id
        req = {
            'start': 0,
            'stop': 2
        }
        resp = requests.get(url, params=req)

        actual = resp.json()
        expect = utils.get_json('json/test_read_home_timeline_rest.json')
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
