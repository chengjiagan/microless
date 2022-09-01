import unittest
import utils
from proto.urlshorten_pb2_grpc import UrlShortenServiceStub
from proto.urlshorten_pb2 import ComposeUrlsRequest, GetExtendedUrlsRequest, GetExtendedUrlsRespond


class TestUrlShorten(utils.TestSocialNetwork):
    stub: UrlShortenServiceStub

    def setUp(self) -> None:
        super().setUp("urlshorten", UrlShortenServiceStub)

    def test_compose_urls(self) -> None:
        urls = ["https://url_0.com", "https://url_1.com", "https://url_2.com"]
        req = ComposeUrlsRequest(urls=urls)
        resp = self.stub.ComposeUrls(req)

        for url in resp.urls:
            actual = self.url_db.find_one({'shortened_url': url.shortened_url})
            del actual['_id']
            expect = {
                'expanded_url': url.expanded_url,
                'shortened_url': url.shortened_url,
            }
            self.assertEqual(expect, actual)
            self.assertTrue(url.shortened_url.startswith('http://short-url/'))

    def test_get_extended_urls(self) -> None:
        self.url_db.insert_many(utils.get_bson(
            'json/test_get_extended_urls_bson.json'))

        urls = [
            "http://short-url/UsaD6HEdz0",
            "http://short-url/ThbXfQ6pYS",
            "http://short-url/Q3n267l1VQ",
        ]
        req = GetExtendedUrlsRequest(shortened_urls=urls)
        actual = self.stub.GetExtendedUrls(req)

        exp_urls = ["https://url_0.com", "https://url_1.com", "https://url_2.com"]
        expect = GetExtendedUrlsRespond(urls=exp_urls)

        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
