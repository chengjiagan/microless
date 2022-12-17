import utils
import requests
from bson import ObjectId
from datetime import datetime
from google.protobuf.timestamp_pb2 import Timestamp
from proto.search_pb2_grpc import SearchServiceStub
from proto.search_pb2 import SearchRequest, SearchRespond

class TestSearch(utils.TestHotelReservation):
    stub: SearchServiceStub

    def setUp(self) -> None:
        super().setUp('search', SearchServiceStub)

    def test_search(self) -> None:
        hotels = utils.get_bson('data/test_search_profile_bson.json')
        self.db['profile'].insert_many(hotels)
        reservations = utils.get_bson('data/test_search_reservation_bson.json')
        self.db['reservation'].insert_many(reservations)
        rates = utils.get_bson('data/test_search_rate_bson.json')
        self.db['hotel-rate'].insert_many(rates)
        geos = utils.get_bson('data/test_search_geo_bson.json')
        self.db['geo'].insert_many(geos)

        in_date = Timestamp()
        in_date.FromDatetime(datetime(2022, 12, 1))
        out_date = Timestamp()
        out_date.FromDatetime(datetime(2022, 12, 5))
        req = SearchRequest(
            lat=31.30, lon=121.51,
            in_date=in_date, out_date=out_date,
            room_number=3
        )
        resp = self.stub.Search(req)

        actual = resp
        expect = utils.get_proto('data/test_search_result_proto.json', SearchRespond)
        self.assertEqual(expect, actual)

    def test_search_rest(self) -> None:
        hotels = utils.get_bson('data/test_search_profile_bson.json')
        self.db['profile'].insert_many(hotels)
        reservations = utils.get_bson('data/test_search_reservation_bson.json')
        self.db['reservation'].insert_many(reservations)
        rates = utils.get_bson('data/test_search_rate_bson.json')
        self.db['hotel-rate'].insert_many(rates)
        geos = utils.get_bson('data/test_search_geo_bson.json')
        self.db['geo'].insert_many(geos)

        url = f'http://{self.gateway}/api/v1/search'
        req = {
            "lat": 31.30,
            "lon": 121.51,
            "in_date": "2022-12-01T00:00:00Z",
            "out_date": "2022-12-05T00:00:00Z",
            "room_number": 3
        }
        resp = requests.post(url, json=req)

        actual = resp.json()
        expect = utils.get_json('data/test_search_result_json.json')
        self.assertEqual(expect, actual)