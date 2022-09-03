import unittest
import utils
import hashlib
import jwt
import requests
from bson import ObjectId
from datetime import datetime
from proto.user_pb2_grpc import UserServiceStub
from proto.user_pb2 import GetUserIdRequest, LoginRequest, RegisterUserRequest


class TestUser(utils.TestSocialNetwork):
    stub: UserServiceStub

    def setUp(self) -> None:
        super().setUp('user', UserServiceStub)

    def test_register_user(self) -> None:
        req = RegisterUserRequest(
            first_name='first', last_name='last', username='username', password='password')
        resp = self.stub.RegisterUser(req)

        # check user database
        actual = self.user_db.find_one()
        pswd = hashlib.sha256(
            ('password' + actual['salt']).encode()).hexdigest()
        expect = {
            '_id': ObjectId(resp.user_id),
            'first_name': 'first',
            'last_name': 'last',
            'username': 'username',
            'salt': actual['salt'],
            'password': pswd,
        }
        self.assertEqual(expect, actual)

        # check social graph database
        actual = self.socialgraph_db.find_one()
        del actual['_id']
        expect = {
            'user_id': ObjectId(resp.user_id),
            'followers': [],
            'followees': [],
        }
        self.assertEqual(expect, actual)

        # check user timeline database
        actual = self.timeline_db.find_one()
        del actual['_id']
        expect = {
            'user_id': ObjectId(resp.user_id),
            'post_ids': [],
        }
        self.assertEqual(expect, actual)

    def test_register_user_rest(self) -> None:
        self.register_user_rest(self.rest['user'])

    def test_register_user_nginx(self) -> None:
        self.register_user_rest(self.nginx)

    def register_user_rest(self, addr: str) -> None:
        url = 'http://' + addr + '/api/v1/user/register'
        req = {
            'firstName': 'first',
            'lastName': 'last',
            'username': 'username',
            'password': 'password'
        }
        resp = requests.post(url, json=req)
        user_id = resp.json()['userId']

        # check user database
        actual = self.user_db.find_one()
        pswd = hashlib.sha256(
            ('password' + actual['salt']).encode()).hexdigest()
        expect = {
            '_id': ObjectId(user_id),
            'first_name': 'first',
            'last_name': 'last',
            'username': 'username',
            'salt': actual['salt'],
            'password': pswd,
        }
        self.assertEqual(expect, actual)

        # check social graph database
        actual = self.socialgraph_db.find_one()
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id),
            'followers': [],
            'followees': [],
        }
        self.assertEqual(expect, actual)

        # check user timeline database
        actual = self.timeline_db.find_one()
        del actual['_id']
        expect = {
            'user_id': ObjectId(user_id),
            'post_ids': [],
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
            'first_name': 'first',
            'last_name': 'last',
            'username': username,
            'salt': salt,
            'password': hashed_pswd,
        }
        self.user_db.insert_one(user)

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
        self.login_rest(self.rest['user'])

    def test_login_nginx(self) -> None:
        self.login_rest(self.nginx)

    def login_rest(self, addr: str) -> None:
        salt = '11111111'
        username = 'user_0'
        passwd = 'password'
        hashed_pswd = hashlib.sha256((passwd + salt).encode()).hexdigest()
        user_id = '000000000000000000000001'
        user = {
            '_id': ObjectId(user_id),
            'first_name': 'first',
            'last_name': 'last',
            'username': username,
            'salt': salt,
            'password': hashed_pswd,
        }
        self.user_db.insert_one(user)

        url = 'http://' + addr + '/api/v1/user/login'
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

    def test_get_user_id(self) -> None:
        salt = '11111111'
        username = 'user_0'
        passwd = 'password'
        hashed_pswd = hashlib.sha256(
            (passwd + salt).encode('utf-8')).hexdigest()
        user_id = '000000000000000000000001'
        user = {
            '_id': ObjectId(user_id),
            'first_name': 'first',
            'last_name': 'last',
            'username': username,
            'salt': salt,
            'password': hashed_pswd,
        }
        self.user_db.insert_one(user)

        req = GetUserIdRequest(username=username)
        resp = self.stub.GetUserId(req)

        self.assertEqual(user_id, resp.user_id)


if __name__ == '__main__':
    unittest.main()
