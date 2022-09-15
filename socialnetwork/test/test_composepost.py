import unittest
import requests
import utils
from proto.composepost_pb2_grpc import ComposePostServiceStub
from proto.composepost_pb2 import ComposePostRequest
from proto.data_pb2 import PostType
from bson import ObjectId


class TestComposePost(utils.TestSocialNetwork):
    stub: ComposePostServiceStub

    def setUp(self) -> None:
        super().setUp('composepost', ComposePostServiceStub)

    def test_compose_post(self) -> None:
        user_id_0 = '000000000000000000000000'
        user_id_1 = '000000000000000000000001'
        user_id_2 = '000000000000000000000002'
        user_id_3 = '000000000000000000000003'
        followers = [user_id_2, user_id_3]
        mentions = [user_id_1, user_id_2]

        self.user_db.insert_many(utils.get_bson(
            'json/test_compose_post_users.json'))
        self.socialgraph_db.insert_one({
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(i) for i in followers],
            'followees': []
        })
        self.timeline_db.insert_one({
            'user_id': ObjectId(user_id_0),
            'post_ids': []
        })

        username = 'username_0'
        text = utils.get_text('json/test_compose_post_text.txt')
        media_ids = [0, 1]
        media_types = ['png', 'png']
        post_type = PostType.POST
        req = ComposePostRequest(username=username, user_id=user_id_0, text=text,
                                 media_ids=media_ids, media_types=media_types, post_type=post_type)
        self.stub.ComposePost(req)

        # shortened urls are generated dynamically
        exp_urls = ["https://url_0.com", "https://url_1.com"]
        urls = []
        for u in exp_urls:
            doc = self.url_db.find_one({'expanded_url': u})
            del doc['_id']
            urls.append(doc)

        post = utils.get_bson('json/test_compose_post_post.json')
        post['urls'] = urls

        # should upload correct post
        actual = self.post_db.find_one()
        post_id = actual['_id']
        del actual['_id']
        expect = post
        self.assertEqual(expect, actual)

        # should update user's timeline
        actual = self.timeline_db.find_one()
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_0),
            'post_ids': [post_id]
        }
        self.assertEqual(expect, actual)

        # should update followers' and mentioned users' timeline
        expect = [str(post_id)]
        for i in set(followers) & set(mentions):
            actual = self.timeline_redis.zrange(i, 0, -1)
            self.assertEqual(expect, actual)

    def test_compose_post_rest(self) -> None:
        self.compose_post_rest(self.gateway)

    def compose_post_rest(self, addr: str) -> None:
        user_id_0 = '000000000000000000000000'
        user_id_1 = '000000000000000000000001'
        user_id_2 = '000000000000000000000002'
        user_id_3 = '000000000000000000000003'
        followers = [user_id_2, user_id_3]
        mentions = [user_id_1, user_id_2]

        self.user_db.insert_many(utils.get_bson(
            'json/test_compose_post_users.json'))
        self.socialgraph_db.insert_one({
            'user_id': ObjectId(user_id_0),
            'followers': [ObjectId(i) for i in followers],
            'followees': []
        })
        self.timeline_db.insert_one({
            'user_id': ObjectId(user_id_0),
            'post_ids': []
        })

        url = 'http://' + addr + '/api/v1/composepost'
        req = {
            'username': 'username_0',
            'text': utils.get_text('json/test_compose_post_text.txt'),
            'mediaIds': [0, 1],
            'mediaTypes': ['png', 'png'],
            'postType': "POST"
        }
        resp = requests.post(url, json=req)

        # shortened urls are generated dynamically
        exp_urls = ["https://url_0.com", "https://url_1.com"]
        urls = []
        for u in exp_urls:
            doc = self.url_db.find_one({'expanded_url': u})
            del doc['_id']
            urls.append(doc)

        post = utils.get_bson('json/test_compose_post_post.json')
        post['urls'] = urls

        # should upload correct post
        actual = self.post_db.find_one()
        post_id = actual['_id']
        del actual['_id']
        expect = post
        self.assertEqual(expect, actual)

        # should update user's timeline
        actual = self.timeline_db.find_one()
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id_0),
            'post_ids': [post_id]
        }
        self.assertEqual(expect, actual)

        # should update followers' and mentioned users' timeline
        expect = [str(post_id)]
        for i in set(followers) & set(mentions):
            actual = self.timeline_redis.zrange(i, 0, -1)
            self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
