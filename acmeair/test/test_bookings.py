import utils
import requests
from bson import ObjectId
from proto.data_pb2 import BookingInfo
from proto.bookings_pb2_grpc import BookingsServiceStub
from proto.bookings_pb2 import BookFlightsRequest, GetBookingByIdRequest, GetBookingByUserRequest, GetBookingByUserRespond, CancelBookingByIdRequest
from google.protobuf.json_format import ParseDict

class TestBookings(utils.TestAcmeair):
    stub: BookingsServiceStub

    def setUp(self) -> None:
        super().setUp('bookings', BookingsServiceStub)
        # insert test flights
        flights = utils.get_bson('data/test_bookings_flights_bson.json')
        self.db['flights'].insert_many(flights)
        # insert test customers
        customers = utils.get_bson('data/test_bookings_customers_bson.json')
        self.db['customer'].insert_many(customers)

    def test_book_flights(self) -> None:
        flight_id = '63bfa0ec6ced8e7900d0516e'
        customer_id = '000000000000000000000001'
        req = BookFlightsRequest(
            customer_id=customer_id,
            to_flight_id=flight_id,
            one_way_flight=True
        )
        resp = self.stub.BookFlights(req)

        # check booking info
        actual = resp.booking
        o = utils.get_json('data/test_book_flights_proto.json')
        o['bookingId'] = actual.booking_id
        o['dateOfBooking'] = ObjectId(actual.booking_id).generation_time.isoformat()
        expect = BookingInfo()
        ParseDict(o, expect)
        self.assertEqual(expect, actual)

        # check booking in mongodb
        booking_oid = ObjectId(resp.booking.booking_id)
        actual = self.db['bookings'].find_one({'_id': booking_oid})
        expect = utils.get_bson('data/test_book_flights_bson.json', oid=booking_oid)
        self.assertEqual(expect, actual)

    def test_get_booking_by_id(self) -> None:
        # insert test bookings
        bookings = utils.get_bson('data/test_bookings_bookings_bson.json')
        self.db['bookings'].insert_many(bookings)

        booking_id = '63bf9410831e972f68508e64'
        req = GetBookingByIdRequest(booking_id=booking_id)
        resp = self.stub.GetBookingById(req)

        actual = resp.booking
        expect = utils.get_proto('data/test_get_booking_by_id_proto.json', BookingInfo)
        self.assertEqual(expect, actual)

    def test_get_booking_by_user(self) -> None:
        # insert test bookings
        bookings = utils.get_bson('data/test_bookings_bookings_bson.json')
        self.db['bookings'].insert_many(bookings)

        customer_id = '000000000000000000000001'
        req = GetBookingByUserRequest(customer_id=customer_id)
        resp = self.stub.GetBookingByUser(req)

        actual = resp
        expect = utils.get_proto('data/test_get_booking_by_user_proto.json', GetBookingByUserRespond)
        self.assertEqual(expect, actual)

    def test_cancel_booking_by_id(self) -> None:
        # insert test bookings
        bookings = utils.get_bson('data/test_bookings_bookings_bson.json')
        self.db['bookings'].insert_many(bookings)

        booking_id = '63bf9410831e972f68508e64'
        customer_id = '000000000000000000000001'
        req = CancelBookingByIdRequest(
            booking_id=booking_id,
            customer_id=customer_id
        )
        self.stub.CancelBookingById(req)

        actual = self.db['bookings'].find_one({'_id': ObjectId(booking_id)})
        self.assertIsNone(actual)

    def test_book_flights_rest(self) -> None:
        url = f'http://{self.gateway}/api/v1/bookings/bookflights'
        customer_id = '000000000000000000000001'
        flight_id = '63bfa0ec6ced8e7900d0516e'
        req = {
            'customer_id': customer_id,
            'to_flight_id': flight_id,
            'one_way_flight': True
        }
        resp = requests.post(url, json=req)

        # check booking info
        actual = resp.json()['booking']
        expect = utils.get_json('data/test_book_flights_rest.json')
        expect['bookingId'] = actual['bookingId']
        expect['dateOfBooking'] = actual['dateOfBooking']
        self.assertEqual(expect, actual)

        # check booking in mongodb
        booking_oid = ObjectId(actual['bookingId'])
        actual = self.db['bookings'].find_one({'_id': booking_oid})
        expect = utils.get_bson('data/test_book_flights_bson.json', oid=booking_oid)
        self.assertEqual(expect, actual)

    def test_get_booking_by_id_rest(self) -> None:
        # insert test bookings
        bookings = utils.get_bson('data/test_bookings_bookings_bson.json')
        self.db['bookings'].insert_many(bookings)

        booking_id = '63bf9410831e972f68508e64'
        url = f'http://{self.gateway}/api/v1/bookings/byid/{booking_id}'
        resp = requests.get(url)

        actual = resp.json()['booking']
        expect = utils.get_json('data/test_get_booking_by_id_rest.json')
        self.assertEqual(expect, actual)

    def test_get_booking_by_user_rest(self) -> None:
        # insert test bookings
        bookings = utils.get_bson('data/test_bookings_bookings_bson.json')
        self.db['bookings'].insert_many(bookings)

        customer_id = '000000000000000000000001'
        url = f'http://{self.gateway}/api/v1/bookings/byuser/{customer_id}'
        resp = requests.get(url)

        actual = resp.json()
        expect = utils.get_json('data/test_get_booking_by_user_rest.json')
        self.assertEqual(expect, actual)

    def test_cancel_booking_by_id_rest(self) -> None:
        # insert test bookings
        bookings = utils.get_bson('data/test_bookings_bookings_bson.json')
        self.db['bookings'].insert_many(bookings)

        url = f'http://{self.gateway}/api/v1/bookings/cancelbooking'
        booking_id = '63bf9410831e972f68508e64'
        customer_id = '000000000000000000000001'
        req = {
            'booking_id': booking_id,
            'customer_id': customer_id
        }
        requests.post(url, json=req)

        actual = self.db['bookings'].find_one({'_id': ObjectId(booking_id)})
        self.assertIsNone(actual)