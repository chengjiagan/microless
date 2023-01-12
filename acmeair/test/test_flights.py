import utils
import requests
from datetime import datetime
from google.protobuf.timestamp_pb2 import Timestamp
from proto.flights_pb2_grpc import FlightsServiceStub
from proto.flights_pb2 import GetTripFlightsRequest, GetTripFlightsRespond, BrowseFlightsRequest, BrowseFlightsRespond, GetFlightByIdRequest
from proto.data_pb2 import FlightInfo

class TestFlights(utils.TestAcmeair):
    stub: FlightsServiceStub

    def setUp(self) -> None:
        super().setUp('flights', FlightsServiceStub)
        # insert test flights
        flights = utils.get_bson('data/test_flights_bson.json')
        self.db['flights'].insert_many(flights)

    def test_get_trip_flights(self) -> None:
        ft = Timestamp()
        ft.FromDatetime(datetime(2023, 1, 1, 12))
        rt = Timestamp()
        rt.FromDatetime(datetime(2023, 1, 2, 12))
        req = GetTripFlightsRequest(
            from_airport='CAN',
            to_airport='SHA',
            from_date=ft,
            return_date=rt,
            one_way_flight=False
        )
        resp = self.stub.GetTripFlights(req)

        actual = resp
        expect = utils.get_proto('data/test_get_trip_flights_proto.json', GetTripFlightsRespond)
        self.assertEqual(actual, expect)

    def test_browse_flights(self) -> None:
        req = BrowseFlightsRequest(
            from_airport='CAN',
            to_airport='SHA',
            one_way_flight=False
        )
        resp = self.stub.BrowseFlights(req)

        actual = resp
        expect = utils.get_proto('data/test_browse_flights_proto.json', BrowseFlightsRespond)
        self.assertEqual(actual, expect)

    def test_get_trip_flights_rest(self) -> None:
        url = f'http://{self.gateway}/api/v1/flights/queryflights'
        req = {
            'from_airport': 'CAN',
            'to_airport': 'SHA',
            'from_date': '2023-01-01T12:00:00Z',
            'return_date': '2023-01-02T12:00:00Z',
            'one_way_flight': False
        }
        resp = requests.post(url, json=req)

        actual = resp.json()
        expect = utils.get_json('data/test_get_trip_flights_rest.json')
        self.assertEqual(actual, expect)

    def test_browse_flights_rest(self) -> None:
        url = f'http://{self.gateway}/api/v1/flights/browseflights'
        req = {
            'from_airport': 'CAN',
            'to_airport': 'SHA',
            'one_way_flight': False
        }
        resp = requests.post(url, json=req)

        actual = resp.json()
        expect = utils.get_json('data/test_browse_flights_rest.json')
        self.assertEqual(actual, expect)

    def test_get_flight_by_id(self) -> None:
        flight_id = '63bfa0ec6ced8e7900d05169'
        req = GetFlightByIdRequest(flight_id=flight_id)
        resp = self.stub.GetFlightById(req)

        actual = resp.flight
        expect = utils.get_proto('data/test_get_flight_by_id_proto.json', FlightInfo)
        self.assertEqual(actual, expect)
