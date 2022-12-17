import utils
import requests
from bson import ObjectId
from datetime import datetime
from google.protobuf.timestamp_pb2 import Timestamp
from proto.rate_pb2_grpc import RateServiceStub
from proto.rate_pb2 import GetRatesRequest, AddRateRequest, HotelRate

class TestRate(utils.TestHotelReservation):
    stub: RateServiceStub

    def setUp(self) -> None:
        super().setUp('rate', RateServiceStub)

    def test_get_rates(self) -> None:
        hotel_id1 = '000000000000000000000001'
        rate1 = {
            'hotel_id': ObjectId(hotel_id1),
            'total_rate': 10,
            'num_rate': 4
        }
        hotel_id2 = '000000000000000000000002'
        rate2 = {
            'hotel_id': ObjectId(hotel_id2),
            'total_rate': 12,
            'num_rate': 3
        }
        self.db['hotel-rate'].insert_many([rate1, rate2])

        req = GetRatesRequest(hotel_ids=[hotel_id1, hotel_id2])
        resp = self.stub.GetRates(req)

        result1 = HotelRate(hotel_id=hotel_id1, rate=2.5)
        result2 = HotelRate(hotel_id=hotel_id2, rate=4.0)

        actual = resp.rates
        expect = [result1, result2]
        self.assertEqual(actual, expect)

    def test_add_rate(self) -> None:
        hotel_id = '000000000000000000000001'
        hotel_rate = {
            'hotel_id': ObjectId(hotel_id),
            'total_rate': 0,
            'num_rate': 0
        }
        self.db['hotel-rate'].insert_one(hotel_rate)

        user_id = '000000000000000000000002'
        in_date = Timestamp()
        in_date.FromDatetime(datetime(2022, 12, 15))
        out_date = Timestamp()
        out_date.FromDatetime(datetime(2022, 12, 18))
        rate = 3
        req = AddRateRequest(
            hotel_id=hotel_id,
            user_id=user_id,
            in_date=in_date,
            out_date=out_date,
            rate=rate
        )
        self.stub.AddRate(req)

        # new rate plan
        actual = self.db['rate-plan'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 15),
            'out_date': datetime(2022, 12, 18),
            'rate': rate
        }
        self.assertEqual(actual, expect)

        # update hotel rate
        actual = self.db['hotel-rate'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'total_rate': 3,
            'num_rate': 1
        }
        self.assertEqual(actual, expect)

    def test_add_rate_rest(self) -> None:
        hotel_id = '000000000000000000000001'
        hotel_rate = {
            'hotel_id': ObjectId(hotel_id),
            'total_rate': 0,
            'num_rate': 0
        }
        self.db['hotel-rate'].insert_one(hotel_rate)

        url = f'http://{self.gateway}/api/v1/rate'
        user_id = '000000000000000000000002'
        in_date = '2022-12-15T00:00:00Z'
        out_date = '2022-12-18T00:00:00Z'
        rate = 3
        req = {
            'hotel_id': hotel_id,
            'user_id': user_id,
            'in_date': in_date,
            'out_date': out_date,
            'rate': rate
        }
        requests.post(url, json=req)

        # new rate plan
        actual = self.db['rate-plan'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 15),
            'out_date': datetime(2022, 12, 18),
            'rate': rate
        }
        self.assertEqual(actual, expect)

        # update hotel rate
        actual = self.db['hotel-rate'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'total_rate': 3,
            'num_rate': 1
        }
        self.assertEqual(actual, expect)
