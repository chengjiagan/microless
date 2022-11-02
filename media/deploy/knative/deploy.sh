#! /bin/bash

# this script is just for reference, please do not run it directly

# Create namespace
kubectl create -f namespace.yaml

# Install service config
kubectl create -f config.yaml

# Install services
kubectl create -f service/cast-info-service.yaml
kubectl create -f service/movie-info-service.yaml
kubectl create -f service/plot-service.yaml
kubectl create -f service/review-storage-service.yaml
kubectl create -f service/movie-review-service.yaml
kubectl create -f service/user-review-service.yaml
kubectl create -f service/rating-service.yaml
kubectl create -f service/compose-review-service.yaml
kubectl create -f service/user-service.yaml
kubectl create -f service/page-service.yaml

# Install restful gateway
kubectl create -f restful-gateway.yaml

# Wait for services
sleep 30

# Run test
kubectl create -f test.yaml