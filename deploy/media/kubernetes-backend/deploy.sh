#! /bin/bash

# Create namespace
kubectl create -f namespace.yaml
# Install redis
kubectl create -f redis.yaml
# Install memcached
kubectl create -f cast-info-memcached.yaml
kubectl create -f movie-info-memcached.yaml
kubectl create -f plot-memcached.yaml
kubectl create -f review-storage-memcached.yaml
kubectl create -f user-memcached.yaml