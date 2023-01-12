import utils
import requests
from proto.data_pb2 import CustomerInfo
from proto.customer_pb2_grpc import CustomerServiceStub
from proto.customer_pb2 import GetCustomerRequest, PutCustomerRequest

class TestCustomer(utils.TestAcmeair):
    stub: CustomerServiceStub

    def setUp(self) -> None:
        super().setUp('customer', CustomerServiceStub)

    def test_get_customer(self) -> None:
        customer = utils.get_bson('data/test_get_customer_bson.json')
        self.db['customer'].insert_one(customer)

        customer_id = '000000000000000000000001'
        req = GetCustomerRequest(customer_id=customer_id)
        resp = self.stub.GetCustomer(req)

        actual = resp.customer
        expect = utils.get_proto('data/test_get_customer_proto.json', CustomerInfo)
        self.assertEqual(actual, expect)

    def test_put_customer(self) -> None:
        self.db['customer'].insert_one(utils.get_bson('data/test_put_customer_old_bson.json'))

        customer_id = '000000000000000000000001'
        req = PutCustomerRequest(
            customer_id=customer_id,
            customer=utils.get_proto('data/test_put_customer_proto.json', CustomerInfo)
        )
        resp = self.stub.PutCustomer(req)

        actual = self.db['customer'].find_one()
        expect = utils.get_bson('data/test_put_customer_new_bson.json')
        self.assertEqual(actual, expect)

    def test_get_customer_rest(self) -> None:
        customer = utils.get_bson('data/test_get_customer_bson.json')
        self.db['customer'].insert_one(customer)

        customer_id = '000000000000000000000001'
        url = f'http://{self.gateway}/api/v1/customer/byid/{customer_id}'
        resp = requests.get(url)

        actual = resp.json()['customer']
        expect = utils.get_json('data/test_get_customer_rest.json')
        self.assertEqual(actual, expect)

    def test_put_customer_rest(self) -> None:
        self.db['customer'].insert_one(utils.get_bson('data/test_put_customer_old_bson.json'))

        customer_id = '000000000000000000000001'
        url = f'http://{self.gateway}/api/v1/customer/byid/{customer_id}'
        customer = utils.get_json('data/test_put_customer_rest.json')
        req = {
            'customer': customer
        }
        requests.post(url, json=req)

        actual = self.db['customer'].find_one()
        expect = utils.get_bson('data/test_put_customer_new_bson.json')
        self.assertEqual(actual, expect)