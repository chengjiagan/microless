import unittest
import utils
import requests


class TestReducer(utils.TestMapReduce):
    def setUp(self) -> None:
        super().setUp('reducer')

    def test_reducer(self) -> None:
        data = utils.get_json('data/test_reducer.json')

        req = {
            'data': data['input']
        }
        resp = requests.post(self.addr, json=req)

        expect = data['output']
        actual = resp.json()['result']
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
