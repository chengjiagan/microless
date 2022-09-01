import unittest
import utils
import requests
from bson import ObjectId
from proto.socialgraph_pb2_grpc import SocialGraphServiceStub
from proto.socialgraph_pb2 import FollowRequest, GetFolloweesRequest, GetFolloweesRespond, GetFollowersRequest, GetFollowersRespond, InsertUserRequest, UnfollowRequest


class TestSocialGraph(utils.TestSocialNetwork):
    stub: SocialGraphServiceStub

    def setUp(self) -> None:
        super().setUp('socialgraph', SocialGraphServiceStub)

    def test_insert_user(self) -> None:
        user_id = '000000000000000000000001'
        req = InsertUserRequest(user_id=user_id)
        self.stub.InsertUser(req)

        actual = self.socialgraph_db.find_one()
        del actual['_id']

        expect = {
            'user_id': ObjectId(user_id),
            'followers': [],
            'followees': [],
        }

        self.assertEqual(expect, actual)

    def test_social_graph(self) -> None:
        user_id_0 = '000000000000000000000000'
        user_id_1 = '000000000000000000000001'
        user_id_2 = '000000000000000000000002'

        self.socialgraph_db.insert_many([
            {'user_id': ObjectId(user_id_0), 'followers': [], 'followees': []},
            {'user_id': ObjectId(user_id_1), 'followers': [], 'followees': []},
            {'user_id': ObjectId(user_id_2), 'followers': [], 'followees': []},
        ])

        # follow
        self.stub.Follow(FollowRequest(
            user_id=user_id_0, followee_id=user_id_1))
        self.stub.Follow(FollowRequest(
            user_id=user_id_0, followee_id=user_id_2))
        self.stub.Follow(FollowRequest(
            user_id=user_id_1, followee_id=user_id_2))
        self.stub.Follow(FollowRequest(
            user_id=user_id_1, followee_id=user_id_0))
        self.stub.Follow(FollowRequest(
            user_id=user_id_2, followee_id=user_id_1))
        self.stub.Follow(FollowRequest(
            user_id=user_id_2, followee_id=user_id_0))
        self.stub.Follow(FollowRequest(
            user_id=user_id_2, followee_id=user_id_0))

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_0)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(user_id_1), ObjectId(user_id_2)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_2)],
        }
        self.assertEqual(expect, actual)

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_1)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_1),
            'followers': [ObjectId(user_id_0), ObjectId(user_id_2)],
            'followees': [ObjectId(user_id_2), ObjectId(user_id_0)],
        }
        self.assertEqual(expect, actual)

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_2)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_2),
            'followers': [ObjectId(user_id_0), ObjectId(user_id_1)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_0)],
        }
        self.assertEqual(expect, actual)

        # unfollow
        self.stub.Unfollow(UnfollowRequest(
            user_id=user_id_1, followee_id=user_id_0))
        self.stub.Unfollow(UnfollowRequest(
            user_id=user_id_1, followee_id=user_id_2))

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_0)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(user_id_2)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_2)],
        }
        self.assertEqual(expect, actual)

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_1)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_1),
            'followers': [ObjectId(user_id_0), ObjectId(user_id_2)],
            'followees': [],
        }
        self.assertEqual(expect, actual)

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_2)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_2),
            'followers': [ObjectId(user_id_0)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_0)],
        }
        self.assertEqual(expect, actual)

        # follow again
        self.stub.Follow(FollowRequest(
            user_id=user_id_1, followee_id=user_id_0))
        self.stub.Follow(FollowRequest(
            user_id=user_id_1, followee_id=user_id_2))

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_0)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(user_id_2), ObjectId(user_id_1)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_2)],
        }
        self.assertEqual(expect, actual)

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_1)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_1),
            'followers': [ObjectId(user_id_0), ObjectId(user_id_2)],
            'followees': [ObjectId(user_id_0), ObjectId(user_id_2)],
        }
        self.assertEqual(expect, actual)

        actual = self.socialgraph_db.find_one({'user_id': ObjectId(user_id_2)})
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_2),
            'followers': [ObjectId(user_id_0), ObjectId(user_id_1)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_0)],
        }
        self.assertEqual(expect, actual)

        actual = self.stub.GetFollowers(GetFollowersRequest(user_id=user_id_0))
        expect = GetFollowersRespond(followers_id=[user_id_2, user_id_1])
        self.assertEqual(expect, actual)

        actual = self.stub.GetFollowees(GetFolloweesRequest(user_id=user_id_0))
        expect = GetFolloweesRespond(followees_id=[user_id_1, user_id_2])
        self.assertEqual(expect, actual)

        actual = self.stub.GetFollowers(GetFollowersRequest(user_id=user_id_1))
        expect = GetFollowersRespond(followers_id=[user_id_0, user_id_2])
        self.assertEqual(expect, actual)

        actual = self.stub.GetFollowees(GetFolloweesRequest(user_id=user_id_1))
        expect = GetFolloweesRespond(followees_id=[user_id_0, user_id_2])
        self.assertEqual(expect, actual)

        actual = self.stub.GetFollowers(GetFollowersRequest(user_id=user_id_2))
        expect = GetFollowersRespond(followers_id=[user_id_0, user_id_1])
        self.assertEqual(expect, actual)

        actual = self.stub.GetFollowees(GetFolloweesRequest(user_id=user_id_2))
        expect = GetFolloweesRespond(followees_id=[user_id_1, user_id_0])
        self.assertEqual(expect, actual)

    def test_get_followers_rest(self) -> None:
        user_id_0 = '000000000000000000000000'
        user_id_1 = '000000000000000000000001'
        user_id_2 = '000000000000000000000002'
        self.socialgraph_db.insert_one({
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(user_id_2), ObjectId(user_id_1)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_2)],
        })

        url = 'http://' + self.rest['socialgraph'] + '/api/v1/socialgraph/followers/' + user_id_0
        resp = requests.get(url)

        actual = resp.json()
        expect = {'followersId': [user_id_2, user_id_1]}
        self.assertEqual(expect, actual)

    def test_get_followees_rest(self) -> None:
        user_id_0 = '000000000000000000000000'
        user_id_1 = '000000000000000000000001'
        user_id_2 = '000000000000000000000002'
        self.socialgraph_db.insert_one({
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(user_id_2), ObjectId(user_id_1)],
            'followees': [ObjectId(user_id_1), ObjectId(user_id_2)],
        })

        url = 'http://' + self.rest['socialgraph'] + '/api/v1/socialgraph/followees/' + user_id_0
        resp = requests.get(url)

        actual = resp.json()
        expect = {'followeesId': [user_id_1, user_id_2]}
        self.assertEqual(expect, actual)

if __name__ == '__main__':
    unittest.main()
