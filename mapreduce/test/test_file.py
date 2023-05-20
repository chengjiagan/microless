import io
import unittest
import utils
import requests


class TestFile(utils.TestMapReduce):
    def setUp(self) -> None:
        super().setUp('file')

    def test_file(self) -> None:
        data = io.BytesIO(b'')
        self.storage.put_object(utils.BUCKET, 'test/1', data, 0)
        self.storage.put_object(utils.BUCKET, 'test/2', data, 0)
        self.storage.put_object(utils.BUCKET, 'test/3', data, 0)
        self.storage.put_object(utils.BUCKET, 'nontest/1', data, 0)

        req = {
            'bucket': utils.BUCKET,
            'prefix': 'test/'
        }
        resp = requests.post(self.addr, json=req)

        expect = ['test/1', 'test/2', 'test/3']
        actual = resp.json()['filenames']
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
