#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install ingress-nginx
helm install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.2.5

# Create socialnetwork namespace
kubectl create -f namespace.yaml

# Install jaeger
kubectl create -f backend/jaeger.yaml
# Install redis
kubectl create -f backend/redis.yaml
# Install mongodb
kubectl create -f backend/mongodb.yaml
# Install memcached
kubectl create -f backend/post-storage-memcached.yaml
kubectl create -f backend/user-memcached.yaml
kubectl create -f backend/url-shorten-memcached.yaml

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