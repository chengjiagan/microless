#! /bin/bash

# this script is just for reference, please do not run it directly

# Create namespace
kubectl create -f namespace.yaml

# Install service config
kubectl create -f config.yaml

# Install services
kubectl create -f service/mapreduce-service.yaml
kubectl create -f service/file-service.yaml
kubectl create -f service/mapper-service.yaml
kubectl create -f service/reducer-service.yaml

# Wait for services
sleep 30

# Run test
kubectl create -f test.yaml