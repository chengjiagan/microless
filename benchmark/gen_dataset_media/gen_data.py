#! /usr/bin/env python

import random
import requests
import json
import pymongo

URL_GATEWAY = 'http://localhost:8080'
URL_REGUSER = URL_GATEWAY + '/api/v1/user/register'
URL_REVIEW = URL_GATEWAY + '/api/v1/composereview'

URL_MONGO = 'mongodb://localhost:27017'
MONGO_DB = 'media-dev'
COL_CAST = 'cast-info'
COL_PLOT = 'plot'
COL_MOVIE = 'movie-info'
COL_REVIEW = 'movie-review'

DATASET_CAST = 'casts.json'
DATASET_MOVIE = 'movies.json'
NUM_USER = 2000
SAVE_MOVIEIDS = 'movie_ids.json'
SAVE_USERIDS = 'user_ids.json'

ALPHANUM = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'

# connect to mongodb
client = pymongo.MongoClient(URL_MONGO)
db = client.get_database(MONGO_DB)

# gen cast infos
cast2castid = {}
cast_col = db.get_collection(COL_CAST)
with open(DATASET_CAST, 'r') as f:
    casts = json.load(f)
for c in casts:
    info = {
        'name': c['name'],
        'gender': c['gender'] == 2,
        'intro': c['biography']
    }
    res = cast_col.insert_one(info)
    cast2castid[c['id']] = res.inserted_id

# gen movie infos and plots
movieids = []
plot_col = db.get_collection(COL_PLOT)
movie_col = db.get_collection(COL_MOVIE)
review_col = db.get_collection(COL_REVIEW)
with open(DATASET_MOVIE, 'r') as f:
    movies = json.load(f)
for m in movies:
    plot = {
        'plot': m['overview']
    }
    res = plot_col.insert_one(plot)

    casts = []
    casts = [
        {
            'cast_id': c['cast_id'],
            'character': c['character'],
            'cast_info_id': cast2castid[c['id']]
        }
        for c in m['cast'] if c['id'] in cast2castid
    ]

    movie = {
        'title': m['title'],
        'casts': casts,
        'plot_id': res.inserted_id,
        'thumbnail_ids': [],
        'photo_ids': [],
        'video_ids': [],
        'avg_rating': 0,
        'num_rating': 0
    }
    res = movie_col.insert_one(movie)
    movieids.append(str(res.inserted_id))

    review = {
        'movie_id': res.inserted_id,
        'review_ids': []
    }
    review_col.insert_one(review)
with open(SAVE_MOVIEIDS, 'w') as f:
    json.dump(movieids, f)

# gen users
userids = [None for _ in range(NUM_USER)]
for i in range(NUM_USER):
    req = {
        'first_name': str(i),
        'last_name': str(i),
        'username': str(i),
        'password': 'password'
    }
    res = requests.post(URL_REGUSER, json=req)
    userids[i] = res.json()['userId']
with open(SAVE_USERIDS, 'w') as f:
    json.dump(userids, f)

# gen reviews
for u in userids:
    num_review = random.randint(1, 100)
    movies = random.choices(movieids, k=num_review)
    for m in movies:
        rate = random.randint(1, 10)
        post_len = random.randint(1, 200)
        text = ''.join(random.choices(ALPHANUM, k=post_len))
        req = {
            'movie_id': m,
            'user_id': u,
            'text': text,
            'rating': rate
        }
        requests.post(URL_REVIEW, json=req)
