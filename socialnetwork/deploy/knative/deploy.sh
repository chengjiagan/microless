#! /bin/bash

# this script is just for reference, please do not run it directly

# Create namespace
kubectl create -f namespace.yaml

# Install service config
kubectl create -f config.yaml

# Install services
kubectl create -f service/post-storage-service.yaml
kubectl create -f service/user-timeline-service.yaml
kubectl create -f service/user-service.yaml
kubectl create -f service/social-graph-service.yaml
kubectl create -f service/home-timeline-service.yaml
kubectl create -f service/media-service.yaml
kubectl create -f service/url-shorten-service.yaml
kubectl create -f service/user-mention-service.yaml
kubectl create -f service/text-service.yaml
kubectl create -f service/compose-post-service.yaml

# Install restful gateway
kubectl create -f restful-gateway.yaml

# Wait for services
sleep 30

# Run test
kubectl create -f test.yaml
