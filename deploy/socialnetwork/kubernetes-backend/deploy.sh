#! /bin/bash
# Create namespace
kubectl create -f namespace.yaml
# Install redis
kubectl create -f redis.yaml
# Install memcached
kubectl create -f post-storage-memcached.yaml
kubectl create -f user-memcached.yaml
kubectl create -f url-shorten-memcached.yaml