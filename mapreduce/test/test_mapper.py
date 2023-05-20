import unittest
import utils
import requests


class TestMapper(utils.TestMapReduce):
    def setUp(self) -> None:
        super().setUp('mapper')

    def test_mapper(self) -> None:
        self.storage.fput_object(utils.BUCKET, 'test_mapper.parquet', 'data/yellow_tripdata_2021-01.parquet')

        req = {
            'bucket': utils.BUCKET,
            'file': 'test_mapper.parquet'
        }
        resp = requests.post(self.addr, json=req)

        expect = utils.get_json('data/test_mapper.json')
        actual = resp.json()['result']
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
