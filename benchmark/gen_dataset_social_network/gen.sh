#! /bin/bash

python ./gen_data.py
mongodump -d='socialnetwork-dev' --gzip --archive='socialnetwork.mongo.gz'