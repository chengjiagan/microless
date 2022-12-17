import utils
import requests
from datetime import datetime
from bson import ObjectId
from google.protobuf.timestamp_pb2 import Timestamp
from proto.reservation_pb2_grpc import ReservationServiceStub
from proto.reservation_pb2 import MakeReservationRequest, CheckAvailabilityRequest

class TestReservation(utils.TestHotelReservation):
    stub: ReservationServiceStub

    def setUp(self) -> None:
        super().setUp('reservation', ReservationServiceStub)

    def test_make_reservation(self) -> None:
        hotel_id = '000000000000000000000001'
        hotel = {
            '_id': ObjectId(hotel_id),
            'name': 'hotel',
            'room_number': 10
        }
        self.db['profile'].insert_one(hotel)

        user_id = '000000000000000000000001'
        in_date = Timestamp()
        in_date.FromDatetime(datetime(2022, 12, 1))
        out_date = Timestamp()
        out_date.FromDatetime(datetime(2022, 12, 5))
        req = MakeReservationRequest(
            hotel_id=hotel_id,
            user_id=user_id,
            in_date=in_date,
            out_date=out_date,
            room_number=3
        )
        self.stub.MakeReservation(req)

        actual = self.db['reservation'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 1),
            'out_date': datetime(2022, 12, 5),
            'room_number': 3
        }
        self.assertEqual(expect, actual)

    def test_make_reservation_rest(self) -> None:
        hotel_id = '000000000000000000000001'
        hotel = {
            '_id': ObjectId(hotel_id),
            'name': 'hotel',
            'room_number': 10
        }
        self.db['profile'].insert_one(hotel)

        url = f'http://{self.gateway}/api/v1/reservation'
        user_id = '000000000000000000000001'
        req = {
            'hotel_id': hotel_id,
            'user_id': user_id,
            'in_date': '2022-12-01T00:00:00Z',
            'out_date': '2022-12-05T00:00:00Z',
            'room_number': 3
        }
        requests.post(url, json=req)

        actual = self.db['reservation'].find_one()
        del actual['_id']
        expect = {
            'hotel_id': ObjectId(hotel_id),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 1),
            'out_date': datetime(2022, 12, 5),
            'room_number': 3
        }
        self.assertEqual(expect, actual)

    def test_check_availability(self) -> None:
        hotel_id1 = '000000000000000000000001'
        hotel1 = {
            '_id': ObjectId(hotel_id1),
            'name': 'hotel1',
            'room_number': 10
        }
        hotel_id2 = '000000000000000000000002'
        hotel2 = {
            '_id': ObjectId(hotel_id2),
            'name': 'hotel2',
            'room_number': 10
        }
        hotel_id3 = '000000000000000000000003'
        hotel3 = {
            '_id': ObjectId(hotel_id3),
            'name': 'hotel3',
            'room_number': 10
        }
        self.db['profile'].insert_many([hotel1, hotel2, hotel3])
        user_id = '000000000000000000000001'
        res1 = {
            'hotel_id': ObjectId(hotel_id1),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 1),
            'out_date': datetime(2022, 12, 5),
            'room_number': 1
        }
        res2 = {
            'hotel_id': ObjectId(hotel_id2),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 1),
            'out_date': datetime(2022, 12, 5),
            'room_number': 8
        }
        res3 = {
            'hotel_id': ObjectId(hotel_id3),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 1),
            'out_date': datetime(2022, 12, 4),
            'room_number': 4
        }
        res4 = {
            'hotel_id': ObjectId(hotel_id3),
            'user_id': ObjectId(user_id),
            'in_date': datetime(2022, 12, 2),
            'out_date': datetime(2022, 12, 5),
            'room_number': 4
        }
        self.db['reservation'].insert_many([res1, res2, res3, res4])

        in_date = Timestamp()
        in_date.FromDatetime(datetime(2022, 12, 1))
        out_date = Timestamp()
        out_date.FromDatetime(datetime(2022, 12, 5))
        req = CheckAvailabilityRequest(
            hotel_ids=[hotel_id1, hotel_id2, hotel_id3],
            in_date=in_date,
            out_date=out_date,
            room_number=3
        )
        resp = self.stub.CheckAvailability(req)

        actual = resp.hotel_ids
        expect = [hotel_id1]
        self.assertEqual(expect, actual)
