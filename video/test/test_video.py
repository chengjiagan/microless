import unittest
import utils
import requests


class TestMapper(utils.TestVideo):
    def setUp(self) -> None:
        super().setUp('video')

    def test_video_simple(self) -> None:
        self.storage.fput_object(utils.BUCKET, 'test.mp4', 'data/test.mp4')

        req = {
            'bucket': utils.BUCKET,
            'input': 'test.mp4',
            'segment_time': 0,
            'type': 'avi',
            'outdir': 'out'
        }
        resp = requests.post(self.addr, json=req)

        expect = {
            'preview': 'out/test_preview.gif',
            'transcoded': 'out/test.avi'
        }
        actual = resp.json()
        self.assertEqual(expect, actual)

        expect = {'out/test.avi', 'out/test_preview.gif'}
        actual = {o.object_name for o in self.storage.list_objects(utils.BUCKET, 'out/')}
        self.assertEqual(expect, actual)

    def test_video_split(self) -> None:
        self.storage.fput_object(utils.BUCKET, 'test.mp4', 'data/test.mp4')

        req = {
            'bucket': utils.BUCKET,
            'input': 'test.mp4',
            'segment_time': 20,
            'type': 'avi',
            'outdir': 'out'
        }
        resp = requests.post(self.addr, json=req)

        expect = {
            'preview': 'out/test_preview.gif',
            'transcoded': 'out/test.avi'
        }
        actual = resp.json()
        self.assertEqual(expect, actual)

        expect = {'out/test.avi', 'out/test_preview.gif'}
        actual = {o.object_name for o in self.storage.list_objects(utils.BUCKET, 'out/')}
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
