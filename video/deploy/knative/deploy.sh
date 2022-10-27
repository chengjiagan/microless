#! /bin/bash

# this script is just for reference, please do not run it directly

# Create namespace
kubectl create -f namespace.yaml

# Install service config
kubectl create -f config.yaml

# Install services
kubectl create -f service/video-service.yaml
kubectl create -f service/preview-service.yaml
kubectl create -f service/split-service.yaml
kubectl create -f service/transcode-service.yaml
kubectl create -f service/merge-service.yaml

# Run test
kubectl create -f test.yaml