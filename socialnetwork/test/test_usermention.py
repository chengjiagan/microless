import unittest
import utils
from proto.usermention_pb2 import ComposeUserMentionsRequest, ComposeUserMentionsRespond
from proto.usermention_pb2_grpc import UserMentionServiceStub


class TestUserMention(utils.TestSocialNetwork):
    stub: UserMentionServiceStub

    def setUp(self) -> None:
        super().setUp("usermention", UserMentionServiceStub)

    def test_compose_user_mentions(self) -> None:
        self.user_db.insert_many(utils.get_bson(
            'json/test_compose_user_mentions_users.json'))

        usernames = ['username_1', 'username_2', 'username_3']
        req = ComposeUserMentionsRequest(usernames=usernames)
        resp = self.stub.ComposeUserMentions(req)

        actual = resp
        expect = utils.get_proto(
            'json/test_compose_user_mentions_proto.json', ComposeUserMentionsRespond)
        self.assertEqual(expect, actual)

if __name__ == '__main__':
    unittest.main()
