import unittest

from bson import ObjectId
import utils
from proto.castinfo_pb2_grpc import CastInfoServiceStub
from proto.castinfo_pb2 import ReadCastInfoRequest, ReadCastInfoRespond, WriteCastInfoRequest


class TestCastInfo(utils.TestMedia):
    stub: CastInfoServiceStub

    def setUp(self) -> None:
        super().setUp('castinfo', CastInfoServiceStub)

    def test_write_cast_info(self) -> None:
        name = 'name_0'
        gender = True
        intro = 'intro_0'
        req = WriteCastInfoRequest(name=name, gender=gender, intro=intro)
        resp = self.stub.WriteCastInfo(req)

        expect = {
            '_id': ObjectId(resp.cast_info_id),
            'name': name,
            'gender': gender,
            'intro': intro
        }
        actual = self.db['cast-info'].find_one()
        self.assertEqual(expect, actual)

    def test_read_cast_info(self) -> None:
        infos = utils.get_bson('data/test_read_cast_info_bson.json')
        self.db['cast-info'].insert_many(infos)

        info_ids = ['000000000000000000000000', '000000000000000000000002']
        req = ReadCastInfoRequest(cast_ids=info_ids)
        resp = self.stub.ReadCastInfo(req)

        expect = utils.get_proto(
            'data/test_read_cast_info_proto.json', ReadCastInfoRespond)
        actual = resp
        self.assertEqual(expect, actual)


if __name__ == '__main__':
    unittest.main()
