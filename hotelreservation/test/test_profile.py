import utils
import requests
from bson import ObjectId
from proto.profile_pb2_grpc import ProfileServiceStub
from proto.profile_pb2 import GetProfilesRequest, AddProfileRequest, GetRoomNumberRequest, GetProfilesRespond

class TestProfile(utils.TestHotelReservation):
    stub: ProfileServiceStub

    def setUp(self) -> None:
        super().setUp('profile', ProfileServiceStub)

    def test_get_room_number(self) -> None:
        hotel_id1 = '000000000000000000000001'
        hotels = utils.get_bson('data/test_get_room_number_bson.json')
        self.db['profile'].insert_many(hotels)

        req = GetRoomNumberRequest(hotel_id=hotel_id1)
        resp = self.stub.GetRoomNumber(req)

        actual = resp.room_number
        expect = 300
        self.assertEqual(actual, expect)

    def test_get_profiles(self) -> None:
        hotel_id1 = '000000000000000000000001'
        hotels = utils.get_bson('data/test_get_profiles_bson.json')
        self.db['profile'].insert_many(hotels)

        req = GetProfilesRequest(hotel_ids=[hotel_id1])
        resp = self.stub.GetProfiles(req)

        actual = resp
        expect = utils.get_proto('data/test_get_profiles_proto.json', GetProfilesRespond)
        self.assertEqual(actual, expect)

    def test_add_profile(self) -> None:
        req = utils.get_proto('data/test_add_profile_proto.json', AddProfileRequest)
        resp = self.stub.AddProfile(req)

        actual = self.db['profile'].find_one()
        expect = utils.get_bson('data/test_add_profile_bson.json', oid=ObjectId(resp.hotel_id))
        self.assertEqual(actual, expect)

    def test_add_profile_rest(self) -> None:
        url = f'http://{self.gateway}/api/v1/profile'
        req = utils.get_json('data/test_add_profile_rest_json.json')
        resp = requests.post(url, json=req).json()

        actual = self.db['profile'].find_one()
        expect = utils.get_bson('data/test_add_profile_bson.json', oid=ObjectId(resp['hotelId']))
        self.assertEqual(actual, expect)
