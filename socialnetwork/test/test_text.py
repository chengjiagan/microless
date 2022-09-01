import unittest
import utils
from proto.text_pb2_grpc import TextServiceStub
from proto.text_pb2 import ComposeTextRequest
from proto.data_pb2 import UserMention


class TestText(utils.TestSocialNetwork):
    stub: TextServiceStub

    def setUp(self) -> None:
        super().setUp('text', TextServiceStub)

    def test_compose_text(self) -> None:
        self.user_db.insert_many(utils.get_bson('json/test_compose_text_users.json'))

        text = utils.get_text('json/test_compose_text_text.txt')
        req = ComposeTextRequest(text=text)
        resp = self.stub.ComposeText(req)

        actual = resp.text
        expect = text
        self.assertEqual(expect, actual)

        mentions = [
            UserMention(user_id='000000000000000000000001', username='username_1'),
            UserMention(user_id='000000000000000000000002', username='username_2'),
            UserMention(user_id='000000000000000000000003', username='username_3'),
            UserMention(user_id='000000000000000000000004', username='username_4'),
        ]
        for actual, expect in zip(resp.user_mention, mentions):
            self.assertEqual(expect, actual)

        urls = ["https://url_0.com", "https://url_1.com", "https://url_2.com"]
        for url, exp_url in zip(resp.urls, urls):
            actual = url.expanded_url
            expect = exp_url
            self.assertEqual(expect, actual)

            actual = self.url_db.find_one({'shortened_url': url.shortened_url})
            del actual['_id']
            expect = {
                'expanded_url': exp_url,
                'shortened_url': url.shortened_url,
            }
            self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
