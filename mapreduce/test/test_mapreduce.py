import unittest
import utils
import requests


class TestMapreduce(utils.TestMapReduce):
    def setUp(self) -> None:
        super().setUp('mapreduce')

    def test_mapreduce(self) -> None:
        self.storage.fput_object(
            utils.BUCKET, 'yellow_tripdata_2021-01.parquet', 'data/yellow_tripdata_2021-01.parquet')
        self.storage.fput_object(
            utils.BUCKET, 'yellow_tripdata_2021-02.parquet', 'data/yellow_tripdata_2021-02.parquet')

        req = {
            'bucket': utils.BUCKET,
            'prefix': ''
        }
        resp = requests.post(self.addr, json=req)

        expect = utils.get_json('data/test_mapreduce.json')
        actual = resp.json()['result']
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
