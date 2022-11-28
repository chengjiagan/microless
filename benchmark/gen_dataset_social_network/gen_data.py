#! /usr/bin/env python

import networkx as nx
import random
import requests
import json

URL_GATEWAY = 'http://localhost:8080'
ALPHANUM = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'

# gen user
NUM_USER = 4039
URL_REGUSER = URL_GATEWAY + '/api/v1/user/register'
SAVE_USERIDS = 'user_ids.json'
pswd = 'password'
user_ids = [None for _ in range(NUM_USER)]
for i in range(NUM_USER):
    req = {
        'first_name': str(i),
        'last_name': str(i),
        'username': str(i),
        'password': pswd
    }
    resp = requests.post(URL_REGUSER, json=req)
    user_ids[i] = resp.json()['userId']
# save user_ids
with open(SAVE_USERIDS, 'w') as f:
    json.dump(user_ids, f)

# gen social graph
SNAP_DATA = 'facebook_combined.txt'
URL_FOLLOW = URL_GATEWAY + '/api/v1/socialgraph/follow'
g = nx.read_edgelist(SNAP_DATA, nodetype=int).to_directed()
for i, u in enumerate(user_ids):
    for followee in g.neighbors(i):
        req = {
            'user_id': u,
            'followee_id': user_ids[followee]
        }
        requests.post(URL_FOLLOW, json=req)

# gen post for users
URL_POST = URL_GATEWAY + '/api/v1/composepost'
for i, user_id in enumerate(user_ids):
    num_post = random.randint(1, 100)
    for _ in range(num_post):
        post_len = random.randint(1, 200)
        text = ''.join(random.choices(ALPHANUM, k=post_len))
        req = {
            'username': str(i),
            'user_id': user_id,
            'text': text,
            'media_ids': [],
            'media_types': [],
            'post_type': 0
        }
        requests.post(URL_POST, json=req)
