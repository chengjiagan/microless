#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install ingress-nginx
helm install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.2.5

# Create namespace
kubectl create -f namespace.yaml

# Install jaeger
kubectl create -f backend/jaeger.yaml
# Install redis
kubectl create -f backend/redis.yaml
# Install mongodb
kubectl create -f backend/mongodb.yaml
# Install memcached
kubectl create -f backend/cast-info-memcached.yaml
kubectl create -f backend/movie-info-memcached.yaml
kubectl create -f backend/plot-memcached.yaml
kubectl create -f backend/review-storage-memcached.yaml
kubectl create -f backend/user-memcached.yaml

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