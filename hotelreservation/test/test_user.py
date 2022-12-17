import unittest
import utils
import hashlib
import jwt
import requests
from bson import ObjectId
from datetime import datetime
from proto.user_pb2_grpc import UserServiceStub
from proto.user_pb2 import LoginRequest, RegisterUserRequest


class TestUser(utils.TestHotelReservation):
    stub: UserServiceStub

    def setUp(self) -> None:
        super().setUp('user', UserServiceStub)

    def test_register_user(self) -> None:
        req = RegisterUserRequest(username='username', password='password')
        resp = self.stub.RegisterUser(req)

        # check user database
        actual = self.db['user'].find_one()
        pswd = hashlib.sha256(
            ('password' + actual['salt']).encode()).hexdigest()
        expect = {
            '_id': ObjectId(resp.user_id),
            'username': 'username',
            'salt': actual['salt'],
            'password': pswd,
        }
        self.assertEqual(expect, actual)

    def test_register_user_rest(self) -> None:
        url = 'http://' + self.gateway + '/api/v1/user/register'
        req = {
            'username': 'username',
            'password': 'password'
        }
        resp = requests.post(url, json=req)
        user_id = resp.json()['userId']

        # check user database
        actual = self.db['user'].find_one()
        pswd = hashlib.sha256(
            ('password' + actual['salt']).encode()).hexdigest()
        expect = {
            '_id': ObjectId(user_id),
            'username': 'username',
            'salt': actual['salt'],
            'password': pswd,
        }
        self.assertEqual(expect, actual)

    def test_login(self) -> None:
        salt = '11111111'
        username = 'user_0'
        passwd = 'password'
        hashed_pswd = hashlib.sha256((passwd + salt).encode()).hexdigest()
        user_id = '000000000000000000000001'
        user = {
            '_id': ObjectId(user_id),
            'username': username,
            'salt': salt,
            'password': hashed_pswd,
        }
        self.db['user'].insert_one(user)

        req = LoginRequest(username=username, password=passwd)
        resp = self.stub.Login(req)

        actual = jwt.decode(resp.token, self.secret, algorithms=["HS256"])

        expect = {
            'user_id': user_id,
            'username': username,
            'timestamp': actual['timestamp'],
            'ttl': '3600',
        }

        self.assertEqual(expect, actual)
        self.assertAlmostEqual(
            int(datetime.utcnow().timestamp()), int(actual['timestamp']))

    def test_login_rest(self) -> None:
        salt = '11111111'
        username = 'user_0'
        passwd = 'password'
        hashed_pswd = hashlib.sha256((passwd + salt).encode()).hexdigest()
        user_id = '000000000000000000000001'
        user = {
            '_id': ObjectId(user_id),
            'username': username,
            'salt': salt,
            'password': hashed_pswd,
        }
        self.db['user'].insert_one(user)

        url = 'http://' + self.gateway + '/api/v1/user/login'
        req = {
            'username': username,
            'password': passwd
        }
        resp = requests.post(url, json=req)
        token = resp.json()['token']

        actual = jwt.decode(token, self.secret, algorithms=["HS256"])
        expect = {
            'user_id': user_id,
            'username': username,
            'timestamp': actual['timestamp'],
            'ttl': '3600',
        }
        self.assertEqual(expect, actual)
        self.assertAlmostEqual(
            int(datetime.utcnow().timestamp()), int(actual['timestamp']))


if __name__ == '__main__':
    unittest.main()
