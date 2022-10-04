#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install ingress-nginx
helm install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.2.5

# Install knative
kubectl apply -f https://github.com/knative/operator/releases/download/knative-v1.7.0/operator.yaml
kubectl create -f knative.yaml

# Create namespace
kubectl create -f namespace.yaml

# Install jaeger
kubectl create -f ../kubernetes/backend/jaeger.yaml
# Install redis
kubectl create -f ../kubernetes/backend/redis.yaml
# Install mongodb
kubectl create -f ../kubernetes/backend/mongodb.yaml
# Install memcached
kubectl create -f ../kubernetes/backend/cast-info-memcached.yaml
kubectl create -f ../kubernetes/backend/movie-info-memcached.yaml
kubectl create -f ../kubernetes/backend/plot-memcached.yaml
kubectl create -f ../kubernetes/backend/review-storage-memcached.yaml
kubectl create -f ../kubernetes/backend/user-memcached.yaml

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