import utils
from bson import ObjectId
from proto.geo_pb2_grpc import GeoServiceStub
from proto.geo_pb2 import NearbyRequest, AddHotelRequest

class TestGeo(utils.TestHotelReservation):
    stub: GeoServiceStub

    def setUp(self) -> None:
        super().setUp('geo', GeoServiceStub)

    def test_add_hotel(self) -> None:
        hotel_id = '000000000000000000000001'
        lat = 0.1
        lon = 0.1

        req = AddHotelRequest(hotel_id=hotel_id, lat=lat, lon=lon)
        self.stub.AddHotel(req)

        actual = self.db['geo'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'location': {
                'type': 'Point',
                'coordinates': [lon, lat]
            }
        }
        self.assertEqual(actual, expect)

    def test_nearby(self) -> None:
        hotel_id1 = '000000000000000000000001'
        location1 = {
            'hotel_id': ObjectId(hotel_id1),
            'location': {
                'type': 'Point',
                'coordinates': [121.44, 31.03]
            }
        }
        hotel_id2 = '000000000000000000000002'
        location2 = {
            'hotel_id': ObjectId(hotel_id2),
            'location': {
                'type': 'Point',
                'coordinates': [121.51, 31.30]
            }
        }
        hotel_id3 = '000000000000000000000003'
        location3 = {
            'hotel_id': ObjectId(hotel_id3),
            'location': {
                'type': 'Point',
                'coordinates': [121.51, 31.29]
            }
        }
        self.db['geo'].insert_many([location1, location2, location3])

        req = NearbyRequest(lat=31.30, lon=121.51)
        resp = self.stub.Nearby(req)

        actual = resp.hotel_ids
        expect = [hotel_id2, hotel_id3]
        self.assertEqual(actual, expect)
